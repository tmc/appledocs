// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRLaundryWasherModeClusterChangeToModeResponseParams */


/* debug [class_header]: Header for MTRLaundryWasherModeClusterChangeToModeResponseParams */
// The class instance for the [MTRLaundryWasherModeClusterChangeToModeResponseParams] class.
var (
	MTRLaundryWasherModeClusterChangeToModeResponseParamsClass     _MTRLaundryWasherModeClusterChangeToModeResponseParamsClass
	MTRLaundryWasherModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTRLaundryWasherModeClusterChangeToModeResponseParamsClass() _MTRLaundryWasherModeClusterChangeToModeResponseParamsClass {
	MTRLaundryWasherModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTRLaundryWasherModeClusterChangeToModeResponseParamsClass = _MTRLaundryWasherModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTRLaundryWasherModeClusterChangeToModeResponseParams")}
	})
	return MTRLaundryWasherModeClusterChangeToModeResponseParamsClass
}

type _MTRLaundryWasherModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRLaundryWasherModeClusterChangeToModeResponseParams */
// An interface definition for the [MTRLaundryWasherModeClusterChangeToModeResponseParams] class.
type IMTRLaundryWasherModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRLaundryWasherModeClusterChangeToModeResponseParams */
	// properties:
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRLaundryWasherModeClusterChangeToModeResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRLaundryWasherModeClusterChangeToModeResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRLaundryWasherModeClusterChangeToModeResponseParamsClass) Alloc() MTRLaundryWasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRLaundryWasherModeClusterChangeToModeResponseParamsClass) New() MTRLaundryWasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLaundryWasherModeClusterChangeToModeResponseParams) Init() MTRLaundryWasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLaundryWasherModeClusterChangeToModeResponseParams) Autorelease() MTRLaundryWasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLaundryWasherModeClusterChangeToModeResponseParams creates a new MTRLaundryWasherModeClusterChangeToModeResponseParams instance.
func NewMTRLaundryWasherModeClusterChangeToModeResponseParams() MTRLaundryWasherModeClusterChangeToModeResponseParams {
	return getMTRLaundryWasherModeClusterChangeToModeResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRLaundryWasherModeClusterChangeToModeResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeResponseParams
type MTRLaundryWasherModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTRLaundryWasherModeClusterChangeToModeResponseParamsFrom constructs a [MTRLaundryWasherModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTRLaundryWasherModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTRLaundryWasherModeClusterChangeToModeResponseParams {
	return MTRLaundryWasherModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRLaundryWasherModeClusterChangeToModeResponseParams */

// Initialize an MTRLaundryWasherModeClusterChangeToModeResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeResponseParams/init(responseValue:)
func NewMTRLaundryWasherModeClusterChangeToModeResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRLaundryWasherModeClusterChangeToModeResponseParams {
	instance := getMTRLaundryWasherModeClusterChangeToModeResponseParamsClass().Alloc()
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRLaundryWasherModeClusterChangeToModeResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRLaundryWasherModeClusterChangeToModeResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRLaundryWasherModeClusterChangeToModeResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRLaundryWasherModeClusterChangeToModeResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRLaundryWasherModeClusterChangeToModeResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlaundrywashermodeclusterchangetomoderesponseparams/status
func (m_ MTRLaundryWasherModeClusterChangeToModeResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlaundrywashermodeclusterchangetomoderesponseparams/status
func (m_ MTRLaundryWasherModeClusterChangeToModeResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlaundrywashermodeclusterchangetomoderesponseparams/statustext
func (m_ MTRLaundryWasherModeClusterChangeToModeResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}/* debug [instance_properties/getter]: statusText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlaundrywashermodeclusterchangetomoderesponseparams/statustext
func (m_ MTRLaundryWasherModeClusterChangeToModeResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}/* debug [instance_properties/setter]: statusText */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRLaundryWasherModeClusterChangeToModeResponseParams */


