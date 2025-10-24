// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRAccessControlClusterFabricRestrictionReviewUpdateEvent */


/* debug [class_header]: Header for MTRAccessControlClusterFabricRestrictionReviewUpdateEvent */
// The class instance for the [MTRAccessControlClusterFabricRestrictionReviewUpdateEvent] class.
var (
	MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass     _MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass
	MTRAccessControlClusterFabricRestrictionReviewUpdateEventClassOnce sync.Once
)

func getMTRAccessControlClusterFabricRestrictionReviewUpdateEventClass() _MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass {
	MTRAccessControlClusterFabricRestrictionReviewUpdateEventClassOnce.Do(func() {
		MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass = _MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass{objc.GetClass("MTRAccessControlClusterFabricRestrictionReviewUpdateEvent")}
	})
	return MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass
}

type _MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAccessControlClusterFabricRestrictionReviewUpdateEvent */
// An interface definition for the [MTRAccessControlClusterFabricRestrictionReviewUpdateEvent] class.
type IMTRAccessControlClusterFabricRestrictionReviewUpdateEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAccessControlClusterFabricRestrictionReviewUpdateEvent */
	// properties:
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	Instruction() objc.IObject /* cross-framework: NSString */
	SetInstruction(value objc.IObject /* cross-framework: NSString */)
	Token() objc.IObject /* cross-framework: NSNumber */
	SetToken(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAccessControlClusterFabricRestrictionReviewUpdateEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAccessControlClusterFabricRestrictionReviewUpdateEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass) Alloc() MTRAccessControlClusterFabricRestrictionReviewUpdateEvent {
	rv := objc.Send[MTRAccessControlClusterFabricRestrictionReviewUpdateEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass) New() MTRAccessControlClusterFabricRestrictionReviewUpdateEvent {
	rv := objc.Send[MTRAccessControlClusterFabricRestrictionReviewUpdateEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) Init() MTRAccessControlClusterFabricRestrictionReviewUpdateEvent {
	rv := objc.Send[MTRAccessControlClusterFabricRestrictionReviewUpdateEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) Autorelease() MTRAccessControlClusterFabricRestrictionReviewUpdateEvent {
	rv := objc.Send[MTRAccessControlClusterFabricRestrictionReviewUpdateEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterFabricRestrictionReviewUpdateEvent creates a new MTRAccessControlClusterFabricRestrictionReviewUpdateEvent instance.
func NewMTRAccessControlClusterFabricRestrictionReviewUpdateEvent() MTRAccessControlClusterFabricRestrictionReviewUpdateEvent {
	return getMTRAccessControlClusterFabricRestrictionReviewUpdateEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAccessControlClusterFabricRestrictionReviewUpdateEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterFabricRestrictionReviewUpdateEvent
type MTRAccessControlClusterFabricRestrictionReviewUpdateEvent struct {
	objectivec.Object
}

// MTRAccessControlClusterFabricRestrictionReviewUpdateEventFrom constructs a [MTRAccessControlClusterFabricRestrictionReviewUpdateEvent] from an unsafe.Pointer.
func MTRAccessControlClusterFabricRestrictionReviewUpdateEventFrom(ptr unsafe.Pointer) MTRAccessControlClusterFabricRestrictionReviewUpdateEvent {
	return MTRAccessControlClusterFabricRestrictionReviewUpdateEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAccessControlClusterFabricRestrictionReviewUpdateEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAccessControlClusterFabricRestrictionReviewUpdateEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAccessControlClusterFabricRestrictionReviewUpdateEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAccessControlClusterFabricRestrictionReviewUpdateEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAccessControlClusterFabricRestrictionReviewUpdateEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterFabricRestrictionReviewUpdateEvent/fabricIndex
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterFabricRestrictionReviewUpdateEvent/fabricIndex
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusterfabricrestrictionreviewupdateevent/instruction
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) Instruction() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("instruction"))
	return rv
}/* debug [instance_properties/getter]: instruction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusterfabricrestrictionreviewupdateevent/instruction
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) SetInstruction(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstruction:"), value)
}/* debug [instance_properties/setter]: instruction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusterfabricrestrictionreviewupdateevent/token
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) Token() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("token"))
	return rv
}/* debug [instance_properties/getter]: token */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusterfabricrestrictionreviewupdateevent/token
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) SetToken(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setToken:"), value)
}/* debug [instance_properties/setter]: token */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAccessControlClusterFabricRestrictionReviewUpdateEvent */



