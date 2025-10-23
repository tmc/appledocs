// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mSizeDescriptor] class.
var (
	MSizeDescriptorClass     _mSizeDescriptorClass
	MSizeDescriptorClassOnce sync.Once
)

func getmSizeDescriptorClass() _mSizeDescriptorClass {
	MSizeDescriptorClassOnce.Do(func() {
		MSizeDescriptorClass = _mSizeDescriptorClass{objc.GetClass("mSizeDescriptor")}
	})
	return MSizeDescriptorClass
}

type _mSizeDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [mSizeDescriptor] class.
type ImSizeDescriptor interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/mSizeDescriptor
type mSizeDescriptor struct {
	objectivec.Object
}

// mSizeDescriptorFrom constructs a [mSizeDescriptor] from an unsafe.Pointer.
func mSizeDescriptorFrom(ptr unsafe.Pointer) mSizeDescriptor {
	return mSizeDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mSizeDescriptorClass) Alloc() mSizeDescriptor {
	rv := objc.Send[mSizeDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mSizeDescriptorClass) New() mSizeDescriptor {
	rv := objc.Send[mSizeDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mSizeDescriptor) Init() mSizeDescriptor {
	rv := objc.Send[mSizeDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mSizeDescriptor) Autorelease() mSizeDescriptor {
	rv := objc.Send[mSizeDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmSizeDescriptor creates a new mSizeDescriptor instance.
func NewmSizeDescriptor() mSizeDescriptor {
	return getmSizeDescriptorClass().New()
}




