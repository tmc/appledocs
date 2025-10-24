// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLAuthenticationChallenge */


/* debug [class_header]: Header for NSURLAuthenticationChallenge */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLAuthenticationChallenge */
// An interface definition for the [URLAuthenticationChallenge] class.
type IURLAuthenticationChallenge interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLAuthenticationChallenge */
	// properties:
	Error() IError
	FailureResponse() IURLResponse
	PreviousFailureCount() int
	ProposedCredential() IURLCredential
	ProtectionSpace() IURLProtectionSpace
	Sender() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLAuthenticationChallenge */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLAuthenticationChallenge */
// Alloc allocates a new instance without initialization.
func (uc _URLAuthenticationChallengeClass) Alloc() URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLAuthenticationChallenge */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLAuthenticationChallenge *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLAuthenticationChallenge */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLAuthenticationChallenge */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLAuthenticationChallenge */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLAuthenticationChallenge */

// The error object representing the last authentication failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/error
func (u_ URLAuthenticationChallenge) Error() IError {
	rv := objc.Send[Error](u_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// The URL response object representing the last authentication failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/failureResponse
func (u_ URLAuthenticationChallenge) FailureResponse() IURLResponse {
	rv := objc.Send[URLResponse](u_.ID, objc.Sel("failureResponse"))
	return rv
}/* debug [instance_properties/getter]: failureResponse */


// The receiver’s count of failed authentication attempts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/previousFailureCount
func (u_ URLAuthenticationChallenge) PreviousFailureCount() int {
	rv := objc.Send[int](u_.ID, objc.Sel("previousFailureCount"))
	return rv
}/* debug [instance_properties/getter]: previousFailureCount */


// The proposed credential for this challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/proposedCredential
func (u_ URLAuthenticationChallenge) ProposedCredential() IURLCredential {
	rv := objc.Send[URLCredential](u_.ID, objc.Sel("proposedCredential"))
	return rv
}/* debug [instance_properties/getter]: proposedCredential */


// The receiver’s protection space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/protectionSpace
func (u_ URLAuthenticationChallenge) ProtectionSpace() IURLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](u_.ID, objc.Sel("protectionSpace"))
	return rv
}/* debug [instance_properties/getter]: protectionSpace */


// The sender of the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/sender
func (u_ URLAuthenticationChallenge) Sender() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("sender"))
	return rv
}/* debug [instance_properties/getter]: sender */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLAuthenticationChallenge */



