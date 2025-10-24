//go:build darwin && ios

// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Preferences


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/isLookToScrollEnabled
func (p_ Preferences) IsLookToScrollEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isLookToScrollEnabled"))
	return rv
}
func (p_ Preferences) SetIsLookToScrollEnabled(value bool) {
	p_.ID.Send(objc.RegisterName("setIsLookToScrollEnabled:"), value)
}





