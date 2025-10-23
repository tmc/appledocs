// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FontManager] class.
var (
	FontManagerClass     _FontManagerClass
	FontManagerClassOnce sync.Once
)

func getFontManagerClass() _FontManagerClass {
	FontManagerClassOnce.Do(func() {
		FontManagerClass = _FontManagerClass{objc.GetClass("NSFontManager")}
	})
	return FontManagerClass
}

type _FontManagerClass struct {
	class objc.Class
}

// An interface definition for the [FontManager] class.
type IFontManager interface {
	objectivec.IObject
	SetFontMenu(newMenu IMenu)
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	AvailableFontFamilies() string
	SetAvailableFontFamilies(value string)
	AvailableFonts() string
	SetAvailableFonts(value string)
	CurrentFontAction() unsafe.Pointer
	SetCurrentFontAction(value unsafe.Pointer)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsMultiple() bool
	SetIsMultiple(value bool)
	SelectedFont() NSFont
	SetSelectedFont(value IFont)
	Target() unsafe.Pointer
	SetTarget(value unsafe.Pointer)
}

// The center of activity for the font-conversion system.
//
// The font manager records the currently selected font, updates the Font panel and Font menu to reflect the selected font, initiates font changes, and converts fonts in response to requests from text-bearing objects. In a more prosaic role, can be queried for the fonts available to the application and for the particular attributes of a font, such as whether it’s condensed or extended. You typically set up a font manager and the Font menu using Interface Builder. However, you can also do so programmatically by getting the shared font manager instance and having it create the standard Font menu at runtime: You can then add the Font menu to your app’s main menu. After the Font menu is installed, your app automatically gains the functionality of both the Font menu and the Font panel. Font collections are managed by .


// The center of activity for the font-conversion system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager
type FontManager struct {
	objectivec.Object
}

// FontManagerFrom constructs a [FontManager] from an unsafe.Pointer.
//
// The center of activity for the font-conversion system.
func FontManagerFrom(ptr unsafe.Pointer) FontManager {
	return FontManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FontManagerClass) Alloc() FontManager {
	rv := objc.Send[FontManager](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FontManagerClass) New() FontManager {
	rv := objc.Send[FontManager](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FontManager) Init() FontManager {
	rv := objc.Send[FontManager](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FontManager) Autorelease() FontManager {
	rv := objc.Send[FontManager](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFontManager creates a new FontManager instance.
func NewFontManager() FontManager {
	return getFontManagerClass().New()
}



// Records the given menu as the application’s Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/setFontMenu(_:)
func (f_ FontManager) SetFontMenu(newMenu IMenu) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFontMenu:"), newMenu)
}


// The action sent to the first responder when the user selects a new font from the Font panel or chooses a command from the Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/action
func (f_ FontManager) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("action"))
	return rv
}


// The action sent to the first responder when the user selects a new font from the Font panel or chooses a command from the Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/action
func (f_ FontManager) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAction:"), value)
}


// The names of the font families available in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/availablefontfamilies
func (f_ FontManager) AvailableFontFamilies() string {
	rv := objc.Send[string](f_.ID, objc.Sel("availableFontFamilies"))
	return rv
}


// The names of the font families available in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/availablefontfamilies
func (f_ FontManager) SetAvailableFontFamilies(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAvailableFontFamilies:"), objc.String(value))
}


// The names of the fonts available in the system (not the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/availablefonts
func (f_ FontManager) AvailableFonts() string {
	rv := objc.Send[string](f_.ID, objc.Sel("availableFonts"))
	return rv
}


// The names of the fonts available in the system (not the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/availablefonts
func (f_ FontManager) SetAvailableFonts(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAvailableFonts:"), objc.String(value))
}


// The current font conversion action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/currentfontaction
func (f_ FontManager) CurrentFontAction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("currentFontAction"))
	return rv
}


// The current font conversion action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/currentfontaction
func (f_ FontManager) SetCurrentFontAction(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setCurrentFontAction:"), value)
}


// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/isenabled
func (f_ FontManager) IsEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/isenabled
func (f_ FontManager) SetIsEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether the currently selected font has multiple fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/ismultiple
func (f_ FontManager) IsMultiple() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isMultiple"))
	return rv
}


// A Boolean value that indicates whether the currently selected font has multiple fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/ismultiple
func (f_ FontManager) SetIsMultiple(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsMultiple:"), value)
}


// The currently selected font object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/selectedfont
func (f_ FontManager) SelectedFont() NSFont {
	rv := objc.Send[NSFont](f_.ID, objc.Sel("selectedFont"))
	return rv
}


// The currently selected font object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/selectedfont
func (f_ FontManager) SetSelectedFont(value IFont) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSelectedFont:"), value)
}


// The object that receives action messages related to the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/target
func (f_ FontManager) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("target"))
	return rv
}


// The object that receives action messages related to the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/target
func (f_ FontManager) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTarget:"), value)
}



