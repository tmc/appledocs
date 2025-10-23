// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLAuthenticationChallenge] class.
var (
	URLAuthenticationChallengeClass     _URLAuthenticationChallengeClass
	URLAuthenticationChallengeClassOnce sync.Once
)

func getURLAuthenticationChallengeClass() _URLAuthenticationChallengeClass {
	URLAuthenticationChallengeClassOnce.Do(func() {
		URLAuthenticationChallengeClass = _URLAuthenticationChallengeClass{objc.GetClass("NSURLAuthenticationChallenge")}
	})
	return URLAuthenticationChallengeClass
}

type _URLAuthenticationChallengeClass struct {
	class objc.Class
}

// An interface definition for the [URLAuthenticationChallenge] class.
type IURLAuthenticationChallenge interface {
	objectivec.IObject
	// properties:
	PreviousFailureCount() int /* primitive/slice/pointer. */
	ProposedCredential() IURLCredential
	Error() IError
	SetError(value IError)
	FailureResponse() IURLResponse
	SetFailureResponse(value IURLResponse)
	ProtectionSpace() IURLProtectionSpace
	SetProtectionSpace(value IURLProtectionSpace)
	Sender() unsafe.Pointer
	SetSender(value unsafe.Pointer)
	// methods:
}

// A challenge from a server requiring authentication from the client.
//
// Your app receives authentication challenges in various , , and delegate methods, such as . These objects provide the information you’ll need when deciding how to handle a server’s request for authentication. At the core of that authentication challenge is a that defines the type of authentication being requested, the host and port number, the networking protocol, and (where applicable) the authentication realm (a group of related URLs on the same server that share a single set of credentials).


// A challenge from a server requiring authentication from the client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge
type URLAuthenticationChallenge struct {
	objectivec.Object
}

// URLAuthenticationChallengeFrom constructs a [URLAuthenticationChallenge] from an unsafe.Pointer.
//
// A challenge from a server requiring authentication from the client.
func URLAuthenticationChallengeFrom(ptr unsafe.Pointer) URLAuthenticationChallenge {
	return URLAuthenticationChallenge{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLAuthenticationChallengeClass) Alloc() URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLAuthenticationChallengeClass) New() URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLAuthenticationChallenge) Init() URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLAuthenticationChallenge) Autorelease() URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLAuthenticationChallenge creates a new URLAuthenticationChallenge instance.
func NewURLAuthenticationChallenge() URLAuthenticationChallenge {
	return getURLAuthenticationChallengeClass().New()
}



// The receiver’s count of failed authentication attempts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/previousFailureCount
func (u_ URLAuthenticationChallenge) PreviousFailureCount() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](u_.ID, objc.Sel("previousFailureCount"))
	return rv
}


// The proposed credential for this challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/proposedCredential
func (u_ URLAuthenticationChallenge) ProposedCredential() IURLCredential {
	rv := objc.Send[URLCredential](u_.ID, objc.Sel("proposedCredential"))
	return rv
}


// The error object representing the last authentication failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlauthenticationchallenge/error
func (u_ URLAuthenticationChallenge) Error() IError {
	rv := objc.Send[Error](u_.ID, objc.Sel("error"))
	return rv
}


// The error object representing the last authentication failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlauthenticationchallenge/error
func (u_ URLAuthenticationChallenge) SetError(value IError) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setError:"), value)
}


// The URL response object representing the last authentication failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlauthenticationchallenge/failureresponse
func (u_ URLAuthenticationChallenge) FailureResponse() IURLResponse {
	rv := objc.Send[URLResponse](u_.ID, objc.Sel("failureResponse"))
	return rv
}


// The URL response object representing the last authentication failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlauthenticationchallenge/failureresponse
func (u_ URLAuthenticationChallenge) SetFailureResponse(value IURLResponse) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setFailureResponse:"), value)
}


// The receiver’s protection space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlauthenticationchallenge/protectionspace
func (u_ URLAuthenticationChallenge) ProtectionSpace() IURLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](u_.ID, objc.Sel("protectionSpace"))
	return rv
}


// The receiver’s protection space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlauthenticationchallenge/protectionspace
func (u_ URLAuthenticationChallenge) SetProtectionSpace(value IURLProtectionSpace) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setProtectionSpace:"), value)
}


// The sender of the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlauthenticationchallenge/sender
func (u_ URLAuthenticationChallenge) Sender() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("sender"))
	return rv
}


// The sender of the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlauthenticationchallenge/sender
func (u_ URLAuthenticationChallenge) SetSender(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSender:"), value)
}



