// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKFetchRecordZoneChangesConfiguration */


/* debug [class_header]: Header for CKFetchRecordZoneChangesConfiguration */
// The class instance for the [CKFetchRecordZoneChangesConfiguration] class.
var (
	CKFetchRecordZoneChangesConfigurationClass     _CKFetchRecordZoneChangesConfigurationClass
	CKFetchRecordZoneChangesConfigurationClassOnce sync.Once
)

func getCKFetchRecordZoneChangesConfigurationClass() _CKFetchRecordZoneChangesConfigurationClass {
	CKFetchRecordZoneChangesConfigurationClassOnce.Do(func() {
		CKFetchRecordZoneChangesConfigurationClass = _CKFetchRecordZoneChangesConfigurationClass{objc.GetClass("CKFetchRecordZoneChangesConfiguration")}
	})
	return CKFetchRecordZoneChangesConfigurationClass
}

type _CKFetchRecordZoneChangesConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKFetchRecordZoneChangesConfiguration */
// An interface definition for the [CKFetchRecordZoneChangesConfiguration] class.
type ICKFetchRecordZoneChangesConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKFetchRecordZoneChangesConfiguration */
	// properties:
	DesiredKeys() []string
	SetDesiredKeys(value []string)
	PreviousServerChangeToken() ICKServerChangeToken
	SetPreviousServerChangeToken(value ICKServerChangeToken)
	ResultsLimit() uint
	SetResultsLimit(value uint)
	ConfigurationsByRecordZoneID() ICKFetchRecordZoneChangesConfiguration
	SetConfigurationsByRecordZoneID(value ICKFetchRecordZoneChangesConfiguration)
	FetchAllChanges() bool
	SetFetchAllChanges(value bool)
	RecordZoneIDs() ICKRecordZoneID
	SetRecordZoneIDs(value ICKRecordZoneID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKFetchRecordZoneChangesConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKFetchRecordZoneChangesConfiguration */
// Alloc allocates a new instance without initialization.
func (cc _CKFetchRecordZoneChangesConfigurationClass) Alloc() CKFetchRecordZoneChangesConfiguration {
	rv := objc.Send[CKFetchRecordZoneChangesConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKFetchRecordZoneChangesConfigurationClass) New() CKFetchRecordZoneChangesConfiguration {
	rv := objc.Send[CKFetchRecordZoneChangesConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKFetchRecordZoneChangesConfiguration) Init() CKFetchRecordZoneChangesConfiguration {
	rv := objc.Send[CKFetchRecordZoneChangesConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKFetchRecordZoneChangesConfiguration) Autorelease() CKFetchRecordZoneChangesConfiguration {
	rv := objc.Send[CKFetchRecordZoneChangesConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchRecordZoneChangesConfiguration creates a new CKFetchRecordZoneChangesConfiguration instance.
func NewCKFetchRecordZoneChangesConfiguration() CKFetchRecordZoneChangesConfiguration {
	return getCKFetchRecordZoneChangesConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKFetchRecordZoneChangesConfiguration */
// A configuration object that describes the information to fetch from a record zone.


// A configuration object that describes the information to fetch from a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration
type CKFetchRecordZoneChangesConfiguration struct {
	objectivec.Object
}

// CKFetchRecordZoneChangesConfigurationFrom constructs a [CKFetchRecordZoneChangesConfiguration] from an unsafe.Pointer.
//
// A configuration object that describes the information to fetch from a record zone.
func CKFetchRecordZoneChangesConfigurationFrom(ptr unsafe.Pointer) CKFetchRecordZoneChangesConfiguration {
	return CKFetchRecordZoneChangesConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKFetchRecordZoneChangesConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKFetchRecordZoneChangesConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKFetchRecordZoneChangesConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKFetchRecordZoneChangesConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKFetchRecordZoneChangesConfiguration */

// An array of the record keys to retrieve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesConfiguration/desiredKeys
func (c_ CKFetchRecordZoneChangesConfiguration) DesiredKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("desiredKeys"))
	return rv
}/* debug [instance_properties/getter]: desiredKeys */


// An array of the record keys to retrieve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesConfiguration/desiredKeys
func (c_ CKFetchRecordZoneChangesConfiguration) SetDesiredKeys(value []string) {
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


// The server change token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration/previousServerChangeToken
func (c_ CKFetchRecordZoneChangesConfiguration) PreviousServerChangeToken() ICKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](c_.ID, objc.Sel("previousServerChangeToken"))
	return rv
}/* debug [instance_properties/getter]: previousServerChangeToken */


// The server change token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration/previousServerChangeToken
func (c_ CKFetchRecordZoneChangesConfiguration) SetPreviousServerChangeToken(value ICKServerChangeToken) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviousServerChangeToken:"), value)
}/* debug [instance_properties/setter]: previousServerChangeToken */


// The maximum number of records that CloudKit retrieves when fetching zone changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration/resultsLimit
func (c_ CKFetchRecordZoneChangesConfiguration) ResultsLimit() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("resultsLimit"))
	return rv
}/* debug [instance_properties/getter]: resultsLimit */


// The maximum number of records that CloudKit retrieves when fetching zone changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration/resultsLimit
func (c_ CKFetchRecordZoneChangesConfiguration) SetResultsLimit(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}/* debug [instance_properties/setter]: resultsLimit */


// A dictionary of configurations for fetching change operations by zone identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/configurationsbyrecordzoneid
func (c_ CKFetchRecordZoneChangesConfiguration) ConfigurationsByRecordZoneID() ICKFetchRecordZoneChangesConfiguration {
	rv := objc.Send[CKFetchRecordZoneChangesConfiguration](c_.ID, objc.Sel("configurationsByRecordZoneID"))
	return rv
}/* debug [instance_properties/getter]: configurationsByRecordZoneID */


// A dictionary of configurations for fetching change operations by zone identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/configurationsbyrecordzoneid
func (c_ CKFetchRecordZoneChangesConfiguration) SetConfigurationsByRecordZoneID(value ICKFetchRecordZoneChangesConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfigurationsByRecordZoneID:"), value)
}/* debug [instance_properties/setter]: configurationsByRecordZoneID */


// A Boolean value that indicates whether to send repeated requests to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/fetchallchanges
func (c_ CKFetchRecordZoneChangesConfiguration) FetchAllChanges() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("fetchAllChanges"))
	return rv
}/* debug [instance_properties/getter]: fetchAllChanges */


// A Boolean value that indicates whether to send repeated requests to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/fetchallchanges
func (c_ CKFetchRecordZoneChangesConfiguration) SetFetchAllChanges(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchAllChanges:"), value)
}/* debug [instance_properties/setter]: fetchAllChanges */


// The IDs of the record zones that contain the records to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/recordzoneids
func (c_ CKFetchRecordZoneChangesConfiguration) RecordZoneIDs() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("recordZoneIDs"))
	return rv
}/* debug [instance_properties/getter]: recordZoneIDs */


// The IDs of the record zones that contain the records to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/recordzoneids
func (c_ CKFetchRecordZoneChangesConfiguration) SetRecordZoneIDs(value ICKRecordZoneID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneIDs:"), value)
}/* debug [instance_properties/setter]: recordZoneIDs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKFetchRecordZoneChangesConfiguration */



