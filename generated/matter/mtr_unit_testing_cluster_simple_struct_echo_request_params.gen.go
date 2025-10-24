// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterSimpleStructEchoRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterSimpleStructEchoRequestParams */
// The class instance for the [MTRUnitTestingClusterSimpleStructEchoRequestParams] class.
var (
	MTRUnitTestingClusterSimpleStructEchoRequestParamsClass     _MTRUnitTestingClusterSimpleStructEchoRequestParamsClass
	MTRUnitTestingClusterSimpleStructEchoRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterSimpleStructEchoRequestParamsClass() _MTRUnitTestingClusterSimpleStructEchoRequestParamsClass {
	MTRUnitTestingClusterSimpleStructEchoRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterSimpleStructEchoRequestParamsClass = _MTRUnitTestingClusterSimpleStructEchoRequestParamsClass{objc.GetClass("MTRUnitTestingClusterSimpleStructEchoRequestParams")}
	})
	return MTRUnitTestingClusterSimpleStructEchoRequestParamsClass
}

type _MTRUnitTestingClusterSimpleStructEchoRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterSimpleStructEchoRequestParams */
// An interface definition for the [MTRUnitTestingClusterSimpleStructEchoRequestParams] class.
type IMTRUnitTestingClusterSimpleStructEchoRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterSimpleStructEchoRequestParams */
	// properties:
	Arg1() IMTRUnitTestingClusterSimpleStruct
	SetArg1(value IMTRUnitTestingClusterSimpleStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterSimpleStructEchoRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterSimpleStructEchoRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterSimpleStructEchoRequestParamsClass) Alloc() MTRUnitTestingClusterSimpleStructEchoRequestParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructEchoRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterSimpleStructEchoRequestParamsClass) New() MTRUnitTestingClusterSimpleStructEchoRequestParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructEchoRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterSimpleStructEchoRequestParams) Init() MTRUnitTestingClusterSimpleStructEchoRequestParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructEchoRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterSimpleStructEchoRequestParams) Autorelease() MTRUnitTestingClusterSimpleStructEchoRequestParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructEchoRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterSimpleStructEchoRequestParams creates a new MTRUnitTestingClusterSimpleStructEchoRequestParams instance.
func NewMTRUnitTestingClusterSimpleStructEchoRequestParams() MTRUnitTestingClusterSimpleStructEchoRequestParams {
	return getMTRUnitTestingClusterSimpleStructEchoRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterSimpleStructEchoRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructEchoRequestParams
type MTRUnitTestingClusterSimpleStructEchoRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterSimpleStructEchoRequestParamsFrom constructs a [MTRUnitTestingClusterSimpleStructEchoRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterSimpleStructEchoRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterSimpleStructEchoRequestParams {
	return MTRUnitTestingClusterSimpleStructEchoRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterSimpleStructEchoRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterSimpleStructEchoRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterSimpleStructEchoRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterSimpleStructEchoRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterSimpleStructEchoRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructEchoRequestParams/arg1
func (m_ MTRUnitTestingClusterSimpleStructEchoRequestParams) Arg1() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructEchoRequestParams/arg1
func (m_ MTRUnitTestingClusterSimpleStructEchoRequestParams) SetArg1(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructEchoRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterSimpleStructEchoRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructEchoRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterSimpleStructEchoRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructEchoRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterSimpleStructEchoRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructEchoRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterSimpleStructEchoRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterSimpleStructEchoRequestParams */



