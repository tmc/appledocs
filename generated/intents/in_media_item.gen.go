// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INMediaItem] class.
var (
	INMediaItemClass     _INMediaItemClass
	INMediaItemClassOnce sync.Once
)

func getINMediaItemClass() _INMediaItemClass {
	INMediaItemClassOnce.Do(func() {
		INMediaItemClass = _INMediaItemClass{objc.GetClass("INMediaItem")}
	})
	return INMediaItemClass
}

type _INMediaItemClass struct {
	class objc.Class
}

// An interface definition for the [INMediaItem] class.
type IINMediaItem interface {
	objectivec.IObject
	Artist() string
	SetArtist(value string)
	Artwork() INImage
	SetArtwork(value INImage)
	Identifier() string
	SetIdentifier(value string)
	Title() string
	SetTitle(value string)
	Type() unsafe.Pointer
	SetType(value unsafe.Pointer)
}

// An object that describes a piece of media content, such as a song, TV show, artist, or podcast playlist.


// An object that describes a piece of media content, such as a song, TV show, artist, or podcast playlist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INMediaItem
type INMediaItem struct {
	objectivec.Object
}

// INMediaItemFrom constructs a [INMediaItem] from an unsafe.Pointer.
//
// An object that describes a piece of media content, such as a song, TV show, artist, or podcast playlist.
func INMediaItemFrom(ptr unsafe.Pointer) INMediaItem {
	return INMediaItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INMediaItemClass) Alloc() INMediaItem {
	rv := objc.Send[INMediaItem](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INMediaItemClass) New() INMediaItem {
	rv := objc.Send[INMediaItem](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INMediaItem) Init() INMediaItem {
	rv := objc.Send[INMediaItem](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INMediaItem) Autorelease() INMediaItem {
	rv := objc.Send[INMediaItem](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINMediaItem creates a new INMediaItem instance.
func NewINMediaItem() INMediaItem {
	return getINMediaItemClass().New()
}



// The artist associated with the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediaitem/artist
func (i_ INMediaItem) Artist() string {
	rv := objc.Send[string](i_.ID, objc.Sel("artist"))
	return rv
}


// The artist associated with the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediaitem/artist
func (i_ INMediaItem) SetArtist(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setArtist:"), objc.String(value))
}


// Artwork for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediaitem/artwork
func (i_ INMediaItem) Artwork() INImage {
	rv := objc.Send[INImage](i_.ID, objc.Sel("artwork"))
	return rv
}


// Artwork for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediaitem/artwork
func (i_ INMediaItem) SetArtwork(value INImage) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setArtwork:"), value)
}


// The value your app uses to identify the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediaitem/identifier
func (i_ INMediaItem) Identifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("identifier"))
	return rv
}


// The value your app uses to identify the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediaitem/identifier
func (i_ INMediaItem) SetIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The media item title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediaitem/title
func (i_ INMediaItem) Title() string {
	rv := objc.Send[string](i_.ID, objc.Sel("title"))
	return rv
}


// The media item title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediaitem/title
func (i_ INMediaItem) SetTitle(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The media item type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediaitem/type
func (i_ INMediaItem) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("type"))
	return rv
}


// The media item type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediaitem/type
func (i_ INMediaItem) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setType:"), value)
}



