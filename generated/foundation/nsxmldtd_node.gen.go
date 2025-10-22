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
	DtdKind() unsafe.Pointer
	SetDtdKind(value unsafe.Pointer)
	IsExternal() bool
	SetIsExternal(value bool)
	NotationName() string
	SetNotationName(value string)
	PublicID() string
	SetPublicID(value string)
	SystemID() string
	SetSystemID(value string)
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





// Returns an object initialized with the DTD declaration in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/init(xmlString:)

func NewXMLDTDNodeWithXMLString(string_ string) XMLDTDNode {
	instance := getXMLDTDNodeClass().Alloc()
	rv := objc.Send[XMLDTDNode](instance.ID, objc.Sel("initWithXMLString:"), objc.String(string_))
	rv.Autorelease()
	return rv
}



// Returns the receiver’s DTD kind.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/dtdkind-swift.property

func (x_ XMLDTDNode) DtdKind() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("dtdKind"))
	return rv
}


// Returns the receiver’s DTD kind.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/dtdkind-swift.property

func (x_ XMLDTDNode) SetDtdKind(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDtdKind:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/isexternal

func (x_ XMLDTDNode) IsExternal() bool {
	rv := objc.Send[bool](x_.ID, objc.Sel("isExternal"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/isexternal

func (x_ XMLDTDNode) SetIsExternal(value bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setIsExternal:"), value)
}


// Returns the name of the notation associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/notationname

func (x_ XMLDTDNode) NotationName() string {
	rv := objc.Send[string](x_.ID, objc.Sel("notationName"))
	return rv
}


// Returns the name of the notation associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/notationname

func (x_ XMLDTDNode) SetNotationName(value string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNotationName:"), objc.String(value))
}


// Returns the public identifier associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/publicid

func (x_ XMLDTDNode) PublicID() string {
	rv := objc.Send[string](x_.ID, objc.Sel("publicID"))
	return rv
}


// Returns the public identifier associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/publicid

func (x_ XMLDTDNode) SetPublicID(value string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setPublicID:"), objc.String(value))
}


// Returns the system identifier associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/systemid

func (x_ XMLDTDNode) SystemID() string {
	rv := objc.Send[string](x_.ID, objc.Sel("systemID"))
	return rv
}


// Returns the system identifier associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/systemid

func (x_ XMLDTDNode) SetSystemID(value string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setSystemID:"), objc.String(value))
}


