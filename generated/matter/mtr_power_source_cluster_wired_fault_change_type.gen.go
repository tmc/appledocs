// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPowerSourceClusterWiredFaultChangeType */


/* debug [class_header]: Header for MTRPowerSourceClusterWiredFaultChangeType */
// The class instance for the [MTRPowerSourceClusterWiredFaultChangeType] class.
var (
	MTRPowerSourceClusterWiredFaultChangeTypeClass     _MTRPowerSourceClusterWiredFaultChangeTypeClass
	MTRPowerSourceClusterWiredFaultChangeTypeClassOnce sync.Once
)

func getMTRPowerSourceClusterWiredFaultChangeTypeClass() _MTRPowerSourceClusterWiredFaultChangeTypeClass {
	MTRPowerSourceClusterWiredFaultChangeTypeClassOnce.Do(func() {
		MTRPowerSourceClusterWiredFaultChangeTypeClass = _MTRPowerSourceClusterWiredFaultChangeTypeClass{objc.GetClass("MTRPowerSourceClusterWiredFaultChangeType")}
	})
	return MTRPowerSourceClusterWiredFaultChangeTypeClass
}

type _MTRPowerSourceClusterWiredFaultChangeTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPowerSourceClusterWiredFaultChangeType */
// An interface definition for the [MTRPowerSourceClusterWiredFaultChangeType] class.
type IMTRPowerSourceClusterWiredFaultChangeType interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPowerSourceClusterWiredFaultChangeType */
	// properties:
	Current() objc.IObject /* cross-framework: NSArray */
	SetCurrent(value objc.IObject /* cross-framework: NSArray */)
	Previous() objc.IObject /* cross-framework: NSArray */
	SetPrevious(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPowerSourceClusterWiredFaultChangeType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPowerSourceClusterWiredFaultChangeType */
// Alloc allocates a new instance without initialization.
func (mc _MTRPowerSourceClusterWiredFaultChangeTypeClass) Alloc() MTRPowerSourceClusterWiredFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPowerSourceClusterWiredFaultChangeTypeClass) New() MTRPowerSourceClusterWiredFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPowerSourceClusterWiredFaultChangeType) Init() MTRPowerSourceClusterWiredFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPowerSourceClusterWiredFaultChangeType) Autorelease() MTRPowerSourceClusterWiredFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPowerSourceClusterWiredFaultChangeType creates a new MTRPowerSourceClusterWiredFaultChangeType instance.
func NewMTRPowerSourceClusterWiredFaultChangeType() MTRPowerSourceClusterWiredFaultChangeType {
	return getMTRPowerSourceClusterWiredFaultChangeTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPowerSourceClusterWiredFaultChangeType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterWiredFaultChangeType
type MTRPowerSourceClusterWiredFaultChangeType struct {
	objectivec.Object
}

// MTRPowerSourceClusterWiredFaultChangeTypeFrom constructs a [MTRPowerSourceClusterWiredFaultChangeType] from an unsafe.Pointer.
func MTRPowerSourceClusterWiredFaultChangeTypeFrom(ptr unsafe.Pointer) MTRPowerSourceClusterWiredFaultChangeType {
	return MTRPowerSourceClusterWiredFaultChangeType{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPowerSourceClusterWiredFaultChangeType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPowerSourceClusterWiredFaultChangeType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPowerSourceClusterWiredFaultChangeType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPowerSourceClusterWiredFaultChangeType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPowerSourceClusterWiredFaultChangeType */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterWiredFaultChangeType/current
func (m_ MTRPowerSourceClusterWiredFaultChangeType) Current() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("current"))
	return rv
}/* debug [instance_properties/getter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterWiredFaultChangeType/current
func (m_ MTRPowerSourceClusterWiredFaultChangeType) SetCurrent(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}/* debug [instance_properties/setter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterWiredFaultChangeType/previous
func (m_ MTRPowerSourceClusterWiredFaultChangeType) Previous() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("previous"))
	return rv
}/* debug [instance_properties/getter]: previous */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterWiredFaultChangeType/previous
func (m_ MTRPowerSourceClusterWiredFaultChangeType) SetPrevious(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}/* debug [instance_properties/setter]: previous */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPowerSourceClusterWiredFaultChangeType */



