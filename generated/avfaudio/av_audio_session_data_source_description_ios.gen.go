//go:build darwin && ios

// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AudioSessionDataSourceDescription


// Selects the preferred directivity configuration for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/setPreferredPolarPattern(_:)
func (a_ AudioSessionDataSourceDescription) SetPreferredPolarPatternError(pattern AudioSessionPolarPattern /* typedef */, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredPolarPattern:error:"), pattern, outError)
	return rv
}

// iOS-only properties

// The system-assigned identifier for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/dataSourceID
func (a_ AudioSessionDataSourceDescription) DataSourceID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](a_.ID, objc.Sel("dataSourceID"))
	return rv
}

// A human-readable name for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/dataSourceName
func (a_ AudioSessionDataSourceDescription) DataSourceName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("dataSourceName"))
	return rv
}

// The location of the data source on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/location
func (a_ AudioSessionDataSourceDescription) Location() AudioSessionLocation /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("location"))
	return rv
}

// The orientation of the data source relative to the device’s natural orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/orientation
func (a_ AudioSessionDataSourceDescription) Orientation() AudioSessionOrientation /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("orientation"))
	return rv
}

// The preferred directivity configuration for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/preferredPolarPattern
func (a_ AudioSessionDataSourceDescription) PreferredPolarPattern() AudioSessionPolarPattern /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("preferredPolarPattern"))
	return rv
}

// The data source’s active polar pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/selectedPolarPattern
func (a_ AudioSessionDataSourceDescription) SelectedPolarPattern() AudioSessionPolarPattern /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("selectedPolarPattern"))
	return rv
}

// The set of directivity configurations supported by the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/supportedPolarPatterns
func (a_ AudioSessionDataSourceDescription) SupportedPolarPatterns() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("supportedPolarPatterns"))
	return rv
}





