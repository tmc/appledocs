//go:build darwin && ios

// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AXFeatureOverrideSessionManager


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionManager/beginOverrideSession(enabling:disabling:)
func (a_ AXFeatureOverrideSessionManager) BeginOverrideSessionEnablingOptionsDisablingOptionsError(enableOptions AXFeatureOverrideSessionOptions, disableOptions AXFeatureOverrideSessionOptions, error_ unsafe.Pointer) IAXFeatureOverrideSession {
	rv := objc.Send[AXFeatureOverrideSession](a_.ID, objc.Sel("beginOverrideSessionEnablingOptions:disablingOptions:error:"), enableOptions, disableOptions, error_)
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionManager/end(_:)
func (a_ AXFeatureOverrideSessionManager) EndOverrideSessionError(session IAXFeatureOverrideSession, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("endOverrideSession:error:"), session, error_)
	return rv
}

// iOS-only properties





