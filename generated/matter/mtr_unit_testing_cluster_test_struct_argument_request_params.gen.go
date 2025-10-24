// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestStructArgumentRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestStructArgumentRequestParams */
// The class instance for the [MTRUnitTestingClusterTestStructArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestStructArgumentRequestParamsClass     _MTRUnitTestingClusterTestStructArgumentRequestParamsClass
	MTRUnitTestingClusterTestStructArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestStructArgumentRequestParamsClass() _MTRUnitTestingClusterTestStructArgumentRequestParamsClass {
	MTRUnitTestingClusterTestStructArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestStructArgumentRequestParamsClass = _MTRUnitTestingClusterTestStructArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestStructArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestStructArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestStructArgumentRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestStructArgumentRequestParams */
// An interface definition for the [MTRUnitTestingClusterTestStructArgumentRequestParams] class.
type IMTRUnitTestingClusterTestStructArgumentRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestStructArgumentRequestParams */
	// properties:
	Arg1() IMTRUnitTestingClusterSimpleStruct
	SetArg1(value IMTRUnitTestingClusterSimpleStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestStructArgumentRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestStructArgumentRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestStructArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestStructArgumentRequestParamsClass) New() MTRUnitTestingClusterTestStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestStructArgumentRequestParams) Init() MTRUnitTestingClusterTestStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestStructArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestStructArgumentRequestParams creates a new MTRUnitTestingClusterTestStructArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestStructArgumentRequestParams() MTRUnitTestingClusterTestStructArgumentRequestParams {
	return getMTRUnitTestingClusterTestStructArgumentRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestStructArgumentRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArgumentRequestParams
type MTRUnitTestingClusterTestStructArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestStructArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestStructArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestStructArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestStructArgumentRequestParams {
	return MTRUnitTestingClusterTestStructArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestStructArgumentRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestStructArgumentRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestStructArgumentRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestStructArgumentRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestStructArgumentRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestStructArgumentRequestParams) Arg1() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestStructArgumentRequestParams) SetArg1(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestStructArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestStructArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestStructArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestStructArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestStructArgumentRequestParams */



