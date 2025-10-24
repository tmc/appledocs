// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterPresetTypeStruct */


/* debug [class_header]: Header for MTRThermostatClusterPresetTypeStruct */
// The class instance for the [MTRThermostatClusterPresetTypeStruct] class.
var (
	MTRThermostatClusterPresetTypeStructClass     _MTRThermostatClusterPresetTypeStructClass
	MTRThermostatClusterPresetTypeStructClassOnce sync.Once
)

func getMTRThermostatClusterPresetTypeStructClass() _MTRThermostatClusterPresetTypeStructClass {
	MTRThermostatClusterPresetTypeStructClassOnce.Do(func() {
		MTRThermostatClusterPresetTypeStructClass = _MTRThermostatClusterPresetTypeStructClass{objc.GetClass("MTRThermostatClusterPresetTypeStruct")}
	})
	return MTRThermostatClusterPresetTypeStructClass
}

type _MTRThermostatClusterPresetTypeStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterPresetTypeStruct */
// An interface definition for the [MTRThermostatClusterPresetTypeStruct] class.
type IMTRThermostatClusterPresetTypeStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterPresetTypeStruct */
	// properties:
	PresetTypeFeatures() objc.IObject /* cross-framework: NSNumber */
	SetPresetTypeFeatures(value objc.IObject /* cross-framework: NSNumber */)
	NumberOfPresets() objc.IObject /* cross-framework: NSNumber */
	SetNumberOfPresets(value objc.IObject /* cross-framework: NSNumber */)
	PresetScenario() objc.IObject /* cross-framework: NSNumber */
	SetPresetScenario(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterPresetTypeStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterPresetTypeStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterPresetTypeStructClass) Alloc() MTRThermostatClusterPresetTypeStruct {
	rv := objc.Send[MTRThermostatClusterPresetTypeStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterPresetTypeStructClass) New() MTRThermostatClusterPresetTypeStruct {
	rv := objc.Send[MTRThermostatClusterPresetTypeStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterPresetTypeStruct) Init() MTRThermostatClusterPresetTypeStruct {
	rv := objc.Send[MTRThermostatClusterPresetTypeStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterPresetTypeStruct) Autorelease() MTRThermostatClusterPresetTypeStruct {
	rv := objc.Send[MTRThermostatClusterPresetTypeStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterPresetTypeStruct creates a new MTRThermostatClusterPresetTypeStruct instance.
func NewMTRThermostatClusterPresetTypeStruct() MTRThermostatClusterPresetTypeStruct {
	return getMTRThermostatClusterPresetTypeStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterPresetTypeStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct
type MTRThermostatClusterPresetTypeStruct struct {
	objectivec.Object
}

// MTRThermostatClusterPresetTypeStructFrom constructs a [MTRThermostatClusterPresetTypeStruct] from an unsafe.Pointer.
func MTRThermostatClusterPresetTypeStructFrom(ptr unsafe.Pointer) MTRThermostatClusterPresetTypeStruct {
	return MTRThermostatClusterPresetTypeStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterPresetTypeStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterPresetTypeStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterPresetTypeStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterPresetTypeStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterPresetTypeStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/presetTypeFeatures
func (m_ MTRThermostatClusterPresetTypeStruct) PresetTypeFeatures() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("presetTypeFeatures"))
	return rv
}/* debug [instance_properties/getter]: presetTypeFeatures */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/presetTypeFeatures
func (m_ MTRThermostatClusterPresetTypeStruct) SetPresetTypeFeatures(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetTypeFeatures:"), value)
}/* debug [instance_properties/setter]: presetTypeFeatures */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresettypestruct/numberofpresets
func (m_ MTRThermostatClusterPresetTypeStruct) NumberOfPresets() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("numberOfPresets"))
	return rv
}/* debug [instance_properties/getter]: numberOfPresets */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresettypestruct/numberofpresets
func (m_ MTRThermostatClusterPresetTypeStruct) SetNumberOfPresets(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfPresets:"), value)
}/* debug [instance_properties/setter]: numberOfPresets */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresettypestruct/presetscenario
func (m_ MTRThermostatClusterPresetTypeStruct) PresetScenario() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("presetScenario"))
	return rv
}/* debug [instance_properties/getter]: presetScenario */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresettypestruct/presetscenario
func (m_ MTRThermostatClusterPresetTypeStruct) SetPresetScenario(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetScenario:"), value)
}/* debug [instance_properties/setter]: presetScenario */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterPresetTypeStruct */



