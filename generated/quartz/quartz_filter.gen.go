// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [QuartzFilter] class.
var (
	QuartzFilterClass     _QuartzFilterClass
	QuartzFilterClassOnce sync.Once
)

func getQuartzFilterClass() _QuartzFilterClass {
	QuartzFilterClassOnce.Do(func() {
		QuartzFilterClass = _QuartzFilterClass{objc.GetClass("QuartzFilter")}
	})
	return QuartzFilterClass
}

type _QuartzFilterClass struct {
	class objc.Class
}

// An interface definition for the [QuartzFilter] class.
type IQuartzFilter interface {
	objectivec.IObject
	Properties() unsafe.Pointer
}

//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter
type QuartzFilter struct {
	objectivec.Object
}

// QuartzFilterFrom constructs a [QuartzFilter] from an unsafe.Pointer.
func QuartzFilterFrom(ptr unsafe.Pointer) QuartzFilter {
	return QuartzFilter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (qc _QuartzFilterClass) Alloc() QuartzFilter {
	rv := objc.Send[QuartzFilter](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QuartzFilterClass) New() QuartzFilter {
	rv := objc.Send[QuartzFilter](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QuartzFilter) Init() QuartzFilter {
	rv := objc.Send[QuartzFilter](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QuartzFilter) Autorelease() QuartzFilter {
	rv := objc.Send[QuartzFilter](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuartzFilter creates a new QuartzFilter instance.
func NewQuartzFilter() QuartzFilter {
	return getQuartzFilterClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/init(properties:)
func NewQuartzFilterWithProperties(properties objc.ID) QuartzFilter {
	rv := objc.Send[QuartzFilter](objc.ID(getQuartzFilterClass().class), objc.Sel("quartzFilterWithProperties:"), properties)
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/init(properties:)
func (qc _QuartzFilterClass) QuartzFilterWithProperties(properties objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(qc.class), objc.Sel("quartzFilterWithProperties:"), properties)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/properties()
func (q_ QuartzFilter) Properties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](q_.ID, objc.Sel("properties"))
	return rv
}


