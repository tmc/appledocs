// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui_test

import (
	"github.com/tmc/appledocs/generated/iobluetoothui"
)

// Suppress unused import errors
var _ = iobluetoothui.NewBluetoothServiceBrowserController

// ExampleNewBluetoothServiceBrowserController demonstrates how to create a BluetoothServiceBrowserController instance using NewBluetoothServiceBrowserController.
// Allocator work Bluetooth Service Browser window controller.
func ExampleNewBluetoothServiceBrowserController() {
	_ = iobluetoothui.NewBluetoothServiceBrowserController(
		iobluetoothui.BluetoothServiceBrowserControllerOptions /* typedef */{}, // inOptions BluetoothServiceBrowserControllerOptions /* typedef */
	)
	// Output:
}
// ExampleBluetoothServiceBrowserController_ClearAllowedUUIDs demonstrates using ClearAllowedUUIDs on a BluetoothServiceBrowserController instance.
// Resets the controller back to the default state where it will accept any device the user selects.
func ExampleBluetoothServiceBrowserController_ClearAllowedUUIDs() {
	obj := iobluetoothui.NewBluetoothServiceBrowserController()
	obj.ClearAllowedUUIDs()
	// Output:
	}

// ExampleBluetoothServiceBrowserController_GetDescriptionText demonstrates using GetDescriptionText on a BluetoothServiceBrowserController instance.
// Returns the description text that appears in the device selector panel.
func ExampleBluetoothServiceBrowserController_GetDescriptionText() {
	obj := iobluetoothui.NewBluetoothServiceBrowserController()
	_ = obj.GetDescriptionText()
	// Output:
	}

// ExampleBluetoothServiceBrowserController_GetOptions demonstrates using GetOptions on a BluetoothServiceBrowserController instance.
// Returns the option bits that control the panel’s behavior.
func ExampleBluetoothServiceBrowserController_GetOptions() {
	obj := iobluetoothui.NewBluetoothServiceBrowserController()
	_ = obj.GetOptions()
	// Output:
	}

// ExampleBluetoothServiceBrowserController_GetPrompt demonstrates using GetPrompt on a BluetoothServiceBrowserController instance.
// Returns the title of the default/select button in the device selector panel.
func ExampleBluetoothServiceBrowserController_GetPrompt() {
	obj := iobluetoothui.NewBluetoothServiceBrowserController()
	_ = obj.GetPrompt()
	// Output:
	}

// ExampleBluetoothServiceBrowserController_GetServiceBrowserControllerRef demonstrates using GetServiceBrowserControllerRef on a BluetoothServiceBrowserController instance.
// Returns an IOBluetoothServiceBrowserControllerRef representation of the target IOBluetoothServiceBrowserController object.
func ExampleBluetoothServiceBrowserController_GetServiceBrowserControllerRef() {
	obj := iobluetoothui.NewBluetoothServiceBrowserController()
	_ = obj.GetServiceBrowserControllerRef()
	// Output:
	}

// ExampleBluetoothServiceBrowserController_GetResults demonstrates using GetResults on a BluetoothServiceBrowserController instance.
// Returns the result of the user’s selection.
func ExampleBluetoothServiceBrowserController_GetResults() {
	obj := iobluetoothui.NewBluetoothServiceBrowserController()
	_ = obj.GetResults()
	// Output:
	}

// ExampleBluetoothServiceBrowserController_GetSearchAttributes demonstrates using GetSearchAttributes on a BluetoothServiceBrowserController instance.
// Returns the search attributes that control the panel’s search/inquiry behavior.
func ExampleBluetoothServiceBrowserController_GetSearchAttributes() {
	obj := iobluetoothui.NewBluetoothServiceBrowserController()
	_ = obj.GetSearchAttributes()
	// Output:
	}

// ExampleBluetoothServiceBrowserController_GetTitle demonstrates using GetTitle on a BluetoothServiceBrowserController instance.
// Returns the title of the device selector panel.
func ExampleBluetoothServiceBrowserController_GetTitle() {
	obj := iobluetoothui.NewBluetoothServiceBrowserController()
	_ = obj.GetTitle()
	// Output:
	}

// ExampleBluetoothServiceBrowserController_RunModal demonstrates using RunModal on a BluetoothServiceBrowserController instance.
// Runs the service browser panel in a modal session to allow the user to select a service on a Bluetooth device.
func ExampleBluetoothServiceBrowserController_RunModal() {
	obj := iobluetoothui.NewBluetoothServiceBrowserController()
	_ = obj.RunModal()
	// Output:
	}




