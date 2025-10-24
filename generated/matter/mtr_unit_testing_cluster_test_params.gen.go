// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestParams */
// The class instance for the [MTRUnitTestingClusterTestParams] class.
var (
	MTRUnitTestingClusterTestParamsClass     _MTRUnitTestingClusterTestParamsClass
	MTRUnitTestingClusterTestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestParamsClass() _MTRUnitTestingClusterTestParamsClass {
	MTRUnitTestingClusterTestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestParamsClass = _MTRUnitTestingClusterTestParamsClass{objc.GetClass("MTRUnitTestingClusterTestParams")}
	})
	return MTRUnitTestingClusterTestParamsClass
}

type _MTRUnitTestingClusterTestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestParams */
// An interface definition for the [MTRUnitTestingClusterTestParams] class.
type IMTRUnitTestingClusterTestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestParamsClass) Alloc() MTRUnitTestingClusterTestParams {
	rv := objc.Send[MTRUnitTestingClusterTestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestParamsClass) New() MTRUnitTestingClusterTestParams {
	rv := objc.Send[MTRUnitTestingClusterTestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestParams) Init() MTRUnitTestingClusterTestParams {
	rv := objc.Send[MTRUnitTestingClusterTestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestParams) Autorelease() MTRUnitTestingClusterTestParams {
	rv := objc.Send[MTRUnitTestingClusterTestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestParams creates a new MTRUnitTestingClusterTestParams instance.
func NewMTRUnitTestingClusterTestParams() MTRUnitTestingClusterTestParams {
	return getMTRUnitTestingClusterTestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestParams
type MTRUnitTestingClusterTestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestParamsFrom constructs a [MTRUnitTestingClusterTestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestParams {
	return MTRUnitTestingClusterTestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestParams */



