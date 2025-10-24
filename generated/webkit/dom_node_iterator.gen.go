// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMNodeIterator */

/* debug [class_header]: Header for DOMNodeIterator */
// The class instance for the [DOMNodeIterator] class.
var (
	DOMNodeIteratorClass     _DOMNodeIteratorClass
	DOMNodeIteratorClassOnce sync.Once
)

func getDOMNodeIteratorClass() _DOMNodeIteratorClass {
	DOMNodeIteratorClassOnce.Do(func() {
		DOMNodeIteratorClass = _DOMNodeIteratorClass{objc.GetClass("DOMNodeIterator")}
	})
	return DOMNodeIteratorClass
}

type _DOMNodeIteratorClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMNodeIterator */
// An interface definition for the [DOMNodeIterator] class.
type IDOMNodeIterator interface {
	IDOMObject

	/* debug [class_interface_properties]: Properties for DOMNodeIterator */
	// properties:
	ExpandEntityReferences() bool
	Filter() unsafe.Pointer
	PointerBeforeReferenceNode() bool
	ReferenceNode() IDOMNode
	Root() IDOMNode
	WhatToShow() unsafe.Pointer
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMNodeIterator */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMNodeIterator */
// Alloc allocates a new instance without initialization.
func (dc _DOMNodeIteratorClass) Alloc() DOMNodeIterator {
	rv := objc.Send[DOMNodeIterator](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMNodeIteratorClass) New() DOMNodeIterator {
	rv := objc.Send[DOMNodeIterator](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMNodeIterator) Init() DOMNodeIterator {
	rv := objc.Send[DOMNodeIterator](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMNodeIterator) Autorelease() DOMNodeIterator {
	rv := objc.Send[DOMNodeIterator](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMNodeIterator creates a new DOMNodeIterator instance.
func NewDOMNodeIterator() DOMNodeIterator {
	return getDOMNodeIteratorClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMNodeIterator */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNodeIterator
type DOMNodeIterator struct {
	DOMObject
}

// DOMNodeIteratorFrom constructs a [DOMNodeIterator] from an unsafe.Pointer.
func DOMNodeIteratorFrom(ptr unsafe.Pointer) DOMNodeIterator {
	return DOMNodeIterator{
		DOMObject: DOMObjectFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMNodeIterator */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMNodeIterator */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMNodeIterator */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMNodeIterator */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMNodeIterator */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNodeIterator/expandEntityReferences
func (d_ DOMNodeIterator) ExpandEntityReferences() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("expandEntityReferences"))
	return rv
} /* debug [instance_properties/getter]: expandEntityReferences */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNodeIterator/filter
func (d_ DOMNodeIterator) Filter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("filter"))
	return rv
} /* debug [instance_properties/getter]: filter */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNodeIterator/pointerBeforeReferenceNode
func (d_ DOMNodeIterator) PointerBeforeReferenceNode() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("pointerBeforeReferenceNode"))
	return rv
} /* debug [instance_properties/getter]: pointerBeforeReferenceNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNodeIterator/referenceNode
func (d_ DOMNodeIterator) ReferenceNode() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("referenceNode"))
	return rv
} /* debug [instance_properties/getter]: referenceNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNodeIterator/root
func (d_ DOMNodeIterator) Root() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("root"))
	return rv
} /* debug [instance_properties/getter]: root */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNodeIterator/whatToShow
func (d_ DOMNodeIterator) WhatToShow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("whatToShow"))
	return rv
} /* debug [instance_properties/getter]: whatToShow */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMNodeIterator */
