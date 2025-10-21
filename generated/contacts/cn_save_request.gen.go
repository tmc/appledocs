// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CNSaveRequest] class.
var (
	CNSaveRequestClass     _CNSaveRequestClass
	CNSaveRequestClassOnce sync.Once
)

func getCNSaveRequestClass() _CNSaveRequestClass {
	CNSaveRequestClassOnce.Do(func() {
		CNSaveRequestClass = _CNSaveRequestClass{objc.GetClass("CNSaveRequest")}
	})
	return CNSaveRequestClass
}

type _CNSaveRequestClass struct {
	class objc.Class
}

// An interface definition for the [CNSaveRequest] class.
type ICNSaveRequest interface {
	objectivec.IObject
	AddGroupToContainerWithIdentifier(group unsafe.Pointer, identifier string)
	AddContactToContainerWithIdentifier(contact unsafe.Pointer, identifier string)
	AddMemberToGroup(contact unsafe.Pointer, group unsafe.Pointer)
	AddSubgroupToGroup(subgroup unsafe.Pointer, group unsafe.Pointer)
	DeleteGroup(group unsafe.Pointer)
	DeleteContact(contact unsafe.Pointer)
	RemoveMemberFromGroup(contact unsafe.Pointer, group unsafe.Pointer)
	RemoveSubgroupFromGroup(subgroup unsafe.Pointer, group unsafe.Pointer)
	UpdateContact(contact unsafe.Pointer)
	UpdateGroup(group unsafe.Pointer)
}

// An object that collects the changes you want to save to the user’s contacts database.
//
// Create a new object for each save operation you want to make. You can batch multiple changes into one save request (note that these changes only apply to objects). In the case of overlapping changes in multiple or concurrent save requests, the last change wins. If you try to add an object (that is, a contact or a group) that already exists in the contact store, you receive the error and the array is updated to contain the object you tried to add. If you try to update or delete an object that is not present in the contact store, the save request does not perform the update or deletion, the error occurs, and the array is updated to contain the object you tried to update or delete. Do not access objects in the save request while that request is executing.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest
type CNSaveRequest struct {
	objectivec.Object
}

// CNSaveRequestFrom constructs a [CNSaveRequest] from an unsafe.Pointer.
//
// An object that collects the changes you want to save to the user’s contacts database.
func CNSaveRequestFrom(ptr unsafe.Pointer) CNSaveRequest {
	return CNSaveRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNSaveRequestClass) Alloc() CNSaveRequest {
	rv := objc.Send[CNSaveRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNSaveRequestClass) New() CNSaveRequest {
	rv := objc.Send[CNSaveRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNSaveRequest) Init() CNSaveRequest {
	rv := objc.Send[CNSaveRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNSaveRequest) Autorelease() CNSaveRequest {
	rv := objc.Send[CNSaveRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNSaveRequest creates a new CNSaveRequest instance.
func NewCNSaveRequest() CNSaveRequest {
	return getCNSaveRequestClass().New()
}


// Adds a group to the contact store.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/add(_:toContainerWithIdentifier:)-4ikaa
func (c_ CNSaveRequest) AddGroupToContainerWithIdentifier(group unsafe.Pointer, identifier string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addGroup:toContainerWithIdentifier:"), group, objc.String(identifier))
}

// Adds the specified contact to the contact store.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/add(_:toContainerWithIdentifier:)-7eut4
func (c_ CNSaveRequest) AddContactToContainerWithIdentifier(contact unsafe.Pointer, identifier string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addContact:toContainerWithIdentifier:"), contact, objc.String(identifier))
}

// Adds a contact as a member of a group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/addMember(_:to:)
func (c_ CNSaveRequest) AddMemberToGroup(contact unsafe.Pointer, group unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addMember:toGroup:"), contact, group)
}

// Add the specified group to a parent group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/addSubgroup(_:to:)
func (c_ CNSaveRequest) AddSubgroupToGroup(subgroup unsafe.Pointer, group unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addSubgroup:toGroup:"), subgroup, group)
}

// Deletes a group from the contact store.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/delete(_:)-29lsm
func (c_ CNSaveRequest) DeleteGroup(group unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteGroup:"), group)
}

// Deletes a contact from the contact store.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/delete(_:)-8m1tc
func (c_ CNSaveRequest) DeleteContact(contact unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteContact:"), contact)
}

// Removes a contact as a member of a group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/removeMember(_:from:)
func (c_ CNSaveRequest) RemoveMemberFromGroup(contact unsafe.Pointer, group unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeMember:fromGroup:"), contact, group)
}

// Remove a subgroup from the specified parent group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/removeSubgroup(_:from:)
func (c_ CNSaveRequest) RemoveSubgroupFromGroup(subgroup unsafe.Pointer, group unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeSubgroup:fromGroup:"), subgroup, group)
}

// Updates an existing contact in the contact store.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/update(_:)-3gaig
func (c_ CNSaveRequest) UpdateContact(contact unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("updateContact:"), contact)
}

// Updates an existing group in the contact store.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/update(_:)-8h6f6
func (c_ CNSaveRequest) UpdateGroup(group unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("updateGroup:"), group)
}

// A Boolean value that indicates whether to refetch the added and updated contacts after the save request executes.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/shouldRefetchContacts
func (c_ CNSaveRequest) ShouldRefetchContacts() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldRefetchContacts"))
	return rv
}


// SetShouldRefetchContacts sets the value of the shouldRefetchContacts property.
// A Boolean value that indicates whether to refetch the added and updated contacts after the save request executes.

//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/shouldRefetchContacts
func (c_ CNSaveRequest) SetShouldRefetchContacts(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldRefetchContacts:"), value)
}

// A string that identifies the author of the transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/transactionAuthor
func (c_ CNSaveRequest) TransactionAuthor() string {
	rv := objc.Send[string](c_.ID, objc.Sel("transactionAuthor"))
	return rv
}


// SetTransactionAuthor sets the value of the transactionAuthor property.
// A string that identifies the author of the transaction.

//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/transactionAuthor
func (c_ CNSaveRequest) SetTransactionAuthor(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTransactionAuthor:"), objc.String(value))
}



