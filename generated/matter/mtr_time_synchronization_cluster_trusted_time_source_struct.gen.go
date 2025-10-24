// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterTrustedTimeSourceStruct */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterTrustedTimeSourceStruct */
// The class instance for the [MTRTimeSynchronizationClusterTrustedTimeSourceStruct] class.
var (
	MTRTimeSynchronizationClusterTrustedTimeSourceStructClass     _MTRTimeSynchronizationClusterTrustedTimeSourceStructClass
	MTRTimeSynchronizationClusterTrustedTimeSourceStructClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterTrustedTimeSourceStructClass() _MTRTimeSynchronizationClusterTrustedTimeSourceStructClass {
	MTRTimeSynchronizationClusterTrustedTimeSourceStructClassOnce.Do(func() {
		MTRTimeSynchronizationClusterTrustedTimeSourceStructClass = _MTRTimeSynchronizationClusterTrustedTimeSourceStructClass{objc.GetClass("MTRTimeSynchronizationClusterTrustedTimeSourceStruct")}
	})
	return MTRTimeSynchronizationClusterTrustedTimeSourceStructClass
}

type _MTRTimeSynchronizationClusterTrustedTimeSourceStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterTrustedTimeSourceStruct */
// An interface definition for the [MTRTimeSynchronizationClusterTrustedTimeSourceStruct] class.
type IMTRTimeSynchronizationClusterTrustedTimeSourceStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterTrustedTimeSourceStruct */
	// properties:
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	NodeID() objc.IObject /* cross-framework: NSNumber */
	SetNodeID(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterTrustedTimeSourceStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterTrustedTimeSourceStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterTrustedTimeSourceStructClass) Alloc() MTRTimeSynchronizationClusterTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTrustedTimeSourceStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterTrustedTimeSourceStructClass) New() MTRTimeSynchronizationClusterTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTrustedTimeSourceStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) Init() MTRTimeSynchronizationClusterTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTrustedTimeSourceStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) Autorelease() MTRTimeSynchronizationClusterTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTrustedTimeSourceStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterTrustedTimeSourceStruct creates a new MTRTimeSynchronizationClusterTrustedTimeSourceStruct instance.
func NewMTRTimeSynchronizationClusterTrustedTimeSourceStruct() MTRTimeSynchronizationClusterTrustedTimeSourceStruct {
	return getMTRTimeSynchronizationClusterTrustedTimeSourceStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterTrustedTimeSourceStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTrustedTimeSourceStruct
type MTRTimeSynchronizationClusterTrustedTimeSourceStruct struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterTrustedTimeSourceStructFrom constructs a [MTRTimeSynchronizationClusterTrustedTimeSourceStruct] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterTrustedTimeSourceStructFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterTrustedTimeSourceStruct {
	return MTRTimeSynchronizationClusterTrustedTimeSourceStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterTrustedTimeSourceStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterTrustedTimeSourceStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterTrustedTimeSourceStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterTrustedTimeSourceStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterTrustedTimeSourceStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTrustedTimeSourceStruct/fabricIndex
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTrustedTimeSourceStruct/fabricIndex
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertrustedtimesourcestruct/endpoint
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertrustedtimesourcestruct/endpoint
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}/* debug [instance_properties/setter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertrustedtimesourcestruct/nodeid
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) NodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeID"))
	return rv
}/* debug [instance_properties/getter]: nodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertrustedtimesourcestruct/nodeid
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) SetNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}/* debug [instance_properties/setter]: nodeID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterTrustedTimeSourceStruct */



