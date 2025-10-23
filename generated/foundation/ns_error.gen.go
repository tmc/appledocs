// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Error] class.
var (
	ErrorClass     _ErrorClass
	ErrorClassOnce sync.Once
)

func getErrorClass() _ErrorClass {
	ErrorClassOnce.Do(func() {
		ErrorClass = _ErrorClass{objc.GetClass("NSError")}
	})
	return ErrorClass
}

type _ErrorClass struct {
	class objc.Class
}

// An interface definition for the [Error] class.
type IError interface {
	objectivec.IObject
	// properties:
	Code() int /* primitive/slice/pointer. */
	Domain() objc.IObject /* cross-framework: ErrorDomain */
	HelpAnchor() IString
	LocalizedDescription() IString
	LocalizedFailureReason() IString
	LocalizedRecoveryOptions() []string /* primitive/slice/pointer. */
	LocalizedRecoverySuggestion() IString
	RecoveryAttempter() objc.ID
	UnderlyingErrors() []Error /* primitive/slice/pointer. */
	UserInfo() IDictionary /* already interface */
	NSCocoaErrorDomain() IString
	NSMachErrorDomain() IString
	NSOSStatusErrorDomain() IString
	NSPOSIXErrorDomain() IString
	NSRecoveryAttempterErrorKey() IString
	NSStreamSOCKSErrorDomain() IString
	NSStreamSocketSSLErrorDomain() IString
	NSURLErrorDomain() IString
	// methods:
}

// Information about an error condition including a domain, a domain-specific error code, and application-specific information.
//
// Objective-C methods can signal an error condition by returning an object by reference, which provides additional information about the kind of error and any underlying cause, if one can be determined. An object may also provide localized error descriptions suitable for display to the user in its user info dictionary. See for more information. Methods in Foundation and other Cocoa frameworks most often produce errors in the Cocoa error domain ( ); error codes for the Cocoa Error Domain are documented in the . There are also predefined domains corresponding to Mach ( ), POSIX ( ), and Carbon ( ) errors. is “toll-free bridged” with its Core Foundation counterpart, . See for more information.


// Information about an error condition including a domain, a domain-specific error code, and application-specific information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError
type Error struct {
	objectivec.Object
}

// ErrorFrom constructs a [Error] from an unsafe.Pointer.
//
// Information about an error condition including a domain, a domain-specific error code, and application-specific information.
func ErrorFrom(ptr unsafe.Pointer) Error {
	return Error{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _ErrorClass) Alloc() Error {
	rv := objc.Send[Error](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _ErrorClass) New() Error {
	rv := objc.Send[Error](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ Error) Init() Error {
	rv := objc.Send[Error](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ Error) Autorelease() Error {
	rv := objc.Send[Error](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewError creates a new Error instance.
func NewError() Error {
	return getErrorClass().New()
}



// Returns an object initialized for a given domain and code with a given dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/init(domain:code:userInfo:)
func NewErrorWithDomainCodeUserInfo(domain objc.IObject /* cross-framework ErrorDomain */, code int /* primitive/slice/pointer. */, dict IDictionary /* already interface */) Error {
	instance := getErrorClass().Alloc()
	rv := objc.Send[Error](instance.ID, objc.Sel("initWithDomain:code:userInfo:"), domain, code, dict)
	rv.Autorelease()
	return rv
}



// Creates and initializes an object for a given domain and code with a given dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/errorWithDomain:code:userInfo:
func (ec _ErrorClass) ErrorWithDomainCodeUserInfo(domain objc.IObject /* cross-framework ErrorDomain */, code int /* primitive/slice/pointer. */, dict IDictionary /* already interface */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("errorWithDomain:code:userInfo:"), domain, code, dict)
	return rv
}


// Returns a properly formatted error object with a error code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/fileProviderErrorForCollision(with:)
func (ec _ErrorClass) FileProviderErrorForCollisionWithItem(existingItem FileProviderItem /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("fileProviderErrorForCollisionWithItem:"), existingItem)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/fileProviderErrorForNonExistentItem(withIdentifier:)
func (ec _ErrorClass) FileProviderErrorForNonExistentItemWithIdentifier(itemIdentifier FileProviderItemIdentifier /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("fileProviderErrorForNonExistentItemWithIdentifier:"), itemIdentifier)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/fileProviderErrorForRejectedDeletion(of:)
func (ec _ErrorClass) FileProviderErrorForRejectedDeletionOfItem(updatedVersion FileProviderItem /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("fileProviderErrorForRejectedDeletionOfItem:"), updatedVersion)
	return rv
}


// Specifies a block to call when the corresponding property is not present in the user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/setUserInfoValueProvider(forDomain:provider:)
func (ec _ErrorClass) SetUserInfoValueProviderForDomainProvider(errorDomain objc.IObject /* cross-framework ErrorDomain */, provider unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("setUserInfoValueProviderForDomain:provider:"), errorDomain, provider)
}


// Returns any user info provider specified for a given error domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfoValueProvider(forDomain:)
func (ec _ErrorClass) UserInfoValueProviderForDomain(errorDomain objc.IObject /* cross-framework ErrorDomain */) {
	objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("userInfoValueProviderForDomain:"), errorDomain)
}


// The error code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/code
func (e_ Error) Code() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("code"))
	return rv
}


// A string containing the error domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/domain
func (e_ Error) Domain() objc.IObject /* cross-framework: ErrorDomain */ {
	rv := objc.Send[ErrorDomain](e_.ID, objc.Sel("domain"))
	return rv
}


// A string to display in response to an alert panel help anchor button being pressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/helpAnchor
func (e_ Error) HelpAnchor() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("helpAnchor"))
	return rv
}


// A string containing the localized description of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/localizedDescription
func (e_ Error) LocalizedDescription() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("localizedDescription"))
	return rv
}


// A string containing the localized explanation of the reason for the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/localizedFailureReason
func (e_ Error) LocalizedFailureReason() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("localizedFailureReason"))
	return rv
}


// An array containing the localized titles of buttons appropriate for displaying in an alert panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/localizedRecoveryOptions
func (e_ Error) LocalizedRecoveryOptions() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](e_.ID, objc.Sel("localizedRecoveryOptions"))
	return rv
}


// A string containing the localized recovery suggestion for the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/localizedRecoverySuggestion
func (e_ Error) LocalizedRecoverySuggestion() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("localizedRecoverySuggestion"))
	return rv
}


// The object in the user info dictionary corresponding to the key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/recoveryAttempter
func (e_ Error) RecoveryAttempter() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("recoveryAttempter"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/underlyingErrors
func (e_ Error) UnderlyingErrors() []Error /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Error](e_.ID, objc.Sel("underlyingErrors"))
	return rv
}


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (e_ Error) UserInfo() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](e_.ID, objc.Sel("userInfo"))
	return rv
}


// Cocoa errors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscocoaerrordomain
func (e_ Error) NSCocoaErrorDomain() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("NSCocoaErrorDomain"))
	return rv
}


// Mach errors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmacherrordomain
func (e_ Error) NSMachErrorDomain() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("NSMachErrorDomain"))
	return rv
}


// Mac OS 9/Carbon errors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsosstatuserrordomain
func (e_ Error) NSOSStatusErrorDomain() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("NSOSStatusErrorDomain"))
	return rv
}


// POSIX/BSD errors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsposixerrordomain
func (e_ Error) NSPOSIXErrorDomain() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("NSPOSIXErrorDomain"))
	return rv
}


// The corresponding value is an object that conforms to the NSErrorRecoveryAttempting informal protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrecoveryattemptererrorkey
func (e_ Error) NSRecoveryAttempterErrorKey() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("NSRecoveryAttempterErrorKey"))
	return rv
}


// The error domain used by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsockserrordomain
func (e_ Error) NSStreamSOCKSErrorDomain() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("NSStreamSOCKSErrorDomain"))
	return rv
}


// The error domain used by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsocketsslerrordomain
func (e_ Error) NSStreamSocketSSLErrorDomain() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("NSStreamSocketSSLErrorDomain"))
	return rv
}


// URL loading system errors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlerrordomain
func (e_ Error) NSURLErrorDomain() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("NSURLErrorDomain"))
	return rv
}


