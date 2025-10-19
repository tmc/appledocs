// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AsynchronousFetchResult] class.
var asynchronousFetchResultClass = _AsynchronousFetchResultClass{objc.GetClass("NSAsynchronousFetchResult")}

type _AsynchronousFetchResultClass struct {
	class objc.Class
}

// An interface definition for the [AsynchronousFetchResult] class.
type IAsynchronousFetchResult interface {
	IPersistentStoreAsynchronousResult
}

// A fetch result object that encompasses the response from an executed asynchronous fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAsynchronousFetchResult

type AsynchronousFetchResult struct {
	PersistentStoreAsynchronousResult
}

// AsynchronousFetchResultFrom constructs a [AsynchronousFetchResult] from an unsafe.Pointer.
//
// A fetch result object that encompasses the response from an executed asynchronous fetch request.
func AsynchronousFetchResultFrom(ptr unsafe.Pointer) AsynchronousFetchResult {
	return AsynchronousFetchResult{
		PersistentStoreAsynchronousResult: PersistentStoreAsynchronousResultFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (ac _AsynchronousFetchResultClass) Alloc() AsynchronousFetchResult {
	rv := objc.Send[AsynchronousFetchResult](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AsynchronousFetchResultClass) New() AsynchronousFetchResult {
	rv := objc.Send[AsynchronousFetchResult](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AsynchronousFetchResult) Init() AsynchronousFetchResult {
	rv := objc.Send[AsynchronousFetchResult](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AsynchronousFetchResult) Autorelease() AsynchronousFetchResult {
	rv := objc.Send[AsynchronousFetchResult](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAsynchronousFetchResult creates a new AsynchronousFetchResult instance.
func NewAsynchronousFetchResult() AsynchronousFetchResult {
	return asynchronousFetchResultClass.New()
}




