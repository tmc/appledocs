// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRServiceAreaClusterSkipAreaParams */


/* debug [class_header]: Header for MTRServiceAreaClusterSkipAreaParams */
// The class instance for the [MTRServiceAreaClusterSkipAreaParams] class.
var (
	MTRServiceAreaClusterSkipAreaParamsClass     _MTRServiceAreaClusterSkipAreaParamsClass
	MTRServiceAreaClusterSkipAreaParamsClassOnce sync.Once
)

func getMTRServiceAreaClusterSkipAreaParamsClass() _MTRServiceAreaClusterSkipAreaParamsClass {
	MTRServiceAreaClusterSkipAreaParamsClassOnce.Do(func() {
		MTRServiceAreaClusterSkipAreaParamsClass = _MTRServiceAreaClusterSkipAreaParamsClass{objc.GetClass("MTRServiceAreaClusterSkipAreaParams")}
	})
	return MTRServiceAreaClusterSkipAreaParamsClass
}

type _MTRServiceAreaClusterSkipAreaParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRServiceAreaClusterSkipAreaParams */
// An interface definition for the [MTRServiceAreaClusterSkipAreaParams] class.
type IMTRServiceAreaClusterSkipAreaParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRServiceAreaClusterSkipAreaParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	SkippedArea() objc.IObject /* cross-framework: NSNumber */
	SetSkippedArea(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRServiceAreaClusterSkipAreaParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRServiceAreaClusterSkipAreaParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterSkipAreaParamsClass) Alloc() MTRServiceAreaClusterSkipAreaParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRServiceAreaClusterSkipAreaParamsClass) New() MTRServiceAreaClusterSkipAreaParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterSkipAreaParams) Init() MTRServiceAreaClusterSkipAreaParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterSkipAreaParams) Autorelease() MTRServiceAreaClusterSkipAreaParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterSkipAreaParams creates a new MTRServiceAreaClusterSkipAreaParams instance.
func NewMTRServiceAreaClusterSkipAreaParams() MTRServiceAreaClusterSkipAreaParams {
	return getMTRServiceAreaClusterSkipAreaParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRServiceAreaClusterSkipAreaParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams
type MTRServiceAreaClusterSkipAreaParams struct {
	objectivec.Object
}

// MTRServiceAreaClusterSkipAreaParamsFrom constructs a [MTRServiceAreaClusterSkipAreaParams] from an unsafe.Pointer.
func MTRServiceAreaClusterSkipAreaParamsFrom(ptr unsafe.Pointer) MTRServiceAreaClusterSkipAreaParams {
	return MTRServiceAreaClusterSkipAreaParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRServiceAreaClusterSkipAreaParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRServiceAreaClusterSkipAreaParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRServiceAreaClusterSkipAreaParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRServiceAreaClusterSkipAreaParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRServiceAreaClusterSkipAreaParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/serverSideProcessingTimeout
func (m_ MTRServiceAreaClusterSkipAreaParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/serverSideProcessingTimeout
func (m_ MTRServiceAreaClusterSkipAreaParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/skippedArea
func (m_ MTRServiceAreaClusterSkipAreaParams) SkippedArea() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("skippedArea"))
	return rv
}/* debug [instance_properties/getter]: skippedArea */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/skippedArea
func (m_ MTRServiceAreaClusterSkipAreaParams) SetSkippedArea(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSkippedArea:"), value)
}/* debug [instance_properties/setter]: skippedArea */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterskipareaparams/timedinvoketimeoutms
func (m_ MTRServiceAreaClusterSkipAreaParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterskipareaparams/timedinvoketimeoutms
func (m_ MTRServiceAreaClusterSkipAreaParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRServiceAreaClusterSkipAreaParams */



