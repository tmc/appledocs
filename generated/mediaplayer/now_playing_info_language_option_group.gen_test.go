// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer_test

import (
	"github.com/tmc/appledocs/generated/mediaplayer"
)

// Suppress unused import errors
var _ = mediaplayer.NewNowPlayingInfoLanguageOptionGroup

// ExampleNewNowPlayingInfoLanguageOptionGroupWithLanguageOptionsDefaultLanguageOptionAllowEmptySelection demonstrates how to create a NowPlayingInfoLanguageOptionGroup instance using NewNowPlayingInfoLanguageOptionGroupWithLanguageOptionsDefaultLanguageOptionAllowEmptySelection.
// Creates a new language option group with the supplied language options.
func ExampleNewNowPlayingInfoLanguageOptionGroupWithLanguageOptionsDefaultLanguageOptionAllowEmptySelection() {
	_ = mediaplayer.NewNowPlayingInfoLanguageOptionGroupWithLanguageOptionsDefaultLanguageOptionAllowEmptySelection(
		[]mediaplayer.NowPlayingInfoLanguageOption{}, // languageOptions []NowPlayingInfoLanguageOption
		mediaplayer.MPNowPlayingInfoLanguageOption{}, // defaultLanguageOption MPNowPlayingInfoLanguageOption
		false, // allowEmptySelection bool
	)
	// Output:
}
