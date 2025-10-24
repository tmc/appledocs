// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRXPCDeviceControllerParameters */


/* debug [class_header]: Header for MTRXPCDeviceControllerParameters */
// The class instance for the [MTRXPCDeviceControllerParameters] class.
var (
	MTRXPCDeviceControllerParametersClass     _MTRXPCDeviceControllerParametersClass
	MTRXPCDeviceControllerParametersClassOnce sync.Once
)

func getMTRXPCDeviceControllerParametersClass() _MTRXPCDeviceControllerParametersClass {
	MTRXPCDeviceControllerParametersClassOnce.Do(func() {
		MTRXPCDeviceControllerParametersClass = _MTRXPCDeviceControllerParametersClass{objc.GetClass("MTRXPCDeviceControllerParameters")}
	})
	return MTRXPCDeviceControllerParametersClass
}

type _MTRXPCDeviceControllerParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRXPCDeviceControllerParameters */
// An interface definition for the [MTRXPCDeviceControllerParameters] class.
type IMTRXPCDeviceControllerParameters interface {
	IMTRDeviceControllerAbstractParameters
	
/* debug [class_interface_properties]: Properties for MTRXPCDeviceControllerParameters */
	// properties:
	XpcConnectionBlock() unsafe.Pointer
	UniqueIdentifier() foundation.UUID
	SetUniqueIdentifier(value foundation.UUID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRXPCDeviceControllerParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRXPCDeviceControllerParameters */
// Alloc allocates a new instance without initialization.
func (mc _MTRXPCDeviceControllerParametersClass) Alloc() MTRXPCDeviceControllerParameters {
	rv := objc.Send[MTRXPCDeviceControllerParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRXPCDeviceControllerParametersClass) New() MTRXPCDeviceControllerParameters {
	rv := objc.Send[MTRXPCDeviceControllerParameters](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRXPCDeviceControllerParameters) Init() MTRXPCDeviceControllerParameters {
	rv := objc.Send[MTRXPCDeviceControllerParameters](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRXPCDeviceControllerParameters) Autorelease() MTRXPCDeviceControllerParameters {
	rv := objc.Send[MTRXPCDeviceControllerParameters](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRXPCDeviceControllerParameters creates a new MTRXPCDeviceControllerParameters instance.
func NewMTRXPCDeviceControllerParameters() MTRXPCDeviceControllerParameters {
	return getMTRXPCDeviceControllerParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRXPCDeviceControllerParameters */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRXPCDeviceControllerParameters
type MTRXPCDeviceControllerParameters struct {
	MTRDeviceControllerAbstractParameters
}

// MTRXPCDeviceControllerParametersFrom constructs a [MTRXPCDeviceControllerParameters] from an unsafe.Pointer.
func MTRXPCDeviceControllerParametersFrom(ptr unsafe.Pointer) MTRXPCDeviceControllerParameters {
	return MTRXPCDeviceControllerParameters{
		MTRDeviceControllerAbstractParameters: MTRDeviceControllerAbstractParametersFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRXPCDeviceControllerParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRXPCDeviceControllerParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRXPCDeviceControllerParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRXPCDeviceControllerParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRXPCDeviceControllerParameters */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRXPCDeviceControllerParameters/xpcConnectionBlock
func (m_ MTRXPCDeviceControllerParameters) XpcConnectionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("xpcConnectionBlock"))
	return rv
}/* debug [instance_properties/getter]: xpcConnectionBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrxpcdevicecontrollerparameters/uniqueidentifier
func (m_ MTRXPCDeviceControllerParameters) UniqueIdentifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](m_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}/* debug [instance_properties/getter]: uniqueIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrxpcdevicecontrollerparameters/uniqueidentifier
func (m_ MTRXPCDeviceControllerParameters) SetUniqueIdentifier(value foundation.UUID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUniqueIdentifier:"), value)
}/* debug [instance_properties/setter]: uniqueIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRXPCDeviceControllerParameters */



