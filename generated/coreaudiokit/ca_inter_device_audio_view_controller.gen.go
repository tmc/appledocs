// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CAInterDeviceAudioViewController */


/* debug [class_header]: Header for CAInterDeviceAudioViewController */
// The class instance for the [InterDeviceAudioViewController] class.
var (
	InterDeviceAudioViewControllerClass     _InterDeviceAudioViewControllerClass
	InterDeviceAudioViewControllerClassOnce sync.Once
)

func getInterDeviceAudioViewControllerClass() _InterDeviceAudioViewControllerClass {
	InterDeviceAudioViewControllerClassOnce.Do(func() {
		InterDeviceAudioViewControllerClass = _InterDeviceAudioViewControllerClass{objc.GetClass("CAInterDeviceAudioViewController")}
	})
	return InterDeviceAudioViewControllerClass
}

type _InterDeviceAudioViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InterDeviceAudioViewController */
// An interface definition for the [InterDeviceAudioViewController] class.
type IInterDeviceAudioViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for InterDeviceAudioViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InterDeviceAudioViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InterDeviceAudioViewController */
// Alloc allocates a new instance without initialization.
func (ic _InterDeviceAudioViewControllerClass) Alloc() InterDeviceAudioViewController {
	rv := objc.Send[InterDeviceAudioViewController](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InterDeviceAudioViewControllerClass) New() InterDeviceAudioViewController {
	rv := objc.Send[InterDeviceAudioViewController](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InterDeviceAudioViewController) Init() InterDeviceAudioViewController {
	rv := objc.Send[InterDeviceAudioViewController](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InterDeviceAudioViewController) Autorelease() InterDeviceAudioViewController {
	rv := objc.Send[InterDeviceAudioViewController](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInterDeviceAudioViewController creates a new InterDeviceAudioViewController instance.
func NewInterDeviceAudioViewController() InterDeviceAudioViewController {
	return getInterDeviceAudioViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InterDeviceAudioViewController */
// A view controller object that displays iOS devices that support inter-device audio.


// A view controller object that displays iOS devices that support inter-device audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterDeviceAudioViewController
type InterDeviceAudioViewController struct {
	ViewController
}

// InterDeviceAudioViewControllerFrom constructs a [InterDeviceAudioViewController] from an unsafe.Pointer.
//
// A view controller object that displays iOS devices that support inter-device audio.
func InterDeviceAudioViewControllerFrom(ptr unsafe.Pointer) InterDeviceAudioViewController {
	return InterDeviceAudioViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InterDeviceAudioViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InterDeviceAudioViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InterDeviceAudioViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InterDeviceAudioViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InterDeviceAudioViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAInterDeviceAudioViewController */



