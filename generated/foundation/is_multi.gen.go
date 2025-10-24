// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class isMulti */


/* debug [class_header]: Header for isMulti */
// The class instance for the [isMulti] class.
var (
	IsMultiClass     _isMultiClass
	IsMultiClassOnce sync.Once
)

func getisMultiClass() _isMultiClass {
	IsMultiClassOnce.Do(func() {
		IsMultiClass = _isMultiClass{objc.GetClass("isMulti")}
	})
	return IsMultiClass
}

type _isMultiClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for isMulti */
// An interface definition for the [isMulti] class.
type IisMulti interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for isMulti */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for isMulti */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for isMulti */
// Alloc allocates a new instance without initialization.
func (ic _isMultiClass) Alloc() isMulti {
	rv := objc.Send[isMulti](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _isMultiClass) New() isMulti {
	rv := objc.Send[isMulti](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ isMulti) Init() isMulti {
	rv := objc.Send[isMulti](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ isMulti) Autorelease() isMulti {
	rv := objc.Send[isMulti](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewisMulti creates a new isMulti instance.
func NewisMulti() isMulti {
	return getisMultiClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for isMulti */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/isMulti
type isMulti struct {
	objectivec.Object
}

// isMultiFrom constructs a [isMulti] from an unsafe.Pointer.
func isMultiFrom(ptr unsafe.Pointer) isMulti {
	return isMulti{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for isMulti *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for isMulti */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for isMulti */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for isMulti */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for isMulti */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class isMulti */



