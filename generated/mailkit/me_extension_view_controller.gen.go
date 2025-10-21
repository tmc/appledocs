// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [MEExtensionViewController] class.
var (
	MEExtensionViewControllerClass     _MEExtensionViewControllerClass
	MEExtensionViewControllerClassOnce sync.Once
)

func getMEExtensionViewControllerClass() _MEExtensionViewControllerClass {
	MEExtensionViewControllerClassOnce.Do(func() {
		MEExtensionViewControllerClass = _MEExtensionViewControllerClass{objc.GetClass("MEExtensionViewController")}
	})
	return MEExtensionViewControllerClass
}

type _MEExtensionViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [MEExtensionViewController] class.
type IMEExtensionViewController interface {
	appkit.IViewController
}

// An object that manages a view for compose session and message security handlers.
//
// Create a custom subclass of to provide MailKit with a view that displays: Additional configuration options in a Mail compose window Details about the user who signed a digitally signed email message
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEExtensionViewController
type MEExtensionViewController struct {
	appkit.ViewController
}

// MEExtensionViewControllerFrom constructs a [MEExtensionViewController] from an unsafe.Pointer.
//
// An object that manages a view for compose session and message security handlers.
func MEExtensionViewControllerFrom(ptr unsafe.Pointer) MEExtensionViewController {
	return MEExtensionViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MEExtensionViewControllerClass) Alloc() MEExtensionViewController {
	rv := objc.Send[MEExtensionViewController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEExtensionViewControllerClass) New() MEExtensionViewController {
	rv := objc.Send[MEExtensionViewController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEExtensionViewController) Init() MEExtensionViewController {
	rv := objc.Send[MEExtensionViewController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEExtensionViewController) Autorelease() MEExtensionViewController {
	rv := objc.Send[MEExtensionViewController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEExtensionViewController creates a new MEExtensionViewController instance.
func NewMEExtensionViewController() MEExtensionViewController {
	return getMEExtensionViewControllerClass().New()
}




