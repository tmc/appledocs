// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNRenderingSessionAttributes] class.
var (
	CNRenderingSessionAttributesClass     _CNRenderingSessionAttributesClass
	CNRenderingSessionAttributesClassOnce sync.Once
)

func getCNRenderingSessionAttributesClass() _CNRenderingSessionAttributesClass {
	CNRenderingSessionAttributesClassOnce.Do(func() {
		CNRenderingSessionAttributesClass = _CNRenderingSessionAttributesClass{objc.GetClass("CNRenderingSessionAttributes")}
	})
	return CNRenderingSessionAttributesClass
}

type _CNRenderingSessionAttributesClass struct {
	class objc.Class
}

// An interface definition for the [CNRenderingSessionAttributes] class.
type ICNRenderingSessionAttributes interface {
	objectivec.IObject
	RenderingVersion() int
}

// A structure for movie-wide attributes required for proper rendering.
//
// The attributes include camera intrinsics from the camera on which the video was originally recorded.


// A structure for movie-wide attributes required for proper rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionAttributes

type CNRenderingSessionAttributes struct {
	objectivec.Object
}

// CNRenderingSessionAttributesFrom constructs a [CNRenderingSessionAttributes] from an unsafe.Pointer.
//
// A structure for movie-wide attributes required for proper rendering.
func CNRenderingSessionAttributesFrom(ptr unsafe.Pointer) CNRenderingSessionAttributes {
	return CNRenderingSessionAttributes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNRenderingSessionAttributesClass) Alloc() CNRenderingSessionAttributes {
	rv := objc.Send[CNRenderingSessionAttributes](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNRenderingSessionAttributesClass) New() CNRenderingSessionAttributes {
	rv := objc.Send[CNRenderingSessionAttributes](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNRenderingSessionAttributes) Init() CNRenderingSessionAttributes {
	rv := objc.Send[CNRenderingSessionAttributes](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNRenderingSessionAttributes) Autorelease() CNRenderingSessionAttributes {
	rv := objc.Send[CNRenderingSessionAttributes](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNRenderingSessionAttributes creates a new CNRenderingSessionAttributes instance.
func NewCNRenderingSessionAttributes() CNRenderingSessionAttributes {
	return getCNRenderingSessionAttributesClass().New()
}



// Loads the rendering session attributes from an asset asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionAttributes/loadFromAsset:completionHandler:

func (cc _CNRenderingSessionAttributesClass) LoadFromAssetCompletionHandler(asset avfoundation.IAsset, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("loadFromAsset:completionHandler:"), asset, completionHandler)
}


// The primary version number used to render the original Cinematic move that determines compatibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionAttributes/renderingVersion

func (c_ CNRenderingSessionAttributes) RenderingVersion() int {
	rv := objc.Send[int](c_.ID, objc.Sel("renderingVersion"))
	return rv
}



