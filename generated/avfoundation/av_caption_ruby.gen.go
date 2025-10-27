// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptionRuby] class.
var (
	CaptionRubyClass     _CaptionRubyClass
	CaptionRubyClassOnce sync.Once
)

func getCaptionRubyClass() _CaptionRubyClass {
	CaptionRubyClassOnce.Do(func() {
		CaptionRubyClass = _CaptionRubyClass{objc.GetClass("AVCaptionRuby")}
	})
	return CaptionRubyClass
}

type _CaptionRubyClass struct {
	class objc.Class
}





// An interface definition for the [CaptionRuby] class.
type ICaptionRuby interface {
	objectivec.IObject
	

	// properties:
	Alignment() CaptionRubyAlignment
	Position() CaptionRubyPosition
	Text() foundation.foundation.INSString


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptionRubyClass) Alloc() CaptionRuby {
	rv := objc.Send[CaptionRuby](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionRubyClass) New() CaptionRuby {
	rv := objc.Send[CaptionRuby](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptionRuby) Init() CaptionRuby {
	rv := objc.Send[CaptionRuby](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptionRuby) Autorelease() CaptionRuby {
	rv := objc.Send[CaptionRuby](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptionRuby creates a new CaptionRuby instance.
func NewCaptionRuby() CaptionRuby {
	return getCaptionRubyClass().New()
}





// An object that presents ruby characters.
//
// Ruby characters are small annotations, typically used in Japanese content, that render alongside the base text.


// An object that presents ruby characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby
type CaptionRuby struct {
	objectivec.Object
}

// CaptionRubyFrom constructs a [CaptionRuby] from an unsafe.Pointer.
//
// An object that presents ruby characters.
func CaptionRubyFrom(ptr unsafe.Pointer) CaptionRuby {
	return CaptionRuby{objectivec.Object{objc.ID(ptr)}}
}






// Creates ruby text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/init(text:)
func NewCaptionRubyWithText(text foundation.foundation.INSString) CaptionRuby {
	instance := getCaptionRubyClass().Alloc()
	rv := objc.Send[CaptionRuby](instance.ID, objc.Sel("initWithText:"), text)
	rv.Autorelease()
	return rv
}


// Creates ruby text with position and alignment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/init(text:position:alignment:)
func NewCaptionRubyWithTextPositionAlignment(text foundation.foundation.INSString, position CaptionRubyPosition, alignment CaptionRubyAlignment) CaptionRuby {
	instance := getCaptionRubyClass().Alloc()
	rv := objc.Send[CaptionRuby](instance.ID, objc.Sel("initWithText:position:alignment:"), text, position, alignment)
	rv.Autorelease()
	return rv
}






















// The ruby text alignment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/alignment
func (c_ CaptionRuby) Alignment() CaptionRubyAlignment {
	rv := objc.Send[CaptionRubyAlignment](c_.ID, objc.Sel("alignment"))
	return rv
}


// The ruby text position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/position
func (c_ CaptionRuby) Position() CaptionRubyPosition {
	rv := objc.Send[CaptionRubyPosition](c_.ID, objc.Sel("position"))
	return rv
}


// The ruby text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/text
func (c_ CaptionRuby) Text() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("text"))
	return rv
}







