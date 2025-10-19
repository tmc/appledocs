// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasMouthPosition] class.
var hasMouthPositionClass = _hasMouthPositionClass{objc.GetClass("hasMouthPosition")}

type _hasMouthPositionClass struct {
	class objc.Class
}

// An interface definition for the [hasMouthPosition] class.
type IhasMouthPosition interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasMouthPosition-c.ivar

type hasMouthPosition struct {
	objectivec.Object
}

// hasMouthPositionFrom constructs a [hasMouthPosition] from an unsafe.Pointer.
func hasMouthPositionFrom(ptr unsafe.Pointer) hasMouthPosition {
	return hasMouthPosition{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (hc _hasMouthPositionClass) Alloc() hasMouthPosition {
	rv := objc.Send[hasMouthPosition](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (hc _hasMouthPositionClass) New() hasMouthPosition {
	rv := objc.Send[hasMouthPosition](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasMouthPosition) Init() hasMouthPosition {
	rv := objc.Send[hasMouthPosition](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasMouthPosition) Autorelease() hasMouthPosition {
	rv := objc.Send[hasMouthPosition](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasMouthPosition creates a new hasMouthPosition instance.
func NewhasMouthPosition() hasMouthPosition {
	return hasMouthPositionClass.New()
}




