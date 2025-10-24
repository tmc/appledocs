//go:build darwin && ios

// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CBCentralManager


// Register for an event notification when the central manager makes a connection matching the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/registerForConnectionEvents(options:)
func (c_ CBCentralManager) RegisterForConnectionEventsWithOptions(options foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerForConnectionEventsWithOptions:"), options)
}

// iOS-only properties




