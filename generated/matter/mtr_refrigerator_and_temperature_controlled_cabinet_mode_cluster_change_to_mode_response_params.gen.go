// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams */


/* debug [class_header]: Header for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams */
// The class instance for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams] class.
var (
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass     _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass() _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass {
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass = _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams")}
	})
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass
}

type _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams */
// An interface definition for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams] class.
type IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams */
	// properties:
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass) Alloc() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass) New() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams) Init() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams) Autorelease() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams creates a new MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams instance.
func NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	return getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams
type MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsFrom constructs a [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams/status
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams/status
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratorandtemperaturecontrolledcabinetmodeclusterchangetomoderesponseparams/statustext
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}/* debug [instance_properties/getter]: statusText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratorandtemperaturecontrolledcabinetmodeclusterchangetomoderesponseparams/statustext
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}/* debug [instance_properties/setter]: statusText */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams */



