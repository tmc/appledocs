// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class width */


/* debug [class_header]: Header for width */
// The class instance for the [width] class.
var (
	WidthClass     _widthClass
	WidthClassOnce sync.Once
)

func getwidthClass() _widthClass {
	WidthClassOnce.Do(func() {
		WidthClass = _widthClass{objc.GetClass("width")}
	})
	return WidthClass
}

type _widthClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for width */
// An interface definition for the [width] class.
type Iwidth interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for width */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for width */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for width */
// Alloc allocates a new instance without initialization.
func (wc _widthClass) Alloc() width {
	rv := objc.Send[width](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _widthClass) New() width {
	rv := objc.Send[width](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ width) Init() width {
	rv := objc.Send[width](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ width) Autorelease() width {
	rv := objc.Send[width](w_.ID, objc.Sel("autorelease"))
	return rv
}

// Newwidth creates a new width instance.
func Newwidth() width {
	return getwidthClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/width-c.ivar
type width struct {
	objectivec.Object
}

// widthFrom constructs a [width] from an unsafe.Pointer.
func widthFrom(ptr unsafe.Pointer) width {
	return width{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for width *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for width */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for width */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for width */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for width */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class width */



