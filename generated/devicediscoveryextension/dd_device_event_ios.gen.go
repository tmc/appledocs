//go:build darwin && ios

// Code generated from Apple documentation for DeviceDiscoveryExtension. DO NOT EDIT.

package devicediscoveryextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for DDDeviceEvent


// iOS-only properties

// An object that describes a third-party media receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/device
func (d_ DDDeviceEvent) Device() IDDDevice {
	rv := objc.Send[DDDevice](d_.ID, objc.Sel("device"))
	return rv
}

// A type for the event that describes the discovery status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/eventType-swift.property
func (d_ DDDeviceEvent) EventType() DDEventType {
	rv := objc.Send[DDEventType](d_.ID, objc.Sel("eventType"))
	return rv
}




