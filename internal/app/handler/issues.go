// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package handler

import (
	"fmt"
)

// Observer interface
type Observer interface {
	Trigger(event interface{})
}

// NewIssue handler
type NewIssue struct {
}

// Trigger triggers when new issue got created
func (o NewIssue) Trigger(event interface{}) {
	fmt.Printf("Payload %s\n", event.(string))
}

// EditIssue handler
type EditIssue struct {
}

// Trigger triggers when new issue got updated
func (o EditIssue) Trigger(event interface{}) {
	fmt.Printf("Payload %s\n", event.(string))
}
