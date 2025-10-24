// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlayableContentManagerContext] class.
var (
	PlayableContentManagerContextClass     _PlayableContentManagerContextClass
	PlayableContentManagerContextClassOnce sync.Once
)

func getPlayableContentManagerContextClass() _PlayableContentManagerContextClass {
	PlayableContentManagerContextClassOnce.Do(func() {
		PlayableContentManagerContextClass = _PlayableContentManagerContextClass{objc.GetClass("MPPlayableContentManagerContext")}
	})
	return PlayableContentManagerContextClass
}

type _PlayableContentManagerContextClass struct {
	class objc.Class
}

// An interface definition for the [PlayableContentManagerContext] class.
type IPlayableContentManagerContext interface {
	objectivec.IObject
	// properties:
	ImageCropRect() objc.IObject /* cross-framework: Rect */
	SetImageCropRect(value objc.IObject /* cross-framework: Rect */)
	ContentLimitsEnabled() bool
	SetContentLimitsEnabled(value bool)
	ContentLimitsEnforced() bool
	SetContentLimitsEnforced(value bool)
	EndpointAvailable() bool
	SetEndpointAvailable(value bool)
	EnforcedContentItemsCount() int
	SetEnforcedContentItemsCount(value int)
	EnforcedContentTreeDepth() int
	SetEnforcedContentTreeDepth(value int)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
	// methods:
}

// An object representing the current state of the playable endpoint.


// An object representing the current state of the playable endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManagerContext
type PlayableContentManagerContext struct {
	objectivec.Object
}

// PlayableContentManagerContextFrom constructs a [PlayableContentManagerContext] from an unsafe.Pointer.
//
// An object representing the current state of the playable endpoint.
func PlayableContentManagerContextFrom(ptr unsafe.Pointer) PlayableContentManagerContext {
	return PlayableContentManagerContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayableContentManagerContextClass) Alloc() PlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayableContentManagerContextClass) New() PlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayableContentManagerContext) Init() PlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayableContentManagerContext) Autorelease() PlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayableContentManagerContext creates a new PlayableContentManagerContext instance.
func NewPlayableContentManagerContext() PlayableContentManagerContext {
	return getPlayableContentManagerContextClass().New()
}



// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (p_ PlayableContentManagerContext) ImageCropRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](p_.ID, objc.Sel("imageCropRect"))
	return rv
}


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (p_ PlayableContentManagerContext) SetImageCropRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImageCropRect:"), value)
}


// A Boolean value that indicates whether the content server enables content limits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanagercontext/contentlimitsenabled
func (p_ PlayableContentManagerContext) ContentLimitsEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("contentLimitsEnabled"))
	return rv
}


// A Boolean value that indicates whether the content server enables content limits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanagercontext/contentlimitsenabled
func (p_ PlayableContentManagerContext) SetContentLimitsEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentLimitsEnabled:"), value)
}


// A Boolean value that indicates whether the content server enforces content limits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanagercontext/contentlimitsenforced
func (p_ PlayableContentManagerContext) ContentLimitsEnforced() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("contentLimitsEnforced"))
	return rv
}


// A Boolean value that indicates whether the content server enforces content limits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanagercontext/contentlimitsenforced
func (p_ PlayableContentManagerContext) SetContentLimitsEnforced(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentLimitsEnforced:"), value)
}


// Returns a Boolean that indicates whether the content server is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanagercontext/endpointavailable
func (p_ PlayableContentManagerContext) EndpointAvailable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("endpointAvailable"))
	return rv
}


// Returns a Boolean that indicates whether the content server is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanagercontext/endpointavailable
func (p_ PlayableContentManagerContext) SetEndpointAvailable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEndpointAvailable:"), value)
}


// Returns the number of content items to display during content limiting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanagercontext/enforcedcontentitemscount
func (p_ PlayableContentManagerContext) EnforcedContentItemsCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("enforcedContentItemsCount"))
	return rv
}


// Returns the number of content items to display during content limiting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanagercontext/enforcedcontentitemscount
func (p_ PlayableContentManagerContext) SetEnforcedContentItemsCount(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEnforcedContentItemsCount:"), value)
}


// The maximum depth of the navigation hierarchy allowed by the content server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanagercontext/enforcedcontenttreedepth
func (p_ PlayableContentManagerContext) EnforcedContentTreeDepth() int {
	rv := objc.Send[int](p_.ID, objc.Sel("enforcedContentTreeDepth"))
	return rv
}


// The maximum depth of the navigation hierarchy allowed by the content server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpplayablecontentmanagercontext/enforcedcontenttreedepth
func (p_ PlayableContentManagerContext) SetEnforcedContentTreeDepth(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEnforcedContentTreeDepth:"), value)
}


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (p_ PlayableContentManagerContext) ShowsRouteButton() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (p_ PlayableContentManagerContext) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsRouteButton:"), value)
}



