// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UIViewController */


/* debug [class_header]: Header for UIViewController */
// The class instance for the [ViewController] class.
var (
	ViewControllerClass     _ViewControllerClass
	ViewControllerClassOnce sync.Once
)

func getViewControllerClass() _ViewControllerClass {
	ViewControllerClassOnce.Do(func() {
		ViewControllerClass = _ViewControllerClass{objc.GetClass("UIViewController")}
	})
	return ViewControllerClass
}

type _ViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ViewController */
// An interface definition for the [ViewController] class.
type IViewController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ViewController */
// Alloc allocates a new instance without initialization.
func (vc _ViewControllerClass) Alloc() ViewController {
	rv := objc.Send[ViewController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _ViewControllerClass) New() ViewController {
	rv := objc.Send[ViewController](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ ViewController) Init() ViewController {
	rv := objc.Send[ViewController](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ ViewController) Autorelease() ViewController {
	rv := objc.Send[ViewController](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewViewController creates a new ViewController instance.
func NewViewController() ViewController {
	return getViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ViewController */
// A parent class referenced by other SafariServices classes.


// A parent class referenced by other SafariServices classes. [Full Topic]
type ViewController struct {
	objectivec.Object
}

// ViewControllerFrom constructs a [ViewController] from an unsafe.Pointer.
//
// A parent class referenced by other SafariServices classes.
func ViewControllerFrom(ptr unsafe.Pointer) ViewController {
	return ViewController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class UIViewController */



