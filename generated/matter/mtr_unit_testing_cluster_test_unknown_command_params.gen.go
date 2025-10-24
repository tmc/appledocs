// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestUnknownCommandParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestUnknownCommandParams */
// The class instance for the [MTRUnitTestingClusterTestUnknownCommandParams] class.
var (
	MTRUnitTestingClusterTestUnknownCommandParamsClass     _MTRUnitTestingClusterTestUnknownCommandParamsClass
	MTRUnitTestingClusterTestUnknownCommandParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestUnknownCommandParamsClass() _MTRUnitTestingClusterTestUnknownCommandParamsClass {
	MTRUnitTestingClusterTestUnknownCommandParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestUnknownCommandParamsClass = _MTRUnitTestingClusterTestUnknownCommandParamsClass{objc.GetClass("MTRUnitTestingClusterTestUnknownCommandParams")}
	})
	return MTRUnitTestingClusterTestUnknownCommandParamsClass
}

type _MTRUnitTestingClusterTestUnknownCommandParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestUnknownCommandParams */
// An interface definition for the [MTRUnitTestingClusterTestUnknownCommandParams] class.
type IMTRUnitTestingClusterTestUnknownCommandParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestUnknownCommandParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestUnknownCommandParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestUnknownCommandParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestUnknownCommandParamsClass) Alloc() MTRUnitTestingClusterTestUnknownCommandParams {
	rv := objc.Send[MTRUnitTestingClusterTestUnknownCommandParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestUnknownCommandParamsClass) New() MTRUnitTestingClusterTestUnknownCommandParams {
	rv := objc.Send[MTRUnitTestingClusterTestUnknownCommandParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestUnknownCommandParams) Init() MTRUnitTestingClusterTestUnknownCommandParams {
	rv := objc.Send[MTRUnitTestingClusterTestUnknownCommandParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestUnknownCommandParams) Autorelease() MTRUnitTestingClusterTestUnknownCommandParams {
	rv := objc.Send[MTRUnitTestingClusterTestUnknownCommandParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestUnknownCommandParams creates a new MTRUnitTestingClusterTestUnknownCommandParams instance.
func NewMTRUnitTestingClusterTestUnknownCommandParams() MTRUnitTestingClusterTestUnknownCommandParams {
	return getMTRUnitTestingClusterTestUnknownCommandParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestUnknownCommandParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestUnknownCommandParams
type MTRUnitTestingClusterTestUnknownCommandParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestUnknownCommandParamsFrom constructs a [MTRUnitTestingClusterTestUnknownCommandParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestUnknownCommandParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestUnknownCommandParams {
	return MTRUnitTestingClusterTestUnknownCommandParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestUnknownCommandParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestUnknownCommandParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestUnknownCommandParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestUnknownCommandParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestUnknownCommandParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestUnknownCommandParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestUnknownCommandParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestUnknownCommandParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestUnknownCommandParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestUnknownCommandParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestUnknownCommandParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestUnknownCommandParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestUnknownCommandParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestUnknownCommandParams */



