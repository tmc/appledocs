// Code generated from Apple documentation for IOSurface. DO NOT EDIT.

package iosurface_test

import (
	"github.com/tmc/appledocs/generated/iosurface"
)

// Suppress unused import errors
var _ = iosurface.NewSurface

// ExampleSurface_AllAttachments demonstrates using AllAttachments on a Surface instance.
func ExampleSurface_AllAttachments() {
	obj := iosurface.NewSurface()
	_ = obj.AllAttachments()
	// Output:
	}

// ExampleSurface_DecrementUseCount demonstrates using DecrementUseCount on a Surface instance.
func ExampleSurface_DecrementUseCount() {
	obj := iosurface.NewSurface()
	obj.DecrementUseCount()
	// Output:
	}

// ExampleSurface_IncrementUseCount demonstrates using IncrementUseCount on a Surface instance.
func ExampleSurface_IncrementUseCount() {
	obj := iosurface.NewSurface()
	obj.IncrementUseCount()
	// Output:
	}

// ExampleSurface_RemoveAllAttachments demonstrates using RemoveAllAttachments on a Surface instance.
func ExampleSurface_RemoveAllAttachments() {
	obj := iosurface.NewSurface()
	obj.RemoveAllAttachments()
	// Output:
	}




