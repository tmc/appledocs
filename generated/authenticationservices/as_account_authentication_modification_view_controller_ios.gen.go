//go:build darwin && ios

// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for AccountAuthenticationModificationViewController


// Prepares the view controller’s interface that displays when converting an account that uses password authentication to use Sign in with Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationViewController/prepareInterfaceToConvertAccountToSignInWithApple(for:existingCredential:userInfo:)
func (a_ AccountAuthenticationModificationViewController) PrepareInterfaceToConvertAccountToSignInWithAppleForServiceIdentifierExistingCredentialUserInfo(serviceIdentifier CredentialServiceIdentifier /* not a class type */, existingCredential IASPasswordCredential, userInfo objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("prepareInterfaceToConvertAccountToSignInWithAppleForServiceIdentifier:existingCredential:userInfo:"), serviceIdentifier, existingCredential, userInfo)
}

// iOS-only properties





