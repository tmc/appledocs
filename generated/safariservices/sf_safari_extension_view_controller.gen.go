// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [SFSafariExtensionViewController] class.
var (
	SFSafariExtensionViewControllerClass     _SFSafariExtensionViewControllerClass
	SFSafariExtensionViewControllerClassOnce sync.Once
)

func getSFSafariExtensionViewControllerClass() _SFSafariExtensionViewControllerClass {
	SFSafariExtensionViewControllerClassOnce.Do(func() {
		SFSafariExtensionViewControllerClass = _SFSafariExtensionViewControllerClass{objc.GetClass("SFSafariExtensionViewController")}
	})
	return SFSafariExtensionViewControllerClass
}

type _SFSafariExtensionViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [SFSafariExtensionViewController] class.
type ISFSafariExtensionViewController interface {
	appkit.IViewController
	// properties:
	// methods:
	DismissPopover()
}

// The view controller for a popover associated with your app extension.
//
// If your toolbar item has a popover, your popover view controller should be a subclass of this class. As with other macOS development, typically you want to add your own outlets and actions to the view controller, and provide an XIB file for its user interface. Your view controller’s contents must use Auto Layout.


// The view controller for a popover associated with your app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionViewController
type SFSafariExtensionViewController struct {
	appkit.ViewController
}

// SFSafariExtensionViewControllerFrom constructs a [SFSafariExtensionViewController] from an unsafe.Pointer.
//
// The view controller for a popover associated with your app extension.
func SFSafariExtensionViewControllerFrom(ptr unsafe.Pointer) SFSafariExtensionViewController {
	return SFSafariExtensionViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariExtensionViewControllerClass) Alloc() SFSafariExtensionViewController {
	rv := objc.Send[SFSafariExtensionViewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSafariExtensionViewControllerClass) New() SFSafariExtensionViewController {
	rv := objc.Send[SFSafariExtensionViewController](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariExtensionViewController) Init() SFSafariExtensionViewController {
	rv := objc.Send[SFSafariExtensionViewController](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariExtensionViewController) Autorelease() SFSafariExtensionViewController {
	rv := objc.Send[SFSafariExtensionViewController](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariExtensionViewController creates a new SFSafariExtensionViewController instance.
func NewSFSafariExtensionViewController() SFSafariExtensionViewController {
	return getSFSafariExtensionViewControllerClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionViewController/dismissPopover()
func (s_ SFSafariExtensionViewController) DismissPopover() {
	objc.Send[objc.ID](s_.ID, objc.Sel("dismissPopover"))
}



