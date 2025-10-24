// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothHandsFreeDevice

// ExampleBluetoothHandsFreeDevice_AcceptCall demonstrates using AcceptCall on a BluetoothHandsFreeDevice instance.
// Accepts an incoming call.
func ExampleBluetoothHandsFreeDevice_AcceptCall() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.AcceptCall()
	// Output:
	}

// ExampleBluetoothHandsFreeDevice_AcceptCallOnPhone demonstrates using AcceptCallOnPhone on a BluetoothHandsFreeDevice instance.
// Accepts an incoming call and transfers the audio to the managed hands-free phone or headset.
func ExampleBluetoothHandsFreeDevice_AcceptCallOnPhone() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.AcceptCallOnPhone()
	// Output:
	}

// ExampleBluetoothHandsFreeDevice_AddHeldCall demonstrates using AddHeldCall on a BluetoothHandsFreeDevice instance.
// Adds held calls to the current conversation.
func ExampleBluetoothHandsFreeDevice_AddHeldCall() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.AddHeldCall()
	// Output:
	}

// ExampleBluetoothHandsFreeDevice_CallTransfer demonstrates using CallTransfer on a BluetoothHandsFreeDevice instance.
// Ends all calls that are active or on hold, and accepts any waiting calls.
func ExampleBluetoothHandsFreeDevice_CallTransfer() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.CallTransfer()
	// Output:
	}

// ExampleBluetoothHandsFreeDevice_CurrentCallList demonstrates using CurrentCallList on a BluetoothHandsFreeDevice instance.
// Requests that the Bluetooth audio gateway send the delegate a list of calls that are active, on hold, or being set up.
func ExampleBluetoothHandsFreeDevice_CurrentCallList() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.CurrentCallList()
	// Output:
	}

// ExampleBluetoothHandsFreeDevice_EndCall demonstrates using EndCall on a BluetoothHandsFreeDevice instance.
// Ends the current call or refuses an incoming call.
func ExampleBluetoothHandsFreeDevice_EndCall() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.EndCall()
	// Output:
	}

// ExampleBluetoothHandsFreeDevice_HoldCall demonstrates using HoldCall on a BluetoothHandsFreeDevice instance.
// Places all active calls on hold and accepts a held or waiting call.
func ExampleBluetoothHandsFreeDevice_HoldCall() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.HoldCall()
	// Output:
	}

// ExampleBluetoothHandsFreeDevice_Redial demonstrates using Redial on a BluetoothHandsFreeDevice instance.
// Calls the number stored on the hands-free phone or headset again.
func ExampleBluetoothHandsFreeDevice_Redial() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.Redial()
	// Output:
	}

// ExampleBluetoothHandsFreeDevice_ReleaseActiveCalls demonstrates using ReleaseActiveCalls on a BluetoothHandsFreeDevice instance.
// Ends all active calls and accepts a held or waiting call.
func ExampleBluetoothHandsFreeDevice_ReleaseActiveCalls() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.ReleaseActiveCalls()
	// Output:
	}

// ExampleBluetoothHandsFreeDevice_ReleaseHeldCalls demonstrates using ReleaseHeldCalls on a BluetoothHandsFreeDevice instance.
// Ends all calls that are on hold or returns a busy signal for a waiting call.
func ExampleBluetoothHandsFreeDevice_ReleaseHeldCalls() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.ReleaseHeldCalls()
	// Output:
	}

// ExampleBluetoothHandsFreeDevice_SubscriberNumber demonstrates using SubscriberNumber on a BluetoothHandsFreeDevice instance.
// Requests that the Bluetooth audio gateway send the subscriber number to the delegate.
func ExampleBluetoothHandsFreeDevice_SubscriberNumber() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.SubscriberNumber()
	// Output:
	}

// ExampleBluetoothHandsFreeDevice_TransferAudioToComputer demonstrates using TransferAudioToComputer on a BluetoothHandsFreeDevice instance.
// Moves the audio for current and future calls to a Mac.
func ExampleBluetoothHandsFreeDevice_TransferAudioToComputer() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.TransferAudioToComputer()
	// Output:
	}

// ExampleBluetoothHandsFreeDevice_TransferAudioToPhone demonstrates using TransferAudioToPhone on a BluetoothHandsFreeDevice instance.
// Moves the audio for current or future calls to a phone.
func ExampleBluetoothHandsFreeDevice_TransferAudioToPhone() {
	obj := iobluetooth.NewBluetoothHandsFreeDevice()
	obj.TransferAudioToPhone()
	// Output:
	}

