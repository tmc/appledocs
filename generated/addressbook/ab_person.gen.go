// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ABPerson] class.
var (
	ABPersonClass     _ABPersonClass
	ABPersonClassOnce sync.Once
)

func getABPersonClass() _ABPersonClass {
	ABPersonClassOnce.Do(func() {
		ABPersonClass = _ABPersonClass{objc.GetClass("ABPerson")}
	})
	return ABPersonClass
}

type _ABPersonClass struct {
	class objc.Class
}

// An interface definition for the [ABPerson] class.
type IABPerson interface {
	IABRecord
	BeginLoadingImageDataForClient(client objectivec.IObject) int
	ImageData() foundation.Data
	LinkedPeople() foundation.Array
	ParentGroups() foundation.Array
	SetImageData(data foundation.IData) bool
	VCardRepresentation() foundation.Data
}

// An object that encapsulates all information about a person in the Address Book database.
//
// An object corresponds to a single person record in the database. A person object contains the person’s name, company, address, email addresses, and phone numbers. The class is “toll-free bridged” with its procedural C opaque-type counterpart. This means that the type is interchangeable in function or method calls with instances of the class.


// An object that encapsulates all information about a person in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson

type ABPerson struct {
	ABRecord
}

// ABPersonFrom constructs a [ABPerson] from an unsafe.Pointer.
//
// An object that encapsulates all information about a person in the Address Book database.
func ABPersonFrom(ptr unsafe.Pointer) ABPerson {
	return ABPerson{
		ABRecord: ABRecordFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _ABPersonClass) Alloc() ABPerson {
	rv := objc.Send[ABPerson](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ABPersonClass) New() ABPerson {
	rv := objc.Send[ABPerson](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ABPerson) Init() ABPerson {
	rv := objc.Send[ABPerson](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ABPerson) Autorelease() ABPerson {
	rv := objc.Send[ABPerson](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewABPerson creates a new ABPerson instance.
func NewABPerson() ABPerson {
	return getABPersonClass().New()
}




// Returns an instance initialized with the given data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/init(VCardRepresentation:)

func NewABPersonWithVCardRepresentation(vCardData foundation.IData) ABPerson {
	instance := getABPersonClass().Alloc()
	rv := objc.Send[ABPerson](instance.ID, objc.Sel("initWithVCardRepresentation:"), vCardData)
	rv.Autorelease()
	return rv
}



// Adds the given properties to all the records of this type in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/addPropertiesAndTypes(_:)

func (ac _ABPersonClass) AddPropertiesAndTypes(properties objectivec.IObject) int {
	rv := objc.Send[int](objc.ID(ac.class), objc.Sel("addPropertiesAndTypes:"), properties)
	return rv
}


// Cancels an asynchronous fetch of the images for a given tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/cancelLoadingImageData(forTag:)

func (ac _ABPersonClass) CancelLoadingImageDataForTag(tag int) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("cancelLoadingImageDataForTag:"), tag)
}


// Returns an array of the names of all the properties for the record in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/properties()

func (ac _ABPersonClass) Properties() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(ac.class), objc.Sel("properties"))
	return rv
}


// Removes the given properties from all the records of this type in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/removeProperties(_:)

func (ac _ABPersonClass) RemoveProperties(properties objectivec.IObject) int {
	rv := objc.Send[int](objc.ID(ac.class), objc.Sel("removeProperties:"), properties)
	return rv
}


// Returns a search element object that specifies a query for records of this type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/searchElement(forProperty:label:key:value:comparison:)

func (ac _ABPersonClass) SearchElementForPropertyLabelKeyValueComparison(property string, label string, key string, value objectivec.IObject, comparison IABSearchComparison) ABSearchElement {
	rv := objc.Send[ABSearchElement](objc.ID(ac.class), objc.Sel("searchElementForProperty:label:key:value:comparison:"), objc.String(property), objc.String(label), objc.String(key), value, comparison)
	return rv
}


// Returns the type of a given property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/type(ofProperty:)

func (ac _ABPersonClass) TypeOfProperty(property string) ABPropertyType {
	rv := objc.Send[ABPropertyType](objc.ID(ac.class), objc.Sel("typeOfProperty:"), objc.String(property))
	return rv
}



// Starts an asynchronous fetch for image data in all locations
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/beginLoadingImageData(for:)

func (a_ ABPerson) BeginLoadingImageDataForClient(client objectivec.IObject) int {
	rv := objc.Send[int](a_.ID, objc.Sel("beginLoadingImageDataForClient:"), client)
	return rv
}



// Returns data that contains a picture of this person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/imageData()

func (a_ ABPerson) ImageData() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("imageData"))
	return rv
}



// Returns the array of all person records that are linked to the person this record represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/linkedPeople()

func (a_ ABPerson) LinkedPeople() foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("linkedPeople"))
	return rv
}



// Returns an array of the address book groups that this person belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/parentGroups()

func (a_ ABPerson) ParentGroups() foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("parentGroups"))
	return rv
}



// Sets the image for this person to the given data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/setImageData(_:)

func (a_ ABPerson) SetImageData(data foundation.IData) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setImageData:"), data)
	return rv
}



// Returns the vCard representation of the person record as a data object in vCard format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/vCardRepresentation()

func (a_ ABPerson) VCardRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("vCardRepresentation"))
	return rv
}


