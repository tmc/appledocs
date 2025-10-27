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
	

	// properties:
	IsShowingAppNames() bool
	SetIsShowingAppNames(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _InterAppAudioSwitcherViewClass) Alloc() InterAppAudioSwitcherView {
	rv := objc.Send[InterAppAudioSwitcherView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A view that provides an audio switcher user interface.


// A view that provides an audio switcher user interface.
//
// [Full Topic]
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

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudioswitcherview/isshowingappnames
func (i_ InterAppAudioSwitcherView) IsShowingAppNames() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isShowingAppNames"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudioswitcherview/isshowingappnames
func (i_ InterAppAudioSwitcherView) SetIsShowingAppNames(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsShowingAppNames:"), value)
}







