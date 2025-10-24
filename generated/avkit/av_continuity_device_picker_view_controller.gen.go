// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [ContinuityDevicePickerViewController] class.
type IContinuityDevicePickerViewController interface {
	IViewController
	

	// properties:


	

	// methods:


}





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















// A Boolean value that indicates whether the system supports connecting to a continuity device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContinuityDevicePickerViewController/isSupported
func (cc _ContinuityDevicePickerViewControllerClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("supported"))
	return rv
}
















