// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"sync"

	"github.com/clivern/pushover-actions/internal/app/handler"
)

// Observable interface
type Observable interface {
	Add(observer handler.Observer)
	Notify(event interface{})
	Remove(event interface{})
}

// WatchTower struct
type WatchTower struct {
	observer sync.Map
}

// Add adds a new handler
func (wt *WatchTower) Add(observer handler.Observer) {
	wt.observer.Store(observer, struct{}{})
}

// Remove removes a handler
func (wt *WatchTower) Remove(observer handler.Observer) {
	wt.observer.Delete(observer)
}

// Notify notify all handlers
func (wt *WatchTower) Notify(event interface{}) {
	wt.observer.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}

		key.(handler.Observer).Trigger(event)
		return true
	})
}
