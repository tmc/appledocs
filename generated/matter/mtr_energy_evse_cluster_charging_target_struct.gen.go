// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterChargingTargetStruct */


/* debug [class_header]: Header for MTREnergyEVSEClusterChargingTargetStruct */
// The class instance for the [MTREnergyEVSEClusterChargingTargetStruct] class.
var (
	MTREnergyEVSEClusterChargingTargetStructClass     _MTREnergyEVSEClusterChargingTargetStructClass
	MTREnergyEVSEClusterChargingTargetStructClassOnce sync.Once
)

func getMTREnergyEVSEClusterChargingTargetStructClass() _MTREnergyEVSEClusterChargingTargetStructClass {
	MTREnergyEVSEClusterChargingTargetStructClassOnce.Do(func() {
		MTREnergyEVSEClusterChargingTargetStructClass = _MTREnergyEVSEClusterChargingTargetStructClass{objc.GetClass("MTREnergyEVSEClusterChargingTargetStruct")}
	})
	return MTREnergyEVSEClusterChargingTargetStructClass
}

type _MTREnergyEVSEClusterChargingTargetStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterChargingTargetStruct */
// An interface definition for the [MTREnergyEVSEClusterChargingTargetStruct] class.
type IMTREnergyEVSEClusterChargingTargetStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterChargingTargetStruct */
	// properties:
	AddedEnergy() objc.IObject /* cross-framework: NSNumber */
	SetAddedEnergy(value objc.IObject /* cross-framework: NSNumber */)
	TargetSoC() objc.IObject /* cross-framework: NSNumber */
	SetTargetSoC(value objc.IObject /* cross-framework: NSNumber */)
	TargetTimeMinutesPastMidnight() objc.IObject /* cross-framework: NSNumber */
	SetTargetTimeMinutesPastMidnight(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterChargingTargetStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterChargingTargetStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterChargingTargetStructClass) Alloc() MTREnergyEVSEClusterChargingTargetStruct {
	rv := objc.Send[MTREnergyEVSEClusterChargingTargetStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterChargingTargetStructClass) New() MTREnergyEVSEClusterChargingTargetStruct {
	rv := objc.Send[MTREnergyEVSEClusterChargingTargetStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterChargingTargetStruct) Init() MTREnergyEVSEClusterChargingTargetStruct {
	rv := objc.Send[MTREnergyEVSEClusterChargingTargetStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterChargingTargetStruct) Autorelease() MTREnergyEVSEClusterChargingTargetStruct {
	rv := objc.Send[MTREnergyEVSEClusterChargingTargetStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterChargingTargetStruct creates a new MTREnergyEVSEClusterChargingTargetStruct instance.
func NewMTREnergyEVSEClusterChargingTargetStruct() MTREnergyEVSEClusterChargingTargetStruct {
	return getMTREnergyEVSEClusterChargingTargetStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterChargingTargetStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetStruct
type MTREnergyEVSEClusterChargingTargetStruct struct {
	objectivec.Object
}

// MTREnergyEVSEClusterChargingTargetStructFrom constructs a [MTREnergyEVSEClusterChargingTargetStruct] from an unsafe.Pointer.
func MTREnergyEVSEClusterChargingTargetStructFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterChargingTargetStruct {
	return MTREnergyEVSEClusterChargingTargetStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterChargingTargetStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterChargingTargetStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterChargingTargetStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterChargingTargetStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterChargingTargetStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetStruct/addedEnergy
func (m_ MTREnergyEVSEClusterChargingTargetStruct) AddedEnergy() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("addedEnergy"))
	return rv
}/* debug [instance_properties/getter]: addedEnergy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetStruct/addedEnergy
func (m_ MTREnergyEVSEClusterChargingTargetStruct) SetAddedEnergy(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAddedEnergy:"), value)
}/* debug [instance_properties/setter]: addedEnergy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterchargingtargetstruct/targetsoc
func (m_ MTREnergyEVSEClusterChargingTargetStruct) TargetSoC() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetSoC"))
	return rv
}/* debug [instance_properties/getter]: targetSoC */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterchargingtargetstruct/targetsoc
func (m_ MTREnergyEVSEClusterChargingTargetStruct) SetTargetSoC(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetSoC:"), value)
}/* debug [instance_properties/setter]: targetSoC */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterchargingtargetstruct/targettimeminutespastmidnight
func (m_ MTREnergyEVSEClusterChargingTargetStruct) TargetTimeMinutesPastMidnight() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetTimeMinutesPastMidnight"))
	return rv
}/* debug [instance_properties/getter]: targetTimeMinutesPastMidnight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterchargingtargetstruct/targettimeminutespastmidnight
func (m_ MTREnergyEVSEClusterChargingTargetStruct) SetTargetTimeMinutesPastMidnight(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetTimeMinutesPastMidnight:"), value)
}/* debug [instance_properties/setter]: targetTimeMinutesPastMidnight */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterChargingTargetStruct */



