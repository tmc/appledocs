// Code generated from Apple documentation for LinkPresentation. DO NOT EDIT.

package linkpresentation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LPLinkMetadata] class.
var (
	LPLinkMetadataClass     _LPLinkMetadataClass
	LPLinkMetadataClassOnce sync.Once
)

func getLPLinkMetadataClass() _LPLinkMetadataClass {
	LPLinkMetadataClassOnce.Do(func() {
		LPLinkMetadataClass = _LPLinkMetadataClass{objc.GetClass("LPLinkMetadata")}
	})
	return LPLinkMetadataClass
}

type _LPLinkMetadataClass struct {
	class objc.Class
}

// An interface definition for the [LPLinkMetadata] class.
type ILPLinkMetadata interface {
	objectivec.IObject
	// properties:
	IconProvider() objc.IObject /* cross-framework: ItemProvider */
	SetIconProvider(value objc.IObject /* cross-framework: ItemProvider */)
	ImageProvider() objc.IObject /* cross-framework: ItemProvider */
	SetImageProvider(value objc.IObject /* cross-framework: ItemProvider */)
	OriginalURL() objc.IObject /* cross-framework: NSURL */
	SetOriginalURL(value objc.IObject /* cross-framework: NSURL */)
	RemoteVideoURL() objc.IObject /* cross-framework: NSURL */
	SetRemoteVideoURL(value objc.IObject /* cross-framework: NSURL */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	URL() objc.IObject /* cross-framework: NSURL */
	SetURL(value objc.IObject /* cross-framework: NSURL */)
	VideoProvider() objc.IObject /* cross-framework: ItemProvider */
	SetVideoProvider(value objc.IObject /* cross-framework: ItemProvider */)
	// methods:
}

// An object that contains metadata about a URL.
//
// Use to store the metadata about a URL, including its title, icon, images and video. Fetch metadata using . For remote URLs, cache the metadata locally to avoid the data and performance cost of fetching it from the internet every time you present it. is serializable with . For local file URLs, the API retrieves a representative thumbnail for the file, if possible.


// An object that contains metadata about a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata
type LPLinkMetadata struct {
	objectivec.Object
}

// LPLinkMetadataFrom constructs a [LPLinkMetadata] from an unsafe.Pointer.
//
// An object that contains metadata about a URL.
func LPLinkMetadataFrom(ptr unsafe.Pointer) LPLinkMetadata {
	return LPLinkMetadata{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LPLinkMetadataClass) Alloc() LPLinkMetadata {
	rv := objc.Send[LPLinkMetadata](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LPLinkMetadataClass) New() LPLinkMetadata {
	rv := objc.Send[LPLinkMetadata](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LPLinkMetadata) Init() LPLinkMetadata {
	rv := objc.Send[LPLinkMetadata](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LPLinkMetadata) Autorelease() LPLinkMetadata {
	rv := objc.Send[LPLinkMetadata](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLPLinkMetadata creates a new LPLinkMetadata instance.
func NewLPLinkMetadata() LPLinkMetadata {
	return getLPLinkMetadataClass().New()
}



// An object that retrieves data corresponding to a representative icon for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/iconProvider
func (l_ LPLinkMetadata) IconProvider() objc.IObject /* cross-framework: ItemProvider */ {
	rv := objc.Send[foundation.ItemProvider](l_.ID, objc.Sel("iconProvider"))
	return rv
}


// An object that retrieves data corresponding to a representative icon for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/iconProvider
func (l_ LPLinkMetadata) SetIconProvider(value objc.IObject /* cross-framework: ItemProvider */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIconProvider:"), value)
}


// An object that retrieves data corresponding to a representative image for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/imageProvider
func (l_ LPLinkMetadata) ImageProvider() objc.IObject /* cross-framework: ItemProvider */ {
	rv := objc.Send[foundation.ItemProvider](l_.ID, objc.Sel("imageProvider"))
	return rv
}


// An object that retrieves data corresponding to a representative image for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/imageProvider
func (l_ LPLinkMetadata) SetImageProvider(value objc.IObject /* cross-framework: ItemProvider */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setImageProvider:"), value)
}


// The original URL of the metadata request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/originalURL
func (l_ LPLinkMetadata) OriginalURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](l_.ID, objc.Sel("originalURL"))
	return rv
}


// The original URL of the metadata request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/originalURL
func (l_ LPLinkMetadata) SetOriginalURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOriginalURL:"), value)
}


// A remote URL corresponding to a representative video for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/remoteVideoURL
func (l_ LPLinkMetadata) RemoteVideoURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](l_.ID, objc.Sel("remoteVideoURL"))
	return rv
}


// A remote URL corresponding to a representative video for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/remoteVideoURL
func (l_ LPLinkMetadata) SetRemoteVideoURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRemoteVideoURL:"), value)
}


// A representative title for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/title
func (l_ LPLinkMetadata) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("title"))
	return rv
}


// A representative title for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/title
func (l_ LPLinkMetadata) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTitle:"), value)
}


// The URL that returned the metadata, taking server-side redirects into account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/url
func (l_ LPLinkMetadata) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](l_.ID, objc.Sel("URL"))
	return rv
}


// The URL that returned the metadata, taking server-side redirects into account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/url
func (l_ LPLinkMetadata) SetURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setURL:"), value)
}


// An object that retrieves data corresponding to a representative video for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/videoProvider
func (l_ LPLinkMetadata) VideoProvider() objc.IObject /* cross-framework: ItemProvider */ {
	rv := objc.Send[foundation.ItemProvider](l_.ID, objc.Sel("videoProvider"))
	return rv
}


// An object that retrieves data corresponding to a representative video for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/videoProvider
func (l_ LPLinkMetadata) SetVideoProvider(value objc.IObject /* cross-framework: ItemProvider */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setVideoProvider:"), value)
}



