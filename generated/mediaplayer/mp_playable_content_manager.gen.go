// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlayableContentManager] class.
var (
	PlayableContentManagerClass     _PlayableContentManagerClass
	PlayableContentManagerClassOnce sync.Once
)

func getPlayableContentManagerClass() _PlayableContentManagerClass {
	PlayableContentManagerClassOnce.Do(func() {
		PlayableContentManagerClass = _PlayableContentManagerClass{objc.GetClass("MPPlayableContentManager")}
	})
	return PlayableContentManagerClass
}

type _PlayableContentManagerClass struct {
	class objc.Class
}

// An interface definition for the [PlayableContentManager] class.
type IPlayableContentManager interface {
	objectivec.IObject
	BeginUpdates()
	EndUpdates()
}

// A shared content manager for controlling interactions between your media app and system-provided or external media player interfaces.
//
// The app provides data to the content manager so that the media player can browse the content provided. A delegate provides the media player the ability to perform actions that manage the app’s playback queue. You don’t create a new content manager directly, instead you grab the shared content manager using the method. After getting the shared content manager, your next step depends on the features your app supports: To provide content navigation and suggested content for CarPlay, immediately set both the and properties. After setting these properties, use the and methods to load the information from the data source. To provide suggested content when the user connects headphones, a Bluetooth stereo, or another output device, set only the property. After you set a delegate, iOS automatically calls methods in the protocol allowing you to suggest appropriate content.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManager
type PlayableContentManager struct {
	objectivec.Object
}

// PlayableContentManagerFrom constructs a [PlayableContentManager] from an unsafe.Pointer.
//
// A shared content manager for controlling interactions between your media app and system-provided or external media player interfaces.
func PlayableContentManagerFrom(ptr unsafe.Pointer) PlayableContentManager {
	return PlayableContentManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayableContentManagerClass) Alloc() PlayableContentManager {
	rv := objc.Send[PlayableContentManager](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayableContentManagerClass) New() PlayableContentManager {
	rv := objc.Send[PlayableContentManager](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayableContentManager) Init() PlayableContentManager {
	rv := objc.Send[PlayableContentManager](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayableContentManager) Autorelease() PlayableContentManager {
	rv := objc.Send[PlayableContentManager](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayableContentManager creates a new PlayableContentManager instance.
func NewPlayableContentManager() PlayableContentManager {
	return getPlayableContentManagerClass().New()
}


// Updates several Media Player content items at once.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManager/beginUpdates()
func (p_ PlayableContentManager) BeginUpdates() {
	objc.Send[objc.ID](p_.ID, objc.Sel("beginUpdates"))
}

// Ends a synchronized update.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManager/endUpdates()
func (p_ PlayableContentManager) EndUpdates() {
	objc.Send[objc.ID](p_.ID, objc.Sel("endUpdates"))
}

// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (p_ PlayableContentManager) ImageCropRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("imageCropRect"))
	return rv
}


// SetImageCropRect sets the value of the imageCropRect property.
// The bounds, in points, of the content area for the full size image associated with the media item artwork.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (p_ PlayableContentManager) SetImageCropRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImageCropRect:"), value)
}

// The current state of the playable content endpoint.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/context
func (p_ PlayableContentManager) Context() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("context"))
	return rv
}


// SetContext sets the value of the context property.
// The current state of the playable content endpoint.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/context
func (p_ PlayableContentManager) SetContext(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContext:"), value)
}

// The data source provided by the app.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/datasource
func (p_ PlayableContentManager) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dataSource"))
	return rv
}


// SetDataSource sets the value of the dataSource property.
// The data source provided by the app.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/datasource
func (p_ PlayableContentManager) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDataSource:"), value)
}

// A delegate that lets the media player manage the app’s playback queue.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/delegate
func (p_ PlayableContentManager) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate that lets the media player manage the app’s playback queue.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/delegate
func (p_ PlayableContentManager) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}

// The content items currently playing based on their identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/nowplayingidentifiers
func (p_ PlayableContentManager) NowPlayingIdentifiers() string {
	rv := objc.Send[string](p_.ID, objc.Sel("nowPlayingIdentifiers"))
	return rv
}


// SetNowPlayingIdentifiers sets the value of the nowPlayingIdentifiers property.
// The content items currently playing based on their identifiers.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/nowplayingidentifiers
func (p_ PlayableContentManager) SetNowPlayingIdentifiers(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNowPlayingIdentifiers:"), objc.String(value))
}

// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (p_ PlayableContentManager) ShowsRouteButton() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// SetShowsRouteButton sets the value of the showsRouteButton property.
// A Boolean value that indicates whether the route button is visible in the volume view.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (p_ PlayableContentManager) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsRouteButton:"), value)
}



