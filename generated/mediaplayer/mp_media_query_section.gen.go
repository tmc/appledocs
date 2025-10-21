// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MediaQuerySection] class.
var (
	MediaQuerySectionClass     _MediaQuerySectionClass
	MediaQuerySectionClassOnce sync.Once
)

func getMediaQuerySectionClass() _MediaQuerySectionClass {
	MediaQuerySectionClassOnce.Do(func() {
		MediaQuerySectionClass = _MediaQuerySectionClass{objc.GetClass("MPMediaQuerySection")}
	})
	return MediaQuerySectionClass
}

type _MediaQuerySectionClass struct {
	class objc.Class
}

// An interface definition for the [MediaQuerySection] class.
type IMediaQuerySection interface {
	objectivec.IObject
}

// A range of media items or media item collections from within a media query.
//
// You can use sections when displaying a query’s items or collections in your app’s user interface. You obtain an array of media query sections by using the or properties of a media query (an instance of the class). The property values of a media query section are read-only.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuerySection
type MediaQuerySection struct {
	objectivec.Object
}

// MediaQuerySectionFrom constructs a [MediaQuerySection] from an unsafe.Pointer.
//
// A range of media items or media item collections from within a media query.
func MediaQuerySectionFrom(ptr unsafe.Pointer) MediaQuerySection {
	return MediaQuerySection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaQuerySectionClass) Alloc() MediaQuerySection {
	rv := objc.Send[MediaQuerySection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaQuerySectionClass) New() MediaQuerySection {
	rv := objc.Send[MediaQuerySection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaQuerySection) Init() MediaQuerySection {
	rv := objc.Send[MediaQuerySection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaQuerySection) Autorelease() MediaQuerySection {
	rv := objc.Send[MediaQuerySection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaQuerySection creates a new MediaQuerySection instance.
func NewMediaQuerySection() MediaQuerySection {
	return getMediaQuerySectionClass().New()
}


// The range in the media query’s items or collections array that the media query section represents.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuerySection/range
func (m_ MediaQuerySection) Range() foundation.Range {
	rv := objc.Send[foundation.Range](m_.ID, objc.Sel("range"))
	return rv
}

// The localized title of the media query section.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuerySection/title
func (m_ MediaQuerySection) Title() string {
	rv := objc.Send[string](m_.ID, objc.Sel("title"))
	return rv
}



