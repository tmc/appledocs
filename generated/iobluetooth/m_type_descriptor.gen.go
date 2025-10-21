// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mTypeDescriptor] class.
var (
	MTypeDescriptorClass     _mTypeDescriptorClass
	MTypeDescriptorClassOnce sync.Once
)

func getmTypeDescriptorClass() _mTypeDescriptorClass {
	MTypeDescriptorClassOnce.Do(func() {
		MTypeDescriptorClass = _mTypeDescriptorClass{objc.GetClass("mTypeDescriptor")}
	})
	return MTypeDescriptorClass
}

type _mTypeDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [mTypeDescriptor] class.
type ImTypeDescriptor interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/mTypeDescriptor
type mTypeDescriptor struct {
	objectivec.Object
}

// mTypeDescriptorFrom constructs a [mTypeDescriptor] from an unsafe.Pointer.
func mTypeDescriptorFrom(ptr unsafe.Pointer) mTypeDescriptor {
	return mTypeDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mTypeDescriptorClass) Alloc() mTypeDescriptor {
	rv := objc.Send[mTypeDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mTypeDescriptorClass) New() mTypeDescriptor {
	rv := objc.Send[mTypeDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mTypeDescriptor) Init() mTypeDescriptor {
	rv := objc.Send[mTypeDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mTypeDescriptor) Autorelease() mTypeDescriptor {
	rv := objc.Send[mTypeDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmTypeDescriptor creates a new mTypeDescriptor instance.
func NewmTypeDescriptor() mTypeDescriptor {
	return getmTypeDescriptorClass().New()
}




