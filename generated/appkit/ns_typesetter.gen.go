// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Typesetter] class.
var (
	TypesetterClass     _TypesetterClass
	TypesetterClassOnce sync.Once
)

func getTypesetterClass() _TypesetterClass {
	TypesetterClassOnce.Do(func() {
		TypesetterClass = _TypesetterClass{objc.GetClass("NSTypesetter")}
	})
	return TypesetterClass
}

type _TypesetterClass struct {
	class objc.Class
}

// An interface definition for the [Typesetter] class.
type ITypesetter interface {
	objectivec.IObject
	// properties:
	AttributedString() objc.IObject /* cross-framework: AttributedString */
	SetAttributedString(value objc.IObject /* cross-framework: AttributedString */)
	AttributesForExtraLineFragment() objc.IObject /* cross-framework: Key */
	SetAttributesForExtraLineFragment(value objc.IObject /* cross-framework: Key */)
	BidiProcessingEnabled() bool
	SetBidiProcessingEnabled(value bool)
	CurrentParagraphStyle() objc.IObject /* cross-framework: ParagraphStyle */
	SetCurrentParagraphStyle(value objc.IObject /* cross-framework: ParagraphStyle */)
	CurrentTextContainer() objc.IObject /* cross-framework: TextContainer */
	SetCurrentTextContainer(value objc.IObject /* cross-framework: TextContainer */)
	HyphenationFactor() float32
	SetHyphenationFactor(value float32)
	LayoutManager() objc.IObject /* cross-framework: LayoutManager */
	SetLayoutManager(value objc.IObject /* cross-framework: LayoutManager */)
	LineFragmentPadding() float64
	SetLineFragmentPadding(value float64)
	ParagraphCharacterRange() objc.IObject /* cross-framework: Range */
	SetParagraphCharacterRange(value objc.IObject /* cross-framework: Range */)
	ParagraphGlyphRange() objc.IObject /* cross-framework: Range */
	SetParagraphGlyphRange(value objc.IObject /* cross-framework: Range */)
	ParagraphSeparatorCharacterRange() objc.IObject /* cross-framework: Range */
	SetParagraphSeparatorCharacterRange(value objc.IObject /* cross-framework: Range */)
	ParagraphSeparatorGlyphRange() objc.IObject /* cross-framework: Range */
	SetParagraphSeparatorGlyphRange(value objc.IObject /* cross-framework: Range */)
	TextContainers() objc.IObject /* cross-framework: NSArray */
	SetTextContainers(value objc.IObject /* cross-framework: NSArray */)
	TypesetterBehavior() unsafe.Pointer
	SetTypesetterBehavior(value unsafe.Pointer)
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	// methods:
}

// An abstract class that performs various type layout tasks.
//
// uses concrete subclasses of to perform line layout, which includes word wrapping, hyphenation, and line breaking in either vertical or horizontal rectangles. By default, the text system uses the concrete subclass .


// An abstract class that performs various type layout tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter
type Typesetter struct {
	objectivec.Object
}

// TypesetterFrom constructs a [Typesetter] from an unsafe.Pointer.
//
// An abstract class that performs various type layout tasks.
func TypesetterFrom(ptr unsafe.Pointer) Typesetter {
	return Typesetter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TypesetterClass) Alloc() Typesetter {
	rv := objc.Send[Typesetter](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TypesetterClass) New() Typesetter {
	rv := objc.Send[Typesetter](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Typesetter) Init() Typesetter {
	rv := objc.Send[Typesetter](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Typesetter) Autorelease() Typesetter {
	rv := objc.Send[Typesetter](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTypesetter creates a new Typesetter instance.
func NewTypesetter() Typesetter {
	return getTypesetterClass().New()
}



// Returns the text backing store, usually an instance of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/attributedstring
func (t_ Typesetter) AttributedString() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("attributedString"))
	return rv
}


// Returns the text backing store, usually an instance of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/attributedstring
func (t_ Typesetter) SetAttributedString(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributedString:"), value)
}


// Returns the attributes used to lay out the extra line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/attributesforextralinefragment
func (t_ Typesetter) AttributesForExtraLineFragment() objc.IObject /* cross-framework: Key */ {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("attributesForExtraLineFragment"))
	return rv
}


// Returns the attributes used to lay out the extra line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/attributesforextralinefragment
func (t_ Typesetter) SetAttributesForExtraLineFragment(value objc.IObject /* cross-framework: Key */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributesForExtraLineFragment:"), value)
}


// Returns whether bidirectional text processing is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/bidiprocessingenabled
func (t_ Typesetter) BidiProcessingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("bidiProcessingEnabled"))
	return rv
}


// Returns whether bidirectional text processing is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/bidiprocessingenabled
func (t_ Typesetter) SetBidiProcessingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBidiProcessingEnabled:"), value)
}


// Returns the paragraph style object for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/currentparagraphstyle
func (t_ Typesetter) CurrentParagraphStyle() objc.IObject /* cross-framework: ParagraphStyle */ {
	rv := objc.Send[ParagraphStyle](t_.ID, objc.Sel("currentParagraphStyle"))
	return rv
}


// Returns the paragraph style object for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/currentparagraphstyle
func (t_ Typesetter) SetCurrentParagraphStyle(value objc.IObject /* cross-framework: ParagraphStyle */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentParagraphStyle:"), value)
}


// Returns the text container for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/currenttextcontainer
func (t_ Typesetter) CurrentTextContainer() objc.IObject /* cross-framework: TextContainer */ {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("currentTextContainer"))
	return rv
}


// Returns the text container for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/currenttextcontainer
func (t_ Typesetter) SetCurrentTextContainer(value objc.IObject /* cross-framework: TextContainer */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentTextContainer:"), value)
}


// Returns the current hyphenation factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/hyphenationfactor
func (t_ Typesetter) HyphenationFactor() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("hyphenationFactor"))
	return rv
}


// Returns the current hyphenation factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/hyphenationfactor
func (t_ Typesetter) SetHyphenationFactor(value float32) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHyphenationFactor:"), value)
}


// Returns the layout manager for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/layoutmanager
func (t_ Typesetter) LayoutManager() objc.IObject /* cross-framework: LayoutManager */ {
	rv := objc.Send[LayoutManager](t_.ID, objc.Sel("layoutManager"))
	return rv
}


// Returns the layout manager for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/layoutmanager
func (t_ Typesetter) SetLayoutManager(value objc.IObject /* cross-framework: LayoutManager */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutManager:"), value)
}


// Returns the current line fragment padding, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/linefragmentpadding
func (t_ Typesetter) LineFragmentPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineFragmentPadding"))
	return rv
}


// Returns the current line fragment padding, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/linefragmentpadding
func (t_ Typesetter) SetLineFragmentPadding(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineFragmentPadding:"), value)
}


// Returns the character range currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphcharacterrange
func (t_ Typesetter) ParagraphCharacterRange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("paragraphCharacterRange"))
	return rv
}


// Returns the character range currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphcharacterrange
func (t_ Typesetter) SetParagraphCharacterRange(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphCharacterRange:"), value)
}


// Returns the glyph range currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphglyphrange
func (t_ Typesetter) ParagraphGlyphRange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("paragraphGlyphRange"))
	return rv
}


// Returns the glyph range currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphglyphrange
func (t_ Typesetter) SetParagraphGlyphRange(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphGlyphRange:"), value)
}


// Returns the current paragraph separator character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphseparatorcharacterrange
func (t_ Typesetter) ParagraphSeparatorCharacterRange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("paragraphSeparatorCharacterRange"))
	return rv
}


// Returns the current paragraph separator character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphseparatorcharacterrange
func (t_ Typesetter) SetParagraphSeparatorCharacterRange(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphSeparatorCharacterRange:"), value)
}


// Returns the current paragraph separator range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphseparatorglyphrange
func (t_ Typesetter) ParagraphSeparatorGlyphRange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("paragraphSeparatorGlyphRange"))
	return rv
}


// Returns the current paragraph separator range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphseparatorglyphrange
func (t_ Typesetter) SetParagraphSeparatorGlyphRange(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphSeparatorGlyphRange:"), value)
}


// Returns an array containing the text containers belonging to the current layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/textcontainers
func (t_ Typesetter) TextContainers() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](t_.ID, objc.Sel("textContainers"))
	return rv
}


// Returns an array containing the text containers belonging to the current layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/textcontainers
func (t_ Typesetter) SetTextContainers(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainers:"), value)
}


// Returns the current typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/typesetterbehavior
func (t_ Typesetter) TypesetterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("typesetterBehavior"))
	return rv
}


// Returns the current typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/typesetterbehavior
func (t_ Typesetter) SetTypesetterBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypesetterBehavior:"), value)
}


// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/usesfontleading
func (t_ Typesetter) UsesFontLeading() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontLeading"))
	return rv
}


// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/usesfontleading
func (t_ Typesetter) SetUsesFontLeading(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontLeading:"), value)
}



