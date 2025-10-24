//go:build darwin && ios

// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ICDeviceBrowser


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/requestControlAuthorization(completion:)
func (i_ ICDeviceBrowser) RequestControlAuthorizationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestControlAuthorizationWithCompletion:"), completion)
}

// iOS-only properties





