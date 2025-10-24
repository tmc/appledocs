// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestAddArgumentsParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestAddArgumentsParams */
// The class instance for the [MTRUnitTestingClusterTestAddArgumentsParams] class.
var (
	MTRUnitTestingClusterTestAddArgumentsParamsClass     _MTRUnitTestingClusterTestAddArgumentsParamsClass
	MTRUnitTestingClusterTestAddArgumentsParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestAddArgumentsParamsClass() _MTRUnitTestingClusterTestAddArgumentsParamsClass {
	MTRUnitTestingClusterTestAddArgumentsParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestAddArgumentsParamsClass = _MTRUnitTestingClusterTestAddArgumentsParamsClass{objc.GetClass("MTRUnitTestingClusterTestAddArgumentsParams")}
	})
	return MTRUnitTestingClusterTestAddArgumentsParamsClass
}

type _MTRUnitTestingClusterTestAddArgumentsParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestAddArgumentsParams */
// An interface definition for the [MTRUnitTestingClusterTestAddArgumentsParams] class.
type IMTRUnitTestingClusterTestAddArgumentsParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestAddArgumentsParams */
	// properties:
	Arg1() objc.IObject /* cross-framework: NSNumber */
	SetArg1(value objc.IObject /* cross-framework: NSNumber */)
	Arg2() objc.IObject /* cross-framework: NSNumber */
	SetArg2(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestAddArgumentsParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestAddArgumentsParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestAddArgumentsParamsClass) Alloc() MTRUnitTestingClusterTestAddArgumentsParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestAddArgumentsParamsClass) New() MTRUnitTestingClusterTestAddArgumentsParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestAddArgumentsParams) Init() MTRUnitTestingClusterTestAddArgumentsParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestAddArgumentsParams) Autorelease() MTRUnitTestingClusterTestAddArgumentsParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestAddArgumentsParams creates a new MTRUnitTestingClusterTestAddArgumentsParams instance.
func NewMTRUnitTestingClusterTestAddArgumentsParams() MTRUnitTestingClusterTestAddArgumentsParams {
	return getMTRUnitTestingClusterTestAddArgumentsParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestAddArgumentsParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsParams
type MTRUnitTestingClusterTestAddArgumentsParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestAddArgumentsParamsFrom constructs a [MTRUnitTestingClusterTestAddArgumentsParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestAddArgumentsParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestAddArgumentsParams {
	return MTRUnitTestingClusterTestAddArgumentsParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestAddArgumentsParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestAddArgumentsParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestAddArgumentsParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestAddArgumentsParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestAddArgumentsParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsParams/arg1
func (m_ MTRUnitTestingClusterTestAddArgumentsParams) Arg1() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsParams/arg1
func (m_ MTRUnitTestingClusterTestAddArgumentsParams) SetArg1(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsParams/arg2
func (m_ MTRUnitTestingClusterTestAddArgumentsParams) Arg2() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg2"))
	return rv
}/* debug [instance_properties/getter]: arg2 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsParams/arg2
func (m_ MTRUnitTestingClusterTestAddArgumentsParams) SetArg2(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}/* debug [instance_properties/setter]: arg2 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestAddArgumentsParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestAddArgumentsParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestAddArgumentsParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestAddArgumentsParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestAddArgumentsParams */



