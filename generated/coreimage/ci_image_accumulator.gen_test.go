// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewImageAccumulator

// ExampleImageAccumulator_Clear demonstrates using Clear on a ImageAccumulator instance.
// Resets the accumulator, discarding any pending updates and the current content.
func ExampleImageAccumulator_Clear() {
	obj := coreimage.NewImageAccumulator()
	obj.Clear()
	// Output:
	}

// ExampleImageAccumulator_Image demonstrates using Image on a ImageAccumulator instance.
// Returns the current contents of the image accumulator.
//
// Note: This example is not executed because Image crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleImageAccumulator_Image() {
	obj := coreimage.NewImageAccumulator()
	_ = obj.Image()
	}

