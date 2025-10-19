// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Gradient] class.
var (
	gradientClass     _GradientClass
	gradientClassOnce sync.Once
)

func getGradientClass() _GradientClass {
	gradientClassOnce.Do(func() {
		gradientClass = _GradientClass{objc.GetClass("NSGradient")}
	})
	return gradientClass
}

type _GradientClass struct {
	class objc.Class
}

// An interface definition for the [Gradient] class.
type IGradient interface {
	objectivec.IObject
}

// An object that can draw gradient fill colors [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGradient

type Gradient struct {
	objectivec.Object
}

// GradientFrom constructs a [Gradient] from an unsafe.Pointer.
//
// An object that can draw gradient fill colors
func GradientFrom(ptr unsafe.Pointer) Gradient {
	return Gradient{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (gc _GradientClass) Alloc() Gradient {
	rv := objc.Send[Gradient](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (gc _GradientClass) New() Gradient {
	rv := objc.Send[Gradient](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ Gradient) Init() Gradient {
	rv := objc.Send[Gradient](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ Gradient) Autorelease() Gradient {
	rv := objc.Send[Gradient](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGradient creates a new Gradient instance.
func NewGradient() Gradient {
	return getGradientClass().New()
}




