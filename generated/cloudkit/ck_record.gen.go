// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKRecord] class.
var (
	CKRecordClass     _CKRecordClass
	CKRecordClassOnce sync.Once
)

func getCKRecordClass() _CKRecordClass {
	CKRecordClassOnce.Do(func() {
		CKRecordClass = _CKRecordClass{objc.GetClass("CKRecord")}
	})
	return CKRecordClass
}

type _CKRecordClass struct {
	class objc.Class
}

// An interface definition for the [CKRecord] class.
type ICKRecord interface {
	objectivec.IObject
	// properties:
	RecordChangeTag() objc.IObject /* cross-framework: NSString */
	CreationDate() objc.IObject /* cross-framework: Date */
	SetCreationDate(value objc.IObject /* cross-framework: Date */)
	CreatorUserRecordID() objc.IObject /* cross-framework: CKRecordID */
	SetCreatorUserRecordID(value objc.IObject /* cross-framework: CKRecordID */)
	EncryptedValues() unsafe.Pointer
	SetEncryptedValues(value unsafe.Pointer)
	LastModifiedUserRecordID() objc.IObject /* cross-framework: CKRecordID */
	SetLastModifiedUserRecordID(value objc.IObject /* cross-framework: CKRecordID */)
	ModificationDate() objc.IObject /* cross-framework: Date */
	SetModificationDate(value objc.IObject /* cross-framework: Date */)
	Parent() ICKReference
	SetParent(value ICKReference)
	RecordID() objc.IObject /* cross-framework: CKRecordID */
	SetRecordID(value objc.IObject /* cross-framework: CKRecordID */)
	RecordType() unsafe.Pointer
	SetRecordType(value unsafe.Pointer)
	Share() ICKReference
	SetShare(value ICKReference)
	// methods:
	EncodeSystemFieldsWithCoder(coder objc.IObject /* cross-framework: Coder */)
}

// A collection of key-value pairs that store your app’s data.
//
// Records are the fundamental objects that manage data in CloudKit. You can define any number of record types for your app, with each record type corresponding to a different type of information. Within a record type, you then define one or more fields, each with a name and a value. Records can contain simple data types, such as strings and numbers, or more complex types, such as geographic locations or pointers to other records. An important step in using CloudKit is defining the record types your app supports. A new record object doesn’t contain any keys or values. During development, you can add new keys and values at any time. The first time you set a value for a key and save the record, the server associates that type with the key for all records of the same type. The class doesn’t impose these type constraints or do any local validation of a record’s contents. CloudKit enforces these constraints when you save the records. Although records behave like dictionaries, there are limitations to the types of values you can assign to keys. The following are the object types that the class supports. Attempting to specify objects of any other type results in failure. Fields of all types are searchable unless otherwise noted.


// A collection of key-value pairs that store your app’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord
type CKRecord struct {
	objectivec.Object
}

// CKRecordFrom constructs a [CKRecord] from an unsafe.Pointer.
//
// A collection of key-value pairs that store your app’s data.
func CKRecordFrom(ptr unsafe.Pointer) CKRecord {
	return CKRecord{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKRecordClass) Alloc() CKRecord {
	rv := objc.Send[CKRecord](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKRecordClass) New() CKRecord {
	rv := objc.Send[CKRecord](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKRecord) Init() CKRecord {
	rv := objc.Send[CKRecord](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKRecord) Autorelease() CKRecord {
	rv := objc.Send[CKRecord](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecord creates a new CKRecord instance.
func NewCKRecord() CKRecord {
	return getCKRecordClass().New()
}



// Encodes the record’s system fields using the specified archiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/encodeSystemFields(with:)
func (c_ CKRecord) EncodeSystemFieldsWithCoder(coder objc.IObject /* cross-framework: Coder */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeSystemFieldsWithCoder:"), coder)
}


// The server change token for the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/recordChangeTag
func (c_ CKRecord) RecordChangeTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("recordChangeTag"))
	return rv
}


// The time when CloudKit first saves the record to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/creationdate
func (c_ CKRecord) CreationDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("creationDate"))
	return rv
}


// The time when CloudKit first saves the record to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/creationdate
func (c_ CKRecord) SetCreationDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreationDate:"), value)
}


// The ID of the user who creates the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/creatoruserrecordid
func (c_ CKRecord) CreatorUserRecordID() objc.IObject /* cross-framework: CKRecordID */ {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("creatorUserRecordID"))
	return rv
}


// The ID of the user who creates the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/creatoruserrecordid
func (c_ CKRecord) SetCreatorUserRecordID(value objc.IObject /* cross-framework: CKRecordID */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreatorUserRecordID:"), value)
}


// An object that manages the record’s encrypted key-value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/encryptedvalues
func (c_ CKRecord) EncryptedValues() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("encryptedValues"))
	return rv
}


// An object that manages the record’s encrypted key-value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/encryptedvalues
func (c_ CKRecord) SetEncryptedValues(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEncryptedValues:"), value)
}


// The ID of the user who most recently modified the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/lastmodifieduserrecordid
func (c_ CKRecord) LastModifiedUserRecordID() objc.IObject /* cross-framework: CKRecordID */ {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("lastModifiedUserRecordID"))
	return rv
}


// The ID of the user who most recently modified the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/lastmodifieduserrecordid
func (c_ CKRecord) SetLastModifiedUserRecordID(value objc.IObject /* cross-framework: CKRecordID */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLastModifiedUserRecordID:"), value)
}


// The most recent time that CloudKit saved the record to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/modificationdate
func (c_ CKRecord) ModificationDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("modificationDate"))
	return rv
}


// The most recent time that CloudKit saved the record to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/modificationdate
func (c_ CKRecord) SetModificationDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModificationDate:"), value)
}


// A reference to the record’s parent record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/parent
func (c_ CKRecord) Parent() ICKReference {
	rv := objc.Send[CKReference](c_.ID, objc.Sel("parent"))
	return rv
}


// A reference to the record’s parent record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/parent
func (c_ CKRecord) SetParent(value ICKReference) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParent:"), value)
}


// The unique ID of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordid
func (c_ CKRecord) RecordID() objc.IObject /* cross-framework: CKRecordID */ {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("recordID"))
	return rv
}


// The unique ID of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordid
func (c_ CKRecord) SetRecordID(value objc.IObject /* cross-framework: CKRecordID */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordID:"), value)
}


// The value that your app defines to identify the type of record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordtype-6v7au
func (c_ CKRecord) RecordType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordType"))
	return rv
}


// The value that your app defines to identify the type of record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordtype-6v7au
func (c_ CKRecord) SetRecordType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordType:"), value)
}


// A reference to the share object that determines the share status of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/share
func (c_ CKRecord) Share() ICKReference {
	rv := objc.Send[CKReference](c_.ID, objc.Sel("share"))
	return rv
}


// A reference to the share object that determines the share status of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/share
func (c_ CKRecord) SetShare(value ICKReference) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShare:"), value)
}



