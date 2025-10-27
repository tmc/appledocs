// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [FontPanel] class.
type IFontPanel interface {
	IPanel
	

	// properties:
	AccessoryView() IView
	SetAccessoryView(value IView)
	Enabled() bool
	SetEnabled(value bool)
	WorksWhenModal() bool
	SetWorksWhenModal(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)


	

	// methods:
	PanelConvertFont(fontObj IFont) IFont
	ReloadDefaultFontFamilies()
	SetPanelFontIsMultiple(fontObj IFont, flag bool)


}





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















// Returns the single instance for the application, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/shared
func (fc _FontPanelClass) SharedFontPanel() FontPanel {
	rv := objc.Send[FontPanel](objc.ID(fc.class), objc.Sel("sharedFontPanel"))
	return rv
}

// A Boolean value that indicates whether the shared Font panel has been created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/sharedFontPanelExists
func (fc _FontPanelClass) SharedFontPanelExists() bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("sharedFontPanelExists"))
	return rv
}






// Converts the specified font using the settings in the receiver, with the aid of the shared if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/convert(_:)
func (f_ FontPanel) PanelConvertFont(fontObj IFont) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("panelConvertFont:"), fontObj)
	return rv
}


// Triggers a reload to the default state, so that the delegate is called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/reloadDefaultFontFamilies()
func (f_ FontPanel) ReloadDefaultFontFamilies() {
	objc.Send[objc.ID](f_.ID, objc.Sel("reloadDefaultFontFamilies"))
}


// Sets the selected font in the receiver to the specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/setPanelFont(_:isMultiple:)
func (f_ FontPanel) SetPanelFontIsMultiple(fontObj IFont, flag bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPanelFont:isMultiple:"), fontObj, flag)
}







// The specified view as the receiver’s accessory view, allowing you to add custom controls to your application’s Font panel without having to create a subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/accessoryView
func (f_ FontPanel) AccessoryView() IView {
	rv := objc.Send[View](f_.ID, objc.Sel("accessoryView"))
	return rv
}


// The specified view as the receiver’s accessory view, allowing you to add custom controls to your application’s Font panel without having to create a subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/accessoryView
func (f_ FontPanel) SetAccessoryView(value IView) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAccessoryView:"), value)
}


// A Boolean that shows whether the receiver’s Set button is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/isEnabled
func (f_ FontPanel) Enabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean that shows whether the receiver’s Set button is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/isEnabled
func (f_ FontPanel) SetEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEnabled:"), value)
}


// Returns the single instance for the application, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/shared
func (f_ FontPanel) SharedFontPanel() IFontPanel {
	rv := objc.Send[FontPanel](f_.ID, objc.Sel("sharedFontPanel"))
	return rv
}


// A Boolean value that indicates whether the shared Font panel has been created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/sharedFontPanelExists
func (f_ FontPanel) SharedFontPanelExists() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("sharedFontPanelExists"))
	return rv
}


// A Boolean that indicates whether the receiver allows fonts to be changed in modal windows and panels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/worksWhenModal
func (f_ FontPanel) WorksWhenModal() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("worksWhenModal"))
	return rv
}


// A Boolean that indicates whether the receiver allows fonts to be changed in modal windows and panels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/worksWhenModal
func (f_ FontPanel) SetWorksWhenModal(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWorksWhenModal:"), value)
}


// A Boolean that shows whether the receiver’s Set button is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/isenabled
func (f_ FontPanel) IsEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean that shows whether the receiver’s Set button is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/isenabled
func (f_ FontPanel) SetIsEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsEnabled:"), value)
}








