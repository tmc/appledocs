// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementModeClusterModeOptionStruct */


/* debug [class_header]: Header for MTRDeviceEnergyManagementModeClusterModeOptionStruct */
// The class instance for the [MTRDeviceEnergyManagementModeClusterModeOptionStruct] class.
var (
	MTRDeviceEnergyManagementModeClusterModeOptionStructClass     _MTRDeviceEnergyManagementModeClusterModeOptionStructClass
	MTRDeviceEnergyManagementModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementModeClusterModeOptionStructClass() _MTRDeviceEnergyManagementModeClusterModeOptionStructClass {
	MTRDeviceEnergyManagementModeClusterModeOptionStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementModeClusterModeOptionStructClass = _MTRDeviceEnergyManagementModeClusterModeOptionStructClass{objc.GetClass("MTRDeviceEnergyManagementModeClusterModeOptionStruct")}
	})
	return MTRDeviceEnergyManagementModeClusterModeOptionStructClass
}

type _MTRDeviceEnergyManagementModeClusterModeOptionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementModeClusterModeOptionStruct */
// An interface definition for the [MTRDeviceEnergyManagementModeClusterModeOptionStruct] class.
type IMTRDeviceEnergyManagementModeClusterModeOptionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementModeClusterModeOptionStruct */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementModeClusterModeOptionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementModeClusterModeOptionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementModeClusterModeOptionStructClass) Alloc() MTRDeviceEnergyManagementModeClusterModeOptionStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementModeClusterModeOptionStructClass) New() MTRDeviceEnergyManagementModeClusterModeOptionStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) Init() MTRDeviceEnergyManagementModeClusterModeOptionStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) Autorelease() MTRDeviceEnergyManagementModeClusterModeOptionStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementModeClusterModeOptionStruct creates a new MTRDeviceEnergyManagementModeClusterModeOptionStruct instance.
func NewMTRDeviceEnergyManagementModeClusterModeOptionStruct() MTRDeviceEnergyManagementModeClusterModeOptionStruct {
	return getMTRDeviceEnergyManagementModeClusterModeOptionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementModeClusterModeOptionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct
type MTRDeviceEnergyManagementModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementModeClusterModeOptionStructFrom constructs a [MTRDeviceEnergyManagementModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementModeClusterModeOptionStruct {
	return MTRDeviceEnergyManagementModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementModeClusterModeOptionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementModeClusterModeOptionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementModeClusterModeOptionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementModeClusterModeOptionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementModeClusterModeOptionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/label
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/label
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementmodeclustermodeoptionstruct/mode
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementmodeclustermodeoptionstruct/mode
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementModeClusterModeOptionStruct */



