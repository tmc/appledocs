// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Shadow] class.
var (
	shadowClass     _ShadowClass
	shadowClassOnce sync.Once
)

func getShadowClass() _ShadowClass {
	shadowClassOnce.Do(func() {
		shadowClass = _ShadowClass{objc.GetClass("NSShadow")}
	})
	return shadowClass
}

type _ShadowClass struct {
	class objc.Class
}

// An interface definition for the [Shadow] class.
type IShadow interface {
	objectivec.IObject
}

// An object you use to specify attributes to create and style a drop shadow during drawing operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShadow

type Shadow struct {
	objectivec.Object
}

// ShadowFrom constructs a [Shadow] from an unsafe.Pointer.
//
// An object you use to specify attributes to create and style a drop shadow during drawing operations.
func ShadowFrom(ptr unsafe.Pointer) Shadow {
	return Shadow{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _ShadowClass) Alloc() Shadow {
	rv := objc.Send[Shadow](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _ShadowClass) New() Shadow {
	rv := objc.Send[Shadow](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Shadow) Init() Shadow {
	rv := objc.Send[Shadow](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Shadow) Autorelease() Shadow {
	rv := objc.Send[Shadow](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewShadow creates a new Shadow instance.
func NewShadow() Shadow {
	return getShadowClass().New()
}




