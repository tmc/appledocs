// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterAtomicResponseParams */


/* debug [class_header]: Header for MTRThermostatClusterAtomicResponseParams */
// The class instance for the [MTRThermostatClusterAtomicResponseParams] class.
var (
	MTRThermostatClusterAtomicResponseParamsClass     _MTRThermostatClusterAtomicResponseParamsClass
	MTRThermostatClusterAtomicResponseParamsClassOnce sync.Once
)

func getMTRThermostatClusterAtomicResponseParamsClass() _MTRThermostatClusterAtomicResponseParamsClass {
	MTRThermostatClusterAtomicResponseParamsClassOnce.Do(func() {
		MTRThermostatClusterAtomicResponseParamsClass = _MTRThermostatClusterAtomicResponseParamsClass{objc.GetClass("MTRThermostatClusterAtomicResponseParams")}
	})
	return MTRThermostatClusterAtomicResponseParamsClass
}

type _MTRThermostatClusterAtomicResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterAtomicResponseParams */
// An interface definition for the [MTRThermostatClusterAtomicResponseParams] class.
type IMTRThermostatClusterAtomicResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterAtomicResponseParams */
	// properties:
	AttributeStatus() objc.IObject /* cross-framework: NSArray */
	SetAttributeStatus(value objc.IObject /* cross-framework: NSArray */)
	StatusCode() objc.IObject /* cross-framework: NSNumber */
	SetStatusCode(value objc.IObject /* cross-framework: NSNumber */)
	Timeout() objc.IObject /* cross-framework: NSNumber */
	SetTimeout(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterAtomicResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterAtomicResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterAtomicResponseParamsClass) Alloc() MTRThermostatClusterAtomicResponseParams {
	rv := objc.Send[MTRThermostatClusterAtomicResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterAtomicResponseParamsClass) New() MTRThermostatClusterAtomicResponseParams {
	rv := objc.Send[MTRThermostatClusterAtomicResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterAtomicResponseParams) Init() MTRThermostatClusterAtomicResponseParams {
	rv := objc.Send[MTRThermostatClusterAtomicResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterAtomicResponseParams) Autorelease() MTRThermostatClusterAtomicResponseParams {
	rv := objc.Send[MTRThermostatClusterAtomicResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterAtomicResponseParams creates a new MTRThermostatClusterAtomicResponseParams instance.
func NewMTRThermostatClusterAtomicResponseParams() MTRThermostatClusterAtomicResponseParams {
	return getMTRThermostatClusterAtomicResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterAtomicResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams
type MTRThermostatClusterAtomicResponseParams struct {
	objectivec.Object
}

// MTRThermostatClusterAtomicResponseParamsFrom constructs a [MTRThermostatClusterAtomicResponseParams] from an unsafe.Pointer.
func MTRThermostatClusterAtomicResponseParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterAtomicResponseParams {
	return MTRThermostatClusterAtomicResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterAtomicResponseParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterAtomicResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterAtomicResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterAtomicResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterAtomicResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/attributeStatus
func (m_ MTRThermostatClusterAtomicResponseParams) AttributeStatus() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("attributeStatus"))
	return rv
}/* debug [instance_properties/getter]: attributeStatus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/attributeStatus
func (m_ MTRThermostatClusterAtomicResponseParams) SetAttributeStatus(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributeStatus:"), value)
}/* debug [instance_properties/setter]: attributeStatus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusteratomicresponseparams/statuscode
func (m_ MTRThermostatClusterAtomicResponseParams) StatusCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("statusCode"))
	return rv
}/* debug [instance_properties/getter]: statusCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusteratomicresponseparams/statuscode
func (m_ MTRThermostatClusterAtomicResponseParams) SetStatusCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusCode:"), value)
}/* debug [instance_properties/setter]: statusCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusteratomicresponseparams/timeout
func (m_ MTRThermostatClusterAtomicResponseParams) Timeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timeout"))
	return rv
}/* debug [instance_properties/getter]: timeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusteratomicresponseparams/timeout
func (m_ MTRThermostatClusterAtomicResponseParams) SetTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeout:"), value)
}/* debug [instance_properties/setter]: timeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterAtomicResponseParams */



