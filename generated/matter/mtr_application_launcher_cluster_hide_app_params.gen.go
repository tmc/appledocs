// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRApplicationLauncherClusterHideAppParams */


/* debug [class_header]: Header for MTRApplicationLauncherClusterHideAppParams */
// The class instance for the [MTRApplicationLauncherClusterHideAppParams] class.
var (
	MTRApplicationLauncherClusterHideAppParamsClass     _MTRApplicationLauncherClusterHideAppParamsClass
	MTRApplicationLauncherClusterHideAppParamsClassOnce sync.Once
)

func getMTRApplicationLauncherClusterHideAppParamsClass() _MTRApplicationLauncherClusterHideAppParamsClass {
	MTRApplicationLauncherClusterHideAppParamsClassOnce.Do(func() {
		MTRApplicationLauncherClusterHideAppParamsClass = _MTRApplicationLauncherClusterHideAppParamsClass{objc.GetClass("MTRApplicationLauncherClusterHideAppParams")}
	})
	return MTRApplicationLauncherClusterHideAppParamsClass
}

type _MTRApplicationLauncherClusterHideAppParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRApplicationLauncherClusterHideAppParams */
// An interface definition for the [MTRApplicationLauncherClusterHideAppParams] class.
type IMTRApplicationLauncherClusterHideAppParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRApplicationLauncherClusterHideAppParams */
	// properties:
	Application() IMTRApplicationLauncherClusterApplicationStruct
	SetApplication(value IMTRApplicationLauncherClusterApplicationStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRApplicationLauncherClusterHideAppParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRApplicationLauncherClusterHideAppParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterHideAppParamsClass) Alloc() MTRApplicationLauncherClusterHideAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterHideAppParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRApplicationLauncherClusterHideAppParamsClass) New() MTRApplicationLauncherClusterHideAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterHideAppParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterHideAppParams) Init() MTRApplicationLauncherClusterHideAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterHideAppParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterHideAppParams) Autorelease() MTRApplicationLauncherClusterHideAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterHideAppParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterHideAppParams creates a new MTRApplicationLauncherClusterHideAppParams instance.
func NewMTRApplicationLauncherClusterHideAppParams() MTRApplicationLauncherClusterHideAppParams {
	return getMTRApplicationLauncherClusterHideAppParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRApplicationLauncherClusterHideAppParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterHideAppParams
type MTRApplicationLauncherClusterHideAppParams struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterHideAppParamsFrom constructs a [MTRApplicationLauncherClusterHideAppParams] from an unsafe.Pointer.
func MTRApplicationLauncherClusterHideAppParamsFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterHideAppParams {
	return MTRApplicationLauncherClusterHideAppParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRApplicationLauncherClusterHideAppParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRApplicationLauncherClusterHideAppParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRApplicationLauncherClusterHideAppParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRApplicationLauncherClusterHideAppParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRApplicationLauncherClusterHideAppParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterHideAppParams/application
func (m_ MTRApplicationLauncherClusterHideAppParams) Application() IMTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](m_.ID, objc.Sel("application"))
	return rv
}/* debug [instance_properties/getter]: application */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterHideAppParams/application
func (m_ MTRApplicationLauncherClusterHideAppParams) SetApplication(value IMTRApplicationLauncherClusterApplicationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplication:"), value)
}/* debug [instance_properties/setter]: application */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterHideAppParams/serverSideProcessingTimeout
func (m_ MTRApplicationLauncherClusterHideAppParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterHideAppParams/serverSideProcessingTimeout
func (m_ MTRApplicationLauncherClusterHideAppParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterHideAppParams/timedInvokeTimeoutMs
func (m_ MTRApplicationLauncherClusterHideAppParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterHideAppParams/timedInvokeTimeoutMs
func (m_ MTRApplicationLauncherClusterHideAppParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRApplicationLauncherClusterHideAppParams */



