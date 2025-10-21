// Code generated from Apple documentation for Accounts. DO NOT EDIT.

package accounts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ACAccountStore] class.
var (
	ACAccountStoreClass     _ACAccountStoreClass
	ACAccountStoreClassOnce sync.Once
)

func getACAccountStoreClass() _ACAccountStoreClass {
	ACAccountStoreClassOnce.Do(func() {
		ACAccountStoreClass = _ACAccountStoreClass{objc.GetClass("ACAccountStore")}
	})
	return ACAccountStoreClass
}

type _ACAccountStoreClass struct {
	class objc.Class
}

// An interface definition for the [ACAccountStore] class.
type IACAccountStore interface {
	objectivec.IObject
	AccountWithIdentifier(identifier appkit.string) ACAccount
	AccountTypeWithAccountTypeIdentifier(typeIdentifier appkit.string) ACAccountType
	AccountsWithAccountType(accountType ACAccountType) foundation.Array
	RemoveAccountWithCompletionHandler(account IACAccount, completionHandler unsafe.Pointer)
	RenewCredentialsForAccountCompletion(account IACAccount, completionHandler unsafe.Pointer)
	RequestAccessToAccountsWithTypeOptionsCompletion(accountType ACAccountType, options objectivec.IObject, completion unsafe.Pointer)
	RequestAccessToAccountsWithTypeWithCompletionHandler(accountType ACAccountType, handler unsafe.Pointer)
	SaveAccountWithCompletionHandler(account IACAccount, completionHandler unsafe.Pointer)
}

// The object you use to request, manage, and store the user’s account information.
//
// The class provides an interface for accessing, managing, and storing accounts. To create and retrieve accounts from the Accounts database, you must create an object. Each object belongs to a single account store object.
//
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore
type ACAccountStore struct {
	objectivec.Object
}

// ACAccountStoreFrom constructs a [ACAccountStore] from an unsafe.Pointer.
//
// The object you use to request, manage, and store the user’s account information.
func ACAccountStoreFrom(ptr unsafe.Pointer) ACAccountStore {
	return ACAccountStore{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ACAccountStoreClass) Alloc() ACAccountStore {
	rv := objc.Send[ACAccountStore](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ACAccountStoreClass) New() ACAccountStore {
	rv := objc.Send[ACAccountStore](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ACAccountStore) Init() ACAccountStore {
	rv := objc.Send[ACAccountStore](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ACAccountStore) Autorelease() ACAccountStore {
	rv := objc.Send[ACAccountStore](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewACAccountStore creates a new ACAccountStore instance.
func NewACAccountStore() ACAccountStore {
	return getACAccountStoreClass().New()
}


// Returns the account with the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore/account(withIdentifier:)
func (a_ ACAccountStore) AccountWithIdentifier(identifier appkit.string) ACAccount {
	rv := objc.Send[ACAccount](a_.ID, objc.Sel("accountWithIdentifier:"), identifier)
	return rv
}

// Returns an account type that matches the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore/accountType(withAccountTypeIdentifier:)
func (a_ ACAccountStore) AccountTypeWithAccountTypeIdentifier(typeIdentifier appkit.string) ACAccountType {
	rv := objc.Send[ACAccountType](a_.ID, objc.Sel("accountTypeWithAccountTypeIdentifier:"), typeIdentifier)
	return rv
}

// Returns all accounts of the specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore/accounts(with:)
func (a_ ACAccountStore) AccountsWithAccountType(accountType ACAccountType) foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("accountsWithAccountType:"), accountType)
	return rv
}

// Removes an account from the account store.
//
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore/removeAccount(_:withCompletionHandler:)
func (a_ ACAccountStore) RemoveAccountWithCompletionHandler(account IACAccount, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeAccount:withCompletionHandler:"), account, completionHandler)
}

// Renews account credentials when the credentials are no longer valid.
//
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore/renewCredentials(for:completion:)
func (a_ ACAccountStore) RenewCredentialsForAccountCompletion(account IACAccount, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("renewCredentialsForAccount:completion:"), account, completionHandler)
}

// Obtains permission to access protected user properties.
//
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore/requestAccessToAccounts(with:options:completion:)
func (a_ ACAccountStore) RequestAccessToAccountsWithTypeOptionsCompletion(accountType ACAccountType, options objectivec.IObject, completion unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("requestAccessToAccountsWithType:options:completion:"), accountType, options, completion)
}

// Requests access to accounts of the specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore/requestAccessToAccountsWithType:withCompletionHandler:
func (a_ ACAccountStore) RequestAccessToAccountsWithTypeWithCompletionHandler(accountType ACAccountType, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("requestAccessToAccountsWithType:withCompletionHandler:"), accountType, handler)
}

// Saves an account to the Accounts database.
//
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore/saveAccount(_:withCompletionHandler:)
func (a_ ACAccountStore) SaveAccountWithCompletionHandler(account IACAccount, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("saveAccount:withCompletionHandler:"), account, completionHandler)
}

// The accounts managed by this account store.
//
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore/accounts
func (a_ ACAccountStore) Accounts() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("accounts"))
	return rv
}



