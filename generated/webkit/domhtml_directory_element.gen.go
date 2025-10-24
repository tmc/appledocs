// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLDirectoryElement */

/* debug [class_header]: Header for DOMHTMLDirectoryElement */
// The class instance for the [DOMHTMLDirectoryElement] class.
var (
	DOMHTMLDirectoryElementClass     _DOMHTMLDirectoryElementClass
	DOMHTMLDirectoryElementClassOnce sync.Once
)

func getDOMHTMLDirectoryElementClass() _DOMHTMLDirectoryElementClass {
	DOMHTMLDirectoryElementClassOnce.Do(func() {
		DOMHTMLDirectoryElementClass = _DOMHTMLDirectoryElementClass{objc.GetClass("DOMHTMLDirectoryElement")}
	})
	return DOMHTMLDirectoryElementClass
}

type _DOMHTMLDirectoryElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLDirectoryElement */
// An interface definition for the [DOMHTMLDirectoryElement] class.
type IDOMHTMLDirectoryElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLDirectoryElement */
	// properties:
	Compact() bool
	SetCompact(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLDirectoryElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLDirectoryElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLDirectoryElementClass) Alloc() DOMHTMLDirectoryElement {
	rv := objc.Send[DOMHTMLDirectoryElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLDirectoryElementClass) New() DOMHTMLDirectoryElement {
	rv := objc.Send[DOMHTMLDirectoryElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLDirectoryElement) Init() DOMHTMLDirectoryElement {
	rv := objc.Send[DOMHTMLDirectoryElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLDirectoryElement) Autorelease() DOMHTMLDirectoryElement {
	rv := objc.Send[DOMHTMLDirectoryElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLDirectoryElement creates a new DOMHTMLDirectoryElement instance.
func NewDOMHTMLDirectoryElement() DOMHTMLDirectoryElement {
	return getDOMHTMLDirectoryElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLDirectoryElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDirectoryElement
type DOMHTMLDirectoryElement struct {
	DOMHTMLElement
}

// DOMHTMLDirectoryElementFrom constructs a [DOMHTMLDirectoryElement] from an unsafe.Pointer.
func DOMHTMLDirectoryElementFrom(ptr unsafe.Pointer) DOMHTMLDirectoryElement {
	return DOMHTMLDirectoryElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLDirectoryElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLDirectoryElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLDirectoryElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLDirectoryElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLDirectoryElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDirectoryElement/compact
func (d_ DOMHTMLDirectoryElement) Compact() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("compact"))
	return rv
} /* debug [instance_properties/getter]: compact */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDirectoryElement/compact
func (d_ DOMHTMLDirectoryElement) SetCompact(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCompact:"), value)
} /* debug [instance_properties/setter]: compact */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLDirectoryElement */
