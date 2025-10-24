// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTROtaSoftwareUpdateRequestorClusterProviderLocation */


/* debug [class_header]: Header for MTROtaSoftwareUpdateRequestorClusterProviderLocation */
// The class instance for the [MTROtaSoftwareUpdateRequestorClusterProviderLocation] class.
var (
	MTROtaSoftwareUpdateRequestorClusterProviderLocationClass     _MTROtaSoftwareUpdateRequestorClusterProviderLocationClass
	MTROtaSoftwareUpdateRequestorClusterProviderLocationClassOnce sync.Once
)

func getMTROtaSoftwareUpdateRequestorClusterProviderLocationClass() _MTROtaSoftwareUpdateRequestorClusterProviderLocationClass {
	MTROtaSoftwareUpdateRequestorClusterProviderLocationClassOnce.Do(func() {
		MTROtaSoftwareUpdateRequestorClusterProviderLocationClass = _MTROtaSoftwareUpdateRequestorClusterProviderLocationClass{objc.GetClass("MTROtaSoftwareUpdateRequestorClusterProviderLocation")}
	})
	return MTROtaSoftwareUpdateRequestorClusterProviderLocationClass
}

type _MTROtaSoftwareUpdateRequestorClusterProviderLocationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROtaSoftwareUpdateRequestorClusterProviderLocation */
// An interface definition for the [MTROtaSoftwareUpdateRequestorClusterProviderLocation] class.
type IMTROtaSoftwareUpdateRequestorClusterProviderLocation interface {
	IMTROTASoftwareUpdateRequestorClusterProviderLocation
	
/* debug [class_interface_properties]: Properties for MTROtaSoftwareUpdateRequestorClusterProviderLocation */
	// properties:
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	ProviderNodeID() objc.IObject /* cross-framework: NSNumber */
	SetProviderNodeID(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROtaSoftwareUpdateRequestorClusterProviderLocation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROtaSoftwareUpdateRequestorClusterProviderLocation */
// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateRequestorClusterProviderLocationClass) Alloc() MTROtaSoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterProviderLocation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROtaSoftwareUpdateRequestorClusterProviderLocationClass) New() MTROtaSoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterProviderLocation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) Init() MTROtaSoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterProviderLocation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) Autorelease() MTROtaSoftwareUpdateRequestorClusterProviderLocation {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterProviderLocation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateRequestorClusterProviderLocation creates a new MTROtaSoftwareUpdateRequestorClusterProviderLocation instance.
func NewMTROtaSoftwareUpdateRequestorClusterProviderLocation() MTROtaSoftwareUpdateRequestorClusterProviderLocation {
	return getMTROtaSoftwareUpdateRequestorClusterProviderLocationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROtaSoftwareUpdateRequestorClusterProviderLocation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterProviderLocation-yhnm
type MTROtaSoftwareUpdateRequestorClusterProviderLocation struct {
	MTROTASoftwareUpdateRequestorClusterProviderLocation
}

// MTROtaSoftwareUpdateRequestorClusterProviderLocationFrom constructs a [MTROtaSoftwareUpdateRequestorClusterProviderLocation] from an unsafe.Pointer.
func MTROtaSoftwareUpdateRequestorClusterProviderLocationFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateRequestorClusterProviderLocation {
	return MTROtaSoftwareUpdateRequestorClusterProviderLocation{
		MTROTASoftwareUpdateRequestorClusterProviderLocation: MTROTASoftwareUpdateRequestorClusterProviderLocationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROtaSoftwareUpdateRequestorClusterProviderLocation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROtaSoftwareUpdateRequestorClusterProviderLocation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROtaSoftwareUpdateRequestorClusterProviderLocation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROtaSoftwareUpdateRequestorClusterProviderLocation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROtaSoftwareUpdateRequestorClusterProviderLocation */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterProviderLocation-yhnm/endpoint
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterProviderLocation-yhnm/endpoint
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}/* debug [instance_properties/setter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterProviderLocation-yhnm/fabricIndex
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterProviderLocation-yhnm/fabricIndex
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterProviderLocation-yhnm/providerNodeID
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) ProviderNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("providerNodeID"))
	return rv
}/* debug [instance_properties/getter]: providerNodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterProviderLocation-yhnm/providerNodeID
func (m_ MTROtaSoftwareUpdateRequestorClusterProviderLocation) SetProviderNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderNodeID:"), value)
}/* debug [instance_properties/setter]: providerNodeID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROtaSoftwareUpdateRequestorClusterProviderLocation */



