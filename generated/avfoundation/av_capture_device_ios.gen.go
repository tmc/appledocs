//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/coremedia"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureDevice


// iOS-only properties

// An array of physical devices that make up a virtual device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/constituentDevices
func (c_ CaptureDevice) ConstituentDevices() []ICaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("constituentDevices"))
	return rv
}

// A key-value observable property indicating the current aspect ratio for a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/dynamicAspectRatio
func (c_ CaptureDevice) DynamicAspectRatio() CaptureAspectRatio /* not a class type */ {
	rv := objc.Send[CaptureAspectRatio](c_.ID, objc.Sel("dynamicAspectRatio"))
	return rv
}

// A Boolean value that indicates whether the device monitors the subject area for changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSubjectAreaChangeMonitoringEnabled
func (c_ CaptureDevice) SubjectAreaChangeMonitoringEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("subjectAreaChangeMonitoringEnabled"))
	return rv
}
func (c_ CaptureDevice) SetSubjectAreaChangeMonitoringEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setSubjectAreaChangeMonitoringEnabled:"), value)
}

// The nominal 35mm equivalent focal length of the capture device’s lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/nominalFocalLengthIn35mmFilm
func (c_ CaptureDevice) NominalFocalLengthIn35mmFilm() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("nominalFocalLengthIn35mmFilm"))
	return rv
}

// A value that indicates the capture device’s current system pressure state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/systemPressureState-swift.property
func (c_ CaptureDevice) SystemPressureState() objc.IObject /* cross-framework: CaptureSystemPressureState */ {
	rv := objc.Send[CaptureSystemPressureState](c_.ID, objc.Sel("systemPressureState"))
	return rv
}




