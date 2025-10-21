// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [bufferSize] class.
var (
	BufferSizeClass     _bufferSizeClass
	BufferSizeClassOnce sync.Once
)

func getbufferSizeClass() _bufferSizeClass {
	BufferSizeClassOnce.Do(func() {
		BufferSizeClass = _bufferSizeClass{objc.GetClass("bufferSize")}
	})
	return BufferSizeClass
}

type _bufferSizeClass struct {
	class objc.Class
}

// An interface definition for the [bufferSize] class.
type IbufferSize interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/bufferSize
type bufferSize struct {
	objectivec.Object
}

// bufferSizeFrom constructs a [bufferSize] from an unsafe.Pointer.
func bufferSizeFrom(ptr unsafe.Pointer) bufferSize {
	return bufferSize{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _bufferSizeClass) Alloc() bufferSize {
	rv := objc.Send[bufferSize](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _bufferSizeClass) New() bufferSize {
	rv := objc.Send[bufferSize](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ bufferSize) Init() bufferSize {
	rv := objc.Send[bufferSize](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ bufferSize) Autorelease() bufferSize {
	rv := objc.Send[bufferSize](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewbufferSize creates a new bufferSize instance.
func NewbufferSize() bufferSize {
	return getbufferSizeClass().New()
}




