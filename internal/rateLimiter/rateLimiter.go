package rateLimiter

import (
	"sync"
	"time"
)

type UserMessageTracker struct {
	MessageCount    int
	LastMessageTime time.Time
}

type RateLimiter struct {
	UserMessageMap map[int]*UserMessageTracker
	Mutex          sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		UserMessageMap: make(map[int]*UserMessageTracker),
	}
}

func (rl *RateLimiter) CheckAndIncrement(userID int) bool {
	rl.Mutex.Lock()
	defer rl.Mutex.Unlock()

	tracker, exists := rl.UserMessageMap[userID]
	if !exists {
		tracker = &UserMessageTracker{}
		rl.UserMessageMap[userID] = tracker
	}

	if time.Since(tracker.LastMessageTime) < 10*time.Second && tracker.MessageCount >= 1 {
		return false
	}

	tracker.MessageCount++
	tracker.LastMessageTime = time.Now()

	return true
}
