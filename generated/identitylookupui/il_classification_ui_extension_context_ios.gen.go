//go:build darwin && ios

// Code generated from Apple documentation for IdentityLookupUI. DO NOT EDIT.

package identitylookupui

import (
	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for ILClassificationUIExtensionContext

// iOS-only properties

// A Boolean value that determines whether the extension has enough information to complete the report.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookupUI/ILClassificationUIExtensionContext/isReadyForClassificationResponse
func (i_ ILClassificationUIExtensionContext) ReadyForClassificationResponse() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("readyForClassificationResponse"))
	return rv
}
func (i_ ILClassificationUIExtensionContext) SetReadyForClassificationResponse(value bool) {
	i_.ID.Send(objc.RegisterName("setReadyForClassificationResponse:"), value)
}
