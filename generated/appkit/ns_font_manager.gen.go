// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AvailableFontFamilies() []string /* primitive/slice/pointer. */
	CollectionNames() objc.ID
	CurrentFontAction() FontAction /* not a class type */
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Enabled() bool /* primitive/slice/pointer. */
	SetEnabled(value bool /* primitive/slice/pointer. */)
	SelectedFont() IFont
	Target() objc.ID
	SetTarget(value objc.ID)
	AvailableFonts() string /* primitive/slice/pointer. */
	SetAvailableFonts(value string /* primitive/slice/pointer. */)
	IsEnabled() bool /* primitive/slice/pointer. */
	SetIsEnabled(value bool /* primitive/slice/pointer. */)
	IsMultiple() bool /* primitive/slice/pointer. */
	SetIsMultiple(value bool /* primitive/slice/pointer. */)
	// methods:
	AvailableMembersOfFontFamily(fam string /* primitive/slice/pointer. */) []foundation.objc.IObject /* cross-framework: Array */
	ConvertFont(fontObj IFont) IFont
	ConvertFontToSize(fontObj IFont, size float64 /* primitive/slice/pointer. */) IFont
	ConvertAttributes(attributes foundation.IDictionary /* already interface */) foundation.IDictionary /* already interface */
	ConvertFontTraits(traits FontTraitMask /* not a class type */) FontTraitMask /* not a class type */
	ConvertWeightOfFont(upFlag bool /* primitive/slice/pointer. */, fontObj IFont) IFont
	FontWithFamilyTraitsWeightSize(family string /* primitive/slice/pointer. */, traits FontTraitMask /* not a class type */, weight int /* primitive/slice/pointer. */, size float64 /* primitive/slice/pointer. */) IFont
	FontMenu(create bool /* primitive/slice/pointer. */) IMenu
	LocalizedNameForFamilyFace(family string /* primitive/slice/pointer. */, faceKey string /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */
	ModifyFontViaPanel(sender objectivec.IObject)
	OrderFrontStylesPanel(sender objectivec.IObject)
	RemoveFontTrait(sender objectivec.IObject)
	SetFontMenu(newMenu IMenu)
	SetSelectedAttributesIsMultiple(attributes foundation.IDictionary /* already interface */, flag bool /* primitive/slice/pointer. */)
	WeightOfFont(fontObj IFont) int /* primitive/slice/pointer. */
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


// Returns an array with one entry for each available member of a font family.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/availableMembers(ofFontFamily:)
func (f_ FontManager) AvailableMembersOfFontFamily(fam string /* primitive/slice/pointer. */) []foundation.objc.IObject /* cross-framework: Array */ {
	rv := objc.Send[[]foundation.Array](f_.ID, objc.Sel("availableMembersOfFontFamily:"), objc.String(fam))
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


// Returns a font object whose traits are the same as those of the given font, except for the size, which is changed to the given size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toSize:)
func (f_ FontManager) ConvertFontToSize(fontObj IFont, size float64 /* primitive/slice/pointer. */) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:toSize:"), fontObj, size)
	return rv
}


// Converts attributes in response to an object initiating an attribute change, typically the Font panel or Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convertAttributes(_:)
func (f_ FontManager) ConvertAttributes(attributes foundation.IDictionary /* already interface */) foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](f_.ID, objc.Sel("convertAttributes:"), attributes)
	return rv
}


// Converts font traits to a new traits mask value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convertFontTraits(_:)
func (f_ FontManager) ConvertFontTraits(traits FontTraitMask /* not a class type */) FontTraitMask /* not a class type */ {
	rv := objc.Send[FontTraitMask](f_.ID, objc.Sel("convertFontTraits:"), traits)
	return rv
}


// Returns a font object whose weight is greater or lesser than that of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convertWeight(_:of:)
func (f_ FontManager) ConvertWeightOfFont(upFlag bool /* primitive/slice/pointer. */, fontObj IFont) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertWeight:ofFont:"), upFlag, fontObj)
	return rv
}


// Attempts to load a font with the specified characteristics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/font(withFamily:traits:weight:size:)
func (f_ FontManager) FontWithFamilyTraitsWeightSize(family string /* primitive/slice/pointer. */, traits FontTraitMask /* not a class type */, weight int /* primitive/slice/pointer. */, size float64 /* primitive/slice/pointer. */) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("fontWithFamily:traits:weight:size:"), objc.String(family), traits, weight, size)
	return rv
}


// Returns the menu that’s connected to the font conversion system, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/fontMenu(_:)
func (f_ FontManager) FontMenu(create bool /* primitive/slice/pointer. */) IMenu {
	rv := objc.Send[Menu](f_.ID, objc.Sel("fontMenu:"), create)
	return rv
}


// Returns a localized string with the name of the specified font family and face, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/localizedName(forFamily:face:)
func (f_ FontManager) LocalizedNameForFamilyFace(family string /* primitive/slice/pointer. */, faceKey string /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](f_.ID, objc.Sel("localizedNameForFamily:face:"), objc.String(family), objc.String(faceKey))
	return rv
}


// Modifies a font trait using input from the Font panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/modifyFontViaPanel(_:)
func (f_ FontManager) ModifyFontViaPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("modifyFontViaPanel:"), sender)
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
func (f_ FontManager) SetSelectedAttributesIsMultiple(attributes foundation.IDictionary /* already interface */, flag bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSelectedAttributes:isMultiple:"), attributes, flag)
}


// Returns an approximation of the specified font’s weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/weight(of:)
func (f_ FontManager) WeightOfFont(fontObj IFont) int /* primitive/slice/pointer. */ {
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
func (f_ FontManager) AvailableFontFamilies() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](f_.ID, objc.Sel("availableFontFamilies"))
	return rv
}


// The names of the currently loaded font collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/collectionNames
func (f_ FontManager) CollectionNames() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("collectionNames"))
	return rv
}


// The current font conversion action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/currentFontAction
func (f_ FontManager) CurrentFontAction() FontAction /* not a class type */ {
	rv := objc.Send[FontAction](f_.ID, objc.Sel("currentFontAction"))
	return rv
}


// The font manager’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/delegate
func (f_ FontManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("delegate"))
	return rv
}


// The font manager’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/delegate
func (f_ FontManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/isEnabled
func (f_ FontManager) Enabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/isEnabled
func (f_ FontManager) SetEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEnabled:"), value)
}


// The currently selected font object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/selectedFont
func (f_ FontManager) SelectedFont() IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("selectedFont"))
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


// The names of the fonts available in the system (not the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/availablefonts
func (f_ FontManager) AvailableFonts() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](f_.ID, objc.Sel("availableFonts"))
	return rv
}


// The names of the fonts available in the system (not the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/availablefonts
func (f_ FontManager) SetAvailableFonts(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAvailableFonts:"), objc.String(value))
}


// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/isenabled
func (f_ FontManager) IsEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/isenabled
func (f_ FontManager) SetIsEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether the currently selected font has multiple fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/ismultiple
func (f_ FontManager) IsMultiple() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isMultiple"))
	return rv
}


// A Boolean value that indicates whether the currently selected font has multiple fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/ismultiple
func (f_ FontManager) SetIsMultiple(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsMultiple:"), value)
}



