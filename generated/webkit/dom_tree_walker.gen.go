// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMTreeWalker */


/* debug [class_header]: Header for DOMTreeWalker */
// The class instance for the [DOMTreeWalker] class.
var (
	DOMTreeWalkerClass     _DOMTreeWalkerClass
	DOMTreeWalkerClassOnce sync.Once
)

func getDOMTreeWalkerClass() _DOMTreeWalkerClass {
	DOMTreeWalkerClassOnce.Do(func() {
		DOMTreeWalkerClass = _DOMTreeWalkerClass{objc.GetClass("DOMTreeWalker")}
	})
	return DOMTreeWalkerClass
}

type _DOMTreeWalkerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMTreeWalker */
// An interface definition for the [DOMTreeWalker] class.
type IDOMTreeWalker interface {
	IDOMObject
	
/* debug [class_interface_properties]: Properties for DOMTreeWalker */
	// properties:
	CurrentNode() IDOMNode
	SetCurrentNode(value IDOMNode)
	ExpandEntityReferences() bool
	Filter() unsafe.Pointer
	Root() IDOMNode
	WhatToShow() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMTreeWalker */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMTreeWalker */
// Alloc allocates a new instance without initialization.
func (dc _DOMTreeWalkerClass) Alloc() DOMTreeWalker {
	rv := objc.Send[DOMTreeWalker](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMTreeWalkerClass) New() DOMTreeWalker {
	rv := objc.Send[DOMTreeWalker](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMTreeWalker) Init() DOMTreeWalker {
	rv := objc.Send[DOMTreeWalker](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMTreeWalker) Autorelease() DOMTreeWalker {
	rv := objc.Send[DOMTreeWalker](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMTreeWalker creates a new DOMTreeWalker instance.
func NewDOMTreeWalker() DOMTreeWalker {
	return getDOMTreeWalkerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMTreeWalker */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMTreeWalker
type DOMTreeWalker struct {
	DOMObject
}

// DOMTreeWalkerFrom constructs a [DOMTreeWalker] from an unsafe.Pointer.
func DOMTreeWalkerFrom(ptr unsafe.Pointer) DOMTreeWalker {
	return DOMTreeWalker{
		DOMObject: DOMObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMTreeWalker *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMTreeWalker */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMTreeWalker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMTreeWalker */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMTreeWalker */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMTreeWalker/currentNode
func (d_ DOMTreeWalker) CurrentNode() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("currentNode"))
	return rv
}/* debug [instance_properties/getter]: currentNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMTreeWalker/currentNode
func (d_ DOMTreeWalker) SetCurrentNode(value IDOMNode) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCurrentNode:"), value)
}/* debug [instance_properties/setter]: currentNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMTreeWalker/expandEntityReferences
func (d_ DOMTreeWalker) ExpandEntityReferences() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("expandEntityReferences"))
	return rv
}/* debug [instance_properties/getter]: expandEntityReferences */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMTreeWalker/filter
func (d_ DOMTreeWalker) Filter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("filter"))
	return rv
}/* debug [instance_properties/getter]: filter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMTreeWalker/root
func (d_ DOMTreeWalker) Root() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("root"))
	return rv
}/* debug [instance_properties/getter]: root */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMTreeWalker/whatToShow
func (d_ DOMTreeWalker) WhatToShow() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("whatToShow"))
	return rv
}/* debug [instance_properties/getter]: whatToShow */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMTreeWalker */



