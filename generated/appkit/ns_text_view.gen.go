// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [TextView] class.
var (
	TextViewClass     _TextViewClass
	TextViewClassOnce sync.Once
)

func getTextViewClass() _TextViewClass {
	TextViewClassOnce.Do(func() {
		TextViewClass = _TextViewClass{objc.GetClass("NSTextView")}
	})
	return TextViewClass
}

type _TextViewClass struct {
	class objc.Class
}

// An interface definition for the [TextView] class.
type ITextView interface {
	IText
}

// A view that draws text and handles user interactions with that text.
//
// The class is the front-end class to the AppKit text system. The class draws the text managed by the back-end components and handles user events to select and modify its text, in addition to supporting rich text, attachments, input management, and key binding, and marked text attributes. is the principal means to obtain a text object that caters to almost all needs for displaying and managing text at the user interface level. While is a subclass of the class — which declares the most general Cocoa interface to the text system — adds major features beyond the capabilities of . You can also do more powerful and more creative text manipulation (such as displaying text in a circle) using , , , and related classes. You’re more likely to use the class than . It’s also important to remember that conforms to a large number of protocols, the methods of which are available to instances of the class. communicates with its delegate through methods declared both by the and by its superclass’s protocol, . All delegation messages come from the first text view. In macOS 12 and later, if you explicitly call the property on a text view or text container, the framework reverts to a compatibility mode that uses . The text view also switches to this compatibility mode when it encounters text content that’s not yet supported, such as .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView
type TextView struct {
	Text
}

// TextViewFrom constructs a [TextView] from an unsafe.Pointer.
//
// A view that draws text and handles user interactions with that text.
func TextViewFrom(ptr unsafe.Pointer) TextView {
	return TextView{
		Text: TextFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextViewClass) Alloc() TextView {
	rv := objc.Send[TextView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextViewClass) New() TextView {
	rv := objc.Send[TextView](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextView) Init() TextView {
	rv := objc.Send[TextView](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextView) Autorelease() TextView {
	rv := objc.Send[TextView](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextView creates a new TextView instance.
func NewTextView() TextView {
	return getTextViewClass().New()
}


// The layout manager that lays out text for the receiver’s text container.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/layoutManager
func (t_ TextView) LayoutManager() NSLayoutManager {
	rv := objc.Send[NSLayoutManager](t_.ID, objc.Sel("layoutManager"))
	return rv
}

// The receiver’s text container.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainer
func (t_ TextView) TextContainer() NSTextContainer {
	rv := objc.Send[NSTextContainer](t_.ID, objc.Sel("textContainer"))
	return rv
}


// SetTextContainer sets the value of the textContainer property.
// The receiver’s text container.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainer
func (t_ TextView) SetTextContainer(value ITextContainer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainer:"), value)
}

// The receiver’s text storage object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textStorage
func (t_ TextView) TextStorage() NSTextStorage {
	rv := objc.Send[NSTextStorage](t_.ID, objc.Sel("textStorage"))
	return rv
}

// The data types that the receiver accepts as the destination view of a dragging operation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/acceptabledragtypes
func (t_ TextView) AcceptableDragTypes() PasteboardType {
	rv := objc.Send[PasteboardType](t_.ID, objc.Sel("acceptableDragTypes"))
	return rv
}


// SetAcceptableDragTypes sets the value of the acceptableDragTypes property.
// The data types that the receiver accepts as the destination view of a dragging operation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/acceptabledragtypes
func (t_ TextView) SetAcceptableDragTypes(value PasteboardType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptableDragTypes:"), value)
}

// A Boolean value that indicates whether the receiver accepts the glyph info attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/acceptsglyphinfo
func (t_ TextView) AcceptsGlyphInfo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("acceptsGlyphInfo"))
	return rv
}


// SetAcceptsGlyphInfo sets the value of the acceptsGlyphInfo property.
// A Boolean value that indicates whether the receiver accepts the glyph info attribute.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/acceptsglyphinfo
func (t_ TextView) SetAcceptsGlyphInfo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptsGlyphInfo:"), value)
}

// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowedinputsourcelocales
func (t_ TextView) AllowedInputSourceLocales() string {
	rv := objc.Send[string](t_.ID, objc.Sel("allowedInputSourceLocales"))
	return rv
}


// SetAllowedInputSourceLocales sets the value of the allowedInputSourceLocales property.
// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowedinputsourcelocales
func (t_ TextView) SetAllowedInputSourceLocales(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedInputSourceLocales:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowedwritingtoolsresultoptions
func (t_ TextView) AllowedWritingToolsResultOptions() WritingToolsResultOptions {
	rv := objc.Send[WritingToolsResultOptions](t_.ID, objc.Sel("allowedWritingToolsResultOptions"))
	return rv
}


// SetAllowedWritingToolsResultOptions sets the value of the allowedWritingToolsResultOptions property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowedwritingtoolsresultoptions
func (t_ TextView) SetAllowedWritingToolsResultOptions(value WritingToolsResultOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedWritingToolsResultOptions:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowscharacterpickertouchbaritem
func (t_ TextView) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsCharacterPickerTouchBarItem"))
	return rv
}


// SetAllowsCharacterPickerTouchBarItem sets the value of the allowsCharacterPickerTouchBarItem property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowscharacterpickertouchbaritem
func (t_ TextView) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsCharacterPickerTouchBarItem:"), value)
}

// A Boolean value that indicates whether the receiver allows its background color to change.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsdocumentbackgroundcolorchange
func (t_ TextView) AllowsDocumentBackgroundColorChange() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDocumentBackgroundColorChange"))
	return rv
}


// SetAllowsDocumentBackgroundColorChange sets the value of the allowsDocumentBackgroundColorChange property.
// A Boolean value that indicates whether the receiver allows its background color to change.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsdocumentbackgroundcolorchange
func (t_ TextView) SetAllowsDocumentBackgroundColorChange(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsDocumentBackgroundColorChange:"), value)
}

// Indicates whether image attachments should permit editing of their images.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsimageediting
func (t_ TextView) AllowsImageEditing() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsImageEditing"))
	return rv
}


// SetAllowsImageEditing sets the value of the allowsImageEditing property.
// Indicates whether image attachments should permit editing of their images.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsimageediting
func (t_ TextView) SetAllowsImageEditing(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsImageEditing:"), value)
}

// A Boolean value that indicates whether the receiver allows undo.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsundo
func (t_ TextView) AllowsUndo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsUndo"))
	return rv
}


// SetAllowsUndo sets the value of the allowsUndo property.
// A Boolean value that indicates whether the receiver allows undo.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsundo
func (t_ TextView) SetAllowsUndo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsUndo:"), value)
}

// The receiver’s background color.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/backgroundcolor
func (t_ TextView) BackgroundColor() NSColor {
	rv := objc.Send[NSColor](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The receiver’s background color.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/backgroundcolor
func (t_ TextView) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/candidatelisttouchbaritem
func (t_ TextView) CandidateListTouchBarItem() NSCandidateListTouchBarItem {
	rv := objc.Send[NSCandidateListTouchBarItem](t_.ID, objc.Sel("candidateListTouchBarItem"))
	return rv
}


// SetCandidateListTouchBarItem sets the value of the candidateListTouchBarItem property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/candidatelisttouchbaritem
func (t_ TextView) SetCandidateListTouchBarItem(value ICandidateListTouchBarItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCandidateListTouchBarItem:"), value)
}

// The receiver’s default paragraph style.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/defaultparagraphstyle
func (t_ TextView) DefaultParagraphStyle() NSParagraphStyle {
	rv := objc.Send[NSParagraphStyle](t_.ID, objc.Sel("defaultParagraphStyle"))
	return rv
}


// SetDefaultParagraphStyle sets the value of the defaultParagraphStyle property.
// The receiver’s default paragraph style.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/defaultparagraphstyle
func (t_ TextView) SetDefaultParagraphStyle(value NSParagraphStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDefaultParagraphStyle:"), value)
}

// The delegate for all text views sharing the receiver’s layout manager.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/delegate
func (t_ TextView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate for all text views sharing the receiver’s layout manager.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/delegate
func (t_ TextView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/displayslinktooltips
func (t_ TextView) DisplaysLinkToolTips() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("displaysLinkToolTips"))
	return rv
}


// SetDisplaysLinkToolTips sets the value of the displaysLinkToolTips property.
// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/displayslinktooltips
func (t_ TextView) SetDisplaysLinkToolTips(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDisplaysLinkToolTips:"), value)
}

// A Boolean value that indicates whether the receiver draws its background.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/drawsbackground
func (t_ TextView) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}


// SetDrawsBackground sets the value of the drawsBackground property.
// A Boolean value that indicates whether the receiver draws its background.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/drawsbackground
func (t_ TextView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}

// The default text checking types.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/enabledtextcheckingtypes
func (t_ TextView) EnabledTextCheckingTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("enabledTextCheckingTypes"))
	return rv
}


// SetEnabledTextCheckingTypes sets the value of the enabledTextCheckingTypes property.
// The default text checking types.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/enabledtextcheckingtypes
func (t_ TextView) SetEnabledTextCheckingTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEnabledTextCheckingTypes:"), value)
}

// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/importsgraphics
func (t_ TextView) ImportsGraphics() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("importsGraphics"))
	return rv
}


// SetImportsGraphics sets the value of the importsGraphics property.
// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/importsgraphics
func (t_ TextView) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImportsGraphics:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/inlinepredictiontype
func (t_ TextView) InlinePredictionType() TextInputTraitType {
	rv := objc.Send[TextInputTraitType](t_.ID, objc.Sel("inlinePredictionType"))
	return rv
}


// SetInlinePredictionType sets the value of the inlinePredictionType property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/inlinepredictiontype
func (t_ TextView) SetInlinePredictionType(value TextInputTraitType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInlinePredictionType:"), value)
}

// The color of the insertion point.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/insertionpointcolor
func (t_ TextView) InsertionPointColor() NSColor {
	rv := objc.Send[NSColor](t_.ID, objc.Sel("insertionPointColor"))
	return rv
}


// SetInsertionPointColor sets the value of the insertionPointColor property.
// The color of the insertion point.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/insertionpointcolor
func (t_ TextView) SetInsertionPointColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInsertionPointColor:"), value)
}

// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdashsubstitutionenabled
func (t_ TextView) IsAutomaticDashSubstitutionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticDashSubstitutionEnabled"))
	return rv
}


// SetIsAutomaticDashSubstitutionEnabled sets the value of the isAutomaticDashSubstitutionEnabled property.
// A Boolean value that indicates whether automatic dash substitution is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdashsubstitutionenabled
func (t_ TextView) SetIsAutomaticDashSubstitutionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticDashSubstitutionEnabled:"), value)
}

// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdatadetectionenabled
func (t_ TextView) IsAutomaticDataDetectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticDataDetectionEnabled"))
	return rv
}


// SetIsAutomaticDataDetectionEnabled sets the value of the isAutomaticDataDetectionEnabled property.
// A Boolean value that indicates whether automatic data detection is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdatadetectionenabled
func (t_ TextView) SetIsAutomaticDataDetectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticDataDetectionEnabled:"), value)
}

// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticlinkdetectionenabled
func (t_ TextView) IsAutomaticLinkDetectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticLinkDetectionEnabled"))
	return rv
}


// SetIsAutomaticLinkDetectionEnabled sets the value of the isAutomaticLinkDetectionEnabled property.
// A Boolean value that enables or disables automatic link detection.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticlinkdetectionenabled
func (t_ TextView) SetIsAutomaticLinkDetectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticLinkDetectionEnabled:"), value)
}

// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticquotesubstitutionenabled
func (t_ TextView) IsAutomaticQuoteSubstitutionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticQuoteSubstitutionEnabled"))
	return rv
}


// SetIsAutomaticQuoteSubstitutionEnabled sets the value of the isAutomaticQuoteSubstitutionEnabled property.
// A Boolean value that enables and disables automatic quotation mark substitution.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticquotesubstitutionenabled
func (t_ TextView) SetIsAutomaticQuoteSubstitutionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticQuoteSubstitutionEnabled:"), value)
}

// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticspellingcorrectionenabled
func (t_ TextView) IsAutomaticSpellingCorrectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticSpellingCorrectionEnabled"))
	return rv
}


// SetIsAutomaticSpellingCorrectionEnabled sets the value of the isAutomaticSpellingCorrectionEnabled property.
// A Boolean value that indicates whether automatic spelling correction is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticspellingcorrectionenabled
func (t_ TextView) SetIsAutomaticSpellingCorrectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticSpellingCorrectionEnabled:"), value)
}

// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextcompletionenabled
func (t_ TextView) IsAutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticTextCompletionEnabled"))
	return rv
}


// SetIsAutomaticTextCompletionEnabled sets the value of the isAutomaticTextCompletionEnabled property.
// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextcompletionenabled
func (t_ TextView) SetIsAutomaticTextCompletionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticTextCompletionEnabled:"), value)
}

// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextreplacementenabled
func (t_ TextView) IsAutomaticTextReplacementEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticTextReplacementEnabled"))
	return rv
}


// SetIsAutomaticTextReplacementEnabled sets the value of the isAutomaticTextReplacementEnabled property.
// A Boolean value that indicates whether automatic text replacement is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextreplacementenabled
func (t_ TextView) SetIsAutomaticTextReplacementEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticTextReplacementEnabled:"), value)
}

// A Boolean value that indicates whether undo coalescing is in progress.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscoalescingundo
func (t_ TextView) IsCoalescingUndo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isCoalescingUndo"))
	return rv
}


// SetIsCoalescingUndo sets the value of the isCoalescingUndo property.
// A Boolean value that indicates whether undo coalescing is in progress.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscoalescingundo
func (t_ TextView) SetIsCoalescingUndo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsCoalescingUndo:"), value)
}

// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscontinuousspellcheckingenabled
func (t_ TextView) IsContinuousSpellCheckingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isContinuousSpellCheckingEnabled"))
	return rv
}


// SetIsContinuousSpellCheckingEnabled sets the value of the isContinuousSpellCheckingEnabled property.
// A Boolean value that indicates whether the receiver has continuous spell checking enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscontinuousspellcheckingenabled
func (t_ TextView) SetIsContinuousSpellCheckingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsContinuousSpellCheckingEnabled:"), value)
}

// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iseditable
func (t_ TextView) IsEditable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEditable"))
	return rv
}


// SetIsEditable sets the value of the isEditable property.
// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iseditable
func (t_ TextView) SetIsEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEditable:"), value)
}

// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isfieldeditor
func (t_ TextView) IsFieldEditor() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFieldEditor"))
	return rv
}


// SetIsFieldEditor sets the value of the isFieldEditor property.
// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isfieldeditor
func (t_ TextView) SetIsFieldEditor(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFieldEditor:"), value)
}

// Enables and disables grammar checking.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isgrammarcheckingenabled
func (t_ TextView) IsGrammarCheckingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isGrammarCheckingEnabled"))
	return rv
}


// SetIsGrammarCheckingEnabled sets the value of the isGrammarCheckingEnabled property.
// Enables and disables grammar checking.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isgrammarcheckingenabled
func (t_ TextView) SetIsGrammarCheckingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsGrammarCheckingEnabled:"), value)
}

// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isincrementalsearchingenabled
func (t_ TextView) IsIncrementalSearchingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isIncrementalSearchingEnabled"))
	return rv
}


// SetIsIncrementalSearchingEnabled sets the value of the isIncrementalSearchingEnabled property.
// A Boolean value that indicates whether incremental searching is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isincrementalsearchingenabled
func (t_ TextView) SetIsIncrementalSearchingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsIncrementalSearchingEnabled:"), value)
}

// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrichtext
func (t_ TextView) IsRichText() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRichText"))
	return rv
}


// SetIsRichText sets the value of the isRichText property.
// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrichtext
func (t_ TextView) SetIsRichText(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRichText:"), value)
}

// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrulervisible
func (t_ TextView) IsRulerVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRulerVisible"))
	return rv
}


// SetIsRulerVisible sets the value of the isRulerVisible property.
// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrulervisible
func (t_ TextView) SetIsRulerVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRulerVisible:"), value)
}

// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isselectable
func (t_ TextView) IsSelectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelectable"))
	return rv
}


// SetIsSelectable sets the value of the isSelectable property.
// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isselectable
func (t_ TextView) SetIsSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSelectable:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iswritingtoolsactive
func (t_ TextView) IsWritingToolsActive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isWritingToolsActive"))
	return rv
}


// SetIsWritingToolsActive sets the value of the isWritingToolsActive property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iswritingtoolsactive
func (t_ TextView) SetIsWritingToolsActive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsWritingToolsActive:"), value)
}

// The attributes used to draw the onscreen presentation of link text.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/linktextattributes
func (t_ TextView) LinkTextAttributes() coreml.Key {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("linkTextAttributes"))
	return rv
}


// SetLinkTextAttributes sets the value of the linkTextAttributes property.
// The attributes used to draw the onscreen presentation of link text.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/linktextattributes
func (t_ TextView) SetLinkTextAttributes(value coreml.IKey) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLinkTextAttributes:"), value)
}

// The attributes used to draw marked text.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/markedtextattributes
func (t_ TextView) MarkedTextAttributes() coreml.Key {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("markedTextAttributes"))
	return rv
}


// SetMarkedTextAttributes sets the value of the markedTextAttributes property.
// The attributes used to draw marked text.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/markedtextattributes
func (t_ TextView) SetMarkedTextAttributes(value coreml.IKey) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMarkedTextAttributes:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/mathexpressioncompletiontype
func (t_ TextView) MathExpressionCompletionType() TextInputTraitType {
	rv := objc.Send[TextInputTraitType](t_.ID, objc.Sel("mathExpressionCompletionType"))
	return rv
}


// SetMathExpressionCompletionType sets the value of the mathExpressionCompletionType property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/mathexpressioncompletiontype
func (t_ TextView) SetMathExpressionCompletionType(value TextInputTraitType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMathExpressionCompletionType:"), value)
}

// The range of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusercharacterattributechange
func (t_ TextView) RangeForUserCharacterAttributeChange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("rangeForUserCharacterAttributeChange"))
	return rv
}


// SetRangeForUserCharacterAttributeChange sets the value of the rangeForUserCharacterAttributeChange property.
// The range of characters affected by an action method that changes character (not paragraph) attributes.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusercharacterattributechange
func (t_ TextView) SetRangeForUserCharacterAttributeChange(value foundation.IRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangeForUserCharacterAttributeChange:"), value)
}

// The partial range from the most recent beginning of a word up to the insertion point.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusercompletion
func (t_ TextView) RangeForUserCompletion() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("rangeForUserCompletion"))
	return rv
}


// SetRangeForUserCompletion sets the value of the rangeForUserCompletion property.
// The partial range from the most recent beginning of a word up to the insertion point.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusercompletion
func (t_ TextView) SetRangeForUserCompletion(value foundation.IRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangeForUserCompletion:"), value)
}

// The range of characters affected by an action method that changes paragraph (not character) attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforuserparagraphattributechange
func (t_ TextView) RangeForUserParagraphAttributeChange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("rangeForUserParagraphAttributeChange"))
	return rv
}


// SetRangeForUserParagraphAttributeChange sets the value of the rangeForUserParagraphAttributeChange property.
// The range of characters affected by an action method that changes paragraph (not character) attributes.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforuserparagraphattributechange
func (t_ TextView) SetRangeForUserParagraphAttributeChange(value foundation.IRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangeForUserParagraphAttributeChange:"), value)
}

// The range of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusertextchange
func (t_ TextView) RangeForUserTextChange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("rangeForUserTextChange"))
	return rv
}


// SetRangeForUserTextChange sets the value of the rangeForUserTextChange property.
// The range of characters affected by a method that changes characters (as opposed to attributes).

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusertextchange
func (t_ TextView) SetRangeForUserTextChange(value foundation.IRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangeForUserTextChange:"), value)
}

// An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforusercharacterattributechange
func (t_ TextView) RangesForUserCharacterAttributeChange() foundation.Value {
	rv := objc.Send[foundation.Value](t_.ID, objc.Sel("rangesForUserCharacterAttributeChange"))
	return rv
}


// SetRangesForUserCharacterAttributeChange sets the value of the rangesForUserCharacterAttributeChange property.
// An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforusercharacterattributechange
func (t_ TextView) SetRangesForUserCharacterAttributeChange(value foundation.IValue) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangesForUserCharacterAttributeChange:"), value)
}

// An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforuserparagraphattributechange
func (t_ TextView) RangesForUserParagraphAttributeChange() foundation.Value {
	rv := objc.Send[foundation.Value](t_.ID, objc.Sel("rangesForUserParagraphAttributeChange"))
	return rv
}


// SetRangesForUserParagraphAttributeChange sets the value of the rangesForUserParagraphAttributeChange property.
// An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforuserparagraphattributechange
func (t_ TextView) SetRangesForUserParagraphAttributeChange(value foundation.IValue) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangesForUserParagraphAttributeChange:"), value)
}

// An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforusertextchange
func (t_ TextView) RangesForUserTextChange() foundation.Value {
	rv := objc.Send[foundation.Value](t_.ID, objc.Sel("rangesForUserTextChange"))
	return rv
}


// SetRangesForUserTextChange sets the value of the rangesForUserTextChange property.
// An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforusertextchange
func (t_ TextView) SetRangesForUserTextChange(value foundation.IValue) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangesForUserTextChange:"), value)
}

// The types this text view can read immediately from the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/readablepasteboardtypes
func (t_ TextView) ReadablePasteboardTypes() PasteboardType {
	rv := objc.Send[PasteboardType](t_.ID, objc.Sel("readablePasteboardTypes"))
	return rv
}


// SetReadablePasteboardTypes sets the value of the readablePasteboardTypes property.
// The types this text view can read immediately from the pasteboard.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/readablepasteboardtypes
func (t_ TextView) SetReadablePasteboardTypes(value PasteboardType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReadablePasteboardTypes:"), value)
}

// An array containing the ranges of characters selected in the receiver’s layout manager.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectedranges
func (t_ TextView) SelectedRanges() foundation.Value {
	rv := objc.Send[foundation.Value](t_.ID, objc.Sel("selectedRanges"))
	return rv
}


// SetSelectedRanges sets the value of the selectedRanges property.
// An array containing the ranges of characters selected in the receiver’s layout manager.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectedranges
func (t_ TextView) SetSelectedRanges(value foundation.IValue) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRanges:"), value)
}

// The attributes used to indicate the selection.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectedtextattributes
func (t_ TextView) SelectedTextAttributes() coreml.Key {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("selectedTextAttributes"))
	return rv
}


// SetSelectedTextAttributes sets the value of the selectedTextAttributes property.
// The attributes used to indicate the selection.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectedtextattributes
func (t_ TextView) SetSelectedTextAttributes(value coreml.IKey) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedTextAttributes:"), value)
}

// The preferred direction of selection.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectionaffinity
func (t_ TextView) SelectionAffinity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("selectionAffinity"))
	return rv
}


// SetSelectionAffinity sets the value of the selectionAffinity property.
// The preferred direction of selection.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectionaffinity
func (t_ TextView) SetSelectionAffinity(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionAffinity:"), value)
}

// The selection granularity for subsequent extension of a selection.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectiongranularity
func (t_ TextView) SelectionGranularity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("selectionGranularity"))
	return rv
}


// SetSelectionGranularity sets the value of the selectionGranularity property.
// The selection granularity for subsequent extension of a selection.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectiongranularity
func (t_ TextView) SetSelectionGranularity(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionGranularity:"), value)
}

// A Boolean value that determines whether the receiver should draw its insertion point.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/shoulddrawinsertionpoint
func (t_ TextView) ShouldDrawInsertionPoint() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldDrawInsertionPoint"))
	return rv
}


// SetShouldDrawInsertionPoint sets the value of the shouldDrawInsertionPoint property.
// A Boolean value that determines whether the receiver should draw its insertion point.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/shoulddrawinsertionpoint
func (t_ TextView) SetShouldDrawInsertionPoint(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShouldDrawInsertionPoint:"), value)
}

// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/smartinsertdeleteenabled
func (t_ TextView) SmartInsertDeleteEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("smartInsertDeleteEnabled"))
	return rv
}


// SetSmartInsertDeleteEnabled sets the value of the smartInsertDeleteEnabled property.
// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/smartinsertdeleteenabled
func (t_ TextView) SetSmartInsertDeleteEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSmartInsertDeleteEnabled:"), value)
}

// A tag identifying the text view’s text as a document for the spell checker server.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/spellcheckerdocumenttag
func (t_ TextView) SpellCheckerDocumentTag() int {
	rv := objc.Send[int](t_.ID, objc.Sel("spellCheckerDocumentTag"))
	return rv
}


// SetSpellCheckerDocumentTag sets the value of the spellCheckerDocumentTag property.
// A tag identifying the text view’s text as a document for the spell checker server.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/spellcheckerdocumenttag
func (t_ TextView) SetSpellCheckerDocumentTag(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSpellCheckerDocumentTag:"), value)
}

// The empty space the receiver leaves around its associated text container.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontainerinset
func (t_ TextView) TextContainerInset() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("textContainerInset"))
	return rv
}


// SetTextContainerInset sets the value of the textContainerInset property.
// The empty space the receiver leaves around its associated text container.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontainerinset
func (t_ TextView) SetTextContainerInset(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainerInset:"), value)
}

// The origin of the receiver’s text container.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontainerorigin
func (t_ TextView) TextContainerOrigin() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](t_.ID, objc.Sel("textContainerOrigin"))
	return rv
}


// SetTextContainerOrigin sets the value of the textContainerOrigin property.
// The origin of the receiver’s text container.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontainerorigin
func (t_ TextView) SetTextContainerOrigin(value coregraphics.CGPoint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainerOrigin:"), value)
}

// The receiver’s text storage object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontentstorage
func (t_ TextView) TextContentStorage() NSTextContentStorage {
	rv := objc.Send[NSTextContentStorage](t_.ID, objc.Sel("textContentStorage"))
	return rv
}


// SetTextContentStorage sets the value of the textContentStorage property.
// The receiver’s text storage object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontentstorage
func (t_ TextView) SetTextContentStorage(value ITextContentStorage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContentStorage:"), value)
}

// ************************* Text Highlight support **************************
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/texthighlightattributes
func (t_ TextView) TextHighlightAttributes() coreml.Key {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("textHighlightAttributes"))
	return rv
}


// SetTextHighlightAttributes sets the value of the textHighlightAttributes property.
// ************************* Text Highlight support **************************

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/texthighlightattributes
func (t_ TextView) SetTextHighlightAttributes(value coreml.IKey) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextHighlightAttributes:"), value)
}

// The manager that lays out text for the receiver’s text container.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textlayoutmanager
func (t_ TextView) TextLayoutManager() NSTextLayoutManager {
	rv := objc.Send[NSTextLayoutManager](t_.ID, objc.Sel("textLayoutManager"))
	return rv
}


// SetTextLayoutManager sets the value of the textLayoutManager property.
// The manager that lays out text for the receiver’s text container.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textlayoutmanager
func (t_ TextView) SetTextLayoutManager(value ITextLayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextLayoutManager:"), value)
}

// The receiver’s typing attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/typingattributes
func (t_ TextView) TypingAttributes() coreml.Key {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("typingAttributes"))
	return rv
}


// SetTypingAttributes sets the value of the typingAttributes property.
// The receiver’s typing attributes.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/typingattributes
func (t_ TextView) SetTypingAttributes(value coreml.IKey) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypingAttributes:"), value)
}

// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesadaptivecolormappingfordarkappearance
func (t_ TextView) UsesAdaptiveColorMappingForDarkAppearance() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesAdaptiveColorMappingForDarkAppearance"))
	return rv
}


// SetUsesAdaptiveColorMappingForDarkAppearance sets the value of the usesAdaptiveColorMappingForDarkAppearance property.
// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesadaptivecolormappingfordarkappearance
func (t_ TextView) SetUsesAdaptiveColorMappingForDarkAppearance(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesAdaptiveColorMappingForDarkAppearance:"), value)
}

// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindbar
func (t_ TextView) UsesFindBar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFindBar"))
	return rv
}


// SetUsesFindBar sets the value of the usesFindBar property.
// A Boolean value that indicates whether to use the find bar for this text view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindbar
func (t_ TextView) SetUsesFindBar(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFindBar:"), value)
}

// A Boolean value that indicates whether the receiver allows for a find panel.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindpanel
func (t_ TextView) UsesFindPanel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFindPanel"))
	return rv
}


// SetUsesFindPanel sets the value of the usesFindPanel property.
// A Boolean value that indicates whether the receiver allows for a find panel.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindpanel
func (t_ TextView) SetUsesFindPanel(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFindPanel:"), value)
}

// A Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfontpanel
func (t_ TextView) UsesFontPanel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontPanel"))
	return rv
}


// SetUsesFontPanel sets the value of the usesFontPanel property.
// A Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfontpanel
func (t_ TextView) SetUsesFontPanel(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontPanel:"), value)
}

// A Boolean value that indicates whether this text view uses the inspector bar.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesinspectorbar
func (t_ TextView) UsesInspectorBar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesInspectorBar"))
	return rv
}


// SetUsesInspectorBar sets the value of the usesInspectorBar property.
// A Boolean value that indicates whether this text view uses the inspector bar.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesinspectorbar
func (t_ TextView) SetUsesInspectorBar(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesInspectorBar:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesrolloverbuttonforselection
func (t_ TextView) UsesRolloverButtonForSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesRolloverButtonForSelection"))
	return rv
}


// SetUsesRolloverButtonForSelection sets the value of the usesRolloverButtonForSelection property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesrolloverbuttonforselection
func (t_ TextView) SetUsesRolloverButtonForSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesRolloverButtonForSelection:"), value)
}

// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesruler
func (t_ TextView) UsesRuler() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesRuler"))
	return rv
}


// SetUsesRuler sets the value of the usesRuler property.
// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesruler
func (t_ TextView) SetUsesRuler(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesRuler:"), value)
}

// The pasteboard types that can be provided from the current selection.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/writablepasteboardtypes
func (t_ TextView) WritablePasteboardTypes() PasteboardType {
	rv := objc.Send[PasteboardType](t_.ID, objc.Sel("writablePasteboardTypes"))
	return rv
}


// SetWritablePasteboardTypes sets the value of the writablePasteboardTypes property.
// The pasteboard types that can be provided from the current selection.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/writablepasteboardtypes
func (t_ TextView) SetWritablePasteboardTypes(value PasteboardType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWritablePasteboardTypes:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/writingtoolsbehavior
func (t_ TextView) WritingToolsBehavior() WritingToolsBehavior {
	rv := objc.Send[WritingToolsBehavior](t_.ID, objc.Sel("writingToolsBehavior"))
	return rv
}


// SetWritingToolsBehavior sets the value of the writingToolsBehavior property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/writingtoolsbehavior
func (t_ TextView) SetWritingToolsBehavior(value WritingToolsBehavior) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWritingToolsBehavior:"), value)
}



