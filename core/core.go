// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core runtime.
package core

import (
	"context"
	"errors"
	"time"

	"github.com/munbot/master/config"
	"github.com/munbot/master/utils/lock"
	"github.com/munbot/master/utils/uuid"
)

type key int
const lockKey key = 0

var _ Runtime = &Core{}

type Core struct {
	ctx      context.Context
	mu       *lock.Locker
	uuid     string
	locked   string
	cfg      *config.Config
	cfgFlags *config.Flags
	flags    *Flags
}

func NewRuntime(ctx context.Context) Runtime {
	return &Core{ctx: ctx, mu: lock.New(), uuid: uuid.Rand()}
}

func (rt *Core) String() string {
	return "Core:" + rt.uuid
}

func (rt *Core) UUID() string {
	return rt.uuid
}

func (rt *Core) Context() context.Context {
	return rt.ctx
}

func (rt *Core) WithContext(ctx context.Context) (context.Context, error) {
	lockId := rt.fromContext(ctx, lockKey)
	if lockId == "" {
		if rt.locked != "" {
			// we are locked but the new context is not, so copy our lock key to
			// the new context
			ctx = rt.ctxCopy(ctx, lockKey)
		}
	} else {
		if rt.locked == "" {
			// the new context is locked but we are not, try to lock ourselves
			if err := rt.Lock(); err != nil {
				return nil, err
			}
		}
		// use new context lock key for future validations
		rt.locked = lockId
	}
	rt.ctx = ctx
	return rt.ctx, nil
}

func (rt *Core) ctxCopy(dst context.Context, k key) context.Context {
	if v, ok := rt.ctx.Value(k).(string); ok {
		dst = context.WithValue(dst, k, v)
	}
	return dst
}

func (rt *Core) fromContext(ctx context.Context, k key) string {
	s, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return s
}

func (rt *Core) Lock() error {
	if err := rt.tryLock(); err != nil {
		return err
	}
	rt.ctx = context.WithValue(rt.ctx, lockKey, rt.locked)
	return nil
}

func (rt *Core) tryLock() error {
	if rt.mu.TryLockTimeout(300*time.Millisecond) {
		rt.locked = uuid.Rand()
		return nil
	}
	return errors.New("core lock timeout")
}

func (rt *Core) Configure(cfg *config.Config, cfl *config.Flags, f *Flags) error {
	if rt.locked == "" {
		return errors.New("core runtime not locked")
	}
	select {
	case <-rt.ctx.Done():
		return rt.ctx.Err()
	default:
		rt.cfg = cfg
		rt.cfgFlags = cfl
		rt.flags = f
	}
	return nil
}
