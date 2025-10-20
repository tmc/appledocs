// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [queryTimeoutInSeconds] class.
var (
	QueryTimeoutInSecondsClass     _queryTimeoutInSecondsClass
	QueryTimeoutInSecondsClassOnce sync.Once
)

func getqueryTimeoutInSecondsClass() _queryTimeoutInSecondsClass {
	QueryTimeoutInSecondsClassOnce.Do(func() {
		QueryTimeoutInSecondsClass = _queryTimeoutInSecondsClass{objc.GetClass("queryTimeoutInSeconds")}
	})
	return QueryTimeoutInSecondsClass
}

type _queryTimeoutInSecondsClass struct {
	class objc.Class
}

// An interface definition for the [queryTimeoutInSeconds] class.
type IqueryTimeoutInSeconds interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/queryTimeoutInSeconds-c.ivar
type queryTimeoutInSeconds struct {
	objectivec.Object
}

// queryTimeoutInSecondsFrom constructs a [queryTimeoutInSeconds] from an unsafe.Pointer.
func queryTimeoutInSecondsFrom(ptr unsafe.Pointer) queryTimeoutInSeconds {
	return queryTimeoutInSeconds{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (qc _queryTimeoutInSecondsClass) Alloc() queryTimeoutInSeconds {
	rv := objc.Send[queryTimeoutInSeconds](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _queryTimeoutInSecondsClass) New() queryTimeoutInSeconds {
	rv := objc.Send[queryTimeoutInSeconds](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ queryTimeoutInSeconds) Init() queryTimeoutInSeconds {
	rv := objc.Send[queryTimeoutInSeconds](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ queryTimeoutInSeconds) Autorelease() queryTimeoutInSeconds {
	rv := objc.Send[queryTimeoutInSeconds](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewqueryTimeoutInSeconds creates a new queryTimeoutInSeconds instance.
func NewqueryTimeoutInSeconds() queryTimeoutInSeconds {
	return getqueryTimeoutInSecondsClass().New()
}




