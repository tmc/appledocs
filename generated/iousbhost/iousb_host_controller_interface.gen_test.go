// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost_test

import (
	"github.com/tmc/appledocs/generated/iousbhost"
)

// Suppress unused import errors
var _ = iousbhost.NewUSBHostControllerInterface

// ExampleUSBHostControllerInterface_Destroy demonstrates using Destroy on a USBHostControllerInterface instance.
func ExampleUSBHostControllerInterface_Destroy() {
	obj := iousbhost.NewUSBHostControllerInterface()
	obj.Destroy()
	// Output:
	}

