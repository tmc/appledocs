// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKReference */


/* debug [class_header]: Header for CKReference */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKReference */
// An interface definition for the [CKReference] class.
type ICKReference interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKReference */
	// properties:
	ReferenceAction() CKReferenceAction
	RecordID() ICKRecordID
	Action() objectivec.IObject
	SetAction(value objectivec.IObject)
	RecordChangeTag() objc.IObject /* cross-framework: NSString */
	SetRecordChangeTag(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKReference */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKReference */
// Alloc allocates a new instance without initialization.
func (cc _CKReferenceClass) Alloc() CKReference {
	rv := objc.Send[CKReference](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKReference */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKReference */

// Creates a reference object that points to the specified record object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/init(record:action:)
func NewCKReferenceWithRecordAction(record objc.IObject /* cross-framework: CKRecord */, action CKReferenceAction) CKReference {
	instance := getCKReferenceClass().Alloc()
	rv := objc.Send[CKReference](instance.ID, objc.Sel("initWithRecord:action:"), record, action)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKReferenceWithRecordAction */


// Creates a reference object that points to the record with the specified ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/init(recordID:action:)
func NewCKReferenceWithRecordIDAction(recordID ICKRecordID, action CKReferenceAction) CKReference {
	instance := getCKReferenceClass().Alloc()
	rv := objc.Send[CKReference](instance.ID, objc.Sel("initWithRecordID:action:"), recordID, action)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKReferenceWithRecordIDAction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKReference */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKReference */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKReference */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKReference */

// The ownership behavior for the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/action-swift.property
func (c_ CKReference) ReferenceAction() CKReferenceAction {
	rv := objc.Send[CKReferenceAction](c_.ID, objc.Sel("referenceAction"))
	return rv
}/* debug [instance_properties/getter]: referenceAction */


// The ID of the referenced record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/recordID
func (c_ CKReference) RecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("recordID"))
	return rv
}/* debug [instance_properties/getter]: recordID */


// The ownership behavior for the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/reference/action-swift.property
func (c_ CKReference) Action() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// The ownership behavior for the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/reference/action-swift.property
func (c_ CKReference) SetAction(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// The server change token for the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordchangetag
func (c_ CKReference) RecordChangeTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("recordChangeTag"))
	return rv
}/* debug [instance_properties/getter]: recordChangeTag */


// The server change token for the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordchangetag
func (c_ CKReference) SetRecordChangeTag(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordChangeTag:"), value)
}/* debug [instance_properties/setter]: recordChangeTag */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKReference */


