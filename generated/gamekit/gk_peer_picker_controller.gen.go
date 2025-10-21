// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PeerPickerController] class.
var (
	PeerPickerControllerClass     _PeerPickerControllerClass
	PeerPickerControllerClassOnce sync.Once
)

func getPeerPickerControllerClass() _PeerPickerControllerClass {
	PeerPickerControllerClassOnce.Do(func() {
		PeerPickerControllerClass = _PeerPickerControllerClass{objc.GetClass("GKPeerPickerController")}
	})
	return PeerPickerControllerClass
}

type _PeerPickerControllerClass struct {
	class objc.Class
}

// An interface definition for the [PeerPickerController] class.
type IPeerPickerController interface {
	objectivec.IObject
	Dismiss()
	Show()
}

// Provides a standard user interface to allow one iOS device to discover and connect to another.
//
// The result is a configured object connecting the two devices. To use a object, your application creates the controller, adds a delegate, configures the allowed connection types, and then shows the peer picker. The delegate is called as the user makes selections within the peer picker interface. In iOS 3.0, the peer picker can be configured to select between Bluetooth and Internet connections. On iOS 3.0, your application should release the peer picker object after it dismisses the peer picker dialog. On iOS 3.1 or later, your application may release the peer picker after it is shown to the user. If you do this, the peer picker controller is automatically deallocated after the dialog is dismissed.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerController
type PeerPickerController struct {
	objectivec.Object
}

// PeerPickerControllerFrom constructs a [PeerPickerController] from an unsafe.Pointer.
//
// Provides a standard user interface to allow one iOS device to discover and connect to another.
func PeerPickerControllerFrom(ptr unsafe.Pointer) PeerPickerController {
	return PeerPickerController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PeerPickerControllerClass) Alloc() PeerPickerController {
	rv := objc.Send[PeerPickerController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PeerPickerControllerClass) New() PeerPickerController {
	rv := objc.Send[PeerPickerController](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PeerPickerController) Init() PeerPickerController {
	rv := objc.Send[PeerPickerController](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PeerPickerController) Autorelease() PeerPickerController {
	rv := objc.Send[PeerPickerController](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPeerPickerController creates a new PeerPickerController instance.
func NewPeerPickerController() PeerPickerController {
	return getPeerPickerControllerClass().New()
}


// Hides the peer picker dialog.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerController/dismiss()
func (p_ PeerPickerController) Dismiss() {
	objc.Send[objc.ID](p_.ID, objc.Sel("dismiss"))
}

// Displays the peer picker dialog to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerController/show()
func (p_ PeerPickerController) Show() {
	objc.Send[objc.ID](p_.ID, objc.Sel("show"))
}

// A mask that determines the types of connections a dialog presents to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerController/connectionTypesMask
func (p_ PeerPickerController) ConnectionTypesMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("connectionTypesMask"))
	return rv
}


// SetConnectionTypesMask sets the value of the connectionTypesMask property.
// A mask that determines the types of connections a dialog presents to the user.

//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerController/connectionTypesMask
func (p_ PeerPickerController) SetConnectionTypesMask(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setConnectionTypesMask:"), value)
}

// The delegate of the peer picker controller.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerController/delegate
func (p_ PeerPickerController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate of the peer picker controller.

//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerController/delegate
func (p_ PeerPickerController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the picker dialog is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerController/isVisible
func (p_ PeerPickerController) Visible() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("visible"))
	return rv
}

// A Boolean value that indicates whether the picker dialog is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkpeerpickercontroller/isvisible
func (p_ PeerPickerController) IsVisible() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isVisible"))
	return rv
}


// SetIsVisible sets the value of the isVisible property.
// A Boolean value that indicates whether the picker dialog is visible.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkpeerpickercontroller/isvisible
func (p_ PeerPickerController) SetIsVisible(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsVisible:"), value)
}



