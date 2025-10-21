// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RelationshipDescription] class.
var (
	RelationshipDescriptionClass     _RelationshipDescriptionClass
	RelationshipDescriptionClassOnce sync.Once
)

func getRelationshipDescriptionClass() _RelationshipDescriptionClass {
	RelationshipDescriptionClassOnce.Do(func() {
		RelationshipDescriptionClass = _RelationshipDescriptionClass{objc.GetClass("NSRelationshipDescription")}
	})
	return RelationshipDescriptionClass
}

type _RelationshipDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [RelationshipDescription] class.
type IRelationshipDescription interface {
	IPropertyDescription
}

// A description of a relationship between two entities.
//
// provides additional attributes that are specific to modeling a relationship between two entities. For the common attributes of all property types, see . For example, use this class to define a relationship’s — the number of managed objects the relationship can reference. For a to-one relationship, set to . For a to-many relationship, set to a number greater than to impose an upper limit; otherwise, use to allow an unlimited number of referenced objects. At runtime, you can modify a relationship description until you associate its owning managed object model with a persistent store coordinator. If you attempt to modify the model after you associate it, Core Data throws an exception. To modify a model that’s in use, create and modify a copy and then discard any objects that belong to the original model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription
type RelationshipDescription struct {
	PropertyDescription
}

// RelationshipDescriptionFrom constructs a [RelationshipDescription] from an unsafe.Pointer.
//
// A description of a relationship between two entities.
func RelationshipDescriptionFrom(ptr unsafe.Pointer) RelationshipDescription {
	return RelationshipDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RelationshipDescriptionClass) Alloc() RelationshipDescription {
	rv := objc.Send[RelationshipDescription](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RelationshipDescriptionClass) New() RelationshipDescription {
	rv := objc.Send[RelationshipDescription](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RelationshipDescription) Init() RelationshipDescription {
	rv := objc.Send[RelationshipDescription](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RelationshipDescription) Autorelease() RelationshipDescription {
	rv := objc.Send[RelationshipDescription](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRelationshipDescription creates a new RelationshipDescription instance.
func NewRelationshipDescription() RelationshipDescription {
	return getRelationshipDescriptionClass().New()
}


// A Boolean value that determines whether the relationship preserves the order of the referenced managed objects.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/isordered
func (r_ RelationshipDescription) IsOrdered() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isOrdered"))
	return rv
}


// SetIsOrdered sets the value of the isOrdered property.
// A Boolean value that determines whether the relationship preserves the order of the referenced managed objects.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/isordered
func (r_ RelationshipDescription) SetIsOrdered(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsOrdered:"), value)
}

// The maximum number of managed objects the relationship can reference.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/maxcount
func (r_ RelationshipDescription) MaxCount() int {
	rv := objc.Send[int](r_.ID, objc.Sel("maxCount"))
	return rv
}


// SetMaxCount sets the value of the maxCount property.
// The maximum number of managed objects the relationship can reference.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/maxcount
func (r_ RelationshipDescription) SetMaxCount(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaxCount:"), value)
}

// The relationship that represents the inverse of the current relationship.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/inverserelationship
func (r_ RelationshipDescription) InverseRelationship() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("inverseRelationship"))
	return rv
}


// SetInverseRelationship sets the value of the inverseRelationship property.
// The relationship that represents the inverse of the current relationship.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/inverserelationship
func (r_ RelationshipDescription) SetInverseRelationship(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInverseRelationship:"), value)
}

// The type of object the relationship contains.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/destinationentity
func (r_ RelationshipDescription) DestinationEntity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("destinationEntity"))
	return rv
}


// SetDestinationEntity sets the value of the destinationEntity property.
// The type of object the relationship contains.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/destinationentity
func (r_ RelationshipDescription) SetDestinationEntity(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDestinationEntity:"), value)
}

// The minimum number of managed objects the relationship can reference.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/mincount
func (r_ RelationshipDescription) MinCount() int {
	rv := objc.Send[int](r_.ID, objc.Sel("minCount"))
	return rv
}


// SetMinCount sets the value of the minCount property.
// The minimum number of managed objects the relationship can reference.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/mincount
func (r_ RelationshipDescription) SetMinCount(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMinCount:"), value)
}

// Returns a Boolean value that indicates whether the relationship can contain many managed objects.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/istomany
func (r_ RelationshipDescription) IsToMany() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isToMany"))
	return rv
}


// SetIsToMany sets the value of the isToMany property.
// Returns a Boolean value that indicates whether the relationship can contain many managed objects.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/istomany
func (r_ RelationshipDescription) SetIsToMany(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsToMany:"), value)
}

// The relationship’s unique identity.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/versionhash
func (r_ RelationshipDescription) VersionHash() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("versionHash"))
	return rv
}


// SetVersionHash sets the value of the versionHash property.
// The relationship’s unique identity.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/versionhash
func (r_ RelationshipDescription) SetVersionHash(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setVersionHash:"), value)
}

// The rule to apply when you delete the relationship’s owning managed object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription/deleteRule
func (r_ RelationshipDescription) DeleteRule() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("deleteRule"))
	return rv
}


// SetDeleteRule sets the value of the deleteRule property.
// The rule to apply when you delete the relationship’s owning managed object.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription/deleteRule
func (r_ RelationshipDescription) SetDeleteRule(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDeleteRule:"), value)
}



