// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURLDownload

// ExampleURLDownload_Cancel demonstrates using Cancel on a URLDownload instance.
// Cancels the receiver’s download and deletes the downloaded file.
func ExampleURLDownload_Cancel() {
	obj := foundation.NewURLDownload()
	obj.Cancel()
	// Output:
	}

