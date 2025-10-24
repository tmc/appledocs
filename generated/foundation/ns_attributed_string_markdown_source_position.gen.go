// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAttributedStringMarkdownSourcePosition */


/* debug [class_header]: Header for NSAttributedStringMarkdownSourcePosition */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AttributedStringMarkdownSourcePosition */
// An interface definition for the [AttributedStringMarkdownSourcePosition] class.
type IAttributedStringMarkdownSourcePosition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AttributedStringMarkdownSourcePosition */
	// properties:
	EndColumn() int
	EndLine() int
	StartColumn() int
	StartLine() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AttributedStringMarkdownSourcePosition */
	// methods:
	RangeInString(string_ IString) objc.IObject /* cross-framework: Range */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AttributedStringMarkdownSourcePosition */
// Alloc allocates a new instance without initialization.
func (ac _AttributedStringMarkdownSourcePositionClass) Alloc() AttributedStringMarkdownSourcePosition {
	rv := objc.Send[AttributedStringMarkdownSourcePosition](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AttributedStringMarkdownSourcePosition */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AttributedStringMarkdownSourcePosition */

// Creates a Markdown source position instance from its start and end line and column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/initWithStartLine:startColumn:endLine:endColumn:
func NewAttributedStringMarkdownSourcePositionWithStartLineStartColumnEndLineEndColumn(startLine int, startColumn int, endLine int, endColumn int) AttributedStringMarkdownSourcePosition {
	instance := getAttributedStringMarkdownSourcePositionClass().Alloc()
	rv := objc.Send[AttributedStringMarkdownSourcePosition](instance.ID, objc.Sel("initWithStartLine:startColumn:endLine:endColumn:"), startLine, startColumn, endLine, endColumn)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAttributedStringMarkdownSourcePositionWithStartLineStartColumnEndLineEndColumn */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AttributedStringMarkdownSourcePosition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AttributedStringMarkdownSourcePosition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AttributedStringMarkdownSourcePosition */

// Returns a range indicating the source portion within a Markdown string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/rangeInString:
func (a_ AttributedStringMarkdownSourcePosition) RangeInString(string_ IString) objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("rangeInString:"), string_)
	return rv
}/* debug [instance_methods/method]: RangeInString */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AttributedStringMarkdownSourcePosition */

// The column where the text ends in the Markdown source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/endColumn
func (a_ AttributedStringMarkdownSourcePosition) EndColumn() int {
	rv := objc.Send[int](a_.ID, objc.Sel("endColumn"))
	return rv
}/* debug [instance_properties/getter]: endColumn */


// The line where the text ends in the Markdown source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/endLine
func (a_ AttributedStringMarkdownSourcePosition) EndLine() int {
	rv := objc.Send[int](a_.ID, objc.Sel("endLine"))
	return rv
}/* debug [instance_properties/getter]: endLine */


// The column where the text begins in the Markdown source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/startColumn
func (a_ AttributedStringMarkdownSourcePosition) StartColumn() int {
	rv := objc.Send[int](a_.ID, objc.Sel("startColumn"))
	return rv
}/* debug [instance_properties/getter]: startColumn */


// The line where the text begins in the Markdown source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/startLine
func (a_ AttributedStringMarkdownSourcePosition) StartLine() int {
	rv := objc.Send[int](a_.ID, objc.Sel("startLine"))
	return rv
}/* debug [instance_properties/getter]: startLine */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAttributedStringMarkdownSourcePosition */


