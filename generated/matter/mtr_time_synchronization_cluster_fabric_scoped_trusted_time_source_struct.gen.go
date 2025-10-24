// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct */
// The class instance for the [MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct] class.
var (
	MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass     _MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass
	MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass() _MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass {
	MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClassOnce.Do(func() {
		MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass = _MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass{objc.GetClass("MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct")}
	})
	return MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass
}

type _MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct */
// An interface definition for the [MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct] class.
type IMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct */
	// properties:
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	NodeID() objc.IObject /* cross-framework: NSNumber */
	SetNodeID(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass) Alloc() MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass) New() MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) Init() MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) Autorelease() MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct creates a new MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct instance.
func NewMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct() MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	return getMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct
type MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructFrom constructs a [MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	return MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct/endpoint
func (m_ MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct/endpoint
func (m_ MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}/* debug [instance_properties/setter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclusterfabricscopedtrustedtimesourcestruct/nodeid
func (m_ MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) NodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeID"))
	return rv
}/* debug [instance_properties/getter]: nodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclusterfabricscopedtrustedtimesourcestruct/nodeid
func (m_ MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) SetNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}/* debug [instance_properties/setter]: nodeID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct */



