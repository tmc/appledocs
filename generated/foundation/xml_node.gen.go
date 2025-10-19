// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [XMLNode] class.
var (
	xMLNodeClass     _XMLNodeClass
	xMLNodeClassOnce sync.Once
)

func getXMLNodeClass() _XMLNodeClass {
	xMLNodeClassOnce.Do(func() {
		xMLNodeClass = _XMLNodeClass{objc.GetClass("NSXMLNode")}
	})
	return xMLNodeClass
}

type _XMLNodeClass struct {
	class objc.Class
}

// An interface definition for the [XMLNode] class.
type IXMLNode interface {
	objectivec.IObject
}

// The nodes in the abstract, logical tree structure that represents an XML document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode
type XMLNode struct {
	objectivec.Object
}

// XMLNodeFrom constructs a [XMLNode] from an unsafe.Pointer.
//
// The nodes in the abstract, logical tree structure that represents an XML document.
func XMLNodeFrom(ptr unsafe.Pointer) XMLNode {
	return XMLNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (xc _XMLNodeClass) Alloc() XMLNode {
	rv := objc.Send[XMLNode](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (xc _XMLNodeClass) New() XMLNode {
	rv := objc.Send[XMLNode](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XMLNode) Init() XMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XMLNode) Autorelease() XMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLNode creates a new XMLNode instance.
func NewXMLNode() XMLNode {
	return getXMLNodeClass().New()
}




