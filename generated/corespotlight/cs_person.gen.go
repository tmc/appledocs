// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CSPerson] class.
var (
	CSPersonClass     _CSPersonClass
	CSPersonClassOnce sync.Once
)

func getCSPersonClass() _CSPersonClass {
	CSPersonClassOnce.Do(func() {
		CSPersonClass = _CSPersonClass{objc.GetClass("CSPerson")}
	})
	return CSPersonClass
}

type _CSPersonClass struct {
	class objc.Class
}

// An interface definition for the [CSPerson] class.
type ICSPerson interface {
	objectivec.IObject
	// properties:
	ContactIdentifier() string /* primitive/slice/pointer. */
	SetContactIdentifier(value string /* primitive/slice/pointer. */)
	DisplayName() string /* primitive/slice/pointer. */
	HandleIdentifier() string /* primitive/slice/pointer. */
	Handles() []string /* primitive/slice/pointer. */
	// methods:
}

// An object that represents a person in the context of search results.
//
// A object represents a person in the context of search results. You can create a object when you have a display name and a contact handle of some kind, such as an email address or phone number. If you create a object to represent a specific contact, you can use the value of the contact’s identifier property for the person object’s property. Using the same value lets you avoid using names or phone numbers to look up the contact that’s associated with a person.


// An object that represents a person in the context of search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson
type CSPerson struct {
	objectivec.Object
}

// CSPersonFrom constructs a [CSPerson] from an unsafe.Pointer.
//
// An object that represents a person in the context of search results.
func CSPersonFrom(ptr unsafe.Pointer) CSPerson {
	return CSPerson{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CSPersonClass) Alloc() CSPerson {
	rv := objc.Send[CSPerson](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CSPersonClass) New() CSPerson {
	rv := objc.Send[CSPerson](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSPerson) Init() CSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSPerson) Autorelease() CSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSPerson creates a new CSPerson instance.
func NewCSPerson() CSPerson {
	return getCSPersonClass().New()
}



// Returns a new object initialized with the specified display name and contact attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/init(displayName:handles:handleIdentifier:)
func NewCSPersonWithDisplayNameHandlesHandleIdentifier(displayName string /* primitive/slice/pointer. */, handles []string /* primitive/slice/pointer. */, handleIdentifier string /* primitive/slice/pointer. */) CSPerson {
	instance := getCSPersonClass().Alloc()
	rv := objc.Send[CSPerson](instance.ID, objc.Sel("initWithDisplayName:handles:handleIdentifier:"), objc.String(displayName), handles, objc.String(handleIdentifier))
	rv.Autorelease()
	return rv
}



// The identifier for the contact associated with the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/contactIdentifier
func (c_ CSPerson) ContactIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("contactIdentifier"))
	return rv
}


// The identifier for the contact associated with the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/contactIdentifier
func (c_ CSPerson) SetContactIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactIdentifier:"), objc.String(value))
}


// A display name for the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/displayName
func (c_ CSPerson) DisplayName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("displayName"))
	return rv
}


// A key that identifies the type of contact property represented by the person object’s handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/handleIdentifier
func (c_ CSPerson) HandleIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("handleIdentifier"))
	return rv
}


// An array of contact handles related to the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/handles
func (c_ CSPerson) Handles() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](c_.ID, objc.Sel("handles"))
	return rv
}


