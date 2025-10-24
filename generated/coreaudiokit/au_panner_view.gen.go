// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

/* debug [class.gen.go]: Generating class AUPannerView */


/* debug [class_header]: Header for AUPannerView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PannerView */
// An interface definition for the [PannerView] class.
type IPannerView interface {
	IView
	
/* debug [class_interface_properties]: Properties for PannerView */
	// properties:
	AudioUnit() audiotoolbox.AudioUnit
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PannerView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PannerView */
// Alloc allocates a new instance without initialization.
func (pc _PannerViewClass) Alloc() PannerView {
	rv := objc.Send[PannerView](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PannerView */
// A view that provides a specialized user interface for a Cocoa-based panner audio unit.


// A view that provides a specialized user interface for a Cocoa-based panner audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUPannerView
type PannerView struct {
	View
}

// PannerViewFrom constructs a [PannerView] from an unsafe.Pointer.
//
// A view that provides a specialized user interface for a Cocoa-based panner audio unit.
func PannerViewFrom(ptr unsafe.Pointer) PannerView {
	return PannerView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PannerView */

// Creates a panner view for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUPannerView/init(audioUnit:)
func NewPannerViewAUPannerViewWithAudioUnit(au audiotoolbox.AudioUnit) PannerView {
	rv := objc.Send[PannerView](objc.ID(getPannerViewClass().class), objc.Sel("AUPannerViewWithAudioUnit:"), au)
	return rv
}/* debug [class_init_methods/constructor]: NewPannerViewAUPannerViewWithAudioUnit */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PannerView */

// Creates a panner view for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUPannerView/init(audioUnit:)
func (pc _PannerViewClass) AUPannerViewWithAudioUnit(au audiotoolbox.AudioUnit) IPannerView {
	rv := objc.Send[PannerView](objc.ID(pc.class), objc.Sel("AUPannerViewWithAudioUnit:"), au)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AUPannerViewWithAudioUnit) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PannerView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PannerView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PannerView */

// The panner audio unit associated with the generic panner view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUPannerView/audioUnit
func (p_ PannerView) AudioUnit() audiotoolbox.AudioUnit {
	rv := objc.Send[audiotoolbox.AudioUnit](p_.ID, objc.Sel("audioUnit"))
	return rv
}/* debug [instance_properties/getter]: audioUnit */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AUPannerView */


