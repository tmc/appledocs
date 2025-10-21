// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CKFetchRecordZoneChangesConfiguration] class.
type ICKFetchRecordZoneChangesConfiguration interface {
	objectivec.IObject
}

// A configuration object that describes the information to fetch from a record zone.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CKFetchRecordZoneChangesConfigurationClass) Alloc() CKFetchRecordZoneChangesConfiguration {
	rv := objc.Send[CKFetchRecordZoneChangesConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// An array of the record keys to retrieve.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesConfiguration/desiredKeys
func (c_ CKFetchRecordZoneChangesConfiguration) DesiredKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("desiredKeys"))
	return rv
}


// SetDesiredKeys sets the value of the desiredKeys property.
// An array of the record keys to retrieve.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesConfiguration/desiredKeys
func (c_ CKFetchRecordZoneChangesConfiguration) SetDesiredKeys(value []string) {
	// Convert Go slice to NSArray
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
}
// The server change token.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration/previousServerChangeToken
func (c_ CKFetchRecordZoneChangesConfiguration) PreviousServerChangeToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("previousServerChangeToken"))
	return rv
}


// SetPreviousServerChangeToken sets the value of the previousServerChangeToken property.
// The server change token.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration/previousServerChangeToken
func (c_ CKFetchRecordZoneChangesConfiguration) SetPreviousServerChangeToken(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviousServerChangeToken:"), value)
}
// The maximum number of records that CloudKit retrieves when fetching zone changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration/resultsLimit
func (c_ CKFetchRecordZoneChangesConfiguration) ResultsLimit() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("resultsLimit"))
	return rv
}


// SetResultsLimit sets the value of the resultsLimit property.
// The maximum number of records that CloudKit retrieves when fetching zone changes.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration/resultsLimit
func (c_ CKFetchRecordZoneChangesConfiguration) SetResultsLimit(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}


