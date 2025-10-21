// Code generated from Apple documentation for IdentityLookupUI. DO NOT EDIT.

package identitylookupui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ILClassificationUIExtensionContext] class.
var (
	ILClassificationUIExtensionContextClass     _ILClassificationUIExtensionContextClass
	ILClassificationUIExtensionContextClassOnce sync.Once
)

func getILClassificationUIExtensionContextClass() _ILClassificationUIExtensionContextClass {
	ILClassificationUIExtensionContextClassOnce.Do(func() {
		ILClassificationUIExtensionContextClass = _ILClassificationUIExtensionContextClass{objc.GetClass("ILClassificationUIExtensionContext")}
	})
	return ILClassificationUIExtensionContextClass
}

type _ILClassificationUIExtensionContextClass struct {
	class objc.Class
}

// An interface definition for the [ILClassificationUIExtensionContext] class.
type IILClassificationUIExtensionContext interface {
	foundation.IExtensionContext
}

// An object that manages the state of the current request.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookupUI/ILClassificationUIExtensionContext
type ILClassificationUIExtensionContext struct {
	foundation.ExtensionContext
}

// ILClassificationUIExtensionContextFrom constructs a [ILClassificationUIExtensionContext] from an unsafe.Pointer.
//
// An object that manages the state of the current request.
func ILClassificationUIExtensionContextFrom(ptr unsafe.Pointer) ILClassificationUIExtensionContext {
	return ILClassificationUIExtensionContext{
		ExtensionContext: foundation.ExtensionContextFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ILClassificationUIExtensionContextClass) Alloc() ILClassificationUIExtensionContext {
	rv := objc.Send[ILClassificationUIExtensionContext](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ILClassificationUIExtensionContextClass) New() ILClassificationUIExtensionContext {
	rv := objc.Send[ILClassificationUIExtensionContext](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ILClassificationUIExtensionContext) Init() ILClassificationUIExtensionContext {
	rv := objc.Send[ILClassificationUIExtensionContext](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ILClassificationUIExtensionContext) Autorelease() ILClassificationUIExtensionContext {
	rv := objc.Send[ILClassificationUIExtensionContext](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewILClassificationUIExtensionContext creates a new ILClassificationUIExtensionContext instance.
func NewILClassificationUIExtensionContext() ILClassificationUIExtensionContext {
	return getILClassificationUIExtensionContextClass().New()
}


// A Boolean value that determines whether the extension has enough information to complete the report.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookupUI/ILClassificationUIExtensionContext/isReadyForClassificationResponse
func (i_ ILClassificationUIExtensionContext) ReadyForClassificationResponse() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("readyForClassificationResponse"))
	return rv
}


// SetReadyForClassificationResponse sets the value of the readyForClassificationResponse property.
// A Boolean value that determines whether the extension has enough information to complete the report.

//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookupUI/ILClassificationUIExtensionContext/isReadyForClassificationResponse
func (i_ ILClassificationUIExtensionContext) SetReadyForClassificationResponse(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReadyForClassificationResponse:"), value)
}


