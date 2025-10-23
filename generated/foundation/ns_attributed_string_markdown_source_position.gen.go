// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AttributedStringMarkdownSourcePosition] class.
var (
	AttributedStringMarkdownSourcePositionClass     _AttributedStringMarkdownSourcePositionClass
	AttributedStringMarkdownSourcePositionClassOnce sync.Once
)

func getAttributedStringMarkdownSourcePositionClass() _AttributedStringMarkdownSourcePositionClass {
	AttributedStringMarkdownSourcePositionClassOnce.Do(func() {
		AttributedStringMarkdownSourcePositionClass = _AttributedStringMarkdownSourcePositionClass{objc.GetClass("NSAttributedStringMarkdownSourcePosition")}
	})
	return AttributedStringMarkdownSourcePositionClass
}

type _AttributedStringMarkdownSourcePositionClass struct {
	class objc.Class
}

// An interface definition for the [AttributedStringMarkdownSourcePosition] class.
type IAttributedStringMarkdownSourcePosition interface {
	objectivec.IObject
	// properties:
	EndColumn() int /* primitive/slice/pointer */
	EndLine() int /* primitive/slice/pointer */
	StartColumn() int /* primitive/slice/pointer */
	StartLine() int /* primitive/slice/pointer */
	// methods:
	RangeInString(string_ string /* primitive/slice/pointer */) Range /* not a class type */
}

// The position of attributed string text in its original Markdown source string.


// The position of attributed string text in its original Markdown source string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition
type AttributedStringMarkdownSourcePosition struct {
	objectivec.Object
}

// AttributedStringMarkdownSourcePositionFrom constructs a [AttributedStringMarkdownSourcePosition] from an unsafe.Pointer.
//
// The position of attributed string text in its original Markdown source string.
func AttributedStringMarkdownSourcePositionFrom(ptr unsafe.Pointer) AttributedStringMarkdownSourcePosition {
	return AttributedStringMarkdownSourcePosition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AttributedStringMarkdownSourcePositionClass) Alloc() AttributedStringMarkdownSourcePosition {
	rv := objc.Send[AttributedStringMarkdownSourcePosition](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AttributedStringMarkdownSourcePositionClass) New() AttributedStringMarkdownSourcePosition {
	rv := objc.Send[AttributedStringMarkdownSourcePosition](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AttributedStringMarkdownSourcePosition) Init() AttributedStringMarkdownSourcePosition {
	rv := objc.Send[AttributedStringMarkdownSourcePosition](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AttributedStringMarkdownSourcePosition) Autorelease() AttributedStringMarkdownSourcePosition {
	rv := objc.Send[AttributedStringMarkdownSourcePosition](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttributedStringMarkdownSourcePosition creates a new AttributedStringMarkdownSourcePosition instance.
func NewAttributedStringMarkdownSourcePosition() AttributedStringMarkdownSourcePosition {
	return getAttributedStringMarkdownSourcePositionClass().New()
}



// Creates a Markdown source position instance from its start and end line and column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/initWithStartLine:startColumn:endLine:endColumn:
func NewAttributedStringMarkdownSourcePositionWithStartLineStartColumnEndLineEndColumn(startLine int /* primitive/slice/pointer */, startColumn int /* primitive/slice/pointer */, endLine int /* primitive/slice/pointer */, endColumn int /* primitive/slice/pointer */) AttributedStringMarkdownSourcePosition {
	instance := getAttributedStringMarkdownSourcePositionClass().Alloc()
	rv := objc.Send[AttributedStringMarkdownSourcePosition](instance.ID, objc.Sel("initWithStartLine:startColumn:endLine:endColumn:"), startLine, startColumn, endLine, endColumn)
	rv.Autorelease()
	return rv
}



// Returns a range indicating the source portion within a Markdown string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/rangeInString:
func (a_ AttributedStringMarkdownSourcePosition) RangeInString(string_ string /* primitive/slice/pointer */) Range /* not a class type */ {
	rv := objc.Send[Range](a_.ID, objc.Sel("rangeInString:"), objc.String(string_))
	return rv
}


// The column where the text ends in the Markdown source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/endColumn
func (a_ AttributedStringMarkdownSourcePosition) EndColumn() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](a_.ID, objc.Sel("endColumn"))
	return rv
}


// The line where the text ends in the Markdown source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/endLine
func (a_ AttributedStringMarkdownSourcePosition) EndLine() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](a_.ID, objc.Sel("endLine"))
	return rv
}


// The column where the text begins in the Markdown source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/startColumn
func (a_ AttributedStringMarkdownSourcePosition) StartColumn() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](a_.ID, objc.Sel("startColumn"))
	return rv
}


// The line where the text begins in the Markdown source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/startLine
func (a_ AttributedStringMarkdownSourcePosition) StartLine() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](a_.ID, objc.Sel("startLine"))
	return rv
}


