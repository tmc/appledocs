// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLCredentialStorage] class.
var uRLCredentialStorageClass = _URLCredentialStorageClass{objc.GetClass("NSURLCredentialStorage")}

type _URLCredentialStorageClass struct {
	class objc.Class
}

// An interface definition for the [URLCredentialStorage] class.
type IURLCredentialStorage interface {
	objectivec.IObject
}

// The manager of a shared credentials cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage

type URLCredentialStorage struct {
	objectivec.Object
}

// URLCredentialStorageFrom constructs a [URLCredentialStorage] from an unsafe.Pointer.
//
// The manager of a shared credentials cache.
func URLCredentialStorageFrom(ptr unsafe.Pointer) URLCredentialStorage {
	return URLCredentialStorage{objectivec.Object{objc.ID(ptr)}}
}



