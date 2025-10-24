// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKPeerPickerController */


/* debug [class_header]: Header for GKPeerPickerController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PeerPickerController */
// An interface definition for the [PeerPickerController] class.
type IPeerPickerController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PeerPickerController */
	// properties:
	IsVisible() bool
	SetIsVisible(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PeerPickerController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PeerPickerController */
// Alloc allocates a new instance without initialization.
func (pc _PeerPickerControllerClass) Alloc() PeerPickerController {
	rv := objc.Send[PeerPickerController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PeerPickerController */
// Provides a standard user interface to allow one iOS device to discover and connect to another.
//
// The result is a configured object connecting the two devices. To use a object, your application creates the controller, adds a delegate, configures the allowed connection types, and then shows the peer picker. The delegate is called as the user makes selections within the peer picker interface. In iOS 3.0, the peer picker can be configured to select between Bluetooth and Internet connections. On iOS 3.0, your application should release the peer picker object after it dismisses the peer picker dialog. On iOS 3.1 or later, your application may release the peer picker after it is shown to the user. If you do this, the peer picker controller is automatically deallocated after the dialog is dismissed.


// Provides a standard user interface to allow one iOS device to discover and connect to another.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PeerPickerController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PeerPickerController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PeerPickerController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PeerPickerController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PeerPickerController */

// A Boolean value that indicates whether the picker dialog is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkpeerpickercontroller/isvisible
func (p_ PeerPickerController) IsVisible() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isVisible"))
	return rv
}/* debug [instance_properties/getter]: isVisible */


// A Boolean value that indicates whether the picker dialog is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkpeerpickercontroller/isvisible
func (p_ PeerPickerController) SetIsVisible(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsVisible:"), value)
}/* debug [instance_properties/setter]: isVisible */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKPeerPickerController */


