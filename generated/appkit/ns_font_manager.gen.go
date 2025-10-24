// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFontManager */


/* debug [class_header]: Header for NSFontManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FontManager */
// An interface definition for the [FontManager] class.
type IFontManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FontManager */
	// properties:
	Action() objc.SEL
	SetAction(value objc.SEL)
	AvailableFontFamilies() []string
	AvailableFonts() []string
	CollectionNames() objc.IObject /* cross-framework: NSArray */
	CurrentFontAction() FontAction
	Delegate() objc.ID
	SetDelegate(value objc.ID)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FontManager */
	// methods:
	AddFontTrait(sender objc.IObject)
	AvailableFontNamesWithTraits(someTraits FontTraitMask) []string
	AvailableMembersOfFontFamily(fam objc.IObject /* cross-framework: NSString */) []foundation.Array
	ConvertFont(fontObj IFont) IFont
	ConvertFontToFace(fontObj IFont, typeface objc.IObject /* cross-framework: NSString */) IFont
	ConvertFontToFamily(fontObj IFont, family objc.IObject /* cross-framework: NSString */) IFont
	ConvertFontToHaveTrait(fontObj IFont, trait FontTraitMask) IFont
	ConvertFontToNotHaveTrait(fontObj IFont, trait FontTraitMask) IFont
	ConvertFontToSize(fontObj IFont, size float64) IFont
	ConvertAttributes(attributes foundation.IDictionary) foundation.IDictionary
	ConvertFontTraits(traits FontTraitMask) FontTraitMask
	ConvertWeightOfFont(upFlag bool, fontObj IFont) IFont
	FontWithFamilyTraitsWeightSize(family objc.IObject /* cross-framework: NSString */, traits FontTraitMask, weight int, size float64) IFont
	FontMenu(create bool) IMenu
	FontNamedHasTraits(fName objc.IObject /* cross-framework: NSString */, someTraits FontTraitMask) bool
	FontPanel(create bool) IFontPanel
	LocalizedNameForFamilyFace(family objc.IObject /* cross-framework: NSString */, faceKey objc.IObject /* cross-framework: NSString */) foundation.String
	ModifyFont(sender objc.IObject)
	ModifyFontViaPanel(sender objc.IObject)
	OrderFrontFontPanel(sender objc.IObject)
	OrderFrontStylesPanel(sender objc.IObject)
	RemoveFontTrait(sender objc.IObject)
	SendAction() bool
	SetFontMenu(newMenu IMenu)
	SetSelectedAttributesIsMultiple(attributes foundation.IDictionary, flag bool)
	SetSelectedFontIsMultiple(fontObj IFont, flag bool)
	TraitsOfFont(fontObj IFont) FontTraitMask
	WeightOfFont(fontObj IFont) int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FontManager */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FontManager */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FontManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FontManager */

// Sets the class that creates the shared font manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/setFontManagerFactory(_:)
func (fc _FontManagerClass) SetFontManagerFactory(factoryId objc.Class) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("setFontManagerFactory:"), factoryId)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetFontManagerFactory) */


// Sets the class that creates the shared Font panel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/setFontPanelFactory(_:)
func (fc _FontManagerClass) SetFontPanelFactory(factoryId objc.Class) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("setFontPanelFactory:"), factoryId)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetFontPanelFactory) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FontManager */

// Returns the shared instance of the font manager for the application, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/shared
func (fc _FontManagerClass) SharedFontManager() FontManager {
	rv := objc.Send[FontManager](objc.ID(fc.class), objc.Sel("sharedFontManager"))
	return rv
}/* debug [class_properties_class/property]: sharedFontManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FontManager */

// Adds a trait to the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/addFontTrait(_:)
func (f_ FontManager) AddFontTrait(sender objc.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("addFontTrait:"), sender)
}/* debug [instance_methods/method]: AddFontTrait */


// Returns the names of the fonts available in the system whose traits are described exactly by the given font trait mask (not the objects themselves).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/availableFontNames(with:)
func (f_ FontManager) AvailableFontNamesWithTraits(someTraits FontTraitMask) []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("availableFontNamesWithTraits:"), someTraits)
	return rv
}/* debug [instance_methods/method]: AvailableFontNamesWithTraits */


// Returns an array with one entry for each available member of a font family.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/availableMembers(ofFontFamily:)
func (f_ FontManager) AvailableMembersOfFontFamily(fam objc.IObject /* cross-framework: NSString */) []foundation.Array {
	rv := objc.Send[[]foundation.Array](f_.ID, objc.Sel("availableMembersOfFontFamily:"), fam)
	return rv
}/* debug [instance_methods/method]: AvailableMembersOfFontFamily */


// Converts the given font according to the object that initiated a font change, typically the Font panel or Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:)
func (f_ FontManager) ConvertFont(fontObj IFont) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:"), fontObj)
	return rv
}/* debug [instance_methods/method]: ConvertFont */


// Returns a font whose traits are as similar as possible to those of the given font except for the typeface, which is changed to the given typeface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toFace:)
func (f_ FontManager) ConvertFontToFace(fontObj IFont, typeface objc.IObject /* cross-framework: NSString */) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:toFace:"), fontObj, typeface)
	return rv
}/* debug [instance_methods/method]: ConvertFontToFace */


// Returns a font whose traits are as similar as possible to those of the given font except for the font family, which is changed to the given family.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toFamily:)
func (f_ FontManager) ConvertFontToFamily(fontObj IFont, family objc.IObject /* cross-framework: NSString */) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:toFamily:"), fontObj, family)
	return rv
}/* debug [instance_methods/method]: ConvertFontToFamily */


// Returns a new version of the font object containing a single additional trait.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toHaveTrait:)
func (f_ FontManager) ConvertFontToHaveTrait(fontObj IFont, trait FontTraitMask) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:toHaveTrait:"), fontObj, trait)
	return rv
}/* debug [instance_methods/method]: ConvertFontToHaveTrait */


// Returns a new version of a font object without the specified traits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toNotHaveTrait:)
func (f_ FontManager) ConvertFontToNotHaveTrait(fontObj IFont, trait FontTraitMask) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:toNotHaveTrait:"), fontObj, trait)
	return rv
}/* debug [instance_methods/method]: ConvertFontToNotHaveTrait */


// Returns a font object whose traits are the same as those of the given font, except for the size, which is changed to the given size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toSize:)
func (f_ FontManager) ConvertFontToSize(fontObj IFont, size float64) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertFont:toSize:"), fontObj, size)
	return rv
}/* debug [instance_methods/method]: ConvertFontToSize */


// Converts attributes in response to an object initiating an attribute change, typically the Font panel or Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convertAttributes(_:)
func (f_ FontManager) ConvertAttributes(attributes foundation.IDictionary) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](f_.ID, objc.Sel("convertAttributes:"), attributes)
	return rv
}/* debug [instance_methods/method]: ConvertAttributes */


// Converts font traits to a new traits mask value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convertFontTraits(_:)
func (f_ FontManager) ConvertFontTraits(traits FontTraitMask) FontTraitMask {
	rv := objc.Send[FontTraitMask](f_.ID, objc.Sel("convertFontTraits:"), traits)
	return rv
}/* debug [instance_methods/method]: ConvertFontTraits */


// Returns a font object whose weight is greater or lesser than that of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/convertWeight(_:of:)
func (f_ FontManager) ConvertWeightOfFont(upFlag bool, fontObj IFont) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("convertWeight:ofFont:"), upFlag, fontObj)
	return rv
}/* debug [instance_methods/method]: ConvertWeightOfFont */


// Attempts to load a font with the specified characteristics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/font(withFamily:traits:weight:size:)
func (f_ FontManager) FontWithFamilyTraitsWeightSize(family objc.IObject /* cross-framework: NSString */, traits FontTraitMask, weight int, size float64) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("fontWithFamily:traits:weight:size:"), family, traits, weight, size)
	return rv
}/* debug [instance_methods/method]: FontWithFamilyTraitsWeightSize */


// Returns the menu that’s connected to the font conversion system, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/fontMenu(_:)
func (f_ FontManager) FontMenu(create bool) IMenu {
	rv := objc.Send[Menu](f_.ID, objc.Sel("fontMenu:"), create)
	return rv
}/* debug [instance_methods/method]: FontMenu */


// Indicates whether the given font has all the specified traits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/fontNamed(_:hasTraits:)
func (f_ FontManager) FontNamedHasTraits(fName objc.IObject /* cross-framework: NSString */, someTraits FontTraitMask) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("fontNamed:hasTraits:"), fName, someTraits)
	return rv
}/* debug [instance_methods/method]: FontNamedHasTraits */


// Returns the application’s shared Font panel object, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/fontPanel(_:)
func (f_ FontManager) FontPanel(create bool) IFontPanel {
	rv := objc.Send[FontPanel](f_.ID, objc.Sel("fontPanel:"), create)
	return rv
}/* debug [instance_methods/method]: FontPanel */


// Returns a localized string with the name of the specified font family and face, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/localizedName(forFamily:face:)
func (f_ FontManager) LocalizedNameForFamilyFace(family objc.IObject /* cross-framework: NSString */, faceKey objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](f_.ID, objc.Sel("localizedNameForFamily:face:"), family, faceKey)
	return rv
}/* debug [instance_methods/method]: LocalizedNameForFamilyFace */


// Modifies a trait of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/modifyFont(_:)
func (f_ FontManager) ModifyFont(sender objc.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("modifyFont:"), sender)
}/* debug [instance_methods/method]: ModifyFont */


// Modifies a font trait using input from the Font panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/modifyFontViaPanel(_:)
func (f_ FontManager) ModifyFontViaPanel(sender objc.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("modifyFontViaPanel:"), sender)
}/* debug [instance_methods/method]: ModifyFontViaPanel */


// Opens the Font panel, creating it if necessary, and displays that panel in front of the app’s windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/orderFrontFontPanel(_:)
func (f_ FontManager) OrderFrontFontPanel(sender objc.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("orderFrontFontPanel:"), sender)
}/* debug [instance_methods/method]: OrderFrontFontPanel */


// Opens the Font Styles panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/orderFrontStylesPanel(_:)
func (f_ FontManager) OrderFrontStylesPanel(sender objc.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("orderFrontStylesPanel:"), sender)
}/* debug [instance_methods/method]: OrderFrontStylesPanel */


// Removes a trait from the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/removeFontTrait(_:)
func (f_ FontManager) RemoveFontTrait(sender objc.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("removeFontTrait:"), sender)
}/* debug [instance_methods/method]: RemoveFontTrait */


// A Boolean value that indicates whether a responder handled the font manager’s action message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/sendAction()
func (f_ FontManager) SendAction() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("sendAction"))
	return rv
}/* debug [instance_methods/method]: SendAction */


// Records the given menu as the application’s Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/setFontMenu(_:)
func (f_ FontManager) SetFontMenu(newMenu IMenu) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFontMenu:"), newMenu)
}/* debug [instance_methods/method]: SetFontMenu */


// Informs the Font panel that the specified font attributes changed for the selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/setSelectedAttributes(_:isMultiple:)
func (f_ FontManager) SetSelectedAttributesIsMultiple(attributes foundation.IDictionary, flag bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSelectedAttributes:isMultiple:"), attributes, flag)
}/* debug [instance_methods/method]: SetSelectedAttributesIsMultiple */


// Records the specified font as the currently selected font and updates the Font panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/setSelectedFont(_:isMultiple:)
func (f_ FontManager) SetSelectedFontIsMultiple(fontObj IFont, flag bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSelectedFont:isMultiple:"), fontObj, flag)
}/* debug [instance_methods/method]: SetSelectedFontIsMultiple */


// Returns the traits of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/traits(of:)
func (f_ FontManager) TraitsOfFont(fontObj IFont) FontTraitMask {
	rv := objc.Send[FontTraitMask](f_.ID, objc.Sel("traitsOfFont:"), fontObj)
	return rv
}/* debug [instance_methods/method]: TraitsOfFont */


// Returns an approximation of the specified font’s weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/weight(of:)
func (f_ FontManager) WeightOfFont(fontObj IFont) int {
	rv := objc.Send[int](f_.ID, objc.Sel("weightOfFont:"), fontObj)
	return rv
}/* debug [instance_methods/method]: WeightOfFont */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FontManager */

// The action sent to the first responder when the user selects a new font from the Font panel or chooses a command from the Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/action
func (f_ FontManager) Action() objc.SEL {
	rv := objc.Send[objc.SEL](f_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// The action sent to the first responder when the user selects a new font from the Font panel or chooses a command from the Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/action
func (f_ FontManager) SetAction(value objc.SEL) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// The names of the font families available in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/availableFontFamilies
func (f_ FontManager) AvailableFontFamilies() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("availableFontFamilies"))
	return rv
}/* debug [instance_properties/getter]: availableFontFamilies */


// The names of the fonts available in the system (not the objects themselves).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/availableFonts
func (f_ FontManager) AvailableFonts() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("availableFonts"))
	return rv
}/* debug [instance_properties/getter]: availableFonts */


// The names of the currently loaded font collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/collectionNames
func (f_ FontManager) CollectionNames() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](f_.ID, objc.Sel("collectionNames"))
	return rv
}/* debug [instance_properties/getter]: collectionNames */


// The current font conversion action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/currentFontAction
func (f_ FontManager) CurrentFontAction() FontAction {
	rv := objc.Send[FontAction](f_.ID, objc.Sel("currentFontAction"))
	return rv
}/* debug [instance_properties/getter]: currentFontAction */


// The font manager’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/delegate
func (f_ FontManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The font manager’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/delegate
func (f_ FontManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/isEnabled
func (f_ FontManager) Enabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/isEnabled
func (f_ FontManager) SetEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// A Boolean value that indicates whether the currently selected font has multiple fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/isMultiple
func (f_ FontManager) Multiple() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("multiple"))
	return rv
}/* debug [instance_properties/getter]: multiple */


// The currently selected font object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/selectedFont
func (f_ FontManager) SelectedFont() IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("selectedFont"))
	return rv
}/* debug [instance_properties/getter]: selectedFont */


// Returns the shared instance of the font manager for the application, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/shared
func (f_ FontManager) SharedFontManager() IFontManager {
	rv := objc.Send[FontManager](f_.ID, objc.Sel("sharedFontManager"))
	return rv
}/* debug [instance_properties/getter]: sharedFontManager */


// The object that receives action messages related to the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/target
func (f_ FontManager) Target() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// The object that receives action messages related to the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager/target
func (f_ FontManager) SetTarget(value objc.ID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */


// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/isenabled
func (f_ FontManager) IsEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/isenabled
func (f_ FontManager) SetIsEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */


// A Boolean value that indicates whether the currently selected font has multiple fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/ismultiple
func (f_ FontManager) IsMultiple() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isMultiple"))
	return rv
}/* debug [instance_properties/getter]: isMultiple */


// A Boolean value that indicates whether the currently selected font has multiple fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/ismultiple
func (f_ FontManager) SetIsMultiple(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsMultiple:"), value)
}/* debug [instance_properties/setter]: isMultiple */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFontManager */



