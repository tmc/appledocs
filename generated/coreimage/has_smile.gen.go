// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasSmile] class.
var hasSmileClass = _hasSmileClass{objc.GetClass("hasSmile")}

type _hasSmileClass struct {
	class objc.Class
}

// An interface definition for the [hasSmile] class.
type IhasSmile interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasSmile-c.ivar

type hasSmile struct {
	objectivec.Object
}

// hasSmileFrom constructs a [hasSmile] from an unsafe.Pointer.
func hasSmileFrom(ptr unsafe.Pointer) hasSmile {
	return hasSmile{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (hc _hasSmileClass) Alloc() hasSmile {
	rv := objc.Send[hasSmile](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (hc _hasSmileClass) New() hasSmile {
	rv := objc.Send[hasSmile](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasSmile) Init() hasSmile {
	rv := objc.Send[hasSmile](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasSmile) Autorelease() hasSmile {
	rv := objc.Send[hasSmile](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasSmile creates a new hasSmile instance.
func NewhasSmile() hasSmile {
	return hasSmileClass.New()
}




