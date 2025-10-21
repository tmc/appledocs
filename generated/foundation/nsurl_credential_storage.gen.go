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
	URLCredentialStorageClass     _URLCredentialStorageClass
	URLCredentialStorageClassOnce sync.Once
)

func getURLCredentialStorageClass() _URLCredentialStorageClass {
	URLCredentialStorageClassOnce.Do(func() {
		URLCredentialStorageClass = _URLCredentialStorageClass{objc.GetClass("NSURLCredentialStorage")}
	})
	return URLCredentialStorageClass
}

type _URLCredentialStorageClass struct {
	class objc.Class
}

// An interface definition for the [URLCredentialStorage] class.
type IURLCredentialStorage interface {
	objectivec.IObject
	CredentialsForProtectionSpace(space unsafe.Pointer) unsafe.Pointer
	SetDefaultCredentialForProtectionSpace(credential unsafe.Pointer, space unsafe.Pointer)
}

// The manager of a shared credentials cache.
//
// The shared cache stores and retrieves instances of . You can store password-based credentials permanently, based on the they were created with. Certificate-based credentials are never stored permanently.
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

// Returns a dictionary containing the credentials for the specified protection space.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/credentials(for:)
func (u_ URLCredentialStorage) CredentialsForProtectionSpace(space unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("credentialsForProtectionSpace:"), space)
	return rv
}

// Sets the default credential for a specified protection space.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/setDefaultCredential(_:for:)
func (u_ URLCredentialStorage) SetDefaultCredentialForProtectionSpace(credential unsafe.Pointer, space unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDefaultCredential:forProtectionSpace:"), credential, space)
}
