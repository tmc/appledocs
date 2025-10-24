// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer_test

import (
	"github.com/tmc/appledocs/generated/mediaplayer"
)

// Suppress unused import errors
var _ = mediaplayer.NewNowPlayingInfoLanguageOption

// ExampleNowPlayingInfoLanguageOption_IsAutomaticAudibleLanguageOption demonstrates using IsAutomaticAudibleLanguageOption on a NowPlayingInfoLanguageOption instance.
// Returns a Boolean value that determines whether to use the best audible language option based on the system preferences.
func ExampleNowPlayingInfoLanguageOption_IsAutomaticAudibleLanguageOption() {
	obj := mediaplayer.NewNowPlayingInfoLanguageOption()
	_ = obj.IsAutomaticAudibleLanguageOption()
	// Output:
	}

// ExampleNowPlayingInfoLanguageOption_IsAutomaticLegibleLanguageOption demonstrates using IsAutomaticLegibleLanguageOption on a NowPlayingInfoLanguageOption instance.
// Returns a Boolean value that determines whether to use the best legible language option based on the system preferences.
func ExampleNowPlayingInfoLanguageOption_IsAutomaticLegibleLanguageOption() {
	obj := mediaplayer.NewNowPlayingInfoLanguageOption()
	_ = obj.IsAutomaticLegibleLanguageOption()
	// Output:
	}

