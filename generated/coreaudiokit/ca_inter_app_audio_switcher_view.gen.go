// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [InterAppAudioSwitcherView] class.
var (
	InterAppAudioSwitcherViewClass     _InterAppAudioSwitcherViewClass
	InterAppAudioSwitcherViewClassOnce sync.Once
)

func getInterAppAudioSwitcherViewClass() _InterAppAudioSwitcherViewClass {
	InterAppAudioSwitcherViewClassOnce.Do(func() {
		InterAppAudioSwitcherViewClass = _InterAppAudioSwitcherViewClass{objc.GetClass("CAInterAppAudioSwitcherView")}
	})
	return InterAppAudioSwitcherViewClass
}

type _InterAppAudioSwitcherViewClass struct {
	class objc.Class
}

// An interface definition for the [InterAppAudioSwitcherView] class.
type IInterAppAudioSwitcherView interface {
	appkit.IView
	ContentWidth() float64
	SetOutputAudioUnit(au unsafe.Pointer)
}

// A view that provides an audio switcher user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioSwitcherView
type InterAppAudioSwitcherView struct {
	appkit.View
}

// InterAppAudioSwitcherViewFrom constructs a [InterAppAudioSwitcherView] from an unsafe.Pointer.
//
// A view that provides an audio switcher user interface.
func InterAppAudioSwitcherViewFrom(ptr unsafe.Pointer) InterAppAudioSwitcherView {
	return InterAppAudioSwitcherView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _InterAppAudioSwitcherViewClass) Alloc() InterAppAudioSwitcherView {
	rv := objc.Send[InterAppAudioSwitcherView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InterAppAudioSwitcherViewClass) New() InterAppAudioSwitcherView {
	rv := objc.Send[InterAppAudioSwitcherView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InterAppAudioSwitcherView) Init() InterAppAudioSwitcherView {
	rv := objc.Send[InterAppAudioSwitcherView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InterAppAudioSwitcherView) Autorelease() InterAppAudioSwitcherView {
	rv := objc.Send[InterAppAudioSwitcherView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInterAppAudioSwitcherView creates a new InterAppAudioSwitcherView instance.
func NewInterAppAudioSwitcherView() InterAppAudioSwitcherView {
	return getInterAppAudioSwitcherViewClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioSwitcherView/contentWidth()
func (i_ InterAppAudioSwitcherView) ContentWidth() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("contentWidth"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioSwitcherView/setOutputAudioUnit(_:)
func (i_ InterAppAudioSwitcherView) SetOutputAudioUnit(au unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOutputAudioUnit:"), au)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioSwitcherView/isShowingAppNames
func (i_ InterAppAudioSwitcherView) ShowingAppNames() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("showingAppNames"))
	return rv
}


// SetShowingAppNames sets the value of the showingAppNames property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioSwitcherView/isShowingAppNames
func (i_ InterAppAudioSwitcherView) SetShowingAppNames(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setShowingAppNames:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudioswitcherview/isshowingappnames
func (i_ InterAppAudioSwitcherView) IsShowingAppNames() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isShowingAppNames"))
	return rv
}


// SetIsShowingAppNames sets the value of the isShowingAppNames property.
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudioswitcherview/isshowingappnames
func (i_ InterAppAudioSwitcherView) SetIsShowingAppNames(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsShowingAppNames:"), value)
}



