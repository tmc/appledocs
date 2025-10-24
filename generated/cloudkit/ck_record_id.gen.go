// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKRecordID */


/* debug [class_header]: Header for CKRecordID */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKRecordID */
// An interface definition for the [CKRecordID] class.
type ICKRecordID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKRecordID */
	// properties:
	RecordName() objc.IObject /* cross-framework: NSString */
	ZoneID() ICKRecordZoneID
	CreationDate() foundation.Date
	SetCreationDate(value foundation.Date)
	CreatorUserRecordID() ICKRecordID
	SetCreatorUserRecordID(value ICKRecordID)
	LastModifiedUserRecordID() ICKRecordID
	SetLastModifiedUserRecordID(value ICKRecordID)
	ModificationDate() foundation.Date
	SetModificationDate(value foundation.Date)
	RecordChangeTag() objc.IObject /* cross-framework: NSString */
	SetRecordChangeTag(value objc.IObject /* cross-framework: NSString */)
	RecordID() ICKRecordID
	SetRecordID(value ICKRecordID)
	RecordType() objectivec.IObject
	SetRecordType(value objectivec.IObject)
	CKRecordNameZoneWideShare() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKRecordID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKRecordID */
// Alloc allocates a new instance without initialization.
func (cc _CKRecordIDClass) Alloc() CKRecordID {
	rv := objc.Send[CKRecordID](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKRecordID */
// An object that uniquely identifies a record in a database.
//
// A record ID object consists of a name string and a zone ID. The name string is an ASCII string that doesn’t exceed 255 characters in length. For automatically created records, the ID name string derives from a UUID and is, therefore, unique. When creating your own record ID objects, you can use names that have more meaning to your app or to the user, as long as each name is unique within the specified zone. For example, you might use a document name for the name string. Record IDs must be unique within the specified database, but you can reuse record IDs in different databases. Each container has a public and a private database, and the private database is different for each unique user. This configuration provides for the reusing of record IDs in each user’s private database, but ensures that only one record uses a specific record ID in the public database. CloudKit generally creates record IDs when it first saves a new record, but you might manually instantiate instances of in specific situations. For example, you must create an instance when saving a record in a zone other than the default zone. You also instantiate instances of when retrieving specific records from a database. Don’t subclass .


// An object that uniquely identifies a record in a database.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKRecordID */

// Creates a new record ID with the specified name in the default zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/ID/init(recordName:)
func NewCKRecordIDWithRecordName(recordName objc.IObject /* cross-framework: NSString */) CKRecordID {
	instance := getCKRecordIDClass().Alloc()
	rv := objc.Send[CKRecordID](instance.ID, objc.Sel("initWithRecordName:"), recordName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKRecordIDWithRecordName */


// Creates a new record ID with the specified name and zone information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordID/initWithRecordName:zoneID:
func NewCKRecordIDWithRecordNameZoneID(recordName objc.IObject /* cross-framework: NSString */, zoneID ICKRecordZoneID) CKRecordID {
	instance := getCKRecordIDClass().Alloc()
	rv := objc.Send[CKRecordID](instance.ID, objc.Sel("initWithRecordName:zoneID:"), recordName, zoneID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKRecordIDWithRecordNameZoneID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKRecordID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKRecordID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKRecordID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKRecordID */

// The unique name of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/ID/recordName
func (c_ CKRecordID) RecordName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("recordName"))
	return rv
}/* debug [instance_properties/getter]: recordName */


// The ID of the zone that contains the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/ID/zoneID
func (c_ CKRecordID) ZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}/* debug [instance_properties/getter]: zoneID */


// The time when CloudKit first saves the record to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/creationdate
func (c_ CKRecordID) CreationDate() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("creationDate"))
	return rv
}/* debug [instance_properties/getter]: creationDate */


// The time when CloudKit first saves the record to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/creationdate
func (c_ CKRecordID) SetCreationDate(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreationDate:"), value)
}/* debug [instance_properties/setter]: creationDate */


// The ID of the user who creates the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/creatoruserrecordid
func (c_ CKRecordID) CreatorUserRecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("creatorUserRecordID"))
	return rv
}/* debug [instance_properties/getter]: creatorUserRecordID */


// The ID of the user who creates the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/creatoruserrecordid
func (c_ CKRecordID) SetCreatorUserRecordID(value ICKRecordID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreatorUserRecordID:"), value)
}/* debug [instance_properties/setter]: creatorUserRecordID */


// The ID of the user who most recently modified the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/lastmodifieduserrecordid
func (c_ CKRecordID) LastModifiedUserRecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("lastModifiedUserRecordID"))
	return rv
}/* debug [instance_properties/getter]: lastModifiedUserRecordID */


// The ID of the user who most recently modified the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/lastmodifieduserrecordid
func (c_ CKRecordID) SetLastModifiedUserRecordID(value ICKRecordID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLastModifiedUserRecordID:"), value)
}/* debug [instance_properties/setter]: lastModifiedUserRecordID */


// The most recent time that CloudKit saved the record to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/modificationdate
func (c_ CKRecordID) ModificationDate() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("modificationDate"))
	return rv
}/* debug [instance_properties/getter]: modificationDate */


// The most recent time that CloudKit saved the record to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/modificationdate
func (c_ CKRecordID) SetModificationDate(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModificationDate:"), value)
}/* debug [instance_properties/setter]: modificationDate */


// The server change token for the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordchangetag
func (c_ CKRecordID) RecordChangeTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("recordChangeTag"))
	return rv
}/* debug [instance_properties/getter]: recordChangeTag */


// The server change token for the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordchangetag
func (c_ CKRecordID) SetRecordChangeTag(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordChangeTag:"), value)
}/* debug [instance_properties/setter]: recordChangeTag */


// The unique ID of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordid
func (c_ CKRecordID) RecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("recordID"))
	return rv
}/* debug [instance_properties/getter]: recordID */


// The unique ID of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordid
func (c_ CKRecordID) SetRecordID(value ICKRecordID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordID:"), value)
}/* debug [instance_properties/setter]: recordID */


// The value that your app defines to identify the type of record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordtype-6v7au
func (c_ CKRecordID) RecordType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("recordType"))
	return rv
}/* debug [instance_properties/getter]: recordType */


// The value that your app defines to identify the type of record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordtype-6v7au
func (c_ CKRecordID) SetRecordType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordType:"), value)
}/* debug [instance_properties/setter]: recordType */


// The name of a share record that manages a shared record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordnamezonewideshare
func (c_ CKRecordID) CKRecordNameZoneWideShare() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CKRecordNameZoneWideShare"))
	return rv
}/* debug [instance_properties/getter]: CKRecordNameZoneWideShare */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKRecordID */


