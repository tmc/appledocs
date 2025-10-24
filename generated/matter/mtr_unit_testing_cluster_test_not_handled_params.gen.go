// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestNotHandledParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestNotHandledParams */
// The class instance for the [MTRUnitTestingClusterTestNotHandledParams] class.
var (
	MTRUnitTestingClusterTestNotHandledParamsClass     _MTRUnitTestingClusterTestNotHandledParamsClass
	MTRUnitTestingClusterTestNotHandledParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestNotHandledParamsClass() _MTRUnitTestingClusterTestNotHandledParamsClass {
	MTRUnitTestingClusterTestNotHandledParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestNotHandledParamsClass = _MTRUnitTestingClusterTestNotHandledParamsClass{objc.GetClass("MTRUnitTestingClusterTestNotHandledParams")}
	})
	return MTRUnitTestingClusterTestNotHandledParamsClass
}

type _MTRUnitTestingClusterTestNotHandledParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestNotHandledParams */
// An interface definition for the [MTRUnitTestingClusterTestNotHandledParams] class.
type IMTRUnitTestingClusterTestNotHandledParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestNotHandledParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestNotHandledParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestNotHandledParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestNotHandledParamsClass) Alloc() MTRUnitTestingClusterTestNotHandledParams {
	rv := objc.Send[MTRUnitTestingClusterTestNotHandledParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestNotHandledParamsClass) New() MTRUnitTestingClusterTestNotHandledParams {
	rv := objc.Send[MTRUnitTestingClusterTestNotHandledParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestNotHandledParams) Init() MTRUnitTestingClusterTestNotHandledParams {
	rv := objc.Send[MTRUnitTestingClusterTestNotHandledParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestNotHandledParams) Autorelease() MTRUnitTestingClusterTestNotHandledParams {
	rv := objc.Send[MTRUnitTestingClusterTestNotHandledParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestNotHandledParams creates a new MTRUnitTestingClusterTestNotHandledParams instance.
func NewMTRUnitTestingClusterTestNotHandledParams() MTRUnitTestingClusterTestNotHandledParams {
	return getMTRUnitTestingClusterTestNotHandledParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestNotHandledParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNotHandledParams
type MTRUnitTestingClusterTestNotHandledParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestNotHandledParamsFrom constructs a [MTRUnitTestingClusterTestNotHandledParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestNotHandledParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestNotHandledParams {
	return MTRUnitTestingClusterTestNotHandledParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestNotHandledParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestNotHandledParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestNotHandledParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestNotHandledParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestNotHandledParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNotHandledParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestNotHandledParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNotHandledParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestNotHandledParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNotHandledParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestNotHandledParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNotHandledParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestNotHandledParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestNotHandledParams */



