// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

/* debug [class.gen.go]: Generating class CAInterAppAudioSwitcherView */


/* debug [class_header]: Header for CAInterAppAudioSwitcherView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InterAppAudioSwitcherView */
// An interface definition for the [InterAppAudioSwitcherView] class.
type IInterAppAudioSwitcherView interface {
	IView
	
/* debug [class_interface_properties]: Properties for InterAppAudioSwitcherView */
	// properties:
	IsShowingAppNames() bool
	SetIsShowingAppNames(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InterAppAudioSwitcherView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InterAppAudioSwitcherView */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InterAppAudioSwitcherView */
// A view that provides an audio switcher user interface.


// A view that provides an audio switcher user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioSwitcherView
type InterAppAudioSwitcherView struct {
	View
}

// InterAppAudioSwitcherViewFrom constructs a [InterAppAudioSwitcherView] from an unsafe.Pointer.
//
// A view that provides an audio switcher user interface.
func InterAppAudioSwitcherViewFrom(ptr unsafe.Pointer) InterAppAudioSwitcherView {
	return InterAppAudioSwitcherView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InterAppAudioSwitcherView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InterAppAudioSwitcherView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InterAppAudioSwitcherView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InterAppAudioSwitcherView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InterAppAudioSwitcherView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudioswitcherview/isshowingappnames
func (i_ InterAppAudioSwitcherView) IsShowingAppNames() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isShowingAppNames"))
	return rv
}/* debug [instance_properties/getter]: isShowingAppNames */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudioswitcherview/isshowingappnames
func (i_ InterAppAudioSwitcherView) SetIsShowingAppNames(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsShowingAppNames:"), value)
}/* debug [instance_properties/setter]: isShowingAppNames */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAInterAppAudioSwitcherView */


