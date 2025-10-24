// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterTimeZoneStruct */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterTimeZoneStruct */
// The class instance for the [MTRTimeSynchronizationClusterTimeZoneStruct] class.
var (
	MTRTimeSynchronizationClusterTimeZoneStructClass     _MTRTimeSynchronizationClusterTimeZoneStructClass
	MTRTimeSynchronizationClusterTimeZoneStructClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterTimeZoneStructClass() _MTRTimeSynchronizationClusterTimeZoneStructClass {
	MTRTimeSynchronizationClusterTimeZoneStructClassOnce.Do(func() {
		MTRTimeSynchronizationClusterTimeZoneStructClass = _MTRTimeSynchronizationClusterTimeZoneStructClass{objc.GetClass("MTRTimeSynchronizationClusterTimeZoneStruct")}
	})
	return MTRTimeSynchronizationClusterTimeZoneStructClass
}

type _MTRTimeSynchronizationClusterTimeZoneStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterTimeZoneStruct */
// An interface definition for the [MTRTimeSynchronizationClusterTimeZoneStruct] class.
type IMTRTimeSynchronizationClusterTimeZoneStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterTimeZoneStruct */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Offset() objc.IObject /* cross-framework: NSNumber */
	SetOffset(value objc.IObject /* cross-framework: NSNumber */)
	ValidAt() objc.IObject /* cross-framework: NSNumber */
	SetValidAt(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterTimeZoneStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterTimeZoneStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterTimeZoneStructClass) Alloc() MTRTimeSynchronizationClusterTimeZoneStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterTimeZoneStructClass) New() MTRTimeSynchronizationClusterTimeZoneStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) Init() MTRTimeSynchronizationClusterTimeZoneStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) Autorelease() MTRTimeSynchronizationClusterTimeZoneStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterTimeZoneStruct creates a new MTRTimeSynchronizationClusterTimeZoneStruct instance.
func NewMTRTimeSynchronizationClusterTimeZoneStruct() MTRTimeSynchronizationClusterTimeZoneStruct {
	return getMTRTimeSynchronizationClusterTimeZoneStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterTimeZoneStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStruct
type MTRTimeSynchronizationClusterTimeZoneStruct struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterTimeZoneStructFrom constructs a [MTRTimeSynchronizationClusterTimeZoneStruct] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterTimeZoneStructFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterTimeZoneStruct {
	return MTRTimeSynchronizationClusterTimeZoneStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterTimeZoneStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterTimeZoneStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterTimeZoneStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterTimeZoneStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterTimeZoneStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStruct/name
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStruct/name
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStruct/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) Offset() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStruct/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) SetOffset(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStruct/validAt
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) ValidAt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("validAt"))
	return rv
}/* debug [instance_properties/getter]: validAt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStruct/validAt
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) SetValidAt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValidAt:"), value)
}/* debug [instance_properties/setter]: validAt */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterTimeZoneStruct */



