// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSXMLDTDNode */


/* debug [class_header]: Header for NSXMLDTDNode */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for XMLDTDNode */
// An interface definition for the [XMLDTDNode] class.
type IXMLDTDNode interface {
	IXMLNode
	
/* debug [class_interface_properties]: Properties for XMLDTDNode */
	// properties:
	DTDKind() XMLDTDNodeKind
	SetDTDKind(value XMLDTDNodeKind)
	External() bool
	NotationName() IString
	SetNotationName(value IString)
	PublicID() IString
	SetPublicID(value IString)
	SystemID() IString
	SetSystemID(value IString)
	IsExternal() bool
	SetIsExternal(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for XMLDTDNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for XMLDTDNode */
// Alloc allocates a new instance without initialization.
func (xc _XMLDTDNodeClass) Alloc() XMLDTDNode {
	rv := objc.Send[XMLDTDNode](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for XMLDTDNode */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for XMLDTDNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/init(kind:options:)
func NewXMLDTDNodeWithKindOptions(kind XMLNodeKind /* not a class type */, options XMLNodeOptions) XMLDTDNode {
	instance := getXMLDTDNodeClass().Alloc()
	rv := objc.Send[XMLDTDNode](instance.ID, objc.Sel("initWithKind:options:"), kind, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewXMLDTDNodeWithKindOptions */


// Returns an object initialized with the DTD declaration in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/init(xmlString:)
func NewXMLDTDNodeWithXMLString(string_ IString) XMLDTDNode {
	instance := getXMLDTDNodeClass().Alloc()
	rv := objc.Send[XMLDTDNode](instance.ID, objc.Sel("initWithXMLString:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewXMLDTDNodeWithXMLString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for XMLDTDNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for XMLDTDNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for XMLDTDNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for XMLDTDNode */

// Returns the receiver’s DTD kind.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/dtdKind-swift.property
func (x_ XMLDTDNode) DTDKind() XMLDTDNodeKind {
	rv := objc.Send[XMLDTDNodeKind](x_.ID, objc.Sel("DTDKind"))
	return rv
}/* debug [instance_properties/getter]: DTDKind */


// Returns the receiver’s DTD kind.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/dtdKind-swift.property
func (x_ XMLDTDNode) SetDTDKind(value XMLDTDNodeKind) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDTDKind:"), value)
}/* debug [instance_properties/setter]: DTDKind */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/isExternal
func (x_ XMLDTDNode) External() bool {
	rv := objc.Send[bool](x_.ID, objc.Sel("external"))
	return rv
}/* debug [instance_properties/getter]: external */


// Returns the name of the notation associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/notationName
func (x_ XMLDTDNode) NotationName() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("notationName"))
	return rv
}/* debug [instance_properties/getter]: notationName */


// Returns the name of the notation associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/notationName
func (x_ XMLDTDNode) SetNotationName(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNotationName:"), value)
}/* debug [instance_properties/setter]: notationName */


// Returns the public identifier associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/publicID
func (x_ XMLDTDNode) PublicID() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("publicID"))
	return rv
}/* debug [instance_properties/getter]: publicID */


// Returns the public identifier associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/publicID
func (x_ XMLDTDNode) SetPublicID(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setPublicID:"), value)
}/* debug [instance_properties/setter]: publicID */


// Returns the system identifier associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/systemID
func (x_ XMLDTDNode) SystemID() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("systemID"))
	return rv
}/* debug [instance_properties/getter]: systemID */


// Returns the system identifier associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/systemID
func (x_ XMLDTDNode) SetSystemID(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setSystemID:"), value)
}/* debug [instance_properties/setter]: systemID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/isexternal
func (x_ XMLDTDNode) IsExternal() bool {
	rv := objc.Send[bool](x_.ID, objc.Sel("isExternal"))
	return rv
}/* debug [instance_properties/getter]: isExternal */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtdnode/isexternal
func (x_ XMLDTDNode) SetIsExternal(value bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setIsExternal:"), value)
}/* debug [instance_properties/setter]: isExternal */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSXMLDTDNode */


