// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GLKEffectProperty */


/* debug [class_header]: Header for GLKEffectProperty */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKEffectProperty */
// An interface definition for the [GLKEffectProperty] class.
type IGLKEffectProperty interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GLKEffectProperty */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKEffectProperty */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKEffectProperty */
// Alloc allocates a new instance without initialization.
func (gc _GLKEffectPropertyClass) Alloc() GLKEffectProperty {
	rv := objc.Send[GLKEffectProperty](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKEffectProperty */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKEffectProperty *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKEffectProperty */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKEffectProperty */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKEffectProperty */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKEffectProperty */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKEffectProperty */



