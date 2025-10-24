// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKFetchRecordZoneChangesOptions */


/* debug [class_header]: Header for CKFetchRecordZoneChangesOptions */
// The class instance for the [CKFetchRecordZoneChangesOptions] class.
var (
	CKFetchRecordZoneChangesOptionsClass     _CKFetchRecordZoneChangesOptionsClass
	CKFetchRecordZoneChangesOptionsClassOnce sync.Once
)

func getCKFetchRecordZoneChangesOptionsClass() _CKFetchRecordZoneChangesOptionsClass {
	CKFetchRecordZoneChangesOptionsClassOnce.Do(func() {
		CKFetchRecordZoneChangesOptionsClass = _CKFetchRecordZoneChangesOptionsClass{objc.GetClass("CKFetchRecordZoneChangesOptions")}
	})
	return CKFetchRecordZoneChangesOptionsClass
}

type _CKFetchRecordZoneChangesOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKFetchRecordZoneChangesOptions */
// An interface definition for the [CKFetchRecordZoneChangesOptions] class.
type ICKFetchRecordZoneChangesOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKFetchRecordZoneChangesOptions */
	// properties:
	DesiredKeys() []string
	SetDesiredKeys(value []string)
	PreviousServerChangeToken() ICKServerChangeToken
	SetPreviousServerChangeToken(value ICKServerChangeToken)
	ResultsLimit() uint
	SetResultsLimit(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKFetchRecordZoneChangesOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKFetchRecordZoneChangesOptions */
// Alloc allocates a new instance without initialization.
func (cc _CKFetchRecordZoneChangesOptionsClass) Alloc() CKFetchRecordZoneChangesOptions {
	rv := objc.Send[CKFetchRecordZoneChangesOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKFetchRecordZoneChangesOptionsClass) New() CKFetchRecordZoneChangesOptions {
	rv := objc.Send[CKFetchRecordZoneChangesOptions](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKFetchRecordZoneChangesOptions) Init() CKFetchRecordZoneChangesOptions {
	rv := objc.Send[CKFetchRecordZoneChangesOptions](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKFetchRecordZoneChangesOptions) Autorelease() CKFetchRecordZoneChangesOptions {
	rv := objc.Send[CKFetchRecordZoneChangesOptions](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchRecordZoneChangesOptions creates a new CKFetchRecordZoneChangesOptions instance.
func NewCKFetchRecordZoneChangesOptions() CKFetchRecordZoneChangesOptions {
	return getCKFetchRecordZoneChangesOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKFetchRecordZoneChangesOptions */
// A configuration object that describes the information to fetch from a record zone.


// A configuration object that describes the information to fetch from a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneOptions
type CKFetchRecordZoneChangesOptions struct {
	objectivec.Object
}

// CKFetchRecordZoneChangesOptionsFrom constructs a [CKFetchRecordZoneChangesOptions] from an unsafe.Pointer.
//
// A configuration object that describes the information to fetch from a record zone.
func CKFetchRecordZoneChangesOptionsFrom(ptr unsafe.Pointer) CKFetchRecordZoneChangesOptions {
	return CKFetchRecordZoneChangesOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKFetchRecordZoneChangesOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKFetchRecordZoneChangesOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKFetchRecordZoneChangesOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKFetchRecordZoneChangesOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKFetchRecordZoneChangesOptions */

// The fields to fetch for the requested records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneOptions/desiredKeys
func (c_ CKFetchRecordZoneChangesOptions) DesiredKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("desiredKeys"))
	return rv
}/* debug [instance_properties/getter]: desiredKeys */


// The fields to fetch for the requested records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneOptions/desiredKeys
func (c_ CKFetchRecordZoneChangesOptions) SetDesiredKeys(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setDesiredKeys:"), nsArray)
}/* debug [instance_properties/setter]: desiredKeys */


// The token that identifies the starting point for retrieving changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneOptions/previousServerChangeToken
func (c_ CKFetchRecordZoneChangesOptions) PreviousServerChangeToken() ICKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](c_.ID, objc.Sel("previousServerChangeToken"))
	return rv
}/* debug [instance_properties/getter]: previousServerChangeToken */


// The token that identifies the starting point for retrieving changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneOptions/previousServerChangeToken
func (c_ CKFetchRecordZoneChangesOptions) SetPreviousServerChangeToken(value ICKServerChangeToken) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviousServerChangeToken:"), value)
}/* debug [instance_properties/setter]: previousServerChangeToken */


// The maximum number of records to fetch from the record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneOptions/resultsLimit
func (c_ CKFetchRecordZoneChangesOptions) ResultsLimit() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("resultsLimit"))
	return rv
}/* debug [instance_properties/getter]: resultsLimit */


// The maximum number of records to fetch from the record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneOptions/resultsLimit
func (c_ CKFetchRecordZoneChangesOptions) SetResultsLimit(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}/* debug [instance_properties/setter]: resultsLimit */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKFetchRecordZoneChangesOptions */



