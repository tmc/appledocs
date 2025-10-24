// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

/* debug [enums.gen.go]: Generating 7 enums for Accessibility */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum AXSettingsFeature (5 cases) */
// AXSettingsFeature - Constants that describe specific Accessibility settings in the Settings app.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/Feature
type AXSettingsFeature uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/Feature/allowAppsToAddAudioToCalls
	AXSettingsFeatureAllowAppsToAddAudioToCalls AXSettingsFeature = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/Feature/assistiveTouch
	AXSettingsFeatureAssistiveTouch AXSettingsFeature = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/Feature/assistiveTouchDevices
	AXSettingsFeatureAssistiveTouchDevices AXSettingsFeature = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/Feature/dwellControl
	AXSettingsFeatureDwellControl AXSettingsFeature = 0
	// AXSettingsFeaturePersonalVoiceAllowAppsToRequestToUse - A constant for opening the Settings app to the setting for Personal Voice > Allow Apps to Request to Use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/Feature/personalVoiceAllowAppsToRequestToUse
	AXSettingsFeaturePersonalVoiceAllowAppsToRequestToUse AXSettingsFeature = 0
)

/* debug [enums.gen.go]: Processing enum AXChartDescriptorContentDirection (6 cases) */
// AXChartDescriptorContentDirection - A constant that describes the content direction of the chart.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/ContentDirection-swift.enum
type AXChartDescriptorContentDirection uint

const (
	// AXChartContentDirectionBottomToTop - A content direction with an x-axis that increases from bottom to top.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/ContentDirection-swift.enum/bottomToTop
	AXChartContentDirectionBottomToTop AXChartDescriptorContentDirection = 0
	// AXChartContentDirectionLeftToRight - A content direction with an x-axis that increases from left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/ContentDirection-swift.enum/leftToRight
	AXChartContentDirectionLeftToRight AXChartDescriptorContentDirection = 0
	// AXChartContentDirectionRadialClockwise - A content direction with a radial x-axis that increases clockwise.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/ContentDirection-swift.enum/radialClockwise
	AXChartContentDirectionRadialClockwise AXChartDescriptorContentDirection = 0
	// AXChartContentDirectionRadialCounterClockwise - A content direction with a radial x-axis that increases counterclockwise.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/ContentDirection-swift.enum/radialCounterClockwise
	AXChartContentDirectionRadialCounterClockwise AXChartDescriptorContentDirection = 0
	// AXChartContentDirectionRightToLeft - A content direction with an x-axis that increases from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/ContentDirection-swift.enum/rightToLeft
	AXChartContentDirectionRightToLeft AXChartDescriptorContentDirection = 0
	// AXChartContentDirectionTopToBottom - A content direction with an x-axis that increases from top to bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/ContentDirection-swift.enum/topToBottom
	AXChartContentDirectionTopToBottom AXChartDescriptorContentDirection = 0
)

/* debug [enums.gen.go]: Processing enum AXCustomContentImportance (2 cases) */
// AXCustomContentImportance - Objects that control the timing of content output.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/Importance-swift.enum
type AXCustomContentImportance uint

const (
	// AXCustomContentImportanceDefault - Output the content to the user on demand.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/Importance-swift.enum/default
	AXCustomContentImportanceDefault AXCustomContentImportance = 0
	// AXCustomContentImportanceHigh - Output the content to the user immediately.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/Importance-swift.enum/high
	AXCustomContentImportanceHigh AXCustomContentImportance = 0
)

/* debug [enums.gen.go]: Processing enum AXFeatureOverrideSessionOptions (5 cases) */
// AXFeatureOverrideSessionOptions - Options indicating which Accessibility features will be turned on or off when an override session is held by your app.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSession/Options
type AXFeatureOverrideSessionOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSession/Options/grayscale
	AXFeatureOverrideSessionOptionsGrayscale AXFeatureOverrideSessionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSession/Options/invertColors
	AXFeatureOverrideSessionOptionsInvertColors AXFeatureOverrideSessionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSession/Options/voiceControl
	AXFeatureOverrideSessionOptionsVoiceControl AXFeatureOverrideSessionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSession/Options/voiceOver
	AXFeatureOverrideSessionOptionsVoiceOver AXFeatureOverrideSessionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSession/Options/zoom
	AXFeatureOverrideSessionOptionsZoom AXFeatureOverrideSessionOptions = 0
)

/* debug [enums.gen.go]: Processing enum AXFeatureOverrideSessionError (4 cases) */
// AXFeatureOverrideSessionError enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionError-swift.struct/Code
type AXFeatureOverrideSessionError uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionError-swift.struct/Code/appNotEntitled
	AXFeatureOverrideSessionErrorAppNotEntitled AXFeatureOverrideSessionError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionError-swift.struct/Code/overrideIsAlreadyActive
	AXFeatureOverrideSessionErrorOverrideIsAlreadyActive AXFeatureOverrideSessionError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionError-swift.struct/Code/overrideNotFoundForUUID
	AXFeatureOverrideSessionErrorOverrideNotFoundForUUID AXFeatureOverrideSessionError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionError-swift.struct/Code/undefined
	AXFeatureOverrideSessionErrorUndefined AXFeatureOverrideSessionError = 0
)

/* debug [enums.gen.go]: Processing enum AXHearingDeviceEar (4 cases) */
// AXHearingDeviceEar - Constants that represent a hearing device ear.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMFiHearingDevice/Ear
type AXHearingDeviceEar uint

const (
	// AXHearingDeviceEarNone - A constant that represents neither ear.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXHearingDeviceEar/AXHearingDeviceEarNone
	AXHearingDeviceEarNone AXHearingDeviceEar = 0
	// AXHearingDeviceEarBoth - A constant that represents both ears.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMFiHearingDevice/Ear/both
	AXHearingDeviceEarBoth AXHearingDeviceEar = 0
	// AXHearingDeviceEarLeft - A constant that represents the left ear.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMFiHearingDevice/Ear/left
	AXHearingDeviceEarLeft AXHearingDeviceEar = 0
	// AXHearingDeviceEarRight - A constant that represents the right ear.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMFiHearingDevice/Ear/right
	AXHearingDeviceEarRight AXHearingDeviceEar = 0
)

/* debug [enums.gen.go]: Processing enum AXNumericDataAxisDescriptorScale (3 cases) */
// AXNumericDataAxisDescriptorScale - Constants that describe the scale of a numeric axis.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/ScaleType-swift.enum
type AXNumericDataAxisDescriptorScale uint

const (
	// AXScaleTypeLinear - A linear scale.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/ScaleType-swift.enum/linear
	AXScaleTypeLinear AXNumericDataAxisDescriptorScale = 0
	// AXScaleTypeLn - A natural log scale.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/ScaleType-swift.enum/ln
	AXScaleTypeLn AXNumericDataAxisDescriptorScale = 0
	// AXScaleTypeLog10 - A log scale.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/ScaleType-swift.enum/log10
	AXScaleTypeLog10 AXNumericDataAxisDescriptorScale = 0
)


