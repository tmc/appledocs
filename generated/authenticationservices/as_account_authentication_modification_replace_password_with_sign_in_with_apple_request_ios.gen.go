//go:build darwin && ios

// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest


// iOS-only properties

// An identifier that represents a particular service that the user needs a credential for, like a web site.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest/serviceIdentifier
func (a_ AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest) ServiceIdentifier() IASCredentialServiceIdentifier {
	rv := objc.Send[CredentialServiceIdentifier](a_.ID, objc.Sel("serviceIdentifier"))
	return rv
}

// The user name of the account to upgrade.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest/user
func (a_ AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest) User() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("user"))
	return rv
}

// A dictionary that contains values to pass to your account modification extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest/userInfo
func (a_ AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest) UserInfo() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](a_.ID, objc.Sel("userInfo"))
	return rv
}




