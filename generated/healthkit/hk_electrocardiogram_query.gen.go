// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKElectrocardiogramQuery] class.
var (
	HKElectrocardiogramQueryClass     _HKElectrocardiogramQueryClass
	HKElectrocardiogramQueryClassOnce sync.Once
)

func getHKElectrocardiogramQueryClass() _HKElectrocardiogramQueryClass {
	HKElectrocardiogramQueryClassOnce.Do(func() {
		HKElectrocardiogramQueryClass = _HKElectrocardiogramQueryClass{objc.GetClass("HKElectrocardiogramQuery")}
	})
	return HKElectrocardiogramQueryClass
}

type _HKElectrocardiogramQueryClass struct {
	class objc.Class
}

// An interface definition for the [HKElectrocardiogramQuery] class.
type IHKElectrocardiogramQuery interface {
	IHKQuery
}

// A query that returns the underlying voltage measurements for an electrocardiogram sample.
//
// Use the query to access the individual voltage measurements associated with an sample. The query calls the data handler once for each voltage measurement, passing a instance that contains the voltage data. After it has sent all the voltage measurements, the query calls the data handler one last time, passing . If an error occurs, it stops collecting voltage data and passes instead. Electrocardiogram queries are immutable: You set query’s properties when you create it, and they don’t change.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogramQuery
type HKElectrocardiogramQuery struct {
	HKQuery
}

// HKElectrocardiogramQueryFrom constructs a [HKElectrocardiogramQuery] from an unsafe.Pointer.
//
// A query that returns the underlying voltage measurements for an electrocardiogram sample.
func HKElectrocardiogramQueryFrom(ptr unsafe.Pointer) HKElectrocardiogramQuery {
	return HKElectrocardiogramQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKElectrocardiogramQueryClass) Alloc() HKElectrocardiogramQuery {
	rv := objc.Send[HKElectrocardiogramQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKElectrocardiogramQueryClass) New() HKElectrocardiogramQuery {
	rv := objc.Send[HKElectrocardiogramQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKElectrocardiogramQuery) Init() HKElectrocardiogramQuery {
	rv := objc.Send[HKElectrocardiogramQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKElectrocardiogramQuery) Autorelease() HKElectrocardiogramQuery {
	rv := objc.Send[HKElectrocardiogramQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKElectrocardiogramQuery creates a new HKElectrocardiogramQuery instance.
func NewHKElectrocardiogramQuery() HKElectrocardiogramQuery {
	return getHKElectrocardiogramQueryClass().New()
}




