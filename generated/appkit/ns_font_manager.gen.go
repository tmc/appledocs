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
	

	// properties:
	Action() objc.SEL
	SetAction(value objc.SEL)
	AvailableFontFamilies() []string
	AvailableFonts() []string
	CollectionNames() foundation.foundation.INSArray
	CurrentFontAction() FontAction
	Enabled() bool
	SetEnabled(value bool)
	Multiple() bool
	SelectedFont() IFont
	Target() objc.ID
	SetTarget(value objc.ID)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsMultiple() bool
	SetIsMultiple(value bool)


	

	// methods:
	AddFontTrait(sender objectivec.IObject)
	AvailableFontNamesWithTraits(someTraits FontTraitMask) []string
	AvailableMembersOfFontFamily(fam foundation.foundation.INSString) []foundation.Array
	ConvertFont(fontObj IFont) IFont
	ConvertFontToFace(fontObj IFont, typeface foundation.foundation.INSString) IFont
	ConvertFontToFamily(fontObj IFont, family foundation.foundation.INSString) IFont
	ConvertFontToHaveTrait(fontObj IFont, trait FontTraitMask) IFont
	ConvertFontToNotHaveTrait(fontObj IFont, trait FontTraitMask) IFont
	ConvertFontToSize(fontObj IFont, size float64) IFont
	ConvertAttributes(attributes foundation.IDictionary) foundation.IDictionary
	ConvertFontTraits(traits FontTraitMask) FontTraitMask
	ConvertWeightOfFont(upFlag bool, fontObj IFont) IFont
	FontWithFamilyTraitsWeightSize(family foundation.foundation.INSString, traits FontTraitMask, weight int, size float64) IFont
	FontMenu(create bool) IMenu
	FontNamedHasTraits(fName foundation.foundation.INSString, someTraits FontTraitMask) bool
	FontPanel(create bool) IFontPanel
	LocalizedNameForFamilyFace(family foundation.foundation.INSString, faceKey foundation.foundation.INSString) foundation.String
	ModifyFont(sender objectivec.IObject)
	ModifyFontViaPanel(sender objectivec.IObject)
	OrderFrontFontPanel(sender objectivec.IObject)
	OrderFrontStylesPanel(sender objectivec.IObject)
	RemoveFontTrait(sender objectivec.IObject)
	SendAction() bool
	SetFontMenu(newMenu IMenu)
	SetSelectedAttributesIsMultiple(attributes foundation.IDictionary, flag bool)
	SetSelectedFontIsMultiple(fontObj IFont, flag bool)
	TraitsOfFont(fontObj IFont) FontTraitMask
	WeightOfFont(fontObj IFont) int


}





// Alloc allocates a new instance without initialization.
func (fc _FontManagerClass) Alloc() FontManager {
	rv := objc.Send[FontManager](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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










// Sets the class that creates the shared font manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/setFontManagerFactory(_:)
func (fc _FontManagerClass) SetFontManagerFactory(factoryId objc.Class) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("setFontManagerFactory:"), factoryId)
}


// Sets the class that creates the shared Font panel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/setFontPanelFactory(_:)
func (fc _FontManagerClass) SetFontPanelFactory(factoryId objc.Class) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("setFontPanelFactory:"), factoryId)
}







// Returns the shared instance of the font manager for the application, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/shared
func (fc _FontManagerClass) SharedFontManager() FontManager {
	rv := objc.Send[FontManager](objc.ID(fc.class), objc.Sel("sharedFontManager"))
	return rv
}






// Adds a trait to the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/addFontTrait(_:)
func (f_ FontManager) AddFontTrait(sender objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("addFontTrait:"), sender)
}


// Returns the names of the fonts available in the system whose traits are described exactly by the given font trait mask (not the objects themselves).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/availableFontNames(with:)
func (f_ FontManager) AvailableFontNamesWithTraits(someTraits FontTraitMask) []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("availableFontNamesWithTraits:"), someTraits)
	return rv
}


// Returns an array with one entry for each available member of a font family.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/availableMembers(ofFontFamily:)
func (f_ FontManager) AvailableMembersOfFontFamily(fam foundation.foundation.INSString) []foundation.Array {
	rv := objc.Send[[]foundation.Array](f_.ID, objc.Sel("availableMembersOfFontFamily:"), fam)
	return rv
}


// Converts the given font according to the object that initiated a font change, typically the Font panel or Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:)
func (f_ FontManager) ConvertFont(fontObj IFont) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:"), fontObj)
	return rv
}


// Returns a font whose traits are as similar as possible to those of the given font except for the typeface, which is changed to the given typeface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toFace:)
func (f_ FontManager) ConvertFontToFace(fontObj IFont, typeface foundation.foundation.INSString) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:toFace:"), fontObj, typeface)
	return rv
}


// Returns a font whose traits are as similar as possible to those of the given font except for the font family, which is changed to the given family.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toFamily:)
func (f_ FontManager) ConvertFontToFamily(fontObj IFont, family foundation.foundation.INSString) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:toFamily:"), fontObj, family)
	return rv
}


// Returns a new version of the font object containing a single additional trait.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toHaveTrait:)
func (f_ FontManager) ConvertFontToHaveTrait(fontObj IFont, trait FontTraitMask) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:toHaveTrait:"), fontObj, trait)
	return rv
}


// Returns a new version of a font object without the specified traits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toNotHaveTrait:)
func (f_ FontManager) ConvertFontToNotHaveTrait(fontObj IFont, trait FontTraitMask) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:toNotHaveTrait:"), fontObj, trait)
	return rv
}


// Returns a font object whose traits are the same as those of the given font, except for the size, which is changed to the given size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toSize:)
func (f_ FontManager) ConvertFontToSize(fontObj IFont, size float64) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:toSize:"), fontObj, size)
	return rv
}


// Converts attributes in response to an object initiating an attribute change, typically the Font panel or Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convertAttributes(_:)
func (f_ FontManager) ConvertAttributes(attributes foundation.IDictionary) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](f_.ID, objc.Sel("convertAttributes:"), attributes)
	return rv
}


// Converts font traits to a new traits mask value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convertFontTraits(_:)
func (f_ FontManager) ConvertFontTraits(traits FontTraitMask) FontTraitMask {
	rv := objc.Send[FontTraitMask](f_.ID, objc.Sel("convertFontTraits:"), traits)
	return rv
}


// Returns a font object whose weight is greater or lesser than that of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convertWeight(_:of:)
func (f_ FontManager) ConvertWeightOfFont(upFlag bool, fontObj IFont) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertWeight:ofFont:"), upFlag, fontObj)
	return rv
}


// Attempts to load a font with the specified characteristics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/font(withFamily:traits:weight:size:)
func (f_ FontManager) FontWithFamilyTraitsWeightSize(family foundation.foundation.INSString, traits FontTraitMask, weight int, size float64) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("fontWithFamily:traits:weight:size:"), family, traits, weight, size)
	return rv
}


// Returns the menu that’s connected to the font conversion system, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/fontMenu(_:)
func (f_ FontManager) FontMenu(create bool) IMenu {
	rv := objc.Send[Menu](f_.ID, objc.Sel("fontMenu:"), create)
	return rv
}


// Indicates whether the given font has all the specified traits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/fontNamed(_:hasTraits:)
func (f_ FontManager) FontNamedHasTraits(fName foundation.foundation.INSString, someTraits FontTraitMask) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("fontNamed:hasTraits:"), fName, someTraits)
	return rv
}


// Returns the application’s shared Font panel object, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/fontPanel(_:)
func (f_ FontManager) FontPanel(create bool) IFontPanel {
	rv := objc.Send[FontPanel](f_.ID, objc.Sel("fontPanel:"), create)
	return rv
}


// Returns a localized string with the name of the specified font family and face, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/localizedName(forFamily:face:)
func (f_ FontManager) LocalizedNameForFamilyFace(family foundation.foundation.INSString, faceKey foundation.foundation.INSString) foundation.String {
	rv := objc.Send[foundation.String](f_.ID, objc.Sel("localizedNameForFamily:face:"), family, faceKey)
	return rv
}


// Modifies a trait of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/modifyFont(_:)
func (f_ FontManager) ModifyFont(sender objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("modifyFont:"), sender)
}


// Modifies a font trait using input from the Font panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/modifyFontViaPanel(_:)
func (f_ FontManager) ModifyFontViaPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("modifyFontViaPanel:"), sender)
}


// Opens the Font panel, creating it if necessary, and displays that panel in front of the app’s windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/orderFrontFontPanel(_:)
func (f_ FontManager) OrderFrontFontPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("orderFrontFontPanel:"), sender)
}


// Opens the Font Styles panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/orderFrontStylesPanel(_:)
func (f_ FontManager) OrderFrontStylesPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("orderFrontStylesPanel:"), sender)
}


// Removes a trait from the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/removeFontTrait(_:)
func (f_ FontManager) RemoveFontTrait(sender objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("removeFontTrait:"), sender)
}


// A Boolean value that indicates whether a responder handled the font manager’s action message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/sendAction()
func (f_ FontManager) SendAction() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("sendAction"))
	return rv
}


// Records the given menu as the application’s Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/setFontMenu(_:)
func (f_ FontManager) SetFontMenu(newMenu IMenu) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFontMenu:"), newMenu)
}


// Informs the Font panel that the specified font attributes changed for the selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/setSelectedAttributes(_:isMultiple:)
func (f_ FontManager) SetSelectedAttributesIsMultiple(attributes foundation.IDictionary, flag bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSelectedAttributes:isMultiple:"), attributes, flag)
}


// Records the specified font as the currently selected font and updates the Font panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/setSelectedFont(_:isMultiple:)
func (f_ FontManager) SetSelectedFontIsMultiple(fontObj IFont, flag bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSelectedFont:isMultiple:"), fontObj, flag)
}


// Returns the traits of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/traits(of:)
func (f_ FontManager) TraitsOfFont(fontObj IFont) FontTraitMask {
	rv := objc.Send[FontTraitMask](f_.ID, objc.Sel("traitsOfFont:"), fontObj)
	return rv
}


// Returns an approximation of the specified font’s weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/weight(of:)
func (f_ FontManager) WeightOfFont(fontObj IFont) int {
	rv := objc.Send[int](f_.ID, objc.Sel("weightOfFont:"), fontObj)
	return rv
}







// The action sent to the first responder when the user selects a new font from the Font panel or chooses a command from the Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/action
func (f_ FontManager) Action() objc.SEL {
	rv := objc.Send[objc.SEL](f_.ID, objc.Sel("action"))
	return rv
}


// The action sent to the first responder when the user selects a new font from the Font panel or chooses a command from the Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/action
func (f_ FontManager) SetAction(value objc.SEL) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAction:"), value)
}


// The names of the font families available in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/availableFontFamilies
func (f_ FontManager) AvailableFontFamilies() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("availableFontFamilies"))
	return rv
}


// The names of the fonts available in the system (not the objects themselves).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/availableFonts
func (f_ FontManager) AvailableFonts() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("availableFonts"))
	return rv
}


// The names of the currently loaded font collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/collectionNames
func (f_ FontManager) CollectionNames() foundation.foundation.INSArray {
	rv := objc.Send[foundation.NSArray](f_.ID, objc.Sel("collectionNames"))
	return rv
}


// The current font conversion action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/currentFontAction
func (f_ FontManager) CurrentFontAction() FontAction {
	rv := objc.Send[FontAction](f_.ID, objc.Sel("currentFontAction"))
	return rv
}


// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/isEnabled
func (f_ FontManager) Enabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/isEnabled
func (f_ FontManager) SetEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEnabled:"), value)
}


// A Boolean value that indicates whether the currently selected font has multiple fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/isMultiple
func (f_ FontManager) Multiple() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("multiple"))
	return rv
}


// The currently selected font object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/selectedFont
func (f_ FontManager) SelectedFont() IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("selectedFont"))
	return rv
}


// Returns the shared instance of the font manager for the application, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/shared
func (f_ FontManager) SharedFontManager() IFontManager {
	rv := objc.Send[FontManager](f_.ID, objc.Sel("sharedFontManager"))
	return rv
}


// The object that receives action messages related to the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/target
func (f_ FontManager) Target() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("target"))
	return rv
}


// The object that receives action messages related to the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/target
func (f_ FontManager) SetTarget(value objc.ID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTarget:"), value)
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








