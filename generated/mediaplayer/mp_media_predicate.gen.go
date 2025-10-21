// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MediaPredicate] class.
var (
	MediaPredicateClass     _MediaPredicateClass
	MediaPredicateClassOnce sync.Once
)

func getMediaPredicateClass() _MediaPredicateClass {
	MediaPredicateClassOnce.Do(func() {
		MediaPredicateClass = _MediaPredicateClass{objc.GetClass("MPMediaPredicate")}
	})
	return MediaPredicateClass
}

type _MediaPredicateClass struct {
	class objc.Class
}

// An interface definition for the [MediaPredicate] class.
type IMediaPredicate interface {
	objectivec.IObject
}

// An abstract class that defines classes for filtering media in a media query.
//
// In media queries, a is a statement of a logical condition that you want to test each media item against. The system returns the media items that satisfy the condition in the query result. Use this class’s concrete subclass, described in , to define the filter in a media query to retrieve a subset of media items from the library. For more information about media queries, see .
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPredicate
type MediaPredicate struct {
	objectivec.Object
}

// MediaPredicateFrom constructs a [MediaPredicate] from an unsafe.Pointer.
//
// An abstract class that defines classes for filtering media in a media query.
func MediaPredicateFrom(ptr unsafe.Pointer) MediaPredicate {
	return MediaPredicate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaPredicateClass) Alloc() MediaPredicate {
	rv := objc.Send[MediaPredicate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaPredicateClass) New() MediaPredicate {
	rv := objc.Send[MediaPredicate](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaPredicate) Init() MediaPredicate {
	rv := objc.Send[MediaPredicate](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaPredicate) Autorelease() MediaPredicate {
	rv := objc.Send[MediaPredicate](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaPredicate creates a new MediaPredicate instance.
func NewMediaPredicate() MediaPredicate {
	return getMediaPredicateClass().New()
}




