// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterTimeZoneType */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterTimeZoneType */
// The class instance for the [MTRTimeSynchronizationClusterTimeZoneType] class.
var (
	MTRTimeSynchronizationClusterTimeZoneTypeClass     _MTRTimeSynchronizationClusterTimeZoneTypeClass
	MTRTimeSynchronizationClusterTimeZoneTypeClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterTimeZoneTypeClass() _MTRTimeSynchronizationClusterTimeZoneTypeClass {
	MTRTimeSynchronizationClusterTimeZoneTypeClassOnce.Do(func() {
		MTRTimeSynchronizationClusterTimeZoneTypeClass = _MTRTimeSynchronizationClusterTimeZoneTypeClass{objc.GetClass("MTRTimeSynchronizationClusterTimeZoneType")}
	})
	return MTRTimeSynchronizationClusterTimeZoneTypeClass
}

type _MTRTimeSynchronizationClusterTimeZoneTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterTimeZoneType */
// An interface definition for the [MTRTimeSynchronizationClusterTimeZoneType] class.
type IMTRTimeSynchronizationClusterTimeZoneType interface {
	IMTRTimeSynchronizationClusterTimeZoneStruct
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterTimeZoneType */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Offset() objc.IObject /* cross-framework: NSNumber */
	SetOffset(value objc.IObject /* cross-framework: NSNumber */)
	ValidAt() objc.IObject /* cross-framework: NSNumber */
	SetValidAt(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterTimeZoneType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterTimeZoneType */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterTimeZoneTypeClass) Alloc() MTRTimeSynchronizationClusterTimeZoneType {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterTimeZoneTypeClass) New() MTRTimeSynchronizationClusterTimeZoneType {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterTimeZoneType) Init() MTRTimeSynchronizationClusterTimeZoneType {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterTimeZoneType) Autorelease() MTRTimeSynchronizationClusterTimeZoneType {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterTimeZoneType creates a new MTRTimeSynchronizationClusterTimeZoneType instance.
func NewMTRTimeSynchronizationClusterTimeZoneType() MTRTimeSynchronizationClusterTimeZoneType {
	return getMTRTimeSynchronizationClusterTimeZoneTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterTimeZoneType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneType
type MTRTimeSynchronizationClusterTimeZoneType struct {
	MTRTimeSynchronizationClusterTimeZoneStruct
}

// MTRTimeSynchronizationClusterTimeZoneTypeFrom constructs a [MTRTimeSynchronizationClusterTimeZoneType] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterTimeZoneTypeFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterTimeZoneType {
	return MTRTimeSynchronizationClusterTimeZoneType{
		MTRTimeSynchronizationClusterTimeZoneStruct: MTRTimeSynchronizationClusterTimeZoneStructFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterTimeZoneType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterTimeZoneType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterTimeZoneType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterTimeZoneType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterTimeZoneType */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneType/name
func (m_ MTRTimeSynchronizationClusterTimeZoneType) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneType/name
func (m_ MTRTimeSynchronizationClusterTimeZoneType) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneType/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneType) Offset() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneType/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneType) SetOffset(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneType/validAt
func (m_ MTRTimeSynchronizationClusterTimeZoneType) ValidAt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("validAt"))
	return rv
}/* debug [instance_properties/getter]: validAt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneType/validAt
func (m_ MTRTimeSynchronizationClusterTimeZoneType) SetValidAt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValidAt:"), value)
}/* debug [instance_properties/setter]: validAt */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterTimeZoneType */



