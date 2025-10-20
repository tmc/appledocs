// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XMLDTDNode] class.
var (
	XMLDTDNodeClass     _XMLDTDNodeClass
	XMLDTDNodeClassOnce sync.Once
)

func getXMLDTDNodeClass() _XMLDTDNodeClass {
	XMLDTDNodeClassOnce.Do(func() {
		XMLDTDNodeClass = _XMLDTDNodeClass{objc.GetClass("NSXMLDTDNode")}
	})
	return XMLDTDNodeClass
}

type _XMLDTDNodeClass struct {
	class objc.Class
}

// An interface definition for the [XMLDTDNode] class.
type IXMLDTDNode interface {
	IXMLNode
}

// A representation of element, attribute-list, entity, and notation declarations in a Document Type Definition.
//
// objects are the sole children of a object (possibly along with comment nodes and processing-instruction nodes). They themselves cannot have any children. objects can be of four kinds—element, attribute-list, entity, or notation declaration—and can also be of a subkind, as specified by a constant. For example, a DTD entity-declaration node could represent an unparsed entity declaration ( ) rather than a parameter entity declaration ( ). You can use a DTD node’s subkind to help determine how to handle the value of the node. You can create an object with the method, the class method , or with the initializer (in the latter method supplying the appropriate constant). Setting the object value or string value of an objects affects different parts of different kinds of declaration. See the related programming topic for more information.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode
type XMLDTDNode struct {
	XMLNode
}

// XMLDTDNodeFrom constructs a [XMLDTDNode] from an unsafe.Pointer.
//
// A representation of element, attribute-list, entity, and notation declarations in a Document Type Definition.
func XMLDTDNodeFrom(ptr unsafe.Pointer) XMLDTDNode {
	return XMLDTDNode{
		XMLNode: XMLNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (xc _XMLDTDNodeClass) Alloc() XMLDTDNode {
	rv := objc.Send[XMLDTDNode](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (xc _XMLDTDNodeClass) New() XMLDTDNode {
	rv := objc.Send[XMLDTDNode](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XMLDTDNode) Init() XMLDTDNode {
	rv := objc.Send[XMLDTDNode](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XMLDTDNode) Autorelease() XMLDTDNode {
	rv := objc.Send[XMLDTDNode](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLDTDNode creates a new XMLDTDNode instance.
func NewXMLDTDNode() XMLDTDNode {
	return getXMLDTDNodeClass().New()
}


// Returns an object initialized with the DTD declaration in a given string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/init(xmlString:)
func NewXMLDTDNodeWithXMLString(string string) XMLDTDNode {
	instance := getXMLDTDNodeClass().Alloc()
	rv := objc.Send[XMLDTDNode](instance.ID, objc.Sel("initWithXMLString:"), objc.String(string))
	rv.Autorelease()
	return rv
}



