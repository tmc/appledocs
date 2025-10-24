// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams */


/* debug [class_header]: Header for MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
// The class instance for the [MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams] class.
var (
	MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass     _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass
	MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass() _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass {
	MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass = _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass{objc.GetClass("MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams")}
	})
	return MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass
}

type _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
// An interface definition for the [MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams] class.
type IMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams interface {
	IMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams
	
/* debug [class_interface_properties]: Properties for MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	SoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UpdateToken() objc.IObject /* cross-framework: NSData */
	SetUpdateToken(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass) New() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) Init() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) Autorelease() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams creates a new MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams instance.
func NewMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	return getMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams-6s1jv
type MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams struct {
	MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams
}

// MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom constructs a [MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	return MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams{
		MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams: MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams-6s1jv/serverSideProcessingTimeout
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams-6s1jv/serverSideProcessingTimeout
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams-6s1jv/softwareVersion
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}/* debug [instance_properties/getter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams-6s1jv/softwareVersion
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: softwareVersion */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams-6s1jv/timedInvokeTimeoutMs
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams-6s1jv/timedInvokeTimeoutMs
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams-6s1jv/updateToken
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) UpdateToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("updateToken"))
	return rv
}/* debug [instance_properties/getter]: updateToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams-6s1jv/updateToken
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetUpdateToken(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}/* debug [instance_properties/setter]: updateToken */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams */



