// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class columnCount */


/* debug [class_header]: Header for columnCount */
// The class instance for the [columnCount] class.
var (
	ColumnCountClass     _columnCountClass
	ColumnCountClassOnce sync.Once
)

func getcolumnCountClass() _columnCountClass {
	ColumnCountClassOnce.Do(func() {
		ColumnCountClass = _columnCountClass{objc.GetClass("columnCount")}
	})
	return ColumnCountClass
}

type _columnCountClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for columnCount */
// An interface definition for the [columnCount] class.
type IcolumnCount interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for columnCount */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for columnCount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for columnCount */
// Alloc allocates a new instance without initialization.
func (cc _columnCountClass) Alloc() columnCount {
	rv := objc.Send[columnCount](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _columnCountClass) New() columnCount {
	rv := objc.Send[columnCount](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ columnCount) Init() columnCount {
	rv := objc.Send[columnCount](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ columnCount) Autorelease() columnCount {
	rv := objc.Send[columnCount](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewcolumnCount creates a new columnCount instance.
func NewcolumnCount() columnCount {
	return getcolumnCountClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for columnCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/columnCount-c.ivar
type columnCount struct {
	objectivec.Object
}

// columnCountFrom constructs a [columnCount] from an unsafe.Pointer.
func columnCountFrom(ptr unsafe.Pointer) columnCount {
	return columnCount{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for columnCount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for columnCount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for columnCount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for columnCount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for columnCount */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class columnCount */



