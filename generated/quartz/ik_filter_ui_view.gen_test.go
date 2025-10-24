// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz_test

import (
	"github.com/tmc/appledocs/generated/quartz"
)

// Suppress unused import errors
var _ = quartz.NewIKFilterUIView

// ExampleIKFilterUIView_Filter demonstrates using Filter on a IKFilterUIView instance.
// Returns the Core Image filter associated with the view.
func ExampleIKFilterUIView_Filter() {
	obj := quartz.NewIKFilterUIView()
	_ = obj.Filter()
	// Output:
	}

// ExampleIKFilterUIView_ObjectController demonstrates using ObjectController on a IKFilterUIView instance.
// Returns the object controller for the bindings between the filter and its view.
func ExampleIKFilterUIView_ObjectController() {
	obj := quartz.NewIKFilterUIView()
	_ = obj.ObjectController()
	// Output:
	}

