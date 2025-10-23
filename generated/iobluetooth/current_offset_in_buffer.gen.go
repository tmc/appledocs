// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [currentOffsetInBuffer] class.
var (
	CurrentOffsetInBufferClass     _currentOffsetInBufferClass
	CurrentOffsetInBufferClassOnce sync.Once
)

func getcurrentOffsetInBufferClass() _currentOffsetInBufferClass {
	CurrentOffsetInBufferClassOnce.Do(func() {
		CurrentOffsetInBufferClass = _currentOffsetInBufferClass{objc.GetClass("currentOffsetInBuffer")}
	})
	return CurrentOffsetInBufferClass
}

type _currentOffsetInBufferClass struct {
	class objc.Class
}

// An interface definition for the [currentOffsetInBuffer] class.
type IcurrentOffsetInBuffer interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/currentOffsetInBuffer
type currentOffsetInBuffer struct {
	objectivec.Object
}

// currentOffsetInBufferFrom constructs a [currentOffsetInBuffer] from an unsafe.Pointer.
func currentOffsetInBufferFrom(ptr unsafe.Pointer) currentOffsetInBuffer {
	return currentOffsetInBuffer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _currentOffsetInBufferClass) Alloc() currentOffsetInBuffer {
	rv := objc.Send[currentOffsetInBuffer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _currentOffsetInBufferClass) New() currentOffsetInBuffer {
	rv := objc.Send[currentOffsetInBuffer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ currentOffsetInBuffer) Init() currentOffsetInBuffer {
	rv := objc.Send[currentOffsetInBuffer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ currentOffsetInBuffer) Autorelease() currentOffsetInBuffer {
	rv := objc.Send[currentOffsetInBuffer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewcurrentOffsetInBuffer creates a new currentOffsetInBuffer instance.
func NewcurrentOffsetInBuffer() currentOffsetInBuffer {
	return getcurrentOffsetInBufferClass().New()
}




