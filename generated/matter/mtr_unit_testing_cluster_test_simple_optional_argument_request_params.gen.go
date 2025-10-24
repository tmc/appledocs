// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams */
// The class instance for the [MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass     _MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass
	MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass() _MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass {
	MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass = _MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams */
// An interface definition for the [MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams] class.
type IMTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams */
	// properties:
	Arg1() objc.IObject /* cross-framework: NSNumber */
	SetArg1(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass) New() MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) Init() MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams creates a new MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams() MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams {
	return getMTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams
type MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams {
	return MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) Arg1() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) SetArg1(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams */



