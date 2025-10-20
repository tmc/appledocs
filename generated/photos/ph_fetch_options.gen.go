// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHFetchOptions] class.
var (
	PHFetchOptionsClass     _PHFetchOptionsClass
	PHFetchOptionsClassOnce sync.Once
)

func getPHFetchOptionsClass() _PHFetchOptionsClass {
	PHFetchOptionsClassOnce.Do(func() {
		PHFetchOptionsClass = _PHFetchOptionsClass{objc.GetClass("PHFetchOptions")}
	})
	return PHFetchOptionsClass
}

type _PHFetchOptionsClass struct {
	class objc.Class
}

// An interface definition for the [PHFetchOptions] class.
type IPHFetchOptions interface {
	objectivec.IObject
}

// A set of options that affect the filtering, sorting, and management of results that Photos returns when you fetch asset or collection objects.
//
// Using class methods on the , , , and classes to fetch assets or collections produces a object containing the requested objects. The options you specify control which objects the fetch result includes, how those objects are arranged in the fetch result, and how Photos should notify your app of changes to the fetch result. Photos supports only a restricted set of keys for the and properties. The set of available keys depends on which class you’re using to fetch assets or collections. The following table lists the keys supported by each class:
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions
type PHFetchOptions struct {
	objectivec.Object
}

// PHFetchOptionsFrom constructs a [PHFetchOptions] from an unsafe.Pointer.
//
// A set of options that affect the filtering, sorting, and management of results that Photos returns when you fetch asset or collection objects.
func PHFetchOptionsFrom(ptr unsafe.Pointer) PHFetchOptions {
	return PHFetchOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHFetchOptionsClass) Alloc() PHFetchOptions {
	rv := objc.Send[PHFetchOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHFetchOptionsClass) New() PHFetchOptions {
	rv := objc.Send[PHFetchOptions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHFetchOptions) Init() PHFetchOptions {
	rv := objc.Send[PHFetchOptions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHFetchOptions) Autorelease() PHFetchOptions {
	rv := objc.Send[PHFetchOptions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHFetchOptions creates a new PHFetchOptions instance.
func NewPHFetchOptions() PHFetchOptions {
	return getPHFetchOptionsClass().New()
}


// The maximum number of objects to include in the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/fetchLimit
func (p_ PHFetchOptions) FetchLimit() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("fetchLimit"))
	return rv
}


// SetFetchLimit sets the value of the fetchLimit property.
// The maximum number of objects to include in the fetch result.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/fetchLimit
func (p_ PHFetchOptions) SetFetchLimit(value uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFetchLimit:"), value)
}
// A Boolean value that determines whether the fetch result includes all assets from burst photo sequences.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeAllBurstAssets
func (p_ PHFetchOptions) IncludeAllBurstAssets() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("includeAllBurstAssets"))
	return rv
}


// SetIncludeAllBurstAssets sets the value of the includeAllBurstAssets property.
// A Boolean value that determines whether the fetch result includes all assets from burst photo sequences.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeAllBurstAssets
func (p_ PHFetchOptions) SetIncludeAllBurstAssets(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIncludeAllBurstAssets:"), value)
}
// The set of source types for which to include assets in the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeAssetSourceTypes
func (p_ PHFetchOptions) IncludeAssetSourceTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("includeAssetSourceTypes"))
	return rv
}


// SetIncludeAssetSourceTypes sets the value of the includeAssetSourceTypes property.
// The set of source types for which to include assets in the fetch result.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeAssetSourceTypes
func (p_ PHFetchOptions) SetIncludeAssetSourceTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIncludeAssetSourceTypes:"), value)
}
// A Boolean value that determines whether the fetch result includes assets marked as hidden.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeHiddenAssets
func (p_ PHFetchOptions) IncludeHiddenAssets() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("includeHiddenAssets"))
	return rv
}


// SetIncludeHiddenAssets sets the value of the includeHiddenAssets property.
// A Boolean value that determines whether the fetch result includes assets marked as hidden.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeHiddenAssets
func (p_ PHFetchOptions) SetIncludeHiddenAssets(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIncludeHiddenAssets:"), value)
}
// A predicate that specifies which properties to select results by and that also specifies any constraints on selection.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/predicate
func (p_ PHFetchOptions) Predicate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("predicate"))
	return rv
}


// SetPredicate sets the value of the predicate property.
// A predicate that specifies which properties to select results by and that also specifies any constraints on selection.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/predicate
func (p_ PHFetchOptions) SetPredicate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPredicate:"), value)
}
// A list of sort descriptors, specifying an order for the fetched objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/sortDescriptors
func (p_ PHFetchOptions) SortDescriptors() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](p_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// SetSortDescriptors sets the value of the sortDescriptors property.
// A list of sort descriptors, specifying an order for the fetched objects.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/sortDescriptors
func (p_ PHFetchOptions) SetSortDescriptors(value []unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSortDescriptors:"), value)
}
// A Boolean value that determines whether your app receives detailed change information for the objects in the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/wantsIncrementalChangeDetails
func (p_ PHFetchOptions) WantsIncrementalChangeDetails() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("wantsIncrementalChangeDetails"))
	return rv
}


// SetWantsIncrementalChangeDetails sets the value of the wantsIncrementalChangeDetails property.
// A Boolean value that determines whether your app receives detailed change information for the objects in the fetch result.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/wantsIncrementalChangeDetails
func (p_ PHFetchOptions) SetWantsIncrementalChangeDetails(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWantsIncrementalChangeDetails:"), value)
}


