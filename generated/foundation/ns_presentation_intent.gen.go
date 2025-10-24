// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPresentationIntent */


/* debug [class_header]: Header for NSPresentationIntent */
// The class instance for the [PresentationIntent] class.
var (
	PresentationIntentClass     _PresentationIntentClass
	PresentationIntentClassOnce sync.Once
)

func getPresentationIntentClass() _PresentationIntentClass {
	PresentationIntentClassOnce.Do(func() {
		PresentationIntentClass = _PresentationIntentClass{objc.GetClass("NSPresentationIntent")}
	})
	return PresentationIntentClass
}

type _PresentationIntentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PresentationIntent */
// An interface definition for the [PresentationIntent] class.
type IPresentationIntent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PresentationIntent */
	// properties:
	Column() int
	ColumnAlignments() []Number
	ColumnCount() int
	HeaderLevel() int
	Identity() int
	IndentationLevel() int
	IntentKind() PresentationIntentKind
	LanguageHint() IString
	Ordinal() int
	ParentIntent() IPresentationIntent
	Row() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PresentationIntent */
	// methods:
	IsEquivalentToPresentationIntent(other IPresentationIntent) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PresentationIntent */
// Alloc allocates a new instance without initialization.
func (pc _PresentationIntentClass) Alloc() PresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PresentationIntentClass) New() PresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PresentationIntent) Init() PresentationIntent {
	rv := objc.Send[PresentationIntent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PresentationIntent) Autorelease() PresentationIntent {
	rv := objc.Send[PresentationIntent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPresentationIntent creates a new PresentationIntent instance.
func NewPresentationIntent() PresentationIntent {
	return getPresentationIntentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PresentationIntent */
// A type that contains the Markdown formatting for blocks of text, like paragraphs, lists, code blocks, and parts of tables.
//
// An object stores the Markdown semantics for a range of characters in an attributed string. When parsing Markdown into an attributed string, the system sets the value of the attribute to an instance of this class. When displaying your string in system views, the system applies a default visual style to match the corresponding information in this type. To replace the system’s default formatting, remove these attributes from your attributed string and apply the formatting you want.


// A type that contains the Markdown formatting for blocks of text, like paragraphs, lists, code blocks, and parts of tables.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent
type PresentationIntent struct {
	objectivec.Object
}

// PresentationIntentFrom constructs a [PresentationIntent] from an unsafe.Pointer.
//
// A type that contains the Markdown formatting for blocks of text, like paragraphs, lists, code blocks, and parts of tables.
func PresentationIntentFrom(ptr unsafe.Pointer) PresentationIntent {
	return PresentationIntent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PresentationIntent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PresentationIntent */

// Creates a block-quote intent with the provided information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/blockQuoteIntentWithIdentity:nestedInsideIntent:
func (pc _PresentationIntentClass) BlockQuoteIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) IPresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("blockQuoteIntentWithIdentity:nestedInsideIntent:"), identity, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BlockQuoteIntentWithIdentityNestedInsideIntent) */


// Creates an code-block intent with the provided information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/codeBlockIntentWithIdentity:languageHint:nestedInsideIntent:
func (pc _PresentationIntentClass) CodeBlockIntentWithIdentityLanguageHintNestedInsideIntent(identity int, languageHint IString, parent IPresentationIntent) IPresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("codeBlockIntentWithIdentity:languageHint:nestedInsideIntent:"), identity, languageHint, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CodeBlockIntentWithIdentityLanguageHintNestedInsideIntent) */


// Creates a header intent with the provided information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/headerIntentWithIdentity:level:nestedInsideIntent:
func (pc _PresentationIntentClass) HeaderIntentWithIdentityLevelNestedInsideIntent(identity int, level int, parent IPresentationIntent) IPresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("headerIntentWithIdentity:level:nestedInsideIntent:"), identity, level, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HeaderIntentWithIdentityLevelNestedInsideIntent) */


// Creates an item for an ordered list with the provided information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/listItemIntentWithIdentity:ordinal:nestedInsideIntent:
func (pc _PresentationIntentClass) ListItemIntentWithIdentityOrdinalNestedInsideIntent(identity int, ordinal int, parent IPresentationIntent) IPresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("listItemIntentWithIdentity:ordinal:nestedInsideIntent:"), identity, ordinal, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ListItemIntentWithIdentityOrdinalNestedInsideIntent) */


// Creates an ordered-list intent with the provided information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/orderedListIntentWithIdentity:nestedInsideIntent:
func (pc _PresentationIntentClass) OrderedListIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) IPresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("orderedListIntentWithIdentity:nestedInsideIntent:"), identity, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OrderedListIntentWithIdentityNestedInsideIntent) */


// Creates a paragraph intent with the provided information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/paragraphIntentWithIdentity:nestedInsideIntent:
func (pc _PresentationIntentClass) ParagraphIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) IPresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("paragraphIntentWithIdentity:nestedInsideIntent:"), identity, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ParagraphIntentWithIdentityNestedInsideIntent) */


// Creates a table cell intent with the provided information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/tableCellIntentWithIdentity:column:nestedInsideIntent:
func (pc _PresentationIntentClass) TableCellIntentWithIdentityColumnNestedInsideIntent(identity int, column int, parent IPresentationIntent) IPresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("tableCellIntentWithIdentity:column:nestedInsideIntent:"), identity, column, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TableCellIntentWithIdentityColumnNestedInsideIntent) */


// Creates a table header intent with the provided information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/tableHeaderRowIntentWithIdentity:nestedInsideIntent:
func (pc _PresentationIntentClass) TableHeaderRowIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) IPresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("tableHeaderRowIntentWithIdentity:nestedInsideIntent:"), identity, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TableHeaderRowIntentWithIdentityNestedInsideIntent) */


// Creates a table intent with the provided information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/tableIntentWithIdentity:columnCount:alignments:nestedInsideIntent:
func (pc _PresentationIntentClass) TableIntentWithIdentityColumnCountAlignmentsNestedInsideIntent(identity int, columnCount int, alignments []Number, parent IPresentationIntent) IPresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("tableIntentWithIdentity:columnCount:alignments:nestedInsideIntent:"), identity, columnCount, alignments, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TableIntentWithIdentityColumnCountAlignmentsNestedInsideIntent) */


// Creates a table row intent with the provided information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/tableRowIntentWithIdentity:row:nestedInsideIntent:
func (pc _PresentationIntentClass) TableRowIntentWithIdentityRowNestedInsideIntent(identity int, row int, parent IPresentationIntent) IPresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("tableRowIntentWithIdentity:row:nestedInsideIntent:"), identity, row, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TableRowIntentWithIdentityRowNestedInsideIntent) */


// Creates a thematic break intent with the provided information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/thematicBreakIntentWithIdentity:nestedInsideIntent:
func (pc _PresentationIntentClass) ThematicBreakIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) IPresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("thematicBreakIntentWithIdentity:nestedInsideIntent:"), identity, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ThematicBreakIntentWithIdentityNestedInsideIntent) */


// Creates an unordered-list intent with the provided information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/unorderedListIntentWithIdentity:nestedInsideIntent:
func (pc _PresentationIntentClass) UnorderedListIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) IPresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("unorderedListIntentWithIdentity:nestedInsideIntent:"), identity, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnorderedListIntentWithIdentityNestedInsideIntent) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PresentationIntent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PresentationIntent */

// Returns a Boolean value that indicates whether the current intent is equivalent to the specified intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/isEquivalentToPresentationIntent:
func (p_ PresentationIntent) IsEquivalentToPresentationIntent(other IPresentationIntent) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isEquivalentToPresentationIntent:"), other)
	return rv
}/* debug [instance_methods/method]: IsEquivalentToPresentationIntent */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PresentationIntent */

// The column number to which the cell belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/column
func (p_ PresentationIntent) Column() int {
	rv := objc.Send[int](p_.ID, objc.Sel("column"))
	return rv
}/* debug [instance_properties/getter]: column */


// The alignments for the columns in a table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/columnAlignments
func (p_ PresentationIntent) ColumnAlignments() []Number {
	rv := objc.Send[[]Number](p_.ID, objc.Sel("columnAlignments"))
	return rv
}/* debug [instance_properties/getter]: columnAlignments */


// The number of columns in a table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/columnCount
func (p_ PresentationIntent) ColumnCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("columnCount"))
	return rv
}/* debug [instance_properties/getter]: columnCount */


// The level of a header section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/headerLevel
func (p_ PresentationIntent) HeaderLevel() int {
	rv := objc.Send[int](p_.ID, objc.Sel("headerLevel"))
	return rv
}/* debug [instance_properties/getter]: headerLevel */


// A unique identifier for the intent in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/identity
func (p_ PresentationIntent) Identity() int {
	rv := objc.Send[int](p_.ID, objc.Sel("identity"))
	return rv
}/* debug [instance_properties/getter]: identity */


// The indentation level of the intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/indentationLevel
func (p_ PresentationIntent) IndentationLevel() int {
	rv := objc.Send[int](p_.ID, objc.Sel("indentationLevel"))
	return rv
}/* debug [instance_properties/getter]: indentationLevel */


// The type of the intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/intentKind
func (p_ PresentationIntent) IntentKind() PresentationIntentKind {
	rv := objc.Send[PresentationIntentKind](p_.ID, objc.Sel("intentKind"))
	return rv
}/* debug [instance_properties/getter]: intentKind */


// The language associated with the code listing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/languageHint
func (p_ PresentationIntent) LanguageHint() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("languageHint"))
	return rv
}/* debug [instance_properties/getter]: languageHint */


// The number for an item in an ordered list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/ordinal
func (p_ PresentationIntent) Ordinal() int {
	rv := objc.Send[int](p_.ID, objc.Sel("ordinal"))
	return rv
}/* debug [instance_properties/getter]: ordinal */


// The parent of the current intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/parentIntent
func (p_ PresentationIntent) ParentIntent() IPresentationIntent {
	rv := objc.Send[PresentationIntent](p_.ID, objc.Sel("parentIntent"))
	return rv
}/* debug [instance_properties/getter]: parentIntent */


// The row number to which this cell belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/row
func (p_ PresentationIntent) Row() int {
	rv := objc.Send[int](p_.ID, objc.Sel("row"))
	return rv
}/* debug [instance_properties/getter]: row */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPresentationIntent */



