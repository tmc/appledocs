// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccountAuthenticationModificationViewController] class.
var (
	AccountAuthenticationModificationViewControllerClass     _AccountAuthenticationModificationViewControllerClass
	AccountAuthenticationModificationViewControllerClassOnce sync.Once
)

func getAccountAuthenticationModificationViewControllerClass() _AccountAuthenticationModificationViewControllerClass {
	AccountAuthenticationModificationViewControllerClassOnce.Do(func() {
		AccountAuthenticationModificationViewControllerClass = _AccountAuthenticationModificationViewControllerClass{objc.GetClass("ASAccountAuthenticationModificationViewController")}
	})
	return AccountAuthenticationModificationViewControllerClass
}

type _AccountAuthenticationModificationViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [AccountAuthenticationModificationViewController] class.
type IAccountAuthenticationModificationViewController interface {
	appkit.IViewController
	// properties:
	ExtensionContext() IASAccountAuthenticationModificationExtensionContext
	SetExtensionContext(value IASAccountAuthenticationModificationExtensionContext)
	// methods:
	PrepareInterfaceToConvertAccountToSignInWithAppleForServiceIdentifierExistingCredentialUserInfo(serviceIdentifier CredentialServiceIdentifier /* not a class type */, existingCredential IASPasswordCredential, userInfo objectivec.IObject)
}

// A view controller that can upgrade user passwords to strong passwords, or convert accounts to use Sign in with Apple.
//
// Adding an account modification extension lets your app seamlessly upgrade user passwords to strong passwords, or convert from using passwords to using Sign in with Apple. The entire process can be automatic, requiring no user interaction, or you can include interactions, such as two-factor authentication confirmation.


// A view controller that can upgrade user passwords to strong passwords, or convert accounts to use Sign in with Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationViewController
type AccountAuthenticationModificationViewController struct {
	appkit.ViewController
}

// AccountAuthenticationModificationViewControllerFrom constructs a [AccountAuthenticationModificationViewController] from an unsafe.Pointer.
//
// A view controller that can upgrade user passwords to strong passwords, or convert accounts to use Sign in with Apple.
func AccountAuthenticationModificationViewControllerFrom(ptr unsafe.Pointer) AccountAuthenticationModificationViewController {
	return AccountAuthenticationModificationViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AccountAuthenticationModificationViewControllerClass) Alloc() AccountAuthenticationModificationViewController {
	rv := objc.Send[AccountAuthenticationModificationViewController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccountAuthenticationModificationViewControllerClass) New() AccountAuthenticationModificationViewController {
	rv := objc.Send[AccountAuthenticationModificationViewController](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccountAuthenticationModificationViewController) Init() AccountAuthenticationModificationViewController {
	rv := objc.Send[AccountAuthenticationModificationViewController](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccountAuthenticationModificationViewController) Autorelease() AccountAuthenticationModificationViewController {
	rv := objc.Send[AccountAuthenticationModificationViewController](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccountAuthenticationModificationViewController creates a new AccountAuthenticationModificationViewController instance.
func NewAccountAuthenticationModificationViewController() AccountAuthenticationModificationViewController {
	return getAccountAuthenticationModificationViewControllerClass().New()
}



// Prepares the view controller’s interface that displays when converting an account that uses password authentication to use Sign in with Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationViewController/prepareInterfaceToConvertAccountToSignInWithApple(for:existingCredential:userInfo:)
func (a_ AccountAuthenticationModificationViewController) PrepareInterfaceToConvertAccountToSignInWithAppleForServiceIdentifierExistingCredentialUserInfo(serviceIdentifier CredentialServiceIdentifier /* not a class type */, existingCredential IASPasswordCredential, userInfo objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("prepareInterfaceToConvertAccountToSignInWithAppleForServiceIdentifier:existingCredential:userInfo:"), serviceIdentifier, existingCredential, userInfo)
}


// The context your account authentication modification extension uses to provide information to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asaccountauthenticationmodificationviewcontroller/extensioncontext
func (a_ AccountAuthenticationModificationViewController) ExtensionContext() IASAccountAuthenticationModificationExtensionContext {
	rv := objc.Send[AccountAuthenticationModificationExtensionContext](a_.ID, objc.Sel("extensionContext"))
	return rv
}


// The context your account authentication modification extension uses to provide information to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asaccountauthenticationmodificationviewcontroller/extensioncontext
func (a_ AccountAuthenticationModificationViewController) SetExtensionContext(value IASAccountAuthenticationModificationExtensionContext) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setExtensionContext:"), value)
}



