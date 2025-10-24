// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterCostStruct */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterCostStruct */
// The class instance for the [MTRDeviceEnergyManagementClusterCostStruct] class.
var (
	MTRDeviceEnergyManagementClusterCostStructClass     _MTRDeviceEnergyManagementClusterCostStructClass
	MTRDeviceEnergyManagementClusterCostStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterCostStructClass() _MTRDeviceEnergyManagementClusterCostStructClass {
	MTRDeviceEnergyManagementClusterCostStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterCostStructClass = _MTRDeviceEnergyManagementClusterCostStructClass{objc.GetClass("MTRDeviceEnergyManagementClusterCostStruct")}
	})
	return MTRDeviceEnergyManagementClusterCostStructClass
}

type _MTRDeviceEnergyManagementClusterCostStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterCostStruct */
// An interface definition for the [MTRDeviceEnergyManagementClusterCostStruct] class.
type IMTRDeviceEnergyManagementClusterCostStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterCostStruct */
	// properties:
	Currency() objc.IObject /* cross-framework: NSNumber */
	SetCurrency(value objc.IObject /* cross-framework: NSNumber */)
	CostType() objc.IObject /* cross-framework: NSNumber */
	SetCostType(value objc.IObject /* cross-framework: NSNumber */)
	DecimalPoints() objc.IObject /* cross-framework: NSNumber */
	SetDecimalPoints(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterCostStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterCostStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterCostStructClass) Alloc() MTRDeviceEnergyManagementClusterCostStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCostStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterCostStructClass) New() MTRDeviceEnergyManagementClusterCostStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCostStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterCostStruct) Init() MTRDeviceEnergyManagementClusterCostStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCostStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterCostStruct) Autorelease() MTRDeviceEnergyManagementClusterCostStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCostStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterCostStruct creates a new MTRDeviceEnergyManagementClusterCostStruct instance.
func NewMTRDeviceEnergyManagementClusterCostStruct() MTRDeviceEnergyManagementClusterCostStruct {
	return getMTRDeviceEnergyManagementClusterCostStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterCostStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct
type MTRDeviceEnergyManagementClusterCostStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterCostStructFrom constructs a [MTRDeviceEnergyManagementClusterCostStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterCostStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterCostStruct {
	return MTRDeviceEnergyManagementClusterCostStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterCostStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterCostStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterCostStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterCostStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterCostStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/currency
func (m_ MTRDeviceEnergyManagementClusterCostStruct) Currency() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("currency"))
	return rv
}/* debug [instance_properties/getter]: currency */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/currency
func (m_ MTRDeviceEnergyManagementClusterCostStruct) SetCurrency(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrency:"), value)
}/* debug [instance_properties/setter]: currency */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustercoststruct/costtype
func (m_ MTRDeviceEnergyManagementClusterCostStruct) CostType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("costType"))
	return rv
}/* debug [instance_properties/getter]: costType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustercoststruct/costtype
func (m_ MTRDeviceEnergyManagementClusterCostStruct) SetCostType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCostType:"), value)
}/* debug [instance_properties/setter]: costType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustercoststruct/decimalpoints
func (m_ MTRDeviceEnergyManagementClusterCostStruct) DecimalPoints() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("decimalPoints"))
	return rv
}/* debug [instance_properties/getter]: decimalPoints */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustercoststruct/decimalpoints
func (m_ MTRDeviceEnergyManagementClusterCostStruct) SetDecimalPoints(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDecimalPoints:"), value)
}/* debug [instance_properties/setter]: decimalPoints */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustercoststruct/value
func (m_ MTRDeviceEnergyManagementClusterCostStruct) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustercoststruct/value
func (m_ MTRDeviceEnergyManagementClusterCostStruct) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterCostStruct */



