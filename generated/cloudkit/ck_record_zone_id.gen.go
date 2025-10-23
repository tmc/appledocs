// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CKRecordZoneID] class.
type ICKRecordZoneID interface {
	objectivec.IObject
	CKCurrentUserDefaultName() string
	OwnerName() string
	SetOwnerName(value string)
	ZoneName() string
	SetZoneName(value string)
}

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

// Alloc allocates a new instance without initialization.
func (cc _CKRecordZoneIDClass) Alloc() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A constant that provides the current user’s default name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckcurrentuserdefaultname
func (c_ CKRecordZoneID) CKCurrentUserDefaultName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CKCurrentUserDefaultName"))
	return rv
}


// The ID of the user who owns the record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzone/id/ownername
func (c_ CKRecordZoneID) OwnerName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("ownerName"))
	return rv
}


// The ID of the user who owns the record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzone/id/ownername
func (c_ CKRecordZoneID) SetOwnerName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOwnerName:"), objc.String(value))
}


// The unique name of the record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzone/id/zonename
func (c_ CKRecordZoneID) ZoneName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("zoneName"))
	return rv
}


// The unique name of the record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzone/id/zonename
func (c_ CKRecordZoneID) SetZoneName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZoneName:"), objc.String(value))
}



