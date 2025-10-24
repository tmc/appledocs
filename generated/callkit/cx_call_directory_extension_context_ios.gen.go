//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CXCallDirectoryExtensionContext


// Adds a blocking entry with the specified phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/addBlockingEntry(withNextSequentialPhoneNumber:)
func (c_ CXCallDirectoryExtensionContext) AddBlockingEntryWithNextSequentialPhoneNumber(phoneNumber CXCallDirectoryPhoneNumber /* typedef */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addBlockingEntryWithNextSequentialPhoneNumber:"), phoneNumber)
}

// Adds an identification entry with the specified phone number and label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/addIdentificationEntry(withNextSequentialPhoneNumber:label:)
func (c_ CXCallDirectoryExtensionContext) AddIdentificationEntryWithNextSequentialPhoneNumberLabel(phoneNumber CXCallDirectoryPhoneNumber /* typedef */, label objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addIdentificationEntryWithNextSequentialPhoneNumber:label:"), phoneNumber, label)
}

// Completes the request to the extension context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/completeRequest(completionHandler:)
func (c_ CXCallDirectoryExtensionContext) CompleteRequestWithCompletionHandler(completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("completeRequestWithCompletionHandler:"), completion)
}

// Removes all stored blocking entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/removeAllBlockingEntries()
func (c_ CXCallDirectoryExtensionContext) RemoveAllBlockingEntries() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllBlockingEntries"))
}

// Removes all stored identification entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/removeAllIdentificationEntries()
func (c_ CXCallDirectoryExtensionContext) RemoveAllIdentificationEntries() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllIdentificationEntries"))
}

// Removes a blocking entry that contains the specified phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/removeBlockingEntry(withPhoneNumber:)
func (c_ CXCallDirectoryExtensionContext) RemoveBlockingEntryWithPhoneNumber(phoneNumber CXCallDirectoryPhoneNumber /* typedef */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeBlockingEntryWithPhoneNumber:"), phoneNumber)
}

// Removes an identification entry that contains the specified phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/removeIdentificationEntry(withPhoneNumber:)
func (c_ CXCallDirectoryExtensionContext) RemoveIdentificationEntryWithPhoneNumber(phoneNumber CXCallDirectoryPhoneNumber /* typedef */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeIdentificationEntryWithPhoneNumber:"), phoneNumber)
}

// iOS-only properties

// Sets a delegate that can handle request failures for the Call Directory extension context object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/delegate
func (c_ CXCallDirectoryExtensionContext) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}
func (c_ CXCallDirectoryExtensionContext) SetDelegate(value objc.ID) {
	c_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// A Boolean value that indicates whether the request provides data incrementally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext/isIncremental
func (c_ CXCallDirectoryExtensionContext) Incremental() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("incremental"))
	return rv
}





