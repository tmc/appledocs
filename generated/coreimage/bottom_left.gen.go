// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [bottomLeft] class.
var (
	BottomLeftClass     _bottomLeftClass
	BottomLeftClassOnce sync.Once
)

func getbottomLeftClass() _bottomLeftClass {
	BottomLeftClassOnce.Do(func() {
		BottomLeftClass = _bottomLeftClass{objc.GetClass("bottomLeft")}
	})
	return BottomLeftClass
}

type _bottomLeftClass struct {
	class objc.Class
}





// An interface definition for the [bottomLeft] class.
type IbottomLeft interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (bc _bottomLeftClass) Alloc() bottomLeft {
	rv := objc.Send[bottomLeft](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _bottomLeftClass) New() bottomLeft {
	rv := objc.Send[bottomLeft](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ bottomLeft) Init() bottomLeft {
	rv := objc.Send[bottomLeft](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ bottomLeft) Autorelease() bottomLeft {
	rv := objc.Send[bottomLeft](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewbottomLeft creates a new bottomLeft instance.
func NewbottomLeft() bottomLeft {
	return getbottomLeftClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/bottomLeft-c.ivar
type bottomLeft struct {
	objectivec.Object
}

// bottomLeftFrom constructs a [bottomLeft] from an unsafe.Pointer.
func bottomLeftFrom(ptr unsafe.Pointer) bottomLeft {
	return bottomLeft{objectivec.Object{objc.ID(ptr)}}
}































