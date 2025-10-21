// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

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

// An interface definition for the [GenericView] class.
type IGenericView interface {
	appkit.IView
}

// A view that provides a generic user interface for a Cocoa audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericView
type GenericView struct {
	appkit.View
}

// GenericViewFrom constructs a [GenericView] from an unsafe.Pointer.
//
// A view that provides a generic user interface for a Cocoa audio unit.
func GenericViewFrom(ptr unsafe.Pointer) GenericView {
	return GenericView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GenericViewClass) Alloc() GenericView {
	rv := objc.Send[GenericView](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a generic view for an audio unit, setting all display flags.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericView/init(audioUnit:)
func NewGenericViewWithAudioUnit(au audiotoolbox.IAudioUnit) GenericView {
	instance := getGenericViewClass().Alloc()
	rv := objc.Send[GenericView](instance.ID, objc.Sel("initWithAudioUnit:"), au)
	rv.Autorelease()
	return rv
}



// Initializes a generic view for an audio unit, setting specific display flags.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericView/init(audioUnit:displayFlags:)
func NewGenericViewWithAudioUnitDisplayFlags(inAudioUnit audiotoolbox.IAudioUnit, inFlags GenericViewDisplayFlags) GenericView {
	instance := getGenericViewClass().Alloc()
	rv := objc.Send[GenericView](instance.ID, objc.Sel("initWithAudioUnit:displayFlags:"), inAudioUnit, inFlags)
	rv.Autorelease()
	return rv
}


// The audio unit associated with the generic view.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericView/audioUnit
func (g_ GenericView) AudioUnit() audiotoolbox.AudioUnit {
	rv := objc.Send[audiotoolbox.AudioUnit](g_.ID, objc.Sel("audioUnit"))
	return rv
}

// Indicates whether or not controls for expert audio unit parameters are displayed in the generic view.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericView/showsExpertParameters
func (g_ GenericView) ShowsExpertParameters() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("showsExpertParameters"))
	return rv
}


// SetShowsExpertParameters sets the value of the showsExpertParameters property.
// Indicates whether or not controls for expert audio unit parameters are displayed in the generic view.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericView/showsExpertParameters
func (g_ GenericView) SetShowsExpertParameters(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShowsExpertParameters:"), value)
}


