// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer_test

import (
	"github.com/tmc/appledocs/generated/mediaplayer"
)

// Suppress unused import errors
var _ = mediaplayer.NewNowPlayingInfoLanguageOption

// ExampleNewNowPlayingInfoLanguageOptionWithTypeLanguageTagCharacteristicsDisplayNameIdentifier demonstrates how to create a NowPlayingInfoLanguageOption instance using NewNowPlayingInfoLanguageOptionWithTypeLanguageTagCharacteristicsDisplayNameIdentifier.
// Creates a single language option.
func ExampleNewNowPlayingInfoLanguageOptionWithTypeLanguageTagCharacteristicsDisplayNameIdentifier() {
	_ = mediaplayer.NewNowPlayingInfoLanguageOptionWithTypeLanguageTagCharacteristicsDisplayNameIdentifier(
		mediaplayer.NowPlayingInfoLanguageOptionType{}, // languageOptionType NowPlayingInfoLanguageOptionType
		"languageTag", // languageTag string
		[]mediaplayer.string{}, // languageOptionCharacteristics []string
		"displayName", // displayName string
		"identifier", // identifier string
	)
	// Output:
}
