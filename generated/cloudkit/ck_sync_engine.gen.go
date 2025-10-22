// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEngine] class.
var (
	CKSyncEngineClass     _CKSyncEngineClass
	CKSyncEngineClassOnce sync.Once
)

func getCKSyncEngineClass() _CKSyncEngineClass {
	CKSyncEngineClassOnce.Do(func() {
		CKSyncEngineClass = _CKSyncEngineClass{objc.GetClass("CKSyncEngine")}
	})
	return CKSyncEngineClass
}

type _CKSyncEngineClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngine] class.
type ICKSyncEngine interface {
	objectivec.IObject
	CancelOperationsWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchChangesWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchChangesWithOptionsCompletionHandler(options CKSyncEngineFetchChangesOptions, completionHandler unsafe.Pointer)
	SendChangesWithCompletionHandler(completionHandler unsafe.Pointer)
	SendChangesWithOptionsCompletionHandler(options CKSyncEngineSendChangesOptions, completionHandler unsafe.Pointer)
	Database() CKDatabase
	State() CKSyncEngineState
}

// An object that manages the synchronization of local and remote record data.
//
// Use to handle your app’s CloudKit sync operations and benefit from the performance and reliability it provides. To use the class, create an instance early in your app’s launch process and specify a database to sync. Thereafter, and depending on good system conditions, the sync engine will periodically push and pull database and record zone changes on the app’s behalf. To participate in those sync operations and to provide the engine with the changes to send, create an object that conforms to and assign an instance of it to the engine’s configuration. You can have multiple instances of in a single process, each targeting a different database. For example, you may have one syncing a person’s private database and another syncing their shared database. Because periodic sync relies on good system conditions — adequate battery charge, an active network connection, a signed-in iCloud account, and so on — the engine’s sync schedule is indeterminate; if you need to sync immediately, like when you need to ensure your app has the most recent changes before continuing, use the and methods. The sync engine uses an opaque type to track its internal state, and it’s your responsibility to persist that state to disk and make it available across app launches so the engine can function properly. For more information, see and . requires the CloudKit and Remote notifications entitlements. For more information, see and .


// An object that manages the synchronization of local and remote record data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9

type CKSyncEngine struct {
	objectivec.Object
}

// CKSyncEngineFrom constructs a [CKSyncEngine] from an unsafe.Pointer.
//
// An object that manages the synchronization of local and remote record data.
func CKSyncEngineFrom(ptr unsafe.Pointer) CKSyncEngine {
	return CKSyncEngine{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineClass) Alloc() CKSyncEngine {
	rv := objc.Send[CKSyncEngine](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineClass) New() CKSyncEngine {
	rv := objc.Send[CKSyncEngine](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngine) Init() CKSyncEngine {
	rv := objc.Send[CKSyncEngine](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngine) Autorelease() CKSyncEngine {
	rv := objc.Send[CKSyncEngine](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngine creates a new CKSyncEngine instance.
func NewCKSyncEngine() CKSyncEngine {
	return getCKSyncEngineClass().New()
}




// Creates a sync engine with the specified configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/initWithConfiguration:

func NewCKSyncEngineWithConfiguration(configuration ICKSyncEngineConfiguration) CKSyncEngine {
	instance := getCKSyncEngineClass().Alloc()
	rv := objc.Send[CKSyncEngine](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}




// Cancels any in-progress or pending sync operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/cancelOperationsWithCompletionHandler:

func (c_ CKSyncEngine) CancelOperationsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("cancelOperationsWithCompletionHandler:"), completionHandler)
}



// Fetches pending remote changes from the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/fetchChangesWithCompletionHandler:

func (c_ CKSyncEngine) FetchChangesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchChangesWithCompletionHandler:"), completionHandler)
}



// Fetches pending remote changes from the server using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/fetchChangesWithOptions:completionHandler:

func (c_ CKSyncEngine) FetchChangesWithOptionsCompletionHandler(options CKSyncEngineFetchChangesOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchChangesWithOptions:completionHandler:"), options, completionHandler)
}



// Sends pending local changes to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/sendChangesWithCompletionHandler:

func (c_ CKSyncEngine) SendChangesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("sendChangesWithCompletionHandler:"), completionHandler)
}



// Sends pending local changes to the server using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/sendChangesWithOptions:completionHandler:

func (c_ CKSyncEngine) SendChangesWithOptionsCompletionHandler(options CKSyncEngineSendChangesOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("sendChangesWithOptions:completionHandler:"), options, completionHandler)
}


// The associated database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/database

func (c_ CKSyncEngine) Database() CKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("database"))
	return rv
}


// The sync engine’s state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/state

func (c_ CKSyncEngine) State() CKSyncEngineState {
	rv := objc.Send[CKSyncEngineState](c_.ID, objc.Sel("state"))
	return rv
}


