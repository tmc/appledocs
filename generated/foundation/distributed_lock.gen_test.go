// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewDistributedLock

// ExampleNewDistributedLockWithPath demonstrates how to create a DistributedLock instance using NewDistributedLockWithPath.
// Initializes an   object to use as the lock the file-system entry specified by a given path.
func ExampleNewDistributedLockWithPath() {
	_ = foundation.NewDistributedLockWithPath(
		"/tmp/test", // path string
	)
	// Output:
}
