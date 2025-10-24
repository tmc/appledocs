// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineConfiguration */


/* debug [class_header]: Header for CKSyncEngineConfiguration */
// The class instance for the [CKSyncEngineConfiguration] class.
var (
	CKSyncEngineConfigurationClass     _CKSyncEngineConfigurationClass
	CKSyncEngineConfigurationClassOnce sync.Once
)

func getCKSyncEngineConfigurationClass() _CKSyncEngineConfigurationClass {
	CKSyncEngineConfigurationClassOnce.Do(func() {
		CKSyncEngineConfigurationClass = _CKSyncEngineConfigurationClass{objc.GetClass("CKSyncEngineConfiguration")}
	})
	return CKSyncEngineConfigurationClass
}

type _CKSyncEngineConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineConfiguration */
// An interface definition for the [CKSyncEngineConfiguration] class.
type ICKSyncEngineConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineConfiguration */
	// properties:
	AutomaticallySync() bool
	SetAutomaticallySync(value bool)
	Database() ICKDatabase
	SetDatabase(value ICKDatabase)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	StateSerialization() ICKSyncEngineStateSerialization
	SetStateSerialization(value ICKSyncEngineStateSerialization)
	SubscriptionID() objectivec.IObject
	SetSubscriptionID(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineConfiguration */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineConfigurationClass) Alloc() CKSyncEngineConfiguration {
	rv := objc.Send[CKSyncEngineConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineConfigurationClass) New() CKSyncEngineConfiguration {
	rv := objc.Send[CKSyncEngineConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineConfiguration) Init() CKSyncEngineConfiguration {
	rv := objc.Send[CKSyncEngineConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineConfiguration) Autorelease() CKSyncEngineConfiguration {
	rv := objc.Send[CKSyncEngineConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineConfiguration creates a new CKSyncEngineConfiguration instance.
func NewCKSyncEngineConfiguration() CKSyncEngineConfiguration {
	return getCKSyncEngineConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineConfiguration */
// A type that configures the attributes and behavior of a sync engine.


// A type that configures the attributes and behavior of a sync engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration
type CKSyncEngineConfiguration struct {
	objectivec.Object
}

// CKSyncEngineConfigurationFrom constructs a [CKSyncEngineConfiguration] from an unsafe.Pointer.
//
// A type that configures the attributes and behavior of a sync engine.
func CKSyncEngineConfigurationFrom(ptr unsafe.Pointer) CKSyncEngineConfiguration {
	return CKSyncEngineConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineConfiguration */

// Creates a configuration for the specified database and serialized state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/initWithDatabase:stateSerialization:delegate:
func NewCKSyncEngineConfigurationWithDatabaseStateSerializationDelegate(database ICKDatabase, stateSerialization ICKSyncEngineStateSerialization, delegate unsafe.Pointer) CKSyncEngineConfiguration {
	instance := getCKSyncEngineConfigurationClass().Alloc()
	rv := objc.Send[CKSyncEngineConfiguration](instance.ID, objc.Sel("initWithDatabase:stateSerialization:delegate:"), database, stateSerialization, delegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEngineConfigurationWithDatabaseStateSerializationDelegate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineConfiguration */

// A Boolean value that determines whether the engine syncs automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/automaticallySync
func (c_ CKSyncEngineConfiguration) AutomaticallySync() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallySync"))
	return rv
}/* debug [instance_properties/getter]: automaticallySync */


// A Boolean value that determines whether the engine syncs automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/automaticallySync
func (c_ CKSyncEngineConfiguration) SetAutomaticallySync(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallySync:"), value)
}/* debug [instance_properties/setter]: automaticallySync */


// The associated database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/database
func (c_ CKSyncEngineConfiguration) Database() ICKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("database"))
	return rv
}/* debug [instance_properties/getter]: database */


// The associated database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/database
func (c_ CKSyncEngineConfiguration) SetDatabase(value ICKDatabase) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDatabase:"), value)
}/* debug [instance_properties/setter]: database */


// The object that provides the records to sync and handles any related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/delegate
func (c_ CKSyncEngineConfiguration) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The object that provides the records to sync and handles any related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/delegate
func (c_ CKSyncEngineConfiguration) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The sync engine’s serialized state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/stateSerialization
func (c_ CKSyncEngineConfiguration) StateSerialization() ICKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](c_.ID, objc.Sel("stateSerialization"))
	return rv
}/* debug [instance_properties/getter]: stateSerialization */


// The sync engine’s serialized state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/stateSerialization
func (c_ CKSyncEngineConfiguration) SetStateSerialization(value ICKSyncEngineStateSerialization) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStateSerialization:"), value)
}/* debug [instance_properties/setter]: stateSerialization */


// The subscription identifier for the associated database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/subscriptionID
func (c_ CKSyncEngineConfiguration) SubscriptionID() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("subscriptionID"))
	return rv
}/* debug [instance_properties/getter]: subscriptionID */


// The subscription identifier for the associated database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/subscriptionID
func (c_ CKSyncEngineConfiguration) SetSubscriptionID(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubscriptionID:"), value)
}/* debug [instance_properties/setter]: subscriptionID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineConfiguration */


