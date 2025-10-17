// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLAuthenticationChallenge] class.
var uRLAuthenticationChallengeClass = _URLAuthenticationChallengeClass{objc.GetClass("NSURLAuthenticationChallenge")}

type _URLAuthenticationChallengeClass struct {
	class objc.Class
}

// An interface definition for the [URLAuthenticationChallenge] class.
type IURLAuthenticationChallenge interface {
	objectivec.IObject
}

// A challenge from a server requiring authentication from the client. [Full Topic]
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



