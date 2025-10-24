// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRAttributePath */


/* debug [class_header]: Header for MTRAttributePath */
// The class instance for the [MTRAttributePath] class.
var (
	MTRAttributePathClass     _MTRAttributePathClass
	MTRAttributePathClassOnce sync.Once
)

func getMTRAttributePathClass() _MTRAttributePathClass {
	MTRAttributePathClassOnce.Do(func() {
		MTRAttributePathClass = _MTRAttributePathClass{objc.GetClass("MTRAttributePath")}
	})
	return MTRAttributePathClass
}

type _MTRAttributePathClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAttributePath */
// An interface definition for the [MTRAttributePath] class.
type IMTRAttributePath interface {
	IMTRClusterPath
	
/* debug [class_interface_properties]: Properties for MTRAttributePath */
	// properties:
	Attribute() objc.IObject /* cross-framework: NSNumber */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAttributePath */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAttributePath */
// Alloc allocates a new instance without initialization.
func (mc _MTRAttributePathClass) Alloc() MTRAttributePath {
	rv := objc.Send[MTRAttributePath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAttributePathClass) New() MTRAttributePath {
	rv := objc.Send[MTRAttributePath](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAttributePath) Init() MTRAttributePath {
	rv := objc.Send[MTRAttributePath](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAttributePath) Autorelease() MTRAttributePath {
	rv := objc.Send[MTRAttributePath](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAttributePath creates a new MTRAttributePath instance.
func NewMTRAttributePath() MTRAttributePath {
	return getMTRAttributePathClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAttributePath */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributePath
type MTRAttributePath struct {
	MTRClusterPath
}

// MTRAttributePathFrom constructs a [MTRAttributePath] from an unsafe.Pointer.
func MTRAttributePathFrom(ptr unsafe.Pointer) MTRAttributePath {
	return MTRAttributePath{
		MTRClusterPath: MTRClusterPathFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAttributePath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributePath/init(endpointID:clusterID:attributeID:)-5tp3a
func NewMTRAttributePathWithEndpointIDClusterIDAttributeID(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */) MTRAttributePath {
	rv := objc.Send[MTRAttributePath](objc.ID(getMTRAttributePathClass().class), objc.Sel("attributePathWithEndpointID:clusterID:attributeID:"), endpointID, clusterID, attributeID)
	return rv
}/* debug [class_init_methods/constructor]: NewMTRAttributePathWithEndpointIDClusterIDAttributeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributePath/init(endpointId:clusterId:attributeId:)-49h2k
func NewMTRAttributePathWithEndpointIdClusterIdAttributeId(endpointId objc.IObject /* cross-framework: NSNumber */, clusterId objc.IObject /* cross-framework: NSNumber */, attributeId objc.IObject /* cross-framework: NSNumber */) MTRAttributePath {
	rv := objc.Send[MTRAttributePath](objc.ID(getMTRAttributePathClass().class), objc.Sel("attributePathWithEndpointId:clusterId:attributeId:"), endpointId, clusterId, attributeId)
	return rv
}/* debug [class_init_methods/constructor]: NewMTRAttributePathWithEndpointIdClusterIdAttributeId */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAttributePath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributePath/init(endpointId:clusterId:attributeId:)-49h2k
func (mc _MTRAttributePathClass) AttributePathWithEndpointIdClusterIdAttributeId(endpointId objc.IObject /* cross-framework: NSNumber */, clusterId objc.IObject /* cross-framework: NSNumber */, attributeId objc.IObject /* cross-framework: NSNumber */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("attributePathWithEndpointId:clusterId:attributeId:"), endpointId, clusterId, attributeId)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AttributePathWithEndpointIdClusterIdAttributeId) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributePath/init(endpointID:clusterID:attributeID:)-5tp3a
func (mc _MTRAttributePathClass) AttributePathWithEndpointIDClusterIDAttributeID(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */) MTRAttributePath {
	rv := objc.Send[MTRAttributePath](objc.ID(mc.class), objc.Sel("attributePathWithEndpointID:clusterID:attributeID:"), endpointID, clusterID, attributeID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AttributePathWithEndpointIDClusterIDAttributeID) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAttributePath */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAttributePath */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAttributePath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributePath/attribute
func (m_ MTRAttributePath) Attribute() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("attribute"))
	return rv
}/* debug [instance_properties/getter]: attribute */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAttributePath */


