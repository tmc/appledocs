// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory_test

import (
	"os"
	"runtime"
	"testing"
)

// TestMain locks the main test goroutine to the OS thread before running tests.
// This is required for OpenDirectory operations which must run on the main thread.
func TestMain(m *testing.M) {
	runtime.LockOSThread()
	os.Exit(m.Run())
}



