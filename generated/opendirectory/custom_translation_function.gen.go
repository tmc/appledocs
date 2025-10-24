// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class customTranslationFunction */


/* debug [class_header]: Header for customTranslationFunction */
// The class instance for the [customTranslationFunction] class.
var (
	CustomTranslationFunctionClass     _customTranslationFunctionClass
	CustomTranslationFunctionClassOnce sync.Once
)

func getcustomTranslationFunctionClass() _customTranslationFunctionClass {
	CustomTranslationFunctionClassOnce.Do(func() {
		CustomTranslationFunctionClass = _customTranslationFunctionClass{objc.GetClass("customTranslationFunction")}
	})
	return CustomTranslationFunctionClass
}

type _customTranslationFunctionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for customTranslationFunction */
// An interface definition for the [customTranslationFunction] class.
type IcustomTranslationFunction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for customTranslationFunction */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for customTranslationFunction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for customTranslationFunction */
// Alloc allocates a new instance without initialization.
func (cc _customTranslationFunctionClass) Alloc() customTranslationFunction {
	rv := objc.Send[customTranslationFunction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _customTranslationFunctionClass) New() customTranslationFunction {
	rv := objc.Send[customTranslationFunction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ customTranslationFunction) Init() customTranslationFunction {
	rv := objc.Send[customTranslationFunction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ customTranslationFunction) Autorelease() customTranslationFunction {
	rv := objc.Send[customTranslationFunction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewcustomTranslationFunction creates a new customTranslationFunction instance.
func NewcustomTranslationFunction() customTranslationFunction {
	return getcustomTranslationFunctionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for customTranslationFunction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customTranslationFunction-c.ivar
type customTranslationFunction struct {
	objectivec.Object
}

// customTranslationFunctionFrom constructs a [customTranslationFunction] from an unsafe.Pointer.
func customTranslationFunctionFrom(ptr unsafe.Pointer) customTranslationFunction {
	return customTranslationFunction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for customTranslationFunction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for customTranslationFunction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for customTranslationFunction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for customTranslationFunction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for customTranslationFunction */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class customTranslationFunction */



