// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Name() string /* primitive/slice/pointer. */
	SetName(value string /* primitive/slice/pointer. */)
	Entity() IEntityDescription
	SetEntity(value IEntityDescription)
	IsIndexed() bool /* primitive/slice/pointer. */
	SetIsIndexed(value bool /* primitive/slice/pointer. */)
	IsIndexedBySpotlight() bool /* primitive/slice/pointer. */
	SetIsIndexedBySpotlight(value bool /* primitive/slice/pointer. */)
	IsOptional() bool /* primitive/slice/pointer. */
	SetIsOptional(value bool /* primitive/slice/pointer. */)
	IsStoredInExternalRecord() bool /* primitive/slice/pointer. */
	SetIsStoredInExternalRecord(value bool /* primitive/slice/pointer. */)
	IsTransient() bool /* primitive/slice/pointer. */
	SetIsTransient(value bool /* primitive/slice/pointer. */)
	RenamingIdentifier() string /* primitive/slice/pointer. */
	SetRenamingIdentifier(value string /* primitive/slice/pointer. */)
	UserInfo() unsafe.Pointer
	SetUserInfo(value unsafe.Pointer)
	ValidationPredicates() objc.IObject /* cross-framework: Predicate */
	SetValidationPredicates(value objc.IObject /* cross-framework: Predicate */)
	ValidationWarnings() unsafe.Pointer
	SetValidationWarnings(value unsafe.Pointer)
	VersionHash() foundation.objc.IObject /* cross-framework: Data */
	SetVersionHash(value foundation.objc.IObject /* cross-framework: Data */)
	VersionHashModifier() string /* primitive/slice/pointer. */
	SetVersionHashModifier(value string /* primitive/slice/pointer. */)
	// methods:
}

// A description of a single property belonging to an entity.
//
// A property describes a single value within an object managed by the Core Data Framework. There are different types of property, each represented by a subclass which encapsulates the specific property behavior—see , , and . Note that a property name cannot be the same as any no-parameter method name of or . For example, you cannot give a property the name “description”. There are hundreds of methods on which may conflict with property names—and this list can grow without warning from frameworks or other libraries. You should avoid very general words (like “font”, and “color”) and words or phrases which overlap with Cocoa paradigms (such as “isEditing” and “objectSpecifier”). Properties—relationships as well as attributes—may be transient. A managed object context knows about transient properties and tracks changes made to them. Transient properties are ignored by the persistent store, and not just during saves: you cannot fetch using a predicate based on transients (although you can use transient properties to filter in memory yourself).


// A description of a single property belonging to an entity.
//
// [Full Topic]
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



// The name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/name
func (p_ PropertyDescription) Name() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("name"))
	return rv
}


// The name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/name
func (p_ PropertyDescription) SetName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setName:"), objc.String(value))
}


// The entity description of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/entity
func (p_ PropertyDescription) Entity() IEntityDescription {
	rv := objc.Send[EntityDescription](p_.ID, objc.Sel("entity"))
	return rv
}


// The entity description of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/entity
func (p_ PropertyDescription) SetEntity(value IEntityDescription) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEntity:"), value)
}


// A Boolean value that indicates whether the receiver should be indexed for searching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isindexed
func (p_ PropertyDescription) IsIndexed() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isIndexed"))
	return rv
}


// A Boolean value that indicates whether the receiver should be indexed for searching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isindexed
func (p_ PropertyDescription) SetIsIndexed(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsIndexed:"), value)
}


// A Boolean value that indicates whether Core Data adds the property’s value to the Core Spotlight index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isindexedbyspotlight
func (p_ PropertyDescription) IsIndexedBySpotlight() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isIndexedBySpotlight"))
	return rv
}


// A Boolean value that indicates whether Core Data adds the property’s value to the Core Spotlight index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isindexedbyspotlight
func (p_ PropertyDescription) SetIsIndexedBySpotlight(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsIndexedBySpotlight:"), value)
}


// A Boolean value that indicates whether the receiver is optional.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isoptional
func (p_ PropertyDescription) IsOptional() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isOptional"))
	return rv
}


// A Boolean value that indicates whether the receiver is optional.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isoptional
func (p_ PropertyDescription) SetIsOptional(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsOptional:"), value)
}


// A Boolean value that indicates whether to write the property’s data in an external record file that corresponds to the managed object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isstoredinexternalrecord
func (p_ PropertyDescription) IsStoredInExternalRecord() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isStoredInExternalRecord"))
	return rv
}


// A Boolean value that indicates whether to write the property’s data in an external record file that corresponds to the managed object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/isstoredinexternalrecord
func (p_ PropertyDescription) SetIsStoredInExternalRecord(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsStoredInExternalRecord:"), value)
}


// A Boolean value that indicates whether the receiver is transient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/istransient
func (p_ PropertyDescription) IsTransient() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isTransient"))
	return rv
}


// A Boolean value that indicates whether the receiver is transient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/istransient
func (p_ PropertyDescription) SetIsTransient(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsTransient:"), value)
}


// The renaming identifier for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/renamingidentifier
func (p_ PropertyDescription) RenamingIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("renamingIdentifier"))
	return rv
}


// The renaming identifier for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/renamingidentifier
func (p_ PropertyDescription) SetRenamingIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRenamingIdentifier:"), objc.String(value))
}


// The user info dictionary of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/userinfo
func (p_ PropertyDescription) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("userInfo"))
	return rv
}


// The user info dictionary of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/userinfo
func (p_ PropertyDescription) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserInfo:"), value)
}


// The validation predicates of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/validationpredicates
func (p_ PropertyDescription) ValidationPredicates() objc.IObject /* cross-framework: Predicate */ {
	rv := objc.Send[Predicate](p_.ID, objc.Sel("validationPredicates"))
	return rv
}


// The validation predicates of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/validationpredicates
func (p_ PropertyDescription) SetValidationPredicates(value objc.IObject /* cross-framework: Predicate */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValidationPredicates:"), value)
}


// The error strings associated with the receiver’s validation predicates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/validationwarnings
func (p_ PropertyDescription) ValidationWarnings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("validationWarnings"))
	return rv
}


// The error strings associated with the receiver’s validation predicates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/validationwarnings
func (p_ PropertyDescription) SetValidationWarnings(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValidationWarnings:"), value)
}


// The version hash for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/versionhash
func (p_ PropertyDescription) VersionHash() foundation.objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("versionHash"))
	return rv
}


// The version hash for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/versionhash
func (p_ PropertyDescription) SetVersionHash(value foundation.objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVersionHash:"), value)
}


// The version hash modifier for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/versionhashmodifier
func (p_ PropertyDescription) VersionHashModifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("versionHashModifier"))
	return rv
}


// The version hash modifier for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/versionhashmodifier
func (p_ PropertyDescription) SetVersionHashModifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVersionHashModifier:"), objc.String(value))
}



