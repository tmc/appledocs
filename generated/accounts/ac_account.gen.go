// Code generated from Apple documentation for Accounts. DO NOT EDIT.

package accounts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ACAccount] class.
var (
	ACAccountClass     _ACAccountClass
	ACAccountClassOnce sync.Once
)

func getACAccountClass() _ACAccountClass {
	ACAccountClassOnce.Do(func() {
		ACAccountClass = _ACAccountClass{objc.GetClass("ACAccount")}
	})
	return ACAccountClass
}

type _ACAccountClass struct {
	class objc.Class
}

// An interface definition for the [ACAccount] class.
type IACAccount interface {
	objectivec.IObject
	// properties:
	AccountDescription() objc.IObject /* cross-framework: NSString */
	SetAccountDescription(value objc.IObject /* cross-framework: NSString */)
	AccountType() IACAccountType
	SetAccountType(value IACAccountType)
	Credential() IACAccountCredential
	SetCredential(value IACAccountCredential)
	Identifier() objc.IObject /* cross-framework: NSString */
	Username() objc.IObject /* cross-framework: NSString */
	SetUsername(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// The information associated with one of the user’s accounts.
//
// An object encapsulates information about a user account stored in the Accounts database. You can create and retrieve accounts using an object. The object provides an interface to the persistent Accounts database. For each user, all account objects belong to a single object.


// The information associated with one of the user’s accounts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccount
type ACAccount struct {
	objectivec.Object
}

// ACAccountFrom constructs a [ACAccount] from an unsafe.Pointer.
//
// The information associated with one of the user’s accounts.
func ACAccountFrom(ptr unsafe.Pointer) ACAccount {
	return ACAccount{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ACAccountClass) Alloc() ACAccount {
	rv := objc.Send[ACAccount](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ACAccountClass) New() ACAccount {
	rv := objc.Send[ACAccount](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ACAccount) Init() ACAccount {
	rv := objc.Send[ACAccount](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ACAccount) Autorelease() ACAccount {
	rv := objc.Send[ACAccount](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewACAccount creates a new ACAccount instance.
func NewACAccount() ACAccount {
	return getACAccountClass().New()
}



// Initializes a new account of the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccount/init(accountType:)
func NewACAccountWithAccountType(type_ IACAccountType) ACAccount {
	instance := getACAccountClass().Alloc()
	rv := objc.Send[ACAccount](instance.ID, objc.Sel("initWithAccountType:"), type_)
	rv.Autorelease()
	return rv
}



// A human-readable description of the account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccount/accountDescription
func (a_ ACAccount) AccountDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("accountDescription"))
	return rv
}


// A human-readable description of the account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccount/accountDescription
func (a_ ACAccount) SetAccountDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccountDescription:"), value)
}


// The type of service account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccount/accountType
func (a_ ACAccount) AccountType() IACAccountType {
	rv := objc.Send[ACAccountType](a_.ID, objc.Sel("accountType"))
	return rv
}


// The type of service account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccount/accountType
func (a_ ACAccount) SetAccountType(value IACAccountType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccountType:"), value)
}


// The credential used to authenticate the user of this account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccount/credential
func (a_ ACAccount) Credential() IACAccountCredential {
	rv := objc.Send[ACAccountCredential](a_.ID, objc.Sel("credential"))
	return rv
}


// The credential used to authenticate the user of this account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccount/credential
func (a_ ACAccount) SetCredential(value IACAccountCredential) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCredential:"), value)
}


// A unique identifier for this account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccount/identifier
func (a_ ACAccount) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("identifier"))
	return rv
}


// The username for this account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccount/username
func (a_ ACAccount) Username() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("username"))
	return rv
}


// The username for this account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccount/username
func (a_ ACAccount) SetUsername(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUsername:"), value)
}


