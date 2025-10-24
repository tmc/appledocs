// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROTASoftwareUpdateRequestorClusterProviderLocation */


/* debug [class_header]: Header for MTROTASoftwareUpdateRequestorClusterProviderLocation */
// The class instance for the [MTROTASoftwareUpdateRequestorClusterProviderLocation] class.
var (
	MTROTASoftwareUpdateRequestorClusterProviderLocationClass     _MTROTASoftwareUpdateRequestorClusterProviderLocationClass
	MTROTASoftwareUpdateRequestorClusterProviderLocationClassOnce sync.Once
)

func getMTROTASoftwareUpdateRequestorClusterProviderLocationClass() _MTROTASoftwareUpdateRequestorClusterProviderLocationClass {
	MTROTASoftwareUpdateRequestorClusterProviderLocationClassOnce.Do(func() {
		MTROTASoftwareUpdateRequestorClusterProviderLocationClass = _MTROTASoftwareUpdateRequestorClusterProviderLocationClass{objc.GetClass("MTROTASoftwareUpdateRequestorClusterProviderLocation")}
	})
	return MTROTASoftwareUpdateRequestorClusterProviderLocationClass
}

type _MTROTASoftwareUpdateRequestorClusterProviderLocationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROTASoftwareUpdateRequestorClusterProviderLocation */
// An interface definition for the [MTROTASoftwareUpdateRequestorClusterProviderLocation] class.
type IMTROTASoftwareUpdateRequestorClusterProviderLocation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROTASoftwareUpdateRequestorClusterProviderLocation */
	// properties:
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	ProviderNodeID() objc.IObject /* cross-framework: NSNumber */
	SetProviderNodeID(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROTASoftwareUpdateRequestorClusterProviderLocation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROTASoftwareUpdateRequestorClusterProviderLocation */
// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateRequestorClusterProviderLocationClass) Alloc() MTROTASoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterProviderLocation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROTASoftwareUpdateRequestorClusterProviderLocationClass) New() MTROTASoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterProviderLocation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) Init() MTROTASoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterProviderLocation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) Autorelease() MTROTASoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterProviderLocation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateRequestorClusterProviderLocation creates a new MTROTASoftwareUpdateRequestorClusterProviderLocation instance.
func NewMTROTASoftwareUpdateRequestorClusterProviderLocation() MTROTASoftwareUpdateRequestorClusterProviderLocation {
	return getMTROTASoftwareUpdateRequestorClusterProviderLocationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROTASoftwareUpdateRequestorClusterProviderLocation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterProviderLocation-76vsq
type MTROTASoftwareUpdateRequestorClusterProviderLocation struct {
	objectivec.Object
}

// MTROTASoftwareUpdateRequestorClusterProviderLocationFrom constructs a [MTROTASoftwareUpdateRequestorClusterProviderLocation] from an unsafe.Pointer.
func MTROTASoftwareUpdateRequestorClusterProviderLocationFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateRequestorClusterProviderLocation {
	return MTROTASoftwareUpdateRequestorClusterProviderLocation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROTASoftwareUpdateRequestorClusterProviderLocation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROTASoftwareUpdateRequestorClusterProviderLocation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROTASoftwareUpdateRequestorClusterProviderLocation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROTASoftwareUpdateRequestorClusterProviderLocation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROTASoftwareUpdateRequestorClusterProviderLocation */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterProviderLocation-76vsq/endpoint
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterProviderLocation-76vsq/endpoint
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}/* debug [instance_properties/setter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterProviderLocation-76vsq/fabricIndex
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterProviderLocation-76vsq/fabricIndex
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterProviderLocation-76vsq/providerNodeID
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) ProviderNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("providerNodeID"))
	return rv
}/* debug [instance_properties/getter]: providerNodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterProviderLocation-76vsq/providerNodeID
func (m_ MTROTASoftwareUpdateRequestorClusterProviderLocation) SetProviderNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderNodeID:"), value)
}/* debug [instance_properties/setter]: providerNodeID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROTASoftwareUpdateRequestorClusterProviderLocation */



