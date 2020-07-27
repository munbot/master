// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core runtime.
package core

import (
	"context"
	"errors"
	"time"

	"github.com/munbot/master/utils/lock"
	"github.com/munbot/master/utils/uuid"
)

type key int

const lockKey key = 0

func lockContext(rt *Runtime) (context.Context, error) {
	var err error
	rt.uuid, err = tryLock(rt.mu)
	if err != nil {
		return nil, err
	}
	return context.WithValue(rt.ctx, lockKey, rt.uuid), nil
}

func tryLock(mu *lock.Locker) (string, error) {
	if mu.TryLockTimeout(time.Second) {
		return uuid.Rand(), nil
	}
	return "", errors.New("core lock timeout")
}

type Runtime struct {
	ctx context.Context
	mu *lock.Locker
	uuid string
}

func NewRuntime(ctx context.Context) *Runtime {
	return &Runtime{ctx: ctx, mu: lock.New()}
}

func (rt *Runtime) Lock() error {
	var err error
	rt.ctx, err = lockContext(rt)
	return err
}
