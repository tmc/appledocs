// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams */


/* debug [class_header]: Header for MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
// The class instance for the [MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams] class.
var (
	MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass     _MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass
	MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClassOnce sync.Once
)

func getMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass() _MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass {
	MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClassOnce.Do(func() {
		MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass = _MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass{objc.GetClass("MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams")}
	})
	return MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass
}

type _MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
// An interface definition for the [MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams] class.
type IMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
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

	
/* debug [class_interface_methods]: Methods for MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass) New() MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) Init() MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) Autorelease() MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams creates a new MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams instance.
func NewMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams() MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	return getMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams-5eau8
type MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	return MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams-5eau8/serverSideProcessingTimeout
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams-5eau8/serverSideProcessingTimeout
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams-5eau8/softwareVersion
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}/* debug [instance_properties/getter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams-5eau8/softwareVersion
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams-5eau8/timedInvokeTimeoutMs
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams-5eau8/timedInvokeTimeoutMs
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams-5eau8/updateToken
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) UpdateToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("updateToken"))
	return rv
}/* debug [instance_properties/getter]: updateToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams-5eau8/updateToken
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetUpdateToken(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}/* debug [instance_properties/setter]: updateToken */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams */



