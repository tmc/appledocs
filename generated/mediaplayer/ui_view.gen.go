// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UIView */

/* debug [class_header]: Header for UIView */
// The class instance for the [View] class.
var (
	ViewClass     _ViewClass
	ViewClassOnce sync.Once
)

func getViewClass() _ViewClass {
	ViewClassOnce.Do(func() {
		ViewClass = _ViewClass{objc.GetClass("UIView")}
	})
	return ViewClass
}

type _ViewClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for View */
// An interface definition for the [View] class.
type IView interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for View */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for View */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for View */
// Alloc allocates a new instance without initialization.
func (vc _ViewClass) Alloc() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _ViewClass) New() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ View) Init() View {
	rv := objc.Send[View](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ View) Autorelease() View {
	rv := objc.Send[View](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewView creates a new View instance.
func NewView() View {
	return getViewClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for View */
// A parent class referenced by other MediaPlayer classes.

// A parent class referenced by other MediaPlayer classes. [Full Topic]
type View struct {
	objectivec.Object
}

// ViewFrom constructs a [View] from an unsafe.Pointer.
//
// A parent class referenced by other MediaPlayer classes.
func ViewFrom(ptr unsafe.Pointer) View {
	return View{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for View */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for View */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for View */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for View */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for View */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class UIView */
