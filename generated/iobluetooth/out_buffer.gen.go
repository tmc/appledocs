// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [outBuffer] class.
var (
	OutBufferClass     _outBufferClass
	OutBufferClassOnce sync.Once
)

func getoutBufferClass() _outBufferClass {
	OutBufferClassOnce.Do(func() {
		OutBufferClass = _outBufferClass{objc.GetClass("outBuffer")}
	})
	return OutBufferClass
}

type _outBufferClass struct {
	class objc.Class
}

// An interface definition for the [outBuffer] class.
type IoutBuffer interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/outBuffer
type outBuffer struct {
	objectivec.Object
}

// outBufferFrom constructs a [outBuffer] from an unsafe.Pointer.
func outBufferFrom(ptr unsafe.Pointer) outBuffer {
	return outBuffer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _outBufferClass) Alloc() outBuffer {
	rv := objc.Send[outBuffer](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _outBufferClass) New() outBuffer {
	rv := objc.Send[outBuffer](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ outBuffer) Init() outBuffer {
	rv := objc.Send[outBuffer](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ outBuffer) Autorelease() outBuffer {
	rv := objc.Send[outBuffer](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewoutBuffer creates a new outBuffer instance.
func NewoutBuffer() outBuffer {
	return getoutBufferClass().New()
}




