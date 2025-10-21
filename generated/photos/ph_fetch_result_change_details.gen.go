// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHFetchResultChangeDetails] class.
var (
	PHFetchResultChangeDetailsClass     _PHFetchResultChangeDetailsClass
	PHFetchResultChangeDetailsClassOnce sync.Once
)

func getPHFetchResultChangeDetailsClass() _PHFetchResultChangeDetailsClass {
	PHFetchResultChangeDetailsClassOnce.Do(func() {
		PHFetchResultChangeDetailsClass = _PHFetchResultChangeDetailsClass{objc.GetClass("PHFetchResultChangeDetails")}
	})
	return PHFetchResultChangeDetailsClass
}

type _PHFetchResultChangeDetailsClass struct {
	class objc.Class
}

// An interface definition for the [PHFetchResultChangeDetails] class.
type IPHFetchResultChangeDetails interface {
	objectivec.IObject
}

// A description of changes that occurred in the set of asset or collection objects listed in a fetch result.
//
// A object provides detailed information about the differences between two fetch results—one that you previously obtained and an updated one that would result if you performed the same fetch again. The change details object provides information useful for updating a UI that lists the contents of a fetch result, such as the indexes of added, removed, and rearranged objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchResultChangeDetails
type PHFetchResultChangeDetails struct {
	objectivec.Object
}

// PHFetchResultChangeDetailsFrom constructs a [PHFetchResultChangeDetails] from an unsafe.Pointer.
//
// A description of changes that occurred in the set of asset or collection objects listed in a fetch result.
func PHFetchResultChangeDetailsFrom(ptr unsafe.Pointer) PHFetchResultChangeDetails {
	return PHFetchResultChangeDetails{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHFetchResultChangeDetailsClass) Alloc() PHFetchResultChangeDetails {
	rv := objc.Send[PHFetchResultChangeDetails](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHFetchResultChangeDetailsClass) New() PHFetchResultChangeDetails {
	rv := objc.Send[PHFetchResultChangeDetails](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHFetchResultChangeDetails) Init() PHFetchResultChangeDetails {
	rv := objc.Send[PHFetchResultChangeDetails](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHFetchResultChangeDetails) Autorelease() PHFetchResultChangeDetails {
	rv := objc.Send[PHFetchResultChangeDetails](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHFetchResultChangeDetails creates a new PHFetchResultChangeDetails instance.
func NewPHFetchResultChangeDetails() PHFetchResultChangeDetails {
	return getPHFetchResultChangeDetailsClass().New()
}


// The indexes of objects in the fetch result whose content or metadata have been updated.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/changedindexes
func (p_ PHFetchResultChangeDetails) ChangedIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](p_.ID, objc.Sel("changedIndexes"))
	return rv
}


// SetChangedIndexes sets the value of the changedIndexes property.
// The indexes of objects in the fetch result whose content or metadata have been updated.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/changedindexes
func (p_ PHFetchResultChangeDetails) SetChangedIndexes(value foundation.IIndexSet) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setChangedIndexes:"), value)
}

// The objects in the fetch result whose content or metadata have been updated.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/changedobjects
func (p_ PHFetchResultChangeDetails) ChangedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("changedObjects"))
	return rv
}


// SetChangedObjects sets the value of the changedObjects property.
// The objects in the fetch result whose content or metadata have been updated.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/changedobjects
func (p_ PHFetchResultChangeDetails) SetChangedObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setChangedObjects:"), value)
}

// The current fetch result, incorporating recent changes.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/fetchresultafterchanges
func (p_ PHFetchResultChangeDetails) FetchResultAfterChanges() PHFetchResult {
	rv := objc.Send[PHFetchResult](p_.ID, objc.Sel("fetchResultAfterChanges"))
	return rv
}


// SetFetchResultAfterChanges sets the value of the fetchResultAfterChanges property.
// The current fetch result, incorporating recent changes.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/fetchresultafterchanges
func (p_ PHFetchResultChangeDetails) SetFetchResultAfterChanges(value IPHFetchResult) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFetchResultAfterChanges:"), value)
}

// The original fetch result, without recent changes.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/fetchresultbeforechanges
func (p_ PHFetchResultChangeDetails) FetchResultBeforeChanges() PHFetchResult {
	rv := objc.Send[PHFetchResult](p_.ID, objc.Sel("fetchResultBeforeChanges"))
	return rv
}


// SetFetchResultBeforeChanges sets the value of the fetchResultBeforeChanges property.
// The original fetch result, without recent changes.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/fetchresultbeforechanges
func (p_ PHFetchResultChangeDetails) SetFetchResultBeforeChanges(value IPHFetchResult) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFetchResultBeforeChanges:"), value)
}

// A Boolean value that indicates whether changes to the fetch result can be described incrementally.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/hasincrementalchanges
func (p_ PHFetchResultChangeDetails) HasIncrementalChanges() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasIncrementalChanges"))
	return rv
}


// SetHasIncrementalChanges sets the value of the hasIncrementalChanges property.
// A Boolean value that indicates whether changes to the fetch result can be described incrementally.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/hasincrementalchanges
func (p_ PHFetchResultChangeDetails) SetHasIncrementalChanges(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHasIncrementalChanges:"), value)
}

// A Boolean value that indicates whether objects have been rearranged in the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/hasmoves
func (p_ PHFetchResultChangeDetails) HasMoves() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasMoves"))
	return rv
}


// SetHasMoves sets the value of the hasMoves property.
// A Boolean value that indicates whether objects have been rearranged in the fetch result.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/hasmoves
func (p_ PHFetchResultChangeDetails) SetHasMoves(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHasMoves:"), value)
}

// The indexes where new objects have been inserted in the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/insertedindexes
func (p_ PHFetchResultChangeDetails) InsertedIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](p_.ID, objc.Sel("insertedIndexes"))
	return rv
}


// SetInsertedIndexes sets the value of the insertedIndexes property.
// The indexes where new objects have been inserted in the fetch result.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/insertedindexes
func (p_ PHFetchResultChangeDetails) SetInsertedIndexes(value foundation.IIndexSet) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInsertedIndexes:"), value)
}

// The new items that have been inserted in the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/insertedobjects
func (p_ PHFetchResultChangeDetails) InsertedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("insertedObjects"))
	return rv
}


// SetInsertedObjects sets the value of the insertedObjects property.
// The new items that have been inserted in the fetch result.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/insertedobjects
func (p_ PHFetchResultChangeDetails) SetInsertedObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInsertedObjects:"), value)
}

// The indexes from which objects have been removed from the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/removedindexes
func (p_ PHFetchResultChangeDetails) RemovedIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](p_.ID, objc.Sel("removedIndexes"))
	return rv
}


// SetRemovedIndexes sets the value of the removedIndexes property.
// The indexes from which objects have been removed from the fetch result.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/removedindexes
func (p_ PHFetchResultChangeDetails) SetRemovedIndexes(value foundation.IIndexSet) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRemovedIndexes:"), value)
}

// The items that have been removed from the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/removedobjects
func (p_ PHFetchResultChangeDetails) RemovedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("removedObjects"))
	return rv
}


// SetRemovedObjects sets the value of the removedObjects property.
// The items that have been removed from the fetch result.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/removedobjects
func (p_ PHFetchResultChangeDetails) SetRemovedObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRemovedObjects:"), value)
}



