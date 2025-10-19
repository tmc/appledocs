// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XMLDTDNode] class.
var (
	xMLDTDNodeClass     _XMLDTDNodeClass
	xMLDTDNodeClassOnce sync.Once
)

func getXMLDTDNodeClass() _XMLDTDNodeClass {
	xMLDTDNodeClassOnce.Do(func() {
		xMLDTDNodeClass = _XMLDTDNodeClass{objc.GetClass("NSXMLDTDNode")}
	})
	return xMLDTDNodeClass
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




