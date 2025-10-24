// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROvenModeClusterChangeToModeResponseParams */


/* debug [class_header]: Header for MTROvenModeClusterChangeToModeResponseParams */
// The class instance for the [MTROvenModeClusterChangeToModeResponseParams] class.
var (
	MTROvenModeClusterChangeToModeResponseParamsClass     _MTROvenModeClusterChangeToModeResponseParamsClass
	MTROvenModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTROvenModeClusterChangeToModeResponseParamsClass() _MTROvenModeClusterChangeToModeResponseParamsClass {
	MTROvenModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTROvenModeClusterChangeToModeResponseParamsClass = _MTROvenModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTROvenModeClusterChangeToModeResponseParams")}
	})
	return MTROvenModeClusterChangeToModeResponseParamsClass
}

type _MTROvenModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROvenModeClusterChangeToModeResponseParams */
// An interface definition for the [MTROvenModeClusterChangeToModeResponseParams] class.
type IMTROvenModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROvenModeClusterChangeToModeResponseParams */
	// properties:
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROvenModeClusterChangeToModeResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROvenModeClusterChangeToModeResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROvenModeClusterChangeToModeResponseParamsClass) Alloc() MTROvenModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROvenModeClusterChangeToModeResponseParamsClass) New() MTROvenModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenModeClusterChangeToModeResponseParams) Init() MTROvenModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenModeClusterChangeToModeResponseParams) Autorelease() MTROvenModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenModeClusterChangeToModeResponseParams creates a new MTROvenModeClusterChangeToModeResponseParams instance.
func NewMTROvenModeClusterChangeToModeResponseParams() MTROvenModeClusterChangeToModeResponseParams {
	return getMTROvenModeClusterChangeToModeResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROvenModeClusterChangeToModeResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeResponseParams
type MTROvenModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTROvenModeClusterChangeToModeResponseParamsFrom constructs a [MTROvenModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTROvenModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTROvenModeClusterChangeToModeResponseParams {
	return MTROvenModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROvenModeClusterChangeToModeResponseParams */

// Initialize an MTROvenModeClusterChangeToModeResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeResponseParams/init(responseValue:)
func NewMTROvenModeClusterChangeToModeResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTROvenModeClusterChangeToModeResponseParams {
	instance := getMTROvenModeClusterChangeToModeResponseParamsClass().Alloc()
	rv := objc.Send[MTROvenModeClusterChangeToModeResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROvenModeClusterChangeToModeResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROvenModeClusterChangeToModeResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROvenModeClusterChangeToModeResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROvenModeClusterChangeToModeResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROvenModeClusterChangeToModeResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovenmodeclusterchangetomoderesponseparams/status
func (m_ MTROvenModeClusterChangeToModeResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovenmodeclusterchangetomoderesponseparams/status
func (m_ MTROvenModeClusterChangeToModeResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovenmodeclusterchangetomoderesponseparams/statustext
func (m_ MTROvenModeClusterChangeToModeResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}/* debug [instance_properties/getter]: statusText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovenmodeclusterchangetomoderesponseparams/statustext
func (m_ MTROvenModeClusterChangeToModeResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}/* debug [instance_properties/setter]: statusText */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROvenModeClusterChangeToModeResponseParams */


