// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRClusterPath */


/* debug [class_header]: Header for MTRClusterPath */
// The class instance for the [MTRClusterPath] class.
var (
	MTRClusterPathClass     _MTRClusterPathClass
	MTRClusterPathClassOnce sync.Once
)

func getMTRClusterPathClass() _MTRClusterPathClass {
	MTRClusterPathClassOnce.Do(func() {
		MTRClusterPathClass = _MTRClusterPathClass{objc.GetClass("MTRClusterPath")}
	})
	return MTRClusterPathClass
}

type _MTRClusterPathClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterPath */
// An interface definition for the [MTRClusterPath] class.
type IMTRClusterPath interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRClusterPath */
	// properties:
	Cluster() objc.IObject /* cross-framework: NSNumber */
	Endpoint() objc.IObject /* cross-framework: NSNumber */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterPath */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterPath */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterPathClass) Alloc() MTRClusterPath {
	rv := objc.Send[MTRClusterPath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterPathClass) New() MTRClusterPath {
	rv := objc.Send[MTRClusterPath](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterPath) Init() MTRClusterPath {
	rv := objc.Send[MTRClusterPath](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterPath) Autorelease() MTRClusterPath {
	rv := objc.Send[MTRClusterPath](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterPath creates a new MTRClusterPath instance.
func NewMTRClusterPath() MTRClusterPath {
	return getMTRClusterPathClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterPath */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPath
type MTRClusterPath struct {
	objectivec.Object
}

// MTRClusterPathFrom constructs a [MTRClusterPath] from an unsafe.Pointer.
func MTRClusterPathFrom(ptr unsafe.Pointer) MTRClusterPath {
	return MTRClusterPath{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPath/init(endpointID:clusterID:)
func NewMTRClusterPathWithEndpointIDClusterID(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */) MTRClusterPath {
	rv := objc.Send[MTRClusterPath](objc.ID(getMTRClusterPathClass().class), objc.Sel("clusterPathWithEndpointID:clusterID:"), endpointID, clusterID)
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterPathWithEndpointIDClusterID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPath/init(endpointID:clusterID:)
func (mc _MTRClusterPathClass) ClusterPathWithEndpointIDClusterID(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */) MTRClusterPath {
	rv := objc.Send[MTRClusterPath](objc.ID(mc.class), objc.Sel("clusterPathWithEndpointID:clusterID:"), endpointID, clusterID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ClusterPathWithEndpointIDClusterID) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterPath */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterPath */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPath/cluster
func (m_ MTRClusterPath) Cluster() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cluster"))
	return rv
}/* debug [instance_properties/getter]: cluster */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPath/endpoint
func (m_ MTRClusterPath) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterPath */


