// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [layerCount] class.
var (
	LayerCountClass     _layerCountClass
	LayerCountClassOnce sync.Once
)

func getlayerCountClass() _layerCountClass {
	LayerCountClassOnce.Do(func() {
		LayerCountClass = _layerCountClass{objc.GetClass("layerCount")}
	})
	return LayerCountClass
}

type _layerCountClass struct {
	class objc.Class
}





// An interface definition for the [layerCount] class.
type IlayerCount interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (lc _layerCountClass) Alloc() layerCount {
	rv := objc.Send[layerCount](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
	return getlayerCountClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/layerCount-c.ivar
type layerCount struct {
	objectivec.Object
}

// layerCountFrom constructs a [layerCount] from an unsafe.Pointer.
func layerCountFrom(ptr unsafe.Pointer) layerCount {
	return layerCount{objectivec.Object{objc.ID(ptr)}}
}































