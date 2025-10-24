// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewFilterGenerator

// ExampleFilterGenerator_Filter demonstrates using Filter on a FilterGenerator instance.
// Creates a filter object based on the filter chain.
func ExampleFilterGenerator_Filter() {
	obj := coreimage.NewFilterGenerator()
	_ = obj.Filter()
	// Output:
	}

