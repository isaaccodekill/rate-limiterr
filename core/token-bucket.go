package core

import (
	"fmt"
	"time"
)

// implement token bucket algorithm
// https://en.wikipedia.org/wiki/Token_bucket

// TokenBucket a struct to hold the number of tokens and the refill rate
// in the future you could use different refill rates for select ip addresses≤

type TokenBucket struct {
	tokens              int
	capacity            int
	refillRatePerSecond int
}

type TokenBucketLimiter struct {
	counterStore map[string]TokenBucket
}

func NewTokenBucketLimiter() TokenBucketLimiter {
	TBL := TokenBucketLimiter{counterStore: make(map[string]TokenBucket)}
	// start a goroutine to refill the tokens
	go func() {
		// call the token filler every second
		for {
			TBL.TokenFiller()
			time.Sleep(time.Second)
		}
	}()
	// return the token bucket limiter
	return TBL
}

// a map of counters for each ip address

func (TBL TokenBucketLimiter) AllowRequest(ipAddr string) bool {

	// when a request comes in, check if the ip address is in the counterStore
	// if it is, check if the counter is greater than 0
	// if it is, decrement the counter and allow the request
	// if it is not, return false
	// if it is not, add the ip address to the counterStore with a counter of 1

	fmt.Println("AllowRequest called", ipAddr, TBL.counterStore)

	store := TBL.counterStore

	val, ok := store[ipAddr]

	if ok && val.tokens < 1 {
		return false
	}

	if !ok {
		store[ipAddr] = TokenBucket{tokens: 10, capacity: 10, refillRatePerSecond: 1}
	}

	val = store[ipAddr]
	val.tokens--

	store[ipAddr] = val
	fmt.Println("finish call", ipAddr, TBL.counterStore)

	return true
}

func (TBL TokenBucketLimiter) TokenFiller() {
	// loop through the counterStore
	// for each ip address, increment the counter by the refill rate
	// if counter is full, do not increment

	store := TBL.counterStore

	for key, val := range store {
		if val.tokens < val.capacity {
			val.tokens += val.refillRatePerSecond
		}
		store[key] = val
	}

	fmt.Println("TokenFiller called")
	fmt.Println(store)
}
