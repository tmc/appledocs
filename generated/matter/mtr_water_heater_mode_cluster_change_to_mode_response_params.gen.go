// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWaterHeaterModeClusterChangeToModeResponseParams */


/* debug [class_header]: Header for MTRWaterHeaterModeClusterChangeToModeResponseParams */
// The class instance for the [MTRWaterHeaterModeClusterChangeToModeResponseParams] class.
var (
	MTRWaterHeaterModeClusterChangeToModeResponseParamsClass     _MTRWaterHeaterModeClusterChangeToModeResponseParamsClass
	MTRWaterHeaterModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTRWaterHeaterModeClusterChangeToModeResponseParamsClass() _MTRWaterHeaterModeClusterChangeToModeResponseParamsClass {
	MTRWaterHeaterModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTRWaterHeaterModeClusterChangeToModeResponseParamsClass = _MTRWaterHeaterModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTRWaterHeaterModeClusterChangeToModeResponseParams")}
	})
	return MTRWaterHeaterModeClusterChangeToModeResponseParamsClass
}

type _MTRWaterHeaterModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWaterHeaterModeClusterChangeToModeResponseParams */
// An interface definition for the [MTRWaterHeaterModeClusterChangeToModeResponseParams] class.
type IMTRWaterHeaterModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWaterHeaterModeClusterChangeToModeResponseParams */
	// properties:
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWaterHeaterModeClusterChangeToModeResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWaterHeaterModeClusterChangeToModeResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterModeClusterChangeToModeResponseParamsClass) Alloc() MTRWaterHeaterModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWaterHeaterModeClusterChangeToModeResponseParamsClass) New() MTRWaterHeaterModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterModeClusterChangeToModeResponseParams) Init() MTRWaterHeaterModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterModeClusterChangeToModeResponseParams) Autorelease() MTRWaterHeaterModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterModeClusterChangeToModeResponseParams creates a new MTRWaterHeaterModeClusterChangeToModeResponseParams instance.
func NewMTRWaterHeaterModeClusterChangeToModeResponseParams() MTRWaterHeaterModeClusterChangeToModeResponseParams {
	return getMTRWaterHeaterModeClusterChangeToModeResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWaterHeaterModeClusterChangeToModeResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeResponseParams
type MTRWaterHeaterModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTRWaterHeaterModeClusterChangeToModeResponseParamsFrom constructs a [MTRWaterHeaterModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTRWaterHeaterModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTRWaterHeaterModeClusterChangeToModeResponseParams {
	return MTRWaterHeaterModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWaterHeaterModeClusterChangeToModeResponseParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWaterHeaterModeClusterChangeToModeResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWaterHeaterModeClusterChangeToModeResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWaterHeaterModeClusterChangeToModeResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWaterHeaterModeClusterChangeToModeResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeResponseParams/status
func (m_ MTRWaterHeaterModeClusterChangeToModeResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeResponseParams/status
func (m_ MTRWaterHeaterModeClusterChangeToModeResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermodeclusterchangetomoderesponseparams/statustext
func (m_ MTRWaterHeaterModeClusterChangeToModeResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}/* debug [instance_properties/getter]: statusText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermodeclusterchangetomoderesponseparams/statustext
func (m_ MTRWaterHeaterModeClusterChangeToModeResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}/* debug [instance_properties/setter]: statusText */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWaterHeaterModeClusterChangeToModeResponseParams */



