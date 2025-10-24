// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMRange */

/* debug [class_header]: Header for DOMRange */
// The class instance for the [DOMRange] class.
var (
	DOMRangeClass     _DOMRangeClass
	DOMRangeClassOnce sync.Once
)

func getDOMRangeClass() _DOMRangeClass {
	DOMRangeClassOnce.Do(func() {
		DOMRangeClass = _DOMRangeClass{objc.GetClass("DOMRange")}
	})
	return DOMRangeClass
}

type _DOMRangeClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMRange */
// An interface definition for the [DOMRange] class.
type IDOMRange interface {
	IDOMObject

	/* debug [class_interface_properties]: Properties for DOMRange */
	// properties:
	Collapsed() bool
	CommonAncestorContainer() IDOMNode
	EndContainer() IDOMNode
	EndOffset() int
	MarkupString() objc.IObject /* cross-framework: NSString */
	StartContainer() IDOMNode
	StartOffset() int
	Text() objc.IObject /* cross-framework: NSString */
	WebArchive() IWebArchive
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMRange */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMRange */
// Alloc allocates a new instance without initialization.
func (dc _DOMRangeClass) Alloc() DOMRange {
	rv := objc.Send[DOMRange](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMRangeClass) New() DOMRange {
	rv := objc.Send[DOMRange](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMRange) Init() DOMRange {
	rv := objc.Send[DOMRange](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMRange) Autorelease() DOMRange {
	rv := objc.Send[DOMRange](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMRange creates a new DOMRange instance.
func NewDOMRange() DOMRange {
	return getDOMRangeClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMRange */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRange
type DOMRange struct {
	DOMObject
}

// DOMRangeFrom constructs a [DOMRange] from an unsafe.Pointer.
func DOMRangeFrom(ptr unsafe.Pointer) DOMRange {
	return DOMRange{
		DOMObject: DOMObjectFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMRange */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMRange */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMRange */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMRange */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMRange */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRange/collapsed
func (d_ DOMRange) Collapsed() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("collapsed"))
	return rv
} /* debug [instance_properties/getter]: collapsed */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRange/commonAncestorContainer
func (d_ DOMRange) CommonAncestorContainer() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("commonAncestorContainer"))
	return rv
} /* debug [instance_properties/getter]: commonAncestorContainer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRange/endContainer
func (d_ DOMRange) EndContainer() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("endContainer"))
	return rv
} /* debug [instance_properties/getter]: endContainer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRange/endOffset
func (d_ DOMRange) EndOffset() int {
	rv := objc.Send[int](d_.ID, objc.Sel("endOffset"))
	return rv
} /* debug [instance_properties/getter]: endOffset */

// A string in markup format corresponding to the content in the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRange/markupString
func (d_ DOMRange) MarkupString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("markupString"))
	return rv
} /* debug [instance_properties/getter]: markupString */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRange/startContainer
func (d_ DOMRange) StartContainer() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("startContainer"))
	return rv
} /* debug [instance_properties/getter]: startContainer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRange/startOffset
func (d_ DOMRange) StartOffset() int {
	rv := objc.Send[int](d_.ID, objc.Sel("startOffset"))
	return rv
} /* debug [instance_properties/getter]: startOffset */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRange/text
func (d_ DOMRange) Text() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("text"))
	return rv
} /* debug [instance_properties/getter]: text */

// A web archive of the content in the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRange/webArchive
func (d_ DOMRange) WebArchive() IWebArchive {
	rv := objc.Send[WebArchive](d_.ID, objc.Sel("webArchive"))
	return rv
} /* debug [instance_properties/getter]: webArchive */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMRange */
