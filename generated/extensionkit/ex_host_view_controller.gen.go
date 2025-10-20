// Code generated from Apple documentation for ExtensionKit. DO NOT EDIT.

package extensionkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [EXHostViewController] class.
var (
	EXHostViewControllerClass     _EXHostViewControllerClass
	EXHostViewControllerClassOnce sync.Once
)

func getEXHostViewControllerClass() _EXHostViewControllerClass {
	EXHostViewControllerClassOnce.Do(func() {
		EXHostViewControllerClass = _EXHostViewControllerClass{objc.GetClass("EXHostViewController")}
	})
	return EXHostViewControllerClass
}

type _EXHostViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [EXHostViewController] class.
type IEXHostViewController interface {
	appkit.IViewController
	MakeXPCConnection()
}

// A view controller that hosts remote views provided by an app extension.
//
// Present this view controller from your app’s interface to display the content for an associated app extension. Configure the view controller with the app extension identity and the specific scene you want to display. Use the associated delegate object to receive notifications when the app extension becomes active or inactive. For more information about presenting this view controller and using it to display an app extension’s UI, see .
//
// [Full Topic]: https://developer.apple.com/documentation/ExtensionKit/EXHostViewController
type EXHostViewController struct {
	appkit.ViewController
}

// EXHostViewControllerFrom constructs a [EXHostViewController] from an unsafe.Pointer.
//
// A view controller that hosts remote views provided by an app extension.
func EXHostViewControllerFrom(ptr unsafe.Pointer) EXHostViewController {
	return EXHostViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _EXHostViewControllerClass) Alloc() EXHostViewController {
	rv := objc.Send[EXHostViewController](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EXHostViewControllerClass) New() EXHostViewController {
	rv := objc.Send[EXHostViewController](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EXHostViewController) Init() EXHostViewController {
	rv := objc.Send[EXHostViewController](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EXHostViewController) Autorelease() EXHostViewController {
	rv := objc.Send[EXHostViewController](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEXHostViewController creates a new EXHostViewController instance.
func NewEXHostViewController() EXHostViewController {
	return getEXHostViewControllerClass().New()
}


// Initiates an XPC connection to the app extension’s scene.
//
// [Full Topic]: https://developer.apple.com/documentation/ExtensionKit/EXHostViewController/makeXPCConnection()
func (e_ EXHostViewController) MakeXPCConnection() {
	objc.Send[objc.ID](e_.ID, objc.Sel("makeXPCConnection"))
}



