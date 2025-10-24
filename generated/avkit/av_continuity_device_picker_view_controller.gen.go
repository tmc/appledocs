// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVContinuityDevicePickerViewController */


/* debug [class_header]: Header for AVContinuityDevicePickerViewController */
// The class instance for the [ContinuityDevicePickerViewController] class.
var (
	ContinuityDevicePickerViewControllerClass     _ContinuityDevicePickerViewControllerClass
	ContinuityDevicePickerViewControllerClassOnce sync.Once
)

func getContinuityDevicePickerViewControllerClass() _ContinuityDevicePickerViewControllerClass {
	ContinuityDevicePickerViewControllerClassOnce.Do(func() {
		ContinuityDevicePickerViewControllerClass = _ContinuityDevicePickerViewControllerClass{objc.GetClass("AVContinuityDevicePickerViewController")}
	})
	return ContinuityDevicePickerViewControllerClass
}

type _ContinuityDevicePickerViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContinuityDevicePickerViewController */
// An interface definition for the [ContinuityDevicePickerViewController] class.
type IContinuityDevicePickerViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for ContinuityDevicePickerViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContinuityDevicePickerViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContinuityDevicePickerViewController */
// Alloc allocates a new instance without initialization.
func (cc _ContinuityDevicePickerViewControllerClass) Alloc() ContinuityDevicePickerViewController {
	rv := objc.Send[ContinuityDevicePickerViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContinuityDevicePickerViewControllerClass) New() ContinuityDevicePickerViewController {
	rv := objc.Send[ContinuityDevicePickerViewController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContinuityDevicePickerViewController) Init() ContinuityDevicePickerViewController {
	rv := objc.Send[ContinuityDevicePickerViewController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContinuityDevicePickerViewController) Autorelease() ContinuityDevicePickerViewController {
	rv := objc.Send[ContinuityDevicePickerViewController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContinuityDevicePickerViewController creates a new ContinuityDevicePickerViewController instance.
func NewContinuityDevicePickerViewController() ContinuityDevicePickerViewController {
	return getContinuityDevicePickerViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContinuityDevicePickerViewController */
// A view controller that provides an interface to a person so they can select and connect a continuity device to the system.
//
// The view controller presents an interface on an Apple TV that lets a person choose a nearby continuity device ( ). Your app can then connect to that device’s cameras and microphones (see and , respectively). To respond to the various outcome events from the picker, your app needs to implement the and assign it to the picker’s property.


// A view controller that provides an interface to a person so they can select and connect a continuity device to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContinuityDevicePickerViewController
type ContinuityDevicePickerViewController struct {
	ViewController
}

// ContinuityDevicePickerViewControllerFrom constructs a [ContinuityDevicePickerViewController] from an unsafe.Pointer.
//
// A view controller that provides an interface to a person so they can select and connect a continuity device to the system.
func ContinuityDevicePickerViewControllerFrom(ptr unsafe.Pointer) ContinuityDevicePickerViewController {
	return ContinuityDevicePickerViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContinuityDevicePickerViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContinuityDevicePickerViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContinuityDevicePickerViewController */

// A Boolean value that indicates whether the system supports connecting to a continuity device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContinuityDevicePickerViewController/isSupported
func (cc _ContinuityDevicePickerViewControllerClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("supported"))
	return rv
}/* debug [class_properties_class/property]: supported */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContinuityDevicePickerViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContinuityDevicePickerViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVContinuityDevicePickerViewController */


