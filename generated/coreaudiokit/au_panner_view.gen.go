// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

// The class instance for the [PannerView] class.
var (
	PannerViewClass     _PannerViewClass
	PannerViewClassOnce sync.Once
)

func getPannerViewClass() _PannerViewClass {
	PannerViewClassOnce.Do(func() {
		PannerViewClass = _PannerViewClass{objc.GetClass("AUPannerView")}
	})
	return PannerViewClass
}

type _PannerViewClass struct {
	class objc.Class
}

// An interface definition for the [PannerView] class.
type IPannerView interface {
	appkit.IView
	AudioUnit() audiotoolbox.AudioUnit
}

// A view that provides a specialized user interface for a Cocoa-based panner audio unit.


// A view that provides a specialized user interface for a Cocoa-based panner audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUPannerView
type PannerView struct {
	appkit.View
}

// PannerViewFrom constructs a [PannerView] from an unsafe.Pointer.
//
// A view that provides a specialized user interface for a Cocoa-based panner audio unit.
func PannerViewFrom(ptr unsafe.Pointer) PannerView {
	return PannerView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PannerViewClass) Alloc() PannerView {
	rv := objc.Send[PannerView](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PannerViewClass) New() PannerView {
	rv := objc.Send[PannerView](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PannerView) Init() PannerView {
	rv := objc.Send[PannerView](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PannerView) Autorelease() PannerView {
	rv := objc.Send[PannerView](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPannerView creates a new PannerView instance.
func NewPannerView() PannerView {
	return getPannerViewClass().New()
}



// Creates a panner view for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUPannerView/init(audioUnit:)
func NewPannerViewAUPannerViewWithAudioUnit(au audiotoolbox.IAudioUnit) PannerView {
	rv := objc.Send[PannerView](objc.ID(getPannerViewClass().class), objc.Sel("AUPannerViewWithAudioUnit:"), au)
	return rv
}



// Creates a panner view for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUPannerView/init(audioUnit:)
func (pc _PannerViewClass) AUPannerViewWithAudioUnit(au audiotoolbox.IAudioUnit) PannerView {
	rv := objc.Send[PannerView](objc.ID(pc.class), objc.Sel("AUPannerViewWithAudioUnit:"), au)
	return rv
}


// The panner audio unit associated with the generic panner view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUPannerView/audioUnit
func (p_ PannerView) AudioUnit() audiotoolbox.AudioUnit {
	rv := objc.Send[audiotoolbox.AudioUnit](p_.ID, objc.Sel("audioUnit"))
	return rv
}


