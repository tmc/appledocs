// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

// Enum types and constants
// AXChartDescriptorContentDirection - A constant that describes the content direction of the chart.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/ContentDirection-swift.enum
type AXChartDescriptorContentDirection uint

// AXFeatureOverrideSessionOptions - Options indicating which Accessibility features will be turned on or off when an override session is held by your app.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSession/Options
type AXFeatureOverrideSessionOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSession/Options/grayscale
	AXFeatureOverrideSessionOptionsGrayscale AXFeatureOverrideSessionOptions = 0
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

// AXFeatureOverrideSessionError enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionError-swift.struct/Code
type AXFeatureOverrideSessionError uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionError-swift.struct/Code/appNotEntitled
	AXFeatureOverrideSessionErrorAppNotEntitled AXFeatureOverrideSessionError = 0
)

// AXHearingDeviceEar - Constants that represent a hearing device ear.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMFiHearingDevice/Ear
type AXHearingDeviceEar uint

// AXNumericDataAxisDescriptorScale - Constants that describe the scale of a numeric axis.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/ScaleType-swift.enum
type AXNumericDataAxisDescriptorScale uint

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


