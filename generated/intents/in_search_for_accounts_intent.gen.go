// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSearchForAccountsIntent] class.
var (
	INSearchForAccountsIntentClass     _INSearchForAccountsIntentClass
	INSearchForAccountsIntentClassOnce sync.Once
)

func getINSearchForAccountsIntentClass() _INSearchForAccountsIntentClass {
	INSearchForAccountsIntentClassOnce.Do(func() {
		INSearchForAccountsIntentClass = _INSearchForAccountsIntentClass{objc.GetClass("INSearchForAccountsIntent")}
	})
	return INSearchForAccountsIntentClass
}

type _INSearchForAccountsIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSearchForAccountsIntent] class.
type IINSearchForAccountsIntent interface {
	IINIntent
}

// A user request for information about their accounts in your app.
//
// Siri creates an instance of when the user asks for information about accounts available in your app. Users can ask for information about monetary and nonmonetary accounts. For example, an airline app might allow the user to search for an account that manages their frequent flier miles. The user can ask for information such as a list of accounts or the balance of a specific account. Use the data the intent provides to find the user’s accounts and return the information the user requests. To process the request, your handler must adopt the protocol. When your implementation confirms the request, provide an instance of that includes the accounts that meet the user’s critieria. If the search is successful, Siri offers the user a way to view the results. is only available to Siri Intents and requires an unlocked device before processing. Siri performs the following actions automatically: Requests the user’s confirmation before passing the request to your app or Intents extension for processing. Asks the user to unlock a locked device.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForAccountsIntent
type INSearchForAccountsIntent struct {
	INIntent
}

// INSearchForAccountsIntentFrom constructs a [INSearchForAccountsIntent] from an unsafe.Pointer.
//
// A user request for information about their accounts in your app.
func INSearchForAccountsIntentFrom(ptr unsafe.Pointer) INSearchForAccountsIntent {
	return INSearchForAccountsIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSearchForAccountsIntentClass) Alloc() INSearchForAccountsIntent {
	rv := objc.Send[INSearchForAccountsIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSearchForAccountsIntentClass) New() INSearchForAccountsIntent {
	rv := objc.Send[INSearchForAccountsIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSearchForAccountsIntent) Init() INSearchForAccountsIntent {
	rv := objc.Send[INSearchForAccountsIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSearchForAccountsIntent) Autorelease() INSearchForAccountsIntent {
	rv := objc.Send[INSearchForAccountsIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSearchForAccountsIntent creates a new INSearchForAccountsIntent instance.
func NewINSearchForAccountsIntent() INSearchForAccountsIntent {
	return getINSearchForAccountsIntentClass().New()
}


// The nickname that the user assigned to the account.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforaccountsintent/accountnickname
func (i_ INSearchForAccountsIntent) AccountNickname() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("accountNickname"))
	return rv
}


// SetAccountNickname sets the value of the accountNickname property.
// The nickname that the user assigned to the account.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforaccountsintent/accountnickname
func (i_ INSearchForAccountsIntent) SetAccountNickname(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAccountNickname:"), value)
}

// The type of the account.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforaccountsintent/accounttype
func (i_ INSearchForAccountsIntent) AccountType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("accountType"))
	return rv
}


// SetAccountType sets the value of the accountType property.
// The type of the account.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforaccountsintent/accounttype
func (i_ INSearchForAccountsIntent) SetAccountType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAccountType:"), value)
}

// The name of an organization attached to the account.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforaccountsintent/organizationname
func (i_ INSearchForAccountsIntent) OrganizationName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("organizationName"))
	return rv
}


// SetOrganizationName sets the value of the organizationName property.
// The name of an organization attached to the account.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforaccountsintent/organizationname
func (i_ INSearchForAccountsIntent) SetOrganizationName(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOrganizationName:"), value)
}

// The balance type requested by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforaccountsintent/requestedbalancetype
func (i_ INSearchForAccountsIntent) RequestedBalanceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("requestedBalanceType"))
	return rv
}


// SetRequestedBalanceType sets the value of the requestedBalanceType property.
// The balance type requested by the user.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforaccountsintent/requestedbalancetype
func (i_ INSearchForAccountsIntent) SetRequestedBalanceType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRequestedBalanceType:"), value)
}



