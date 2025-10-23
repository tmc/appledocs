// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ABGroup] class.
var (
	ABGroupClass     _ABGroupClass
	ABGroupClassOnce sync.Once
)

func getABGroupClass() _ABGroupClass {
	ABGroupClassOnce.Do(func() {
		ABGroupClass = _ABGroupClass{objc.GetClass("ABGroup")}
	})
	return ABGroupClass
}

type _ABGroupClass struct {
	class objc.Class
}

// An interface definition for the [ABGroup] class.
type IABGroup interface {
	IABRecord
	AddMember(person IABPerson) bool
	AddSubgroup(group IABGroup) bool
	DistributionIdentifierForPropertyPerson(property string, person IABPerson) foundation.String
	Members() foundation.Array
	ParentGroups() foundation.Array
	RemoveMember(person IABPerson) bool
	RemoveSubgroup(group IABGroup) bool
	SetDistributionIdentifierForPropertyPerson(identifier string, property string, person IABPerson) bool
	Subgroups() foundation.Array
}

// An object that represents a group of records in the Address Book database.
//
// The class supports the concept of a “group” containing one or more persons. People may belong to multiple groups, and groups may also belong to other groups unless the relationship causes a circular reference. The only predefined property of a group is its name. However, similar to person records, you can add your own properties to group records. Groups not only help to organize person records, but also allow you to create email distribution lists. The class is “toll-free bridged” with its procedural C opaque-type counterpart. This means that the type is interchangeable in function or method calls with instances of the class.


// An object that represents a group of records in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup
type ABGroup struct {
	ABRecord
}

// ABGroupFrom constructs a [ABGroup] from an unsafe.Pointer.
//
// An object that represents a group of records in the Address Book database.
func ABGroupFrom(ptr unsafe.Pointer) ABGroup {
	return ABGroup{
		ABRecord: ABRecordFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _ABGroupClass) Alloc() ABGroup {
	rv := objc.Send[ABGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ABGroupClass) New() ABGroup {
	rv := objc.Send[ABGroup](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ABGroup) Init() ABGroup {
	rv := objc.Send[ABGroup](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ABGroup) Autorelease() ABGroup {
	rv := objc.Send[ABGroup](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewABGroup creates a new ABGroup instance.
func NewABGroup() ABGroup {
	return getABGroupClass().New()
}



// Adds the given properties to all records of this type in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/addPropertiesAndTypes(_:)
func (ac _ABGroupClass) AddPropertiesAndTypes(properties objectivec.IObject) int {
	rv := objc.Send[int](objc.ID(ac.class), objc.Sel("addPropertiesAndTypes:"), properties)
	return rv
}


// Returns an array of the names of all the properties for this record type in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/properties()
func (ac _ABGroupClass) Properties() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(ac.class), objc.Sel("properties"))
	return rv
}


// Removes the given properties from all the records of this type in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/removeProperties(_:)
func (ac _ABGroupClass) RemoveProperties(properties objectivec.IObject) int {
	rv := objc.Send[int](objc.ID(ac.class), objc.Sel("removeProperties:"), properties)
	return rv
}


// Returns a search element object that searches for records of this type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/searchElement(forProperty:label:key:value:comparison:)
func (ac _ABGroupClass) SearchElementForPropertyLabelKeyValueComparison(property string, label string, key string, value objectivec.IObject, comparison IABSearchComparison) ABSearchElement {
	rv := objc.Send[ABSearchElement](objc.ID(ac.class), objc.Sel("searchElementForProperty:label:key:value:comparison:"), objc.String(property), objc.String(label), objc.String(key), value, comparison)
	return rv
}


// Returns the type for a given property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/type(ofProperty:)
func (ac _ABGroupClass) TypeOfProperty(property string) ABPropertyType {
	rv := objc.Send[ABPropertyType](objc.ID(ac.class), objc.Sel("typeOfProperty:"), objc.String(property))
	return rv
}


// Adds a person to a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/addMember(_:)
func (a_ ABGroup) AddMember(person IABPerson) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("addMember:"), person)
	return rv
}


// Adds a subgroup to another group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/addSubgroup(_:)
func (a_ ABGroup) AddSubgroup(group IABGroup) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("addSubgroup:"), group)
	return rv
}


// Returns the distribution identifier for the given property and person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/distributionIdentifier(forProperty:person:)
func (a_ ABGroup) DistributionIdentifierForPropertyPerson(property string, person IABPerson) foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("distributionIdentifierForProperty:person:"), objc.String(property), person)
	return rv
}


// Returns an array of persons in a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/members()
func (a_ ABGroup) Members() foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("members"))
	return rv
}


// Returns an array containing a group’s parents—that is, the groups that a group belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/parentGroups()
func (a_ ABGroup) ParentGroups() foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("parentGroups"))
	return rv
}


// Removes a person from a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/removeMember(_:)
func (a_ ABGroup) RemoveMember(person IABPerson) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeMember:"), person)
	return rv
}


// Removes a subgroup from a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/removeSubgroup(_:)
func (a_ ABGroup) RemoveSubgroup(group IABGroup) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeSubgroup:"), group)
	return rv
}


// Assigns a specific distribution identifier for a person’s multivalue list property so that the group can be used as a distribution list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/setDistributionIdentifier(_:forProperty:person:)
func (a_ ABGroup) SetDistributionIdentifierForPropertyPerson(identifier string, property string, person IABPerson) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setDistributionIdentifier:forProperty:person:"), objc.String(identifier), objc.String(property), person)
	return rv
}


// Returns an array containing a group’s subgroups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABGroup/subgroups()
func (a_ ABGroup) Subgroups() foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("subgroups"))
	return rv
}



