// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKRecordZone] class.
var (
	CKRecordZoneClass     _CKRecordZoneClass
	CKRecordZoneClassOnce sync.Once
)

func getCKRecordZoneClass() _CKRecordZoneClass {
	CKRecordZoneClassOnce.Do(func() {
		CKRecordZoneClass = _CKRecordZoneClass{objc.GetClass("CKRecordZone")}
	})
	return CKRecordZoneClass
}

type _CKRecordZoneClass struct {
	class objc.Class
}

// An interface definition for the [CKRecordZone] class.
type ICKRecordZone interface {
	objectivec.IObject
	Capabilities() CKRecordZoneCapabilities
	EncryptionScope() CKRecordZoneEncryptionScope
	SetEncryptionScope(value ICKRecordZoneEncryptionScope)
	Share() CKReference
	ZoneID() CKRecordZoneID
}

// A database partition that contains related records.
//
// Zones are an important part of how you organize your data. The public and private databases each have a single default zone. In the private database, you can use objects to create additional custom zones as necessary. Use custom zones to arrange and encapsulate groups of related records in the private database. Custom zones support other capabilities too, such as the ability to write multiple records as a single atomic transaction. Treat each custom zone as a single unit of data that is separate from every other zone in the database. Inside the zone, you add records as you would anywhere else. You can also create links between the records inside a zone by using the class. However, the class doesn’t support cross-zone linking, so each reference object must point to a record in the same zone as the current record. Use the class as-is and don’t subclass it.


// A database partition that contains related records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone
type CKRecordZone struct {
	objectivec.Object
}

// CKRecordZoneFrom constructs a [CKRecordZone] from an unsafe.Pointer.
//
// A database partition that contains related records.
func CKRecordZoneFrom(ptr unsafe.Pointer) CKRecordZone {
	return CKRecordZone{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKRecordZoneClass) Alloc() CKRecordZone {
	rv := objc.Send[CKRecordZone](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKRecordZoneClass) New() CKRecordZone {
	rv := objc.Send[CKRecordZone](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKRecordZone) Init() CKRecordZone {
	rv := objc.Send[CKRecordZone](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKRecordZone) Autorelease() CKRecordZone {
	rv := objc.Send[CKRecordZone](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecordZone creates a new CKRecordZone instance.
func NewCKRecordZone() CKRecordZone {
	return getCKRecordZoneClass().New()
}



// Creates a record zone object with the specified zone ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/init(zoneID:)
func NewCKRecordZoneWithZoneID(zoneID ICKRecordZoneID) CKRecordZone {
	instance := getCKRecordZoneClass().Alloc()
	rv := objc.Send[CKRecordZone](instance.ID, objc.Sel("initWithZoneID:"), zoneID)
	rv.Autorelease()
	return rv
}


// Creates a record zone object with the specified zone name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/init(zoneName:)
func NewCKRecordZoneWithZoneName(zoneName string) CKRecordZone {
	instance := getCKRecordZoneClass().Alloc()
	rv := objc.Send[CKRecordZone](instance.ID, objc.Sel("initWithZoneName:"), objc.String(zoneName))
	rv.Autorelease()
	return rv
}



// Returns the default record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/default()
func (cc _CKRecordZoneClass) DefaultRecordZone() CKRecordZone {
	rv := objc.Send[CKRecordZone](objc.ID(cc.class), objc.Sel("defaultRecordZone"))
	return rv
}


// The capabilities that the zone supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/capabilities-swift.property
func (c_ CKRecordZone) Capabilities() CKRecordZoneCapabilities {
	rv := objc.Send[CKRecordZoneCapabilities](c_.ID, objc.Sel("capabilities"))
	return rv
}


// The encryption scope determines the granularity at which encryption keys are stored within the zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/encryptionScope-swift.property
func (c_ CKRecordZone) EncryptionScope() CKRecordZoneEncryptionScope {
	rv := objc.Send[CKRecordZoneEncryptionScope](c_.ID, objc.Sel("encryptionScope"))
	return rv
}


// The encryption scope determines the granularity at which encryption keys are stored within the zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/encryptionScope-swift.property
func (c_ CKRecordZone) SetEncryptionScope(value ICKRecordZoneEncryptionScope) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEncryptionScope:"), value)
}


// A reference to the record zone’s share record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/share
func (c_ CKRecordZone) Share() CKReference {
	rv := objc.Send[CKReference](c_.ID, objc.Sel("share"))
	return rv
}


// The unique ID of the zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/zoneID
func (c_ CKRecordZone) ZoneID() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}


