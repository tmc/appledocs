// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestStructArrayArgumentResponseParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestStructArrayArgumentResponseParams */
// The class instance for the [MTRUnitTestingClusterTestStructArrayArgumentResponseParams] class.
var (
	MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass     _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass
	MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass() _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass {
	MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass = _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestStructArrayArgumentResponseParams")}
	})
	return MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass
}

type _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestStructArrayArgumentResponseParams */
// An interface definition for the [MTRUnitTestingClusterTestStructArrayArgumentResponseParams] class.
type IMTRUnitTestingClusterTestStructArrayArgumentResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestStructArrayArgumentResponseParams */
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
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestStructArrayArgumentResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestStructArrayArgumentResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass) Alloc() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass) New() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Init() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Autorelease() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestStructArrayArgumentResponseParams creates a new MTRUnitTestingClusterTestStructArrayArgumentResponseParams instance.
func NewMTRUnitTestingClusterTestStructArrayArgumentResponseParams() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	return getMTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestStructArrayArgumentResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams
type MTRUnitTestingClusterTestStructArrayArgumentResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestStructArrayArgumentResponseParamsFrom constructs a [MTRUnitTestingClusterTestStructArrayArgumentResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestStructArrayArgumentResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	return MTRUnitTestingClusterTestStructArrayArgumentResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestStructArrayArgumentResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/init(responseValue:)
func NewMTRUnitTestingClusterTestStructArrayArgumentResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	instance := getMTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass().Alloc()
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRUnitTestingClusterTestStructArrayArgumentResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestStructArrayArgumentResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestStructArrayArgumentResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestStructArrayArgumentResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestStructArrayArgumentResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/arg1
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Arg1() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/arg1
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetArg1(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/arg2
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Arg2() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg2"))
	return rv
}/* debug [instance_properties/getter]: arg2 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/arg2
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetArg2(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}/* debug [instance_properties/setter]: arg2 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/arg3
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Arg3() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg3"))
	return rv
}/* debug [instance_properties/getter]: arg3 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/arg3
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetArg3(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg3:"), value)
}/* debug [instance_properties/setter]: arg3 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/arg4
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Arg4() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg4"))
	return rv
}/* debug [instance_properties/getter]: arg4 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/arg4
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetArg4(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg4:"), value)
}/* debug [instance_properties/setter]: arg4 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/arg5
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Arg5() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg5"))
	return rv
}/* debug [instance_properties/getter]: arg5 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/arg5
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetArg5(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg5:"), value)
}/* debug [instance_properties/setter]: arg5 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/arg6
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Arg6() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg6"))
	return rv
}/* debug [instance_properties/getter]: arg6 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/arg6
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetArg6(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg6:"), value)
}/* debug [instance_properties/setter]: arg6 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestStructArrayArgumentResponseParams */


