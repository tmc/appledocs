// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestListStructArgumentRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestListStructArgumentRequestParams */
// The class instance for the [MTRUnitTestingClusterTestListStructArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestListStructArgumentRequestParamsClass     _MTRUnitTestingClusterTestListStructArgumentRequestParamsClass
	MTRUnitTestingClusterTestListStructArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestListStructArgumentRequestParamsClass() _MTRUnitTestingClusterTestListStructArgumentRequestParamsClass {
	MTRUnitTestingClusterTestListStructArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestListStructArgumentRequestParamsClass = _MTRUnitTestingClusterTestListStructArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestListStructArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestListStructArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestListStructArgumentRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestListStructArgumentRequestParams */
// An interface definition for the [MTRUnitTestingClusterTestListStructArgumentRequestParams] class.
type IMTRUnitTestingClusterTestListStructArgumentRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestListStructArgumentRequestParams */
	// properties:
	Arg1() objc.IObject /* cross-framework: NSArray */
	SetArg1(value objc.IObject /* cross-framework: NSArray */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestListStructArgumentRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestListStructArgumentRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestListStructArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestListStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestListStructArgumentRequestParamsClass) New() MTRUnitTestingClusterTestListStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) Init() MTRUnitTestingClusterTestListStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListStructArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestListStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListStructArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestListStructArgumentRequestParams creates a new MTRUnitTestingClusterTestListStructArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestListStructArgumentRequestParams() MTRUnitTestingClusterTestListStructArgumentRequestParams {
	return getMTRUnitTestingClusterTestListStructArgumentRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestListStructArgumentRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructArgumentRequestParams
type MTRUnitTestingClusterTestListStructArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestListStructArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestListStructArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestListStructArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestListStructArgumentRequestParams {
	return MTRUnitTestingClusterTestListStructArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestListStructArgumentRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestListStructArgumentRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestListStructArgumentRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestListStructArgumentRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestListStructArgumentRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) Arg1() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) SetArg1(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestListStructArgumentRequestParams */



