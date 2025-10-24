// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost_test

import (
	"github.com/tmc/appledocs/generated/iousbhost"
)

// Suppress unused import errors
var _ = iousbhost.NewUSBHostObject

// ExampleUSBHostObject_Destroy demonstrates using Destroy on a USBHostObject instance.
// Removes underlying allocations and connections from the USB host object.
func ExampleUSBHostObject_Destroy() {
	obj := iousbhost.NewUSBHostObject()
	obj.Destroy()
	// Output:
	}

