// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct */


/* debug [class_header]: Header for MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct */
// The class instance for the [MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct] class.
var (
	MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass     _MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass
	MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClassOnce sync.Once
)

func getMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass() _MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass {
	MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClassOnce.Do(func() {
		MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass = _MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass{objc.GetClass("MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct")}
	})
	return MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass
}

type _MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct */
// An interface definition for the [MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct] class.
type IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct */
	// properties:
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	EmergencyBoost() objc.IObject /* cross-framework: NSNumber */
	SetEmergencyBoost(value objc.IObject /* cross-framework: NSNumber */)
	OneShot() objc.IObject /* cross-framework: NSNumber */
	SetOneShot(value objc.IObject /* cross-framework: NSNumber */)
	TargetPercentage() objc.IObject /* cross-framework: NSNumber */
	SetTargetPercentage(value objc.IObject /* cross-framework: NSNumber */)
	TargetReheat() objc.IObject /* cross-framework: NSNumber */
	SetTargetReheat(value objc.IObject /* cross-framework: NSNumber */)
	TemporarySetpoint() objc.IObject /* cross-framework: NSNumber */
	SetTemporarySetpoint(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass) Alloc() MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	rv := objc.Send[MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass) New() MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	rv := objc.Send[MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) Init() MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	rv := objc.Send[MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) Autorelease() MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	rv := objc.Send[MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct creates a new MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct instance.
func NewMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct() MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	return getMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct
type MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct struct {
	objectivec.Object
}

// MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructFrom constructs a [MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct] from an unsafe.Pointer.
func MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructFrom(ptr unsafe.Pointer) MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	return MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/duration
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/duration
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterwaterheaterboostinfostruct/emergencyboost
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) EmergencyBoost() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("emergencyBoost"))
	return rv
}/* debug [instance_properties/getter]: emergencyBoost */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterwaterheaterboostinfostruct/emergencyboost
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetEmergencyBoost(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEmergencyBoost:"), value)
}/* debug [instance_properties/setter]: emergencyBoost */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterwaterheaterboostinfostruct/oneshot
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) OneShot() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("oneShot"))
	return rv
}/* debug [instance_properties/getter]: oneShot */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterwaterheaterboostinfostruct/oneshot
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetOneShot(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOneShot:"), value)
}/* debug [instance_properties/setter]: oneShot */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterwaterheaterboostinfostruct/targetpercentage
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) TargetPercentage() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetPercentage"))
	return rv
}/* debug [instance_properties/getter]: targetPercentage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterwaterheaterboostinfostruct/targetpercentage
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetTargetPercentage(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetPercentage:"), value)
}/* debug [instance_properties/setter]: targetPercentage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterwaterheaterboostinfostruct/targetreheat
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) TargetReheat() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetReheat"))
	return rv
}/* debug [instance_properties/getter]: targetReheat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterwaterheaterboostinfostruct/targetreheat
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetTargetReheat(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetReheat:"), value)
}/* debug [instance_properties/setter]: targetReheat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterwaterheaterboostinfostruct/temporarysetpoint
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) TemporarySetpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("temporarySetpoint"))
	return rv
}/* debug [instance_properties/getter]: temporarySetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterwaterheaterboostinfostruct/temporarysetpoint
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetTemporarySetpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTemporarySetpoint:"), value)
}/* debug [instance_properties/setter]: temporarySetpoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct */



