// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLCredential] class.
var URLCredentialClass = _URLCredentialClass{objc.GetClass("NSURLCredential")}

type _URLCredentialClass struct {
	class objc.Class
}

type URLCredential struct {
	objc.ID
}

func URLCredentialFrom(ptr unsafe.Pointer) URLCredential {
	return URLCredential{
		ID: objc.ID(ptr),
	}
}




