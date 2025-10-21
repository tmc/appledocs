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
	MakeSelectionView() ScrubberSelectionView
}

// An abstract class that provides decorative accessory views for selected and highlighted items within a scrubber control.
//
// Choose a selection style ( or ), or create a custom selection style by subclassing and overriding .
//
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

// Alloc allocates a new instance without initialization.
func (sc _ScrubberSelectionStyleClass) Alloc() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a scrubber selection style when included from a nib or Storyboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/init(coder:)
func NewScrubberSelectionStyleWithCoder(coder ICoder) ScrubberSelectionStyle {
	instance := getScrubberSelectionStyleClass().Alloc()
	rv := objc.Send[ScrubberSelectionStyle](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// A built-in selection style that draws the outline of the scrubber item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/outlineOverlay
func (sc _ScrubberSelectionStyleClass) OutlineOverlayStyle() NSScrubberSelectionStyle {
	rv := objc.Send[NSScrubberSelectionStyle](objc.ID(sc.class), objc.Sel("outlineOverlayStyle"))
	return rv
}
// Provides an opportunity to create a customized scrubber selection style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/makeSelectionView()
func (s_ ScrubberSelectionStyle) MakeSelectionView() ScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](s_.ID, objc.Sel("makeSelectionView"))
	return rv
}

// A built-in selection style that draws the outline of the scrubber item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/outlineOverlay
func (s_ ScrubberSelectionStyle) OutlineOverlayStyle() NSScrubberSelectionStyle {
	rv := objc.Send[NSScrubberSelectionStyle](s_.ID, objc.Sel("outlineOverlayStyle"))
	return rv
}


