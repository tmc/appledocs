// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLCredential] class.
var uRLCredentialClass = _URLCredentialClass{objc.GetClass("NSURLCredential")}

type _URLCredentialClass struct {
	class objc.Class
}

// n authentication credential consisting of information specific to the type of credential and the type of persistent storage to use, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential

type URLCredential struct {
	objectivec.Object
}

// URLCredentialFrom constructs a [URLCredential] from an unsafe.Pointer.
//
// n authentication credential consisting of information specific to the type of credential and the type of persistent storage to use, if any.
func URLCredentialFrom(ptr unsafe.Pointer) URLCredential {
	return URLCredential{objectivec.Object{objc.ID(ptr)}}
}



