//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureSmartFramingMonitor


// Begins monitoring the device’s active scene and making framing recommendations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSmartFramingMonitor/startMonitoring()
func (c_ CaptureSmartFramingMonitor) StartMonitoringWithError(outError objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("startMonitoringWithError:"), outError)
	return rv
}

// Stops monitoring the device’s active scene and making framing recommendations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSmartFramingMonitor/stopMonitoring()
func (c_ CaptureSmartFramingMonitor) StopMonitoring() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopMonitoring"))
}

// iOS-only properties

// An array of framings that the monitor is allowed to suggest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSmartFramingMonitor/enabledFramings
func (c_ CaptureSmartFramingMonitor) EnabledFramings() []CaptureFraming {
	rv := objc.Send[[]CaptureFraming](c_.ID, objc.Sel("enabledFramings"))
	return rv
}
func (c_ CaptureSmartFramingMonitor) SetEnabledFramings(value []CaptureFraming) {
	c_.ID.Send(objc.RegisterName("setEnabledFramings:"), value)
}

// Yes when the receiver is actively monitoring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSmartFramingMonitor/isMonitoring
func (c_ CaptureSmartFramingMonitor) Monitoring() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("monitoring"))
	return rv
}

// The latest recommended framing from the monitor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSmartFramingMonitor/recommendedFraming
func (c_ CaptureSmartFramingMonitor) RecommendedFraming() IAVCaptureFraming {
	rv := objc.Send[CaptureFraming](c_.ID, objc.Sel("recommendedFraming"))
	return rv
}

// An array of framings supported by the monitor in its current configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSmartFramingMonitor/supportedFramings
func (c_ CaptureSmartFramingMonitor) SupportedFramings() []CaptureFraming {
	rv := objc.Send[[]CaptureFraming](c_.ID, objc.Sel("supportedFramings"))
	return rv
}





