//go:build darwin && ios

// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AudioSessionDataSourceDescription


// iOS-only properties

// The set of directivity configurations supported by the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/supportedPolarPatterns
func (a_ AudioSessionDataSourceDescription) SupportedPolarPatterns() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("supportedPolarPatterns"))
	return rv
}





