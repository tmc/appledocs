// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Error] class.
var (
	errorClass     _ErrorClass
	errorClassOnce sync.Once
)

func getErrorClass() _ErrorClass {
	errorClassOnce.Do(func() {
		errorClass = _ErrorClass{objc.GetClass("NSError")}
	})
	return errorClass
}

type _ErrorClass struct {
	class objc.Class
}

// An interface definition for the [Error] class.
type IError interface {
	objectivec.IObject
}

// Information about an error condition including a domain, a domain-specific error code, and application-specific information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError
type Error struct {
	objectivec.Object
}

// ErrorFrom constructs a [Error] from an unsafe.Pointer.
//
// Information about an error condition including a domain, a domain-specific error code, and application-specific information.
func ErrorFrom(ptr unsafe.Pointer) Error {
	return Error{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _ErrorClass) Alloc() Error {
	rv := objc.Send[Error](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _ErrorClass) New() Error {
	rv := objc.Send[Error](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ Error) Init() Error {
	rv := objc.Send[Error](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ Error) Autorelease() Error {
	rv := objc.Send[Error](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewError creates a new Error instance.
func NewError() Error {
	return getErrorClass().New()
}


// Returns an object initialized for a given domain and code with a given dictionary. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/init(domain:code:userInfo:)
func NewErrorWithDomainCodeUserInfo(domain unsafe.Pointer, code int, dict unsafe.Pointer) Error {
	instance := getErrorClass().Alloc()
	rv := objc.Send[Error](instance.ID, objc.Sel("initWithDomain:code:userInfo:"), domain, code, dict)
	rv.Autorelease()
	return rv
}


// Creates and initializes an object for a given domain and code with a given dictionary. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/errorWithDomain:code:userInfo:
func (ec _ErrorClass) ErrorWithDomainCodeUserInfo(domain unsafe.Pointer, code int, dict unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("errorWithDomain:code:userInfo:"), domain, code, dict)
	return rv
}
// Returns a properly formatted error object with a error code. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/fileProviderErrorForCollision(with:)
func (ec _ErrorClass) FileProviderErrorForCollisionWithItem(existingItem unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("fileProviderErrorForCollisionWithItem:"), existingItem)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/fileProviderErrorForNonExistentItem(withIdentifier:)
func (ec _ErrorClass) FileProviderErrorForNonExistentItemWithIdentifier(itemIdentifier unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("fileProviderErrorForNonExistentItemWithIdentifier:"), itemIdentifier)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/fileProviderErrorForRejectedDeletion(of:)
func (ec _ErrorClass) FileProviderErrorForRejectedDeletionOfItem(updatedVersion unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("fileProviderErrorForRejectedDeletionOfItem:"), updatedVersion)
	return rv
}
// Specifies a block to call when the corresponding property is not present in the user info dictionary. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/setUserInfoValueProvider(forDomain:provider:)
func (ec _ErrorClass) SetUserInfoValueProviderForDomainProvider(errorDomain unsafe.Pointer, provider unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("setUserInfoValueProviderForDomain:provider:"), errorDomain, provider)
}
// Returns any user info provider specified for a given error domain. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfoValueProvider(forDomain:)
func (ec _ErrorClass) UserInfoValueProviderForDomain(errorDomain unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("userInfoValueProviderForDomain:"), errorDomain)
}

