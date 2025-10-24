// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTREventPath */


/* debug [class_header]: Header for MTREventPath */
// The class instance for the [MTREventPath] class.
var (
	MTREventPathClass     _MTREventPathClass
	MTREventPathClassOnce sync.Once
)

func getMTREventPathClass() _MTREventPathClass {
	MTREventPathClassOnce.Do(func() {
		MTREventPathClass = _MTREventPathClass{objc.GetClass("MTREventPath")}
	})
	return MTREventPathClass
}

type _MTREventPathClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREventPath */
// An interface definition for the [MTREventPath] class.
type IMTREventPath interface {
	IMTRClusterPath
	
/* debug [class_interface_properties]: Properties for MTREventPath */
	// properties:
	Event() objc.IObject /* cross-framework: NSNumber */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREventPath */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREventPath */
// Alloc allocates a new instance without initialization.
func (mc _MTREventPathClass) Alloc() MTREventPath {
	rv := objc.Send[MTREventPath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREventPathClass) New() MTREventPath {
	rv := objc.Send[MTREventPath](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREventPath) Init() MTREventPath {
	rv := objc.Send[MTREventPath](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREventPath) Autorelease() MTREventPath {
	rv := objc.Send[MTREventPath](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREventPath creates a new MTREventPath instance.
func NewMTREventPath() MTREventPath {
	return getMTREventPathClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREventPath */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventPath
type MTREventPath struct {
	MTRClusterPath
}

// MTREventPathFrom constructs a [MTREventPath] from an unsafe.Pointer.
func MTREventPathFrom(ptr unsafe.Pointer) MTREventPath {
	return MTREventPath{
		MTRClusterPath: MTRClusterPathFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREventPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventPath/init(endpointID:clusterID:eventID:)-8cd6r
func NewMTREventPathWithEndpointIDClusterIDEventID(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, eventID objc.IObject /* cross-framework: NSNumber */) MTREventPath {
	rv := objc.Send[MTREventPath](objc.ID(getMTREventPathClass().class), objc.Sel("eventPathWithEndpointID:clusterID:eventID:"), endpointID, clusterID, eventID)
	return rv
}/* debug [class_init_methods/constructor]: NewMTREventPathWithEndpointIDClusterIDEventID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventPath/init(endpointId:clusterId:eventId:)-2i9w4
func NewMTREventPathWithEndpointIdClusterIdEventId(endpointId objc.IObject /* cross-framework: NSNumber */, clusterId objc.IObject /* cross-framework: NSNumber */, eventId objc.IObject /* cross-framework: NSNumber */) MTREventPath {
	rv := objc.Send[MTREventPath](objc.ID(getMTREventPathClass().class), objc.Sel("eventPathWithEndpointId:clusterId:eventId:"), endpointId, clusterId, eventId)
	return rv
}/* debug [class_init_methods/constructor]: NewMTREventPathWithEndpointIdClusterIdEventId */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREventPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventPath/init(endpointId:clusterId:eventId:)-2i9w4
func (mc _MTREventPathClass) EventPathWithEndpointIdClusterIdEventId(endpointId objc.IObject /* cross-framework: NSNumber */, clusterId objc.IObject /* cross-framework: NSNumber */, eventId objc.IObject /* cross-framework: NSNumber */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("eventPathWithEndpointId:clusterId:eventId:"), endpointId, clusterId, eventId)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EventPathWithEndpointIdClusterIdEventId) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventPath/init(endpointID:clusterID:eventID:)-8cd6r
func (mc _MTREventPathClass) EventPathWithEndpointIDClusterIDEventID(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, eventID objc.IObject /* cross-framework: NSNumber */) MTREventPath {
	rv := objc.Send[MTREventPath](objc.ID(mc.class), objc.Sel("eventPathWithEndpointID:clusterID:eventID:"), endpointID, clusterID, eventID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EventPathWithEndpointIDClusterIDEventID) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREventPath */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREventPath */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREventPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventPath/event
func (m_ MTREventPath) Event() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("event"))
	return rv
}/* debug [instance_properties/getter]: event */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREventPath */


