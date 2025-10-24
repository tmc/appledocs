// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROvenCavityOperationalStateClusterOperationalCommandResponseParams */


/* debug [class_header]: Header for MTROvenCavityOperationalStateClusterOperationalCommandResponseParams */
// The class instance for the [MTROvenCavityOperationalStateClusterOperationalCommandResponseParams] class.
var (
	MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass     _MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass
	MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass() _MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass {
	MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass = _MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass{objc.GetClass("MTROvenCavityOperationalStateClusterOperationalCommandResponseParams")}
	})
	return MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass
}

type _MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROvenCavityOperationalStateClusterOperationalCommandResponseParams */
// An interface definition for the [MTROvenCavityOperationalStateClusterOperationalCommandResponseParams] class.
type IMTROvenCavityOperationalStateClusterOperationalCommandResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROvenCavityOperationalStateClusterOperationalCommandResponseParams */
	// properties:
	CommandResponseState() IMTROvenCavityOperationalStateClusterErrorStateStruct
	SetCommandResponseState(value IMTROvenCavityOperationalStateClusterErrorStateStruct)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROvenCavityOperationalStateClusterOperationalCommandResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROvenCavityOperationalStateClusterOperationalCommandResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass) Alloc() MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalCommandResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass) New() MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalCommandResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterOperationalCommandResponseParams) Init() MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalCommandResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterOperationalCommandResponseParams) Autorelease() MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalCommandResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterOperationalCommandResponseParams creates a new MTROvenCavityOperationalStateClusterOperationalCommandResponseParams instance.
func NewMTROvenCavityOperationalStateClusterOperationalCommandResponseParams() MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	return getMTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROvenCavityOperationalStateClusterOperationalCommandResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalCommandResponseParams
type MTROvenCavityOperationalStateClusterOperationalCommandResponseParams struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsFrom constructs a [MTROvenCavityOperationalStateClusterOperationalCommandResponseParams] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	return MTROvenCavityOperationalStateClusterOperationalCommandResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROvenCavityOperationalStateClusterOperationalCommandResponseParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROvenCavityOperationalStateClusterOperationalCommandResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROvenCavityOperationalStateClusterOperationalCommandResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROvenCavityOperationalStateClusterOperationalCommandResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROvenCavityOperationalStateClusterOperationalCommandResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalCommandResponseParams/commandResponseState
func (m_ MTROvenCavityOperationalStateClusterOperationalCommandResponseParams) CommandResponseState() IMTROvenCavityOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("commandResponseState"))
	return rv
}/* debug [instance_properties/getter]: commandResponseState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalCommandResponseParams/commandResponseState
func (m_ MTROvenCavityOperationalStateClusterOperationalCommandResponseParams) SetCommandResponseState(value IMTROvenCavityOperationalStateClusterErrorStateStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommandResponseState:"), value)
}/* debug [instance_properties/setter]: commandResponseState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROvenCavityOperationalStateClusterOperationalCommandResponseParams */



