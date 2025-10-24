// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ABPerson */


/* debug [class_header]: Header for ABPerson */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ABPerson */
// An interface definition for the [ABPerson] class.
type IABPerson interface {
	IABRecord
	
/* debug [class_interface_properties]: Properties for ABPerson */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ABPerson */
	// methods:
	BeginLoadingImageDataForClient(client unsafe.Pointer) int
	ImageData() foundation.Data
	LinkedPeople() foundation.Array
	ParentGroups() foundation.Array
	SetImageData(data objc.IObject /* cross-framework: NSData */) bool
	VCardRepresentation() foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ABPerson */
// Alloc allocates a new instance without initialization.
func (ac _ABPersonClass) Alloc() ABPerson {
	rv := objc.Send[ABPerson](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ABPerson */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ABPerson */

// Returns an instance initialized with the given data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/init(VCardRepresentation:)
func NewABPersonWithVCardRepresentation(vCardData objc.IObject /* cross-framework: NSData */) ABPerson {
	instance := getABPersonClass().Alloc()
	rv := objc.Send[ABPerson](instance.ID, objc.Sel("initWithVCardRepresentation:"), vCardData)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewABPersonWithVCardRepresentation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ABPerson */

// Adds the given properties to all the records of this type in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/addPropertiesAndTypes(_:)
func (ac _ABPersonClass) AddPropertiesAndTypes(properties objc.IObject /* cross-framework: NSDictionary */) int {
	rv := objc.Send[int](objc.ID(ac.class), objc.Sel("addPropertiesAndTypes:"), properties)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AddPropertiesAndTypes) */


// Cancels an asynchronous fetch of the images for a given tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/cancelLoadingImageData(forTag:)
func (ac _ABPersonClass) CancelLoadingImageDataForTag(tag int) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("cancelLoadingImageDataForTag:"), tag)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CancelLoadingImageDataForTag) */


// Returns an array of the names of all the properties for the record in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/properties()
func (ac _ABPersonClass) Properties() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(ac.class), objc.Sel("properties"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Properties) */


// Removes the given properties from all the records of this type in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/removeProperties(_:)
func (ac _ABPersonClass) RemoveProperties(properties objc.IObject /* cross-framework: NSArray */) int {
	rv := objc.Send[int](objc.ID(ac.class), objc.Sel("removeProperties:"), properties)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RemoveProperties) */


// Returns a search element object that specifies a query for records of this type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/searchElement(forProperty:label:key:value:comparison:)
func (ac _ABPersonClass) SearchElementForPropertyLabelKeyValueComparison(property objc.IObject /* cross-framework: NSString */, label objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, value objc.IObject, comparison ABSearchComparison /* typedef */) IABSearchElement {
	rv := objc.Send[ABSearchElement](objc.ID(ac.class), objc.Sel("searchElementForProperty:label:key:value:comparison:"), property, label, key, value, comparison)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SearchElementForPropertyLabelKeyValueComparison) */


// Returns the type of a given property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/type(ofProperty:)
func (ac _ABPersonClass) TypeOfProperty(property objc.IObject /* cross-framework: NSString */) ABPropertyType /* typedef */ {
	rv := objc.Send[uint32](objc.ID(ac.class), objc.Sel("typeOfProperty:"), property)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TypeOfProperty) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ABPerson */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ABPerson */

// Starts an asynchronous fetch for image data in all locations
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/beginLoadingImageData(for:)
func (a_ ABPerson) BeginLoadingImageDataForClient(client unsafe.Pointer) int {
	rv := objc.Send[int](a_.ID, objc.Sel("beginLoadingImageDataForClient:"), client)
	return rv
}/* debug [instance_methods/method]: BeginLoadingImageDataForClient */


// Returns data that contains a picture of this person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/imageData()
func (a_ ABPerson) ImageData() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("imageData"))
	return rv
}/* debug [instance_methods/method]: ImageData */


// Returns the array of all person records that are linked to the person this record represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/linkedPeople()
func (a_ ABPerson) LinkedPeople() foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("linkedPeople"))
	return rv
}/* debug [instance_methods/method]: LinkedPeople */


// Returns an array of the address book groups that this person belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/parentGroups()
func (a_ ABPerson) ParentGroups() foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("parentGroups"))
	return rv
}/* debug [instance_methods/method]: ParentGroups */


// Sets the image for this person to the given data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/setImageData(_:)
func (a_ ABPerson) SetImageData(data objc.IObject /* cross-framework: NSData */) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setImageData:"), data)
	return rv
}/* debug [instance_methods/method]: SetImageData */


// Returns the vCard representation of the person record as a data object in vCard format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPerson/vCardRepresentation()
func (a_ ABPerson) VCardRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("vCardRepresentation"))
	return rv
}/* debug [instance_methods/method]: VCardRepresentation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ABPerson */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ABPerson */


