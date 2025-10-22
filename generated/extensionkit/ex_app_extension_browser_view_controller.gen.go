// Code generated from Apple documentation for ExtensionKit. DO NOT EDIT.

package extensionkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [EXAppExtensionBrowserViewController] class.
var (
	EXAppExtensionBrowserViewControllerClass     _EXAppExtensionBrowserViewControllerClass
	EXAppExtensionBrowserViewControllerClassOnce sync.Once
)

func getEXAppExtensionBrowserViewControllerClass() _EXAppExtensionBrowserViewControllerClass {
	EXAppExtensionBrowserViewControllerClassOnce.Do(func() {
		EXAppExtensionBrowserViewControllerClass = _EXAppExtensionBrowserViewControllerClass{objc.GetClass("EXAppExtensionBrowserViewController")}
	})
	return EXAppExtensionBrowserViewControllerClass
}

type _EXAppExtensionBrowserViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [EXAppExtensionBrowserViewController] class.
type IEXAppExtensionBrowserViewController interface {
	appkit.IViewController
}

// A view controller that displays an interface to enable or disable the host app’s extensions.
//
// When your host app supports app extensions, use this view controller to give people a way to enable or disable those extensions. When you present this view controller, the system displays an out-of-process UI with a list of all app extensions that support your app’s extension points. Someone using your app can use the presented interface to enable or disable extensions selectively. App extensions you include inside your host app’s bundle are enabled by default, but extensions that ship in separate apps are disabled by default. Present this view controller modally from your app, or embed the view controller as a child in one of your existing view controller interfaces. For example, you might choose to embed the view controller in a tab of your app’s preferences interface.


// A view controller that displays an interface to enable or disable the host app’s extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExtensionKit/EXAppExtensionBrowserViewController

type EXAppExtensionBrowserViewController struct {
	appkit.ViewController
}

// EXAppExtensionBrowserViewControllerFrom constructs a [EXAppExtensionBrowserViewController] from an unsafe.Pointer.
//
// A view controller that displays an interface to enable or disable the host app’s extensions.
func EXAppExtensionBrowserViewControllerFrom(ptr unsafe.Pointer) EXAppExtensionBrowserViewController {
	return EXAppExtensionBrowserViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _EXAppExtensionBrowserViewControllerClass) Alloc() EXAppExtensionBrowserViewController {
	rv := objc.Send[EXAppExtensionBrowserViewController](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EXAppExtensionBrowserViewControllerClass) New() EXAppExtensionBrowserViewController {
	rv := objc.Send[EXAppExtensionBrowserViewController](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EXAppExtensionBrowserViewController) Init() EXAppExtensionBrowserViewController {
	rv := objc.Send[EXAppExtensionBrowserViewController](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EXAppExtensionBrowserViewController) Autorelease() EXAppExtensionBrowserViewController {
	rv := objc.Send[EXAppExtensionBrowserViewController](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEXAppExtensionBrowserViewController creates a new EXAppExtensionBrowserViewController instance.
func NewEXAppExtensionBrowserViewController() EXAppExtensionBrowserViewController {
	return getEXAppExtensionBrowserViewControllerClass().New()
}




