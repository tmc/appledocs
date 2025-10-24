// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRApplicationLauncherClusterStopAppParams */


/* debug [class_header]: Header for MTRApplicationLauncherClusterStopAppParams */
// The class instance for the [MTRApplicationLauncherClusterStopAppParams] class.
var (
	MTRApplicationLauncherClusterStopAppParamsClass     _MTRApplicationLauncherClusterStopAppParamsClass
	MTRApplicationLauncherClusterStopAppParamsClassOnce sync.Once
)

func getMTRApplicationLauncherClusterStopAppParamsClass() _MTRApplicationLauncherClusterStopAppParamsClass {
	MTRApplicationLauncherClusterStopAppParamsClassOnce.Do(func() {
		MTRApplicationLauncherClusterStopAppParamsClass = _MTRApplicationLauncherClusterStopAppParamsClass{objc.GetClass("MTRApplicationLauncherClusterStopAppParams")}
	})
	return MTRApplicationLauncherClusterStopAppParamsClass
}

type _MTRApplicationLauncherClusterStopAppParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRApplicationLauncherClusterStopAppParams */
// An interface definition for the [MTRApplicationLauncherClusterStopAppParams] class.
type IMTRApplicationLauncherClusterStopAppParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRApplicationLauncherClusterStopAppParams */
	// properties:
	Application() IMTRApplicationLauncherClusterApplicationStruct
	SetApplication(value IMTRApplicationLauncherClusterApplicationStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRApplicationLauncherClusterStopAppParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRApplicationLauncherClusterStopAppParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterStopAppParamsClass) Alloc() MTRApplicationLauncherClusterStopAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterStopAppParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRApplicationLauncherClusterStopAppParamsClass) New() MTRApplicationLauncherClusterStopAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterStopAppParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterStopAppParams) Init() MTRApplicationLauncherClusterStopAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterStopAppParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterStopAppParams) Autorelease() MTRApplicationLauncherClusterStopAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterStopAppParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterStopAppParams creates a new MTRApplicationLauncherClusterStopAppParams instance.
func NewMTRApplicationLauncherClusterStopAppParams() MTRApplicationLauncherClusterStopAppParams {
	return getMTRApplicationLauncherClusterStopAppParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRApplicationLauncherClusterStopAppParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterStopAppParams
type MTRApplicationLauncherClusterStopAppParams struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterStopAppParamsFrom constructs a [MTRApplicationLauncherClusterStopAppParams] from an unsafe.Pointer.
func MTRApplicationLauncherClusterStopAppParamsFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterStopAppParams {
	return MTRApplicationLauncherClusterStopAppParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRApplicationLauncherClusterStopAppParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRApplicationLauncherClusterStopAppParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRApplicationLauncherClusterStopAppParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRApplicationLauncherClusterStopAppParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRApplicationLauncherClusterStopAppParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterStopAppParams/application
func (m_ MTRApplicationLauncherClusterStopAppParams) Application() IMTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](m_.ID, objc.Sel("application"))
	return rv
}/* debug [instance_properties/getter]: application */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterStopAppParams/application
func (m_ MTRApplicationLauncherClusterStopAppParams) SetApplication(value IMTRApplicationLauncherClusterApplicationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplication:"), value)
}/* debug [instance_properties/setter]: application */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterStopAppParams/serverSideProcessingTimeout
func (m_ MTRApplicationLauncherClusterStopAppParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterStopAppParams/serverSideProcessingTimeout
func (m_ MTRApplicationLauncherClusterStopAppParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterStopAppParams/timedInvokeTimeoutMs
func (m_ MTRApplicationLauncherClusterStopAppParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterStopAppParams/timedInvokeTimeoutMs
func (m_ MTRApplicationLauncherClusterStopAppParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRApplicationLauncherClusterStopAppParams */



