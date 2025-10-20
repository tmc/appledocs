// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AttributeDescription] class.
var (
	attributeDescriptionClass     _AttributeDescriptionClass
	attributeDescriptionClassOnce sync.Once
)

func getAttributeDescriptionClass() _AttributeDescriptionClass {
	attributeDescriptionClassOnce.Do(func() {
		attributeDescriptionClass = _AttributeDescriptionClass{objc.GetClass("NSAttributeDescription")}
	})
	return attributeDescriptionClass
}

type _AttributeDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [AttributeDescription] class.
type IAttributeDescription interface {
	IPropertyDescription
}

// A description of a single attribute belonging to an entity.
//
// inherits from , which provides most of the basic behavior. Instances of are used to describe attributes, as distinct from relationships. The class adds the ability to specify the attribute type, and to specify a default value. In a managed object model, you must specify the type of all attributes—you can only use the undefined attribute type ( ) for transient attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription
type AttributeDescription struct {
	PropertyDescription
}

// AttributeDescriptionFrom constructs a [AttributeDescription] from an unsafe.Pointer.
//
// A description of a single attribute belonging to an entity.
func AttributeDescriptionFrom(ptr unsafe.Pointer) AttributeDescription {
	return AttributeDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AttributeDescriptionClass) Alloc() AttributeDescription {
	rv := objc.Send[AttributeDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AttributeDescriptionClass) New() AttributeDescription {
	rv := objc.Send[AttributeDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AttributeDescription) Init() AttributeDescription {
	rv := objc.Send[AttributeDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AttributeDescription) Autorelease() AttributeDescription {
	rv := objc.Send[AttributeDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttributeDescription creates a new AttributeDescription instance.
func NewAttributeDescription() AttributeDescription {
	return getAttributeDescriptionClass().New()
}


// A Boolean value that determines whether to encrypt the attribute’s value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/allowsCloudEncryption
func (a_ AttributeDescription) AllowsCloudEncryption() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsCloudEncryption"))
	return rv
}

// SetAllowsCloudEncryption sets the value of the allowsCloudEncryption property.
// A Boolean value that determines whether to encrypt the attribute’s value.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/allowsCloudEncryption
func (a_ AttributeDescription) SetAllowsCloudEncryption(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsCloudEncryption:"), value)
}
// The attribute’s type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/attributeType-swift.property
func (a_ AttributeDescription) AttributeType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributeType"))
	return rv
}

// SetAttributeType sets the value of the attributeType property.
// The attribute’s type.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/attributeType-swift.property
func (a_ AttributeDescription) SetAttributeType(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributeType:"), value)
}


