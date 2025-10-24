// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class alphaState */


/* debug [class_header]: Header for alphaState */
// The class instance for the [alphaState] class.
var (
	AlphaStateClass     _alphaStateClass
	AlphaStateClassOnce sync.Once
)

func getalphaStateClass() _alphaStateClass {
	AlphaStateClassOnce.Do(func() {
		AlphaStateClass = _alphaStateClass{objc.GetClass("alphaState")}
	})
	return AlphaStateClass
}

type _alphaStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for alphaState */
// An interface definition for the [alphaState] class.
type IalphaState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for alphaState */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for alphaState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for alphaState */
// Alloc allocates a new instance without initialization.
func (ac _alphaStateClass) Alloc() alphaState {
	rv := objc.Send[alphaState](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _alphaStateClass) New() alphaState {
	rv := objc.Send[alphaState](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ alphaState) Init() alphaState {
	rv := objc.Send[alphaState](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ alphaState) Autorelease() alphaState {
	rv := objc.Send[alphaState](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewalphaState creates a new alphaState instance.
func NewalphaState() alphaState {
	return getalphaStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for alphaState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/alphaState-c.ivar
type alphaState struct {
	objectivec.Object
}

// alphaStateFrom constructs a [alphaState] from an unsafe.Pointer.
func alphaStateFrom(ptr unsafe.Pointer) alphaState {
	return alphaState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for alphaState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for alphaState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for alphaState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for alphaState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for alphaState */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class alphaState */



