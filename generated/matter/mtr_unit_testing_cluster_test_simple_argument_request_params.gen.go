// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestSimpleArgumentRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestSimpleArgumentRequestParams */
// The class instance for the [MTRUnitTestingClusterTestSimpleArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass     _MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass
	MTRUnitTestingClusterTestSimpleArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestSimpleArgumentRequestParamsClass() _MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass {
	MTRUnitTestingClusterTestSimpleArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass = _MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestSimpleArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestSimpleArgumentRequestParams */
// An interface definition for the [MTRUnitTestingClusterTestSimpleArgumentRequestParams] class.
type IMTRUnitTestingClusterTestSimpleArgumentRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestSimpleArgumentRequestParams */
	// properties:
	Arg1() objc.IObject /* cross-framework: NSNumber */
	SetArg1(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestSimpleArgumentRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestSimpleArgumentRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestSimpleArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass) New() MTRUnitTestingClusterTestSimpleArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) Init() MTRUnitTestingClusterTestSimpleArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestSimpleArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestSimpleArgumentRequestParams creates a new MTRUnitTestingClusterTestSimpleArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestSimpleArgumentRequestParams() MTRUnitTestingClusterTestSimpleArgumentRequestParams {
	return getMTRUnitTestingClusterTestSimpleArgumentRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestSimpleArgumentRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentRequestParams
type MTRUnitTestingClusterTestSimpleArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestSimpleArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestSimpleArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestSimpleArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestSimpleArgumentRequestParams {
	return MTRUnitTestingClusterTestSimpleArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestSimpleArgumentRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestSimpleArgumentRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestSimpleArgumentRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestSimpleArgumentRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestSimpleArgumentRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) Arg1() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) SetArg1(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestSimpleArgumentRequestParams */



