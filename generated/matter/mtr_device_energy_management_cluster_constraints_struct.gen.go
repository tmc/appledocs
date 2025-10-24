// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterConstraintsStruct */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterConstraintsStruct */
// The class instance for the [MTRDeviceEnergyManagementClusterConstraintsStruct] class.
var (
	MTRDeviceEnergyManagementClusterConstraintsStructClass     _MTRDeviceEnergyManagementClusterConstraintsStructClass
	MTRDeviceEnergyManagementClusterConstraintsStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterConstraintsStructClass() _MTRDeviceEnergyManagementClusterConstraintsStructClass {
	MTRDeviceEnergyManagementClusterConstraintsStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterConstraintsStructClass = _MTRDeviceEnergyManagementClusterConstraintsStructClass{objc.GetClass("MTRDeviceEnergyManagementClusterConstraintsStruct")}
	})
	return MTRDeviceEnergyManagementClusterConstraintsStructClass
}

type _MTRDeviceEnergyManagementClusterConstraintsStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterConstraintsStruct */
// An interface definition for the [MTRDeviceEnergyManagementClusterConstraintsStruct] class.
type IMTRDeviceEnergyManagementClusterConstraintsStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterConstraintsStruct */
	// properties:
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	LoadControl() objc.IObject /* cross-framework: NSNumber */
	SetLoadControl(value objc.IObject /* cross-framework: NSNumber */)
	MaximumEnergy() objc.IObject /* cross-framework: NSNumber */
	SetMaximumEnergy(value objc.IObject /* cross-framework: NSNumber */)
	NominalPower() objc.IObject /* cross-framework: NSNumber */
	SetNominalPower(value objc.IObject /* cross-framework: NSNumber */)
	StartTime() objc.IObject /* cross-framework: NSNumber */
	SetStartTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterConstraintsStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterConstraintsStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterConstraintsStructClass) Alloc() MTRDeviceEnergyManagementClusterConstraintsStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterConstraintsStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterConstraintsStructClass) New() MTRDeviceEnergyManagementClusterConstraintsStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterConstraintsStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) Init() MTRDeviceEnergyManagementClusterConstraintsStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterConstraintsStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) Autorelease() MTRDeviceEnergyManagementClusterConstraintsStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterConstraintsStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterConstraintsStruct creates a new MTRDeviceEnergyManagementClusterConstraintsStruct instance.
func NewMTRDeviceEnergyManagementClusterConstraintsStruct() MTRDeviceEnergyManagementClusterConstraintsStruct {
	return getMTRDeviceEnergyManagementClusterConstraintsStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterConstraintsStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct
type MTRDeviceEnergyManagementClusterConstraintsStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterConstraintsStructFrom constructs a [MTRDeviceEnergyManagementClusterConstraintsStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterConstraintsStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterConstraintsStruct {
	return MTRDeviceEnergyManagementClusterConstraintsStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterConstraintsStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterConstraintsStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterConstraintsStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterConstraintsStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterConstraintsStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct/duration
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct/duration
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterconstraintsstruct/loadcontrol
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) LoadControl() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("loadControl"))
	return rv
}/* debug [instance_properties/getter]: loadControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterconstraintsstruct/loadcontrol
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) SetLoadControl(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLoadControl:"), value)
}/* debug [instance_properties/setter]: loadControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterconstraintsstruct/maximumenergy
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) MaximumEnergy() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maximumEnergy"))
	return rv
}/* debug [instance_properties/getter]: maximumEnergy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterconstraintsstruct/maximumenergy
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) SetMaximumEnergy(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumEnergy:"), value)
}/* debug [instance_properties/setter]: maximumEnergy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterconstraintsstruct/nominalpower
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) NominalPower() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nominalPower"))
	return rv
}/* debug [instance_properties/getter]: nominalPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterconstraintsstruct/nominalpower
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) SetNominalPower(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalPower:"), value)
}/* debug [instance_properties/setter]: nominalPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterconstraintsstruct/starttime
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}/* debug [instance_properties/getter]: startTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterconstraintsstruct/starttime
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}/* debug [instance_properties/setter]: startTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterConstraintsStruct */



