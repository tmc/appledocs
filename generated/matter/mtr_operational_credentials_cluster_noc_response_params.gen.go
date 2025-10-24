// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterNOCResponseParams */


/* debug [class_header]: Header for MTROperationalCredentialsClusterNOCResponseParams */
// The class instance for the [MTROperationalCredentialsClusterNOCResponseParams] class.
var (
	MTROperationalCredentialsClusterNOCResponseParamsClass     _MTROperationalCredentialsClusterNOCResponseParamsClass
	MTROperationalCredentialsClusterNOCResponseParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterNOCResponseParamsClass() _MTROperationalCredentialsClusterNOCResponseParamsClass {
	MTROperationalCredentialsClusterNOCResponseParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterNOCResponseParamsClass = _MTROperationalCredentialsClusterNOCResponseParamsClass{objc.GetClass("MTROperationalCredentialsClusterNOCResponseParams")}
	})
	return MTROperationalCredentialsClusterNOCResponseParamsClass
}

type _MTROperationalCredentialsClusterNOCResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterNOCResponseParams */
// An interface definition for the [MTROperationalCredentialsClusterNOCResponseParams] class.
type IMTROperationalCredentialsClusterNOCResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterNOCResponseParams */
	// properties:
	DebugText() objc.IObject /* cross-framework: NSString */
	SetDebugText(value objc.IObject /* cross-framework: NSString */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	StatusCode() objc.IObject /* cross-framework: NSNumber */
	SetStatusCode(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterNOCResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterNOCResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterNOCResponseParamsClass) Alloc() MTROperationalCredentialsClusterNOCResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterNOCResponseParamsClass) New() MTROperationalCredentialsClusterNOCResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterNOCResponseParams) Init() MTROperationalCredentialsClusterNOCResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterNOCResponseParams) Autorelease() MTROperationalCredentialsClusterNOCResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterNOCResponseParams creates a new MTROperationalCredentialsClusterNOCResponseParams instance.
func NewMTROperationalCredentialsClusterNOCResponseParams() MTROperationalCredentialsClusterNOCResponseParams {
	return getMTROperationalCredentialsClusterNOCResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterNOCResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCResponseParams
type MTROperationalCredentialsClusterNOCResponseParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterNOCResponseParamsFrom constructs a [MTROperationalCredentialsClusterNOCResponseParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterNOCResponseParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterNOCResponseParams {
	return MTROperationalCredentialsClusterNOCResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterNOCResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCResponseParams/init(responseValue:)
func NewMTROperationalCredentialsClusterNOCResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTROperationalCredentialsClusterNOCResponseParams {
	instance := getMTROperationalCredentialsClusterNOCResponseParamsClass().Alloc()
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROperationalCredentialsClusterNOCResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterNOCResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterNOCResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterNOCResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterNOCResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCResponseParams/debugText
func (m_ MTROperationalCredentialsClusterNOCResponseParams) DebugText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("debugText"))
	return rv
}/* debug [instance_properties/getter]: debugText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCResponseParams/debugText
func (m_ MTROperationalCredentialsClusterNOCResponseParams) SetDebugText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDebugText:"), value)
}/* debug [instance_properties/setter]: debugText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCResponseParams/fabricIndex
func (m_ MTROperationalCredentialsClusterNOCResponseParams) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCResponseParams/fabricIndex
func (m_ MTROperationalCredentialsClusterNOCResponseParams) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCResponseParams/statusCode
func (m_ MTROperationalCredentialsClusterNOCResponseParams) StatusCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("statusCode"))
	return rv
}/* debug [instance_properties/getter]: statusCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCResponseParams/statusCode
func (m_ MTROperationalCredentialsClusterNOCResponseParams) SetStatusCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusCode:"), value)
}/* debug [instance_properties/setter]: statusCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCResponseParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterNOCResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCResponseParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterNOCResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterNOCResponseParams */


