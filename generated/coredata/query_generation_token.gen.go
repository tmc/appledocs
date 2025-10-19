// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [QueryGenerationToken] class.
var queryGenerationTokenClass = _QueryGenerationTokenClass{objc.GetClass("NSQueryGenerationToken")}

type _QueryGenerationTokenClass struct {
	class objc.Class
}

// A token that indicates which generation of the persistent store is being accessed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSQueryGenerationToken

type QueryGenerationToken struct {
	objectivec.Object
}

// QueryGenerationTokenFrom constructs a [QueryGenerationToken] from an unsafe.Pointer.
//
// A token that indicates which generation of the persistent store is being accessed.
func QueryGenerationTokenFrom(ptr unsafe.Pointer) QueryGenerationToken {
	return QueryGenerationToken{objectivec.Object{objc.ID(ptr)}}
}



