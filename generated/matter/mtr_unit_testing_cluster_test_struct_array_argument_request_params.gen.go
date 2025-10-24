// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestStructArrayArgumentRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestStructArrayArgumentRequestParams */
// The class instance for the [MTRUnitTestingClusterTestStructArrayArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass     _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass
	MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass() _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass {
	MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass = _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestStructArrayArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestStructArrayArgumentRequestParams */
// An interface definition for the [MTRUnitTestingClusterTestStructArrayArgumentRequestParams] class.
type IMTRUnitTestingClusterTestStructArrayArgumentRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestStructArrayArgumentRequestParams */
	// properties:
	Arg1() objc.IObject /* cross-framework: NSArray */
	SetArg1(value objc.IObject /* cross-framework: NSArray */)
	Arg2() objc.IObject /* cross-framework: NSArray */
	SetArg2(value objc.IObject /* cross-framework: NSArray */)
	Arg3() objc.IObject /* cross-framework: NSArray */
	SetArg3(value objc.IObject /* cross-framework: NSArray */)
	Arg4() objc.IObject /* cross-framework: NSArray */
	SetArg4(value objc.IObject /* cross-framework: NSArray */)
	Arg5() objc.IObject /* cross-framework: NSNumber */
	SetArg5(value objc.IObject /* cross-framework: NSNumber */)
	Arg6() objc.IObject /* cross-framework: NSNumber */
	SetArg6(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestStructArrayArgumentRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestStructArrayArgumentRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass) New() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Init() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestStructArrayArgumentRequestParams creates a new MTRUnitTestingClusterTestStructArrayArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestStructArrayArgumentRequestParams() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	return getMTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestStructArrayArgumentRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams
type MTRUnitTestingClusterTestStructArrayArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestStructArrayArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestStructArrayArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestStructArrayArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	return MTRUnitTestingClusterTestStructArrayArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestStructArrayArgumentRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestStructArrayArgumentRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestStructArrayArgumentRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestStructArrayArgumentRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestStructArrayArgumentRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Arg1() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetArg1(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/arg2
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Arg2() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg2"))
	return rv
}/* debug [instance_properties/getter]: arg2 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/arg2
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetArg2(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}/* debug [instance_properties/setter]: arg2 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/arg3
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Arg3() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg3"))
	return rv
}/* debug [instance_properties/getter]: arg3 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/arg3
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetArg3(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg3:"), value)
}/* debug [instance_properties/setter]: arg3 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/arg4
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Arg4() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg4"))
	return rv
}/* debug [instance_properties/getter]: arg4 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/arg4
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetArg4(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg4:"), value)
}/* debug [instance_properties/setter]: arg4 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/arg5
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Arg5() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg5"))
	return rv
}/* debug [instance_properties/getter]: arg5 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/arg5
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetArg5(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg5:"), value)
}/* debug [instance_properties/setter]: arg5 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/arg6
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Arg6() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg6"))
	return rv
}/* debug [instance_properties/getter]: arg6 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/arg6
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetArg6(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg6:"), value)
}/* debug [instance_properties/setter]: arg6 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestStructArrayArgumentRequestParams */



