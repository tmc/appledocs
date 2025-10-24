// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui_test

import (
	"github.com/tmc/appledocs/generated/iobluetoothui"
)

// Suppress unused import errors
var _ = iobluetoothui.NewBluetoothObjectPushUIController

// ExampleBluetoothObjectPushUIController_GetDevice demonstrates using GetDevice on a BluetoothObjectPushUIController instance.
// Gets the object representing the remote target device in the transfer.
func ExampleBluetoothObjectPushUIController_GetDevice() {
	obj := iobluetoothui.NewBluetoothObjectPushUIController()
	_ = obj.GetDevice()
	// Output:
	}

// ExampleBluetoothObjectPushUIController_GetTitle demonstrates using GetTitle on a BluetoothObjectPushUIController instance.
// Returns the title of the transfer panel.
func ExampleBluetoothObjectPushUIController_GetTitle() {
	obj := iobluetoothui.NewBluetoothObjectPushUIController()
	_ = obj.GetTitle()
	// Output:
	}

// ExampleBluetoothObjectPushUIController_IsTransferInProgress demonstrates using IsTransferInProgress on a BluetoothObjectPushUIController instance.
// Gets state of the transfer
func ExampleBluetoothObjectPushUIController_IsTransferInProgress() {
	obj := iobluetoothui.NewBluetoothObjectPushUIController()
	_ = obj.IsTransferInProgress()
	// Output:
	}

// ExampleBluetoothObjectPushUIController_RunModal demonstrates using RunModal on a BluetoothObjectPushUIController instance.
// Runs the transfer UI panel in a modal session
func ExampleBluetoothObjectPushUIController_RunModal() {
	obj := iobluetoothui.NewBluetoothObjectPushUIController()
	obj.RunModal()
	// Output:
	}

// ExampleBluetoothObjectPushUIController_RunPanel demonstrates using RunPanel on a BluetoothObjectPushUIController instance.
// Runs the transfer UI as a panel with no modal session
func ExampleBluetoothObjectPushUIController_RunPanel() {
	obj := iobluetoothui.NewBluetoothObjectPushUIController()
	obj.RunPanel()
	// Output:
	}

// ExampleBluetoothObjectPushUIController_Stop demonstrates using Stop on a BluetoothObjectPushUIController instance.
// Stops the transfer UI
func ExampleBluetoothObjectPushUIController_Stop() {
	obj := iobluetoothui.NewBluetoothObjectPushUIController()
	obj.Stop()
	// Output:
	}

