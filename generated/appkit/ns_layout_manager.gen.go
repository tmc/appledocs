// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LayoutManager] class.
var (
	LayoutManagerClass     _LayoutManagerClass
	LayoutManagerClassOnce sync.Once
)

func getLayoutManagerClass() _LayoutManagerClass {
	LayoutManagerClassOnce.Do(func() {
		LayoutManagerClass = _LayoutManagerClass{objc.GetClass("NSLayoutManager")}
	})
	return LayoutManagerClass
}

type _LayoutManagerClass struct {
	class objc.Class
}

// An interface definition for the [LayoutManager] class.
type ILayoutManager interface {
	objectivec.IObject
	AddTemporaryAttributeValueForCharacterRange(attrName unsafe.Pointer, value objc.ID, charRange foundation.Range)
	AddTemporaryAttributesForCharacterRange(attrs unsafe.Pointer, charRange foundation.Range)
	RemoveTemporaryAttributeForCharacterRange(attrName unsafe.Pointer, charRange foundation.Range)
}

// An object that coordinates the layout and display of text characters.
//
// maps Unicode character codes to glyphs, sets the glyphs in a series of objects, and displays them in a series of objects. In addition to its core function of laying out text, a layout manager object coordinates its text view objects, provides services to those text views to support instances for editing paragraph styles, and handles the layout and display of text attributes not inherent in glyphs (such as underline or strikethrough). You can create a subclass of to handle additional text attributes, whether inherent or not.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager
type LayoutManager struct {
	objectivec.Object
}

// LayoutManagerFrom constructs a [LayoutManager] from an unsafe.Pointer.
//
// An object that coordinates the layout and display of text characters.
func LayoutManagerFrom(ptr unsafe.Pointer) LayoutManager {
	return LayoutManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutManagerClass) Alloc() LayoutManager {
	rv := objc.Send[LayoutManager](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LayoutManagerClass) New() LayoutManager {
	rv := objc.Send[LayoutManager](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LayoutManager) Init() LayoutManager {
	rv := objc.Send[LayoutManager](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LayoutManager) Autorelease() LayoutManager {
	rv := objc.Send[LayoutManager](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayoutManager creates a new LayoutManager instance.
func NewLayoutManager() LayoutManager {
	return getLayoutManagerClass().New()
}


// Adds a temporary attribute to the characters in the specified range.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/addTemporaryAttribute(_:value:forCharacterRange:)
func (l_ LayoutManager) AddTemporaryAttributeValueForCharacterRange(attrName unsafe.Pointer, value objc.ID, charRange foundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addTemporaryAttribute:value:forCharacterRange:"), attrName, value, charRange)
}

// Appends one or more temporary attributes to the attributes dictionary of the specified character range.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/addTemporaryAttributes(_:forCharacterRange:)
func (l_ LayoutManager) AddTemporaryAttributesForCharacterRange(attrs unsafe.Pointer, charRange foundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addTemporaryAttributes:forCharacterRange:"), attrs, charRange)
}

// Removes a temporary attribute from the list of attributes for the specified character range.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/removeTemporaryAttribute(_:forCharacterRange:)
func (l_ LayoutManager) RemoveTemporaryAttributeForCharacterRange(attrName unsafe.Pointer, charRange foundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("removeTemporaryAttribute:forCharacterRange:"), attrName, charRange)
}

// The default typesetter behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/typesetterBehavior-swift.property
func (l_ LayoutManager) TypesetterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("typesetterBehavior"))
	return rv
}


// SetTypesetterBehavior sets the value of the typesetterBehavior property.
// The default typesetter behavior.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/typesetterBehavior-swift.property
func (l_ LayoutManager) SetTypesetterBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTypesetterBehavior:"), value)
}

// A Boolean value that indicates whether the layout manager allows noncontiguous layout.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/allowsnoncontiguouslayout
func (l_ LayoutManager) AllowsNonContiguousLayout() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("allowsNonContiguousLayout"))
	return rv
}


// SetAllowsNonContiguousLayout sets the value of the allowsNonContiguousLayout property.
// A Boolean value that indicates whether the layout manager allows noncontiguous layout.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/allowsnoncontiguouslayout
func (l_ LayoutManager) SetAllowsNonContiguousLayout(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAllowsNonContiguousLayout:"), value)
}

// A Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app’s run loop is idle.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/backgroundlayoutenabled
func (l_ LayoutManager) BackgroundLayoutEnabled() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("backgroundLayoutEnabled"))
	return rv
}


// SetBackgroundLayoutEnabled sets the value of the backgroundLayoutEnabled property.
// A Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app’s run loop is idle.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/backgroundlayoutenabled
func (l_ LayoutManager) SetBackgroundLayoutEnabled(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBackgroundLayoutEnabled:"), value)
}

// The default amount of scaling to apply when an attachment image is too large to fit in a text container.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/defaultattachmentscaling
func (l_ LayoutManager) DefaultAttachmentScaling() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("defaultAttachmentScaling"))
	return rv
}


// SetDefaultAttachmentScaling sets the value of the defaultAttachmentScaling property.
// The default amount of scaling to apply when an attachment image is too large to fit in a text container.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/defaultattachmentscaling
func (l_ LayoutManager) SetDefaultAttachmentScaling(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDefaultAttachmentScaling:"), value)
}

// The layout manager’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/delegate
func (l_ LayoutManager) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The layout manager’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/delegate
func (l_ LayoutManager) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelegate:"), value)
}

// The rectangle for the extra line fragment at the end of a document.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/extralinefragmentrect
func (l_ LayoutManager) ExtraLineFragmentRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](l_.ID, objc.Sel("extraLineFragmentRect"))
	return rv
}


// SetExtraLineFragmentRect sets the value of the extraLineFragmentRect property.
// The rectangle for the extra line fragment at the end of a document.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/extralinefragmentrect
func (l_ LayoutManager) SetExtraLineFragmentRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setExtraLineFragmentRect:"), value)
}

// The text container for the extra line fragment rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/extralinefragmenttextcontainer
func (l_ LayoutManager) ExtraLineFragmentTextContainer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("extraLineFragmentTextContainer"))
	return rv
}


// SetExtraLineFragmentTextContainer sets the value of the extraLineFragmentTextContainer property.
// The text container for the extra line fragment rectangle.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/extralinefragmenttextcontainer
func (l_ LayoutManager) SetExtraLineFragmentTextContainer(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setExtraLineFragmentTextContainer:"), value)
}

// The rectangle that encloses the insertion point in the extra line fragment rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/extralinefragmentusedrect
func (l_ LayoutManager) ExtraLineFragmentUsedRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](l_.ID, objc.Sel("extraLineFragmentUsedRect"))
	return rv
}


// SetExtraLineFragmentUsedRect sets the value of the extraLineFragmentUsedRect property.
// The rectangle that encloses the insertion point in the extra line fragment rectangle.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/extralinefragmentusedrect
func (l_ LayoutManager) SetExtraLineFragmentUsedRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setExtraLineFragmentUsedRect:"), value)
}

// The first text view in the layout manager’s series of text views.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/firsttextview
func (l_ LayoutManager) FirstTextView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("firstTextView"))
	return rv
}


// SetFirstTextView sets the value of the firstTextView property.
// The first text view in the layout manager’s series of text views.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/firsttextview
func (l_ LayoutManager) SetFirstTextView(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setFirstTextView:"), value)
}

// The glyph generator that the layout manager uses.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/glyphgenerator
func (l_ LayoutManager) GlyphGenerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("glyphGenerator"))
	return rv
}


// SetGlyphGenerator sets the value of the glyphGenerator property.
// The glyph generator that the layout manager uses.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/glyphgenerator
func (l_ LayoutManager) SetGlyphGenerator(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setGlyphGenerator:"), value)
}

// A Boolean value that indicates whether the layout manager currently has any areas of noncontiguous layout.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/hasnoncontiguouslayout
func (l_ LayoutManager) HasNonContiguousLayout() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("hasNonContiguousLayout"))
	return rv
}


// SetHasNonContiguousLayout sets the value of the hasNonContiguousLayout property.
// A Boolean value that indicates whether the layout manager currently has any areas of noncontiguous layout.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/hasnoncontiguouslayout
func (l_ LayoutManager) SetHasNonContiguousLayout(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setHasNonContiguousLayout:"), value)
}

// A Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/limitslayoutforsuspiciouscontents
func (l_ LayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("limitsLayoutForSuspiciousContents"))
	return rv
}


// SetLimitsLayoutForSuspiciousContents sets the value of the limitsLayoutForSuspiciousContents property.
// A Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/limitslayoutforsuspiciouscontents
func (l_ LayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLimitsLayoutForSuspiciousContents:"), value)
}

// The number of glyphs in the layout manager.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/numberofglyphs
func (l_ LayoutManager) NumberOfGlyphs() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfGlyphs"))
	return rv
}


// SetNumberOfGlyphs sets the value of the numberOfGlyphs property.
// The number of glyphs in the layout manager.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/numberofglyphs
func (l_ LayoutManager) SetNumberOfGlyphs(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfGlyphs:"), value)
}

// A Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/showscontrolcharacters
func (l_ LayoutManager) ShowsControlCharacters() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("showsControlCharacters"))
	return rv
}


// SetShowsControlCharacters sets the value of the showsControlCharacters property.
// A Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/showscontrolcharacters
func (l_ LayoutManager) SetShowsControlCharacters(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShowsControlCharacters:"), value)
}

// A Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/showsinvisiblecharacters
func (l_ LayoutManager) ShowsInvisibleCharacters() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("showsInvisibleCharacters"))
	return rv
}


// SetShowsInvisibleCharacters sets the value of the showsInvisibleCharacters property.
// A Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/showsinvisiblecharacters
func (l_ LayoutManager) SetShowsInvisibleCharacters(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShowsInvisibleCharacters:"), value)
}

// The current text containers of the layout manager.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/textcontainers
func (l_ LayoutManager) TextContainers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("textContainers"))
	return rv
}


// SetTextContainers sets the value of the textContainers property.
// The current text containers of the layout manager.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/textcontainers
func (l_ LayoutManager) SetTextContainers(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTextContainers:"), value)
}

// The text storage object that contains the content to lay out.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/textstorage
func (l_ LayoutManager) TextStorage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("textStorage"))
	return rv
}


// SetTextStorage sets the value of the textStorage property.
// The text storage object that contains the content to lay out.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/textstorage
func (l_ LayoutManager) SetTextStorage(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTextStorage:"), value)
}

// The text view that contains the first glyph in the selection.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/textviewforbeginningofselection
func (l_ LayoutManager) TextViewForBeginningOfSelection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("textViewForBeginningOfSelection"))
	return rv
}


// SetTextViewForBeginningOfSelection sets the value of the textViewForBeginningOfSelection property.
// The text view that contains the first glyph in the selection.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/textviewforbeginningofselection
func (l_ LayoutManager) SetTextViewForBeginningOfSelection(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTextViewForBeginningOfSelection:"), value)
}

// The current typesetter.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/typesetter
func (l_ LayoutManager) Typesetter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("typesetter"))
	return rv
}


// SetTypesetter sets the value of the typesetter property.
// The current typesetter.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/typesetter
func (l_ LayoutManager) SetTypesetter(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTypesetter:"), value)
}

// A Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/usesdefaulthyphenation
func (l_ LayoutManager) UsesDefaultHyphenation() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("usesDefaultHyphenation"))
	return rv
}


// SetUsesDefaultHyphenation sets the value of the usesDefaultHyphenation property.
// A Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/usesdefaulthyphenation
func (l_ LayoutManager) SetUsesDefaultHyphenation(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setUsesDefaultHyphenation:"), value)
}

// A Boolean value that indicates whether the layout manager uses the leading of the font.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/usesfontleading
func (l_ LayoutManager) UsesFontLeading() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("usesFontLeading"))
	return rv
}


// SetUsesFontLeading sets the value of the usesFontLeading property.
// A Boolean value that indicates whether the layout manager uses the leading of the font.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/usesfontleading
func (l_ LayoutManager) SetUsesFontLeading(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setUsesFontLeading:"), value)
}



