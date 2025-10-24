// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore_test

import (
	"github.com/tmc/appledocs/generated/quartzcore"
)

// Suppress unused import errors
var _ = quartzcore.NewLayer

// ExampleNewLayer demonstrates how to create a Layer instance.
// Returns an initialized   object.
func ExampleNewLayer() {
	_ = quartzcore.NewLayer()
	// Output:
}

// ExampleNewLayerWithRemoteClientId demonstrates how to create a Layer instance using NewLayerWithRemoteClientId.
// Initializes a layer with a remote client ID.
func ExampleNewLayerWithRemoteClientId() {
	_ = quartzcore.NewLayerWithRemoteClientId(
		0, // client_id uint32
	)
	// Output:
}
