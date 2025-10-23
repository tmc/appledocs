// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CKSyncEngineState] class.
type ICKSyncEngineState interface {
	objectivec.IObject
	// properties:
	UserRecordID() objc.IObject /* cross-framework: CKRecordID */
	SetUserRecordID(value objc.IObject /* cross-framework: CKRecordID */)
	// methods:
	AddPendingDatabaseChanges(changes []CKSyncEnginePendingDatabaseChange /* primitive/slice/pointer. */)
	AddPendingRecordZoneChanges(changes []CKSyncEnginePendingRecordZoneChange /* primitive/slice/pointer. */)
}

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

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineStateClass) Alloc() CKSyncEngineState {
	rv := objc.Send[CKSyncEngineState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Adds the specified database changes to the state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/addPendingDatabaseChanges:
func (c_ CKSyncEngineState) AddPendingDatabaseChanges(changes []CKSyncEnginePendingDatabaseChange /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addPendingDatabaseChanges:"), changes)
}


// Adds the specified record zone changes to the state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/addPendingRecordZoneChanges:
func (c_ CKSyncEngineState) AddPendingRecordZoneChanges(changes []CKSyncEnginePendingRecordZoneChange /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addPendingRecordZoneChanges:"), changes)
}


// The user record ID for the corresponding user record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/userrecordid
func (c_ CKSyncEngineState) UserRecordID() objc.IObject /* cross-framework: CKRecordID */ {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("userRecordID"))
	return rv
}


// The user record ID for the corresponding user record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/userrecordid
func (c_ CKSyncEngineState) SetUserRecordID(value objc.IObject /* cross-framework: CKRecordID */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserRecordID:"), value)
}



