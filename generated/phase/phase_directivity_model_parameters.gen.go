// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEDirectivityModelParameters */


/* debug [class_header]: Header for PHASEDirectivityModelParameters */
// The class instance for the [PHASEDirectivityModelParameters] class.
var (
	PHASEDirectivityModelParametersClass     _PHASEDirectivityModelParametersClass
	PHASEDirectivityModelParametersClassOnce sync.Once
)

func getPHASEDirectivityModelParametersClass() _PHASEDirectivityModelParametersClass {
	PHASEDirectivityModelParametersClassOnce.Do(func() {
		PHASEDirectivityModelParametersClass = _PHASEDirectivityModelParametersClass{objc.GetClass("PHASEDirectivityModelParameters")}
	})
	return PHASEDirectivityModelParametersClass
}

type _PHASEDirectivityModelParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEDirectivityModelParameters */
// An interface definition for the [PHASEDirectivityModelParameters] class.
type IPHASEDirectivityModelParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEDirectivityModelParameters */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEDirectivityModelParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEDirectivityModelParameters */
// Alloc allocates a new instance without initialization.
func (pc _PHASEDirectivityModelParametersClass) Alloc() PHASEDirectivityModelParameters {
	rv := objc.Send[PHASEDirectivityModelParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEDirectivityModelParametersClass) New() PHASEDirectivityModelParameters {
	rv := objc.Send[PHASEDirectivityModelParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEDirectivityModelParameters) Init() PHASEDirectivityModelParameters {
	rv := objc.Send[PHASEDirectivityModelParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEDirectivityModelParameters) Autorelease() PHASEDirectivityModelParameters {
	rv := objc.Send[PHASEDirectivityModelParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEDirectivityModelParameters creates a new PHASEDirectivityModelParameters instance.
func NewPHASEDirectivityModelParameters() PHASEDirectivityModelParameters {
	return getPHASEDirectivityModelParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEDirectivityModelParameters */
// A base class for objects that direct sound.
//
// Several classes derive from this class that implement a unique strategy to direct sound. Rather than create an instance of this class, instantiate a subclass, such as or .


// A base class for objects that direct sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDirectivityModelParameters
type PHASEDirectivityModelParameters struct {
	objectivec.Object
}

// PHASEDirectivityModelParametersFrom constructs a [PHASEDirectivityModelParameters] from an unsafe.Pointer.
//
// A base class for objects that direct sound.
func PHASEDirectivityModelParametersFrom(ptr unsafe.Pointer) PHASEDirectivityModelParameters {
	return PHASEDirectivityModelParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEDirectivityModelParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEDirectivityModelParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEDirectivityModelParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEDirectivityModelParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEDirectivityModelParameters */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEDirectivityModelParameters */



