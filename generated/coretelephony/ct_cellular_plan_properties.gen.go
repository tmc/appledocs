// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CTCellularPlanProperties */


/* debug [class_header]: Header for CTCellularPlanProperties */
// The class instance for the [CellularPlanProperties] class.
var (
	CellularPlanPropertiesClass     _CellularPlanPropertiesClass
	CellularPlanPropertiesClassOnce sync.Once
)

func getCellularPlanPropertiesClass() _CellularPlanPropertiesClass {
	CellularPlanPropertiesClassOnce.Do(func() {
		CellularPlanPropertiesClass = _CellularPlanPropertiesClass{objc.GetClass("CTCellularPlanProperties")}
	})
	return CellularPlanPropertiesClass
}

type _CellularPlanPropertiesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CellularPlanProperties */
// An interface definition for the [CellularPlanProperties] class.
type ICellularPlanProperties interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CellularPlanProperties */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CellularPlanProperties */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CellularPlanProperties */
// Alloc allocates a new instance without initialization.
func (cc _CellularPlanPropertiesClass) Alloc() CellularPlanProperties {
	rv := objc.Send[CellularPlanProperties](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CellularPlanPropertiesClass) New() CellularPlanProperties {
	rv := objc.Send[CellularPlanProperties](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CellularPlanProperties) Init() CellularPlanProperties {
	rv := objc.Send[CellularPlanProperties](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CellularPlanProperties) Autorelease() CellularPlanProperties {
	rv := objc.Send[CellularPlanProperties](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCellularPlanProperties creates a new CellularPlanProperties instance.
func NewCellularPlanProperties() CellularPlanProperties {
	return getCellularPlanPropertiesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CellularPlanProperties */
// An object you use for an eSIM.


// An object you use for an eSIM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProperties
type CellularPlanProperties struct {
	objectivec.Object
}

// CellularPlanPropertiesFrom constructs a [CellularPlanProperties] from an unsafe.Pointer.
//
// An object you use for an eSIM.
func CellularPlanPropertiesFrom(ptr unsafe.Pointer) CellularPlanProperties {
	return CellularPlanProperties{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CellularPlanProperties *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CellularPlanProperties */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CellularPlanProperties */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CellularPlanProperties */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CellularPlanProperties */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CTCellularPlanProperties */


