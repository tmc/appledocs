// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRServiceAreaClusterSkipAreaResponseParams */


/* debug [class_header]: Header for MTRServiceAreaClusterSkipAreaResponseParams */
// The class instance for the [MTRServiceAreaClusterSkipAreaResponseParams] class.
var (
	MTRServiceAreaClusterSkipAreaResponseParamsClass     _MTRServiceAreaClusterSkipAreaResponseParamsClass
	MTRServiceAreaClusterSkipAreaResponseParamsClassOnce sync.Once
)

func getMTRServiceAreaClusterSkipAreaResponseParamsClass() _MTRServiceAreaClusterSkipAreaResponseParamsClass {
	MTRServiceAreaClusterSkipAreaResponseParamsClassOnce.Do(func() {
		MTRServiceAreaClusterSkipAreaResponseParamsClass = _MTRServiceAreaClusterSkipAreaResponseParamsClass{objc.GetClass("MTRServiceAreaClusterSkipAreaResponseParams")}
	})
	return MTRServiceAreaClusterSkipAreaResponseParamsClass
}

type _MTRServiceAreaClusterSkipAreaResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRServiceAreaClusterSkipAreaResponseParams */
// An interface definition for the [MTRServiceAreaClusterSkipAreaResponseParams] class.
type IMTRServiceAreaClusterSkipAreaResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRServiceAreaClusterSkipAreaResponseParams */
	// properties:
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRServiceAreaClusterSkipAreaResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRServiceAreaClusterSkipAreaResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterSkipAreaResponseParamsClass) Alloc() MTRServiceAreaClusterSkipAreaResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRServiceAreaClusterSkipAreaResponseParamsClass) New() MTRServiceAreaClusterSkipAreaResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) Init() MTRServiceAreaClusterSkipAreaResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) Autorelease() MTRServiceAreaClusterSkipAreaResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterSkipAreaResponseParams creates a new MTRServiceAreaClusterSkipAreaResponseParams instance.
func NewMTRServiceAreaClusterSkipAreaResponseParams() MTRServiceAreaClusterSkipAreaResponseParams {
	return getMTRServiceAreaClusterSkipAreaResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRServiceAreaClusterSkipAreaResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams
type MTRServiceAreaClusterSkipAreaResponseParams struct {
	objectivec.Object
}

// MTRServiceAreaClusterSkipAreaResponseParamsFrom constructs a [MTRServiceAreaClusterSkipAreaResponseParams] from an unsafe.Pointer.
func MTRServiceAreaClusterSkipAreaResponseParamsFrom(ptr unsafe.Pointer) MTRServiceAreaClusterSkipAreaResponseParams {
	return MTRServiceAreaClusterSkipAreaResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRServiceAreaClusterSkipAreaResponseParams */

// Initialize an MTRServiceAreaClusterSkipAreaResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams/init(responseValue:)
func NewMTRServiceAreaClusterSkipAreaResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRServiceAreaClusterSkipAreaResponseParams {
	instance := getMTRServiceAreaClusterSkipAreaResponseParamsClass().Alloc()
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRServiceAreaClusterSkipAreaResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRServiceAreaClusterSkipAreaResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRServiceAreaClusterSkipAreaResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRServiceAreaClusterSkipAreaResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRServiceAreaClusterSkipAreaResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterskiparearesponseparams/status
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterskiparearesponseparams/status
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterskiparearesponseparams/statustext
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}/* debug [instance_properties/getter]: statusText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterskiparearesponseparams/statustext
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}/* debug [instance_properties/setter]: statusText */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRServiceAreaClusterSkipAreaResponseParams */


