// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNSaveRequest */


/* debug [class_header]: Header for CNSaveRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNSaveRequest */
// An interface definition for the [CNSaveRequest] class.
type ICNSaveRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNSaveRequest */
	// properties:
	ShouldRefetchContacts() bool
	SetShouldRefetchContacts(value bool)
	TransactionAuthor() objc.IObject /* cross-framework: NSString */
	SetTransactionAuthor(value objc.IObject /* cross-framework: NSString */)
	CNErrorUserInfoAffectedRecordsKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNSaveRequest */
	// methods:
	AddGroupToContainerWithIdentifier(group ICNMutableGroup, identifier objc.IObject /* cross-framework: NSString */)
	AddContactToContainerWithIdentifier(contact ICNMutableContact, identifier objc.IObject /* cross-framework: NSString */)
	AddMemberToGroup(contact ICNContact, group ICNGroup)
	AddSubgroupToGroup(subgroup ICNGroup, group ICNGroup)
	DeleteGroup(group ICNMutableGroup)
	DeleteContact(contact ICNMutableContact)
	RemoveMemberFromGroup(contact ICNContact, group ICNGroup)
	RemoveSubgroupFromGroup(subgroup ICNGroup, group ICNGroup)
	UpdateContact(contact ICNMutableContact)
	UpdateGroup(group ICNMutableGroup)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNSaveRequest */
// Alloc allocates a new instance without initialization.
func (cc _CNSaveRequestClass) Alloc() CNSaveRequest {
	rv := objc.Send[CNSaveRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNSaveRequest */
// An object that collects the changes you want to save to the user’s contacts database.
//
// Create a new object for each save operation you want to make. You can batch multiple changes into one save request (note that these changes only apply to objects). In the case of overlapping changes in multiple or concurrent save requests, the last change wins. If you try to add an object (that is, a contact or a group) that already exists in the contact store, you receive the error and the array is updated to contain the object you tried to add. If you try to update or delete an object that is not present in the contact store, the save request does not perform the update or deletion, the error occurs, and the array is updated to contain the object you tried to update or delete. Do not access objects in the save request while that request is executing.


// An object that collects the changes you want to save to the user’s contacts database.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNSaveRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNSaveRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNSaveRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNSaveRequest */

// Adds a group to the contact store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/add(_:toContainerWithIdentifier:)-4ikaa
func (c_ CNSaveRequest) AddGroupToContainerWithIdentifier(group ICNMutableGroup, identifier objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addGroup:toContainerWithIdentifier:"), group, identifier)
}/* debug [instance_methods/method]: AddGroupToContainerWithIdentifier */


// Adds the specified contact to the contact store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/add(_:toContainerWithIdentifier:)-7eut4
func (c_ CNSaveRequest) AddContactToContainerWithIdentifier(contact ICNMutableContact, identifier objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addContact:toContainerWithIdentifier:"), contact, identifier)
}/* debug [instance_methods/method]: AddContactToContainerWithIdentifier */


// Adds a contact as a member of a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/addMember(_:to:)
func (c_ CNSaveRequest) AddMemberToGroup(contact ICNContact, group ICNGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addMember:toGroup:"), contact, group)
}/* debug [instance_methods/method]: AddMemberToGroup */


// Add the specified group to a parent group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/addSubgroup(_:to:)
func (c_ CNSaveRequest) AddSubgroupToGroup(subgroup ICNGroup, group ICNGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addSubgroup:toGroup:"), subgroup, group)
}/* debug [instance_methods/method]: AddSubgroupToGroup */


// Deletes a group from the contact store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/delete(_:)-29lsm
func (c_ CNSaveRequest) DeleteGroup(group ICNMutableGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteGroup:"), group)
}/* debug [instance_methods/method]: DeleteGroup */


// Deletes a contact from the contact store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/delete(_:)-8m1tc
func (c_ CNSaveRequest) DeleteContact(contact ICNMutableContact) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteContact:"), contact)
}/* debug [instance_methods/method]: DeleteContact */


// Removes a contact as a member of a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/removeMember(_:from:)
func (c_ CNSaveRequest) RemoveMemberFromGroup(contact ICNContact, group ICNGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeMember:fromGroup:"), contact, group)
}/* debug [instance_methods/method]: RemoveMemberFromGroup */


// Remove a subgroup from the specified parent group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/removeSubgroup(_:from:)
func (c_ CNSaveRequest) RemoveSubgroupFromGroup(subgroup ICNGroup, group ICNGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeSubgroup:fromGroup:"), subgroup, group)
}/* debug [instance_methods/method]: RemoveSubgroupFromGroup */


// Updates an existing contact in the contact store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/update(_:)-3gaig
func (c_ CNSaveRequest) UpdateContact(contact ICNMutableContact) {
	objc.Send[objc.ID](c_.ID, objc.Sel("updateContact:"), contact)
}/* debug [instance_methods/method]: UpdateContact */


// Updates an existing group in the contact store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/update(_:)-8h6f6
func (c_ CNSaveRequest) UpdateGroup(group ICNMutableGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("updateGroup:"), group)
}/* debug [instance_methods/method]: UpdateGroup */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNSaveRequest */

// A Boolean value that indicates whether to refetch the added and updated contacts after the save request executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/shouldRefetchContacts
func (c_ CNSaveRequest) ShouldRefetchContacts() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldRefetchContacts"))
	return rv
}/* debug [instance_properties/getter]: shouldRefetchContacts */


// A Boolean value that indicates whether to refetch the added and updated contacts after the save request executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/shouldRefetchContacts
func (c_ CNSaveRequest) SetShouldRefetchContacts(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldRefetchContacts:"), value)
}/* debug [instance_properties/setter]: shouldRefetchContacts */


// A string that identifies the author of the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/transactionAuthor
func (c_ CNSaveRequest) TransactionAuthor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("transactionAuthor"))
	return rv
}/* debug [instance_properties/getter]: transactionAuthor */


// A string that identifies the author of the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest/transactionAuthor
func (c_ CNSaveRequest) SetTransactionAuthor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTransactionAuthor:"), value)
}/* debug [instance_properties/setter]: transactionAuthor */


// The contact, group, and container objects for which the error code applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnerroruserinfoaffectedrecordskey
func (c_ CNSaveRequest) CNErrorUserInfoAffectedRecordsKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNErrorUserInfoAffectedRecordsKey"))
	return rv
}/* debug [instance_properties/getter]: CNErrorUserInfoAffectedRecordsKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNSaveRequest */



