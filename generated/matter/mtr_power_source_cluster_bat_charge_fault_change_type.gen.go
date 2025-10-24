// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPowerSourceClusterBatChargeFaultChangeType */


/* debug [class_header]: Header for MTRPowerSourceClusterBatChargeFaultChangeType */
// The class instance for the [MTRPowerSourceClusterBatChargeFaultChangeType] class.
var (
	MTRPowerSourceClusterBatChargeFaultChangeTypeClass     _MTRPowerSourceClusterBatChargeFaultChangeTypeClass
	MTRPowerSourceClusterBatChargeFaultChangeTypeClassOnce sync.Once
)

func getMTRPowerSourceClusterBatChargeFaultChangeTypeClass() _MTRPowerSourceClusterBatChargeFaultChangeTypeClass {
	MTRPowerSourceClusterBatChargeFaultChangeTypeClassOnce.Do(func() {
		MTRPowerSourceClusterBatChargeFaultChangeTypeClass = _MTRPowerSourceClusterBatChargeFaultChangeTypeClass{objc.GetClass("MTRPowerSourceClusterBatChargeFaultChangeType")}
	})
	return MTRPowerSourceClusterBatChargeFaultChangeTypeClass
}

type _MTRPowerSourceClusterBatChargeFaultChangeTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPowerSourceClusterBatChargeFaultChangeType */
// An interface definition for the [MTRPowerSourceClusterBatChargeFaultChangeType] class.
type IMTRPowerSourceClusterBatChargeFaultChangeType interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPowerSourceClusterBatChargeFaultChangeType */
	// properties:
	Current() objc.IObject /* cross-framework: NSArray */
	SetCurrent(value objc.IObject /* cross-framework: NSArray */)
	Previous() objc.IObject /* cross-framework: NSArray */
	SetPrevious(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPowerSourceClusterBatChargeFaultChangeType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPowerSourceClusterBatChargeFaultChangeType */
// Alloc allocates a new instance without initialization.
func (mc _MTRPowerSourceClusterBatChargeFaultChangeTypeClass) Alloc() MTRPowerSourceClusterBatChargeFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPowerSourceClusterBatChargeFaultChangeTypeClass) New() MTRPowerSourceClusterBatChargeFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPowerSourceClusterBatChargeFaultChangeType) Init() MTRPowerSourceClusterBatChargeFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPowerSourceClusterBatChargeFaultChangeType) Autorelease() MTRPowerSourceClusterBatChargeFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPowerSourceClusterBatChargeFaultChangeType creates a new MTRPowerSourceClusterBatChargeFaultChangeType instance.
func NewMTRPowerSourceClusterBatChargeFaultChangeType() MTRPowerSourceClusterBatChargeFaultChangeType {
	return getMTRPowerSourceClusterBatChargeFaultChangeTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPowerSourceClusterBatChargeFaultChangeType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatChargeFaultChangeType
type MTRPowerSourceClusterBatChargeFaultChangeType struct {
	objectivec.Object
}

// MTRPowerSourceClusterBatChargeFaultChangeTypeFrom constructs a [MTRPowerSourceClusterBatChargeFaultChangeType] from an unsafe.Pointer.
func MTRPowerSourceClusterBatChargeFaultChangeTypeFrom(ptr unsafe.Pointer) MTRPowerSourceClusterBatChargeFaultChangeType {
	return MTRPowerSourceClusterBatChargeFaultChangeType{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPowerSourceClusterBatChargeFaultChangeType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPowerSourceClusterBatChargeFaultChangeType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPowerSourceClusterBatChargeFaultChangeType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPowerSourceClusterBatChargeFaultChangeType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPowerSourceClusterBatChargeFaultChangeType */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatChargeFaultChangeType/current
func (m_ MTRPowerSourceClusterBatChargeFaultChangeType) Current() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("current"))
	return rv
}/* debug [instance_properties/getter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatChargeFaultChangeType/current
func (m_ MTRPowerSourceClusterBatChargeFaultChangeType) SetCurrent(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}/* debug [instance_properties/setter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatChargeFaultChangeType/previous
func (m_ MTRPowerSourceClusterBatChargeFaultChangeType) Previous() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("previous"))
	return rv
}/* debug [instance_properties/getter]: previous */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatChargeFaultChangeType/previous
func (m_ MTRPowerSourceClusterBatChargeFaultChangeType) SetPrevious(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}/* debug [instance_properties/setter]: previous */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPowerSourceClusterBatChargeFaultChangeType */



