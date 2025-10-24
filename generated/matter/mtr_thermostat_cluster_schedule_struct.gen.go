// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterScheduleStruct */


/* debug [class_header]: Header for MTRThermostatClusterScheduleStruct */
// The class instance for the [MTRThermostatClusterScheduleStruct] class.
var (
	MTRThermostatClusterScheduleStructClass     _MTRThermostatClusterScheduleStructClass
	MTRThermostatClusterScheduleStructClassOnce sync.Once
)

func getMTRThermostatClusterScheduleStructClass() _MTRThermostatClusterScheduleStructClass {
	MTRThermostatClusterScheduleStructClassOnce.Do(func() {
		MTRThermostatClusterScheduleStructClass = _MTRThermostatClusterScheduleStructClass{objc.GetClass("MTRThermostatClusterScheduleStruct")}
	})
	return MTRThermostatClusterScheduleStructClass
}

type _MTRThermostatClusterScheduleStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterScheduleStruct */
// An interface definition for the [MTRThermostatClusterScheduleStruct] class.
type IMTRThermostatClusterScheduleStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterScheduleStruct */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	BuiltIn() objc.IObject /* cross-framework: NSNumber */
	SetBuiltIn(value objc.IObject /* cross-framework: NSNumber */)
	PresetHandle() foundation.Data
	SetPresetHandle(value foundation.Data)
	ScheduleHandle() foundation.Data
	SetScheduleHandle(value foundation.Data)
	SystemMode() objc.IObject /* cross-framework: NSNumber */
	SetSystemMode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterScheduleStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterScheduleStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterScheduleStructClass) Alloc() MTRThermostatClusterScheduleStruct {
	rv := objc.Send[MTRThermostatClusterScheduleStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterScheduleStructClass) New() MTRThermostatClusterScheduleStruct {
	rv := objc.Send[MTRThermostatClusterScheduleStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterScheduleStruct) Init() MTRThermostatClusterScheduleStruct {
	rv := objc.Send[MTRThermostatClusterScheduleStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterScheduleStruct) Autorelease() MTRThermostatClusterScheduleStruct {
	rv := objc.Send[MTRThermostatClusterScheduleStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterScheduleStruct creates a new MTRThermostatClusterScheduleStruct instance.
func NewMTRThermostatClusterScheduleStruct() MTRThermostatClusterScheduleStruct {
	return getMTRThermostatClusterScheduleStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterScheduleStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct
type MTRThermostatClusterScheduleStruct struct {
	objectivec.Object
}

// MTRThermostatClusterScheduleStructFrom constructs a [MTRThermostatClusterScheduleStruct] from an unsafe.Pointer.
func MTRThermostatClusterScheduleStructFrom(ptr unsafe.Pointer) MTRThermostatClusterScheduleStruct {
	return MTRThermostatClusterScheduleStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterScheduleStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterScheduleStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterScheduleStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterScheduleStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterScheduleStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/name
func (m_ MTRThermostatClusterScheduleStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/name
func (m_ MTRThermostatClusterScheduleStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterschedulestruct/builtin
func (m_ MTRThermostatClusterScheduleStruct) BuiltIn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("builtIn"))
	return rv
}/* debug [instance_properties/getter]: builtIn */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterschedulestruct/builtin
func (m_ MTRThermostatClusterScheduleStruct) SetBuiltIn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBuiltIn:"), value)
}/* debug [instance_properties/setter]: builtIn */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterschedulestruct/presethandle
func (m_ MTRThermostatClusterScheduleStruct) PresetHandle() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("presetHandle"))
	return rv
}/* debug [instance_properties/getter]: presetHandle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterschedulestruct/presethandle
func (m_ MTRThermostatClusterScheduleStruct) SetPresetHandle(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetHandle:"), value)
}/* debug [instance_properties/setter]: presetHandle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterschedulestruct/schedulehandle
func (m_ MTRThermostatClusterScheduleStruct) ScheduleHandle() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("scheduleHandle"))
	return rv
}/* debug [instance_properties/getter]: scheduleHandle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterschedulestruct/schedulehandle
func (m_ MTRThermostatClusterScheduleStruct) SetScheduleHandle(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScheduleHandle:"), value)
}/* debug [instance_properties/setter]: scheduleHandle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterschedulestruct/systemmode
func (m_ MTRThermostatClusterScheduleStruct) SystemMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("systemMode"))
	return rv
}/* debug [instance_properties/getter]: systemMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterschedulestruct/systemmode
func (m_ MTRThermostatClusterScheduleStruct) SetSystemMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSystemMode:"), value)
}/* debug [instance_properties/setter]: systemMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterScheduleStruct */



