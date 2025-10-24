// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode */


/* debug [class_header]: Header for MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode */
// The class instance for the [MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode] class.
var (
	MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass     _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass
	MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClassOnce sync.Once
)

func getMTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass() _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass {
	MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClassOnce.Do(func() {
		MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass = _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass{objc.GetClass("MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode")}
	})
	return MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass
}

type _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode */
// An interface definition for the [MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode] class.
type IMTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode */
	// methods:
	ChangeToModeWithParamsCompletion(params IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass) Alloc() MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass) New() MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) Init() MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) Autorelease() MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode creates a new MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode instance.
func NewMTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode() MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	return getMTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode */
// Cluster Refrigerator And Temperature Controlled Cabinet Mode
//
// Attributes and commands for selecting a mode from a list of supported options.


// Cluster Refrigerator And Temperature Controlled Cabinet Mode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode
type MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeFrom constructs a [MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode] from an unsafe.Pointer.
//
// Cluster Refrigerator And Temperature Controlled Cabinet Mode
func MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeFrom(ptr unsafe.Pointer) MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	return MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode */

// Command ChangeToMode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/changeToMode(with:completion:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) ChangeToModeWithParamsCompletion(params IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ChangeToModeWithParamsCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode */



