// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mValue] class.
var (
	MValueClass     _mValueClass
	MValueClassOnce sync.Once
)

func getmValueClass() _mValueClass {
	MValueClassOnce.Do(func() {
		MValueClass = _mValueClass{objc.GetClass("mValue")}
	})
	return MValueClass
}

type _mValueClass struct {
	class objc.Class
}

// An interface definition for the [mValue] class.
type ImValue interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/mValue
type mValue struct {
	objectivec.Object
}

// mValueFrom constructs a [mValue] from an unsafe.Pointer.
func mValueFrom(ptr unsafe.Pointer) mValue {
	return mValue{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mValueClass) Alloc() mValue {
	rv := objc.Send[mValue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mValueClass) New() mValue {
	rv := objc.Send[mValue](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mValue) Init() mValue {
	rv := objc.Send[mValue](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mValue) Autorelease() mValue {
	rv := objc.Send[mValue](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmValue creates a new mValue instance.
func NewmValue() mValue {
	return getmValueClass().New()
}




