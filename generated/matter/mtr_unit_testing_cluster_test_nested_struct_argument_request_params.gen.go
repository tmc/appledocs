// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestNestedStructArgumentRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestNestedStructArgumentRequestParams */
// The class instance for the [MTRUnitTestingClusterTestNestedStructArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass     _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass
	MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass() _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass {
	MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass = _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestNestedStructArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestNestedStructArgumentRequestParams */
// An interface definition for the [MTRUnitTestingClusterTestNestedStructArgumentRequestParams] class.
type IMTRUnitTestingClusterTestNestedStructArgumentRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestNestedStructArgumentRequestParams */
	// properties:
	Arg1() IMTRUnitTestingClusterNestedStruct
	SetArg1(value IMTRUnitTestingClusterNestedStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestNestedStructArgumentRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestNestedStructArgumentRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass) New() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) Init() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestNestedStructArgumentRequestParams creates a new MTRUnitTestingClusterTestNestedStructArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestNestedStructArgumentRequestParams() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	return getMTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestNestedStructArgumentRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructArgumentRequestParams
type MTRUnitTestingClusterTestNestedStructArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestNestedStructArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestNestedStructArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestNestedStructArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	return MTRUnitTestingClusterTestNestedStructArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestNestedStructArgumentRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestNestedStructArgumentRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestNestedStructArgumentRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestNestedStructArgumentRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestNestedStructArgumentRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) Arg1() IMTRUnitTestingClusterNestedStruct {
	rv := objc.Send[MTRUnitTestingClusterNestedStruct](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) SetArg1(value IMTRUnitTestingClusterNestedStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestNestedStructArgumentRequestParams */



