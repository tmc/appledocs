// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [layerCount] class.
var layerCountClass = _layerCountClass{objc.GetClass("layerCount")}

type _layerCountClass struct {
	class objc.Class
}

// An interface definition for the [layerCount] class.
type IlayerCount interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/layerCount-c.ivar

type layerCount struct {
	objectivec.Object
}

// layerCountFrom constructs a [layerCount] from an unsafe.Pointer.
func layerCountFrom(ptr unsafe.Pointer) layerCount {
	return layerCount{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (lc _layerCountClass) Alloc() layerCount {
	rv := objc.Send[layerCount](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (lc _layerCountClass) New() layerCount {
	rv := objc.Send[layerCount](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ layerCount) Init() layerCount {
	rv := objc.Send[layerCount](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ layerCount) Autorelease() layerCount {
	rv := objc.Send[layerCount](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewlayerCount creates a new layerCount instance.
func NewlayerCount() layerCount {
	return layerCountClass.New()
}




