// Code generated from Apple documentation for FileProviderUI. DO NOT EDIT.

package fileproviderui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	foundation.IExtensionContext
	// properties:
	DomainIdentifier() FileProviderDomainIdentifier /* not a class type */
	SetDomainIdentifier(value FileProviderDomainIdentifier /* not a class type */)
	// methods:
	CancelRequestWithError(error_ Error /* not a class type */)
	CompleteRequest()
}

// An extension context provided to File Provider UI extensions.


// An extension context provided to File Provider UI extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionContext
type FPUIActionExtensionContext struct {
	foundation.ExtensionContext
}

// FPUIActionExtensionContextFrom constructs a [FPUIActionExtensionContext] from an unsafe.Pointer.
//
// An extension context provided to File Provider UI extensions.
func FPUIActionExtensionContextFrom(ptr unsafe.Pointer) FPUIActionExtensionContext {
	return FPUIActionExtensionContext{
		ExtensionContext: foundation.ExtensionContextFrom(ptr),
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionContext/cancelRequest(withError:)
func (f_ FPUIActionExtensionContext) CancelRequestWithError(error_ Error /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("cancelRequestWithError:"), error_)
}


// Marks the action as complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionContext/completeRequest()
func (f_ FPUIActionExtensionContext) CompleteRequest() {
	objc.Send[objc.ID](f_.ID, objc.Sel("completeRequest"))
}


// The identifier for the domain managed by the current file provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileproviderui/fpuiactionextensioncontext/domainidentifier
func (f_ FPUIActionExtensionContext) DomainIdentifier() FileProviderDomainIdentifier /* not a class type */ {
	rv := objc.Send[FileProviderDomainIdentifier](f_.ID, objc.Sel("domainIdentifier"))
	return rv
}


// The identifier for the domain managed by the current file provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileproviderui/fpuiactionextensioncontext/domainidentifier
func (f_ FPUIActionExtensionContext) SetDomainIdentifier(value FileProviderDomainIdentifier /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDomainIdentifier:"), value)
}



