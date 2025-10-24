// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	ImageCropRect() objc.IObject /* cross-framework: Rect */
	SetImageCropRect(value objc.IObject /* cross-framework: Rect */)
	Context() IMPPlayableContentManagerContext
	SetContext(value IMPPlayableContentManagerContext)
	DataSource() PlayableContentDataSource /* not a class type */
	SetDataSource(value PlayableContentDataSource /* not a class type */)
	Delegate() PlayableContentDelegate /* not a class type */
	SetDelegate(value PlayableContentDelegate /* not a class type */)
	NowPlayingIdentifiers() objc.IObject /* cross-framework: NSString */
	SetNowPlayingIdentifiers(value objc.IObject /* cross-framework: NSString */)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
	// methods:
}

// A shared content manager for controlling interactions between your media app and system-provided or external media player interfaces.
//
// The app provides data to the content manager so that the media player can browse the content provided. A delegate provides the media player the ability to perform actions that manage the app’s playback queue. You don’t create a new content manager directly, instead you grab the shared content manager using the method. After getting the shared content manager, your next step depends on the features your app supports: To provide content navigation and suggested content for CarPlay, immediately set both the and properties. After setting these properties, use the and methods to load the information from the data source. To provide suggested content when the user connects headphones, a Bluetooth stereo, or another output device, set only the property. After you set a delegate, iOS automatically calls methods in the protocol allowing you to suggest appropriate content.


// A shared content manager for controlling interactions between your media app and system-provided or external media player interfaces.
//
// [Full Topic]
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



// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (p_ PlayableContentManager) ImageCropRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](p_.ID, objc.Sel("imageCropRect"))
	return rv
}


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (p_ PlayableContentManager) SetImageCropRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImageCropRect:"), value)
}


// The current state of the playable content endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/context
func (p_ PlayableContentManager) Context() IMPPlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](p_.ID, objc.Sel("context"))
	return rv
}


// The current state of the playable content endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/context
func (p_ PlayableContentManager) SetContext(value IMPPlayableContentManagerContext) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContext:"), value)
}


// The data source provided by the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/datasource
func (p_ PlayableContentManager) DataSource() PlayableContentDataSource /* not a class type */ {
	rv := objc.Send[PlayableContentDataSource](p_.ID, objc.Sel("dataSource"))
	return rv
}


// The data source provided by the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/datasource
func (p_ PlayableContentManager) SetDataSource(value PlayableContentDataSource /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDataSource:"), value)
}


// A delegate that lets the media player manage the app’s playback queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/delegate
func (p_ PlayableContentManager) Delegate() PlayableContentDelegate /* not a class type */ {
	rv := objc.Send[PlayableContentDelegate](p_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate that lets the media player manage the app’s playback queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/delegate
func (p_ PlayableContentManager) SetDelegate(value PlayableContentDelegate /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}


// The content items currently playing based on their identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/nowplayingidentifiers
func (p_ PlayableContentManager) NowPlayingIdentifiers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("nowPlayingIdentifiers"))
	return rv
}


// The content items currently playing based on their identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanager/nowplayingidentifiers
func (p_ PlayableContentManager) SetNowPlayingIdentifiers(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNowPlayingIdentifiers:"), value)
}


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (p_ PlayableContentManager) ShowsRouteButton() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (p_ PlayableContentManager) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsRouteButton:"), value)
}


