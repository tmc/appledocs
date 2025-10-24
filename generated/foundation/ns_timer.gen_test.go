// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewTimer

// ExampleTimer_Fire demonstrates using Fire on a Timer instance.
// Causes the timer’s message to be sent to its target.
func ExampleTimer_Fire() {
	obj := foundation.NewTimer()
	obj.Fire()
	// Output:
	}

// ExampleTimer_Invalidate demonstrates using Invalidate on a Timer instance.
// Stops the timer from ever firing again and requests its removal from its run loop.
func ExampleTimer_Invalidate() {
	obj := foundation.NewTimer()
	obj.Invalidate()
	// Output:
	}

