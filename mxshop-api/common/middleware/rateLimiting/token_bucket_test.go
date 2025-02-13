package ratelimit

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// TestNewTokenBucket 测试令牌桶的初始状态
func TestNewTokenBucket(t *testing.T) {
	tb := NewTokenBucket(10, 20)
	defer tb.Stop()

	if tb.tokens != 20 {
		t.Errorf("Expected initial tokens to be 20, got %d", tb.tokens)
	}
	if tb.rate != 10 {
		t.Errorf("Expected rate to be 10, got %d", tb.rate)
	}
	if tb.capacity != 20 {
		t.Errorf("Expected capacity to be 20, got %d", tb.capacity)
	}
}

// TestAllow 测试令牌消耗和补充
func TestAllow(t *testing.T) {
	tb := NewTokenBucket(1, 2) // 速率 1 令牌/秒，容量 2
	defer tb.Stop()

	// 初始状态，应该允许 2 次请求
	if !tb.Allow() {
		t.Error("Expected first request to be allowed")
	}
	if !tb.Allow() {
		t.Error("Expected second request to be allowed")
	}
	if tb.Allow() {
		t.Error("Expected third request to be denied")
	}

	// 等待 1.2 秒，确保令牌补充
	time.Sleep(1200 * time.Millisecond)

	// 重新获取令牌
	if !tb.Allow() {
		t.Error("Expected request after refill to be allowed")
	}
}

// TestRateLimiting 测试限流功能
func TestRateLimiting(t *testing.T) {
	tb := NewTokenBucket(2, 5) // 速率 2 令牌/秒，容量 5
	defer tb.Stop()

	allowed := 0
	for i := 0; i < 10; i++ {
		if tb.Allow() {
			allowed++
		}
		time.Sleep(200 * time.Millisecond) // 模拟请求间隔
	}

	// 在 2 秒内，最多允许 4 个请求（初始 5 个令牌，但每秒补充 2 个）
	if allowed < 4 || allowed > 5 {
		t.Errorf("Expected allowed requests between 4 and 5, got %d", allowed)
	}
}

// TestSetRateAndCapacity 测试动态调整速率和容量
func TestSetRateAndCapacity(t *testing.T) {
	tb := NewTokenBucket(1, 2) // 速率 1 令牌/秒，容量 2
	defer tb.Stop()

	// 初始状态
	if !tb.Allow() || !tb.Allow() {
		t.Error("Expected initial requests to be allowed")
	}
	if tb.Allow() {
		t.Error("Expected request to be denied after tokens exhausted")
	}

	// 调整速率和容量
	tb.SetRate(2)
	tb.SetCapacity(3)

	// 等待 1 秒，补充 2 个令牌
	time.Sleep(1100 * time.Millisecond)
	if !tb.Allow() || !tb.Allow() || !tb.Allow() {
		t.Error("Expected requests to be allowed after rate and capacity adjustment")
	}
	if tb.Allow() {
		t.Error("Expected request to be denied after tokens exhausted")
	}
}

// TestStop 测试停止功能
func TestStop(t *testing.T) {
	tb := NewTokenBucket(1, 2)

	// 先消耗掉所有令牌
	tb.Allow()
	tb.Allow()

	tb.Stop()

	// 等待 1.5 秒，确保令牌不会再补充
	time.Sleep(1500 * time.Millisecond)

	// `Allow` 应该始终返回 false
	if tb.Allow() {
		t.Error("Expected request to be denied after stop")
	}
}

// TestConcurrentAccess 测试并发访问
func TestConcurrentAccess(t *testing.T) {
	tb := NewTokenBucket(10, 20) // 速率 10 令牌/秒，容量 20
	defer tb.Stop()

	var wg sync.WaitGroup
	allowed := int64(0)

	// 启动 50 个并发请求
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if tb.Allow() {
				atomic.AddInt64(&allowed, 1)
			}
		}()
	}

	wg.Wait()

	// 初始令牌为 20，最多允许 20 个请求
	if allowed != 20 {
		t.Errorf("Expected 20 allowed requests, got %d", allowed)
	}
}
