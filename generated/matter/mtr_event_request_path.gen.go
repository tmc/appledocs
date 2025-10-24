// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREventRequestPath */


/* debug [class_header]: Header for MTREventRequestPath */
// The class instance for the [MTREventRequestPath] class.
var (
	MTREventRequestPathClass     _MTREventRequestPathClass
	MTREventRequestPathClassOnce sync.Once
)

func getMTREventRequestPathClass() _MTREventRequestPathClass {
	MTREventRequestPathClassOnce.Do(func() {
		MTREventRequestPathClass = _MTREventRequestPathClass{objc.GetClass("MTREventRequestPath")}
	})
	return MTREventRequestPathClass
}

type _MTREventRequestPathClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREventRequestPath */
// An interface definition for the [MTREventRequestPath] class.
type IMTREventRequestPath interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREventRequestPath */
	// properties:
	Cluster() objc.IObject /* cross-framework: NSNumber */
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	Event() objc.IObject /* cross-framework: NSNumber */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREventRequestPath */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREventRequestPath */
// Alloc allocates a new instance without initialization.
func (mc _MTREventRequestPathClass) Alloc() MTREventRequestPath {
	rv := objc.Send[MTREventRequestPath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREventRequestPathClass) New() MTREventRequestPath {
	rv := objc.Send[MTREventRequestPath](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREventRequestPath) Init() MTREventRequestPath {
	rv := objc.Send[MTREventRequestPath](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREventRequestPath) Autorelease() MTREventRequestPath {
	rv := objc.Send[MTREventRequestPath](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREventRequestPath creates a new MTREventRequestPath instance.
func NewMTREventRequestPath() MTREventRequestPath {
	return getMTREventRequestPathClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREventRequestPath */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventRequestPath
type MTREventRequestPath struct {
	objectivec.Object
}

// MTREventRequestPathFrom constructs a [MTREventRequestPath] from an unsafe.Pointer.
func MTREventRequestPathFrom(ptr unsafe.Pointer) MTREventRequestPath {
	return MTREventRequestPath{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREventRequestPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventRequestPath/init(endpointID:clusterID:eventID:)
func NewMTREventRequestPathWithEndpointIDClusterIDEventID(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, eventID objc.IObject /* cross-framework: NSNumber */) MTREventRequestPath {
	rv := objc.Send[MTREventRequestPath](objc.ID(getMTREventRequestPathClass().class), objc.Sel("requestPathWithEndpointID:clusterID:eventID:"), endpointID, clusterID, eventID)
	return rv
}/* debug [class_init_methods/constructor]: NewMTREventRequestPathWithEndpointIDClusterIDEventID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREventRequestPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventRequestPath/init(endpointID:clusterID:eventID:)
func (mc _MTREventRequestPathClass) RequestPathWithEndpointIDClusterIDEventID(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, eventID objc.IObject /* cross-framework: NSNumber */) MTREventRequestPath {
	rv := objc.Send[MTREventRequestPath](objc.ID(mc.class), objc.Sel("requestPathWithEndpointID:clusterID:eventID:"), endpointID, clusterID, eventID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestPathWithEndpointIDClusterIDEventID) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREventRequestPath */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREventRequestPath */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREventRequestPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventRequestPath/cluster
func (m_ MTREventRequestPath) Cluster() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cluster"))
	return rv
}/* debug [instance_properties/getter]: cluster */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventRequestPath/endpoint
func (m_ MTREventRequestPath) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventRequestPath/event
func (m_ MTREventRequestPath) Event() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("event"))
	return rv
}/* debug [instance_properties/getter]: event */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREventRequestPath */


