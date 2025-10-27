// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ScrubberSelectionStyle] class.
var (
	ScrubberSelectionStyleClass     _ScrubberSelectionStyleClass
	ScrubberSelectionStyleClassOnce sync.Once
)

func getScrubberSelectionStyleClass() _ScrubberSelectionStyleClass {
	ScrubberSelectionStyleClassOnce.Do(func() {
		ScrubberSelectionStyleClass = _ScrubberSelectionStyleClass{objc.GetClass("NSScrubberSelectionStyle")}
	})
	return ScrubberSelectionStyleClass
}

type _ScrubberSelectionStyleClass struct {
	class objc.Class
}





// An interface definition for the [ScrubberSelectionStyle] class.
type IScrubberSelectionStyle interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	MakeSelectionView() IScrubberSelectionView


}





// Alloc allocates a new instance without initialization.
func (sc _ScrubberSelectionStyleClass) Alloc() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberSelectionStyleClass) New() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberSelectionStyle) Init() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberSelectionStyle) Autorelease() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberSelectionStyle creates a new ScrubberSelectionStyle instance.
func NewScrubberSelectionStyle() ScrubberSelectionStyle {
	return getScrubberSelectionStyleClass().New()
}





// An abstract class that provides decorative accessory views for selected and highlighted items within a scrubber control.
//
// Choose a selection style ( or ), or create a custom selection style by subclassing and overriding .


// An abstract class that provides decorative accessory views for selected and highlighted items within a scrubber control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle
type ScrubberSelectionStyle struct {
	objectivec.Object
}

// ScrubberSelectionStyleFrom constructs a [ScrubberSelectionStyle] from an unsafe.Pointer.
//
// An abstract class that provides decorative accessory views for selected and highlighted items within a scrubber control.
func ScrubberSelectionStyleFrom(ptr unsafe.Pointer) ScrubberSelectionStyle {
	return ScrubberSelectionStyle{objectivec.Object{objc.ID(ptr)}}
}






// Initializes a scrubber selection style when included from a nib or Storyboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/init(coder:)
func NewScrubberSelectionStyleWithCoder(coder foundation.foundation.INSCoder) ScrubberSelectionStyle {
	instance := getScrubberSelectionStyleClass().Alloc()
	rv := objc.Send[ScrubberSelectionStyle](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}












// A built-in selection style that draws the outline of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/outlineOverlay
func (sc _ScrubberSelectionStyleClass) OutlineOverlayStyle() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](objc.ID(sc.class), objc.Sel("outlineOverlayStyle"))
	return rv
}

// A built-in selection style that draws a rounded rectangle as the background of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/roundedBackground
func (sc _ScrubberSelectionStyleClass) RoundedBackgroundStyle() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](objc.ID(sc.class), objc.Sel("roundedBackgroundStyle"))
	return rv
}






// Provides an opportunity to create a customized scrubber selection style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/makeSelectionView()
func (s_ ScrubberSelectionStyle) MakeSelectionView() IScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](s_.ID, objc.Sel("makeSelectionView"))
	return rv
}







// A built-in selection style that draws the outline of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/outlineOverlay
func (s_ ScrubberSelectionStyle) OutlineOverlayStyle() IScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID, objc.Sel("outlineOverlayStyle"))
	return rv
}


// A built-in selection style that draws a rounded rectangle as the background of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/roundedBackground
func (s_ ScrubberSelectionStyle) RoundedBackgroundStyle() IScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID, objc.Sel("roundedBackgroundStyle"))
	return rv
}







