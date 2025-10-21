// Code generated from Apple documentation for DeviceDiscoveryExtension. DO NOT EDIT.

package devicediscoveryextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DDDiscoverySession] class.
var (
	DDDiscoverySessionClass     _DDDiscoverySessionClass
	DDDiscoverySessionClassOnce sync.Once
)

func getDDDiscoverySessionClass() _DDDiscoverySessionClass {
	DDDiscoverySessionClassOnce.Do(func() {
		DDDiscoverySessionClass = _DDDiscoverySessionClass{objc.GetClass("DDDiscoverySession")}
	})
	return DDDiscoverySessionClass
}

type _DDDiscoverySessionClass struct {
	class objc.Class
}

// An interface definition for the [DDDiscoverySession] class.
type IDDDiscoverySession interface {
	objectivec.IObject
	ReportEvent(inEvent IDDDeviceEvent)
}

// An object that relays device discovery events from the extension to the system.
//
// The system passes the extension an instance of this class when it attempts to discover a device. Device discovery starts when an app displays and the system calls the extension’s implementation.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDiscoverySession
type DDDiscoverySession struct {
	objectivec.Object
}

// DDDiscoverySessionFrom constructs a [DDDiscoverySession] from an unsafe.Pointer.
//
// An object that relays device discovery events from the extension to the system.
func DDDiscoverySessionFrom(ptr unsafe.Pointer) DDDiscoverySession {
	return DDDiscoverySession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DDDiscoverySessionClass) Alloc() DDDiscoverySession {
	rv := objc.Send[DDDiscoverySession](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DDDiscoverySessionClass) New() DDDiscoverySession {
	rv := objc.Send[DDDiscoverySession](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDDiscoverySession) Init() DDDiscoverySession {
	rv := objc.Send[DDDiscoverySession](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDDiscoverySession) Autorelease() DDDiscoverySession {
	rv := objc.Send[DDDiscoverySession](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDDiscoverySession creates a new DDDiscoverySession instance.
func NewDDDiscoverySession() DDDiscoverySession {
	return getDDDiscoverySessionClass().New()
}


// Reports an event to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDiscoverySession/report(_:)
func (d_ DDDiscoverySession) ReportEvent(inEvent IDDDeviceEvent) {
	objc.Send[objc.ID](d_.ID, objc.Sel("reportEvent:"), inEvent)
}




