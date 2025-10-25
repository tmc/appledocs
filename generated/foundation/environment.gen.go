// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class environment */


/* debug [class_header]: Header for environment */
// The class instance for the [environment] class.
var (
	EnvironmentClass     _environmentClass
	EnvironmentClassOnce sync.Once
)

func getenvironmentClass() _environmentClass {
	EnvironmentClassOnce.Do(func() {
		EnvironmentClass = _environmentClass{objc.GetClass("environment")}
	})
	return EnvironmentClass
}

type _environmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for environment */
// An interface definition for the [environment] class.
type Ienvironment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for environment */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for environment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for environment */
// Alloc allocates a new instance without initialization.
func (ec _environmentClass) Alloc() environment {
	rv := objc.Send[environment](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _environmentClass) New() environment {
	rv := objc.Send[environment](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ environment) Init() environment {
	rv := objc.Send[environment](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ environment) Autorelease() environment {
	rv := objc.Send[environment](e_.ID, objc.Sel("autorelease"))
	return rv
}

// Newenvironment creates a new environment instance.
func Newenvironment() environment {
	return getenvironmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for environment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProcessInfo/environment-c.ivar
type environment struct {
	objectivec.Object
}

// environmentFrom constructs a [environment] from an unsafe.Pointer.
func environmentFrom(ptr unsafe.Pointer) environment {
	return environment{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for environment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for environment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for environment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for environment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for environment */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class environment */



