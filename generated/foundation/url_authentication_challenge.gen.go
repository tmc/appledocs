// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLAuthenticationChallenge] class.
var URLAuthenticationChallengeClass = _URLAuthenticationChallengeClass{objc.GetClass("NSURLAuthenticationChallenge")}

type _URLAuthenticationChallengeClass struct {
	class objc.Class
}

type URLAuthenticationChallenge struct {
	objc.ID
}

func URLAuthenticationChallengeFrom(ptr unsafe.Pointer) URLAuthenticationChallenge {
	return URLAuthenticationChallenge{
		ID: objc.ID(ptr),
	}
}




