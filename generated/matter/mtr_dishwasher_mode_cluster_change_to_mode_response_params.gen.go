// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDishwasherModeClusterChangeToModeResponseParams */


/* debug [class_header]: Header for MTRDishwasherModeClusterChangeToModeResponseParams */
// The class instance for the [MTRDishwasherModeClusterChangeToModeResponseParams] class.
var (
	MTRDishwasherModeClusterChangeToModeResponseParamsClass     _MTRDishwasherModeClusterChangeToModeResponseParamsClass
	MTRDishwasherModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTRDishwasherModeClusterChangeToModeResponseParamsClass() _MTRDishwasherModeClusterChangeToModeResponseParamsClass {
	MTRDishwasherModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTRDishwasherModeClusterChangeToModeResponseParamsClass = _MTRDishwasherModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTRDishwasherModeClusterChangeToModeResponseParams")}
	})
	return MTRDishwasherModeClusterChangeToModeResponseParamsClass
}

type _MTRDishwasherModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDishwasherModeClusterChangeToModeResponseParams */
// An interface definition for the [MTRDishwasherModeClusterChangeToModeResponseParams] class.
type IMTRDishwasherModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDishwasherModeClusterChangeToModeResponseParams */
	// properties:
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDishwasherModeClusterChangeToModeResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDishwasherModeClusterChangeToModeResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherModeClusterChangeToModeResponseParamsClass) Alloc() MTRDishwasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDishwasherModeClusterChangeToModeResponseParamsClass) New() MTRDishwasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherModeClusterChangeToModeResponseParams) Init() MTRDishwasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherModeClusterChangeToModeResponseParams) Autorelease() MTRDishwasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherModeClusterChangeToModeResponseParams creates a new MTRDishwasherModeClusterChangeToModeResponseParams instance.
func NewMTRDishwasherModeClusterChangeToModeResponseParams() MTRDishwasherModeClusterChangeToModeResponseParams {
	return getMTRDishwasherModeClusterChangeToModeResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDishwasherModeClusterChangeToModeResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeResponseParams
type MTRDishwasherModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTRDishwasherModeClusterChangeToModeResponseParamsFrom constructs a [MTRDishwasherModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTRDishwasherModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTRDishwasherModeClusterChangeToModeResponseParams {
	return MTRDishwasherModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDishwasherModeClusterChangeToModeResponseParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDishwasherModeClusterChangeToModeResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDishwasherModeClusterChangeToModeResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDishwasherModeClusterChangeToModeResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDishwasherModeClusterChangeToModeResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeResponseParams/status
func (m_ MTRDishwasherModeClusterChangeToModeResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeResponseParams/status
func (m_ MTRDishwasherModeClusterChangeToModeResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwashermodeclusterchangetomoderesponseparams/statustext
func (m_ MTRDishwasherModeClusterChangeToModeResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}/* debug [instance_properties/getter]: statusText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwashermodeclusterchangetomoderesponseparams/statustext
func (m_ MTRDishwasherModeClusterChangeToModeResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}/* debug [instance_properties/setter]: statusText */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDishwasherModeClusterChangeToModeResponseParams */



