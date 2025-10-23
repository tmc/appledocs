// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Animation() appkit.Animation /* not a class type */
	SetAnimation(value appkit.Animation /* not a class type */)
	Region() CaptionRegion /* not a class type */
	SetRegion(value CaptionRegion /* not a class type */)
	Text() string /* primitive/slice/pointer */
	SetText(value string /* primitive/slice/pointer */)
	TextAlignment() unsafe.Pointer
	SetTextAlignment(value unsafe.Pointer)
	TimeRange() TimeRange /* not a class type */
	SetTimeRange(value TimeRange /* not a class type */)
	// methods:
}

// An object that represents text to present over a time range.
//
// A caption contains a cue, which is a single sentence or paragraph of text for a time range in the video timeline. Within the active range, the caption may animate (for example, Karaoke lyrics) by rolling-up, changing visibility, or using other dynamic styling.


// An object that represents text to present over a time range.
//
// [Full Topic]
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



// The animation that the system applies to this caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/animation-swift.property
func (c_ Caption) Animation() appkit.Animation /* not a class type */ {
	rv := objc.Send[appkit.Animation](c_.ID, objc.Sel("animation"))
	return rv
}


// The animation that the system applies to this caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/animation-swift.property
func (c_ Caption) SetAnimation(value appkit.Animation /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAnimation:"), value)
}


// The region in which the caption exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/region
func (c_ Caption) Region() CaptionRegion /* not a class type */ {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("region"))
	return rv
}


// The region in which the caption exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/region
func (c_ Caption) SetRegion(value CaptionRegion /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRegion:"), value)
}


// The caption text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/text
func (c_ Caption) Text() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](c_.ID, objc.Sel("text"))
	return rv
}


// The caption text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/text
func (c_ Caption) SetText(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setText:"), objc.String(value))
}


// The alignment for the caption text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/textalignment-swift.property
func (c_ Caption) TextAlignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("textAlignment"))
	return rv
}


// The alignment for the caption text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/textalignment-swift.property
func (c_ Caption) SetTextAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextAlignment:"), value)
}


// The time range over which the system presents the caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/timerange
func (c_ Caption) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("timeRange"))
	return rv
}


// The time range over which the system presents the caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/timerange
func (c_ Caption) SetTimeRange(value TimeRange /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeRange:"), value)
}



