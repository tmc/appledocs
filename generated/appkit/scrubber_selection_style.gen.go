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
	scrubberSelectionStyleClass     _ScrubberSelectionStyleClass
	scrubberSelectionStyleClassOnce sync.Once
)

func getScrubberSelectionStyleClass() _ScrubberSelectionStyleClass {
	scrubberSelectionStyleClassOnce.Do(func() {
		scrubberSelectionStyleClass = _ScrubberSelectionStyleClass{objc.GetClass("NSScrubberSelectionStyle")}
	})
	return scrubberSelectionStyleClass
}

type _ScrubberSelectionStyleClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberSelectionStyle] class.
type IScrubberSelectionStyle interface {
	objectivec.IObject
}

// An abstract class that provides decorative accessory views for selected and highlighted items within a scrubber control. [Full Topic]
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




