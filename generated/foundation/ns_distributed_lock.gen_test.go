// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewDistributedLock

// ExampleDistributedLock_BreakLock demonstrates using BreakLock on a DistributedLock instance.
// Forces the lock to be relinquished.
func ExampleDistributedLock_BreakLock() {
	obj := foundation.NewDistributedLock()
	obj.BreakLock()
	// Output:
	}

// ExampleDistributedLock_TryLock demonstrates using TryLock on a DistributedLock instance.
// Attempts to acquire the receiver and immediately returns a Boolean value that indicates whether the attempt was successful.
func ExampleDistributedLock_TryLock() {
	obj := foundation.NewDistributedLock()
	_ = obj.TryLock()
	// Output:
	}

// ExampleDistributedLock_Unlock demonstrates using Unlock on a DistributedLock instance.
// Relinquishes the receiver.
func ExampleDistributedLock_Unlock() {
	obj := foundation.NewDistributedLock()
	obj.Unlock()
	// Output:
	}

