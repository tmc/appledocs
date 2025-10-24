// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKRecordZoneID */


/* debug [class_header]: Header for CKRecordZoneID */
// The class instance for the [CKRecordZoneID] class.
var (
	CKRecordZoneIDClass     _CKRecordZoneIDClass
	CKRecordZoneIDClassOnce sync.Once
)

func getCKRecordZoneIDClass() _CKRecordZoneIDClass {
	CKRecordZoneIDClassOnce.Do(func() {
		CKRecordZoneIDClass = _CKRecordZoneIDClass{objc.GetClass("CKRecordZoneID")}
	})
	return CKRecordZoneIDClass
}

type _CKRecordZoneIDClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKRecordZoneID */
// An interface definition for the [CKRecordZoneID] class.
type ICKRecordZoneID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKRecordZoneID */
	// properties:
	OwnerName() objc.IObject /* cross-framework: NSString */
	ZoneName() objc.IObject /* cross-framework: NSString */
	CKCurrentUserDefaultName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKRecordZoneID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKRecordZoneID */
// Alloc allocates a new instance without initialization.
func (cc _CKRecordZoneIDClass) Alloc() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKRecordZoneIDClass) New() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKRecordZoneID) Init() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKRecordZoneID) Autorelease() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecordZoneID creates a new CKRecordZoneID instance.
func NewCKRecordZoneID() CKRecordZoneID {
	return getCKRecordZoneIDClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKRecordZoneID */
// An object that uniquely identifies a record zone in a database.
//
// Zones are a mechanism for grouping related records together. You create zone ID objects when you want to fetch an existing zone object or create a new zone with a specific name. A record zone ID distinguishes one zone from another by a name string and the ID of the user who creates the zone. Both strings must be ASCII strings that don’t exceed 255 characters. When creating your own record zone ID objects, you can use names that have more meaning to your app or to the user, providing each zone name is unique within the specified database. The owner name must be either the current user name or the name of another user. Get the current user name from or by calling . When creating new record zones, make the name string in the record zone ID unique in the target database. Public databases don’t support custom zones, and only the user who owns the database can create zones in private databases. Don’t create subclasses of this class.


// An object that uniquely identifies a record zone in a database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/ID
type CKRecordZoneID struct {
	objectivec.Object
}

// CKRecordZoneIDFrom constructs a [CKRecordZoneID] from an unsafe.Pointer.
//
// An object that uniquely identifies a record zone in a database.
func CKRecordZoneIDFrom(ptr unsafe.Pointer) CKRecordZoneID {
	return CKRecordZoneID{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKRecordZoneID */

// Creates a record zone ID with the specified name and owner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneID/initWithZoneName:ownerName:
func NewCKRecordZoneIDWithZoneNameOwnerName(zoneName objc.IObject /* cross-framework: NSString */, ownerName objc.IObject /* cross-framework: NSString */) CKRecordZoneID {
	instance := getCKRecordZoneIDClass().Alloc()
	rv := objc.Send[CKRecordZoneID](instance.ID, objc.Sel("initWithZoneName:ownerName:"), zoneName, ownerName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKRecordZoneIDWithZoneNameOwnerName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKRecordZoneID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKRecordZoneID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKRecordZoneID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKRecordZoneID */

// The ID of the user who owns the record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/ID/ownerName
func (c_ CKRecordZoneID) OwnerName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("ownerName"))
	return rv
}/* debug [instance_properties/getter]: ownerName */


// The unique name of the record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/ID/zoneName
func (c_ CKRecordZoneID) ZoneName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("zoneName"))
	return rv
}/* debug [instance_properties/getter]: zoneName */


// A constant that provides the current user’s default name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcurrentuserdefaultname
func (c_ CKRecordZoneID) CKCurrentUserDefaultName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CKCurrentUserDefaultName"))
	return rv
}/* debug [instance_properties/getter]: CKCurrentUserDefaultName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKRecordZoneID */


