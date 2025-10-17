// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLCredentialStorage] class.
var URLCredentialStorageClass objc.Class

func init() {
	URLCredentialStorageClass = objc.GetClass("NSURLCredentialStorage")
}

type URLCredentialStorage struct {
	objc.ID
}

func URLCredentialStorageFrom(ptr unsafe.Pointer) URLCredentialStorage {
	return URLCredentialStorage{
		ID: objc.ID(ptr),
	}
}




