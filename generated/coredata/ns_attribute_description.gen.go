// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AttributeDescription] class.
var (
	AttributeDescriptionClass     _AttributeDescriptionClass
	AttributeDescriptionClassOnce sync.Once
)

func getAttributeDescriptionClass() _AttributeDescriptionClass {
	AttributeDescriptionClassOnce.Do(func() {
		AttributeDescriptionClass = _AttributeDescriptionClass{objc.GetClass("NSAttributeDescription")}
	})
	return AttributeDescriptionClass
}

type _AttributeDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [AttributeDescription] class.
type IAttributeDescription interface {
	IPropertyDescription
	// properties:
	AllowsCloudEncryption() bool
	SetAllowsCloudEncryption(value bool)
	AttributeType() AttributeType
	SetAttributeType(value AttributeType)
	DefaultValue() objc.ID
	SetDefaultValue(value objc.ID)
	PreservesValueInHistoryOnDeletion() bool
	SetPreservesValueInHistoryOnDeletion(value bool)
	AllowsExternalBinaryDataStorage() bool
	SetAllowsExternalBinaryDataStorage(value bool)
	AttributeValueClassName() objc.IObject /* cross-framework: NSString */
	SetAttributeValueClassName(value objc.IObject /* cross-framework: NSString */)
	Type() AttributeType
	SetType(value AttributeType)
	ValueTransformerName() objc.IObject /* cross-framework: NSString */
	SetValueTransformerName(value objc.IObject /* cross-framework: NSString */)
	VersionHash() objc.IObject /* cross-framework: Data */
	SetVersionHash(value objc.IObject /* cross-framework: Data */)
	// methods:
}

// A description of a single attribute belonging to an entity.
//
// inherits from , which provides most of the basic behavior. Instances of are used to describe attributes, as distinct from relationships. The class adds the ability to specify the attribute type, and to specify a default value. In a managed object model, you must specify the type of all attributes—you can only use the undefined attribute type ( ) for transient attributes.


// A description of a single attribute belonging to an entity.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/allowsCloudEncryption
func (a_ AttributeDescription) AllowsCloudEncryption() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsCloudEncryption"))
	return rv
}


// A Boolean value that determines whether to encrypt the attribute’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/allowsCloudEncryption
func (a_ AttributeDescription) SetAllowsCloudEncryption(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsCloudEncryption:"), value)
}


// The attribute’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/attributeType-swift.property
func (a_ AttributeDescription) AttributeType() AttributeType {
	rv := objc.Send[AttributeType](a_.ID, objc.Sel("attributeType"))
	return rv
}


// The attribute’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/attributeType-swift.property
func (a_ AttributeDescription) SetAttributeType(value AttributeType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributeType:"), value)
}


// The default value of the attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/defaultValue
func (a_ AttributeDescription) DefaultValue() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("defaultValue"))
	return rv
}


// The default value of the attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/defaultValue
func (a_ AttributeDescription) SetDefaultValue(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDefaultValue:"), value)
}


// A Boolean value that indicates whether the attribute records its value in the persistent history transaction for a managed object’s deletion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/preservesValueInHistoryOnDeletion
func (a_ AttributeDescription) PreservesValueInHistoryOnDeletion() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("preservesValueInHistoryOnDeletion"))
	return rv
}


// A Boolean value that indicates whether the attribute records its value in the persistent history transaction for a managed object’s deletion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/preservesValueInHistoryOnDeletion
func (a_ AttributeDescription) SetPreservesValueInHistoryOnDeletion(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreservesValueInHistoryOnDeletion:"), value)
}


// A Boolean value that indicates whether the attribute allows external binary storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/allowsexternalbinarydatastorage
func (a_ AttributeDescription) AllowsExternalBinaryDataStorage() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsExternalBinaryDataStorage"))
	return rv
}


// A Boolean value that indicates whether the attribute allows external binary storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/allowsexternalbinarydatastorage
func (a_ AttributeDescription) SetAllowsExternalBinaryDataStorage(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsExternalBinaryDataStorage:"), value)
}


// The class name that represents the attribute’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/attributevalueclassname
func (a_ AttributeDescription) AttributeValueClassName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("attributeValueClassName"))
	return rv
}


// The class name that represents the attribute’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/attributevalueclassname
func (a_ AttributeDescription) SetAttributeValueClassName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributeValueClassName:"), value)
}


// The attribute’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/type
func (a_ AttributeDescription) Type() AttributeType {
	rv := objc.Send[AttributeType](a_.ID, objc.Sel("type"))
	return rv
}


// The attribute’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/type
func (a_ AttributeDescription) SetType(value AttributeType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setType:"), value)
}


// The name of the transformer to use for the attribute value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/valuetransformername
func (a_ AttributeDescription) ValueTransformerName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("valueTransformerName"))
	return rv
}


// The name of the transformer to use for the attribute value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/valuetransformername
func (a_ AttributeDescription) SetValueTransformerName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setValueTransformerName:"), value)
}


// The version hash for the attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/versionhash
func (a_ AttributeDescription) VersionHash() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("versionHash"))
	return rv
}


// The version hash for the attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/versionhash
func (a_ AttributeDescription) SetVersionHash(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVersionHash:"), value)
}



