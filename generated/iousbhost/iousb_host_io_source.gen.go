// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [USBHostIOSource] class.
var (
	USBHostIOSourceClass     _USBHostIOSourceClass
	USBHostIOSourceClassOnce sync.Once
)

func getUSBHostIOSourceClass() _USBHostIOSourceClass {
	USBHostIOSourceClassOnce.Do(func() {
		USBHostIOSourceClass = _USBHostIOSourceClass{objc.GetClass("IOUSBHostIOSource")}
	})
	return USBHostIOSourceClass
}

type _USBHostIOSourceClass struct {
	class objc.Class
}

// An interface definition for the [USBHostIOSource] class.
type IUSBHostIOSource interface {
	objectivec.IObject
	DeviceAddress() uint
	EndpointAddress() uint
	HostInterface() IOUSBHostInterface
}

// This class provides basic functionality for deriving pipe and stream classes.
//
// Don’t create objects of this class or use this class as a subclass. Instead, use and when creating an .


// This class provides basic functionality for deriving pipe and stream classes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSource
type USBHostIOSource struct {
	objectivec.Object
}

// USBHostIOSourceFrom constructs a [USBHostIOSource] from an unsafe.Pointer.
//
// This class provides basic functionality for deriving pipe and stream classes.
func USBHostIOSourceFrom(ptr unsafe.Pointer) USBHostIOSource {
	return USBHostIOSource{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _USBHostIOSourceClass) Alloc() USBHostIOSource {
	rv := objc.Send[USBHostIOSource](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _USBHostIOSourceClass) New() USBHostIOSource {
	rv := objc.Send[USBHostIOSource](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ USBHostIOSource) Init() USBHostIOSource {
	rv := objc.Send[USBHostIOSource](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ USBHostIOSource) Autorelease() USBHostIOSource {
	rv := objc.Send[USBHostIOSource](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUSBHostIOSource creates a new USBHostIOSource instance.
func NewUSBHostIOSource() USBHostIOSource {
	return getUSBHostIOSourceClass().New()
}



// The device’s bus address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSource/deviceAddress
func (u_ USBHostIOSource) DeviceAddress() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("deviceAddress"))
	return rv
}


// The pipe or stream’s endpoint address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSource/endpointAddress
func (u_ USBHostIOSource) EndpointAddress() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("endpointAddress"))
	return rv
}


// The interface for the input/output source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSource/hostInterface
func (u_ USBHostIOSource) HostInterface() IOUSBHostInterface {
	rv := objc.Send[IOUSBHostInterface](u_.ID, objc.Sel("hostInterface"))
	return rv
}



