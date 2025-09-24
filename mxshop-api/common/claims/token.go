package claims

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type TokenManager struct {
	rdb   *redis.Client
	ctx   context.Context
	ttl   time.Duration
	limit int64 // 最大 token 数量，1=单点登录，>1=多终端限制
}

func NewTokenManager(rdb *redis.Client, ttl time.Duration, limit int64) *TokenManager {
	return &TokenManager{
		rdb:   rdb,
		ctx:   context.Background(),
		ttl:   ttl,
		limit: limit,
	}
}

func (tm *TokenManager) key(userID string) string {
	return fmt.Sprintf("user:%s:tokens", userID)
}

// SaveToken 保存 token
func (tm *TokenManager) SaveToken(userID string, token string) error {
	key := tm.key(userID)
	expireAt := time.Now().Add(tm.ttl).Unix()

	// 单点登录模式，清理旧 token
	if tm.limit == 1 {
		if err := tm.rdb.Del(tm.ctx, key).Err(); err != nil {
			return err
		}
	}

	// 添加新 token
	if err := tm.rdb.ZAdd(tm.ctx, key, &redis.Z{
		Score:  float64(expireAt),
		Member: token,
	}).Err(); err != nil {
		return err
	}

	// 设置集合的过期时间兜底
	if err := tm.rdb.Expire(tm.ctx, key, tm.ttl).Err(); err != nil {
		return err
	}

	// 多终端限制，删除最旧的
	if tm.limit > 1 {
		count, _ := tm.rdb.ZCard(tm.ctx, key).Result()
		if count > int64(tm.limit) {
			// 删除过多的旧 token（score 最小的）
			if err := tm.rdb.ZRemRangeByRank(tm.ctx, key, 0, int64(count-tm.limit-1)).Err(); err != nil {
				return err
			}
		}
	}

	return nil
}

// ValidateToken 校验 token 是否有效
func (tm *TokenManager) ValidateToken(userID string, token string) (bool, error) {
	key := tm.key(userID)

	expire, err := tm.rdb.ZScore(tm.ctx, key, token).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil // token 不存在
	}
	if err != nil {
		return false, err
	}

	if int64(expire) < time.Now().Unix() {
		// 已过期，顺便删除
		_ = tm.rdb.ZRem(tm.ctx, key, token).Err()
		return false, nil
	}

	return true, nil
}

// RevokeToken 删除指定 token
func (tm *TokenManager) RevokeToken(userID string, token string) error {
	key := tm.key(userID)
	return tm.rdb.ZRem(tm.ctx, key, token).Err()
}

// RevokeAll 清除该用户的所有 token
func (tm *TokenManager) RevokeAll(userID string) error {
	key := tm.key(userID)
	return tm.rdb.Del(tm.ctx, key).Err()
}

// BlacklistToken 将 Token 加入黑名单
func (tm *TokenManager) BlacklistToken(token string, expiry time.Time) error {
	ctx := context.Background()
	key := fmt.Sprintf("blacklist:%s", token)

	// 计算剩余过期时间
	remaining := time.Until(expiry)
	if remaining <= 0 {
		return nil // Token 已过期，无需加入黑名单
	}
	// 尝试使用 Redis
	err := tm.rdb.Set(ctx, key, "1", remaining).Err()
	if err != nil {
		// Redis 失败，使用内存存储
		fmt.Printf("Redis 黑名单存储失败，使用内存: %v\n", err)
	}
	return nil

}

// IsTokenBlacklisted 检查 Token 是否在黑名单中
func (tm *TokenManager) IsTokenBlacklisted(token string) bool {
	ctx := context.Background()
	key := fmt.Sprintf("blacklist:%s", token)

	// 检查 Redis 中是否存在
	_, err := tm.rdb.Get(ctx, key).Result()
	if err == nil {
		return true // 在黑名单中
	}
	return false
}
