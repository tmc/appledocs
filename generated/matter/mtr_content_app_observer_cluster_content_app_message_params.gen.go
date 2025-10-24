// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRContentAppObserverClusterContentAppMessageParams */


/* debug [class_header]: Header for MTRContentAppObserverClusterContentAppMessageParams */
// The class instance for the [MTRContentAppObserverClusterContentAppMessageParams] class.
var (
	MTRContentAppObserverClusterContentAppMessageParamsClass     _MTRContentAppObserverClusterContentAppMessageParamsClass
	MTRContentAppObserverClusterContentAppMessageParamsClassOnce sync.Once
)

func getMTRContentAppObserverClusterContentAppMessageParamsClass() _MTRContentAppObserverClusterContentAppMessageParamsClass {
	MTRContentAppObserverClusterContentAppMessageParamsClassOnce.Do(func() {
		MTRContentAppObserverClusterContentAppMessageParamsClass = _MTRContentAppObserverClusterContentAppMessageParamsClass{objc.GetClass("MTRContentAppObserverClusterContentAppMessageParams")}
	})
	return MTRContentAppObserverClusterContentAppMessageParamsClass
}

type _MTRContentAppObserverClusterContentAppMessageParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRContentAppObserverClusterContentAppMessageParams */
// An interface definition for the [MTRContentAppObserverClusterContentAppMessageParams] class.
type IMTRContentAppObserverClusterContentAppMessageParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRContentAppObserverClusterContentAppMessageParams */
	// properties:
	Data() objc.IObject /* cross-framework: NSString */
	SetData(value objc.IObject /* cross-framework: NSString */)
	EncodingHint() objc.IObject /* cross-framework: NSString */
	SetEncodingHint(value objc.IObject /* cross-framework: NSString */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRContentAppObserverClusterContentAppMessageParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRContentAppObserverClusterContentAppMessageParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRContentAppObserverClusterContentAppMessageParamsClass) Alloc() MTRContentAppObserverClusterContentAppMessageParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRContentAppObserverClusterContentAppMessageParamsClass) New() MTRContentAppObserverClusterContentAppMessageParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentAppObserverClusterContentAppMessageParams) Init() MTRContentAppObserverClusterContentAppMessageParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentAppObserverClusterContentAppMessageParams) Autorelease() MTRContentAppObserverClusterContentAppMessageParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentAppObserverClusterContentAppMessageParams creates a new MTRContentAppObserverClusterContentAppMessageParams instance.
func NewMTRContentAppObserverClusterContentAppMessageParams() MTRContentAppObserverClusterContentAppMessageParams {
	return getMTRContentAppObserverClusterContentAppMessageParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRContentAppObserverClusterContentAppMessageParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams
type MTRContentAppObserverClusterContentAppMessageParams struct {
	objectivec.Object
}

// MTRContentAppObserverClusterContentAppMessageParamsFrom constructs a [MTRContentAppObserverClusterContentAppMessageParams] from an unsafe.Pointer.
func MTRContentAppObserverClusterContentAppMessageParamsFrom(ptr unsafe.Pointer) MTRContentAppObserverClusterContentAppMessageParams {
	return MTRContentAppObserverClusterContentAppMessageParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRContentAppObserverClusterContentAppMessageParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRContentAppObserverClusterContentAppMessageParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRContentAppObserverClusterContentAppMessageParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRContentAppObserverClusterContentAppMessageParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRContentAppObserverClusterContentAppMessageParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/data
func (m_ MTRContentAppObserverClusterContentAppMessageParams) Data() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/data
func (m_ MTRContentAppObserverClusterContentAppMessageParams) SetData(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentappobserverclustercontentappmessageparams/encodinghint
func (m_ MTRContentAppObserverClusterContentAppMessageParams) EncodingHint() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("encodingHint"))
	return rv
}/* debug [instance_properties/getter]: encodingHint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentappobserverclustercontentappmessageparams/encodinghint
func (m_ MTRContentAppObserverClusterContentAppMessageParams) SetEncodingHint(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEncodingHint:"), value)
}/* debug [instance_properties/setter]: encodingHint */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentappobserverclustercontentappmessageparams/serversideprocessingtimeout
func (m_ MTRContentAppObserverClusterContentAppMessageParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentappobserverclustercontentappmessageparams/serversideprocessingtimeout
func (m_ MTRContentAppObserverClusterContentAppMessageParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentappobserverclustercontentappmessageparams/timedinvoketimeoutms
func (m_ MTRContentAppObserverClusterContentAppMessageParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentappobserverclustercontentappmessageparams/timedinvoketimeoutms
func (m_ MTRContentAppObserverClusterContentAppMessageParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRContentAppObserverClusterContentAppMessageParams */



