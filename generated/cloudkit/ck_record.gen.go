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
	AllKeys() []string
	AllTokens() []string
	ChangedKeys() []string
	EncodeSystemFieldsWithCoder(coder foundation.ICoder)
	ObjectForKey(key unsafe.Pointer) objc.ID
	SetObjectForKey(object objectivec.IObject, key unsafe.Pointer)
	SetObjectForKeyedSubscript(object objectivec.IObject, key unsafe.Pointer)
	SetParentReferenceFromRecord(parentRecord ICKRecord)
	SetParentReferenceFromRecordID(parentRecordID ICKRecordID)
	ObjectForKeyedSubscript(key unsafe.Pointer) objc.ID
}

// A collection of key-value pairs that store your app’s data.
//
// Records are the fundamental objects that manage data in CloudKit. You can define any number of record types for your app, with each record type corresponding to a different type of information. Within a record type, you then define one or more fields, each with a name and a value. Records can contain simple data types, such as strings and numbers, or more complex types, such as geographic locations or pointers to other records. An important step in using CloudKit is defining the record types your app supports. A new record object doesn’t contain any keys or values. During development, you can add new keys and values at any time. The first time you set a value for a key and save the record, the server associates that type with the key for all records of the same type. The class doesn’t impose these type constraints or do any local validation of a record’s contents. CloudKit enforces these constraints when you save the records. Although records behave like dictionaries, there are limitations to the types of values you can assign to keys. The following are the object types that the class supports. Attempting to specify objects of any other type results in failure. Fields of all types are searchable unless otherwise noted.
//
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




// Creates a new record of the specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/initWithRecordType:
func NewCKRecordWithRecordType(recordType unsafe.Pointer) CKRecord {
	instance := getCKRecordClass().Alloc()
	rv := objc.Send[CKRecord](instance.ID, objc.Sel("initWithRecordType:"), recordType)
	rv.Autorelease()
	return rv
}



// Creates a record using an ID that you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/initWithRecordType:recordID:
func NewCKRecordWithRecordTypeRecordID(recordType unsafe.Pointer, recordID ICKRecordID) CKRecord {
	instance := getCKRecordClass().Alloc()
	rv := objc.Send[CKRecord](instance.ID, objc.Sel("initWithRecordType:recordID:"), recordType, recordID)
	rv.Autorelease()
	return rv
}



// Creates a record in the specified zone.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/initWithRecordType:zoneID:
func NewCKRecordWithRecordTypeZoneID(recordType unsafe.Pointer, zoneID ICKRecordZoneID) CKRecord {
	instance := getCKRecordClass().Alloc()
	rv := objc.Send[CKRecord](instance.ID, objc.Sel("initWithRecordType:zoneID:"), recordType, zoneID)
	rv.Autorelease()
	return rv
}


// Returns an array of the record’s keys.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/allKeys
func (c_ CKRecord) AllKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("allKeys"))
	return rv
}

// Returns an array of strings to use for full-text searches of the field’s string-based values.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/allTokens()
func (c_ CKRecord) AllTokens() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("allTokens"))
	return rv
}

// Returns an array of keys with recent changes to their values.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/changedKeys
func (c_ CKRecord) ChangedKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("changedKeys"))
	return rv
}

// Encodes the record’s system fields using the specified archiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/encodeSystemFields(with:)
func (c_ CKRecord) EncodeSystemFieldsWithCoder(coder foundation.ICoder) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeSystemFieldsWithCoder:"), coder)
}

// Returns the object that the record stores for the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/objectForKey:
func (c_ CKRecord) ObjectForKey(key unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("objectForKey:"), key)
	return rv
}

// Stores an object in the record using the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/setObject:forKey:
func (c_ CKRecord) SetObjectForKey(object objectivec.IObject, key unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObject:forKey:"), object, key)
}

// Stores an object in the record using the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/setObject:forKeyedSubscript:
func (c_ CKRecord) SetObjectForKeyedSubscript(object objectivec.IObject, key unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObject:forKeyedSubscript:"), object, key)
}

// Creates and sets a reference object for a parent from its record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/setParent(_:)-23du1
func (c_ CKRecord) SetParentReferenceFromRecord(parentRecord ICKRecord) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParentReferenceFromRecord:"), parentRecord)
}

// Creates and sets a reference object for a parent from the parent’s record ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/setParent(_:)-7egcx
func (c_ CKRecord) SetParentReferenceFromRecordID(parentRecordID ICKRecordID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParentReferenceFromRecordID:"), parentRecordID)
}

// Returns the object that the record stores for the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/subscript(_:)-51whk
func (c_ CKRecord) ObjectForKeyedSubscript(key unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}

// The time when CloudKit first saves the record to the server.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/creationDate
func (c_ CKRecord) CreationDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("creationDate"))
	return rv
}

// The ID of the user who creates the record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/creatorUserRecordID
func (c_ CKRecord) CreatorUserRecordID() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("creatorUserRecordID"))
	return rv
}

// An object that manages the record’s encrypted key-value pairs.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/encryptedValues
func (c_ CKRecord) EncryptedValues() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("encryptedValues"))
	return rv
}

// The ID of the user who most recently modified the record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/lastModifiedUserRecordID
func (c_ CKRecord) LastModifiedUserRecordID() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("lastModifiedUserRecordID"))
	return rv
}

// The most recent time that CloudKit saved the record to the server.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/modificationDate
func (c_ CKRecord) ModificationDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("modificationDate"))
	return rv
}

// A reference to the record’s parent record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/parent
func (c_ CKRecord) Parent() CKReference {
	rv := objc.Send[CKReference](c_.ID, objc.Sel("parent"))
	return rv
}


// SetParent sets the value of the parent property.
// A reference to the record’s parent record.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/parent
func (c_ CKRecord) SetParent(value ICKReference) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParent:"), value)
}

// The server change token for the record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/recordChangeTag
func (c_ CKRecord) RecordChangeTag() string {
	rv := objc.Send[string](c_.ID, objc.Sel("recordChangeTag"))
	return rv
}

// The unique ID of the record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/recordID
func (c_ CKRecord) RecordID() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("recordID"))
	return rv
}

// The value that your app defines to identify the type of record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/recordType-9s09b
func (c_ CKRecord) RecordType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordType"))
	return rv
}

// A reference to the share object that determines the share status of the record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/share
func (c_ CKRecord) Share() CKReference {
	rv := objc.Send[CKReference](c_.ID, objc.Sel("share"))
	return rv
}


