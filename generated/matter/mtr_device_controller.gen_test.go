// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter_test

import (
	"github.com/tmc/appledocs/generated/matter"
)

// Suppress unused import errors
var _ = matter.NewMTRDeviceController

// ExampleMTRDeviceController_Resume demonstrates using Resume on a MTRDeviceController instance.
// Resume the controller.  This has no effect if the controller is not   suspended.
func ExampleMTRDeviceController_Resume() {
	obj := matter.NewMTRDeviceController()
	obj.Resume()
	// Output:
	}

// ExampleMTRDeviceController_Shutdown demonstrates using Shutdown on a MTRDeviceController instance.
func ExampleMTRDeviceController_Shutdown() {
	obj := matter.NewMTRDeviceController()
	obj.Shutdown()
	// Output:
	}

// ExampleMTRDeviceController_StopBrowseForCommissionables demonstrates using StopBrowseForCommissionables on a MTRDeviceController instance.
func ExampleMTRDeviceController_StopBrowseForCommissionables() {
	obj := matter.NewMTRDeviceController()
	_ = obj.StopBrowseForCommissionables()
	// Output:
	}

// ExampleMTRDeviceController_Suspend demonstrates using Suspend on a MTRDeviceController instance.
// Suspend the controller.  This will attempt to stop all network traffic associated   with the controller.  The controller will remain suspended until it is   resumed.
func ExampleMTRDeviceController_Suspend() {
	obj := matter.NewMTRDeviceController()
	obj.Suspend()
	// Output:
	}

