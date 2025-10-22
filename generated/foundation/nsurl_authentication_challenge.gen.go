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
	FailureResponse() NSURLResponse
	PreviousFailureCount() int
	ProposedCredential() NSURLCredential
	Sender() objc.ID
	Error() Error
	SetError(value IError)
	ProtectionSpace() NSURLProtectionSpace
	SetProtectionSpace(value IURLProtectionSpace)
}

// A challenge from a server requiring authentication from the client.
//
// Your app receives authentication challenges in various , , and delegate methods, such as . These objects provide the information you’ll need when deciding how to handle a server’s request for authentication. At the core of that authentication challenge is a that defines the type of authentication being requested, the host and port number, the networking protocol, and (where applicable) the authentication realm (a group of related URLs on the same server that share a single set of credentials).
//
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


// The URL response object representing the last authentication failure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/failureResponse
func (u_ URLAuthenticationChallenge) FailureResponse() NSURLResponse {
	rv := objc.Send[NSURLResponse](u_.ID, objc.Sel("failureResponse"))
	return rv
}

// The receiver’s count of failed authentication attempts.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/previousFailureCount
func (u_ URLAuthenticationChallenge) PreviousFailureCount() int {
	rv := objc.Send[int](u_.ID, objc.Sel("previousFailureCount"))
	return rv
}

// The proposed credential for this challenge.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/proposedCredential
func (u_ URLAuthenticationChallenge) ProposedCredential() NSURLCredential {
	rv := objc.Send[NSURLCredential](u_.ID, objc.Sel("proposedCredential"))
	return rv
}

// The sender of the challenge.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/sender
func (u_ URLAuthenticationChallenge) Sender() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("sender"))
	return rv
}

// The error object representing the last authentication failure.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlauthenticationchallenge/error
func (u_ URLAuthenticationChallenge) Error() Error {
	rv := objc.Send[Error](u_.ID, objc.Sel("error"))
	return rv
}


// SetError sets the value of the error property.
// The error object representing the last authentication failure.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlauthenticationchallenge/error
func (u_ URLAuthenticationChallenge) SetError(value IError) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setError:"), value)
}

// The receiver’s protection space.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlauthenticationchallenge/protectionspace
func (u_ URLAuthenticationChallenge) ProtectionSpace() NSURLProtectionSpace {
	rv := objc.Send[NSURLProtectionSpace](u_.ID, objc.Sel("protectionSpace"))
	return rv
}


// SetProtectionSpace sets the value of the protectionSpace property.
// The receiver’s protection space.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlauthenticationchallenge/protectionspace
func (u_ URLAuthenticationChallenge) SetProtectionSpace(value IURLProtectionSpace) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setProtectionSpace:"), value)
}



