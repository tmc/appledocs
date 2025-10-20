// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [QueryGenerationToken] class.
var (
	queryGenerationTokenClass     _QueryGenerationTokenClass
	queryGenerationTokenClassOnce sync.Once
)

func getQueryGenerationTokenClass() _QueryGenerationTokenClass {
	queryGenerationTokenClassOnce.Do(func() {
		queryGenerationTokenClass = _QueryGenerationTokenClass{objc.GetClass("NSQueryGenerationToken")}
	})
	return queryGenerationTokenClass
}

type _QueryGenerationTokenClass struct {
	class objc.Class
}

// An interface definition for the [QueryGenerationToken] class.
type IQueryGenerationToken interface {
	objectivec.IObject
}

// A token that indicates which generation of the persistent store is being accessed.
//
// When a managed object context is pinned to a specific generation of the app data, a query generation token will be associated with that context.
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

// Alloc allocates a new instance without initialization.
func (qc _QueryGenerationTokenClass) Alloc() QueryGenerationToken {
	rv := objc.Send[QueryGenerationToken](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QueryGenerationTokenClass) New() QueryGenerationToken {
	rv := objc.Send[QueryGenerationToken](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QueryGenerationToken) Init() QueryGenerationToken {
	rv := objc.Send[QueryGenerationToken](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QueryGenerationToken) Autorelease() QueryGenerationToken {
	rv := objc.Send[QueryGenerationToken](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQueryGenerationToken creates a new QueryGenerationToken instance.
func NewQueryGenerationToken() QueryGenerationToken {
	return getQueryGenerationTokenClass().New()
}




