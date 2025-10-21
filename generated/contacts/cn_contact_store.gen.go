// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CNContactStore] class.
var (
	CNContactStoreClass     _CNContactStoreClass
	CNContactStoreClassOnce sync.Once
)

func getCNContactStoreClass() _CNContactStoreClass {
	CNContactStoreClassOnce.Do(func() {
		CNContactStoreClass = _CNContactStoreClass{objc.GetClass("CNContactStore")}
	})
	return CNContactStoreClass
}

type _CNContactStoreClass struct {
	class objc.Class
}

// An interface definition for the [CNContactStore] class.
type ICNContactStore interface {
	objectivec.IObject
	ContainersMatchingPredicateError(predicate unsafe.Pointer, error_ unsafe.Pointer) []CNContainer
	DefaultContainerIdentifier() string
	EnumerateContactsWithFetchRequestErrorUsingBlock(fetchRequest unsafe.Pointer, error_ unsafe.Pointer, block unsafe.Pointer) bool
	EnumeratorForChangeHistoryFetchRequestError(request unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	EnumeratorForContactFetchRequestError(request unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	ExecuteSaveRequestError(saveRequest unsafe.Pointer, error_ unsafe.Pointer) bool
	GroupsMatchingPredicateError(predicate unsafe.Pointer, error_ unsafe.Pointer) []CNGroup
	RequestAccessForEntityTypeCompletionHandler(entityType unsafe.Pointer, completionHandler unsafe.Pointer)
	UnifiedContactWithIdentifierKeysToFetchError(identifier string, keys unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	UnifiedContactsMatchingPredicateKeysToFetchError(predicate unsafe.Pointer, keys unsafe.Pointer, error_ unsafe.Pointer) []CNContact
	UnifiedMeContactWithKeysToFetchError(keys unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
}

// The object that fetches and saves contacts, groups, and containers from the user’s Contacts database.
//
// The object represents the user’s contacts store database, and you use it to fetch information from that database and save changes back to it. There are a few recommended ways you can implement fetch and save requests in your app: Fetch only the properties that you need for contacts. When fetching all contacts and caching the results, first fetch all contacts identifiers, then fetch batches of detailed contacts by identifiers as required. To aggregate several contacts fetches, first collect a set of unique identifiers from the fetches. Then fetch batches of detailed contacts by those unique identifiers. If you cache the fetched contacts, groups, or containers, you need to refetch these objects (and release the old cached objects) when is posted. Because fetch methods perform I/O, it’s recommended that you avoid using the main thread to execute fetches.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore
type CNContactStore struct {
	objectivec.Object
}

// CNContactStoreFrom constructs a [CNContactStore] from an unsafe.Pointer.
//
// The object that fetches and saves contacts, groups, and containers from the user’s Contacts database.
func CNContactStoreFrom(ptr unsafe.Pointer) CNContactStore {
	return CNContactStore{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNContactStoreClass) Alloc() CNContactStore {
	rv := objc.Send[CNContactStore](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNContactStoreClass) New() CNContactStore {
	rv := objc.Send[CNContactStore](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContactStore) Init() CNContactStore {
	rv := objc.Send[CNContactStore](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContactStore) Autorelease() CNContactStore {
	rv := objc.Send[CNContactStore](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContactStore creates a new CNContactStore instance.
func NewCNContactStore() CNContactStore {
	return getCNContactStoreClass().New()
}


// Returns the current authorization status to access the contact data.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/authorizationStatus(for:)
func (cc _CNContactStoreClass) AuthorizationStatusForEntityType(entityType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("authorizationStatusForEntityType:"), entityType)
	return rv
}

// Fetches all containers matching the specified predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/containers(matching:)
func (c_ CNContactStore) ContainersMatchingPredicateError(predicate unsafe.Pointer, error_ unsafe.Pointer) []CNContainer {
	rv := objc.Send[[]CNContainer](c_.ID, objc.Sel("containersMatchingPredicate:error:"), predicate, error_)
	return rv
}

// Returns the identifier of the default container.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/defaultContainerIdentifier()
func (c_ CNContactStore) DefaultContainerIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("defaultContainerIdentifier"))
	return rv
}

// Returns a Boolean value that indicates whether the enumeration of all contacts matching a contact fetch request executes successfully.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/enumerateContacts(with:usingBlock:)
func (c_ CNContactStore) EnumerateContactsWithFetchRequestErrorUsingBlock(fetchRequest unsafe.Pointer, error_ unsafe.Pointer, block unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enumerateContactsWithFetchRequest:error:usingBlock:"), fetchRequest, error_, block)
	return rv
}

// Enumerates a change history fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/enumeratorForChangeHistoryFetchRequest:error:
func (c_ CNContactStore) EnumeratorForChangeHistoryFetchRequestError(request unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("enumeratorForChangeHistoryFetchRequest:error:"), request, error_)
	return rv
}

// Enumerates a contact fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/enumeratorForContactFetchRequest:error:
func (c_ CNContactStore) EnumeratorForContactFetchRequestError(request unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("enumeratorForContactFetchRequest:error:"), request, error_)
	return rv
}

// Executes a save request and returns success or failure.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/execute(_:)
func (c_ CNContactStore) ExecuteSaveRequestError(saveRequest unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("executeSaveRequest:error:"), saveRequest, error_)
	return rv
}

// Fetches all groups matching the specified predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/groups(matching:)
func (c_ CNContactStore) GroupsMatchingPredicateError(predicate unsafe.Pointer, error_ unsafe.Pointer) []CNGroup {
	rv := objc.Send[[]CNGroup](c_.ID, objc.Sel("groupsMatchingPredicate:error:"), predicate, error_)
	return rv
}

// Requests access to the user’s contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/requestAccess(for:completionHandler:)
func (c_ CNContactStore) RequestAccessForEntityTypeCompletionHandler(entityType unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("requestAccessForEntityType:completionHandler:"), entityType, completionHandler)
}

// Fetches a unified contact for the specified contact identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/unifiedContact(withIdentifier:keysToFetch:)
func (c_ CNContactStore) UnifiedContactWithIdentifierKeysToFetchError(identifier string, keys unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("unifiedContactWithIdentifier:keysToFetch:error:"), objc.String(identifier), keys, error_)
	return rv
}

// Fetches all unified contacts matching the specified predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/unifiedContacts(matching:keysToFetch:)
func (c_ CNContactStore) UnifiedContactsMatchingPredicateKeysToFetchError(predicate unsafe.Pointer, keys unsafe.Pointer, error_ unsafe.Pointer) []CNContact {
	rv := objc.Send[[]CNContact](c_.ID, objc.Sel("unifiedContactsMatchingPredicate:keysToFetch:error:"), predicate, keys, error_)
	return rv
}

// Fetches the unified contact that’s the card.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/unifiedMeContactWithKeys(toFetch:)
func (c_ CNContactStore) UnifiedMeContactWithKeysToFetchError(keys unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("unifiedMeContactWithKeysToFetch:error:"), keys, error_)
	return rv
}

// The current history token.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactStore/currentHistoryToken
func (c_ CNContactStore) CurrentHistoryToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("currentHistoryToken"))
	return rv
}



