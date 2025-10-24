// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestNestedStructListArgumentRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestNestedStructListArgumentRequestParams */
// The class instance for the [MTRUnitTestingClusterTestNestedStructListArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass     _MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass
	MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass() _MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass {
	MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass = _MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestNestedStructListArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestNestedStructListArgumentRequestParams */
// An interface definition for the [MTRUnitTestingClusterTestNestedStructListArgumentRequestParams] class.
type IMTRUnitTestingClusterTestNestedStructListArgumentRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestNestedStructListArgumentRequestParams */
	// properties:
	Arg1() IMTRUnitTestingClusterNestedStructList
	SetArg1(value IMTRUnitTestingClusterNestedStructList)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestNestedStructListArgumentRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestNestedStructListArgumentRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructListArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass) New() MTRUnitTestingClusterTestNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructListArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) Init() MTRUnitTestingClusterTestNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructListArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructListArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestNestedStructListArgumentRequestParams creates a new MTRUnitTestingClusterTestNestedStructListArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestNestedStructListArgumentRequestParams() MTRUnitTestingClusterTestNestedStructListArgumentRequestParams {
	return getMTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestNestedStructListArgumentRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructListArgumentRequestParams
type MTRUnitTestingClusterTestNestedStructListArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestNestedStructListArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestNestedStructListArgumentRequestParams {
	return MTRUnitTestingClusterTestNestedStructListArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestNestedStructListArgumentRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestNestedStructListArgumentRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestNestedStructListArgumentRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestNestedStructListArgumentRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestNestedStructListArgumentRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructListArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) Arg1() IMTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructListArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) SetArg1(value IMTRUnitTestingClusterNestedStructList) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructListArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructListArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructListArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructListArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestNestedStructListArgumentRequestParams */



