// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GLKEffectProperty] class.
var (
	GLKEffectPropertyClass     _GLKEffectPropertyClass
	GLKEffectPropertyClassOnce sync.Once
)

func getGLKEffectPropertyClass() _GLKEffectPropertyClass {
	GLKEffectPropertyClassOnce.Do(func() {
		GLKEffectPropertyClass = _GLKEffectPropertyClass{objc.GetClass("GLKEffectProperty")}
	})
	return GLKEffectPropertyClass
}

type _GLKEffectPropertyClass struct {
	class objc.Class
}

// An interface definition for the [GLKEffectProperty] class.
type IGLKEffectProperty interface {
	objectivec.IObject
}

// The abstract superclass for configuration information used in GLKit rendering effects.
//
// Subclasses of provide one or more Objective-C properties that define how that state can be configured for an effect.


// The abstract superclass for configuration information used in GLKit rendering effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectProperty
type GLKEffectProperty struct {
	objectivec.Object
}

// GLKEffectPropertyFrom constructs a [GLKEffectProperty] from an unsafe.Pointer.
//
// The abstract superclass for configuration information used in GLKit rendering effects.
func GLKEffectPropertyFrom(ptr unsafe.Pointer) GLKEffectProperty {
	return GLKEffectProperty{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GLKEffectPropertyClass) Alloc() GLKEffectProperty {
	rv := objc.Send[GLKEffectProperty](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GLKEffectPropertyClass) New() GLKEffectProperty {
	rv := objc.Send[GLKEffectProperty](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKEffectProperty) Init() GLKEffectProperty {
	rv := objc.Send[GLKEffectProperty](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKEffectProperty) Autorelease() GLKEffectProperty {
	rv := objc.Send[GLKEffectProperty](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKEffectProperty creates a new GLKEffectProperty instance.
func NewGLKEffectProperty() GLKEffectProperty {
	return getGLKEffectPropertyClass().New()
}




