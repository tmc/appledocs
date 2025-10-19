// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [symbolDescriptor] class.
var (
	symbolDescriptorClass     _symbolDescriptorClass
	symbolDescriptorClassOnce sync.Once
)

func getsymbolDescriptorClass() _symbolDescriptorClass {
	symbolDescriptorClassOnce.Do(func() {
		symbolDescriptorClass = _symbolDescriptorClass{objc.GetClass("symbolDescriptor")}
	})
	return symbolDescriptorClass
}

type _symbolDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [symbolDescriptor] class.
type IsymbolDescriptor interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/symbolDescriptor-c.ivar
type symbolDescriptor struct {
	objectivec.Object
}

// symbolDescriptorFrom constructs a [symbolDescriptor] from an unsafe.Pointer.
func symbolDescriptorFrom(ptr unsafe.Pointer) symbolDescriptor {
	return symbolDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _symbolDescriptorClass) Alloc() symbolDescriptor {
	rv := objc.Send[symbolDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _symbolDescriptorClass) New() symbolDescriptor {
	rv := objc.Send[symbolDescriptor](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ symbolDescriptor) Init() symbolDescriptor {
	rv := objc.Send[symbolDescriptor](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ symbolDescriptor) Autorelease() symbolDescriptor {
	rv := objc.Send[symbolDescriptor](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewsymbolDescriptor creates a new symbolDescriptor instance.
func NewsymbolDescriptor() symbolDescriptor {
	return getsymbolDescriptorClass().New()
}




