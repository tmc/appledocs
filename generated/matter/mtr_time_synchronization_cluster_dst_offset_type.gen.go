// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterDstOffsetType */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterDstOffsetType */
// The class instance for the [MTRTimeSynchronizationClusterDstOffsetType] class.
var (
	MTRTimeSynchronizationClusterDstOffsetTypeClass     _MTRTimeSynchronizationClusterDstOffsetTypeClass
	MTRTimeSynchronizationClusterDstOffsetTypeClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterDstOffsetTypeClass() _MTRTimeSynchronizationClusterDstOffsetTypeClass {
	MTRTimeSynchronizationClusterDstOffsetTypeClassOnce.Do(func() {
		MTRTimeSynchronizationClusterDstOffsetTypeClass = _MTRTimeSynchronizationClusterDstOffsetTypeClass{objc.GetClass("MTRTimeSynchronizationClusterDstOffsetType")}
	})
	return MTRTimeSynchronizationClusterDstOffsetTypeClass
}

type _MTRTimeSynchronizationClusterDstOffsetTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterDstOffsetType */
// An interface definition for the [MTRTimeSynchronizationClusterDstOffsetType] class.
type IMTRTimeSynchronizationClusterDstOffsetType interface {
	IMTRTimeSynchronizationClusterDSTOffsetStruct
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterDstOffsetType */
	// properties:
	Offset() objc.IObject /* cross-framework: NSNumber */
	SetOffset(value objc.IObject /* cross-framework: NSNumber */)
	ValidStarting() objc.IObject /* cross-framework: NSNumber */
	SetValidStarting(value objc.IObject /* cross-framework: NSNumber */)
	ValidUntil() objc.IObject /* cross-framework: NSNumber */
	SetValidUntil(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterDstOffsetType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterDstOffsetType */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterDstOffsetTypeClass) Alloc() MTRTimeSynchronizationClusterDstOffsetType {
	rv := objc.Send[MTRTimeSynchronizationClusterDstOffsetType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterDstOffsetTypeClass) New() MTRTimeSynchronizationClusterDstOffsetType {
	rv := objc.Send[MTRTimeSynchronizationClusterDstOffsetType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterDstOffsetType) Init() MTRTimeSynchronizationClusterDstOffsetType {
	rv := objc.Send[MTRTimeSynchronizationClusterDstOffsetType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterDstOffsetType) Autorelease() MTRTimeSynchronizationClusterDstOffsetType {
	rv := objc.Send[MTRTimeSynchronizationClusterDstOffsetType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterDstOffsetType creates a new MTRTimeSynchronizationClusterDstOffsetType instance.
func NewMTRTimeSynchronizationClusterDstOffsetType() MTRTimeSynchronizationClusterDstOffsetType {
	return getMTRTimeSynchronizationClusterDstOffsetTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterDstOffsetType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDstOffsetType
type MTRTimeSynchronizationClusterDstOffsetType struct {
	MTRTimeSynchronizationClusterDSTOffsetStruct
}

// MTRTimeSynchronizationClusterDstOffsetTypeFrom constructs a [MTRTimeSynchronizationClusterDstOffsetType] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterDstOffsetTypeFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterDstOffsetType {
	return MTRTimeSynchronizationClusterDstOffsetType{
		MTRTimeSynchronizationClusterDSTOffsetStruct: MTRTimeSynchronizationClusterDSTOffsetStructFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterDstOffsetType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterDstOffsetType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterDstOffsetType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterDstOffsetType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterDstOffsetType */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDstOffsetType/offset
func (m_ MTRTimeSynchronizationClusterDstOffsetType) Offset() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDstOffsetType/offset
func (m_ MTRTimeSynchronizationClusterDstOffsetType) SetOffset(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDstOffsetType/validStarting
func (m_ MTRTimeSynchronizationClusterDstOffsetType) ValidStarting() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("validStarting"))
	return rv
}/* debug [instance_properties/getter]: validStarting */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDstOffsetType/validStarting
func (m_ MTRTimeSynchronizationClusterDstOffsetType) SetValidStarting(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValidStarting:"), value)
}/* debug [instance_properties/setter]: validStarting */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDstOffsetType/validUntil
func (m_ MTRTimeSynchronizationClusterDstOffsetType) ValidUntil() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("validUntil"))
	return rv
}/* debug [instance_properties/getter]: validUntil */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDstOffsetType/validUntil
func (m_ MTRTimeSynchronizationClusterDstOffsetType) SetValidUntil(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValidUntil:"), value)
}/* debug [instance_properties/setter]: validUntil */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterDstOffsetType */



