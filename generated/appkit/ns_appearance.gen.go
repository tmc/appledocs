// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAppearance */


/* debug [class_header]: Header for NSAppearance */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Appearance */
// An interface definition for the [Appearance] class.
type IAppearance interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Appearance */
	// properties:
	AllowsVibrancy() bool
	Name() AppearanceName /* typedef */
	Appearance() IAppearance
	SetAppearance(value IAppearance)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Appearance */
	// methods:
	BestMatchFromAppearancesWithNames(appearances []string) AppearanceName /* typedef */
	PerformAsCurrentDrawingAppearance(block unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Appearance */
// Alloc allocates a new instance without initialization.
func (ac _AppearanceClass) Alloc() Appearance {
	rv := objc.Send[Appearance](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Appearance */
// An object that manages standard appearance attributes for UI elements in an app.
//
// An object manages how AppKit renders your app’s UI elements. Specifically, appearance objects determine which colors and images AppKit uses when drawing windows, views, and controls. Although you can use an appearance object to determine how to draw custom views and controls, a better approach is to choose colors and images that adapt automatically to the current appearance. For example, define a color asset whose actual color value changes for light and dark appearances. You can assign specific appearances to your views in Interface Builder. The user chooses the default appearance for the system, but you can override that appearance for all or part of your app. Apps inherit the default system appearance, windows inherit their app’s appearance, and views inherit the appearance of their nearest ancestor (either a superview or window). To force a window or view to adopt an appearance, assign a specific appearance object to its property. When AppKit draws a control, it automatically sets the current appearance on the current thread to the control’s appearance. The current appearance influences the drawing path and return values you get when you access system fonts and colors. The current appearance also affects the appearance of text and images, such as the text and template images in a toolbar.


// An object that manages standard appearance attributes for UI elements in an app.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Appearance */

// Creates an appearance object based on the name of one of the standard system appearances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/init(named:)
func NewAppearanceNamed(name AppearanceName /* typedef */) Appearance {
	rv := objc.Send[Appearance](objc.ID(getAppearanceClass().class), objc.Sel("appearanceNamed:"), name)
	return rv
}/* debug [class_init_methods/constructor]: NewAppearanceNamed */


// Creates an appearance object from the named appearance file located in the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/init(appearanceNamed:bundle:)
func NewAppearanceWithAppearanceNamedBundle(name AppearanceName /* typedef */, bundle foundation.Bundle) Appearance {
	instance := getAppearanceClass().Alloc()
	rv := objc.Send[Appearance](instance.ID, objc.Sel("initWithAppearanceNamed:bundle:"), name, bundle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAppearanceWithAppearanceNamedBundle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/init(coder:)
func NewAppearanceWithCoder(coder foundation.Coder) Appearance {
	instance := getAppearanceClass().Alloc()
	rv := objc.Send[Appearance](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAppearanceWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Appearance */

// The appearance that the system uses for color and asset resolution, and that’s active for drawing, usually from locking focus on a view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/currentDrawing()
func (ac _AppearanceClass) CurrentDrawing() {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("currentDrawing"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CurrentDrawing) */


// Creates an appearance object based on the name of one of the standard system appearances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/init(named:)
func (ac _AppearanceClass) AppearanceNamed(name AppearanceName /* typedef */) IAppearance {
	rv := objc.Send[Appearance](objc.ID(ac.class), objc.Sel("appearanceNamed:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AppearanceNamed) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Appearance */

// Returns the appearance object that’s active on the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/current
func (ac _AppearanceClass) CurrentAppearance() Appearance {
	rv := objc.Send[Appearance](objc.ID(ac.class), objc.Sel("currentAppearance"))
	return rv
}/* debug [class_properties_class/property]: currentAppearance */

// The appearance that the system uses for color and asset resolution, and that’s active for drawing, usually from locking focus on a view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/currentDrawingAppearance
func (ac _AppearanceClass) CurrentDrawingAppearance() Appearance {
	rv := objc.Send[Appearance](objc.ID(ac.class), objc.Sel("currentDrawingAppearance"))
	return rv
}/* debug [class_properties_class/property]: currentDrawingAppearance */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Appearance */

// Returns the appearance name that most closely matches the current appearance object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/bestMatch(from:)
func (a_ Appearance) BestMatchFromAppearancesWithNames(appearances []string) AppearanceName /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("bestMatchFromAppearancesWithNames:"), appearances)
	return rv
}/* debug [instance_methods/method]: BestMatchFromAppearancesWithNames */


// Sets the appearance to be the active drawing appearance and perform the specified block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/performAsCurrentDrawingAppearance(_:)
func (a_ Appearance) PerformAsCurrentDrawingAppearance(block unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("performAsCurrentDrawingAppearance:"), block)
}/* debug [instance_methods/method]: PerformAsCurrentDrawingAppearance */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Appearance */

// Specifies whether the current appearance allows vibrancy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/allowsVibrancy
func (a_ Appearance) AllowsVibrancy() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsVibrancy"))
	return rv
}/* debug [instance_properties/getter]: allowsVibrancy */


// Returns the appearance object that’s active on the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/current
func (a_ Appearance) CurrentAppearance() IAppearance {
	rv := objc.Send[Appearance](a_.ID, objc.Sel("currentAppearance"))
	return rv
}/* debug [instance_properties/getter]: currentAppearance */


// Returns the appearance object that’s active on the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/current
func (a_ Appearance) SetCurrentAppearance(value IAppearance) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentAppearance:"), value)
}/* debug [instance_properties/setter]: currentAppearance */


// The appearance that the system uses for color and asset resolution, and that’s active for drawing, usually from locking focus on a view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/currentDrawingAppearance
func (a_ Appearance) CurrentDrawingAppearance() IAppearance {
	rv := objc.Send[Appearance](a_.ID, objc.Sel("currentDrawingAppearance"))
	return rv
}/* debug [instance_properties/getter]: currentDrawingAppearance */


// The name of the appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance/name-swift.property
func (a_ Appearance) Name() AppearanceName /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The appearance of the receiver, in an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearancecustomization/appearance
func (a_ Appearance) Appearance() IAppearance {
	rv := objc.Send[Appearance](a_.ID, objc.Sel("appearance"))
	return rv
}/* debug [instance_properties/getter]: appearance */


// The appearance of the receiver, in an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearancecustomization/appearance
func (a_ Appearance) SetAppearance(value IAppearance) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAppearance:"), value)
}/* debug [instance_properties/setter]: appearance */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAppearance */


