//go:build darwin && ios

// Code generated from Apple documentation for IdentityLookupUI. DO NOT EDIT.

package identitylookupui

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// iOS-only methods for ILClassificationUIExtensionViewController


// Notifies the view controller when the user finishes entering data and presses the Done button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookupUI/ILClassificationUIExtensionViewController/classificationResponse(for:)
func (i_ ILClassificationUIExtensionViewController) ClassificationResponseForRequest(request unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("classificationResponseForRequest:"), request)
	return rv
}

// Notifies the view controller just before the system presents it to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookupUI/ILClassificationUIExtensionViewController/prepare(for:)
func (i_ ILClassificationUIExtensionViewController) PrepareForClassificationRequest(request unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("prepareForClassificationRequest:"), request)
}

// iOS-only properties

// The context for the current request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookupUI/ILClassificationUIExtensionViewController/extensionContext
func (i_ ILClassificationUIExtensionViewController) ExtensionContext() ILClassificationUIExtensionContext {
	rv := objc.Send[ILClassificationUIExtensionContext](i_.ID, objc.Sel("extensionContext"))
	return rv
}






