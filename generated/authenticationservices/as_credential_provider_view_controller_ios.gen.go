//go:build darwin && ios

// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CredentialProviderViewController


// Prepare the view controller to show a list of all insertable text with user selectable fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderViewController/prepareInterfaceForUserChoosingTextToInsert()
func (c_ CredentialProviderViewController) PrepareInterfaceForUserChoosingTextToInsert() {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareInterfaceForUserChoosingTextToInsert"))
}

// iOS-only properties





