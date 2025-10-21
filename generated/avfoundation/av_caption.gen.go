// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Caption] class.
var (
	CaptionClass     _CaptionClass
	CaptionClassOnce sync.Once
)

func getCaptionClass() _CaptionClass {
	CaptionClassOnce.Do(func() {
		CaptionClass = _CaptionClass{objc.GetClass("AVCaption")}
	})
	return CaptionClass
}

type _CaptionClass struct {
	class objc.Class
}

// An interface definition for the [Caption] class.
type ICaption interface {
	objectivec.IObject
}

// An object that represents text to present over a time range.
//
// A caption contains a cue, which is a single sentence or paragraph of text for a time range in the video timeline. Within the active range, the caption may animate (for example, Karaoke lyrics) by rolling-up, changing visibility, or using other dynamic styling.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption
type Caption struct {
	objectivec.Object
}

// CaptionFrom constructs a [Caption] from an unsafe.Pointer.
//
// An object that represents text to present over a time range.
func CaptionFrom(ptr unsafe.Pointer) Caption {
	return Caption{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptionClass) Alloc() Caption {
	rv := objc.Send[Caption](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptionClass) New() Caption {
	rv := objc.Send[Caption](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Caption) Init() Caption {
	rv := objc.Send[Caption](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Caption) Autorelease() Caption {
	rv := objc.Send[Caption](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaption creates a new Caption instance.
func NewCaption() Caption {
	return getCaptionClass().New()
}




