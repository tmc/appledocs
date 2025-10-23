// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EKVirtualConferenceProvider] class.
var (
	EKVirtualConferenceProviderClass     _EKVirtualConferenceProviderClass
	EKVirtualConferenceProviderClassOnce sync.Once
)

func getEKVirtualConferenceProviderClass() _EKVirtualConferenceProviderClass {
	EKVirtualConferenceProviderClassOnce.Do(func() {
		EKVirtualConferenceProviderClass = _EKVirtualConferenceProviderClass{objc.GetClass("EKVirtualConferenceProvider")}
	})
	return EKVirtualConferenceProviderClass
}

type _EKVirtualConferenceProviderClass struct {
	class objc.Class
}

// An interface definition for the [EKVirtualConferenceProvider] class.
type IEKVirtualConferenceProvider interface {
	objectivec.IObject
	FetchAvailableRoomTypesWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchVirtualConferenceForIdentifierCompletionHandler(identifier EKVirtualConferenceRoomTypeIdentifier, completionHandler unsafe.Pointer)
}

// An object that associates virtual conferencing details with an event object in a user’s calendar.
//
// lets apps that offer virtual conferencing services to integrate directly with events in users’ calendars. To add this support to your app, add a virtual conference extension. The principal class of the app extension is a custom subclass of that you create that provides the following: A list of room types where events take place, such as Personal Room or Team Room A descriptor for a virtual conference, including a user-visible title, one or more URLs, and additional details


// An object that associates virtual conferencing details with an event object in a user’s calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceProvider
type EKVirtualConferenceProvider struct {
	objectivec.Object
}

// EKVirtualConferenceProviderFrom constructs a [EKVirtualConferenceProvider] from an unsafe.Pointer.
//
// An object that associates virtual conferencing details with an event object in a user’s calendar.
func EKVirtualConferenceProviderFrom(ptr unsafe.Pointer) EKVirtualConferenceProvider {
	return EKVirtualConferenceProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EKVirtualConferenceProviderClass) Alloc() EKVirtualConferenceProvider {
	rv := objc.Send[EKVirtualConferenceProvider](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EKVirtualConferenceProviderClass) New() EKVirtualConferenceProvider {
	rv := objc.Send[EKVirtualConferenceProvider](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKVirtualConferenceProvider) Init() EKVirtualConferenceProvider {
	rv := objc.Send[EKVirtualConferenceProvider](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKVirtualConferenceProvider) Autorelease() EKVirtualConferenceProvider {
	rv := objc.Send[EKVirtualConferenceProvider](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKVirtualConferenceProvider creates a new EKVirtualConferenceProvider instance.
func NewEKVirtualConferenceProvider() EKVirtualConferenceProvider {
	return getEKVirtualConferenceProviderClass().New()
}



// Provides an array of room types where events take place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceProvider/fetchAvailableRoomTypes(completionHandler:)
func (e_ EKVirtualConferenceProvider) FetchAvailableRoomTypesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("fetchAvailableRoomTypesWithCompletionHandler:"), completionHandler)
}


// Provides details about a virtual conference that takes place in a room the user selects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceProvider/fetchVirtualConference(identifier:completionHandler:)
func (e_ EKVirtualConferenceProvider) FetchVirtualConferenceForIdentifierCompletionHandler(identifier EKVirtualConferenceRoomTypeIdentifier, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("fetchVirtualConferenceForIdentifier:completionHandler:"), identifier, completionHandler)
}



