// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AppleEventDescriptor] class.
var (
	appleEventDescriptorClass     _AppleEventDescriptorClass
	appleEventDescriptorClassOnce sync.Once
)

func getAppleEventDescriptorClass() _AppleEventDescriptorClass {
	appleEventDescriptorClassOnce.Do(func() {
		appleEventDescriptorClass = _AppleEventDescriptorClass{objc.GetClass("NSAppleEventDescriptor")}
	})
	return appleEventDescriptorClass
}

type _AppleEventDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AppleEventDescriptor] class.
type IAppleEventDescriptor interface {
	objectivec.IObject
}

// A wrapper for the Apple event descriptor data type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor
type AppleEventDescriptor struct {
	objectivec.Object
}

// AppleEventDescriptorFrom constructs a [AppleEventDescriptor] from an unsafe.Pointer.
//
// A wrapper for the Apple event descriptor data type.
func AppleEventDescriptorFrom(ptr unsafe.Pointer) AppleEventDescriptor {
	return AppleEventDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AppleEventDescriptorClass) Alloc() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AppleEventDescriptorClass) New() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AppleEventDescriptor) Init() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AppleEventDescriptor) Autorelease() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAppleEventDescriptor creates a new AppleEventDescriptor instance.
func NewAppleEventDescriptor() AppleEventDescriptor {
	return getAppleEventDescriptorClass().New()
}




