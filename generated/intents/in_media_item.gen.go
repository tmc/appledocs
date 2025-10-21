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
}

// An object that describes a piece of media content, such as a song, TV show, artist, or podcast playlist.
//
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




