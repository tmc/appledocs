// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class wantsInvalid */


/* debug [class_header]: Header for wantsInvalid */
// The class instance for the [wantsInvalid] class.
var (
	WantsInvalidClass     _wantsInvalidClass
	WantsInvalidClassOnce sync.Once
)

func getwantsInvalidClass() _wantsInvalidClass {
	WantsInvalidClassOnce.Do(func() {
		WantsInvalidClass = _wantsInvalidClass{objc.GetClass("wantsInvalid")}
	})
	return WantsInvalidClass
}

type _wantsInvalidClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for wantsInvalid */
// An interface definition for the [wantsInvalid] class.
type IwantsInvalid interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for wantsInvalid */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for wantsInvalid */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for wantsInvalid */
// Alloc allocates a new instance without initialization.
func (wc _wantsInvalidClass) Alloc() wantsInvalid {
	rv := objc.Send[wantsInvalid](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _wantsInvalidClass) New() wantsInvalid {
	rv := objc.Send[wantsInvalid](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ wantsInvalid) Init() wantsInvalid {
	rv := objc.Send[wantsInvalid](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ wantsInvalid) Autorelease() wantsInvalid {
	rv := objc.Send[wantsInvalid](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewwantsInvalid creates a new wantsInvalid instance.
func NewwantsInvalid() wantsInvalid {
	return getwantsInvalidClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for wantsInvalid */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/wantsInvalid
type wantsInvalid struct {
	objectivec.Object
}

// wantsInvalidFrom constructs a [wantsInvalid] from an unsafe.Pointer.
func wantsInvalidFrom(ptr unsafe.Pointer) wantsInvalid {
	return wantsInvalid{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for wantsInvalid *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for wantsInvalid */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for wantsInvalid */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for wantsInvalid */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for wantsInvalid */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class wantsInvalid */



