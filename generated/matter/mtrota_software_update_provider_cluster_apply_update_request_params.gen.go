// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams */


/* debug [class_header]: Header for MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams */
// The class instance for the [MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams] class.
var (
	MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass     _MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass
	MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClassOnce sync.Once
)

func getMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass() _MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass {
	MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClassOnce.Do(func() {
		MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass = _MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass{objc.GetClass("MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams")}
	})
	return MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass
}

type _MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams */
// An interface definition for the [MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams] class.
type IMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams */
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

	
/* debug [class_interface_methods]: Methods for MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass) New() MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams) Init() MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams) Autorelease() MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams creates a new MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams instance.
func NewMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams() MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams {
	return getMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams-1mlcr
type MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams {
	return MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams-1mlcr/newVersion
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams) NewVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newVersion"))
	return rv
}/* debug [instance_properties/getter]: newVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams-1mlcr/newVersion
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams) SetNewVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewVersion:"), value)
}/* debug [instance_properties/setter]: newVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams-1mlcr/serverSideProcessingTimeout
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams-1mlcr/serverSideProcessingTimeout
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams-1mlcr/timedInvokeTimeoutMs
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams-1mlcr/timedInvokeTimeoutMs
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams-1mlcr/updateToken
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams) UpdateToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("updateToken"))
	return rv
}/* debug [instance_properties/getter]: updateToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams-1mlcr/updateToken
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams) SetUpdateToken(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}/* debug [instance_properties/setter]: updateToken */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams */



