// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ParagraphStyle] class.
var (
	ParagraphStyleClass     _ParagraphStyleClass
	ParagraphStyleClassOnce sync.Once
)

func getParagraphStyleClass() _ParagraphStyleClass {
	ParagraphStyleClassOnce.Do(func() {
		ParagraphStyleClass = _ParagraphStyleClass{objc.GetClass("NSParagraphStyle")}
	})
	return ParagraphStyleClass
}

type _ParagraphStyleClass struct {
	class objc.Class
}

// An interface definition for the [ParagraphStyle] class.
type IParagraphStyle interface {
	objectivec.IObject
}

// The paragraph or ruler attributes for an attributed string.
//
// An object stores formatting information for a paragraph of text. The formatting information includes the amount of space between lines, indentations for lines of text, line heights, tab-stop positions, and more. Apply paragraph styles to the text of an attributed string by adding the attribute and setting its value to an instance of this class. The text-rendering system uses the paragraph style information in an attributed string to lay out and render the text. The class manages an immutable set of style information, but you can create an when you want to modify the style information before applying it to your text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle
type ParagraphStyle struct {
	objectivec.Object
}

// ParagraphStyleFrom constructs a [ParagraphStyle] from an unsafe.Pointer.
//
// The paragraph or ruler attributes for an attributed string.
func ParagraphStyleFrom(ptr unsafe.Pointer) ParagraphStyle {
	return ParagraphStyle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _ParagraphStyleClass) Alloc() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ParagraphStyleClass) New() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ParagraphStyle) Init() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ParagraphStyle) Autorelease() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewParagraphStyle creates a new ParagraphStyle instance.
func NewParagraphStyle() ParagraphStyle {
	return getParagraphStyleClass().New()
}


// A Boolean value that indicates whether the system tightens character spacing before truncating text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/allowsDefaultTighteningForTruncation
func (p_ ParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}

// The paragraph’s threshold for hyphenation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/hyphenationFactor
func (p_ ParagraphStyle) HyphenationFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("hyphenationFactor"))
	return rv
}

// The mode for breaking lines in the paragraph that don’t fit within a container.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/lineBreakMode
func (p_ ParagraphStyle) LineBreakMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("lineBreakMode"))
	return rv
}

// The strategy for breaking lines while laying out paragraphs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/lineBreakStrategy-swift.property
func (p_ ParagraphStyle) LineBreakStrategy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("lineBreakStrategy"))
	return rv
}

// The threshold for using tightening as an alternative to truncation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/tighteningFactorForTruncation
func (p_ ParagraphStyle) TighteningFactorForTruncation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("tighteningFactorForTruncation"))
	return rv
}

// A Boolean value that indicates whether the paragraph style uses the system hyphenation settings.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/usesDefaultHyphenation
func (p_ ParagraphStyle) UsesDefaultHyphenation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesDefaultHyphenation"))
	return rv
}



