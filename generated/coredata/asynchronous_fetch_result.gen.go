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



