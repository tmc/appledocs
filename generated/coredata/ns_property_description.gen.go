// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PropertyDescription] class.
var (
	PropertyDescriptionClass     _PropertyDescriptionClass
	PropertyDescriptionClassOnce sync.Once
)

func getPropertyDescriptionClass() _PropertyDescriptionClass {
	PropertyDescriptionClassOnce.Do(func() {
		PropertyDescriptionClass = _PropertyDescriptionClass{objc.GetClass("NSPropertyDescription")}
	})
	return PropertyDescriptionClass
}

type _PropertyDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [PropertyDescription] class.
type IPropertyDescription interface {
	objectivec.IObject
	SetValidationPredicatesWithValidationWarnings(validationPredicates unsafe.Pointer, validationWarnings unsafe.Pointer)
}

// A description of a single property belonging to an entity.
//
// A property describes a single value within an object managed by the Core Data Framework. There are different types of property, each represented by a subclass which encapsulates the specific property behavior—see , , and . Note that a property name cannot be the same as any no-parameter method name of or . For example, you cannot give a property the name “description”. There are hundreds of methods on which may conflict with property names—and this list can grow without warning from frameworks or other libraries. You should avoid very general words (like “font”, and “color”) and words or phrases which overlap with Cocoa paradigms (such as “isEditing” and “objectSpecifier”). Properties—relationships as well as attributes—may be transient. A managed object context knows about transient properties and tracks changes made to them. Transient properties are ignored by the persistent store, and not just during saves: you cannot fetch using a predicate based on transients (although you can use transient properties to filter in memory yourself).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription
type PropertyDescription struct {
	objectivec.Object
}

// PropertyDescriptionFrom constructs a [PropertyDescription] from an unsafe.Pointer.
//
// A description of a single property belonging to an entity.
func PropertyDescriptionFrom(ptr unsafe.Pointer) PropertyDescription {
	return PropertyDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PropertyDescriptionClass) Alloc() PropertyDescription {
	rv := objc.Send[PropertyDescription](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PropertyDescriptionClass) New() PropertyDescription {
	rv := objc.Send[PropertyDescription](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PropertyDescription) Init() PropertyDescription {
	rv := objc.Send[PropertyDescription](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PropertyDescription) Autorelease() PropertyDescription {
	rv := objc.Send[PropertyDescription](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPropertyDescription creates a new PropertyDescription instance.
func NewPropertyDescription() PropertyDescription {
	return getPropertyDescriptionClass().New()
}


// Sets the validation predicates and warnings of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/setValidationPredicates(_:withValidationWarnings:)
func (p_ PropertyDescription) SetValidationPredicatesWithValidationWarnings(validationPredicates unsafe.Pointer, validationWarnings unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValidationPredicates:withValidationWarnings:"), validationPredicates, validationWarnings)
}

// The entity description of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/entity
func (p_ PropertyDescription) Entity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("entity"))
	return rv
}

// A Boolean value that indicates whether the receiver should be indexed for searching.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isIndexed
func (p_ PropertyDescription) Indexed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("indexed"))
	return rv
}


// SetIndexed sets the value of the indexed property.
// A Boolean value that indicates whether the receiver should be indexed for searching.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isIndexed
func (p_ PropertyDescription) SetIndexed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexed:"), value)
}

// A Boolean value that indicates whether Core Data adds the property’s value to the Core Spotlight index.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isIndexedBySpotlight
func (p_ PropertyDescription) IndexedBySpotlight() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("indexedBySpotlight"))
	return rv
}


// SetIndexedBySpotlight sets the value of the indexedBySpotlight property.
// A Boolean value that indicates whether Core Data adds the property’s value to the Core Spotlight index.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isIndexedBySpotlight
func (p_ PropertyDescription) SetIndexedBySpotlight(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexedBySpotlight:"), value)
}

// A Boolean value that indicates whether the receiver is optional.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isOptional
func (p_ PropertyDescription) Optional() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("optional"))
	return rv
}


// SetOptional sets the value of the optional property.
// A Boolean value that indicates whether the receiver is optional.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isOptional
func (p_ PropertyDescription) SetOptional(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOptional:"), value)
}

// A Boolean value that indicates whether to write the property’s data in an external record file that corresponds to the managed object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isStoredInExternalRecord
func (p_ PropertyDescription) StoredInExternalRecord() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("storedInExternalRecord"))
	return rv
}


// SetStoredInExternalRecord sets the value of the storedInExternalRecord property.
// A Boolean value that indicates whether to write the property’s data in an external record file that corresponds to the managed object.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isStoredInExternalRecord
func (p_ PropertyDescription) SetStoredInExternalRecord(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStoredInExternalRecord:"), value)
}

// A Boolean value that indicates whether the receiver is transient.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isTransient
func (p_ PropertyDescription) Transient() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("transient"))
	return rv
}


// SetTransient sets the value of the transient property.
// A Boolean value that indicates whether the receiver is transient.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isTransient
func (p_ PropertyDescription) SetTransient(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransient:"), value)
}

// The name of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/name
func (p_ PropertyDescription) Name() string {
	rv := objc.Send[string](p_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/name
func (p_ PropertyDescription) SetName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setName:"), objc.String(value))
}

// The renaming identifier for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/renamingIdentifier
func (p_ PropertyDescription) RenamingIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("renamingIdentifier"))
	return rv
}


// SetRenamingIdentifier sets the value of the renamingIdentifier property.
// The renaming identifier for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/renamingIdentifier
func (p_ PropertyDescription) SetRenamingIdentifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRenamingIdentifier:"), objc.String(value))
}

// The user info dictionary of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/userInfo
func (p_ PropertyDescription) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("userInfo"))
	return rv
}


// SetUserInfo sets the value of the userInfo property.
// The user info dictionary of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/userInfo
func (p_ PropertyDescription) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserInfo:"), value)
}

// The validation predicates of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/validationPredicates
func (p_ PropertyDescription) ValidationPredicates() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](p_.ID, objc.Sel("validationPredicates"))
	return rv
}

// The error strings associated with the receiver’s validation predicates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/validationWarnings
func (p_ PropertyDescription) ValidationWarnings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("validationWarnings"))
	return rv
}

// The version hash for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/versionHash
func (p_ PropertyDescription) VersionHash() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("versionHash"))
	return rv
}

// The version hash modifier for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/versionHashModifier
func (p_ PropertyDescription) VersionHashModifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("versionHashModifier"))
	return rv
}


// SetVersionHashModifier sets the value of the versionHashModifier property.
// The version hash modifier for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/versionHashModifier
func (p_ PropertyDescription) SetVersionHashModifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVersionHashModifier:"), objc.String(value))
}

// A Boolean value that indicates whether the receiver should be indexed for searching.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isindexed
func (p_ PropertyDescription) IsIndexed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isIndexed"))
	return rv
}


// SetIsIndexed sets the value of the isIndexed property.
// A Boolean value that indicates whether the receiver should be indexed for searching.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isindexed
func (p_ PropertyDescription) SetIsIndexed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsIndexed:"), value)
}

// A Boolean value that indicates whether Core Data adds the property’s value to the Core Spotlight index.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isindexedbyspotlight
func (p_ PropertyDescription) IsIndexedBySpotlight() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isIndexedBySpotlight"))
	return rv
}


// SetIsIndexedBySpotlight sets the value of the isIndexedBySpotlight property.
// A Boolean value that indicates whether Core Data adds the property’s value to the Core Spotlight index.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isindexedbyspotlight
func (p_ PropertyDescription) SetIsIndexedBySpotlight(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsIndexedBySpotlight:"), value)
}

// A Boolean value that indicates whether the receiver is optional.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isoptional
func (p_ PropertyDescription) IsOptional() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isOptional"))
	return rv
}


// SetIsOptional sets the value of the isOptional property.
// A Boolean value that indicates whether the receiver is optional.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isoptional
func (p_ PropertyDescription) SetIsOptional(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsOptional:"), value)
}

// A Boolean value that indicates whether to write the property’s data in an external record file that corresponds to the managed object.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isstoredinexternalrecord
func (p_ PropertyDescription) IsStoredInExternalRecord() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isStoredInExternalRecord"))
	return rv
}


// SetIsStoredInExternalRecord sets the value of the isStoredInExternalRecord property.
// A Boolean value that indicates whether to write the property’s data in an external record file that corresponds to the managed object.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isstoredinexternalrecord
func (p_ PropertyDescription) SetIsStoredInExternalRecord(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsStoredInExternalRecord:"), value)
}

// A Boolean value that indicates whether the receiver is transient.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/istransient
func (p_ PropertyDescription) IsTransient() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isTransient"))
	return rv
}


// SetIsTransient sets the value of the isTransient property.
// A Boolean value that indicates whether the receiver is transient.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/istransient
func (p_ PropertyDescription) SetIsTransient(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsTransient:"), value)
}



