// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKRecordID] class.
var (
	CKRecordIDClass     _CKRecordIDClass
	CKRecordIDClassOnce sync.Once
)

func getCKRecordIDClass() _CKRecordIDClass {
	CKRecordIDClassOnce.Do(func() {
		CKRecordIDClass = _CKRecordIDClass{objc.GetClass("CKRecordID")}
	})
	return CKRecordIDClass
}

type _CKRecordIDClass struct {
	class objc.Class
}

// An interface definition for the [CKRecordID] class.
type ICKRecordID interface {
	objectivec.IObject
}

// An object that uniquely identifies a record in a database.
//
// A record ID object consists of a name string and a zone ID. The name string is an ASCII string that doesn’t exceed 255 characters in length. For automatically created records, the ID name string derives from a UUID and is, therefore, unique. When creating your own record ID objects, you can use names that have more meaning to your app or to the user, as long as each name is unique within the specified zone. For example, you might use a document name for the name string. Record IDs must be unique within the specified database, but you can reuse record IDs in different databases. Each container has a public and a private database, and the private database is different for each unique user. This configuration provides for the reusing of record IDs in each user’s private database, but ensures that only one record uses a specific record ID in the public database. CloudKit generally creates record IDs when it first saves a new record, but you might manually instantiate instances of in specific situations. For example, you must create an instance when saving a record in a zone other than the default zone. You also instantiate instances of when retrieving specific records from a database. Don’t subclass .
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/ID
type CKRecordID struct {
	objectivec.Object
}

// CKRecordIDFrom constructs a [CKRecordID] from an unsafe.Pointer.
//
// An object that uniquely identifies a record in a database.
func CKRecordIDFrom(ptr unsafe.Pointer) CKRecordID {
	return CKRecordID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKRecordIDClass) Alloc() CKRecordID {
	rv := objc.Send[CKRecordID](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKRecordIDClass) New() CKRecordID {
	rv := objc.Send[CKRecordID](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKRecordID) Init() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKRecordID) Autorelease() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecordID creates a new CKRecordID instance.
func NewCKRecordID() CKRecordID {
	return getCKRecordIDClass().New()
}


// Creates a new record ID with the specified name in the default zone.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/ID/init(recordName:)
func NewCKRecordIDWithRecordName(recordName string) CKRecordID {
	instance := getCKRecordIDClass().Alloc()
	rv := objc.Send[CKRecordID](instance.ID, objc.Sel("initWithRecordName:"), objc.String(recordName))
	rv.Autorelease()
	return rv
}

// Creates a new record ID with the specified name and zone information.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordID/initWithRecordName:zoneID:
func NewCKRecordIDWithRecordNameZoneID(recordName string, zoneID unsafe.Pointer) CKRecordID {
	instance := getCKRecordIDClass().Alloc()
	rv := objc.Send[CKRecordID](instance.ID, objc.Sel("initWithRecordName:zoneID:"), objc.String(recordName), zoneID)
	rv.Autorelease()
	return rv
}


// The unique name of the record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/ID/recordName
func (c_ CKRecordID) RecordName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordName"))
	return rv
}

// The ID of the zone that contains the record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/ID/zoneID
func (c_ CKRecordID) ZoneID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("zoneID"))
	return rv
}


