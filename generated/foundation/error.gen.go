// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Error] class.
var ErrorClass objc.Class

func init() {
	ErrorClass = objc.GetClass("NSError")
}

type Error struct {
	objc.ID
}

func ErrorFrom(ptr unsafe.Pointer) Error {
	return Error{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec Error) Alloc() Error {
	ret := objc.ID(ErrorClass).Send(objc.RegisterName("alloc"))
	return Error{ret}
}

// Init initializes the instance.
func (e_ Error) Init() Error {
	ret := e_.ID.Send(objc.RegisterName("init"))
	return Error{ret}
}
// Returns an   object initialized for a given domain and code with a given   dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSError/init(domain:code:userInfo:)
func NewErrorWithDomainCodeUserInfo(domain unsafe.Pointer, code int, dict unsafe.Pointer) Error {
	instance := Error{}.Alloc()
	sel := objc.RegisterName("initWithDomain:code:userInfo:")
	ret := instance.ID.Send(sel, domain, code, dict)
	instance = Error{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Creates and initializes an   object for a given domain and code with a given   dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSError/errorWithDomain:code:userInfo:
func (ec Error) ErrorWithDomainCodeUserInfo(domain unsafe.Pointer, code int, dict unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("errorWithDomain:code:userInfo:")
	ret := objc.ID(ErrorClass).Send(sel, domain, code, dict)
	return unsafe.Pointer(ret)
}
// Returns a properly formatted error object with a   error code. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSError/fileProviderErrorForCollision(with:)
func (ec Error) FileProviderErrorForCollisionWithItem(existingItem unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("fileProviderErrorForCollisionWithItem:")
	ret := objc.ID(ErrorClass).Send(sel, existingItem)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSError/fileProviderErrorForNonExistentItem(withIdentifier:)
func (ec Error) FileProviderErrorForNonExistentItemWithIdentifier(itemIdentifier unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("fileProviderErrorForNonExistentItemWithIdentifier:")
	ret := objc.ID(ErrorClass).Send(sel, itemIdentifier)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSError/fileProviderErrorForRejectedDeletion(of:)
func (ec Error) FileProviderErrorForRejectedDeletionOfItem(updatedVersion unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("fileProviderErrorForRejectedDeletionOfItem:")
	ret := objc.ID(ErrorClass).Send(sel, updatedVersion)
	return unsafe.Pointer(ret)
}
// Specifies a block to call when the corresponding property is not present in the user info dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSError/setUserInfoValueProvider(forDomain:provider:)
func (ec Error) SetUserInfoValueProviderForDomainProvider(errorDomain unsafe.Pointer, provider unsafe.Pointer) {
	sel := objc.RegisterName("setUserInfoValueProviderForDomain:provider:")
	objc.ID(ErrorClass).Send(sel, errorDomain, provider)
}
// Returns any user info provider specified for a given error domain. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSError/userInfoValueProvider(forDomain:)
func (ec Error) UserInfoValueProviderForDomain(errorDomain unsafe.Pointer) {
	sel := objc.RegisterName("userInfoValueProviderForDomain:")
	objc.ID(ErrorClass).Send(sel, errorDomain)
}

