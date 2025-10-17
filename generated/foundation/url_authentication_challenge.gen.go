// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLAuthenticationChallenge] class.
var URLAuthenticationChallengeClass objc.Class

func init() {
	URLAuthenticationChallengeClass = objc.GetClass("NSURLAuthenticationChallenge")
}

type URLAuthenticationChallenge struct {
	objc.ID
}

func URLAuthenticationChallengeFrom(ptr unsafe.Pointer) URLAuthenticationChallenge {
	return URLAuthenticationChallenge{
		ID: objc.ID(ptr),
	}
}



