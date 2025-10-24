// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CXCallDirectoryExtensionContext] class.
var (
	CXCallDirectoryExtensionContextClass     _CXCallDirectoryExtensionContextClass
	CXCallDirectoryExtensionContextClassOnce sync.Once
)

func getCXCallDirectoryExtensionContextClass() _CXCallDirectoryExtensionContextClass {
	CXCallDirectoryExtensionContextClassOnce.Do(func() {
		CXCallDirectoryExtensionContextClass = _CXCallDirectoryExtensionContextClass{objc.GetClass("CXCallDirectoryExtensionContext")}
	})
	return CXCallDirectoryExtensionContextClass
}

type _CXCallDirectoryExtensionContextClass struct {
	class objc.Class
}

// An interface definition for the [CXCallDirectoryExtensionContext] class.
type ICXCallDirectoryExtensionContext interface {
	foundation.IExtensionContext
	// properties:
	IsIncremental() bool
	SetIsIncremental(value bool)
	CXCallDirectoryPhoneNumberMax() CXCallDirectoryPhoneNumber /* typedef */
	// methods:
}

// A programmatic interface for adding identification and blocking entries to a Call Directory app extension.
//
// The system doesn’t initialize objects directly, but instead passes them as arguments to the instance method .


// A programmatic interface for adding identification and blocking entries to a Call Directory app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext
type CXCallDirectoryExtensionContext struct {
	foundation.ExtensionContext
}

// CXCallDirectoryExtensionContextFrom constructs a [CXCallDirectoryExtensionContext] from an unsafe.Pointer.
//
// A programmatic interface for adding identification and blocking entries to a Call Directory app extension.
func CXCallDirectoryExtensionContextFrom(ptr unsafe.Pointer) CXCallDirectoryExtensionContext {
	return CXCallDirectoryExtensionContext{
		ExtensionContext: foundation.ExtensionContextFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CXCallDirectoryExtensionContextClass) Alloc() CXCallDirectoryExtensionContext {
	rv := objc.Send[CXCallDirectoryExtensionContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXCallDirectoryExtensionContextClass) New() CXCallDirectoryExtensionContext {
	rv := objc.Send[CXCallDirectoryExtensionContext](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallDirectoryExtensionContext) Init() CXCallDirectoryExtensionContext {
	rv := objc.Send[CXCallDirectoryExtensionContext](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallDirectoryExtensionContext) Autorelease() CXCallDirectoryExtensionContext {
	rv := objc.Send[CXCallDirectoryExtensionContext](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallDirectoryExtensionContext creates a new CXCallDirectoryExtensionContext instance.
func NewCXCallDirectoryExtensionContext() CXCallDirectoryExtensionContext {
	return getCXCallDirectoryExtensionContextClass().New()
}



// A Boolean value that indicates whether the request provides data incrementally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcalldirectoryextensioncontext/isincremental
func (c_ CXCallDirectoryExtensionContext) IsIncremental() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isIncremental"))
	return rv
}


// A Boolean value that indicates whether the request provides data incrementally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcalldirectoryextensioncontext/isincremental
func (c_ CXCallDirectoryExtensionContext) SetIsIncremental(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsIncremental:"), value)
}


// The maximum allowable value for a phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcalldirectoryphonenumbermax
func (c_ CXCallDirectoryExtensionContext) CXCallDirectoryPhoneNumberMax() CXCallDirectoryPhoneNumber /* typedef */ {
	rv := objc.Send[CXCallDirectoryPhoneNumber](c_.ID, objc.Sel("CXCallDirectoryPhoneNumberMax"))
	return rv
}


