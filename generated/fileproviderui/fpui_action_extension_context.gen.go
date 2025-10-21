// Code generated from Apple documentation for FileProviderUI. DO NOT EDIT.

package fileproviderui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FPUIActionExtensionContext] class.
var (
	FPUIActionExtensionContextClass     _FPUIActionExtensionContextClass
	FPUIActionExtensionContextClassOnce sync.Once
)

func getFPUIActionExtensionContextClass() _FPUIActionExtensionContextClass {
	FPUIActionExtensionContextClassOnce.Do(func() {
		FPUIActionExtensionContextClass = _FPUIActionExtensionContextClass{objc.GetClass("FPUIActionExtensionContext")}
	})
	return FPUIActionExtensionContextClass
}

type _FPUIActionExtensionContextClass struct {
	class objc.Class
}

// An interface definition for the [FPUIActionExtensionContext] class.
type IFPUIActionExtensionContext interface {
	IExtensionContext
	CancelRequestWithError(error_ unsafe.Pointer)
	CompleteRequest()
}

// An extension context provided to File Provider UI extensions.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionContext
type FPUIActionExtensionContext struct {
	ExtensionContext
}

// FPUIActionExtensionContextFrom constructs a [FPUIActionExtensionContext] from an unsafe.Pointer.
//
// An extension context provided to File Provider UI extensions.
func FPUIActionExtensionContextFrom(ptr unsafe.Pointer) FPUIActionExtensionContext {
	return FPUIActionExtensionContext{
		ExtensionContext: ExtensionContextFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FPUIActionExtensionContextClass) Alloc() FPUIActionExtensionContext {
	rv := objc.Send[FPUIActionExtensionContext](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FPUIActionExtensionContextClass) New() FPUIActionExtensionContext {
	rv := objc.Send[FPUIActionExtensionContext](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FPUIActionExtensionContext) Init() FPUIActionExtensionContext {
	rv := objc.Send[FPUIActionExtensionContext](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FPUIActionExtensionContext) Autorelease() FPUIActionExtensionContext {
	rv := objc.Send[FPUIActionExtensionContext](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFPUIActionExtensionContext creates a new FPUIActionExtensionContext instance.
func NewFPUIActionExtensionContext() FPUIActionExtensionContext {
	return getFPUIActionExtensionContextClass().New()
}


// Cancels the action and returns the provided error.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionContext/cancelRequest(withError:)
func (f_ FPUIActionExtensionContext) CancelRequestWithError(error_ unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("cancelRequestWithError:"), error_)
}

// Marks the action as complete.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionContext/completeRequest()
func (f_ FPUIActionExtensionContext) CompleteRequest() {
	objc.Send[objc.ID](f_.ID, objc.Sel("completeRequest"))
}

// The identifier for the domain managed by the current file provider.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionContext/domainIdentifier
func (f_ FPUIActionExtensionContext) DomainIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("domainIdentifier"))
	return rv
}

// The extension context provided by the host app.
//
// [Full Topic]: https://developer.apple.com/documentation/fileproviderui/fpuiactionextensionviewcontroller/extensioncontext
func (f_ FPUIActionExtensionContext) ExtensionContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("extensionContext"))
	return rv
}


// SetExtensionContext sets the value of the extensionContext property.
// The extension context provided by the host app.

//
// [Full Topic]: https://developer.apple.com/documentation/fileproviderui/fpuiactionextensionviewcontroller/extensioncontext
func (f_ FPUIActionExtensionContext) SetExtensionContext(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setExtensionContext:"), value)
}



