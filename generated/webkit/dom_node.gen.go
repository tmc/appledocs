// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMNode */

/* debug [class_header]: Header for DOMNode */
// The class instance for the [DOMNode] class.
var (
	DOMNodeClass     _DOMNodeClass
	DOMNodeClassOnce sync.Once
)

func getDOMNodeClass() _DOMNodeClass {
	DOMNodeClassOnce.Do(func() {
		DOMNodeClass = _DOMNodeClass{objc.GetClass("DOMNode")}
	})
	return DOMNodeClass
}

type _DOMNodeClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMNode */
// An interface definition for the [DOMNode] class.
type IDOMNode interface {
	IDOMObject

	/* debug [class_interface_properties]: Properties for DOMNode */
	// properties:
	Attributes() IDOMNamedNodeMap
	BaseURI() objc.IObject /* cross-framework: NSString */
	ChildNodes() IDOMNodeList
	FirstChild() IDOMNode
	IsContentEditable() bool
	LastChild() IDOMNode
	LocalName() objc.IObject    /* cross-framework: NSString */
	NamespaceURI() objc.IObject /* cross-framework: NSString */
	NextSibling() IDOMNode
	NodeName() objc.IObject /* cross-framework: NSString */
	NodeType() unsafe.Pointer
	NodeValue() objc.IObject /* cross-framework: NSString */
	SetNodeValue(value objc.IObject /* cross-framework: NSString */)
	OwnerDocument() IDOMDocument
	ParentNode() IDOMNode
	ParentElement() IDOMElement
	Prefix() objc.IObject /* cross-framework: NSString */
	SetPrefix(value objc.IObject /* cross-framework: NSString */)
	PreviousSibling() IDOMNode
	TextContent() objc.IObject /* cross-framework: NSString */
	SetTextContent(value objc.IObject /* cross-framework: NSString */)
	WebArchive() IWebArchive
	Parent() IDOMNode
	SetParent(value IDOMNode)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMNode */
	// methods:
	BoundingBox() objc.IObject /* cross-framework: Rect */
	LineBoxRects() foundation.Array
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMNode */
// Alloc allocates a new instance without initialization.
func (dc _DOMNodeClass) Alloc() DOMNode {
	rv := objc.Send[DOMNode](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMNodeClass) New() DOMNode {
	rv := objc.Send[DOMNode](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMNode) Init() DOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMNode) Autorelease() DOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMNode creates a new DOMNode instance.
func NewDOMNode() DOMNode {
	return getDOMNodeClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode
type DOMNode struct {
	DOMObject
}

// DOMNodeFrom constructs a [DOMNode] from an unsafe.Pointer.
func DOMNodeFrom(ptr unsafe.Pointer) DOMNode {
	return DOMNode{
		DOMObject: DOMObjectFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMNode */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMNode */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMNode */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMNode */

// Returns a rectangle that bounds the onscreen rendering of the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/boundingBox()
func (d_ DOMNode) BoundingBox() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](d_.ID, objc.Sel("boundingBox"))
	return rv
} /* debug [instance_methods/method]: BoundingBox */

// Returns the rectangles that bound each line of text in the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/lineBoxRects()
func (d_ DOMNode) LineBoxRects() foundation.Array {
	rv := objc.Send[foundation.Array](d_.ID, objc.Sel("lineBoxRects"))
	return rv
} /* debug [instance_methods/method]: LineBoxRects */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/attributes
func (d_ DOMNode) Attributes() IDOMNamedNodeMap {
	rv := objc.Send[DOMNamedNodeMap](d_.ID, objc.Sel("attributes"))
	return rv
} /* debug [instance_properties/getter]: attributes */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/baseURI
func (d_ DOMNode) BaseURI() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("baseURI"))
	return rv
} /* debug [instance_properties/getter]: baseURI */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/childNodes
func (d_ DOMNode) ChildNodes() IDOMNodeList {
	rv := objc.Send[DOMNodeList](d_.ID, objc.Sel("childNodes"))
	return rv
} /* debug [instance_properties/getter]: childNodes */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/firstChild
func (d_ DOMNode) FirstChild() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("firstChild"))
	return rv
} /* debug [instance_properties/getter]: firstChild */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/isContentEditable
func (d_ DOMNode) IsContentEditable() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isContentEditable"))
	return rv
} /* debug [instance_properties/getter]: isContentEditable */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/lastChild
func (d_ DOMNode) LastChild() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("lastChild"))
	return rv
} /* debug [instance_properties/getter]: lastChild */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/localName
func (d_ DOMNode) LocalName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("localName"))
	return rv
} /* debug [instance_properties/getter]: localName */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/namespaceURI
func (d_ DOMNode) NamespaceURI() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("namespaceURI"))
	return rv
} /* debug [instance_properties/getter]: namespaceURI */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/nextSibling
func (d_ DOMNode) NextSibling() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("nextSibling"))
	return rv
} /* debug [instance_properties/getter]: nextSibling */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/nodeName
func (d_ DOMNode) NodeName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("nodeName"))
	return rv
} /* debug [instance_properties/getter]: nodeName */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/nodeType
func (d_ DOMNode) NodeType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("nodeType"))
	return rv
} /* debug [instance_properties/getter]: nodeType */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/nodeValue
func (d_ DOMNode) NodeValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("nodeValue"))
	return rv
} /* debug [instance_properties/getter]: nodeValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/nodeValue
func (d_ DOMNode) SetNodeValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNodeValue:"), value)
} /* debug [instance_properties/setter]: nodeValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/ownerDocument
func (d_ DOMNode) OwnerDocument() IDOMDocument {
	rv := objc.Send[DOMDocument](d_.ID, objc.Sel("ownerDocument"))
	return rv
} /* debug [instance_properties/getter]: ownerDocument */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/parent
func (d_ DOMNode) ParentNode() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("parentNode"))
	return rv
} /* debug [instance_properties/getter]: parentNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/parentElement
func (d_ DOMNode) ParentElement() IDOMElement {
	rv := objc.Send[DOMElement](d_.ID, objc.Sel("parentElement"))
	return rv
} /* debug [instance_properties/getter]: parentElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/prefix
func (d_ DOMNode) Prefix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("prefix"))
	return rv
} /* debug [instance_properties/getter]: prefix */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/prefix
func (d_ DOMNode) SetPrefix(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPrefix:"), value)
} /* debug [instance_properties/setter]: prefix */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/previousSibling
func (d_ DOMNode) PreviousSibling() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("previousSibling"))
	return rv
} /* debug [instance_properties/getter]: previousSibling */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/textContent
func (d_ DOMNode) TextContent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("textContent"))
	return rv
} /* debug [instance_properties/getter]: textContent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/textContent
func (d_ DOMNode) SetTextContent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTextContent:"), value)
} /* debug [instance_properties/setter]: textContent */

// A web archive of the content of the node and its children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNode/webArchive
func (d_ DOMNode) WebArchive() IWebArchive {
	rv := objc.Send[WebArchive](d_.ID, objc.Sel("webArchive"))
	return rv
} /* debug [instance_properties/getter]: webArchive */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/domnode/parent
func (d_ DOMNode) Parent() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("parent"))
	return rv
} /* debug [instance_properties/getter]: parent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/domnode/parent
func (d_ DOMNode) SetParent(value IDOMNode) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setParent:"), value)
} /* debug [instance_properties/setter]: parent */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMNode */
