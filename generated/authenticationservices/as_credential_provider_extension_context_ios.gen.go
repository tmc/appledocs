//go:build darwin && ios

// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CredentialProviderExtensionContext


// Provides the user-selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderExtensionContext/completeRequest(withTextToInsert:completionHandler:)
func (c_ CredentialProviderExtensionContext) CompleteRequestWithTextToInsertCompletionHandler(text objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("completeRequestWithTextToInsert:completionHandler:"), text, completionHandler)
}

// iOS-only properties





