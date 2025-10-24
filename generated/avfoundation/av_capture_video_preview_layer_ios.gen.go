//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureVideoPreviewLayer


// iOS-only properties

// Indicates whether the layer display automatically adjusts mirroring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/automaticallyAdjustsMirroring
func (c_ CaptureVideoPreviewLayer) AutomaticallyAdjustsMirroring() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyAdjustsMirroring"))
	return rv
}
func (c_ CaptureVideoPreviewLayer) SetAutomaticallyAdjustsMirroring(value bool) {
	c_.ID.Send(objc.RegisterName("setAutomaticallyAdjustsMirroring:"), value)
}

// A Boolean value that indicates whether the layer is rendering video frames from its source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/isPreviewing
func (c_ CaptureVideoPreviewLayer) Previewing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("previewing"))
	return rv
}

// Indicates whether the layer display is mirrored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/mirrored
func (c_ CaptureVideoPreviewLayer) Mirrored() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("mirrored"))
	return rv
}
func (c_ CaptureVideoPreviewLayer) SetMirrored(value bool) {
	c_.ID.Send(objc.RegisterName("setMirrored:"), value)
}

// Indicates whether the layer display supports mirroring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/mirroringSupported
func (c_ CaptureVideoPreviewLayer) MirroringSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("mirroringSupported"))
	return rv
}

// The layer’s orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/orientation
func (c_ CaptureVideoPreviewLayer) Orientation() CaptureVideoOrientation {
	rv := objc.Send[CaptureVideoOrientation](c_.ID, objc.Sel("orientation"))
	return rv
}
func (c_ CaptureVideoPreviewLayer) SetOrientation(value CaptureVideoOrientation) {
	c_.ID.Send(objc.RegisterName("setOrientation:"), value)
}

// Indicates whether the layer display supports changing the orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/orientationSupported
func (c_ CaptureVideoPreviewLayer) OrientationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("orientationSupported"))
	return rv
}




