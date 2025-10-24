// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceControllerAbstractParameters */


/* debug [class_header]: Header for MTRDeviceControllerAbstractParameters */
// The class instance for the [MTRDeviceControllerAbstractParameters] class.
var (
	MTRDeviceControllerAbstractParametersClass     _MTRDeviceControllerAbstractParametersClass
	MTRDeviceControllerAbstractParametersClassOnce sync.Once
)

func getMTRDeviceControllerAbstractParametersClass() _MTRDeviceControllerAbstractParametersClass {
	MTRDeviceControllerAbstractParametersClassOnce.Do(func() {
		MTRDeviceControllerAbstractParametersClass = _MTRDeviceControllerAbstractParametersClass{objc.GetClass("MTRDeviceControllerAbstractParameters")}
	})
	return MTRDeviceControllerAbstractParametersClass
}

type _MTRDeviceControllerAbstractParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceControllerAbstractParameters */
// An interface definition for the [MTRDeviceControllerAbstractParameters] class.
type IMTRDeviceControllerAbstractParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceControllerAbstractParameters */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceControllerAbstractParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceControllerAbstractParameters */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerAbstractParametersClass) Alloc() MTRDeviceControllerAbstractParameters {
	rv := objc.Send[MTRDeviceControllerAbstractParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceControllerAbstractParametersClass) New() MTRDeviceControllerAbstractParameters {
	rv := objc.Send[MTRDeviceControllerAbstractParameters](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceControllerAbstractParameters) Init() MTRDeviceControllerAbstractParameters {
	rv := objc.Send[MTRDeviceControllerAbstractParameters](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceControllerAbstractParameters) Autorelease() MTRDeviceControllerAbstractParameters {
	rv := objc.Send[MTRDeviceControllerAbstractParameters](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceControllerAbstractParameters creates a new MTRDeviceControllerAbstractParameters instance.
func NewMTRDeviceControllerAbstractParameters() MTRDeviceControllerAbstractParameters {
	return getMTRDeviceControllerAbstractParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceControllerAbstractParameters */
// A parent class referenced by other Matter classes.


// A parent class referenced by other Matter classes. [Full Topic]
type MTRDeviceControllerAbstractParameters struct {
	objectivec.Object
}

// MTRDeviceControllerAbstractParametersFrom constructs a [MTRDeviceControllerAbstractParameters] from an unsafe.Pointer.
//
// A parent class referenced by other Matter classes.
func MTRDeviceControllerAbstractParametersFrom(ptr unsafe.Pointer) MTRDeviceControllerAbstractParameters {
	return MTRDeviceControllerAbstractParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceControllerAbstractParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceControllerAbstractParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceControllerAbstractParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceControllerAbstractParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceControllerAbstractParameters */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceControllerAbstractParameters */



