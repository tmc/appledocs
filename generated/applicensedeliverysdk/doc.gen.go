// Code generated from Apple documentation for AppLicenseDeliverySDK. DO NOT EDIT.

// Package applicensedeliverysdk provides Go bindings for the AppLicenseDeliverySDK framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AppLicenseDeliverySDK without requiring cgo.
package applicensedeliverysdk

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AppLicenseDeliverySDK.framework/AppLicenseDeliverySDK"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
