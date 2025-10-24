// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

/* debug [class.gen.go]: Generating class AUGenericView */


/* debug [class_header]: Header for AUGenericView */
// The class instance for the [GenericView] class.
var (
	GenericViewClass     _GenericViewClass
	GenericViewClassOnce sync.Once
)

func getGenericViewClass() _GenericViewClass {
	GenericViewClassOnce.Do(func() {
		GenericViewClass = _GenericViewClass{objc.GetClass("AUGenericView")}
	})
	return GenericViewClass
}

type _GenericViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GenericView */
// An interface definition for the [GenericView] class.
type IGenericView interface {
	IView
	
/* debug [class_interface_properties]: Properties for GenericView */
	// properties:
	AudioUnit() audiotoolbox.AudioUnit
	ShowsExpertParameters() bool
	SetShowsExpertParameters(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GenericView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GenericView */
// Alloc allocates a new instance without initialization.
func (gc _GenericViewClass) Alloc() GenericView {
	rv := objc.Send[GenericView](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GenericViewClass) New() GenericView {
	rv := objc.Send[GenericView](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GenericView) Init() GenericView {
	rv := objc.Send[GenericView](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GenericView) Autorelease() GenericView {
	rv := objc.Send[GenericView](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGenericView creates a new GenericView instance.
func NewGenericView() GenericView {
	return getGenericViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GenericView */
// A view that provides a generic user interface for a Cocoa audio unit.


// A view that provides a generic user interface for a Cocoa audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericView
type GenericView struct {
	View
}

// GenericViewFrom constructs a [GenericView] from an unsafe.Pointer.
//
// A view that provides a generic user interface for a Cocoa audio unit.
func GenericViewFrom(ptr unsafe.Pointer) GenericView {
	return GenericView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GenericView */

// Creates a generic view for an audio unit, setting all display flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericView/init(audioUnit:)
func NewGenericViewWithAudioUnit(au audiotoolbox.AudioUnit) GenericView {
	instance := getGenericViewClass().Alloc()
	rv := objc.Send[GenericView](instance.ID, objc.Sel("initWithAudioUnit:"), au)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGenericViewWithAudioUnit */


// Initializes a generic view for an audio unit, setting specific display flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericView/init(audioUnit:displayFlags:)
func NewGenericViewWithAudioUnitDisplayFlags(inAudioUnit audiotoolbox.AudioUnit, inFlags GenericViewDisplayFlags) GenericView {
	instance := getGenericViewClass().Alloc()
	rv := objc.Send[GenericView](instance.ID, objc.Sel("initWithAudioUnit:displayFlags:"), inAudioUnit, inFlags)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGenericViewWithAudioUnitDisplayFlags */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GenericView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GenericView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GenericView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GenericView */

// The audio unit associated with the generic view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericView/audioUnit
func (g_ GenericView) AudioUnit() audiotoolbox.AudioUnit {
	rv := objc.Send[audiotoolbox.AudioUnit](g_.ID, objc.Sel("audioUnit"))
	return rv
}/* debug [instance_properties/getter]: audioUnit */


// Indicates whether or not controls for expert audio unit parameters are displayed in the generic view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericView/showsExpertParameters
func (g_ GenericView) ShowsExpertParameters() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("showsExpertParameters"))
	return rv
}/* debug [instance_properties/getter]: showsExpertParameters */


// Indicates whether or not controls for expert audio unit parameters are displayed in the generic view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericView/showsExpertParameters
func (g_ GenericView) SetShowsExpertParameters(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShowsExpertParameters:"), value)
}/* debug [instance_properties/setter]: showsExpertParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AUGenericView */


