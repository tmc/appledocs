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
	// properties:
	DTDKind() XMLDTDNodeKind
	SetDTDKind(value XMLDTDNodeKind)
	External() bool /* primitive/slice/pointer */
	NotationName() string /* primitive/slice/pointer */
	SetNotationName(value string /* primitive/slice/pointer */)
	PublicID() string /* primitive/slice/pointer */
	SetPublicID(value string /* primitive/slice/pointer */)
	SystemID() string /* primitive/slice/pointer */
	SetSystemID(value string /* primitive/slice/pointer */)
	IsExternal() bool /* primitive/slice/pointer */
	SetIsExternal(value bool /* primitive/slice/pointer */)
	// methods:
}

// A representation of element, attribute-list, entity, and notation declarations in a Document Type Definition.
//
// objects are the sole children of a object (possibly along with comment nodes and processing-instruction nodes). They themselves cannot have any children. objects can be of four kinds—element, attribute-list, entity, or notation declaration—and can also be of a subkind, as specified by a constant. For example, a DTD entity-declaration node could represent an unparsed entity declaration ( ) rather than a parameter entity declaration ( ). You can use a DTD node’s subkind to help determine how to handle the value of the node. You can create an object with the method, the class method , or with the initializer (in the latter method supplying the appropriate constant). Setting the object value or string value of an objects affects different parts of different kinds of declaration. See the related programming topic for more information.


// A representation of element, attribute-list, entity, and notation declarations in a Document Type Definition.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/init(kind:options:)
func NewXMLDTDNodeWithKindOptions(kind XMLNodeKind, options XMLNodeOptions) XMLDTDNode {
	instance := getXMLDTDNodeClass().Alloc()
	rv := objc.Send[XMLDTDNode](instance.ID, objc.Sel("initWithKind:options:"), kind, options)
	rv.Autorelease()
	return rv
}


// Returns an object initialized with the DTD declaration in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/init(xmlString:)
func NewXMLDTDNodeWithXMLString(string_ string /* primitive/slice/pointer */) XMLDTDNode {
	instance := getXMLDTDNodeClass().Alloc()
	rv := objc.Send[XMLDTDNode](instance.ID, objc.Sel("initWithXMLString:"), objc.String(string_))
	rv.Autorelease()
	return rv
}



// Returns the receiver’s DTD kind.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/dtdKind-swift.property
func (x_ XMLDTDNode) DTDKind() XMLDTDNodeKind {
	rv := objc.Send[XMLDTDNodeKind](x_.ID, objc.Sel("DTDKind"))
	return rv
}


// Returns the receiver’s DTD kind.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/dtdKind-swift.property
func (x_ XMLDTDNode) SetDTDKind(value XMLDTDNodeKind) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDTDKind:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/isExternal
func (x_ XMLDTDNode) External() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](x_.ID, objc.Sel("external"))
	return rv
}


// Returns the name of the notation associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/notationName
func (x_ XMLDTDNode) NotationName() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](x_.ID, objc.Sel("notationName"))
	return rv
}


// Returns the name of the notation associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/notationName
func (x_ XMLDTDNode) SetNotationName(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNotationName:"), objc.String(value))
}


// Returns the public identifier associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/publicID
func (x_ XMLDTDNode) PublicID() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](x_.ID, objc.Sel("publicID"))
	return rv
}


// Returns the public identifier associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/publicID
func (x_ XMLDTDNode) SetPublicID(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setPublicID:"), objc.String(value))
}


// Returns the system identifier associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/systemID
func (x_ XMLDTDNode) SystemID() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](x_.ID, objc.Sel("systemID"))
	return rv
}


// Returns the system identifier associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/systemID
func (x_ XMLDTDNode) SetSystemID(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setSystemID:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/isexternal
func (x_ XMLDTDNode) IsExternal() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](x_.ID, objc.Sel("isExternal"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/isexternal
func (x_ XMLDTDNode) SetIsExternal(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setIsExternal:"), value)
}


