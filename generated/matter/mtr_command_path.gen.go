// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRCommandPath */


/* debug [class_header]: Header for MTRCommandPath */
// The class instance for the [MTRCommandPath] class.
var (
	MTRCommandPathClass     _MTRCommandPathClass
	MTRCommandPathClassOnce sync.Once
)

func getMTRCommandPathClass() _MTRCommandPathClass {
	MTRCommandPathClassOnce.Do(func() {
		MTRCommandPathClass = _MTRCommandPathClass{objc.GetClass("MTRCommandPath")}
	})
	return MTRCommandPathClass
}

type _MTRCommandPathClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRCommandPath */
// An interface definition for the [MTRCommandPath] class.
type IMTRCommandPath interface {
	IMTRClusterPath
	
/* debug [class_interface_properties]: Properties for MTRCommandPath */
	// properties:
	Command() objc.IObject /* cross-framework: NSNumber */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRCommandPath */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRCommandPath */
// Alloc allocates a new instance without initialization.
func (mc _MTRCommandPathClass) Alloc() MTRCommandPath {
	rv := objc.Send[MTRCommandPath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRCommandPathClass) New() MTRCommandPath {
	rv := objc.Send[MTRCommandPath](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommandPath) Init() MTRCommandPath {
	rv := objc.Send[MTRCommandPath](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommandPath) Autorelease() MTRCommandPath {
	rv := objc.Send[MTRCommandPath](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommandPath creates a new MTRCommandPath instance.
func NewMTRCommandPath() MTRCommandPath {
	return getMTRCommandPathClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRCommandPath */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandPath
type MTRCommandPath struct {
	MTRClusterPath
}

// MTRCommandPathFrom constructs a [MTRCommandPath] from an unsafe.Pointer.
func MTRCommandPathFrom(ptr unsafe.Pointer) MTRCommandPath {
	return MTRCommandPath{
		MTRClusterPath: MTRClusterPathFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRCommandPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandPath/init(endpointID:clusterID:commandID:)-98znp
func NewMTRCommandPathWithEndpointIDClusterIDCommandID(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, commandID objc.IObject /* cross-framework: NSNumber */) MTRCommandPath {
	rv := objc.Send[MTRCommandPath](objc.ID(getMTRCommandPathClass().class), objc.Sel("commandPathWithEndpointID:clusterID:commandID:"), endpointID, clusterID, commandID)
	return rv
}/* debug [class_init_methods/constructor]: NewMTRCommandPathWithEndpointIDClusterIDCommandID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandPath/init(endpointId:clusterId:commandId:)-99j08
func NewMTRCommandPathWithEndpointIdClusterIdCommandId(endpointId objc.IObject /* cross-framework: NSNumber */, clusterId objc.IObject /* cross-framework: NSNumber */, commandId objc.IObject /* cross-framework: NSNumber */) MTRCommandPath {
	rv := objc.Send[MTRCommandPath](objc.ID(getMTRCommandPathClass().class), objc.Sel("commandPathWithEndpointId:clusterId:commandId:"), endpointId, clusterId, commandId)
	return rv
}/* debug [class_init_methods/constructor]: NewMTRCommandPathWithEndpointIdClusterIdCommandId */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRCommandPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandPath/init(endpointID:clusterID:commandID:)-98znp
func (mc _MTRCommandPathClass) CommandPathWithEndpointIDClusterIDCommandID(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, commandID objc.IObject /* cross-framework: NSNumber */) MTRCommandPath {
	rv := objc.Send[MTRCommandPath](objc.ID(mc.class), objc.Sel("commandPathWithEndpointID:clusterID:commandID:"), endpointID, clusterID, commandID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CommandPathWithEndpointIDClusterIDCommandID) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandPath/init(endpointId:clusterId:commandId:)-99j08
func (mc _MTRCommandPathClass) CommandPathWithEndpointIdClusterIdCommandId(endpointId objc.IObject /* cross-framework: NSNumber */, clusterId objc.IObject /* cross-framework: NSNumber */, commandId objc.IObject /* cross-framework: NSNumber */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("commandPathWithEndpointId:clusterId:commandId:"), endpointId, clusterId, commandId)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CommandPathWithEndpointIdClusterIdCommandId) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRCommandPath */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRCommandPath */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRCommandPath */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandPath/command
func (m_ MTRCommandPath) Command() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("command"))
	return rv
}/* debug [instance_properties/getter]: command */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRCommandPath */


