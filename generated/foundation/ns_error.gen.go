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
	HelpAnchor() string /* primitive/slice/pointer. */
	LocalizedDescription() string /* primitive/slice/pointer. */
	LocalizedFailureReason() string /* primitive/slice/pointer. */
	UserInfo() IDictionary /* already interface */
	NSCocoaErrorDomain() string /* primitive/slice/pointer. */
	LocalizedRecoveryOptions() string /* primitive/slice/pointer. */
	SetLocalizedRecoveryOptions(value string /* primitive/slice/pointer. */)
	LocalizedRecoverySuggestion() string /* primitive/slice/pointer. */
	SetLocalizedRecoverySuggestion(value string /* primitive/slice/pointer. */)
	RecoveryAttempter() unsafe.Pointer
	SetRecoveryAttempter(value unsafe.Pointer)
	UnderlyingErrors() IError
	SetUnderlyingErrors(value IError)
	NSMachErrorDomain() string /* primitive/slice/pointer. */
	NSOSStatusErrorDomain() string /* primitive/slice/pointer. */
	NSPOSIXErrorDomain() string /* primitive/slice/pointer. */
	NSRecoveryAttempterErrorKey() string /* primitive/slice/pointer. */
	NSStreamSOCKSErrorDomain() string /* primitive/slice/pointer. */
	NSStreamSocketSSLErrorDomain() string /* primitive/slice/pointer. */
	NSURLErrorDomain() string /* primitive/slice/pointer. */
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
func (e_ Error) HelpAnchor() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("helpAnchor"))
	return rv
}


// A string containing the localized description of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/localizedDescription
func (e_ Error) LocalizedDescription() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("localizedDescription"))
	return rv
}


// A string containing the localized explanation of the reason for the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/localizedFailureReason
func (e_ Error) LocalizedFailureReason() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("localizedFailureReason"))
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
func (e_ Error) NSCocoaErrorDomain() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("NSCocoaErrorDomain"))
	return rv
}


// An array containing the localized titles of buttons appropriate for displaying in an alert panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/localizedrecoveryoptions
func (e_ Error) LocalizedRecoveryOptions() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("localizedRecoveryOptions"))
	return rv
}


// An array containing the localized titles of buttons appropriate for displaying in an alert panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/localizedrecoveryoptions
func (e_ Error) SetLocalizedRecoveryOptions(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLocalizedRecoveryOptions:"), objc.String(value))
}


// A string containing the localized recovery suggestion for the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/localizedrecoverysuggestion
func (e_ Error) LocalizedRecoverySuggestion() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("localizedRecoverySuggestion"))
	return rv
}


// A string containing the localized recovery suggestion for the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/localizedrecoverysuggestion
func (e_ Error) SetLocalizedRecoverySuggestion(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLocalizedRecoverySuggestion:"), objc.String(value))
}


// The object in the user info dictionary corresponding to the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/recoveryattempter
func (e_ Error) RecoveryAttempter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("recoveryAttempter"))
	return rv
}


// The object in the user info dictionary corresponding to the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/recoveryattempter
func (e_ Error) SetRecoveryAttempter(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRecoveryAttempter:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/underlyingerrors
func (e_ Error) UnderlyingErrors() IError {
	rv := objc.Send[Error](e_.ID, objc.Sel("underlyingErrors"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/underlyingerrors
func (e_ Error) SetUnderlyingErrors(value IError) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUnderlyingErrors:"), value)
}


// Mach errors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmacherrordomain
func (e_ Error) NSMachErrorDomain() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("NSMachErrorDomain"))
	return rv
}


// Mac OS 9/Carbon errors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsosstatuserrordomain
func (e_ Error) NSOSStatusErrorDomain() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("NSOSStatusErrorDomain"))
	return rv
}


// POSIX/BSD errors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsposixerrordomain
func (e_ Error) NSPOSIXErrorDomain() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("NSPOSIXErrorDomain"))
	return rv
}


// The corresponding value is an object that conforms to the NSErrorRecoveryAttempting informal protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrecoveryattemptererrorkey
func (e_ Error) NSRecoveryAttempterErrorKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("NSRecoveryAttempterErrorKey"))
	return rv
}


// The error domain used by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsockserrordomain
func (e_ Error) NSStreamSOCKSErrorDomain() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("NSStreamSOCKSErrorDomain"))
	return rv
}


// The error domain used by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsocketsslerrordomain
func (e_ Error) NSStreamSocketSSLErrorDomain() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("NSStreamSocketSSLErrorDomain"))
	return rv
}


// URL loading system errors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlerrordomain
func (e_ Error) NSURLErrorDomain() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("NSURLErrorDomain"))
	return rv
}



