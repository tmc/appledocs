// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPowerSourceClusterBatFaultChangeType */


/* debug [class_header]: Header for MTRPowerSourceClusterBatFaultChangeType */
// The class instance for the [MTRPowerSourceClusterBatFaultChangeType] class.
var (
	MTRPowerSourceClusterBatFaultChangeTypeClass     _MTRPowerSourceClusterBatFaultChangeTypeClass
	MTRPowerSourceClusterBatFaultChangeTypeClassOnce sync.Once
)

func getMTRPowerSourceClusterBatFaultChangeTypeClass() _MTRPowerSourceClusterBatFaultChangeTypeClass {
	MTRPowerSourceClusterBatFaultChangeTypeClassOnce.Do(func() {
		MTRPowerSourceClusterBatFaultChangeTypeClass = _MTRPowerSourceClusterBatFaultChangeTypeClass{objc.GetClass("MTRPowerSourceClusterBatFaultChangeType")}
	})
	return MTRPowerSourceClusterBatFaultChangeTypeClass
}

type _MTRPowerSourceClusterBatFaultChangeTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPowerSourceClusterBatFaultChangeType */
// An interface definition for the [MTRPowerSourceClusterBatFaultChangeType] class.
type IMTRPowerSourceClusterBatFaultChangeType interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPowerSourceClusterBatFaultChangeType */
	// properties:
	Current() objc.IObject /* cross-framework: NSArray */
	SetCurrent(value objc.IObject /* cross-framework: NSArray */)
	Previous() objc.IObject /* cross-framework: NSArray */
	SetPrevious(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPowerSourceClusterBatFaultChangeType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPowerSourceClusterBatFaultChangeType */
// Alloc allocates a new instance without initialization.
func (mc _MTRPowerSourceClusterBatFaultChangeTypeClass) Alloc() MTRPowerSourceClusterBatFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPowerSourceClusterBatFaultChangeTypeClass) New() MTRPowerSourceClusterBatFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPowerSourceClusterBatFaultChangeType) Init() MTRPowerSourceClusterBatFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPowerSourceClusterBatFaultChangeType) Autorelease() MTRPowerSourceClusterBatFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPowerSourceClusterBatFaultChangeType creates a new MTRPowerSourceClusterBatFaultChangeType instance.
func NewMTRPowerSourceClusterBatFaultChangeType() MTRPowerSourceClusterBatFaultChangeType {
	return getMTRPowerSourceClusterBatFaultChangeTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPowerSourceClusterBatFaultChangeType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatFaultChangeType
type MTRPowerSourceClusterBatFaultChangeType struct {
	objectivec.Object
}

// MTRPowerSourceClusterBatFaultChangeTypeFrom constructs a [MTRPowerSourceClusterBatFaultChangeType] from an unsafe.Pointer.
func MTRPowerSourceClusterBatFaultChangeTypeFrom(ptr unsafe.Pointer) MTRPowerSourceClusterBatFaultChangeType {
	return MTRPowerSourceClusterBatFaultChangeType{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPowerSourceClusterBatFaultChangeType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPowerSourceClusterBatFaultChangeType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPowerSourceClusterBatFaultChangeType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPowerSourceClusterBatFaultChangeType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPowerSourceClusterBatFaultChangeType */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatFaultChangeType/current
func (m_ MTRPowerSourceClusterBatFaultChangeType) Current() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("current"))
	return rv
}/* debug [instance_properties/getter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatFaultChangeType/current
func (m_ MTRPowerSourceClusterBatFaultChangeType) SetCurrent(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}/* debug [instance_properties/setter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatFaultChangeType/previous
func (m_ MTRPowerSourceClusterBatFaultChangeType) Previous() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("previous"))
	return rv
}/* debug [instance_properties/getter]: previous */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatFaultChangeType/previous
func (m_ MTRPowerSourceClusterBatFaultChangeType) SetPrevious(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}/* debug [instance_properties/setter]: previous */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPowerSourceClusterBatFaultChangeType */



