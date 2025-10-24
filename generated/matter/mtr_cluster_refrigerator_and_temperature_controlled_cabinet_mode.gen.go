// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterRefrigeratorAndTemperatureControlledCabinetMode */


/* debug [class_header]: Header for MTRClusterRefrigeratorAndTemperatureControlledCabinetMode */
// The class instance for the [MTRClusterRefrigeratorAndTemperatureControlledCabinetMode] class.
var (
	MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass     _MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass
	MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClassOnce sync.Once
)

func getMTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass() _MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass {
	MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClassOnce.Do(func() {
		MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass = _MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass{objc.GetClass("MTRClusterRefrigeratorAndTemperatureControlledCabinetMode")}
	})
	return MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass
}

type _MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterRefrigeratorAndTemperatureControlledCabinetMode */
// An interface definition for the [MTRClusterRefrigeratorAndTemperatureControlledCabinetMode] class.
type IMTRClusterRefrigeratorAndTemperatureControlledCabinetMode interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterRefrigeratorAndTemperatureControlledCabinetMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterRefrigeratorAndTemperatureControlledCabinetMode */
	// methods:
	ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterRefrigeratorAndTemperatureControlledCabinetMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass) Alloc() MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRClusterRefrigeratorAndTemperatureControlledCabinetMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass) New() MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRClusterRefrigeratorAndTemperatureControlledCabinetMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) Init() MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRClusterRefrigeratorAndTemperatureControlledCabinetMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) Autorelease() MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRClusterRefrigeratorAndTemperatureControlledCabinetMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterRefrigeratorAndTemperatureControlledCabinetMode creates a new MTRClusterRefrigeratorAndTemperatureControlledCabinetMode instance.
func NewMTRClusterRefrigeratorAndTemperatureControlledCabinetMode() MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	return getMTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterRefrigeratorAndTemperatureControlledCabinetMode */
// Cluster Refrigerator And Temperature Controlled Cabinet Mode Attributes and commands for selecting a mode from a list of supported options.


// Cluster Refrigerator And Temperature Controlled Cabinet Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAndTemperatureControlledCabinetMode
type MTRClusterRefrigeratorAndTemperatureControlledCabinetMode struct {
	MTRGenericCluster
}

// MTRClusterRefrigeratorAndTemperatureControlledCabinetModeFrom constructs a [MTRClusterRefrigeratorAndTemperatureControlledCabinetMode] from an unsafe.Pointer.
//
// Cluster Refrigerator And Temperature Controlled Cabinet Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterRefrigeratorAndTemperatureControlledCabinetModeFrom(ptr unsafe.Pointer) MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	return MTRClusterRefrigeratorAndTemperatureControlledCabinetMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterRefrigeratorAndTemperatureControlledCabinetMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterRefrigeratorAndTemperatureControlledCabinetMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterRefrigeratorAndTemperatureControlledCabinetMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterRefrigeratorAndTemperatureControlledCabinetMode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAndTemperatureControlledCabinetMode/changeToMode(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}/* debug [instance_methods/method]: ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterRefrigeratorAndTemperatureControlledCabinetMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterRefrigeratorAndTemperatureControlledCabinetMode */



