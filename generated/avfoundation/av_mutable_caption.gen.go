// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)





// The class instance for the [MutableCaption] class.
var (
	MutableCaptionClass     _MutableCaptionClass
	MutableCaptionClassOnce sync.Once
)

func getMutableCaptionClass() _MutableCaptionClass {
	MutableCaptionClassOnce.Do(func() {
		MutableCaptionClass = _MutableCaptionClass{objc.GetClass("AVMutableCaption")}
	})
	return MutableCaptionClass
}

type _MutableCaptionClass struct {
	class objc.Class
}





// An interface definition for the [MutableCaption] class.
type IMutableCaption interface {
	ICaption
	

	// properties:
	Animation() CaptionAnimation
	SetAnimation(value CaptionAnimation)
	Region() IAVCaptionRegion
	SetRegion(value IAVCaptionRegion)
	Text() objc.IObject /* cross-framework: NSString */
	SetText(value objc.IObject /* cross-framework: NSString */)
	TextAlignment() CaptionTextAlignment
	SetTextAlignment(value CaptionTextAlignment)
	TimeRange() TimeRange /* not a class type */
	SetTimeRange(value TimeRange /* not a class type */)


	

	// methods:
	RemoveBackgroundColorInRange(range_ corefoundation.Range)
	RemoveDecorationInRange(range_ corefoundation.Range)
	RemoveFontStyleInRange(range_ corefoundation.Range)
	RemoveFontWeightInRange(range_ corefoundation.Range)
	RemoveRubyInRange(range_ corefoundation.Range)
	RemoveTextColorInRange(range_ corefoundation.Range)
	RemoveTextCombineInRange(range_ corefoundation.Range)
	SetBackgroundColorInRange(color ColorRef /* not a class type */, range_ corefoundation.Range)
	SetDecorationInRange(decoration CaptionDecoration, range_ corefoundation.Range)
	SetFontStyleInRange(fontStyle CaptionFontStyle, range_ corefoundation.Range)
	SetFontWeightInRange(fontWeight CaptionFontWeight, range_ corefoundation.Range)
	SetRubyInRange(ruby IAVCaptionRuby, range_ corefoundation.Range)
	SetTextColorInRange(color ColorRef /* not a class type */, range_ corefoundation.Range)
	SetTextCombineInRange(textCombine CaptionTextCombine, range_ corefoundation.Range)


}





// Alloc allocates a new instance without initialization.
func (mc _MutableCaptionClass) Alloc() MutableCaption {
	rv := objc.Send[MutableCaption](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableCaptionClass) New() MutableCaption {
	rv := objc.Send[MutableCaption](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableCaption) Init() MutableCaption {
	rv := objc.Send[MutableCaption](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableCaption) Autorelease() MutableCaption {
	rv := objc.Send[MutableCaption](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableCaption creates a new MutableCaption instance.
func NewMutableCaption() MutableCaption {
	return getMutableCaptionClass().New()
}





// A mutable caption subclass that you use to create new captions.


// A mutable caption subclass that you use to create new captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption
type MutableCaption struct {
	Caption
}

// MutableCaptionFrom constructs a [MutableCaption] from an unsafe.Pointer.
//
// A mutable caption subclass that you use to create new captions.
func MutableCaptionFrom(ptr unsafe.Pointer) MutableCaption {
	return MutableCaption{
		Caption: CaptionFrom(ptr),
	}
}




















// Removes a background color from a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeBackgroundColorInRange:
func (m_ MutableCaption) RemoveBackgroundColorInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeBackgroundColorInRange:"), range_)
}


// Removes a decoration from a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeDecorationInRange:
func (m_ MutableCaption) RemoveDecorationInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeDecorationInRange:"), range_)
}


// Removes a font style from a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeFontStyleInRange:
func (m_ MutableCaption) RemoveFontStyleInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeFontStyleInRange:"), range_)
}


// Removes a font weight from a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeFontWeightInRange:
func (m_ MutableCaption) RemoveFontWeightInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeFontWeightInRange:"), range_)
}


// Removes ruby text from a range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeRubyInRange:
func (m_ MutableCaption) RemoveRubyInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeRubyInRange:"), range_)
}


// Removes the text color for a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeTextColorInRange:
func (m_ MutableCaption) RemoveTextColorInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeTextColorInRange:"), range_)
}


// Removes text combine from a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeTextCombineInRange:
func (m_ MutableCaption) RemoveTextCombineInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeTextCombineInRange:"), range_)
}


// Sets the background color for a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setBackgroundColor:inRange:
func (m_ MutableCaption) SetBackgroundColorInRange(color ColorRef /* not a class type */, range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackgroundColor:inRange:"), color, range_)
}


// Sets a decoration for a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setDecoration:inRange:
func (m_ MutableCaption) SetDecorationInRange(decoration CaptionDecoration, range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDecoration:inRange:"), decoration, range_)
}


// Sets the font style for a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setFontStyle:inRange:
func (m_ MutableCaption) SetFontStyleInRange(fontStyle CaptionFontStyle, range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFontStyle:inRange:"), fontStyle, range_)
}


// Sets the font weight for a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setFontWeight:inRange:
func (m_ MutableCaption) SetFontWeightInRange(fontWeight CaptionFontWeight, range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFontWeight:inRange:"), fontWeight, range_)
}


// Sets ruby text for a range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setRuby:inRange:
func (m_ MutableCaption) SetRubyInRange(ruby IAVCaptionRuby, range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRuby:inRange:"), ruby, range_)
}


// Sets the text color for a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setTextColor:inRange:
func (m_ MutableCaption) SetTextColorInRange(color ColorRef /* not a class type */, range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTextColor:inRange:"), color, range_)
}


// Sets text combine for a range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setTextCombine:inRange:
func (m_ MutableCaption) SetTextCombineInRange(textCombine CaptionTextCombine, range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTextCombine:inRange:"), textCombine, range_)
}







// Animations to apply to the caption text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/animation
func (m_ MutableCaption) Animation() CaptionAnimation {
	rv := objc.Send[CaptionAnimation](m_.ID, objc.Sel("animation"))
	return rv
}


// Animations to apply to the caption text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/animation
func (m_ MutableCaption) SetAnimation(value CaptionAnimation) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnimation:"), value)
}


// The region in which the caption exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/region
func (m_ MutableCaption) Region() IAVCaptionRegion {
	rv := objc.Send[CaptionRegion](m_.ID, objc.Sel("region"))
	return rv
}


// The region in which the caption exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/region
func (m_ MutableCaption) SetRegion(value IAVCaptionRegion) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}


// The caption text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/text
func (m_ MutableCaption) Text() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("text"))
	return rv
}


// The caption text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/text
func (m_ MutableCaption) SetText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setText:"), value)
}


// The alignment of the caption text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/textAlignment
func (m_ MutableCaption) TextAlignment() CaptionTextAlignment {
	rv := objc.Send[CaptionTextAlignment](m_.ID, objc.Sel("textAlignment"))
	return rv
}


// The alignment of the caption text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/textAlignment
func (m_ MutableCaption) SetTextAlignment(value CaptionTextAlignment) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTextAlignment:"), value)
}


// The time range over which the system presents the caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/timeRange
func (m_ MutableCaption) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](m_.ID, objc.Sel("timeRange"))
	return rv
}


// The time range over which the system presents the caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/timeRange
func (m_ MutableCaption) SetTimeRange(value TimeRange /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeRange:"), value)
}








