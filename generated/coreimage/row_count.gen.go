// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class rowCount */


/* debug [class_header]: Header for rowCount */
// The class instance for the [rowCount] class.
var (
	RowCountClass     _rowCountClass
	RowCountClassOnce sync.Once
)

func getrowCountClass() _rowCountClass {
	RowCountClassOnce.Do(func() {
		RowCountClass = _rowCountClass{objc.GetClass("rowCount")}
	})
	return RowCountClass
}

type _rowCountClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for rowCount */
// An interface definition for the [rowCount] class.
type IrowCount interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for rowCount */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for rowCount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for rowCount */
// Alloc allocates a new instance without initialization.
func (rc _rowCountClass) Alloc() rowCount {
	rv := objc.Send[rowCount](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _rowCountClass) New() rowCount {
	rv := objc.Send[rowCount](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ rowCount) Init() rowCount {
	rv := objc.Send[rowCount](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ rowCount) Autorelease() rowCount {
	rv := objc.Send[rowCount](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrowCount creates a new rowCount instance.
func NewrowCount() rowCount {
	return getrowCountClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for rowCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/rowCount-c.ivar
type rowCount struct {
	objectivec.Object
}

// rowCountFrom constructs a [rowCount] from an unsafe.Pointer.
func rowCountFrom(ptr unsafe.Pointer) rowCount {
	return rowCount{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for rowCount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for rowCount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for rowCount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for rowCount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for rowCount */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class rowCount */



