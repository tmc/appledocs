// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterForecastStruct */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterForecastStruct */
// The class instance for the [MTRDeviceEnergyManagementClusterForecastStruct] class.
var (
	MTRDeviceEnergyManagementClusterForecastStructClass     _MTRDeviceEnergyManagementClusterForecastStructClass
	MTRDeviceEnergyManagementClusterForecastStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterForecastStructClass() _MTRDeviceEnergyManagementClusterForecastStructClass {
	MTRDeviceEnergyManagementClusterForecastStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterForecastStructClass = _MTRDeviceEnergyManagementClusterForecastStructClass{objc.GetClass("MTRDeviceEnergyManagementClusterForecastStruct")}
	})
	return MTRDeviceEnergyManagementClusterForecastStructClass
}

type _MTRDeviceEnergyManagementClusterForecastStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterForecastStruct */
// An interface definition for the [MTRDeviceEnergyManagementClusterForecastStruct] class.
type IMTRDeviceEnergyManagementClusterForecastStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterForecastStruct */
	// properties:
	LatestEndTime() objc.IObject /* cross-framework: NSNumber */
	SetLatestEndTime(value objc.IObject /* cross-framework: NSNumber */)
	ActiveSlotNumber() objc.IObject /* cross-framework: NSNumber */
	SetActiveSlotNumber(value objc.IObject /* cross-framework: NSNumber */)
	EarliestStartTime() objc.IObject /* cross-framework: NSNumber */
	SetEarliestStartTime(value objc.IObject /* cross-framework: NSNumber */)
	EndTime() objc.IObject /* cross-framework: NSNumber */
	SetEndTime(value objc.IObject /* cross-framework: NSNumber */)
	ForecastID() objc.IObject /* cross-framework: NSNumber */
	SetForecastID(value objc.IObject /* cross-framework: NSNumber */)
	ForecastUpdateReason() objc.IObject /* cross-framework: NSNumber */
	SetForecastUpdateReason(value objc.IObject /* cross-framework: NSNumber */)
	IsPausable() objc.IObject /* cross-framework: NSNumber */
	SetIsPausable(value objc.IObject /* cross-framework: NSNumber */)
	StartTime() objc.IObject /* cross-framework: NSNumber */
	SetStartTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterForecastStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterForecastStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterForecastStructClass) Alloc() MTRDeviceEnergyManagementClusterForecastStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterForecastStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterForecastStructClass) New() MTRDeviceEnergyManagementClusterForecastStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterForecastStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) Init() MTRDeviceEnergyManagementClusterForecastStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterForecastStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) Autorelease() MTRDeviceEnergyManagementClusterForecastStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterForecastStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterForecastStruct creates a new MTRDeviceEnergyManagementClusterForecastStruct instance.
func NewMTRDeviceEnergyManagementClusterForecastStruct() MTRDeviceEnergyManagementClusterForecastStruct {
	return getMTRDeviceEnergyManagementClusterForecastStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterForecastStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct
type MTRDeviceEnergyManagementClusterForecastStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterForecastStructFrom constructs a [MTRDeviceEnergyManagementClusterForecastStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterForecastStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterForecastStruct {
	return MTRDeviceEnergyManagementClusterForecastStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterForecastStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterForecastStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterForecastStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterForecastStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterForecastStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/latestEndTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) LatestEndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("latestEndTime"))
	return rv
}/* debug [instance_properties/getter]: latestEndTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/latestEndTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetLatestEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLatestEndTime:"), value)
}/* debug [instance_properties/setter]: latestEndTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/activeslotnumber
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) ActiveSlotNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("activeSlotNumber"))
	return rv
}/* debug [instance_properties/getter]: activeSlotNumber */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/activeslotnumber
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetActiveSlotNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveSlotNumber:"), value)
}/* debug [instance_properties/setter]: activeSlotNumber */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/earlieststarttime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) EarliestStartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("earliestStartTime"))
	return rv
}/* debug [instance_properties/getter]: earliestStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/earlieststarttime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetEarliestStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEarliestStartTime:"), value)
}/* debug [instance_properties/setter]: earliestStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/endtime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) EndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endTime"))
	return rv
}/* debug [instance_properties/getter]: endTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/endtime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:"), value)
}/* debug [instance_properties/setter]: endTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/forecastid
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) ForecastID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("forecastID"))
	return rv
}/* debug [instance_properties/getter]: forecastID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/forecastid
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetForecastID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setForecastID:"), value)
}/* debug [instance_properties/setter]: forecastID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/forecastupdatereason
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) ForecastUpdateReason() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("forecastUpdateReason"))
	return rv
}/* debug [instance_properties/getter]: forecastUpdateReason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/forecastupdatereason
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetForecastUpdateReason(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setForecastUpdateReason:"), value)
}/* debug [instance_properties/setter]: forecastUpdateReason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/ispausable
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) IsPausable() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("isPausable"))
	return rv
}/* debug [instance_properties/getter]: isPausable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/ispausable
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetIsPausable(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsPausable:"), value)
}/* debug [instance_properties/setter]: isPausable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/starttime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}/* debug [instance_properties/getter]: startTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterforecaststruct/starttime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}/* debug [instance_properties/setter]: startTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterForecastStruct */



