// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRApplicationLauncherClusterLaunchAppParams */


/* debug [class_header]: Header for MTRApplicationLauncherClusterLaunchAppParams */
// The class instance for the [MTRApplicationLauncherClusterLaunchAppParams] class.
var (
	MTRApplicationLauncherClusterLaunchAppParamsClass     _MTRApplicationLauncherClusterLaunchAppParamsClass
	MTRApplicationLauncherClusterLaunchAppParamsClassOnce sync.Once
)

func getMTRApplicationLauncherClusterLaunchAppParamsClass() _MTRApplicationLauncherClusterLaunchAppParamsClass {
	MTRApplicationLauncherClusterLaunchAppParamsClassOnce.Do(func() {
		MTRApplicationLauncherClusterLaunchAppParamsClass = _MTRApplicationLauncherClusterLaunchAppParamsClass{objc.GetClass("MTRApplicationLauncherClusterLaunchAppParams")}
	})
	return MTRApplicationLauncherClusterLaunchAppParamsClass
}

type _MTRApplicationLauncherClusterLaunchAppParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRApplicationLauncherClusterLaunchAppParams */
// An interface definition for the [MTRApplicationLauncherClusterLaunchAppParams] class.
type IMTRApplicationLauncherClusterLaunchAppParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRApplicationLauncherClusterLaunchAppParams */
	// properties:
	Application() IMTRApplicationLauncherClusterApplicationStruct
	SetApplication(value IMTRApplicationLauncherClusterApplicationStruct)
	Data() objc.IObject /* cross-framework: NSData */
	SetData(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRApplicationLauncherClusterLaunchAppParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRApplicationLauncherClusterLaunchAppParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterLaunchAppParamsClass) Alloc() MTRApplicationLauncherClusterLaunchAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterLaunchAppParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRApplicationLauncherClusterLaunchAppParamsClass) New() MTRApplicationLauncherClusterLaunchAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterLaunchAppParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterLaunchAppParams) Init() MTRApplicationLauncherClusterLaunchAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterLaunchAppParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterLaunchAppParams) Autorelease() MTRApplicationLauncherClusterLaunchAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterLaunchAppParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterLaunchAppParams creates a new MTRApplicationLauncherClusterLaunchAppParams instance.
func NewMTRApplicationLauncherClusterLaunchAppParams() MTRApplicationLauncherClusterLaunchAppParams {
	return getMTRApplicationLauncherClusterLaunchAppParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRApplicationLauncherClusterLaunchAppParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLaunchAppParams
type MTRApplicationLauncherClusterLaunchAppParams struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterLaunchAppParamsFrom constructs a [MTRApplicationLauncherClusterLaunchAppParams] from an unsafe.Pointer.
func MTRApplicationLauncherClusterLaunchAppParamsFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterLaunchAppParams {
	return MTRApplicationLauncherClusterLaunchAppParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRApplicationLauncherClusterLaunchAppParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRApplicationLauncherClusterLaunchAppParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRApplicationLauncherClusterLaunchAppParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRApplicationLauncherClusterLaunchAppParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRApplicationLauncherClusterLaunchAppParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLaunchAppParams/application
func (m_ MTRApplicationLauncherClusterLaunchAppParams) Application() IMTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](m_.ID, objc.Sel("application"))
	return rv
}/* debug [instance_properties/getter]: application */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLaunchAppParams/application
func (m_ MTRApplicationLauncherClusterLaunchAppParams) SetApplication(value IMTRApplicationLauncherClusterApplicationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplication:"), value)
}/* debug [instance_properties/setter]: application */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLaunchAppParams/data
func (m_ MTRApplicationLauncherClusterLaunchAppParams) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLaunchAppParams/data
func (m_ MTRApplicationLauncherClusterLaunchAppParams) SetData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLaunchAppParams/serverSideProcessingTimeout
func (m_ MTRApplicationLauncherClusterLaunchAppParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLaunchAppParams/serverSideProcessingTimeout
func (m_ MTRApplicationLauncherClusterLaunchAppParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLaunchAppParams/timedInvokeTimeoutMs
func (m_ MTRApplicationLauncherClusterLaunchAppParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLaunchAppParams/timedInvokeTimeoutMs
func (m_ MTRApplicationLauncherClusterLaunchAppParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRApplicationLauncherClusterLaunchAppParams */



