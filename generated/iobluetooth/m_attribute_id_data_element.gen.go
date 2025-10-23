// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mAttributeIDDataElement] class.
var (
	MAttributeIDDataElementClass     _mAttributeIDDataElementClass
	MAttributeIDDataElementClassOnce sync.Once
)

func getmAttributeIDDataElementClass() _mAttributeIDDataElementClass {
	MAttributeIDDataElementClassOnce.Do(func() {
		MAttributeIDDataElementClass = _mAttributeIDDataElementClass{objc.GetClass("mAttributeIDDataElement")}
	})
	return MAttributeIDDataElementClass
}

type _mAttributeIDDataElementClass struct {
	class objc.Class
}

// An interface definition for the [mAttributeIDDataElement] class.
type ImAttributeIDDataElement interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/mAttributeIDDataElement
type mAttributeIDDataElement struct {
	objectivec.Object
}

// mAttributeIDDataElementFrom constructs a [mAttributeIDDataElement] from an unsafe.Pointer.
func mAttributeIDDataElementFrom(ptr unsafe.Pointer) mAttributeIDDataElement {
	return mAttributeIDDataElement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mAttributeIDDataElementClass) Alloc() mAttributeIDDataElement {
	rv := objc.Send[mAttributeIDDataElement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mAttributeIDDataElementClass) New() mAttributeIDDataElement {
	rv := objc.Send[mAttributeIDDataElement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mAttributeIDDataElement) Init() mAttributeIDDataElement {
	rv := objc.Send[mAttributeIDDataElement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mAttributeIDDataElement) Autorelease() mAttributeIDDataElement {
	rv := objc.Send[mAttributeIDDataElement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmAttributeIDDataElement creates a new mAttributeIDDataElement instance.
func NewmAttributeIDDataElement() mAttributeIDDataElement {
	return getmAttributeIDDataElementClass().New()
}




