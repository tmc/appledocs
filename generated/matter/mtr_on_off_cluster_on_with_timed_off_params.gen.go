// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROnOffClusterOnWithTimedOffParams */


/* debug [class_header]: Header for MTROnOffClusterOnWithTimedOffParams */
// The class instance for the [MTROnOffClusterOnWithTimedOffParams] class.
var (
	MTROnOffClusterOnWithTimedOffParamsClass     _MTROnOffClusterOnWithTimedOffParamsClass
	MTROnOffClusterOnWithTimedOffParamsClassOnce sync.Once
)

func getMTROnOffClusterOnWithTimedOffParamsClass() _MTROnOffClusterOnWithTimedOffParamsClass {
	MTROnOffClusterOnWithTimedOffParamsClassOnce.Do(func() {
		MTROnOffClusterOnWithTimedOffParamsClass = _MTROnOffClusterOnWithTimedOffParamsClass{objc.GetClass("MTROnOffClusterOnWithTimedOffParams")}
	})
	return MTROnOffClusterOnWithTimedOffParamsClass
}

type _MTROnOffClusterOnWithTimedOffParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROnOffClusterOnWithTimedOffParams */
// An interface definition for the [MTROnOffClusterOnWithTimedOffParams] class.
type IMTROnOffClusterOnWithTimedOffParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROnOffClusterOnWithTimedOffParams */
	// properties:
	OffWaitTime() objc.IObject /* cross-framework: NSNumber */
	SetOffWaitTime(value objc.IObject /* cross-framework: NSNumber */)
	OnOffControl() objc.IObject /* cross-framework: NSNumber */
	SetOnOffControl(value objc.IObject /* cross-framework: NSNumber */)
	OnTime() objc.IObject /* cross-framework: NSNumber */
	SetOnTime(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROnOffClusterOnWithTimedOffParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROnOffClusterOnWithTimedOffParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROnOffClusterOnWithTimedOffParamsClass) Alloc() MTROnOffClusterOnWithTimedOffParams {
	rv := objc.Send[MTROnOffClusterOnWithTimedOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROnOffClusterOnWithTimedOffParamsClass) New() MTROnOffClusterOnWithTimedOffParams {
	rv := objc.Send[MTROnOffClusterOnWithTimedOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnOffClusterOnWithTimedOffParams) Init() MTROnOffClusterOnWithTimedOffParams {
	rv := objc.Send[MTROnOffClusterOnWithTimedOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnOffClusterOnWithTimedOffParams) Autorelease() MTROnOffClusterOnWithTimedOffParams {
	rv := objc.Send[MTROnOffClusterOnWithTimedOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnOffClusterOnWithTimedOffParams creates a new MTROnOffClusterOnWithTimedOffParams instance.
func NewMTROnOffClusterOnWithTimedOffParams() MTROnOffClusterOnWithTimedOffParams {
	return getMTROnOffClusterOnWithTimedOffParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROnOffClusterOnWithTimedOffParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithTimedOffParams
type MTROnOffClusterOnWithTimedOffParams struct {
	objectivec.Object
}

// MTROnOffClusterOnWithTimedOffParamsFrom constructs a [MTROnOffClusterOnWithTimedOffParams] from an unsafe.Pointer.
func MTROnOffClusterOnWithTimedOffParamsFrom(ptr unsafe.Pointer) MTROnOffClusterOnWithTimedOffParams {
	return MTROnOffClusterOnWithTimedOffParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROnOffClusterOnWithTimedOffParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROnOffClusterOnWithTimedOffParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROnOffClusterOnWithTimedOffParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROnOffClusterOnWithTimedOffParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROnOffClusterOnWithTimedOffParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithTimedOffParams/offWaitTime
func (m_ MTROnOffClusterOnWithTimedOffParams) OffWaitTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("offWaitTime"))
	return rv
}/* debug [instance_properties/getter]: offWaitTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithTimedOffParams/offWaitTime
func (m_ MTROnOffClusterOnWithTimedOffParams) SetOffWaitTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffWaitTime:"), value)
}/* debug [instance_properties/setter]: offWaitTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithTimedOffParams/onOffControl
func (m_ MTROnOffClusterOnWithTimedOffParams) OnOffControl() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("onOffControl"))
	return rv
}/* debug [instance_properties/getter]: onOffControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithTimedOffParams/onOffControl
func (m_ MTROnOffClusterOnWithTimedOffParams) SetOnOffControl(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOnOffControl:"), value)
}/* debug [instance_properties/setter]: onOffControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithTimedOffParams/onTime
func (m_ MTROnOffClusterOnWithTimedOffParams) OnTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("onTime"))
	return rv
}/* debug [instance_properties/getter]: onTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithTimedOffParams/onTime
func (m_ MTROnOffClusterOnWithTimedOffParams) SetOnTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOnTime:"), value)
}/* debug [instance_properties/setter]: onTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithTimedOffParams/serverSideProcessingTimeout
func (m_ MTROnOffClusterOnWithTimedOffParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithTimedOffParams/serverSideProcessingTimeout
func (m_ MTROnOffClusterOnWithTimedOffParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithTimedOffParams/timedInvokeTimeoutMs
func (m_ MTROnOffClusterOnWithTimedOffParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithTimedOffParams/timedInvokeTimeoutMs
func (m_ MTROnOffClusterOnWithTimedOffParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROnOffClusterOnWithTimedOffParams */



