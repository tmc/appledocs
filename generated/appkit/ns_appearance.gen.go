// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Appearance] class.
var (
	AppearanceClass     _AppearanceClass
	AppearanceClassOnce sync.Once
)

func getAppearanceClass() _AppearanceClass {
	AppearanceClassOnce.Do(func() {
		AppearanceClass = _AppearanceClass{objc.GetClass("NSAppearance")}
	})
	return AppearanceClass
}

type _AppearanceClass struct {
	class objc.Class
}

// An interface definition for the [Appearance] class.
type IAppearance interface {
	objectivec.IObject
	BestMatchFromAppearancesWithNames(appearances unsafe.Pointer) unsafe.Pointer
}

// An object that manages standard appearance attributes for UI elements in an app.
//
// An object manages how AppKit renders your app’s UI elements. Specifically, appearance objects determine which colors and images AppKit uses when drawing windows, views, and controls. Although you can use an appearance object to determine how to draw custom views and controls, a better approach is to choose colors and images that adapt automatically to the current appearance. For example, define a color asset whose actual color value changes for light and dark appearances. You can assign specific appearances to your views in Interface Builder. The user chooses the default appearance for the system, but you can override that appearance for all or part of your app. Apps inherit the default system appearance, windows inherit their app’s appearance, and views inherit the appearance of their nearest ancestor (either a superview or window). To force a window or view to adopt an appearance, assign a specific appearance object to its property. When AppKit draws a control, it automatically sets the current appearance on the current thread to the control’s appearance. The current appearance influences the drawing path and return values you get when you access system fonts and colors. The current appearance also affects the appearance of text and images, such as the text and template images in a toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance
type Appearance struct {
	objectivec.Object
}

// AppearanceFrom constructs a [Appearance] from an unsafe.Pointer.
//
// An object that manages standard appearance attributes for UI elements in an app.
func AppearanceFrom(ptr unsafe.Pointer) Appearance {
	return Appearance{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AppearanceClass) Alloc() Appearance {
	rv := objc.Send[Appearance](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AppearanceClass) New() Appearance {
	rv := objc.Send[Appearance](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Appearance) Init() Appearance {
	rv := objc.Send[Appearance](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Appearance) Autorelease() Appearance {
	rv := objc.Send[Appearance](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAppearance creates a new Appearance instance.
func NewAppearance() Appearance {
	return getAppearanceClass().New()
}




// Creates an appearance object based on the name of one of the standard system appearances.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/init(named:)
func NewAppearanceNamed(name unsafe.Pointer) Appearance {
	rv := objc.Send[Appearance](objc.ID(getAppearanceClass().class), objc.Sel("appearanceNamed:"), name)
	return rv
}



// Creates an appearance object from the named appearance file located in the specified bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/init(appearanceNamed:bundle:)
func NewAppearanceWithAppearanceNamedBundle(name unsafe.Pointer, bundle unsafe.Pointer) Appearance {
	instance := getAppearanceClass().Alloc()
	rv := objc.Send[Appearance](instance.ID, objc.Sel("initWithAppearanceNamed:bundle:"), name, bundle)
	rv.Autorelease()
	return rv
}


// Creates an appearance object based on the name of one of the standard system appearances.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/init(named:)
func (ac _AppearanceClass) AppearanceNamed(name unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("appearanceNamed:"), name)
	return rv
}

// Returns the appearance name that most closely matches the current appearance object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/bestMatch(from:)
func (a_ Appearance) BestMatchFromAppearancesWithNames(appearances unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("bestMatchFromAppearancesWithNames:"), appearances)
	return rv
}

// The name of the appearance.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/name-swift.property
func (a_ Appearance) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("name"))
	return rv
}


