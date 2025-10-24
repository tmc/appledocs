// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRApplicationLauncherClusterLauncherResponseParams */


/* debug [class_header]: Header for MTRApplicationLauncherClusterLauncherResponseParams */
// The class instance for the [MTRApplicationLauncherClusterLauncherResponseParams] class.
var (
	MTRApplicationLauncherClusterLauncherResponseParamsClass     _MTRApplicationLauncherClusterLauncherResponseParamsClass
	MTRApplicationLauncherClusterLauncherResponseParamsClassOnce sync.Once
)

func getMTRApplicationLauncherClusterLauncherResponseParamsClass() _MTRApplicationLauncherClusterLauncherResponseParamsClass {
	MTRApplicationLauncherClusterLauncherResponseParamsClassOnce.Do(func() {
		MTRApplicationLauncherClusterLauncherResponseParamsClass = _MTRApplicationLauncherClusterLauncherResponseParamsClass{objc.GetClass("MTRApplicationLauncherClusterLauncherResponseParams")}
	})
	return MTRApplicationLauncherClusterLauncherResponseParamsClass
}

type _MTRApplicationLauncherClusterLauncherResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRApplicationLauncherClusterLauncherResponseParams */
// An interface definition for the [MTRApplicationLauncherClusterLauncherResponseParams] class.
type IMTRApplicationLauncherClusterLauncherResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRApplicationLauncherClusterLauncherResponseParams */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	SetData(value objc.IObject /* cross-framework: NSData */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRApplicationLauncherClusterLauncherResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRApplicationLauncherClusterLauncherResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterLauncherResponseParamsClass) Alloc() MTRApplicationLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRApplicationLauncherClusterLauncherResponseParamsClass) New() MTRApplicationLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) Init() MTRApplicationLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) Autorelease() MTRApplicationLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterLauncherResponseParams creates a new MTRApplicationLauncherClusterLauncherResponseParams instance.
func NewMTRApplicationLauncherClusterLauncherResponseParams() MTRApplicationLauncherClusterLauncherResponseParams {
	return getMTRApplicationLauncherClusterLauncherResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRApplicationLauncherClusterLauncherResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLauncherResponseParams
type MTRApplicationLauncherClusterLauncherResponseParams struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterLauncherResponseParamsFrom constructs a [MTRApplicationLauncherClusterLauncherResponseParams] from an unsafe.Pointer.
func MTRApplicationLauncherClusterLauncherResponseParamsFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterLauncherResponseParams {
	return MTRApplicationLauncherClusterLauncherResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRApplicationLauncherClusterLauncherResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLauncherResponseParams/init(responseValue:)
func NewMTRApplicationLauncherClusterLauncherResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRApplicationLauncherClusterLauncherResponseParams {
	instance := getMTRApplicationLauncherClusterLauncherResponseParamsClass().Alloc()
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRApplicationLauncherClusterLauncherResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRApplicationLauncherClusterLauncherResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRApplicationLauncherClusterLauncherResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRApplicationLauncherClusterLauncherResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRApplicationLauncherClusterLauncherResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLauncherResponseParams/data
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLauncherResponseParams/data
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) SetData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLauncherResponseParams/status
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLauncherResponseParams/status
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLauncherResponseParams/timedInvokeTimeoutMs
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLauncherResponseParams/timedInvokeTimeoutMs
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRApplicationLauncherClusterLauncherResponseParams */


