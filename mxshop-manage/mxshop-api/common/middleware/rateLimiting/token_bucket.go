package ratelimit

import (
	"sync"
	"time"
)

// https://mp.weixin.qq.com/s/xa4HzJnuErYsNr4vcYIbYA

// TokenBucket 实现了基于令牌桶的限流算法。
// 令牌桶算法允许突发流量，但总体速率受限于令牌生成速率。
type TokenBucket struct {
	rate      int64         // 生成令牌的速率，单位是令牌/秒
	capacity  int64         // 桶容量
	tokens    int64         // 当前令牌数量
	lastToken time.Time     // 上次生成令牌的时间
	mutex     sync.Mutex    //用于并发控制
	ticker    *time.Ticker  // 定时器，用于定期生成令牌
	stopChan  chan struct{} // 用于停止定时器
	stopped   bool          // 标记是否已停止
}

// NewTokenBucket 创建一个新的令牌桶实例。
// rate: 令牌生成速率（令牌/秒）
// capacity: 桶的容量（最大令牌数）
func NewTokenBucket(rate, capacity int) *TokenBucket {
	tb := &TokenBucket{
		rate:      int64(rate),
		capacity:  int64(capacity),
		tokens:    int64(capacity), // 初始时，桶里有满令牌
		lastToken: time.Now(),
		stopChan:  make(chan struct{}),
		stopped:   false,
	}
	//使用定时器定期补充令牌
	tb.ticker = time.NewTicker(10 * time.Millisecond)
	go func() {
		for {
			select {
			case <-tb.ticker.C:
				tb.refill()
			case <-tb.stopChan:
				return
			}
		}
	}()
	return tb
}

// Allow 检查是否允许通过一个请求。
// 如果桶中有令牌，则消耗一个令牌并返回 true；否则返回 false。
func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	// 补充令牌
	tb.refill()

	// 检查是否有足够的令牌
	if tb.tokens > 0 {
		// 有令牌可以消耗
		tb.tokens--
		return true
	}
	// 没有令牌可用，限制请求
	return false
}

// Stop 停止令牌桶的定时器，释放资源
func (tb *TokenBucket) Stop() {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	if !tb.stopped {
		close(tb.stopChan)
		tb.ticker.Stop()
		tb.stopped = true
	}

}

// refill 根据时间差补充令牌。
// 该方法在定时器触发时调用，或者在 Allow 方法中调用。
func (tb *TokenBucket) refill() {
	// 计算过去时间段内生成的令牌数
	now := time.Now()
	elapsed := now.Sub(tb.lastToken).Seconds()
	tb.lastToken = now

	// 按速率生成令牌
	newTokens := int64(elapsed * float64(tb.rate))
	if newTokens > 0 {
		// 补充令牌
		tb.tokens += newTokens
		// 保证令牌数量不会超过桶的容量
		if tb.tokens > tb.capacity {
			// 确保令牌数量不超过桶的容量
			tb.tokens = tb.capacity
		}
	}
}

// SetRate 动态调整令牌生成速率。
func (tb *TokenBucket) SetRate(rate int64) {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	tb.rate = rate
}

// SetCapacity 动态调整桶的容量。
func (tb *TokenBucket) SetCapacity(capacity int64) {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	tb.capacity = capacity
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
}
