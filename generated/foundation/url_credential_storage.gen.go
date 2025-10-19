// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLCredentialStorage] class.
var (
	uRLCredentialStorageClass     _URLCredentialStorageClass
	uRLCredentialStorageClassOnce sync.Once
)

func getURLCredentialStorageClass() _URLCredentialStorageClass {
	uRLCredentialStorageClassOnce.Do(func() {
		uRLCredentialStorageClass = _URLCredentialStorageClass{objc.GetClass("NSURLCredentialStorage")}
	})
	return uRLCredentialStorageClass
}

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

// Alloc allocates a new instance without initialization.
func (uc _URLCredentialStorageClass) Alloc() URLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLCredentialStorageClass) New() URLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLCredentialStorage) Init() URLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLCredentialStorage) Autorelease() URLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLCredentialStorage creates a new URLCredentialStorage instance.
func NewURLCredentialStorage() URLCredentialStorage {
	return getURLCredentialStorageClass().New()
}




