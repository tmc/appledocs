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
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/3650393-requestcontentsauthorizationwith
func (i_ ICDeviceBrowser) RequestContentsAuthorizationWithCompletion(completion ICAuthorizationStatus) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestContentsAuthorizationWithCompletion:"), completion)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/3650394-requestcontrolauthorizationwithc
func (i_ ICDeviceBrowser) RequestControlAuthorizationWithCompletion(completion ICAuthorizationStatus) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestControlAuthorizationWithCompletion:"), completion)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/3778550-resetcontentsauthorizationwithco
func (i_ ICDeviceBrowser) ResetContentsAuthorizationWithCompletion(completion ICAuthorizationStatus) {
	objc.Send[objc.ID](i_.ID, objc.Sel("resetContentsAuthorizationWithCompletion:"), completion)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/3778551-resetcontrolauthorizationwithcom
func (i_ ICDeviceBrowser) ResetControlAuthorizationWithCompletion(completion ICAuthorizationStatus) {
	objc.Send[objc.ID](i_.ID, objc.Sel("resetControlAuthorizationWithCompletion:"), completion)
}

// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/3650391-contentsauthorizationstatus
func (i_ ICDeviceBrowser) ContentsAuthorizationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("contentsAuthorizationStatus"))
	return rv
}
func (i_ ICDeviceBrowser) SetContentsAuthorizationStatus(value unsafe.Pointer) {
	i_.ID.Send(objc.RegisterName("setContentsAuthorizationStatus:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/3650392-controlauthorizationstatus
func (i_ ICDeviceBrowser) ControlAuthorizationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("controlAuthorizationStatus"))
	return rv
}
func (i_ ICDeviceBrowser) SetControlAuthorizationStatus(value unsafe.Pointer) {
	i_.ID.Send(objc.RegisterName("setControlAuthorizationStatus:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/3650395-issuspended
func (i_ ICDeviceBrowser) IsSuspended() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("isSuspended"))
	return rv
}
func (i_ ICDeviceBrowser) SetIsSuspended(value unsafe.Pointer) {
	i_.ID.Send(objc.RegisterName("setIsSuspended:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/isSuspended
func (i_ ICDeviceBrowser) Suspended() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("suspended"))
	return rv
}




