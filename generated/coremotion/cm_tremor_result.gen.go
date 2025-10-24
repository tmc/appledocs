// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TremorResult] class.
var (
	TremorResultClass     _TremorResultClass
	TremorResultClassOnce sync.Once
)

func getTremorResultClass() _TremorResultClass {
	TremorResultClassOnce.Do(func() {
		TremorResultClass = _TremorResultClass{objc.GetClass("CMTremorResult")}
	})
	return TremorResultClass
}

type _TremorResultClass struct {
	class objc.Class
}

// An interface definition for the [TremorResult] class.
type ITremorResult interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A result object that contains data about the presence and strength of tremors during a one-minute interval.
//
// The following equation is always true: .


// A result object that contains data about the presence and strength of tremors during a one-minute interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult
type TremorResult struct {
	objectivec.Object
}

// TremorResultFrom constructs a [TremorResult] from an unsafe.Pointer.
//
// A result object that contains data about the presence and strength of tremors during a one-minute interval.
func TremorResultFrom(ptr unsafe.Pointer) TremorResult {
	return TremorResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TremorResultClass) Alloc() TremorResult {
	rv := objc.Send[TremorResult](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TremorResultClass) New() TremorResult {
	rv := objc.Send[TremorResult](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TremorResult) Init() TremorResult {
	rv := objc.Send[TremorResult](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TremorResult) Autorelease() TremorResult {
	rv := objc.Send[TremorResult](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTremorResult creates a new TremorResult instance.
func NewTremorResult() TremorResult {
	return getTremorResultClass().New()
}



