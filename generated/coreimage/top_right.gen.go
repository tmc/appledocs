// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [topRight] class.
var (
	TopRightClass     _topRightClass
	TopRightClassOnce sync.Once
)

func gettopRightClass() _topRightClass {
	TopRightClassOnce.Do(func() {
		TopRightClass = _topRightClass{objc.GetClass("topRight")}
	})
	return TopRightClass
}

type _topRightClass struct {
	class objc.Class
}





// An interface definition for the [topRight] class.
type ItopRight interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _topRightClass) Alloc() topRight {
	rv := objc.Send[topRight](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _topRightClass) New() topRight {
	rv := objc.Send[topRight](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ topRight) Init() topRight {
	rv := objc.Send[topRight](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ topRight) Autorelease() topRight {
	rv := objc.Send[topRight](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtopRight creates a new topRight instance.
func NewtopRight() topRight {
	return gettopRightClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/topRight-c.ivar
type topRight struct {
	objectivec.Object
}

// topRightFrom constructs a [topRight] from an unsafe.Pointer.
func topRightFrom(ptr unsafe.Pointer) topRight {
	return topRight{objectivec.Object{objc.ID(ptr)}}
}































