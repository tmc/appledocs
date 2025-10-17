// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLCredential] class.
var URLCredentialClass objc.Class

func init() {
	URLCredentialClass = objc.GetClass("NSURLCredential")
}

type URLCredential struct {
	objc.ID
}

func URLCredentialFrom(ptr unsafe.Pointer) URLCredential {
	return URLCredential{
		ID: objc.ID(ptr),
	}
}




