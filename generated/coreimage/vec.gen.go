// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [vec] class.
var (
	VecClass     _vecClass
	VecClassOnce sync.Once
)

func getvecClass() _vecClass {
	VecClassOnce.Do(func() {
		VecClass = _vecClass{objc.GetClass("vec")}
	})
	return VecClass
}

type _vecClass struct {
	class objc.Class
}





// An interface definition for the [vec] class.
type Ivec interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _vecClass) Alloc() vec {
	rv := objc.Send[vec](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _vecClass) New() vec {
	rv := objc.Send[vec](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ vec) Init() vec {
	rv := objc.Send[vec](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ vec) Autorelease() vec {
	rv := objc.Send[vec](v_.ID, objc.Sel("autorelease"))
	return rv
}

// Newvec creates a new vec instance.
func Newvec() vec {
	return getvecClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/union_(unnamed)/vec
type vec struct {
	objectivec.Object
}

// vecFrom constructs a [vec] from an unsafe.Pointer.
func vecFrom(ptr unsafe.Pointer) vec {
	return vec{objectivec.Object{objc.ID(ptr)}}
}































