// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineState */


/* debug [class_header]: Header for CKSyncEngineState */
// The class instance for the [CKSyncEngineState] class.
var (
	CKSyncEngineStateClass     _CKSyncEngineStateClass
	CKSyncEngineStateClassOnce sync.Once
)

func getCKSyncEngineStateClass() _CKSyncEngineStateClass {
	CKSyncEngineStateClassOnce.Do(func() {
		CKSyncEngineStateClass = _CKSyncEngineStateClass{objc.GetClass("CKSyncEngineState")}
	})
	return CKSyncEngineStateClass
}

type _CKSyncEngineStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineState */
// An interface definition for the [CKSyncEngineState] class.
type ICKSyncEngineState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineState */
	// properties:
	HasPendingUntrackedChanges() bool
	SetHasPendingUntrackedChanges(value bool)
	PendingDatabaseChanges() []CKSyncEnginePendingDatabaseChange
	PendingRecordZoneChanges() []CKSyncEnginePendingRecordZoneChange
	ZoneIDsWithUnfetchedServerChanges() []CKRecordZoneID
	UserRecordID() ICKRecordID
	SetUserRecordID(value ICKRecordID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineState */
	// methods:
	AddPendingDatabaseChanges(changes []CKSyncEnginePendingDatabaseChange)
	AddPendingRecordZoneChanges(changes []CKSyncEnginePendingRecordZoneChange)
	RemovePendingDatabaseChanges(changes []CKSyncEnginePendingDatabaseChange)
	RemovePendingRecordZoneChanges(changes []CKSyncEnginePendingRecordZoneChange)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineState */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineStateClass) Alloc() CKSyncEngineState {
	rv := objc.Send[CKSyncEngineState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineStateClass) New() CKSyncEngineState {
	rv := objc.Send[CKSyncEngineState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineState) Init() CKSyncEngineState {
	rv := objc.Send[CKSyncEngineState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineState) Autorelease() CKSyncEngineState {
	rv := objc.Send[CKSyncEngineState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineState creates a new CKSyncEngineState instance.
func NewCKSyncEngineState() CKSyncEngineState {
	return getCKSyncEngineStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineState */
// An object that manages the sync engine’s state.
//
// To reliably and consistently sync your app’s data, a sync engine keeps a record of several important pieces of data, such as server changes tokens (for databases and record zones), subscription identifiers, the most recent , and so on. This class automatically manages that state on behalf of your app, but there are certain elements you can modify. For example, you control the list of pending changes to send to the iCloud servers and manipulate that list using the and methods. If there aren’t any scheduled sync operations when you invoke these methods, the engine automatically schedules one. An engine’s state changes periodically and, when it does, the sync engine dispatches an event of type to your delegate. The event contains an instance of and, on receipt of such an event, it’s your responsibility to persist the serialized state to disk so that it’s available across app launches. On the next initialization of the sync engine, you provide the most recently persisted state as part of the engine’s configuration. For more information, see .


// An object that manages the sync engine’s state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState
type CKSyncEngineState struct {
	objectivec.Object
}

// CKSyncEngineStateFrom constructs a [CKSyncEngineState] from an unsafe.Pointer.
//
// An object that manages the sync engine’s state.
func CKSyncEngineStateFrom(ptr unsafe.Pointer) CKSyncEngineState {
	return CKSyncEngineState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineState */

// Adds the specified database changes to the state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/addPendingDatabaseChanges:
func (c_ CKSyncEngineState) AddPendingDatabaseChanges(changes []CKSyncEnginePendingDatabaseChange) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addPendingDatabaseChanges:"), changes)
}/* debug [instance_methods/method]: AddPendingDatabaseChanges */


// Adds the specified record zone changes to the state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/addPendingRecordZoneChanges:
func (c_ CKSyncEngineState) AddPendingRecordZoneChanges(changes []CKSyncEnginePendingRecordZoneChange) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addPendingRecordZoneChanges:"), changes)
}/* debug [instance_methods/method]: AddPendingRecordZoneChanges */


// Removes the specified database changes from the state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/removePendingDatabaseChanges:
func (c_ CKSyncEngineState) RemovePendingDatabaseChanges(changes []CKSyncEnginePendingDatabaseChange) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removePendingDatabaseChanges:"), changes)
}/* debug [instance_methods/method]: RemovePendingDatabaseChanges */


// Removes the specified record zone changes from the state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/removePendingRecordZoneChanges:
func (c_ CKSyncEngineState) RemovePendingRecordZoneChanges(changes []CKSyncEnginePendingRecordZoneChange) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removePendingRecordZoneChanges:"), changes)
}/* debug [instance_methods/method]: RemovePendingRecordZoneChanges */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineState */

// A Boolean value that indicates whether there are pending changes that the sync engine is unaware of.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/hasPendingUntrackedChanges
func (c_ CKSyncEngineState) HasPendingUntrackedChanges() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasPendingUntrackedChanges"))
	return rv
}/* debug [instance_properties/getter]: hasPendingUntrackedChanges */


// A Boolean value that indicates whether there are pending changes that the sync engine is unaware of.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/hasPendingUntrackedChanges
func (c_ CKSyncEngineState) SetHasPendingUntrackedChanges(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasPendingUntrackedChanges:"), value)
}/* debug [instance_properties/setter]: hasPendingUntrackedChanges */


// The database changes that the sync engine has yet to send to the iCloud servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/pendingDatabaseChanges
func (c_ CKSyncEngineState) PendingDatabaseChanges() []CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[[]CKSyncEnginePendingDatabaseChange](c_.ID, objc.Sel("pendingDatabaseChanges"))
	return rv
}/* debug [instance_properties/getter]: pendingDatabaseChanges */


// The record zone changes that the sync engine has yet to send to the iCloud servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/pendingRecordZoneChanges
func (c_ CKSyncEngineState) PendingRecordZoneChanges() []CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[[]CKSyncEnginePendingRecordZoneChange](c_.ID, objc.Sel("pendingRecordZoneChanges"))
	return rv
}/* debug [instance_properties/getter]: pendingRecordZoneChanges */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/zoneIDsWithUnfetchedServerChanges
func (c_ CKSyncEngineState) ZoneIDsWithUnfetchedServerChanges() []CKRecordZoneID {
	rv := objc.Send[[]CKRecordZoneID](c_.ID, objc.Sel("zoneIDsWithUnfetchedServerChanges"))
	return rv
}/* debug [instance_properties/getter]: zoneIDsWithUnfetchedServerChanges */


// The user record ID for the corresponding user record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/userrecordid
func (c_ CKSyncEngineState) UserRecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("userRecordID"))
	return rv
}/* debug [instance_properties/getter]: userRecordID */


// The user record ID for the corresponding user record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/userrecordid
func (c_ CKSyncEngineState) SetUserRecordID(value ICKRecordID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserRecordID:"), value)
}/* debug [instance_properties/setter]: userRecordID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineState */



