//go:build darwin && ios

// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for AccountAuthenticationModificationViewController


// Cancels a request that the user initiated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationViewController/cancelRequest()
func (a_ AccountAuthenticationModificationViewController) CancelRequest() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelRequest"))
}

// Upgrades a user’s weak password to a strong password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationViewController/changePasswordWithoutUserInteraction(for:existingCredential:newPassword:userInfo:)
func (a_ AccountAuthenticationModificationViewController) ChangePasswordWithoutUserInteractionForServiceIdentifierExistingCredentialNewPasswordUserInfo(serviceIdentifier IASCredentialServiceIdentifier, existingCredential IASPasswordCredential, newPassword objc.IObject /* cross-framework: NSString */, userInfo objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("changePasswordWithoutUserInteractionForServiceIdentifier:existingCredential:newPassword:userInfo:"), serviceIdentifier, existingCredential, newPassword, userInfo)
}

// Converts an account’s authentication mechanism from using passwords to using Sign in with Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationViewController/convertAccountToSignInWithAppleWithoutUserInteraction(for:existingCredential:userInfo:)
func (a_ AccountAuthenticationModificationViewController) ConvertAccountToSignInWithAppleWithoutUserInteractionForServiceIdentifierExistingCredentialUserInfo(serviceIdentifier IASCredentialServiceIdentifier, existingCredential IASPasswordCredential, userInfo objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("convertAccountToSignInWithAppleWithoutUserInteractionForServiceIdentifier:existingCredential:userInfo:"), serviceIdentifier, existingCredential, userInfo)
}

// Prepares the view controller’s interface that displays when upgrading from a weak password to a strong password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationViewController/prepareInterfaceToChangePassword(for:existingCredential:newPassword:userInfo:)
func (a_ AccountAuthenticationModificationViewController) PrepareInterfaceToChangePasswordForServiceIdentifierExistingCredentialNewPasswordUserInfo(serviceIdentifier IASCredentialServiceIdentifier, existingCredential IASPasswordCredential, newPassword objc.IObject /* cross-framework: NSString */, userInfo objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("prepareInterfaceToChangePasswordForServiceIdentifier:existingCredential:newPassword:userInfo:"), serviceIdentifier, existingCredential, newPassword, userInfo)
}

// Prepares the view controller’s interface that displays when converting an account that uses password authentication to use Sign in with Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationViewController/prepareInterfaceToConvertAccountToSignInWithApple(for:existingCredential:userInfo:)
func (a_ AccountAuthenticationModificationViewController) PrepareInterfaceToConvertAccountToSignInWithAppleForServiceIdentifierExistingCredentialUserInfo(serviceIdentifier IASCredentialServiceIdentifier, existingCredential IASPasswordCredential, userInfo objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("prepareInterfaceToConvertAccountToSignInWithAppleForServiceIdentifier:existingCredential:userInfo:"), serviceIdentifier, existingCredential, userInfo)
}

// iOS-only properties

// The context your account authentication modification extension uses to provide information to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationViewController/extensionContext
func (a_ AccountAuthenticationModificationViewController) ExtensionContext() IASAccountAuthenticationModificationExtensionContext {
	rv := objc.Send[AccountAuthenticationModificationExtensionContext](a_.ID, objc.Sel("extensionContext"))
	return rv
}





