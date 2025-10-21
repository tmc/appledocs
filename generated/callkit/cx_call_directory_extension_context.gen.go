// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	IExtensionContext
	AddBlockingEntryWithNextSequentialPhoneNumber(phoneNumber ICXCallDirectoryPhoneNumber)
	AddIdentificationEntryWithNextSequentialPhoneNumberLabel(phoneNumber ICXCallDirectoryPhoneNumber, label string)
	CompleteRequestWithCompletionHandler(completion unsafe.Pointer)
	RemoveAllBlockingEntries()
	RemoveAllIdentificationEntries()
	RemoveBlockingEntryWithPhoneNumber(phoneNumber ICXCallDirectoryPhoneNumber)
	RemoveIdentificationEntryWithPhoneNumber(phoneNumber ICXCallDirectoryPhoneNumber)
}

// A programmatic interface for adding identification and blocking entries to a Call Directory app extension.
//
// The system doesn’t initialize objects directly, but instead passes them as arguments to the instance method .
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext
type CXCallDirectoryExtensionContext struct {
	ExtensionContext
}

// CXCallDirectoryExtensionContextFrom constructs a [CXCallDirectoryExtensionContext] from an unsafe.Pointer.
//
// A programmatic interface for adding identification and blocking entries to a Call Directory app extension.
func CXCallDirectoryExtensionContextFrom(ptr unsafe.Pointer) CXCallDirectoryExtensionContext {
	return CXCallDirectoryExtensionContext{
		ExtensionContext: ExtensionContextFrom(ptr),
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


// Adds a blocking entry with the specified phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/addBlockingEntry(withNextSequentialPhoneNumber:)
func (c_ CXCallDirectoryExtensionContext) AddBlockingEntryWithNextSequentialPhoneNumber(phoneNumber ICXCallDirectoryPhoneNumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addBlockingEntryWithNextSequentialPhoneNumber:"), phoneNumber)
}

// Adds an identification entry with the specified phone number and label.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/addIdentificationEntry(withNextSequentialPhoneNumber:label:)
func (c_ CXCallDirectoryExtensionContext) AddIdentificationEntryWithNextSequentialPhoneNumberLabel(phoneNumber ICXCallDirectoryPhoneNumber, label string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addIdentificationEntryWithNextSequentialPhoneNumber:label:"), phoneNumber, objc.String(label))
}

// Completes the request to the extension context.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/completeRequest(completionHandler:)
func (c_ CXCallDirectoryExtensionContext) CompleteRequestWithCompletionHandler(completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("completeRequestWithCompletionHandler:"), completion)
}

// Removes all stored blocking entries.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/removeAllBlockingEntries()
func (c_ CXCallDirectoryExtensionContext) RemoveAllBlockingEntries() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllBlockingEntries"))
}

// Removes all stored identification entries.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/removeAllIdentificationEntries()
func (c_ CXCallDirectoryExtensionContext) RemoveAllIdentificationEntries() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllIdentificationEntries"))
}

// Removes a blocking entry that contains the specified phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/removeBlockingEntry(withPhoneNumber:)
func (c_ CXCallDirectoryExtensionContext) RemoveBlockingEntryWithPhoneNumber(phoneNumber ICXCallDirectoryPhoneNumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeBlockingEntryWithPhoneNumber:"), phoneNumber)
}

// Removes an identification entry that contains the specified phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/removeIdentificationEntry(withPhoneNumber:)
func (c_ CXCallDirectoryExtensionContext) RemoveIdentificationEntryWithPhoneNumber(phoneNumber ICXCallDirectoryPhoneNumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeIdentificationEntryWithPhoneNumber:"), phoneNumber)
}

// Sets a delegate that can handle request failures for the Call Directory extension context object.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/delegate
func (c_ CXCallDirectoryExtensionContext) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// Sets a delegate that can handle request failures for the Call Directory extension context object.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/delegate
func (c_ CXCallDirectoryExtensionContext) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the request provides data incrementally.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/isIncremental
func (c_ CXCallDirectoryExtensionContext) Incremental() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("incremental"))
	return rv
}

// A Boolean value that indicates whether the request provides data incrementally.
//
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcalldirectoryextensioncontext/isincremental
func (c_ CXCallDirectoryExtensionContext) IsIncremental() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isIncremental"))
	return rv
}


// SetIsIncremental sets the value of the isIncremental property.
// A Boolean value that indicates whether the request provides data incrementally.

//
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcalldirectoryextensioncontext/isincremental
func (c_ CXCallDirectoryExtensionContext) SetIsIncremental(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsIncremental:"), value)
}

// The maximum allowable value for a phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcalldirectoryphonenumbermax
func (c_ CXCallDirectoryExtensionContext) CXCallDirectoryPhoneNumberMax() CXCallDirectoryPhoneNumber {
	rv := objc.Send[CXCallDirectoryPhoneNumber](c_.ID, objc.Sel("CXCallDirectoryPhoneNumberMax"))
	return rv
}



