//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for PlayableContentManager


// iOS-only properties

// The current state of the playable content endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManager/context
func (p_ PlayableContentManager) Context() IMPPlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](p_.ID, objc.Sel("context"))
	return rv
}

// The data source provided by the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManager/dataSource
func (p_ PlayableContentManager) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dataSource"))
	return rv
}
func (p_ PlayableContentManager) SetDataSource(value unsafe.Pointer) {
	p_.ID.Send(objc.RegisterName("setDataSource:"), value)
}

// A delegate that lets the media player manage the app’s playback queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManager/delegate
func (p_ PlayableContentManager) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}
func (p_ PlayableContentManager) SetDelegate(value unsafe.Pointer) {
	p_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The content items currently playing based on their identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManager/nowPlayingIdentifiers
func (p_ PlayableContentManager) NowPlayingIdentifiers() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("nowPlayingIdentifiers"))
	return rv
}
func (p_ PlayableContentManager) SetNowPlayingIdentifiers(value []string) {
	p_.ID.Send(objc.RegisterName("setNowPlayingIdentifiers:"), value)
}





