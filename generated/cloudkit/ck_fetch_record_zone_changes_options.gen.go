// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CKFetchRecordZoneChangesOptions] class.
type ICKFetchRecordZoneChangesOptions interface {
	objectivec.IObject
	DesiredKeys() []string
	SetDesiredKeys(value []string)
	PreviousServerChangeToken() CKServerChangeToken
	SetPreviousServerChangeToken(value ICKServerChangeToken)
	ResultsLimit() int
	SetResultsLimit(value int)
}

// A configuration object that describes the information to fetch from a record zone.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CKFetchRecordZoneChangesOptionsClass) Alloc() CKFetchRecordZoneChangesOptions {
	rv := objc.Send[CKFetchRecordZoneChangesOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The fields to fetch for the requested records.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneOptions/desiredKeys
func (c_ CKFetchRecordZoneChangesOptions) DesiredKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("desiredKeys"))
	return rv
}


// SetDesiredKeys sets the value of the desiredKeys property.
// The fields to fetch for the requested records.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneOptions/desiredKeys
func (c_ CKFetchRecordZoneChangesOptions) SetDesiredKeys(value []string) {
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

// The token that identifies the starting point for retrieving changes.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/zoneoptions/previousserverchangetoken
func (c_ CKFetchRecordZoneChangesOptions) PreviousServerChangeToken() CKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](c_.ID, objc.Sel("previousServerChangeToken"))
	return rv
}


// SetPreviousServerChangeToken sets the value of the previousServerChangeToken property.
// The token that identifies the starting point for retrieving changes.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/zoneoptions/previousserverchangetoken
func (c_ CKFetchRecordZoneChangesOptions) SetPreviousServerChangeToken(value ICKServerChangeToken) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviousServerChangeToken:"), value)
}

// The maximum number of records to fetch from the record zone.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/zoneoptions/resultslimit
func (c_ CKFetchRecordZoneChangesOptions) ResultsLimit() int {
	rv := objc.Send[int](c_.ID, objc.Sel("resultsLimit"))
	return rv
}


// SetResultsLimit sets the value of the resultsLimit property.
// The maximum number of records to fetch from the record zone.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/zoneoptions/resultslimit
func (c_ CKFetchRecordZoneChangesOptions) SetResultsLimit(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}



