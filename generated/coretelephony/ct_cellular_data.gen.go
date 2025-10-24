// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CTCellularData */


/* debug [class_header]: Header for CTCellularData */
// The class instance for the [CellularData] class.
var (
	CellularDataClass     _CellularDataClass
	CellularDataClassOnce sync.Once
)

func getCellularDataClass() _CellularDataClass {
	CellularDataClassOnce.Do(func() {
		CellularDataClass = _CellularDataClass{objc.GetClass("CTCellularData")}
	})
	return CellularDataClass
}

type _CellularDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CellularData */
// An interface definition for the [CellularData] class.
type ICellularData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CellularData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CellularData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CellularData */
// Alloc allocates a new instance without initialization.
func (cc _CellularDataClass) Alloc() CellularData {
	rv := objc.Send[CellularData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CellularDataClass) New() CellularData {
	rv := objc.Send[CellularData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CellularData) Init() CellularData {
	rv := objc.Send[CellularData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CellularData) Autorelease() CellularData {
	rv := objc.Send[CellularData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCellularData creates a new CellularData instance.
func NewCellularData() CellularData {
	return getCellularDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CellularData */
// An object indicating whether the app can access cellular data.
//
// This property represents all access to cellular data. If the is , the app cannot use the cellular network.


// An object indicating whether the app can access cellular data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularData
type CellularData struct {
	objectivec.Object
}

// CellularDataFrom constructs a [CellularData] from an unsafe.Pointer.
//
// An object indicating whether the app can access cellular data.
func CellularDataFrom(ptr unsafe.Pointer) CellularData {
	return CellularData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CellularData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CellularData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CellularData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CellularData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CellularData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CTCellularData */


