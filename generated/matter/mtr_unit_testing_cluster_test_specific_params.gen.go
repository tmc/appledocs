// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestSpecificParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestSpecificParams */
// The class instance for the [MTRUnitTestingClusterTestSpecificParams] class.
var (
	MTRUnitTestingClusterTestSpecificParamsClass     _MTRUnitTestingClusterTestSpecificParamsClass
	MTRUnitTestingClusterTestSpecificParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestSpecificParamsClass() _MTRUnitTestingClusterTestSpecificParamsClass {
	MTRUnitTestingClusterTestSpecificParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestSpecificParamsClass = _MTRUnitTestingClusterTestSpecificParamsClass{objc.GetClass("MTRUnitTestingClusterTestSpecificParams")}
	})
	return MTRUnitTestingClusterTestSpecificParamsClass
}

type _MTRUnitTestingClusterTestSpecificParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestSpecificParams */
// An interface definition for the [MTRUnitTestingClusterTestSpecificParams] class.
type IMTRUnitTestingClusterTestSpecificParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestSpecificParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestSpecificParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestSpecificParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestSpecificParamsClass) Alloc() MTRUnitTestingClusterTestSpecificParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestSpecificParamsClass) New() MTRUnitTestingClusterTestSpecificParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestSpecificParams) Init() MTRUnitTestingClusterTestSpecificParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestSpecificParams) Autorelease() MTRUnitTestingClusterTestSpecificParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestSpecificParams creates a new MTRUnitTestingClusterTestSpecificParams instance.
func NewMTRUnitTestingClusterTestSpecificParams() MTRUnitTestingClusterTestSpecificParams {
	return getMTRUnitTestingClusterTestSpecificParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestSpecificParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificParams
type MTRUnitTestingClusterTestSpecificParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestSpecificParamsFrom constructs a [MTRUnitTestingClusterTestSpecificParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestSpecificParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestSpecificParams {
	return MTRUnitTestingClusterTestSpecificParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestSpecificParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestSpecificParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestSpecificParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestSpecificParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestSpecificParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestSpecificParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestSpecificParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestSpecificParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestSpecificParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestSpecificParams */



