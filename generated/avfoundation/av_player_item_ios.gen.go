//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for PlayerItem


// Cancels a pending authorization request and dismisses the passcode entry, if displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/cancelPlaybackRestrictionsAuthorizationRequest()
func (p_ PlayerItem) CancelPlaybackRestrictionsAuthorizationRequest() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancelPlaybackRestrictionsAuthorizationRequest"))
}

// Determines whether this item is subject to parental restrictions, and, if so, prompts the user to enter the restrictions passcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/requestPlaybackRestrictionsAuthorization(_:)
func (p_ PlayerItem) RequestPlaybackRestrictionsAuthorization(completion unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("requestPlaybackRestrictionsAuthorization:"), completion)
}

// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/externalSubtitleOptionLanguages
func (p_ PlayerItem) ExternalSubtitleOptionLanguages() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("externalSubtitleOptionLanguages"))
	return rv
}
func (p_ PlayerItem) SetExternalSubtitleOptionLanguages(value []string) {
	p_.ID.Send(objc.RegisterName("setExternalSubtitleOptionLanguages:"), value)
}

// An array of additional metadata for the player item to supplement or replace an asset’s embedded metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/externalMetadata
func (p_ PlayerItem) ExternalMetadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](p_.ID, objc.Sel("externalMetadata"))
	return rv
}
func (p_ PlayerItem) SetExternalMetadata(value []MetadataItem) {
	p_.ID.Send(objc.RegisterName("setExternalMetadata:"), value)
}

// An array of time ranges that identify interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/interstitialTimeRanges
func (p_ PlayerItem) InterstitialTimeRanges() []InterstitialTimeRange /* not a class type */ {
	rv := objc.Send[[]InterstitialTimeRange](p_.ID, objc.Sel("interstitialTimeRanges"))
	return rv
}

// The time marker groups that provide ways to navigate the player item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/navigationMarkerGroups
func (p_ PlayerItem) NavigationMarkerGroups() []NavigationMarkersGroup /* not a class type */ {
	rv := objc.Send[[]NavigationMarkersGroup](p_.ID, objc.Sel("navigationMarkerGroups"))
	return rv
}
func (p_ PlayerItem) SetNavigationMarkerGroups(value []NavigationMarkersGroup /* not a class type */) {
	p_.ID.Send(objc.RegisterName("setNavigationMarkerGroups:"), value)
}

// The item proposed to follow the current content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/nextContentProposal
func (p_ PlayerItem) NextContentProposal() ContentProposal /* not a class type */ {
	rv := objc.Send[ContentProposal](p_.ID, objc.Sel("nextContentProposal"))
	return rv
}
func (p_ PlayerItem) SetNextContentProposal(value ContentProposal /* not a class type */) {
	p_.ID.Send(objc.RegisterName("setNextContentProposal:"), value)
}

// The current now playing information for the player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/nowPlayingInfo
func (p_ PlayerItem) NowPlayingInfo() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("nowPlayingInfo"))
	return rv
}
func (p_ PlayerItem) SetNowPlayingInfo(value foundation.IDictionary) {
	p_.ID.Send(objc.RegisterName("setNowPlayingInfo:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectedExternalSubtitleOptionLanguage
func (p_ PlayerItem) SelectedExternalSubtitleOptionLanguage() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("selectedExternalSubtitleOptionLanguage"))
	return rv
}
func (p_ PlayerItem) SetSelectedExternalSubtitleOptionLanguage(value foundation.foundation.INSString) {
	p_.ID.Send(objc.RegisterName("setSelectedExternalSubtitleOptionLanguage:"), value)
}

// A Boolean value that indicates whether the player translates interstitial events to interstitial time ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/translatesPlayerInterstitialEvents
func (p_ PlayerItem) TranslatesPlayerInterstitialEvents() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("translatesPlayerInterstitialEvents"))
	return rv
}
func (p_ PlayerItem) SetTranslatesPlayerInterstitialEvents(value bool) {
	p_.ID.Send(objc.RegisterName("setTranslatesPlayerInterstitialEvents:"), value)
}




