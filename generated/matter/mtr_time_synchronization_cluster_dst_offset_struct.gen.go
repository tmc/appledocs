// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterDSTOffsetStruct */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterDSTOffsetStruct */
// The class instance for the [MTRTimeSynchronizationClusterDSTOffsetStruct] class.
var (
	MTRTimeSynchronizationClusterDSTOffsetStructClass     _MTRTimeSynchronizationClusterDSTOffsetStructClass
	MTRTimeSynchronizationClusterDSTOffsetStructClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterDSTOffsetStructClass() _MTRTimeSynchronizationClusterDSTOffsetStructClass {
	MTRTimeSynchronizationClusterDSTOffsetStructClassOnce.Do(func() {
		MTRTimeSynchronizationClusterDSTOffsetStructClass = _MTRTimeSynchronizationClusterDSTOffsetStructClass{objc.GetClass("MTRTimeSynchronizationClusterDSTOffsetStruct")}
	})
	return MTRTimeSynchronizationClusterDSTOffsetStructClass
}

type _MTRTimeSynchronizationClusterDSTOffsetStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterDSTOffsetStruct */
// An interface definition for the [MTRTimeSynchronizationClusterDSTOffsetStruct] class.
type IMTRTimeSynchronizationClusterDSTOffsetStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterDSTOffsetStruct */
	// properties:
	Offset() objc.IObject /* cross-framework: NSNumber */
	SetOffset(value objc.IObject /* cross-framework: NSNumber */)
	ValidStarting() objc.IObject /* cross-framework: NSNumber */
	SetValidStarting(value objc.IObject /* cross-framework: NSNumber */)
	ValidUntil() objc.IObject /* cross-framework: NSNumber */
	SetValidUntil(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterDSTOffsetStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterDSTOffsetStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterDSTOffsetStructClass) Alloc() MTRTimeSynchronizationClusterDSTOffsetStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTOffsetStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterDSTOffsetStructClass) New() MTRTimeSynchronizationClusterDSTOffsetStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTOffsetStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterDSTOffsetStruct) Init() MTRTimeSynchronizationClusterDSTOffsetStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTOffsetStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterDSTOffsetStruct) Autorelease() MTRTimeSynchronizationClusterDSTOffsetStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTOffsetStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterDSTOffsetStruct creates a new MTRTimeSynchronizationClusterDSTOffsetStruct instance.
func NewMTRTimeSynchronizationClusterDSTOffsetStruct() MTRTimeSynchronizationClusterDSTOffsetStruct {
	return getMTRTimeSynchronizationClusterDSTOffsetStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterDSTOffsetStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTOffsetStruct
type MTRTimeSynchronizationClusterDSTOffsetStruct struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterDSTOffsetStructFrom constructs a [MTRTimeSynchronizationClusterDSTOffsetStruct] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterDSTOffsetStructFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterDSTOffsetStruct {
	return MTRTimeSynchronizationClusterDSTOffsetStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterDSTOffsetStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterDSTOffsetStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterDSTOffsetStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterDSTOffsetStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterDSTOffsetStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTOffsetStruct/offset
func (m_ MTRTimeSynchronizationClusterDSTOffsetStruct) Offset() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTOffsetStruct/offset
func (m_ MTRTimeSynchronizationClusterDSTOffsetStruct) SetOffset(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTOffsetStruct/validStarting
func (m_ MTRTimeSynchronizationClusterDSTOffsetStruct) ValidStarting() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("validStarting"))
	return rv
}/* debug [instance_properties/getter]: validStarting */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTOffsetStruct/validStarting
func (m_ MTRTimeSynchronizationClusterDSTOffsetStruct) SetValidStarting(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValidStarting:"), value)
}/* debug [instance_properties/setter]: validStarting */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTOffsetStruct/validUntil
func (m_ MTRTimeSynchronizationClusterDSTOffsetStruct) ValidUntil() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("validUntil"))
	return rv
}/* debug [instance_properties/getter]: validUntil */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTOffsetStruct/validUntil
func (m_ MTRTimeSynchronizationClusterDSTOffsetStruct) SetValidUntil(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValidUntil:"), value)
}/* debug [instance_properties/setter]: validUntil */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterDSTOffsetStruct */



