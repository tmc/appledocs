// Code generated from Apple documentation for Accounts. DO NOT EDIT.

package accounts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ACAccountType] class.
var (
	ACAccountTypeClass     _ACAccountTypeClass
	ACAccountTypeClassOnce sync.Once
)

func getACAccountTypeClass() _ACAccountTypeClass {
	ACAccountTypeClassOnce.Do(func() {
		ACAccountTypeClass = _ACAccountTypeClass{objc.GetClass("ACAccountType")}
	})
	return ACAccountTypeClass
}

type _ACAccountTypeClass struct {
	class objc.Class
}

// An interface definition for the [ACAccountType] class.
type IACAccountType interface {
	objectivec.IObject
	// properties:
	AccessGranted() bool
	AccountTypeDescription() objc.IObject /* cross-framework: NSString */
	Identifier() objc.IObject /* cross-framework: NSString */
	AccountType() IACAccountType
	SetAccountType(value IACAccountType)
	// methods:
}

// An object that encapsulates information about all accounts of a particular type.
//
// You don’t create account type objects directly. To obtain an account type object, use the method or the property of an account object. Use the method to obtain all accounts of a particular type.


// An object that encapsulates information about all accounts of a particular type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountType
type ACAccountType struct {
	objectivec.Object
}

// ACAccountTypeFrom constructs a [ACAccountType] from an unsafe.Pointer.
//
// An object that encapsulates information about all accounts of a particular type.
func ACAccountTypeFrom(ptr unsafe.Pointer) ACAccountType {
	return ACAccountType{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ACAccountTypeClass) Alloc() ACAccountType {
	rv := objc.Send[ACAccountType](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ACAccountTypeClass) New() ACAccountType {
	rv := objc.Send[ACAccountType](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ACAccountType) Init() ACAccountType {
	rv := objc.Send[ACAccountType](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ACAccountType) Autorelease() ACAccountType {
	rv := objc.Send[ACAccountType](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewACAccountType creates a new ACAccountType instance.
func NewACAccountType() ACAccountType {
	return getACAccountTypeClass().New()
}



// A Boolean value indicating whether the user granted the application access to accounts of this type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountType/accessGranted
func (a_ ACAccountType) AccessGranted() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("accessGranted"))
	return rv
}


// A human-readable description of the account type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountType/accountTypeDescription
func (a_ ACAccountType) AccountTypeDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("accountTypeDescription"))
	return rv
}


// The unique identifier for the account type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountType/identifier
func (a_ ACAccountType) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("identifier"))
	return rv
}


// The type of service account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accounts/acaccount/accounttype
func (a_ ACAccountType) AccountType() IACAccountType {
	rv := objc.Send[ACAccountType](a_.ID, objc.Sel("accountType"))
	return rv
}


// The type of service account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accounts/acaccount/accounttype
func (a_ ACAccountType) SetAccountType(value IACAccountType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccountType:"), value)
}




