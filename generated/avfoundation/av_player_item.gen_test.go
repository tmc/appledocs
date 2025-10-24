// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewPlayerItem

// ExamplePlayerItem_AccessLog demonstrates using AccessLog on a PlayerItem instance.
// Returns an object that represents a snapshot of the network access log.
func ExamplePlayerItem_AccessLog() {
	obj := avfoundation.NewPlayerItem()
	_ = obj.AccessLog()
	// Output:
	}

// ExamplePlayerItem_CancelContentAuthorizationRequest demonstrates using CancelContentAuthorizationRequest on a PlayerItem instance.
// Cancels the currently outstanding content authorization request.
func ExamplePlayerItem_CancelContentAuthorizationRequest() {
	obj := avfoundation.NewPlayerItem()
	obj.CancelContentAuthorizationRequest()
	// Output:
	}

// ExamplePlayerItem_CancelPendingSeeks demonstrates using CancelPendingSeeks on a PlayerItem instance.
// Cancels any pending seek requests and invokes the corresponding completion handlers if present.
func ExamplePlayerItem_CancelPendingSeeks() {
	obj := avfoundation.NewPlayerItem()
	obj.CancelPendingSeeks()
	// Output:
	}

// ExamplePlayerItem_CurrentDate demonstrates using CurrentDate on a PlayerItem instance.
// Returns the current time of the item as a date.
func ExamplePlayerItem_CurrentDate() {
	obj := avfoundation.NewPlayerItem()
	_ = obj.CurrentDate()
	// Output:
	}

// ExamplePlayerItem_CurrentTime demonstrates using CurrentTime on a PlayerItem instance.
// Returns the current time of the item.
func ExamplePlayerItem_CurrentTime() {
	obj := avfoundation.NewPlayerItem()
	_ = obj.CurrentTime()
	// Output:
	}

// ExamplePlayerItem_ErrorLog demonstrates using ErrorLog on a PlayerItem instance.
// Returns an object that represents a snapshot of the error log.
func ExamplePlayerItem_ErrorLog() {
	obj := avfoundation.NewPlayerItem()
	_ = obj.ErrorLog()
	// Output:
	}

