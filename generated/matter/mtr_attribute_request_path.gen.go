// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRAttributeRequestPath */


/* debug [class_header]: Header for MTRAttributeRequestPath */
// The class instance for the [MTRAttributeRequestPath] class.
var (
	MTRAttributeRequestPathClass     _MTRAttributeRequestPathClass
	MTRAttributeRequestPathClassOnce sync.Once
)

func getMTRAttributeRequestPathClass() _MTRAttributeRequestPathClass {
	MTRAttributeRequestPathClassOnce.Do(func() {
		MTRAttributeRequestPathClass = _MTRAttributeRequestPathClass{objc.GetClass("MTRAttributeRequestPath")}
	})
	return MTRAttributeRequestPathClass
}

type _MTRAttributeRequestPathClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAttributeRequestPath */
// An interface definition for the [MTRAttributeRequestPath] class.
type IMTRAttributeRequestPath interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAttributeRequestPath */
	// properties:
	Attribute() objc.IObject /* cross-framework: NSNumber */
	Cluster() objc.IObject /* cross-framework: NSNumber */
	Endpoint() objc.IObject /* cross-framework: NSNumber */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAttributeRequestPath */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAttributeRequestPath */
// Alloc allocates a new instance without initialization.
func (mc _MTRAttributeRequestPathClass) Alloc() MTRAttributeRequestPath {
	rv := objc.Send[MTRAttributeRequestPath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAttributeRequestPathClass) New() MTRAttributeRequestPath {
	rv := objc.Send[MTRAttributeRequestPath](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAttributeRequestPath) Init() MTRAttributeRequestPath {
	rv := objc.Send[MTRAttributeRequestPath](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAttributeRequestPath) Autorelease() MTRAttributeRequestPath {
	rv := objc.Send[MTRAttributeRequestPath](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAttributeRequestPath creates a new MTRAttributeRequestPath instance.
func NewMTRAttributeRequestPath() MTRAttributeRequestPath {
	return getMTRAttributeRequestPathClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAttributeRequestPath */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeRequestPath
type MTRAttributeRequestPath struct {
	objectivec.Object
}

// MTRAttributeRequestPathFrom constructs a [MTRAttributeRequestPath] from an unsafe.Pointer.
func MTRAttributeRequestPathFrom(ptr unsafe.Pointer) MTRAttributeRequestPath {
	return MTRAttributeRequestPath{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAttributeRequestPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeRequestPath/init(endpointID:clusterID:attributeID:)
func NewMTRAttributeRequestPathWithEndpointIDClusterIDAttributeID(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */) MTRAttributeRequestPath {
	rv := objc.Send[MTRAttributeRequestPath](objc.ID(getMTRAttributeRequestPathClass().class), objc.Sel("requestPathWithEndpointID:clusterID:attributeID:"), endpointID, clusterID, attributeID)
	return rv
}/* debug [class_init_methods/constructor]: NewMTRAttributeRequestPathWithEndpointIDClusterIDAttributeID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAttributeRequestPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeRequestPath/init(endpointID:clusterID:attributeID:)
func (mc _MTRAttributeRequestPathClass) RequestPathWithEndpointIDClusterIDAttributeID(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */) MTRAttributeRequestPath {
	rv := objc.Send[MTRAttributeRequestPath](objc.ID(mc.class), objc.Sel("requestPathWithEndpointID:clusterID:attributeID:"), endpointID, clusterID, attributeID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestPathWithEndpointIDClusterIDAttributeID) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAttributeRequestPath */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAttributeRequestPath */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAttributeRequestPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeRequestPath/attribute
func (m_ MTRAttributeRequestPath) Attribute() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("attribute"))
	return rv
}/* debug [instance_properties/getter]: attribute */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeRequestPath/cluster
func (m_ MTRAttributeRequestPath) Cluster() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cluster"))
	return rv
}/* debug [instance_properties/getter]: cluster */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeRequestPath/endpoint
func (m_ MTRAttributeRequestPath) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAttributeRequestPath */


