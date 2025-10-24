// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementModeClusterModeTagStruct */


/* debug [class_header]: Header for MTRDeviceEnergyManagementModeClusterModeTagStruct */
// The class instance for the [MTRDeviceEnergyManagementModeClusterModeTagStruct] class.
var (
	MTRDeviceEnergyManagementModeClusterModeTagStructClass     _MTRDeviceEnergyManagementModeClusterModeTagStructClass
	MTRDeviceEnergyManagementModeClusterModeTagStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementModeClusterModeTagStructClass() _MTRDeviceEnergyManagementModeClusterModeTagStructClass {
	MTRDeviceEnergyManagementModeClusterModeTagStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementModeClusterModeTagStructClass = _MTRDeviceEnergyManagementModeClusterModeTagStructClass{objc.GetClass("MTRDeviceEnergyManagementModeClusterModeTagStruct")}
	})
	return MTRDeviceEnergyManagementModeClusterModeTagStructClass
}

type _MTRDeviceEnergyManagementModeClusterModeTagStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementModeClusterModeTagStruct */
// An interface definition for the [MTRDeviceEnergyManagementModeClusterModeTagStruct] class.
type IMTRDeviceEnergyManagementModeClusterModeTagStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementModeClusterModeTagStruct */
	// properties:
	MfgCode() objc.IObject /* cross-framework: NSNumber */
	SetMfgCode(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementModeClusterModeTagStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementModeClusterModeTagStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementModeClusterModeTagStructClass) Alloc() MTRDeviceEnergyManagementModeClusterModeTagStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementModeClusterModeTagStructClass) New() MTRDeviceEnergyManagementModeClusterModeTagStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementModeClusterModeTagStruct) Init() MTRDeviceEnergyManagementModeClusterModeTagStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementModeClusterModeTagStruct) Autorelease() MTRDeviceEnergyManagementModeClusterModeTagStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementModeClusterModeTagStruct creates a new MTRDeviceEnergyManagementModeClusterModeTagStruct instance.
func NewMTRDeviceEnergyManagementModeClusterModeTagStruct() MTRDeviceEnergyManagementModeClusterModeTagStruct {
	return getMTRDeviceEnergyManagementModeClusterModeTagStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementModeClusterModeTagStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeTagStruct
type MTRDeviceEnergyManagementModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementModeClusterModeTagStructFrom constructs a [MTRDeviceEnergyManagementModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementModeClusterModeTagStruct {
	return MTRDeviceEnergyManagementModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementModeClusterModeTagStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementModeClusterModeTagStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementModeClusterModeTagStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementModeClusterModeTagStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementModeClusterModeTagStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeTagStruct/mfgCode
func (m_ MTRDeviceEnergyManagementModeClusterModeTagStruct) MfgCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mfgCode"))
	return rv
}/* debug [instance_properties/getter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeTagStruct/mfgCode
func (m_ MTRDeviceEnergyManagementModeClusterModeTagStruct) SetMfgCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}/* debug [instance_properties/setter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementmodeclustermodetagstruct/value
func (m_ MTRDeviceEnergyManagementModeClusterModeTagStruct) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementmodeclustermodetagstruct/value
func (m_ MTRDeviceEnergyManagementModeClusterModeTagStruct) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementModeClusterModeTagStruct */



