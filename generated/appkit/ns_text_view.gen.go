// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	// properties:
	AcceptableDragTypes() unsafe.Pointer
	SetAcceptableDragTypes(value unsafe.Pointer)
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	AllowedInputSourceLocales() objc.IObject /* cross-framework: NSString */
	SetAllowedInputSourceLocales(value objc.IObject /* cross-framework: NSString */)
	AllowedWritingToolsResultOptions() WritingToolsResultOptions
	SetAllowedWritingToolsResultOptions(value WritingToolsResultOptions)
	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)
	AllowsDocumentBackgroundColorChange() bool
	SetAllowsDocumentBackgroundColorChange(value bool)
	AllowsImageEditing() bool
	SetAllowsImageEditing(value bool)
	AllowsUndo() bool
	SetAllowsUndo(value bool)
	BackgroundColor() objc.IObject /* cross-framework: Color */
	SetBackgroundColor(value objc.IObject /* cross-framework: Color */)
	CandidateListTouchBarItem() ICandidateListTouchBarItem
	SetCandidateListTouchBarItem(value ICandidateListTouchBarItem)
	DefaultParagraphStyle() objc.IObject /* cross-framework: ParagraphStyle */
	SetDefaultParagraphStyle(value objc.IObject /* cross-framework: ParagraphStyle */)
	Delegate() TextViewDelegate /* not a class type */
	SetDelegate(value TextViewDelegate /* not a class type */)
	DisplaysLinkToolTips() bool
	SetDisplaysLinkToolTips(value bool)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	EnabledTextCheckingTypes() TextCheckingTypes /* not a class type */
	SetEnabledTextCheckingTypes(value TextCheckingTypes /* not a class type */)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	InlinePredictionType() TextInputTraitType /* not a class type */
	SetInlinePredictionType(value TextInputTraitType /* not a class type */)
	InsertionPointColor() objc.IObject /* cross-framework: Color */
	SetInsertionPointColor(value objc.IObject /* cross-framework: Color */)
	IsAutomaticDashSubstitutionEnabled() bool
	SetIsAutomaticDashSubstitutionEnabled(value bool)
	IsAutomaticDataDetectionEnabled() bool
	SetIsAutomaticDataDetectionEnabled(value bool)
	IsAutomaticLinkDetectionEnabled() bool
	SetIsAutomaticLinkDetectionEnabled(value bool)
	IsAutomaticQuoteSubstitutionEnabled() bool
	SetIsAutomaticQuoteSubstitutionEnabled(value bool)
	IsAutomaticSpellingCorrectionEnabled() bool
	SetIsAutomaticSpellingCorrectionEnabled(value bool)
	IsAutomaticTextCompletionEnabled() bool
	SetIsAutomaticTextCompletionEnabled(value bool)
	IsAutomaticTextReplacementEnabled() bool
	SetIsAutomaticTextReplacementEnabled(value bool)
	IsCoalescingUndo() bool
	SetIsCoalescingUndo(value bool)
	IsContinuousSpellCheckingEnabled() bool
	SetIsContinuousSpellCheckingEnabled(value bool)
	IsEditable() bool
	SetIsEditable(value bool)
	IsFieldEditor() bool
	SetIsFieldEditor(value bool)
	IsGrammarCheckingEnabled() bool
	SetIsGrammarCheckingEnabled(value bool)
	IsIncrementalSearchingEnabled() bool
	SetIsIncrementalSearchingEnabled(value bool)
	IsRichText() bool
	SetIsRichText(value bool)
	IsRulerVisible() bool
	SetIsRulerVisible(value bool)
	IsSelectable() bool
	SetIsSelectable(value bool)
	IsWritingToolsActive() bool
	SetIsWritingToolsActive(value bool)
	LayoutManager() objc.IObject /* cross-framework: LayoutManager */
	SetLayoutManager(value objc.IObject /* cross-framework: LayoutManager */)
	LinkTextAttributes() objc.IObject /* cross-framework: Key */
	SetLinkTextAttributes(value objc.IObject /* cross-framework: Key */)
	MarkedTextAttributes() objc.IObject /* cross-framework: Key */
	SetMarkedTextAttributes(value objc.IObject /* cross-framework: Key */)
	MathExpressionCompletionType() TextInputTraitType /* not a class type */
	SetMathExpressionCompletionType(value TextInputTraitType /* not a class type */)
	RangeForUserCharacterAttributeChange() objc.IObject /* cross-framework: Range */
	SetRangeForUserCharacterAttributeChange(value objc.IObject /* cross-framework: Range */)
	RangeForUserCompletion() objc.IObject /* cross-framework: Range */
	SetRangeForUserCompletion(value objc.IObject /* cross-framework: Range */)
	RangeForUserParagraphAttributeChange() objc.IObject /* cross-framework: Range */
	SetRangeForUserParagraphAttributeChange(value objc.IObject /* cross-framework: Range */)
	RangeForUserTextChange() objc.IObject /* cross-framework: Range */
	SetRangeForUserTextChange(value objc.IObject /* cross-framework: Range */)
	RangesForUserCharacterAttributeChange() objc.IObject /* cross-framework: Value */
	SetRangesForUserCharacterAttributeChange(value objc.IObject /* cross-framework: Value */)
	RangesForUserParagraphAttributeChange() objc.IObject /* cross-framework: Value */
	SetRangesForUserParagraphAttributeChange(value objc.IObject /* cross-framework: Value */)
	RangesForUserTextChange() objc.IObject /* cross-framework: Value */
	SetRangesForUserTextChange(value objc.IObject /* cross-framework: Value */)
	ReadablePasteboardTypes() unsafe.Pointer
	SetReadablePasteboardTypes(value unsafe.Pointer)
	SelectedRanges() objc.IObject /* cross-framework: Value */
	SetSelectedRanges(value objc.IObject /* cross-framework: Value */)
	SelectedTextAttributes() objc.IObject /* cross-framework: Key */
	SetSelectedTextAttributes(value objc.IObject /* cross-framework: Key */)
	SelectionAffinity() SelectionAffinity /* not a class type */
	SetSelectionAffinity(value SelectionAffinity /* not a class type */)
	SelectionGranularity() SelectionGranularity /* not a class type */
	SetSelectionGranularity(value SelectionGranularity /* not a class type */)
	ShouldDrawInsertionPoint() bool
	SetShouldDrawInsertionPoint(value bool)
	SmartInsertDeleteEnabled() bool
	SetSmartInsertDeleteEnabled(value bool)
	SpellCheckerDocumentTag() int
	SetSpellCheckerDocumentTag(value int)
	TextContainer() objc.IObject /* cross-framework: TextContainer */
	SetTextContainer(value objc.IObject /* cross-framework: TextContainer */)
	TextContainerInset() objc.IObject /* cross-framework: Size */
	SetTextContainerInset(value objc.IObject /* cross-framework: Size */)
	TextContainerOrigin() objc.IObject /* cross-framework: Point */
	SetTextContainerOrigin(value objc.IObject /* cross-framework: Point */)
	TextContentStorage() ITextContentStorage
	SetTextContentStorage(value ITextContentStorage)
	TextHighlightAttributes() objc.IObject /* cross-framework: Key */
	SetTextHighlightAttributes(value objc.IObject /* cross-framework: Key */)
	TextLayoutManager() objc.IObject /* cross-framework: TextLayoutManager */
	SetTextLayoutManager(value objc.IObject /* cross-framework: TextLayoutManager */)
	TextStorage() ITextStorage
	SetTextStorage(value ITextStorage)
	TypingAttributes() objc.IObject /* cross-framework: Key */
	SetTypingAttributes(value objc.IObject /* cross-framework: Key */)
	UsesAdaptiveColorMappingForDarkAppearance() bool
	SetUsesAdaptiveColorMappingForDarkAppearance(value bool)
	UsesFindBar() bool
	SetUsesFindBar(value bool)
	UsesFindPanel() bool
	SetUsesFindPanel(value bool)
	UsesFontPanel() bool
	SetUsesFontPanel(value bool)
	UsesInspectorBar() bool
	SetUsesInspectorBar(value bool)
	UsesRolloverButtonForSelection() bool
	SetUsesRolloverButtonForSelection(value bool)
	UsesRuler() bool
	SetUsesRuler(value bool)
	WritablePasteboardTypes() unsafe.Pointer
	SetWritablePasteboardTypes(value unsafe.Pointer)
	WritingToolsBehavior() WritingToolsBehavior
	SetWritingToolsBehavior(value WritingToolsBehavior)
	// methods:
}

// A view that draws text and handles user interactions with that text.
//
// The class is the front-end class to the AppKit text system. The class draws the text managed by the back-end components and handles user events to select and modify its text, in addition to supporting rich text, attachments, input management, and key binding, and marked text attributes. is the principal means to obtain a text object that caters to almost all needs for displaying and managing text at the user interface level. While is a subclass of the class — which declares the most general Cocoa interface to the text system — adds major features beyond the capabilities of . You can also do more powerful and more creative text manipulation (such as displaying text in a circle) using , , , and related classes. You’re more likely to use the class than . It’s also important to remember that conforms to a large number of protocols, the methods of which are available to instances of the class. communicates with its delegate through methods declared both by the and by its superclass’s protocol, . All delegation messages come from the first text view. In macOS 12 and later, if you explicitly call the property on a text view or text container, the framework reverts to a compatibility mode that uses . The text view also switches to this compatibility mode when it encounters text content that’s not yet supported, such as .


// A view that draws text and handles user interactions with that text.
//
// [Full Topic]
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



// The data types that the receiver accepts as the destination view of a dragging operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/acceptabledragtypes
func (t_ TextView) AcceptableDragTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("acceptableDragTypes"))
	return rv
}


// The data types that the receiver accepts as the destination view of a dragging operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/acceptabledragtypes
func (t_ TextView) SetAcceptableDragTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptableDragTypes:"), value)
}


// A Boolean value that indicates whether the receiver accepts the glyph info attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/acceptsglyphinfo
func (t_ TextView) AcceptsGlyphInfo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("acceptsGlyphInfo"))
	return rv
}


// A Boolean value that indicates whether the receiver accepts the glyph info attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/acceptsglyphinfo
func (t_ TextView) SetAcceptsGlyphInfo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptsGlyphInfo:"), value)
}


// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowedinputsourcelocales
func (t_ TextView) AllowedInputSourceLocales() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("allowedInputSourceLocales"))
	return rv
}


// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowedinputsourcelocales
func (t_ TextView) SetAllowedInputSourceLocales(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedInputSourceLocales:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowedwritingtoolsresultoptions
func (t_ TextView) AllowedWritingToolsResultOptions() WritingToolsResultOptions {
	rv := objc.Send[WritingToolsResultOptions](t_.ID, objc.Sel("allowedWritingToolsResultOptions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowedwritingtoolsresultoptions
func (t_ TextView) SetAllowedWritingToolsResultOptions(value WritingToolsResultOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedWritingToolsResultOptions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowscharacterpickertouchbaritem
func (t_ TextView) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsCharacterPickerTouchBarItem"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowscharacterpickertouchbaritem
func (t_ TextView) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsCharacterPickerTouchBarItem:"), value)
}


// A Boolean value that indicates whether the receiver allows its background color to change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsdocumentbackgroundcolorchange
func (t_ TextView) AllowsDocumentBackgroundColorChange() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDocumentBackgroundColorChange"))
	return rv
}


// A Boolean value that indicates whether the receiver allows its background color to change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsdocumentbackgroundcolorchange
func (t_ TextView) SetAllowsDocumentBackgroundColorChange(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsDocumentBackgroundColorChange:"), value)
}


// Indicates whether image attachments should permit editing of their images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsimageediting
func (t_ TextView) AllowsImageEditing() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsImageEditing"))
	return rv
}


// Indicates whether image attachments should permit editing of their images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsimageediting
func (t_ TextView) SetAllowsImageEditing(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsImageEditing:"), value)
}


// A Boolean value that indicates whether the receiver allows undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsundo
func (t_ TextView) AllowsUndo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsUndo"))
	return rv
}


// A Boolean value that indicates whether the receiver allows undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsundo
func (t_ TextView) SetAllowsUndo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsUndo:"), value)
}


// The receiver’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/backgroundcolor
func (t_ TextView) BackgroundColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The receiver’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/backgroundcolor
func (t_ TextView) SetBackgroundColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/candidatelisttouchbaritem
func (t_ TextView) CandidateListTouchBarItem() ICandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](t_.ID, objc.Sel("candidateListTouchBarItem"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/candidatelisttouchbaritem
func (t_ TextView) SetCandidateListTouchBarItem(value ICandidateListTouchBarItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCandidateListTouchBarItem:"), value)
}


// The receiver’s default paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/defaultparagraphstyle
func (t_ TextView) DefaultParagraphStyle() objc.IObject /* cross-framework: ParagraphStyle */ {
	rv := objc.Send[ParagraphStyle](t_.ID, objc.Sel("defaultParagraphStyle"))
	return rv
}


// The receiver’s default paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/defaultparagraphstyle
func (t_ TextView) SetDefaultParagraphStyle(value objc.IObject /* cross-framework: ParagraphStyle */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDefaultParagraphStyle:"), value)
}


// The delegate for all text views sharing the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/delegate
func (t_ TextView) Delegate() TextViewDelegate /* not a class type */ {
	rv := objc.Send[TextViewDelegate](t_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for all text views sharing the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/delegate
func (t_ TextView) SetDelegate(value TextViewDelegate /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/displayslinktooltips
func (t_ TextView) DisplaysLinkToolTips() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("displaysLinkToolTips"))
	return rv
}


// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/displayslinktooltips
func (t_ TextView) SetDisplaysLinkToolTips(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDisplaysLinkToolTips:"), value)
}


// A Boolean value that indicates whether the receiver draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/drawsbackground
func (t_ TextView) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value that indicates whether the receiver draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/drawsbackground
func (t_ TextView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}


// The default text checking types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/enabledtextcheckingtypes
func (t_ TextView) EnabledTextCheckingTypes() TextCheckingTypes /* not a class type */ {
	rv := objc.Send[TextCheckingTypes](t_.ID, objc.Sel("enabledTextCheckingTypes"))
	return rv
}


// The default text checking types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/enabledtextcheckingtypes
func (t_ TextView) SetEnabledTextCheckingTypes(value TextCheckingTypes /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEnabledTextCheckingTypes:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/importsgraphics
func (t_ TextView) ImportsGraphics() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("importsGraphics"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/importsgraphics
func (t_ TextView) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImportsGraphics:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/inlinepredictiontype
func (t_ TextView) InlinePredictionType() TextInputTraitType /* not a class type */ {
	rv := objc.Send[TextInputTraitType](t_.ID, objc.Sel("inlinePredictionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/inlinepredictiontype
func (t_ TextView) SetInlinePredictionType(value TextInputTraitType /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInlinePredictionType:"), value)
}


// The color of the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/insertionpointcolor
func (t_ TextView) InsertionPointColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](t_.ID, objc.Sel("insertionPointColor"))
	return rv
}


// The color of the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/insertionpointcolor
func (t_ TextView) SetInsertionPointColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInsertionPointColor:"), value)
}


// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdashsubstitutionenabled
func (t_ TextView) IsAutomaticDashSubstitutionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticDashSubstitutionEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdashsubstitutionenabled
func (t_ TextView) SetIsAutomaticDashSubstitutionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticDashSubstitutionEnabled:"), value)
}


// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdatadetectionenabled
func (t_ TextView) IsAutomaticDataDetectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticDataDetectionEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdatadetectionenabled
func (t_ TextView) SetIsAutomaticDataDetectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticDataDetectionEnabled:"), value)
}


// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticlinkdetectionenabled
func (t_ TextView) IsAutomaticLinkDetectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticLinkDetectionEnabled"))
	return rv
}


// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticlinkdetectionenabled
func (t_ TextView) SetIsAutomaticLinkDetectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticLinkDetectionEnabled:"), value)
}


// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticquotesubstitutionenabled
func (t_ TextView) IsAutomaticQuoteSubstitutionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticQuoteSubstitutionEnabled"))
	return rv
}


// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticquotesubstitutionenabled
func (t_ TextView) SetIsAutomaticQuoteSubstitutionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticQuoteSubstitutionEnabled:"), value)
}


// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticspellingcorrectionenabled
func (t_ TextView) IsAutomaticSpellingCorrectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticSpellingCorrectionEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticspellingcorrectionenabled
func (t_ TextView) SetIsAutomaticSpellingCorrectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticSpellingCorrectionEnabled:"), value)
}


// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextcompletionenabled
func (t_ TextView) IsAutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticTextCompletionEnabled"))
	return rv
}


// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextcompletionenabled
func (t_ TextView) SetIsAutomaticTextCompletionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticTextCompletionEnabled:"), value)
}


// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextreplacementenabled
func (t_ TextView) IsAutomaticTextReplacementEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticTextReplacementEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextreplacementenabled
func (t_ TextView) SetIsAutomaticTextReplacementEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticTextReplacementEnabled:"), value)
}


// A Boolean value that indicates whether undo coalescing is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscoalescingundo
func (t_ TextView) IsCoalescingUndo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isCoalescingUndo"))
	return rv
}


// A Boolean value that indicates whether undo coalescing is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscoalescingundo
func (t_ TextView) SetIsCoalescingUndo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsCoalescingUndo:"), value)
}


// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscontinuousspellcheckingenabled
func (t_ TextView) IsContinuousSpellCheckingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isContinuousSpellCheckingEnabled"))
	return rv
}


// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscontinuousspellcheckingenabled
func (t_ TextView) SetIsContinuousSpellCheckingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsContinuousSpellCheckingEnabled:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iseditable
func (t_ TextView) IsEditable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEditable"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iseditable
func (t_ TextView) SetIsEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEditable:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isfieldeditor
func (t_ TextView) IsFieldEditor() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFieldEditor"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isfieldeditor
func (t_ TextView) SetIsFieldEditor(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFieldEditor:"), value)
}


// Enables and disables grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isgrammarcheckingenabled
func (t_ TextView) IsGrammarCheckingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isGrammarCheckingEnabled"))
	return rv
}


// Enables and disables grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isgrammarcheckingenabled
func (t_ TextView) SetIsGrammarCheckingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsGrammarCheckingEnabled:"), value)
}


// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isincrementalsearchingenabled
func (t_ TextView) IsIncrementalSearchingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isIncrementalSearchingEnabled"))
	return rv
}


// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isincrementalsearchingenabled
func (t_ TextView) SetIsIncrementalSearchingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsIncrementalSearchingEnabled:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrichtext
func (t_ TextView) IsRichText() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRichText"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrichtext
func (t_ TextView) SetIsRichText(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRichText:"), value)
}


// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrulervisible
func (t_ TextView) IsRulerVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRulerVisible"))
	return rv
}


// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrulervisible
func (t_ TextView) SetIsRulerVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRulerVisible:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isselectable
func (t_ TextView) IsSelectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelectable"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isselectable
func (t_ TextView) SetIsSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSelectable:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iswritingtoolsactive
func (t_ TextView) IsWritingToolsActive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isWritingToolsActive"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iswritingtoolsactive
func (t_ TextView) SetIsWritingToolsActive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsWritingToolsActive:"), value)
}


// The layout manager that lays out text for the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/layoutmanager
func (t_ TextView) LayoutManager() objc.IObject /* cross-framework: LayoutManager */ {
	rv := objc.Send[LayoutManager](t_.ID, objc.Sel("layoutManager"))
	return rv
}


// The layout manager that lays out text for the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/layoutmanager
func (t_ TextView) SetLayoutManager(value objc.IObject /* cross-framework: LayoutManager */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutManager:"), value)
}


// The attributes used to draw the onscreen presentation of link text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/linktextattributes
func (t_ TextView) LinkTextAttributes() objc.IObject /* cross-framework: Key */ {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("linkTextAttributes"))
	return rv
}


// The attributes used to draw the onscreen presentation of link text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/linktextattributes
func (t_ TextView) SetLinkTextAttributes(value objc.IObject /* cross-framework: Key */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLinkTextAttributes:"), value)
}


// The attributes used to draw marked text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/markedtextattributes
func (t_ TextView) MarkedTextAttributes() objc.IObject /* cross-framework: Key */ {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("markedTextAttributes"))
	return rv
}


// The attributes used to draw marked text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/markedtextattributes
func (t_ TextView) SetMarkedTextAttributes(value objc.IObject /* cross-framework: Key */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMarkedTextAttributes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/mathexpressioncompletiontype
func (t_ TextView) MathExpressionCompletionType() TextInputTraitType /* not a class type */ {
	rv := objc.Send[TextInputTraitType](t_.ID, objc.Sel("mathExpressionCompletionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/mathexpressioncompletiontype
func (t_ TextView) SetMathExpressionCompletionType(value TextInputTraitType /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMathExpressionCompletionType:"), value)
}


// The range of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusercharacterattributechange
func (t_ TextView) RangeForUserCharacterAttributeChange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rangeForUserCharacterAttributeChange"))
	return rv
}


// The range of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusercharacterattributechange
func (t_ TextView) SetRangeForUserCharacterAttributeChange(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangeForUserCharacterAttributeChange:"), value)
}


// The partial range from the most recent beginning of a word up to the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusercompletion
func (t_ TextView) RangeForUserCompletion() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rangeForUserCompletion"))
	return rv
}


// The partial range from the most recent beginning of a word up to the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusercompletion
func (t_ TextView) SetRangeForUserCompletion(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangeForUserCompletion:"), value)
}


// The range of characters affected by an action method that changes paragraph (not character) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforuserparagraphattributechange
func (t_ TextView) RangeForUserParagraphAttributeChange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rangeForUserParagraphAttributeChange"))
	return rv
}


// The range of characters affected by an action method that changes paragraph (not character) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforuserparagraphattributechange
func (t_ TextView) SetRangeForUserParagraphAttributeChange(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangeForUserParagraphAttributeChange:"), value)
}


// The range of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusertextchange
func (t_ TextView) RangeForUserTextChange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rangeForUserTextChange"))
	return rv
}


// The range of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusertextchange
func (t_ TextView) SetRangeForUserTextChange(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangeForUserTextChange:"), value)
}


// An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforusercharacterattributechange
func (t_ TextView) RangesForUserCharacterAttributeChange() objc.IObject /* cross-framework: Value */ {
	rv := objc.Send[foundation.Value](t_.ID, objc.Sel("rangesForUserCharacterAttributeChange"))
	return rv
}


// An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforusercharacterattributechange
func (t_ TextView) SetRangesForUserCharacterAttributeChange(value objc.IObject /* cross-framework: Value */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangesForUserCharacterAttributeChange:"), value)
}


// An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforuserparagraphattributechange
func (t_ TextView) RangesForUserParagraphAttributeChange() objc.IObject /* cross-framework: Value */ {
	rv := objc.Send[foundation.Value](t_.ID, objc.Sel("rangesForUserParagraphAttributeChange"))
	return rv
}


// An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforuserparagraphattributechange
func (t_ TextView) SetRangesForUserParagraphAttributeChange(value objc.IObject /* cross-framework: Value */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangesForUserParagraphAttributeChange:"), value)
}


// An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforusertextchange
func (t_ TextView) RangesForUserTextChange() objc.IObject /* cross-framework: Value */ {
	rv := objc.Send[foundation.Value](t_.ID, objc.Sel("rangesForUserTextChange"))
	return rv
}


// An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforusertextchange
func (t_ TextView) SetRangesForUserTextChange(value objc.IObject /* cross-framework: Value */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangesForUserTextChange:"), value)
}


// The types this text view can read immediately from the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/readablepasteboardtypes
func (t_ TextView) ReadablePasteboardTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("readablePasteboardTypes"))
	return rv
}


// The types this text view can read immediately from the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/readablepasteboardtypes
func (t_ TextView) SetReadablePasteboardTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReadablePasteboardTypes:"), value)
}


// An array containing the ranges of characters selected in the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectedranges
func (t_ TextView) SelectedRanges() objc.IObject /* cross-framework: Value */ {
	rv := objc.Send[foundation.Value](t_.ID, objc.Sel("selectedRanges"))
	return rv
}


// An array containing the ranges of characters selected in the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectedranges
func (t_ TextView) SetSelectedRanges(value objc.IObject /* cross-framework: Value */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRanges:"), value)
}


// The attributes used to indicate the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectedtextattributes
func (t_ TextView) SelectedTextAttributes() objc.IObject /* cross-framework: Key */ {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("selectedTextAttributes"))
	return rv
}


// The attributes used to indicate the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectedtextattributes
func (t_ TextView) SetSelectedTextAttributes(value objc.IObject /* cross-framework: Key */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedTextAttributes:"), value)
}


// The preferred direction of selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectionaffinity
func (t_ TextView) SelectionAffinity() SelectionAffinity /* not a class type */ {
	rv := objc.Send[SelectionAffinity](t_.ID, objc.Sel("selectionAffinity"))
	return rv
}


// The preferred direction of selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectionaffinity
func (t_ TextView) SetSelectionAffinity(value SelectionAffinity /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionAffinity:"), value)
}


// The selection granularity for subsequent extension of a selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectiongranularity
func (t_ TextView) SelectionGranularity() SelectionGranularity /* not a class type */ {
	rv := objc.Send[SelectionGranularity](t_.ID, objc.Sel("selectionGranularity"))
	return rv
}


// The selection granularity for subsequent extension of a selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectiongranularity
func (t_ TextView) SetSelectionGranularity(value SelectionGranularity /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionGranularity:"), value)
}


// A Boolean value that determines whether the receiver should draw its insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/shoulddrawinsertionpoint
func (t_ TextView) ShouldDrawInsertionPoint() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldDrawInsertionPoint"))
	return rv
}


// A Boolean value that determines whether the receiver should draw its insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/shoulddrawinsertionpoint
func (t_ TextView) SetShouldDrawInsertionPoint(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShouldDrawInsertionPoint:"), value)
}


// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/smartinsertdeleteenabled
func (t_ TextView) SmartInsertDeleteEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("smartInsertDeleteEnabled"))
	return rv
}


// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/smartinsertdeleteenabled
func (t_ TextView) SetSmartInsertDeleteEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSmartInsertDeleteEnabled:"), value)
}


// A tag identifying the text view’s text as a document for the spell checker server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/spellcheckerdocumenttag
func (t_ TextView) SpellCheckerDocumentTag() int {
	rv := objc.Send[int](t_.ID, objc.Sel("spellCheckerDocumentTag"))
	return rv
}


// A tag identifying the text view’s text as a document for the spell checker server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/spellcheckerdocumenttag
func (t_ TextView) SetSpellCheckerDocumentTag(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSpellCheckerDocumentTag:"), value)
}


// The receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontainer
func (t_ TextView) TextContainer() objc.IObject /* cross-framework: TextContainer */ {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("textContainer"))
	return rv
}


// The receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontainer
func (t_ TextView) SetTextContainer(value objc.IObject /* cross-framework: TextContainer */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainer:"), value)
}


// The empty space the receiver leaves around its associated text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontainerinset
func (t_ TextView) TextContainerInset() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](t_.ID, objc.Sel("textContainerInset"))
	return rv
}


// The empty space the receiver leaves around its associated text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontainerinset
func (t_ TextView) SetTextContainerInset(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainerInset:"), value)
}


// The origin of the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontainerorigin
func (t_ TextView) TextContainerOrigin() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[corefoundation.Point](t_.ID, objc.Sel("textContainerOrigin"))
	return rv
}


// The origin of the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontainerorigin
func (t_ TextView) SetTextContainerOrigin(value objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainerOrigin:"), value)
}


// The receiver’s text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontentstorage
func (t_ TextView) TextContentStorage() ITextContentStorage {
	rv := objc.Send[TextContentStorage](t_.ID, objc.Sel("textContentStorage"))
	return rv
}


// The receiver’s text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontentstorage
func (t_ TextView) SetTextContentStorage(value ITextContentStorage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContentStorage:"), value)
}


// ************************* Text Highlight support **************************
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/texthighlightattributes
func (t_ TextView) TextHighlightAttributes() objc.IObject /* cross-framework: Key */ {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("textHighlightAttributes"))
	return rv
}


// ************************* Text Highlight support **************************
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/texthighlightattributes
func (t_ TextView) SetTextHighlightAttributes(value objc.IObject /* cross-framework: Key */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextHighlightAttributes:"), value)
}


// The manager that lays out text for the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textlayoutmanager
func (t_ TextView) TextLayoutManager() objc.IObject /* cross-framework: TextLayoutManager */ {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("textLayoutManager"))
	return rv
}


// The manager that lays out text for the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textlayoutmanager
func (t_ TextView) SetTextLayoutManager(value objc.IObject /* cross-framework: TextLayoutManager */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextLayoutManager:"), value)
}


// The receiver’s text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textstorage
func (t_ TextView) TextStorage() ITextStorage {
	rv := objc.Send[TextStorage](t_.ID, objc.Sel("textStorage"))
	return rv
}


// The receiver’s text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textstorage
func (t_ TextView) SetTextStorage(value ITextStorage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextStorage:"), value)
}


// The receiver’s typing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/typingattributes
func (t_ TextView) TypingAttributes() objc.IObject /* cross-framework: Key */ {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("typingAttributes"))
	return rv
}


// The receiver’s typing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/typingattributes
func (t_ TextView) SetTypingAttributes(value objc.IObject /* cross-framework: Key */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypingAttributes:"), value)
}


// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesadaptivecolormappingfordarkappearance
func (t_ TextView) UsesAdaptiveColorMappingForDarkAppearance() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesAdaptiveColorMappingForDarkAppearance"))
	return rv
}


// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesadaptivecolormappingfordarkappearance
func (t_ TextView) SetUsesAdaptiveColorMappingForDarkAppearance(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesAdaptiveColorMappingForDarkAppearance:"), value)
}


// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindbar
func (t_ TextView) UsesFindBar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFindBar"))
	return rv
}


// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindbar
func (t_ TextView) SetUsesFindBar(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFindBar:"), value)
}


// A Boolean value that indicates whether the receiver allows for a find panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindpanel
func (t_ TextView) UsesFindPanel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFindPanel"))
	return rv
}


// A Boolean value that indicates whether the receiver allows for a find panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindpanel
func (t_ TextView) SetUsesFindPanel(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFindPanel:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfontpanel
func (t_ TextView) UsesFontPanel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontPanel"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfontpanel
func (t_ TextView) SetUsesFontPanel(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontPanel:"), value)
}


// A Boolean value that indicates whether this text view uses the inspector bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesinspectorbar
func (t_ TextView) UsesInspectorBar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesInspectorBar"))
	return rv
}


// A Boolean value that indicates whether this text view uses the inspector bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesinspectorbar
func (t_ TextView) SetUsesInspectorBar(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesInspectorBar:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesrolloverbuttonforselection
func (t_ TextView) UsesRolloverButtonForSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesRolloverButtonForSelection"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesrolloverbuttonforselection
func (t_ TextView) SetUsesRolloverButtonForSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesRolloverButtonForSelection:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesruler
func (t_ TextView) UsesRuler() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesRuler"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesruler
func (t_ TextView) SetUsesRuler(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesRuler:"), value)
}


// The pasteboard types that can be provided from the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/writablepasteboardtypes
func (t_ TextView) WritablePasteboardTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("writablePasteboardTypes"))
	return rv
}


// The pasteboard types that can be provided from the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/writablepasteboardtypes
func (t_ TextView) SetWritablePasteboardTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWritablePasteboardTypes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/writingtoolsbehavior
func (t_ TextView) WritingToolsBehavior() WritingToolsBehavior {
	rv := objc.Send[WritingToolsBehavior](t_.ID, objc.Sel("writingToolsBehavior"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/writingtoolsbehavior
func (t_ TextView) SetWritingToolsBehavior(value WritingToolsBehavior) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWritingToolsBehavior:"), value)
}



