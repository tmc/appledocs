// Code generated from Apple documentation for IdentityLookupUI. DO NOT EDIT.

package identitylookupui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [ILClassificationUIExtensionViewController] class.
var (
	ILClassificationUIExtensionViewControllerClass     _ILClassificationUIExtensionViewControllerClass
	ILClassificationUIExtensionViewControllerClassOnce sync.Once
)

func getILClassificationUIExtensionViewControllerClass() _ILClassificationUIExtensionViewControllerClass {
	ILClassificationUIExtensionViewControllerClassOnce.Do(func() {
		ILClassificationUIExtensionViewControllerClass = _ILClassificationUIExtensionViewControllerClass{objc.GetClass("ILClassificationUIExtensionViewController")}
	})
	return ILClassificationUIExtensionViewControllerClass
}

type _ILClassificationUIExtensionViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [ILClassificationUIExtensionViewController] class.
type IILClassificationUIExtensionViewController interface {
	appkit.IViewController
	ClassificationResponseForRequest(request unsafe.Pointer) unsafe.Pointer
	PrepareForClassificationRequest(request unsafe.Pointer)
	ExtensionContext() ILClassificationUIExtensionContext
}

// The superclass for an Unwanted Communication Reporting extension’s principal view controller.
//
// Subclass this view controller to create a user interface that gathers additional information from the user about the reported communication.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookupUI/ILClassificationUIExtensionViewController
type ILClassificationUIExtensionViewController struct {
	appkit.ViewController
}

// ILClassificationUIExtensionViewControllerFrom constructs a [ILClassificationUIExtensionViewController] from an unsafe.Pointer.
//
// The superclass for an Unwanted Communication Reporting extension’s principal view controller.
func ILClassificationUIExtensionViewControllerFrom(ptr unsafe.Pointer) ILClassificationUIExtensionViewController {
	return ILClassificationUIExtensionViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ILClassificationUIExtensionViewControllerClass) Alloc() ILClassificationUIExtensionViewController {
	rv := objc.Send[ILClassificationUIExtensionViewController](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ILClassificationUIExtensionViewControllerClass) New() ILClassificationUIExtensionViewController {
	rv := objc.Send[ILClassificationUIExtensionViewController](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ILClassificationUIExtensionViewController) Init() ILClassificationUIExtensionViewController {
	rv := objc.Send[ILClassificationUIExtensionViewController](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ILClassificationUIExtensionViewController) Autorelease() ILClassificationUIExtensionViewController {
	rv := objc.Send[ILClassificationUIExtensionViewController](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewILClassificationUIExtensionViewController creates a new ILClassificationUIExtensionViewController instance.
func NewILClassificationUIExtensionViewController() ILClassificationUIExtensionViewController {
	return getILClassificationUIExtensionViewControllerClass().New()
}


// Notifies the view controller when the user finishes entering data and presses the Done button.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookupUI/ILClassificationUIExtensionViewController/classificationResponse(for:)
func (i_ ILClassificationUIExtensionViewController) ClassificationResponseForRequest(request unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("classificationResponseForRequest:"), request)
	return rv
}

// Notifies the view controller just before the system presents it to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookupUI/ILClassificationUIExtensionViewController/prepare(for:)
func (i_ ILClassificationUIExtensionViewController) PrepareForClassificationRequest(request unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("prepareForClassificationRequest:"), request)
}

// The context for the current request.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookupUI/ILClassificationUIExtensionViewController/extensionContext
func (i_ ILClassificationUIExtensionViewController) ExtensionContext() ILClassificationUIExtensionContext {
	rv := objc.Send[ILClassificationUIExtensionContext](i_.ID, objc.Sel("extensionContext"))
	return rv
}




