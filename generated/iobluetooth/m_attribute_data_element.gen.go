// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mAttributeDataElement] class.
var (
	MAttributeDataElementClass     _mAttributeDataElementClass
	MAttributeDataElementClassOnce sync.Once
)

func getmAttributeDataElementClass() _mAttributeDataElementClass {
	MAttributeDataElementClassOnce.Do(func() {
		MAttributeDataElementClass = _mAttributeDataElementClass{objc.GetClass("mAttributeDataElement")}
	})
	return MAttributeDataElementClass
}

type _mAttributeDataElementClass struct {
	class objc.Class
}

// An interface definition for the [mAttributeDataElement] class.
type ImAttributeDataElement interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/mAttributeDataElement
type mAttributeDataElement struct {
	objectivec.Object
}

// mAttributeDataElementFrom constructs a [mAttributeDataElement] from an unsafe.Pointer.
func mAttributeDataElementFrom(ptr unsafe.Pointer) mAttributeDataElement {
	return mAttributeDataElement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mAttributeDataElementClass) Alloc() mAttributeDataElement {
	rv := objc.Send[mAttributeDataElement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mAttributeDataElementClass) New() mAttributeDataElement {
	rv := objc.Send[mAttributeDataElement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mAttributeDataElement) Init() mAttributeDataElement {
	rv := objc.Send[mAttributeDataElement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mAttributeDataElement) Autorelease() mAttributeDataElement {
	rv := objc.Send[mAttributeDataElement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmAttributeDataElement creates a new mAttributeDataElement instance.
func NewmAttributeDataElement() mAttributeDataElement {
	return getmAttributeDataElementClass().New()
}




