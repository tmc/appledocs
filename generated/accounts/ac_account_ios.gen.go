//go:build darwin && ios

// Code generated from Apple documentation for Accounts. DO NOT EDIT.

package accounts

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ACAccount


// iOS-only properties

// The full name associated with the user account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccount/userFullName
func (a_ ACAccount) UserFullName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("userFullName"))
	return rv
}




