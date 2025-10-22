// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [width] class.
var (
	WidthClass     _widthClass
	WidthClassOnce sync.Once
)

func getwidthClass() _widthClass {
	WidthClassOnce.Do(func() {
		WidthClass = _widthClass{objc.GetClass("width")}
	})
	return WidthClass
}

type _widthClass struct {
	class objc.Class
}

// An interface definition for the [width] class.
type Iwidth interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/width-c.ivar

type width struct {
	objectivec.Object
}

// widthFrom constructs a [width] from an unsafe.Pointer.
func widthFrom(ptr unsafe.Pointer) width {
	return width{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _widthClass) Alloc() width {
	rv := objc.Send[width](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _widthClass) New() width {
	rv := objc.Send[width](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ width) Init() width {
	rv := objc.Send[width](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ width) Autorelease() width {
	rv := objc.Send[width](w_.ID, objc.Sel("autorelease"))
	return rv
}

// Newwidth creates a new width instance.
func Newwidth() width {
	return getwidthClass().New()
}




