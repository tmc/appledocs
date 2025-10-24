// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterUpdateFabricLabelParams */


/* debug [class_header]: Header for MTROperationalCredentialsClusterUpdateFabricLabelParams */
// The class instance for the [MTROperationalCredentialsClusterUpdateFabricLabelParams] class.
var (
	MTROperationalCredentialsClusterUpdateFabricLabelParamsClass     _MTROperationalCredentialsClusterUpdateFabricLabelParamsClass
	MTROperationalCredentialsClusterUpdateFabricLabelParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterUpdateFabricLabelParamsClass() _MTROperationalCredentialsClusterUpdateFabricLabelParamsClass {
	MTROperationalCredentialsClusterUpdateFabricLabelParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterUpdateFabricLabelParamsClass = _MTROperationalCredentialsClusterUpdateFabricLabelParamsClass{objc.GetClass("MTROperationalCredentialsClusterUpdateFabricLabelParams")}
	})
	return MTROperationalCredentialsClusterUpdateFabricLabelParamsClass
}

type _MTROperationalCredentialsClusterUpdateFabricLabelParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterUpdateFabricLabelParams */
// An interface definition for the [MTROperationalCredentialsClusterUpdateFabricLabelParams] class.
type IMTROperationalCredentialsClusterUpdateFabricLabelParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterUpdateFabricLabelParams */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterUpdateFabricLabelParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterUpdateFabricLabelParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterUpdateFabricLabelParamsClass) Alloc() MTROperationalCredentialsClusterUpdateFabricLabelParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateFabricLabelParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterUpdateFabricLabelParamsClass) New() MTROperationalCredentialsClusterUpdateFabricLabelParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateFabricLabelParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) Init() MTROperationalCredentialsClusterUpdateFabricLabelParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateFabricLabelParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) Autorelease() MTROperationalCredentialsClusterUpdateFabricLabelParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateFabricLabelParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterUpdateFabricLabelParams creates a new MTROperationalCredentialsClusterUpdateFabricLabelParams instance.
func NewMTROperationalCredentialsClusterUpdateFabricLabelParams() MTROperationalCredentialsClusterUpdateFabricLabelParams {
	return getMTROperationalCredentialsClusterUpdateFabricLabelParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterUpdateFabricLabelParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateFabricLabelParams
type MTROperationalCredentialsClusterUpdateFabricLabelParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterUpdateFabricLabelParamsFrom constructs a [MTROperationalCredentialsClusterUpdateFabricLabelParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterUpdateFabricLabelParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterUpdateFabricLabelParams {
	return MTROperationalCredentialsClusterUpdateFabricLabelParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterUpdateFabricLabelParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterUpdateFabricLabelParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterUpdateFabricLabelParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterUpdateFabricLabelParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterUpdateFabricLabelParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateFabricLabelParams/label
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateFabricLabelParams/label
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateFabricLabelParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateFabricLabelParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateFabricLabelParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateFabricLabelParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterUpdateFabricLabelParams */



