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
	AllowsCloudEncryption() bool
	SetAllowsCloudEncryption(value bool)
	AllowsExternalBinaryDataStorage() bool
	SetAllowsExternalBinaryDataStorage(value bool)
	AttributeType() AttributeType
	SetAttributeType(value AttributeType)
	AttributeValueClassName() string
	SetAttributeValueClassName(value string)
	DefaultValue() unsafe.Pointer
	SetDefaultValue(value unsafe.Pointer)
	PreservesValueInHistoryOnDeletion() bool
	SetPreservesValueInHistoryOnDeletion(value bool)
	Type() AttributeType
	SetType(value AttributeType)
	ValueTransformerName() string
	SetValueTransformerName(value string)
	VersionHash() foundation.Data
	SetVersionHash(value foundation.IData)
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

// A Boolean value that indicates whether the attribute allows external binary storage.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/allowsExternalBinaryDataStorage
func (a_ AttributeDescription) AllowsExternalBinaryDataStorage() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsExternalBinaryDataStorage"))
	return rv
}


// SetAllowsExternalBinaryDataStorage sets the value of the allowsExternalBinaryDataStorage property.
// A Boolean value that indicates whether the attribute allows external binary storage.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/allowsExternalBinaryDataStorage
func (a_ AttributeDescription) SetAllowsExternalBinaryDataStorage(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsExternalBinaryDataStorage:"), value)
}

// The attribute’s type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/attributeType-swift.property
func (a_ AttributeDescription) AttributeType() AttributeType {
	rv := objc.Send[AttributeType](a_.ID, objc.Sel("attributeType"))
	return rv
}


// SetAttributeType sets the value of the attributeType property.
// The attribute’s type.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/attributeType-swift.property
func (a_ AttributeDescription) SetAttributeType(value AttributeType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributeType:"), value)
}

// The class name that represents the attribute’s value.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/attributevalueclassname
func (a_ AttributeDescription) AttributeValueClassName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("attributeValueClassName"))
	return rv
}


// SetAttributeValueClassName sets the value of the attributeValueClassName property.
// The class name that represents the attribute’s value.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/attributevalueclassname
func (a_ AttributeDescription) SetAttributeValueClassName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributeValueClassName:"), objc.String(value))
}

// The default value of the attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/defaultvalue
func (a_ AttributeDescription) DefaultValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("defaultValue"))
	return rv
}


// SetDefaultValue sets the value of the defaultValue property.
// The default value of the attribute.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/defaultvalue
func (a_ AttributeDescription) SetDefaultValue(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDefaultValue:"), value)
}

// A Boolean value that indicates whether the attribute records its value in the persistent history transaction for a managed object’s deletion.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/preservesvalueinhistoryondeletion
func (a_ AttributeDescription) PreservesValueInHistoryOnDeletion() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("preservesValueInHistoryOnDeletion"))
	return rv
}


// SetPreservesValueInHistoryOnDeletion sets the value of the preservesValueInHistoryOnDeletion property.
// A Boolean value that indicates whether the attribute records its value in the persistent history transaction for a managed object’s deletion.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/preservesvalueinhistoryondeletion
func (a_ AttributeDescription) SetPreservesValueInHistoryOnDeletion(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreservesValueInHistoryOnDeletion:"), value)
}

// The attribute’s type.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/type
func (a_ AttributeDescription) Type() AttributeType {
	rv := objc.Send[AttributeType](a_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The attribute’s type.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/type
func (a_ AttributeDescription) SetType(value AttributeType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setType:"), value)
}

// The name of the transformer to use for the attribute value.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/valuetransformername
func (a_ AttributeDescription) ValueTransformerName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("valueTransformerName"))
	return rv
}


// SetValueTransformerName sets the value of the valueTransformerName property.
// The name of the transformer to use for the attribute value.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/valuetransformername
func (a_ AttributeDescription) SetValueTransformerName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setValueTransformerName:"), objc.String(value))
}

// The version hash for the attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/versionhash
func (a_ AttributeDescription) VersionHash() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("versionHash"))
	return rv
}


// SetVersionHash sets the value of the versionHash property.
// The version hash for the attribute.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/versionhash
func (a_ AttributeDescription) SetVersionHash(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVersionHash:"), value)
}



