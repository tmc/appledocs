// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AttributedStringMarkdownParsingOptions] class.
var (
	AttributedStringMarkdownParsingOptionsClass     _AttributedStringMarkdownParsingOptionsClass
	AttributedStringMarkdownParsingOptionsClassOnce sync.Once
)

func getAttributedStringMarkdownParsingOptionsClass() _AttributedStringMarkdownParsingOptionsClass {
	AttributedStringMarkdownParsingOptionsClassOnce.Do(func() {
		AttributedStringMarkdownParsingOptionsClass = _AttributedStringMarkdownParsingOptionsClass{objc.GetClass("NSAttributedStringMarkdownParsingOptions")}
	})
	return AttributedStringMarkdownParsingOptionsClass
}

type _AttributedStringMarkdownParsingOptionsClass struct {
	class objc.Class
}

// An interface definition for the [AttributedStringMarkdownParsingOptions] class.
type IAttributedStringMarkdownParsingOptions interface {
	objectivec.IObject
	// properties:
	AllowsExtendedAttributes() bool
	SetAllowsExtendedAttributes(value bool)
	AppliesSourcePositionAttributes() bool
	SetAppliesSourcePositionAttributes(value bool)
	FailurePolicy() AttributedStringMarkdownParsingFailurePolicy
	SetFailurePolicy(value AttributedStringMarkdownParsingFailurePolicy)
	InterpretedSyntax() AttributedStringMarkdownInterpretedSyntax
	SetInterpretedSyntax(value AttributedStringMarkdownInterpretedSyntax)
	LanguageCode() IString
	SetLanguageCode(value IString)
	// methods:
}

// Options that affect the parsing of Markdown content into an attributed string.


// Options that affect the parsing of Markdown content into an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions
type AttributedStringMarkdownParsingOptions struct {
	objectivec.Object
}

// AttributedStringMarkdownParsingOptionsFrom constructs a [AttributedStringMarkdownParsingOptions] from an unsafe.Pointer.
//
// Options that affect the parsing of Markdown content into an attributed string.
func AttributedStringMarkdownParsingOptionsFrom(ptr unsafe.Pointer) AttributedStringMarkdownParsingOptions {
	return AttributedStringMarkdownParsingOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AttributedStringMarkdownParsingOptionsClass) Alloc() AttributedStringMarkdownParsingOptions {
	rv := objc.Send[AttributedStringMarkdownParsingOptions](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AttributedStringMarkdownParsingOptionsClass) New() AttributedStringMarkdownParsingOptions {
	rv := objc.Send[AttributedStringMarkdownParsingOptions](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AttributedStringMarkdownParsingOptions) Init() AttributedStringMarkdownParsingOptions {
	rv := objc.Send[AttributedStringMarkdownParsingOptions](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AttributedStringMarkdownParsingOptions) Autorelease() AttributedStringMarkdownParsingOptions {
	rv := objc.Send[AttributedStringMarkdownParsingOptions](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttributedStringMarkdownParsingOptions creates a new AttributedStringMarkdownParsingOptions instance.
func NewAttributedStringMarkdownParsingOptions() AttributedStringMarkdownParsingOptions {
	return getAttributedStringMarkdownParsingOptionsClass().New()
}




// A Boolean value that indicates whether parsing allows extensions to Markdown that specify extended attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/allowsExtendedAttributes
func (a_ AttributedStringMarkdownParsingOptions) AllowsExtendedAttributes() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsExtendedAttributes"))
	return rv
}


// A Boolean value that indicates whether parsing allows extensions to Markdown that specify extended attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/allowsExtendedAttributes
func (a_ AttributedStringMarkdownParsingOptions) SetAllowsExtendedAttributes(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsExtendedAttributes:"), value)
}


// A Boolean value that indicates whether parsing applies attributes that indicate the position of attributed text in the original Markdown string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/appliesSourcePositionAttributes
func (a_ AttributedStringMarkdownParsingOptions) AppliesSourcePositionAttributes() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("appliesSourcePositionAttributes"))
	return rv
}


// A Boolean value that indicates whether parsing applies attributes that indicate the position of attributed text in the original Markdown string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/appliesSourcePositionAttributes
func (a_ AttributedStringMarkdownParsingOptions) SetAppliesSourcePositionAttributes(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAppliesSourcePositionAttributes:"), value)
}


// The policy for handling a parsing failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/failurePolicy
func (a_ AttributedStringMarkdownParsingOptions) FailurePolicy() AttributedStringMarkdownParsingFailurePolicy {
	rv := objc.Send[AttributedStringMarkdownParsingFailurePolicy](a_.ID, objc.Sel("failurePolicy"))
	return rv
}


// The policy for handling a parsing failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/failurePolicy
func (a_ AttributedStringMarkdownParsingOptions) SetFailurePolicy(value AttributedStringMarkdownParsingFailurePolicy) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFailurePolicy:"), value)
}


// The syntax for intepreting a Markdown string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/interpretedSyntax
func (a_ AttributedStringMarkdownParsingOptions) InterpretedSyntax() AttributedStringMarkdownInterpretedSyntax {
	rv := objc.Send[AttributedStringMarkdownInterpretedSyntax](a_.ID, objc.Sel("interpretedSyntax"))
	return rv
}


// The syntax for intepreting a Markdown string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/interpretedSyntax
func (a_ AttributedStringMarkdownParsingOptions) SetInterpretedSyntax(value AttributedStringMarkdownInterpretedSyntax) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInterpretedSyntax:"), value)
}


// The BCP-47 language code for this document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/languageCode
func (a_ AttributedStringMarkdownParsingOptions) LanguageCode() IString {
	rv := objc.Send[String](a_.ID, objc.Sel("languageCode"))
	return rv
}


// The BCP-47 language code for this document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/languageCode
func (a_ AttributedStringMarkdownParsingOptions) SetLanguageCode(value IString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLanguageCode:"), value)
}


