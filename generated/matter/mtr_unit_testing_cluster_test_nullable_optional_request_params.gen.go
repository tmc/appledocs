// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestNullableOptionalRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestNullableOptionalRequestParams */
// The class instance for the [MTRUnitTestingClusterTestNullableOptionalRequestParams] class.
var (
	MTRUnitTestingClusterTestNullableOptionalRequestParamsClass     _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass
	MTRUnitTestingClusterTestNullableOptionalRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestNullableOptionalRequestParamsClass() _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass {
	MTRUnitTestingClusterTestNullableOptionalRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestNullableOptionalRequestParamsClass = _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestNullableOptionalRequestParams")}
	})
	return MTRUnitTestingClusterTestNullableOptionalRequestParamsClass
}

type _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestNullableOptionalRequestParams */
// An interface definition for the [MTRUnitTestingClusterTestNullableOptionalRequestParams] class.
type IMTRUnitTestingClusterTestNullableOptionalRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestNullableOptionalRequestParams */
	// properties:
	Arg1() objc.IObject /* cross-framework: NSNumber */
	SetArg1(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestNullableOptionalRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestNullableOptionalRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass) Alloc() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass) New() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) Init() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) Autorelease() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestNullableOptionalRequestParams creates a new MTRUnitTestingClusterTestNullableOptionalRequestParams instance.
func NewMTRUnitTestingClusterTestNullableOptionalRequestParams() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	return getMTRUnitTestingClusterTestNullableOptionalRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestNullableOptionalRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalRequestParams
type MTRUnitTestingClusterTestNullableOptionalRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestNullableOptionalRequestParamsFrom constructs a [MTRUnitTestingClusterTestNullableOptionalRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestNullableOptionalRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestNullableOptionalRequestParams {
	return MTRUnitTestingClusterTestNullableOptionalRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestNullableOptionalRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestNullableOptionalRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestNullableOptionalRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestNullableOptionalRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestNullableOptionalRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalRequestParams/arg1
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) Arg1() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalRequestParams/arg1
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) SetArg1(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestNullableOptionalRequestParams */



