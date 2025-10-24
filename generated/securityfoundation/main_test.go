// Code generated from Apple documentation for SecurityFoundation. DO NOT EDIT.

package securityfoundation_test

import (
	"os"
	"runtime"
	"testing"
)

// TestMain locks the main test goroutine to the OS thread before running tests.
// This is required for SecurityFoundation operations which must run on the main thread.
func TestMain(m *testing.M) {
	runtime.LockOSThread()
	os.Exit(m.Run())
}



