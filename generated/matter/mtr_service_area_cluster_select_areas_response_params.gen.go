// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRServiceAreaClusterSelectAreasResponseParams */


/* debug [class_header]: Header for MTRServiceAreaClusterSelectAreasResponseParams */
// The class instance for the [MTRServiceAreaClusterSelectAreasResponseParams] class.
var (
	MTRServiceAreaClusterSelectAreasResponseParamsClass     _MTRServiceAreaClusterSelectAreasResponseParamsClass
	MTRServiceAreaClusterSelectAreasResponseParamsClassOnce sync.Once
)

func getMTRServiceAreaClusterSelectAreasResponseParamsClass() _MTRServiceAreaClusterSelectAreasResponseParamsClass {
	MTRServiceAreaClusterSelectAreasResponseParamsClassOnce.Do(func() {
		MTRServiceAreaClusterSelectAreasResponseParamsClass = _MTRServiceAreaClusterSelectAreasResponseParamsClass{objc.GetClass("MTRServiceAreaClusterSelectAreasResponseParams")}
	})
	return MTRServiceAreaClusterSelectAreasResponseParamsClass
}

type _MTRServiceAreaClusterSelectAreasResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRServiceAreaClusterSelectAreasResponseParams */
// An interface definition for the [MTRServiceAreaClusterSelectAreasResponseParams] class.
type IMTRServiceAreaClusterSelectAreasResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRServiceAreaClusterSelectAreasResponseParams */
	// properties:
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRServiceAreaClusterSelectAreasResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRServiceAreaClusterSelectAreasResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterSelectAreasResponseParamsClass) Alloc() MTRServiceAreaClusterSelectAreasResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRServiceAreaClusterSelectAreasResponseParamsClass) New() MTRServiceAreaClusterSelectAreasResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) Init() MTRServiceAreaClusterSelectAreasResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) Autorelease() MTRServiceAreaClusterSelectAreasResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterSelectAreasResponseParams creates a new MTRServiceAreaClusterSelectAreasResponseParams instance.
func NewMTRServiceAreaClusterSelectAreasResponseParams() MTRServiceAreaClusterSelectAreasResponseParams {
	return getMTRServiceAreaClusterSelectAreasResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRServiceAreaClusterSelectAreasResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams
type MTRServiceAreaClusterSelectAreasResponseParams struct {
	objectivec.Object
}

// MTRServiceAreaClusterSelectAreasResponseParamsFrom constructs a [MTRServiceAreaClusterSelectAreasResponseParams] from an unsafe.Pointer.
func MTRServiceAreaClusterSelectAreasResponseParamsFrom(ptr unsafe.Pointer) MTRServiceAreaClusterSelectAreasResponseParams {
	return MTRServiceAreaClusterSelectAreasResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRServiceAreaClusterSelectAreasResponseParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRServiceAreaClusterSelectAreasResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRServiceAreaClusterSelectAreasResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRServiceAreaClusterSelectAreasResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRServiceAreaClusterSelectAreasResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams/statusText
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}/* debug [instance_properties/getter]: statusText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams/statusText
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}/* debug [instance_properties/setter]: statusText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterselectareasresponseparams/status
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterselectareasresponseparams/status
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRServiceAreaClusterSelectAreasResponseParams */



