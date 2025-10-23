// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKReference] class.
var (
	CKReferenceClass     _CKReferenceClass
	CKReferenceClassOnce sync.Once
)

func getCKReferenceClass() _CKReferenceClass {
	CKReferenceClassOnce.Do(func() {
		CKReferenceClass = _CKReferenceClass{objc.GetClass("CKReference")}
	})
	return CKReferenceClass
}

type _CKReferenceClass struct {
	class objc.Class
}

// An interface definition for the [CKReference] class.
type ICKReference interface {
	objectivec.IObject
	ReferenceAction() unsafe.Pointer
	RecordID() CKRecordID
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	RecordChangeTag() string
	SetRecordChangeTag(value string)
}

// A relationship between two records in a record zone.
//
// A object creates a many-to-one relationship between records in your database. Each reference object stores information about the one record that is the target of the reference. You then save the reference object in the fields of one or more records to create a link from those records to the target. Both records must be in the same zone of the same database. References create a stronger relationship between records than just saving the ID of a record as a string. Specifically, you can use references to create an ownership model between two records. When the reference object’s action is , the target of the reference—that is, the record in the reference’s property—becomes the owner of the source record. Deleting the target (owner) record deletes all its source records. The deletion of any owned records can trigger further deletions if those records are the owners of other records. If a record contains two or more objects with an action of , CloudKit deletes the record when it deletes any of the objects it references. To save multiple records that contain references between them, save the target records first or save all the records in one batch operation using .


// A relationship between two records in a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference
type CKReference struct {
	objectivec.Object
}

// CKReferenceFrom constructs a [CKReference] from an unsafe.Pointer.
//
// A relationship between two records in a record zone.
func CKReferenceFrom(ptr unsafe.Pointer) CKReference {
	return CKReference{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKReferenceClass) Alloc() CKReference {
	rv := objc.Send[CKReference](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKReferenceClass) New() CKReference {
	rv := objc.Send[CKReference](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKReference) Init() CKReference {
	rv := objc.Send[CKReference](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKReference) Autorelease() CKReference {
	rv := objc.Send[CKReference](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKReference creates a new CKReference instance.
func NewCKReference() CKReference {
	return getCKReferenceClass().New()
}



// Creates a reference object that points to the specified record object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/init(record:action:)
func NewCKReferenceWithRecordAction(record ICKRecord, action unsafe.Pointer) CKReference {
	instance := getCKReferenceClass().Alloc()
	rv := objc.Send[CKReference](instance.ID, objc.Sel("initWithRecord:action:"), record, action)
	rv.Autorelease()
	return rv
}


// Creates a reference object that points to the record with the specified ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/init(recordID:action:)
func NewCKReferenceWithRecordIDAction(recordID CKRecordID, action unsafe.Pointer) CKReference {
	instance := getCKReferenceClass().Alloc()
	rv := objc.Send[CKReference](instance.ID, objc.Sel("initWithRecordID:action:"), recordID, action)
	rv.Autorelease()
	return rv
}



// The ownership behavior for the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/action-swift.property
func (c_ CKReference) ReferenceAction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("referenceAction"))
	return rv
}


// The ID of the referenced record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/recordID
func (c_ CKReference) RecordID() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("recordID"))
	return rv
}


// The ownership behavior for the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/reference/action-swift.property
func (c_ CKReference) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("action"))
	return rv
}


// The ownership behavior for the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/reference/action-swift.property
func (c_ CKReference) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}


// The server change token for the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordchangetag
func (c_ CKReference) RecordChangeTag() string {
	rv := objc.Send[string](c_.ID, objc.Sel("recordChangeTag"))
	return rv
}


// The server change token for the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordchangetag
func (c_ CKReference) SetRecordChangeTag(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordChangeTag:"), objc.String(value))
}


