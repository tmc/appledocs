// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit_test

import (
	"os"
	"runtime"
	"testing"
)

// TestMain locks the main test goroutine to the OS thread before running tests.
// This is required for CryptoTokenKit operations which must run on the main thread.
func TestMain(m *testing.M) {
	runtime.LockOSThread()
	os.Exit(m.Run())
}



