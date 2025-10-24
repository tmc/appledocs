// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSFontPanel */


/* debug [class_header]: Header for NSFontPanel */
// The class instance for the [FontPanel] class.
var (
	FontPanelClass     _FontPanelClass
	FontPanelClassOnce sync.Once
)

func getFontPanelClass() _FontPanelClass {
	FontPanelClassOnce.Do(func() {
		FontPanelClass = _FontPanelClass{objc.GetClass("NSFontPanel")}
	})
	return FontPanelClass
}

type _FontPanelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FontPanel */
// An interface definition for the [FontPanel] class.
type IFontPanel interface {
	IPanel
	
/* debug [class_interface_properties]: Properties for FontPanel */
	// properties:
	AccessoryView() IView
	SetAccessoryView(value IView)
	Enabled() bool
	SetEnabled(value bool)
	WorksWhenModal() bool
	SetWorksWhenModal(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FontPanel */
	// methods:
	PanelConvertFont(fontObj IFont) IFont
	ReloadDefaultFontFamilies()
	SetPanelFontIsMultiple(fontObj IFont, flag bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FontPanel */
// Alloc allocates a new instance without initialization.
func (fc _FontPanelClass) Alloc() FontPanel {
	rv := objc.Send[FontPanel](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FontPanelClass) New() FontPanel {
	rv := objc.Send[FontPanel](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FontPanel) Init() FontPanel {
	rv := objc.Send[FontPanel](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FontPanel) Autorelease() FontPanel {
	rv := objc.Send[FontPanel](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFontPanel creates a new FontPanel instance.
func NewFontPanel() FontPanel {
	return getFontPanelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FontPanel */
// The Font panel—a user interface object that displays a list of available fonts, letting the user preview them and change the font used to display text.
//
// Actual changes to the font panel are made through conversion messages sent to the shared instance. There’s only one Font panel for each app.


// The Font panel—a user interface object that displays a list of available fonts, letting the user preview them and change the font used to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel
type FontPanel struct {
	Panel
}

// FontPanelFrom constructs a [FontPanel] from an unsafe.Pointer.
//
// The Font panel—a user interface object that displays a list of available fonts, letting the user preview them and change the font used to display text.
func FontPanelFrom(ptr unsafe.Pointer) FontPanel {
	return FontPanel{
		Panel: PanelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FontPanel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FontPanel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FontPanel */

// Returns the single instance for the application, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/shared
func (fc _FontPanelClass) SharedFontPanel() FontPanel {
	rv := objc.Send[FontPanel](objc.ID(fc.class), objc.Sel("sharedFontPanel"))
	return rv
}/* debug [class_properties_class/property]: sharedFontPanel */

// A Boolean value that indicates whether the shared Font panel has been created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/sharedFontPanelExists
func (fc _FontPanelClass) SharedFontPanelExists() bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("sharedFontPanelExists"))
	return rv
}/* debug [class_properties_class/property]: sharedFontPanelExists */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FontPanel */

// Converts the specified font using the settings in the receiver, with the aid of the shared if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/convert(_:)
func (f_ FontPanel) PanelConvertFont(fontObj IFont) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("panelConvertFont:"), fontObj)
	return rv
}/* debug [instance_methods/method]: PanelConvertFont */


// Triggers a reload to the default state, so that the delegate is called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/reloadDefaultFontFamilies()
func (f_ FontPanel) ReloadDefaultFontFamilies() {
	objc.Send[objc.ID](f_.ID, objc.Sel("reloadDefaultFontFamilies"))
}/* debug [instance_methods/method]: ReloadDefaultFontFamilies */


// Sets the selected font in the receiver to the specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/setPanelFont(_:isMultiple:)
func (f_ FontPanel) SetPanelFontIsMultiple(fontObj IFont, flag bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPanelFont:isMultiple:"), fontObj, flag)
}/* debug [instance_methods/method]: SetPanelFontIsMultiple */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FontPanel */

// The specified view as the receiver’s accessory view, allowing you to add custom controls to your application’s Font panel without having to create a subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/accessoryView
func (f_ FontPanel) AccessoryView() IView {
	rv := objc.Send[View](f_.ID, objc.Sel("accessoryView"))
	return rv
}/* debug [instance_properties/getter]: accessoryView */


// The specified view as the receiver’s accessory view, allowing you to add custom controls to your application’s Font panel without having to create a subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/accessoryView
func (f_ FontPanel) SetAccessoryView(value IView) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAccessoryView:"), value)
}/* debug [instance_properties/setter]: accessoryView */


// A Boolean that shows whether the receiver’s Set button is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/isEnabled
func (f_ FontPanel) Enabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean that shows whether the receiver’s Set button is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/isEnabled
func (f_ FontPanel) SetEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// Returns the single instance for the application, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/shared
func (f_ FontPanel) SharedFontPanel() IFontPanel {
	rv := objc.Send[FontPanel](f_.ID, objc.Sel("sharedFontPanel"))
	return rv
}/* debug [instance_properties/getter]: sharedFontPanel */


// A Boolean value that indicates whether the shared Font panel has been created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/sharedFontPanelExists
func (f_ FontPanel) SharedFontPanelExists() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("sharedFontPanelExists"))
	return rv
}/* debug [instance_properties/getter]: sharedFontPanelExists */


// A Boolean that indicates whether the receiver allows fonts to be changed in modal windows and panels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/worksWhenModal
func (f_ FontPanel) WorksWhenModal() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("worksWhenModal"))
	return rv
}/* debug [instance_properties/getter]: worksWhenModal */


// A Boolean that indicates whether the receiver allows fonts to be changed in modal windows and panels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/worksWhenModal
func (f_ FontPanel) SetWorksWhenModal(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWorksWhenModal:"), value)
}/* debug [instance_properties/setter]: worksWhenModal */


// A Boolean that shows whether the receiver’s Set button is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/isenabled
func (f_ FontPanel) IsEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean that shows whether the receiver’s Set button is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/isenabled
func (f_ FontPanel) SetIsEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFontPanel */



