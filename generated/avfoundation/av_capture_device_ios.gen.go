//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureDevice


// Smoothly ends a zoom transition in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/cancelVideoZoomRamp()
func (c_ CaptureDevice) CancelVideoZoomRamp() {
	objc.Send[objc.ID](c_.ID, objc.Sel("cancelVideoZoomRamp"))
}

// Converts device-specific white balance RGB gain values to device-independent chromaticity values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/chromaticityValues(for:)
func (c_ CaptureDevice) ChromaticityValuesForDeviceWhiteBalanceGains(whiteBalanceGains CaptureWhiteBalanceGains /* not a class type */) CaptureWhiteBalanceChromaticityValues /* not a class type */ {
	rv := objc.Send[CaptureWhiteBalanceChromaticityValues](c_.ID, objc.Sel("chromaticityValuesForDeviceWhiteBalanceGains:"), whiteBalanceGains)
	return rv
}

// Converts device-independent temperature and tint values to device-specific white balance RGB gain values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/deviceWhiteBalanceGains(for:)-3wtsa
func (c_ CaptureDevice) DeviceWhiteBalanceGainsForTemperatureAndTintValues(tempAndTintValues CaptureWhiteBalanceTemperatureAndTintValues /* not a class type */) CaptureWhiteBalanceGains /* not a class type */ {
	rv := objc.Send[CaptureWhiteBalanceGains](c_.ID, objc.Sel("deviceWhiteBalanceGainsForTemperatureAndTintValues:"), tempAndTintValues)
	return rv
}

// Converts device-independent chromaticity values to device-specific white balance RGB gain values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/deviceWhiteBalanceGains(for:)-9gdtw
func (c_ CaptureDevice) DeviceWhiteBalanceGainsForChromaticityValues(chromaticityValues CaptureWhiteBalanceChromaticityValues /* not a class type */) CaptureWhiteBalanceGains /* not a class type */ {
	rv := objc.Send[CaptureWhiteBalanceGains](c_.ID, objc.Sel("deviceWhiteBalanceGainsForChromaticityValues:"), chromaticityValues)
	return rv
}

// Begins a smooth transition from the current zoom factor to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/ramp(toVideoZoomFactor:withRate:)
func (c_ CaptureDevice) RampToVideoZoomFactorWithRate(factor float64, rate float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("rampToVideoZoomFactor:withRate:"), factor, rate)
}

// Updates the dynamic aspect ratio of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setDynamicAspectRatio(_:completionHandler:)
func (c_ CaptureDevice) SetDynamicAspectRatioCompletionHandler(dynamicAspectRatio CaptureAspectRatio /* typedef */, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDynamicAspectRatio:completionHandler:"), dynamicAspectRatio, handler)
}

// Sets the exposure mode to a custom state, and locks exposure duration and ISO at explicit values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setExposureModeCustom(duration:iso:completionHandler:)
func (c_ CaptureDevice) SetExposureModeCustomWithDurationISOCompletionHandler(duration objc.IObject /* cross-framework: Time */, ISO float32, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureModeCustomWithDuration:ISO:completionHandler:"), duration, ISO, handler)
}

// Sets the bias to apply to the target exposure value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setExposureTargetBias(_:completionHandler:)
func (c_ CaptureDevice) SetExposureTargetBiasCompletionHandler(bias float32, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTargetBias:completionHandler:"), bias, handler)
}

// Locks the lens position at the specified value, and sets the focus mode to a locked state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setFocusModeLocked(lensPosition:completionHandler:)
func (c_ CaptureDevice) SetFocusModeLockedWithLensPositionCompletionHandler(lensPosition float32, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocusModeLockedWithLensPosition:completionHandler:"), lensPosition, handler)
}

// Sets white balance to locked mode with explicit temperature and tint values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setWhiteBalanceModeLocked(whiteBalanceTemperatureAndTintValues:handler:)
func (c_ CaptureDevice) SetWhiteBalanceModeLockedWithDeviceWhiteBalanceTemperatureAndTintValuesCompletionHandler(whiteBalanceTemperatureAndTintValues CaptureWhiteBalanceTemperatureAndTintValues /* not a class type */, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhiteBalanceModeLockedWithDeviceWhiteBalanceTemperatureAndTintValues:completionHandler:"), whiteBalanceTemperatureAndTintValues, handler)
}

// Sets the white balance to locked mode with the specified white balance gains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setWhiteBalanceModeLocked(with:completionHandler:)
func (c_ CaptureDevice) SetWhiteBalanceModeLockedWithDeviceWhiteBalanceGainsCompletionHandler(whiteBalanceGains CaptureWhiteBalanceGains /* not a class type */, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhiteBalanceModeLockedWithDeviceWhiteBalanceGains:completionHandler:"), whiteBalanceGains, handler)
}

// Converts device-specific white balance RGB gain values to device-independent temperature and tint values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/temperatureAndTintValues(for:)
func (c_ CaptureDevice) TemperatureAndTintValuesForDeviceWhiteBalanceGains(whiteBalanceGains CaptureWhiteBalanceGains /* not a class type */) CaptureWhiteBalanceTemperatureAndTintValues /* not a class type */ {
	rv := objc.Send[CaptureWhiteBalanceTemperatureAndTintValues](c_.ID, objc.Sel("temperatureAndTintValuesForDeviceWhiteBalanceGains:"), whiteBalanceGains)
	return rv
}

// iOS-only properties

// The currently active depth data format of the capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeDepthDataFormat
func (c_ CaptureDevice) ActiveDepthDataFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeDepthDataFormat"))
	return rv
}
func (c_ CaptureDevice) SetActiveDepthDataFormat(value IAVCaptureDeviceFormat) {
	c_.ID.Send(objc.RegisterName("setActiveDepthDataFormat:"), value)
}

// The minimum frame duration of depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeDepthDataMinFrameDuration
func (c_ CaptureDevice) ActiveDepthDataMinFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("activeDepthDataMinFrameDuration"))
	return rv
}
func (c_ CaptureDevice) SetActiveDepthDataMinFrameDuration(value objc.IObject /* cross-framework: Time */) {
	c_.ID.Send(objc.RegisterName("setActiveDepthDataMinFrameDuration:"), value)
}

// The maximum exposure duration, in seconds, defined in the autoexposure algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeMaxExposureDuration
func (c_ CaptureDevice) ActiveMaxExposureDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("activeMaxExposureDuration"))
	return rv
}
func (c_ CaptureDevice) SetActiveMaxExposureDuration(value objc.IObject /* cross-framework: Time */) {
	c_.ID.Send(objc.RegisterName("setActiveMaxExposureDuration:"), value)
}

// A value that controls the allowable range for automatic focusing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/autoFocusRangeRestriction-swift.property
func (c_ CaptureDevice) AutoFocusRangeRestriction() CaptureAutoFocusRangeRestriction {
	rv := objc.Send[CaptureAutoFocusRangeRestriction](c_.ID, objc.Sel("autoFocusRangeRestriction"))
	return rv
}
func (c_ CaptureDevice) SetAutoFocusRangeRestriction(value CaptureAutoFocusRangeRestriction) {
	c_.ID.Send(objc.RegisterName("setAutoFocusRangeRestriction:"), value)
}

// A Boolean value that indicates whether the device automatically adjusts face-driven autoexposure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/automaticallyAdjustsFaceDrivenAutoExposureEnabled
func (c_ CaptureDevice) AutomaticallyAdjustsFaceDrivenAutoExposureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyAdjustsFaceDrivenAutoExposureEnabled"))
	return rv
}
func (c_ CaptureDevice) SetAutomaticallyAdjustsFaceDrivenAutoExposureEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setAutomaticallyAdjustsFaceDrivenAutoExposureEnabled:"), value)
}

// A Boolean value that indicates whether the device automatically adjusts face-driven autofocus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/automaticallyAdjustsFaceDrivenAutoFocusEnabled
func (c_ CaptureDevice) AutomaticallyAdjustsFaceDrivenAutoFocusEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyAdjustsFaceDrivenAutoFocusEnabled"))
	return rv
}
func (c_ CaptureDevice) SetAutomaticallyAdjustsFaceDrivenAutoFocusEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setAutomaticallyAdjustsFaceDrivenAutoFocusEnabled:"), value)
}

// A Boolean value that indicates whether the device automatically manages the state of high dynamic range (HDR) video streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/automaticallyAdjustsVideoHDREnabled
func (c_ CaptureDevice) AutomaticallyAdjustsVideoHDREnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyAdjustsVideoHDREnabled"))
	return rv
}
func (c_ CaptureDevice) SetAutomaticallyAdjustsVideoHDREnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setAutomaticallyAdjustsVideoHDREnabled:"), value)
}

// A Boolean value that indicates whether the capture device automatically switches to low-light boost mode when necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/automaticallyEnablesLowLightBoostWhenAvailable
func (c_ CaptureDevice) AutomaticallyEnablesLowLightBoostWhenAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyEnablesLowLightBoostWhenAvailable"))
	return rv
}
func (c_ CaptureDevice) SetAutomaticallyEnablesLowLightBoostWhenAvailable(value bool) {
	c_.ID.Send(objc.RegisterName("setAutomaticallyEnablesLowLightBoostWhenAvailable:"), value)
}

// An array of physical devices that make up a virtual device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/constituentDevices
func (c_ CaptureDevice) ConstituentDevices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("constituentDevices"))
	return rv
}

// The current device-specific RGB white balance gain values in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/deviceWhiteBalanceGains
func (c_ CaptureDevice) DeviceWhiteBalanceGains() CaptureWhiteBalanceGains /* not a class type */ {
	rv := objc.Send[CaptureWhiteBalanceGains](c_.ID, objc.Sel("deviceWhiteBalanceGains"))
	return rv
}

// The video zoom factor at which a dual camera device can automatically switch between cameras.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/dualCameraSwitchOverVideoZoomFactor
func (c_ CaptureDevice) DualCameraSwitchOverVideoZoomFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("dualCameraSwitchOverVideoZoomFactor"))
	return rv
}

// A key-value observable property indicating the current aspect ratio for a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/dynamicAspectRatio
func (c_ CaptureDevice) DynamicAspectRatio() CaptureAspectRatio /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("dynamicAspectRatio"))
	return rv
}

// A key-value observable property describing the output dimensions of the video buffer based on the device’s dynamic aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/dynamicDimensions
func (c_ CaptureDevice) DynamicDimensions() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("dynamicDimensions"))
	return rv
}

// The length of time over which exposure takes place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureDuration
func (c_ CaptureDevice) ExposureDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("exposureDuration"))
	return rv
}

// The bias to apply to the target exposure value, in exposure value (EV) units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureTargetBias
func (c_ CaptureDevice) ExposureTargetBias() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("exposureTargetBias"))
	return rv
}

// The metered exposure level’s offset from the target exposure value, in exposure value (EV) units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureTargetOffset
func (c_ CaptureDevice) ExposureTargetOffset() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("exposureTargetOffset"))
	return rv
}

// The current device-specific white balance values required for a neutral gray white point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/grayWorldDeviceWhiteBalanceGains
func (c_ CaptureDevice) GrayWorldDeviceWhiteBalanceGains() CaptureWhiteBalanceGains /* not a class type */ {
	rv := objc.Send[CaptureWhiteBalanceGains](c_.ID, objc.Sel("grayWorldDeviceWhiteBalanceGains"))
	return rv
}

// A Boolean value that indicates whether the device supports focus range restrictions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAutoFocusRangeRestrictionSupported
func (c_ CaptureDevice) AutoFocusRangeRestrictionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoFocusRangeRestrictionSupported"))
	return rv
}

// A Boolean value that indicates whether the device has face-driven autoexposure enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFaceDrivenAutoExposureEnabled
func (c_ CaptureDevice) FaceDrivenAutoExposureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("faceDrivenAutoExposureEnabled"))
	return rv
}
func (c_ CaptureDevice) SetFaceDrivenAutoExposureEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setFaceDrivenAutoExposureEnabled:"), value)
}

// A Boolean value that indicates whether the device has face-driven autofocus enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFaceDrivenAutoFocusEnabled
func (c_ CaptureDevice) FaceDrivenAutoFocusEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("faceDrivenAutoFocusEnabled"))
	return rv
}
func (c_ CaptureDevice) SetFaceDrivenAutoFocusEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setFaceDrivenAutoFocusEnabled:"), value)
}

// A Boolean value that indicates whether the flash is currently active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFlashActive
func (c_ CaptureDevice) FlashActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("flashActive"))
	return rv
}

// A Boolean value that indicates whether geometric distortion correction is enabled for this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isGeometricDistortionCorrectionEnabled
func (c_ CaptureDevice) GeometricDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("geometricDistortionCorrectionEnabled"))
	return rv
}
func (c_ CaptureDevice) SetGeometricDistortionCorrectionEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setGeometricDistortionCorrectionEnabled:"), value)
}

// A Boolean value that indicates whether this device supports geometric distortion correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isGeometricDistortionCorrectionSupported
func (c_ CaptureDevice) GeometricDistortionCorrectionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("geometricDistortionCorrectionSupported"))
	return rv
}

// A Boolean value that indicates whether the device should use global tone mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isGlobalToneMappingEnabled
func (c_ CaptureDevice) GlobalToneMappingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("globalToneMappingEnabled"))
	return rv
}
func (c_ CaptureDevice) SetGlobalToneMappingEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setGlobalToneMappingEnabled:"), value)
}

// A Boolean value that indicates whether the device supports locking focus to a specific lens position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isLockingFocusWithCustomLensPositionSupported
func (c_ CaptureDevice) LockingFocusWithCustomLensPositionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lockingFocusWithCustomLensPositionSupported"))
	return rv
}

// A Boolean value that indicates whether the device supports locking white balance to specific gain values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isLockingWhiteBalanceWithCustomDeviceGainsSupported
func (c_ CaptureDevice) LockingWhiteBalanceWithCustomDeviceGainsSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lockingWhiteBalanceWithCustomDeviceGainsSupported"))
	return rv
}

// A Boolean value that indicates whether the capture device’s low light boost feature is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isLowLightBoostEnabled
func (c_ CaptureDevice) LowLightBoostEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lowLightBoostEnabled"))
	return rv
}

// A Boolean value that indicates whether the capture device supports boosting images in low-light conditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isLowLightBoostSupported
func (c_ CaptureDevice) LowLightBoostSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lowLightBoostSupported"))
	return rv
}

// The current exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/iso
func (c_ CaptureDevice) ISO() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("ISO"))
	return rv
}

// A Boolean value that indicates whether a zoom transition is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isRampingVideoZoom
func (c_ CaptureDevice) RampingVideoZoom() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rampingVideoZoom"))
	return rv
}

// A Boolean value that indicates whether smooth autofocus is in an enabled state on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSmoothAutoFocusEnabled
func (c_ CaptureDevice) SmoothAutoFocusEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("smoothAutoFocusEnabled"))
	return rv
}
func (c_ CaptureDevice) SetSmoothAutoFocusEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setSmoothAutoFocusEnabled:"), value)
}

// A Boolean value that indicates whether the device supports smooth autofocus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSmoothAutoFocusSupported
func (c_ CaptureDevice) SmoothAutoFocusSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("smoothAutoFocusSupported"))
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

// A Boolean value that indicates whether the device streams high dynamic range video buffers, also known as extended dynamic range (EDR).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isVideoHDREnabled
func (c_ CaptureDevice) VideoHDREnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("videoHDREnabled"))
	return rv
}
func (c_ CaptureDevice) SetVideoHDREnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setVideoHDREnabled:"), value)
}

// A Boolean value that indicates whether the device consists of two or more physical devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isVirtualDevice
func (c_ CaptureDevice) VirtualDevice() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("virtualDevice"))
	return rv
}

// The size of the lens diaphragm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/lensAperture
func (c_ CaptureDevice) LensAperture() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("lensAperture"))
	return rv
}

// The current focus position of the lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/lensPosition
func (c_ CaptureDevice) LensPosition() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("lensPosition"))
	return rv
}

// The maximum zoom factor allowed in the current capture configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/maxAvailableVideoZoomFactor
func (c_ CaptureDevice) MaxAvailableVideoZoomFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("maxAvailableVideoZoomFactor"))
	return rv
}

// The maximum supported exposure bias, in exposure value (EV) units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/maxExposureTargetBias
func (c_ CaptureDevice) MaxExposureTargetBias() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("maxExposureTargetBias"))
	return rv
}

// The maximum supported value to which you can set a color channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/maxWhiteBalanceGain
func (c_ CaptureDevice) MaxWhiteBalanceGain() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("maxWhiteBalanceGain"))
	return rv
}

// The minimum zoom factor allowed in the current capture configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minAvailableVideoZoomFactor
func (c_ CaptureDevice) MinAvailableVideoZoomFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("minAvailableVideoZoomFactor"))
	return rv
}

// The minimum supported exposure bias, in exposure value (EV) units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minExposureTargetBias
func (c_ CaptureDevice) MinExposureTargetBias() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("minExposureTargetBias"))
	return rv
}

// The nominal 35mm equivalent focal length of the capture device’s lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/nominalFocalLengthIn35mmFilm
func (c_ CaptureDevice) NominalFocalLengthIn35mmFilm() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("nominalFocalLengthIn35mmFilm"))
	return rv
}

// A monitor owned by the device that recommends an optimal framing based on the content in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/smartFramingMonitor
func (c_ CaptureDevice) SmartFramingMonitor() IAVCaptureSmartFramingMonitor {
	rv := objc.Send[CaptureSmartFramingMonitor](c_.ID, objc.Sel("smartFramingMonitor"))
	return rv
}

// A value that indicates the capture device’s current system pressure state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/systemPressureState-swift.property
func (c_ CaptureDevice) SystemPressureState() IAVCaptureSystemPressureState {
	rv := objc.Send[CaptureSystemPressureState](c_.ID, objc.Sel("systemPressureState"))
	return rv
}

// A value that controls the cropping and enlargement of images captured by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/videoZoomFactor
func (c_ CaptureDevice) VideoZoomFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoZoomFactor"))
	return rv
}
func (c_ CaptureDevice) SetVideoZoomFactor(value float64) {
	c_.ID.Send(objc.RegisterName("setVideoZoomFactor:"), value)
}

// An array of video zoom factors at or above which a virtual device, such as the dual camera, may switch to its next constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/virtualDeviceSwitchOverVideoZoomFactors
func (c_ CaptureDevice) VirtualDeviceSwitchOverVideoZoomFactors() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("virtualDeviceSwitchOverVideoZoomFactors"))
	return rv
}




