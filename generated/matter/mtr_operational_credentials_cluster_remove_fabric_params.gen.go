// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterRemoveFabricParams */


/* debug [class_header]: Header for MTROperationalCredentialsClusterRemoveFabricParams */
// The class instance for the [MTROperationalCredentialsClusterRemoveFabricParams] class.
var (
	MTROperationalCredentialsClusterRemoveFabricParamsClass     _MTROperationalCredentialsClusterRemoveFabricParamsClass
	MTROperationalCredentialsClusterRemoveFabricParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterRemoveFabricParamsClass() _MTROperationalCredentialsClusterRemoveFabricParamsClass {
	MTROperationalCredentialsClusterRemoveFabricParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterRemoveFabricParamsClass = _MTROperationalCredentialsClusterRemoveFabricParamsClass{objc.GetClass("MTROperationalCredentialsClusterRemoveFabricParams")}
	})
	return MTROperationalCredentialsClusterRemoveFabricParamsClass
}

type _MTROperationalCredentialsClusterRemoveFabricParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterRemoveFabricParams */
// An interface definition for the [MTROperationalCredentialsClusterRemoveFabricParams] class.
type IMTROperationalCredentialsClusterRemoveFabricParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterRemoveFabricParams */
	// properties:
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterRemoveFabricParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterRemoveFabricParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterRemoveFabricParamsClass) Alloc() MTROperationalCredentialsClusterRemoveFabricParams {
	rv := objc.Send[MTROperationalCredentialsClusterRemoveFabricParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterRemoveFabricParamsClass) New() MTROperationalCredentialsClusterRemoveFabricParams {
	rv := objc.Send[MTROperationalCredentialsClusterRemoveFabricParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) Init() MTROperationalCredentialsClusterRemoveFabricParams {
	rv := objc.Send[MTROperationalCredentialsClusterRemoveFabricParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) Autorelease() MTROperationalCredentialsClusterRemoveFabricParams {
	rv := objc.Send[MTROperationalCredentialsClusterRemoveFabricParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterRemoveFabricParams creates a new MTROperationalCredentialsClusterRemoveFabricParams instance.
func NewMTROperationalCredentialsClusterRemoveFabricParams() MTROperationalCredentialsClusterRemoveFabricParams {
	return getMTROperationalCredentialsClusterRemoveFabricParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterRemoveFabricParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterRemoveFabricParams
type MTROperationalCredentialsClusterRemoveFabricParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterRemoveFabricParamsFrom constructs a [MTROperationalCredentialsClusterRemoveFabricParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterRemoveFabricParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterRemoveFabricParams {
	return MTROperationalCredentialsClusterRemoveFabricParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterRemoveFabricParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterRemoveFabricParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterRemoveFabricParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterRemoveFabricParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterRemoveFabricParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterRemoveFabricParams/fabricIndex
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterRemoveFabricParams/fabricIndex
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterRemoveFabricParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterRemoveFabricParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterRemoveFabricParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterRemoveFabricParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterRemoveFabricParams */



