// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROccupancySensingClusterHoldTimeLimitsStruct */


/* debug [class_header]: Header for MTROccupancySensingClusterHoldTimeLimitsStruct */
// The class instance for the [MTROccupancySensingClusterHoldTimeLimitsStruct] class.
var (
	MTROccupancySensingClusterHoldTimeLimitsStructClass     _MTROccupancySensingClusterHoldTimeLimitsStructClass
	MTROccupancySensingClusterHoldTimeLimitsStructClassOnce sync.Once
)

func getMTROccupancySensingClusterHoldTimeLimitsStructClass() _MTROccupancySensingClusterHoldTimeLimitsStructClass {
	MTROccupancySensingClusterHoldTimeLimitsStructClassOnce.Do(func() {
		MTROccupancySensingClusterHoldTimeLimitsStructClass = _MTROccupancySensingClusterHoldTimeLimitsStructClass{objc.GetClass("MTROccupancySensingClusterHoldTimeLimitsStruct")}
	})
	return MTROccupancySensingClusterHoldTimeLimitsStructClass
}

type _MTROccupancySensingClusterHoldTimeLimitsStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROccupancySensingClusterHoldTimeLimitsStruct */
// An interface definition for the [MTROccupancySensingClusterHoldTimeLimitsStruct] class.
type IMTROccupancySensingClusterHoldTimeLimitsStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROccupancySensingClusterHoldTimeLimitsStruct */
	// properties:
	HoldTimeMax() objc.IObject /* cross-framework: NSNumber */
	SetHoldTimeMax(value objc.IObject /* cross-framework: NSNumber */)
	HoldTimeDefault() objc.IObject /* cross-framework: NSNumber */
	SetHoldTimeDefault(value objc.IObject /* cross-framework: NSNumber */)
	HoldTimeMin() objc.IObject /* cross-framework: NSNumber */
	SetHoldTimeMin(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROccupancySensingClusterHoldTimeLimitsStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROccupancySensingClusterHoldTimeLimitsStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTROccupancySensingClusterHoldTimeLimitsStructClass) Alloc() MTROccupancySensingClusterHoldTimeLimitsStruct {
	rv := objc.Send[MTROccupancySensingClusterHoldTimeLimitsStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROccupancySensingClusterHoldTimeLimitsStructClass) New() MTROccupancySensingClusterHoldTimeLimitsStruct {
	rv := objc.Send[MTROccupancySensingClusterHoldTimeLimitsStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) Init() MTROccupancySensingClusterHoldTimeLimitsStruct {
	rv := objc.Send[MTROccupancySensingClusterHoldTimeLimitsStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) Autorelease() MTROccupancySensingClusterHoldTimeLimitsStruct {
	rv := objc.Send[MTROccupancySensingClusterHoldTimeLimitsStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROccupancySensingClusterHoldTimeLimitsStruct creates a new MTROccupancySensingClusterHoldTimeLimitsStruct instance.
func NewMTROccupancySensingClusterHoldTimeLimitsStruct() MTROccupancySensingClusterHoldTimeLimitsStruct {
	return getMTROccupancySensingClusterHoldTimeLimitsStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROccupancySensingClusterHoldTimeLimitsStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterHoldTimeLimitsStruct
type MTROccupancySensingClusterHoldTimeLimitsStruct struct {
	objectivec.Object
}

// MTROccupancySensingClusterHoldTimeLimitsStructFrom constructs a [MTROccupancySensingClusterHoldTimeLimitsStruct] from an unsafe.Pointer.
func MTROccupancySensingClusterHoldTimeLimitsStructFrom(ptr unsafe.Pointer) MTROccupancySensingClusterHoldTimeLimitsStruct {
	return MTROccupancySensingClusterHoldTimeLimitsStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROccupancySensingClusterHoldTimeLimitsStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROccupancySensingClusterHoldTimeLimitsStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROccupancySensingClusterHoldTimeLimitsStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROccupancySensingClusterHoldTimeLimitsStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROccupancySensingClusterHoldTimeLimitsStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterHoldTimeLimitsStruct/holdTimeMax
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) HoldTimeMax() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("holdTimeMax"))
	return rv
}/* debug [instance_properties/getter]: holdTimeMax */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterHoldTimeLimitsStruct/holdTimeMax
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) SetHoldTimeMax(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHoldTimeMax:"), value)
}/* debug [instance_properties/setter]: holdTimeMax */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroccupancysensingclusterholdtimelimitsstruct/holdtimedefault
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) HoldTimeDefault() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("holdTimeDefault"))
	return rv
}/* debug [instance_properties/getter]: holdTimeDefault */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroccupancysensingclusterholdtimelimitsstruct/holdtimedefault
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) SetHoldTimeDefault(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHoldTimeDefault:"), value)
}/* debug [instance_properties/setter]: holdTimeDefault */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroccupancysensingclusterholdtimelimitsstruct/holdtimemin
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) HoldTimeMin() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("holdTimeMin"))
	return rv
}/* debug [instance_properties/getter]: holdTimeMin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroccupancysensingclusterholdtimelimitsstruct/holdtimemin
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) SetHoldTimeMin(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHoldTimeMin:"), value)
}/* debug [instance_properties/setter]: holdTimeMin */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROccupancySensingClusterHoldTimeLimitsStruct */



