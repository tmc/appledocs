// Code generated from Apple documentation for FileProviderUI. DO NOT EDIT.

package fileproviderui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [FPUIActionExtensionViewController] class.
var (
	FPUIActionExtensionViewControllerClass     _FPUIActionExtensionViewControllerClass
	FPUIActionExtensionViewControllerClassOnce sync.Once
)

func getFPUIActionExtensionViewControllerClass() _FPUIActionExtensionViewControllerClass {
	FPUIActionExtensionViewControllerClassOnce.Do(func() {
		FPUIActionExtensionViewControllerClass = _FPUIActionExtensionViewControllerClass{objc.GetClass("FPUIActionExtensionViewController")}
	})
	return FPUIActionExtensionViewControllerClass
}

type _FPUIActionExtensionViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [FPUIActionExtensionViewController] class.
type IFPUIActionExtensionViewController interface {
	appkit.IViewController
	PrepareForActionWithIdentifierItemIdentifiers(actionIdentifier string, itemIdentifiers unsafe.Pointer)
	PrepareForError(error_ unsafe.Pointer)
}

// The custom user interface used to perform a selected action.
//
// Subclass this view controller to provide the user interface for your actions. No matter how many actions you define, your File Provider UI extension has only one subclass. When the user selects one of your actions, the system instantiates a copy of your subclass, calls its method, and presents it to the user. Your subclass must do the following: Override the method to check the action identifiers and present an appropriate user interface for the selected actions. Provide some sort of feedback, even if the action doesn’t require interaction with the user. For example, present a view that quickly fades out and automatically completes the action. Call the object’s or method when the action is finished to complete the action.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionViewController
type FPUIActionExtensionViewController struct {
	appkit.ViewController
}

// FPUIActionExtensionViewControllerFrom constructs a [FPUIActionExtensionViewController] from an unsafe.Pointer.
//
// The custom user interface used to perform a selected action.
func FPUIActionExtensionViewControllerFrom(ptr unsafe.Pointer) FPUIActionExtensionViewController {
	return FPUIActionExtensionViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FPUIActionExtensionViewControllerClass) Alloc() FPUIActionExtensionViewController {
	rv := objc.Send[FPUIActionExtensionViewController](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FPUIActionExtensionViewControllerClass) New() FPUIActionExtensionViewController {
	rv := objc.Send[FPUIActionExtensionViewController](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FPUIActionExtensionViewController) Init() FPUIActionExtensionViewController {
	rv := objc.Send[FPUIActionExtensionViewController](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FPUIActionExtensionViewController) Autorelease() FPUIActionExtensionViewController {
	rv := objc.Send[FPUIActionExtensionViewController](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFPUIActionExtensionViewController creates a new FPUIActionExtensionViewController instance.
func NewFPUIActionExtensionViewController() FPUIActionExtensionViewController {
	return getFPUIActionExtensionViewControllerClass().New()
}


// Performs any necessary setup or configuration for the specified action.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionViewController/prepare(forAction:itemIdentifiers:)
func (f_ FPUIActionExtensionViewController) PrepareForActionWithIdentifierItemIdentifiers(actionIdentifier string, itemIdentifiers unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("prepareForActionWithIdentifier:itemIdentifiers:"), objc.String(actionIdentifier), itemIdentifiers)
}

// Performs any necessary setup or configuration when an authentication error occurs.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionViewController/prepare(forError:)
func (f_ FPUIActionExtensionViewController) PrepareForError(error_ unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("prepareForError:"), error_)
}

// The extension context provided by the host app.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionViewController/extensionContext
func (f_ FPUIActionExtensionViewController) ExtensionContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("extensionContext"))
	return rv
}


