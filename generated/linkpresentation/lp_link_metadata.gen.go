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
	IconProvider() ItemProvider /* not a class type */
	SetIconProvider(value ItemProvider /* not a class type */)
	ImageProvider() ItemProvider /* not a class type */
	SetImageProvider(value ItemProvider /* not a class type */)
	OriginalURL() foundation.objc.IObject /* cross-framework: URL */
	SetOriginalURL(value foundation.objc.IObject /* cross-framework: URL */)
	RemoteVideoURL() foundation.objc.IObject /* cross-framework: URL */
	SetRemoteVideoURL(value foundation.objc.IObject /* cross-framework: URL */)
	Title() string /* primitive/slice/pointer. */
	SetTitle(value string /* primitive/slice/pointer. */)
	URL() foundation.objc.IObject /* cross-framework: URL */
	SetURL(value foundation.objc.IObject /* cross-framework: URL */)
	VideoProvider() ItemProvider /* not a class type */
	SetVideoProvider(value ItemProvider /* not a class type */)
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
func (l_ LPLinkMetadata) IconProvider() ItemProvider /* not a class type */ {
	rv := objc.Send[ItemProvider](l_.ID, objc.Sel("iconProvider"))
	return rv
}


// An object that retrieves data corresponding to a representative icon for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/iconProvider
func (l_ LPLinkMetadata) SetIconProvider(value ItemProvider /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIconProvider:"), value)
}


// An object that retrieves data corresponding to a representative image for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/imageProvider
func (l_ LPLinkMetadata) ImageProvider() ItemProvider /* not a class type */ {
	rv := objc.Send[ItemProvider](l_.ID, objc.Sel("imageProvider"))
	return rv
}


// An object that retrieves data corresponding to a representative image for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/imageProvider
func (l_ LPLinkMetadata) SetImageProvider(value ItemProvider /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setImageProvider:"), value)
}


// The original URL of the metadata request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/originalURL
func (l_ LPLinkMetadata) OriginalURL() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](l_.ID, objc.Sel("originalURL"))
	return rv
}


// The original URL of the metadata request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/originalURL
func (l_ LPLinkMetadata) SetOriginalURL(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOriginalURL:"), value)
}


// A remote URL corresponding to a representative video for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/remoteVideoURL
func (l_ LPLinkMetadata) RemoteVideoURL() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](l_.ID, objc.Sel("remoteVideoURL"))
	return rv
}


// A remote URL corresponding to a representative video for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/remoteVideoURL
func (l_ LPLinkMetadata) SetRemoteVideoURL(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRemoteVideoURL:"), value)
}


// A representative title for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/title
func (l_ LPLinkMetadata) Title() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](l_.ID, objc.Sel("title"))
	return rv
}


// A representative title for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/title
func (l_ LPLinkMetadata) SetTitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The URL that returned the metadata, taking server-side redirects into account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/url
func (l_ LPLinkMetadata) URL() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](l_.ID, objc.Sel("URL"))
	return rv
}


// The URL that returned the metadata, taking server-side redirects into account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/url
func (l_ LPLinkMetadata) SetURL(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setURL:"), value)
}


// An object that retrieves data corresponding to a representative video for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/videoProvider
func (l_ LPLinkMetadata) VideoProvider() ItemProvider /* not a class type */ {
	rv := objc.Send[ItemProvider](l_.ID, objc.Sel("videoProvider"))
	return rv
}


// An object that retrieves data corresponding to a representative video for the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkMetadata/videoProvider
func (l_ LPLinkMetadata) SetVideoProvider(value ItemProvider /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setVideoProvider:"), value)
}



