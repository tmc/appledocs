// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams */


/* debug [class_header]: Header for MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams */
// The class instance for the [MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams] class.
var (
	MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass     _MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass
	MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass() _MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass {
	MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass = _MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass{objc.GetClass("MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams")}
	})
	return MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass
}

type _MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams */
// An interface definition for the [MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams] class.
type IMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams interface {
	IMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams
	
/* debug [class_interface_properties]: Properties for MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams */
	// properties:
	NewVersion() objc.IObject /* cross-framework: NSNumber */
	SetNewVersion(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UpdateToken() objc.IObject /* cross-framework: NSData */
	SetUpdateToken(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass) New() MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) Init() MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) Autorelease() MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams creates a new MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams instance.
func NewMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams() MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams {
	return getMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams-5k4nj
type MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams struct {
	MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams
}

// MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsFrom constructs a [MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams {
	return MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams{
		MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams: MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams-5k4nj/newVersion
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) NewVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newVersion"))
	return rv
}/* debug [instance_properties/getter]: newVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams-5k4nj/newVersion
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) SetNewVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewVersion:"), value)
}/* debug [instance_properties/setter]: newVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams-5k4nj/serverSideProcessingTimeout
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams-5k4nj/serverSideProcessingTimeout
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams-5k4nj/timedInvokeTimeoutMs
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams-5k4nj/timedInvokeTimeoutMs
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams-5k4nj/updateToken
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) UpdateToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("updateToken"))
	return rv
}/* debug [instance_properties/getter]: updateToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams-5k4nj/updateToken
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) SetUpdateToken(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}/* debug [instance_properties/setter]: updateToken */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams */



