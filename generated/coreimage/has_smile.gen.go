// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class hasSmile */


/* debug [class_header]: Header for hasSmile */
// The class instance for the [hasSmile] class.
var (
	HasSmileClass     _hasSmileClass
	HasSmileClassOnce sync.Once
)

func gethasSmileClass() _hasSmileClass {
	HasSmileClassOnce.Do(func() {
		HasSmileClass = _hasSmileClass{objc.GetClass("hasSmile")}
	})
	return HasSmileClass
}

type _hasSmileClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for hasSmile */
// An interface definition for the [hasSmile] class.
type IhasSmile interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for hasSmile */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for hasSmile */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for hasSmile */
// Alloc allocates a new instance without initialization.
func (hc _hasSmileClass) Alloc() hasSmile {
	rv := objc.Send[hasSmile](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _hasSmileClass) New() hasSmile {
	rv := objc.Send[hasSmile](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasSmile) Init() hasSmile {
	rv := objc.Send[hasSmile](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasSmile) Autorelease() hasSmile {
	rv := objc.Send[hasSmile](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasSmile creates a new hasSmile instance.
func NewhasSmile() hasSmile {
	return gethasSmileClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for hasSmile */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasSmile-c.ivar
type hasSmile struct {
	objectivec.Object
}

// hasSmileFrom constructs a [hasSmile] from an unsafe.Pointer.
func hasSmileFrom(ptr unsafe.Pointer) hasSmile {
	return hasSmile{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for hasSmile *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for hasSmile */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for hasSmile */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for hasSmile */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for hasSmile */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class hasSmile */



