//go:build darwin && ios

// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for AccountAuthenticationModificationExtensionContext


// Cancels a request with an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationExtensionContext/cancelRequest(withError:)
func (a_ AccountAuthenticationModificationExtensionContext) CancelRequestWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelRequestWithError:"), error_)
}

// Completes a request to update an account’s authentication credentials from using a weak password to using a strong password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationExtensionContext/completeChangePasswordRequest(updatedCredential:userInfo:)
func (a_ AccountAuthenticationModificationExtensionContext) CompleteChangePasswordRequestWithUpdatedCredentialUserInfo(updatedCredential IASPasswordCredential, userInfo objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("completeChangePasswordRequestWithUpdatedCredential:userInfo:"), updatedCredential, userInfo)
}

// Completes the process of upgrading an account’s authentication credentials from using passwords to using Sign in with Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationExtensionContext/completeUpgradeToSignInWithApple(userInfo:)
func (a_ AccountAuthenticationModificationExtensionContext) CompleteUpgradeToSignInWithAppleWithUserInfo(userInfo objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("completeUpgradeToSignInWithAppleWithUserInfo:"), userInfo)
}

// Retrieves the user’s current Sign in with Apple authorization credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationExtensionContext/getSignInWithAppleUpgradeAuthorization(state:nonce:completionHandler:)
func (a_ AccountAuthenticationModificationExtensionContext) GetSignInWithAppleUpgradeAuthorizationWithStateNonceCompletionHandler(state objc.IObject /* cross-framework: NSString */, nonce objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("getSignInWithAppleUpgradeAuthorizationWithState:nonce:completionHandler:"), state, nonce, completionHandler)
}

// iOS-only properties





