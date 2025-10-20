// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var URLCredentialStorageClass _URLCredentialStorageClass

func init() {
	URLCredentialStorageClass = _URLCredentialStorageClass{objc.GetClass("NSURLCredentialStorage")}
}

type _URLCredentialStorageClass struct {
	class objc.Class
}

type URLCredentialStorage struct {
	objc.ID
}

func URLCredentialStorageFrom(ptr unsafe.Pointer) URLCredentialStorage {
	return URLCredentialStorage{
		ID: objc.ID(ptr),
	}
}




