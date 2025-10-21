// Code generated from Apple documentation for ApplePayMerchantTokenUsageInformation. DO NOT EDIT.

// Package applepaymerchanttokenusageinformation provides Go bindings for the ApplePayMerchantTokenUsageInformation framework.
//
// Add details about your merchant token usage information package. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ApplePayMerchantTokenUsageInformation without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/ApplePayMerchantTokenUsageInformation
package applepaymerchanttokenusageinformation

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ApplePayMerchantTokenUsageInformation.framework/ApplePayMerchantTokenUsageInformation"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


