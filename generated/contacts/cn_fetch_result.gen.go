// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CNFetchResult] class.
var (
	CNFetchResultClass     _CNFetchResultClass
	CNFetchResultClassOnce sync.Once
)

func getCNFetchResultClass() _CNFetchResultClass {
	CNFetchResultClassOnce.Do(func() {
		CNFetchResultClass = _CNFetchResultClass{objc.GetClass("CNFetchResult")}
	})
	return CNFetchResultClass
}

type _CNFetchResultClass struct {
	class objc.Class
}

// An interface definition for the [CNFetchResult] class.
type ICNFetchResult interface {
	objectivec.IObject
}

// An object that represents the result of a change-history fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNFetchResult
type CNFetchResult struct {
	objectivec.Object
}

// CNFetchResultFrom constructs a [CNFetchResult] from an unsafe.Pointer.
//
// An object that represents the result of a change-history fetch request.
func CNFetchResultFrom(ptr unsafe.Pointer) CNFetchResult {
	return CNFetchResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNFetchResultClass) Alloc() CNFetchResult {
	rv := objc.Send[CNFetchResult](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNFetchResultClass) New() CNFetchResult {
	rv := objc.Send[CNFetchResult](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNFetchResult) Init() CNFetchResult {
	rv := objc.Send[CNFetchResult](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNFetchResult) Autorelease() CNFetchResult {
	rv := objc.Send[CNFetchResult](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNFetchResult creates a new CNFetchResult instance.
func NewCNFetchResult() CNFetchResult {
	return getCNFetchResultClass().New()
}


// An opaque token that indicates a point in history in the user’s Contacts database.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNFetchResult/currentHistoryToken
func (c_ CNFetchResult) CurrentHistoryToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("currentHistoryToken"))
	return rv
}

// The result of the fetch request, expressed as the value type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNFetchResult/value
func (c_ CNFetchResult) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("value"))
	return rv
}



