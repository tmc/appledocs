// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestListInt8UReverseRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestListInt8UReverseRequestParams */
// The class instance for the [MTRUnitTestingClusterTestListInt8UReverseRequestParams] class.
var (
	MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass     _MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass
	MTRUnitTestingClusterTestListInt8UReverseRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestListInt8UReverseRequestParamsClass() _MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass {
	MTRUnitTestingClusterTestListInt8UReverseRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass = _MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestListInt8UReverseRequestParams")}
	})
	return MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass
}

type _MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestListInt8UReverseRequestParams */
// An interface definition for the [MTRUnitTestingClusterTestListInt8UReverseRequestParams] class.
type IMTRUnitTestingClusterTestListInt8UReverseRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestListInt8UReverseRequestParams */
	// properties:
	Arg1() objc.IObject /* cross-framework: NSArray */
	SetArg1(value objc.IObject /* cross-framework: NSArray */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestListInt8UReverseRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestListInt8UReverseRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass) Alloc() MTRUnitTestingClusterTestListInt8UReverseRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass) New() MTRUnitTestingClusterTestListInt8UReverseRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) Init() MTRUnitTestingClusterTestListInt8UReverseRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) Autorelease() MTRUnitTestingClusterTestListInt8UReverseRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestListInt8UReverseRequestParams creates a new MTRUnitTestingClusterTestListInt8UReverseRequestParams instance.
func NewMTRUnitTestingClusterTestListInt8UReverseRequestParams() MTRUnitTestingClusterTestListInt8UReverseRequestParams {
	return getMTRUnitTestingClusterTestListInt8UReverseRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestListInt8UReverseRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseRequestParams
type MTRUnitTestingClusterTestListInt8UReverseRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestListInt8UReverseRequestParamsFrom constructs a [MTRUnitTestingClusterTestListInt8UReverseRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestListInt8UReverseRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestListInt8UReverseRequestParams {
	return MTRUnitTestingClusterTestListInt8UReverseRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestListInt8UReverseRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestListInt8UReverseRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestListInt8UReverseRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestListInt8UReverseRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestListInt8UReverseRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseRequestParams/arg1
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) Arg1() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseRequestParams/arg1
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) SetArg1(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestListInt8UReverseRequestParams */



