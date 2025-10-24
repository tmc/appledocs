// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CKDatabaseOperation */


/* debug [class_header]: Header for CKDatabaseOperation */
// The class instance for the [CKDatabaseOperation] class.
var (
	CKDatabaseOperationClass     _CKDatabaseOperationClass
	CKDatabaseOperationClassOnce sync.Once
)

func getCKDatabaseOperationClass() _CKDatabaseOperationClass {
	CKDatabaseOperationClassOnce.Do(func() {
		CKDatabaseOperationClass = _CKDatabaseOperationClass{objc.GetClass("CKDatabaseOperation")}
	})
	return CKDatabaseOperationClass
}

type _CKDatabaseOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKDatabaseOperation */
// An interface definition for the [CKDatabaseOperation] class.
type ICKDatabaseOperation interface {
	ICKOperation
	
/* debug [class_interface_properties]: Properties for CKDatabaseOperation */
	// properties:
	Database() ICKDatabase
	SetDatabase(value ICKDatabase)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKDatabaseOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKDatabaseOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKDatabaseOperationClass) Alloc() CKDatabaseOperation {
	rv := objc.Send[CKDatabaseOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKDatabaseOperationClass) New() CKDatabaseOperation {
	rv := objc.Send[CKDatabaseOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKDatabaseOperation) Init() CKDatabaseOperation {
	rv := objc.Send[CKDatabaseOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKDatabaseOperation) Autorelease() CKDatabaseOperation {
	rv := objc.Send[CKDatabaseOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKDatabaseOperation creates a new CKDatabaseOperation instance.
func NewCKDatabaseOperation() CKDatabaseOperation {
	return getCKDatabaseOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKDatabaseOperation */
// The abstract base class for operations that act upon databases in CloudKit.
//
// Database operations typically involve fetching and saving records and other database objects, as well as executing queries on the contents of the database. Use this class’s property to tell the operation which database to use when you execute it. Don’t subclass this class or create instances of it. Instead, create instances of one of its concrete subclasses.


// The abstract base class for operations that act upon databases in CloudKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabaseOperation
type CKDatabaseOperation struct {
	CKOperation
}

// CKDatabaseOperationFrom constructs a [CKDatabaseOperation] from an unsafe.Pointer.
//
// The abstract base class for operations that act upon databases in CloudKit.
func CKDatabaseOperationFrom(ptr unsafe.Pointer) CKDatabaseOperation {
	return CKDatabaseOperation{
		CKOperation: CKOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKDatabaseOperation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKDatabaseOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKDatabaseOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKDatabaseOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKDatabaseOperation */

// The database that the operation uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabaseOperation/database
func (c_ CKDatabaseOperation) Database() ICKDatabase {
	rv := objc.Send[CKDatabase](c_.ID, objc.Sel("database"))
	return rv
}/* debug [instance_properties/getter]: database */


// The database that the operation uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabaseOperation/database
func (c_ CKDatabaseOperation) SetDatabase(value ICKDatabase) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDatabase:"), value)
}/* debug [instance_properties/setter]: database */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKDatabaseOperation */



