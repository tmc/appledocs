// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter_test

import (
	"github.com/tmc/appledocs/generated/matter"
)

// Suppress unused import errors
var _ = matter.NewMTRDevice

// ExampleMTRDevice_DescriptorClusters demonstrates using DescriptorClusters on a MTRDevice instance.
// Read all known attributes from descriptor clusters on all known endpoints.
func ExampleMTRDevice_DescriptorClusters() {
	obj := matter.NewMTRDevice()
	_ = obj.DescriptorClusters()
	// Output:
	}

