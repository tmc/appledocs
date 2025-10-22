// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [QuartzFilterView] class.
var (
	QuartzFilterViewClass     _QuartzFilterViewClass
	QuartzFilterViewClassOnce sync.Once
)

func getQuartzFilterViewClass() _QuartzFilterViewClass {
	QuartzFilterViewClassOnce.Do(func() {
		QuartzFilterViewClass = _QuartzFilterViewClass{objc.GetClass("QuartzFilterView")}
	})
	return QuartzFilterViewClass
}

type _QuartzFilterViewClass struct {
	class objc.Class
}

// An interface definition for the [QuartzFilterView] class.
type IQuartzFilterView interface {
	IView
	SizeToFit()
}

//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterView
type QuartzFilterView struct {
	View
}

// QuartzFilterViewFrom constructs a [QuartzFilterView] from an unsafe.Pointer.
func QuartzFilterViewFrom(ptr unsafe.Pointer) QuartzFilterView {
	return QuartzFilterView{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (qc _QuartzFilterViewClass) Alloc() QuartzFilterView {
	rv := objc.Send[QuartzFilterView](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QuartzFilterViewClass) New() QuartzFilterView {
	rv := objc.Send[QuartzFilterView](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QuartzFilterView) Init() QuartzFilterView {
	rv := objc.Send[QuartzFilterView](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QuartzFilterView) Autorelease() QuartzFilterView {
	rv := objc.Send[QuartzFilterView](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuartzFilterView creates a new QuartzFilterView instance.
func NewQuartzFilterView() QuartzFilterView {
	return getQuartzFilterViewClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterView/sizeToFit()
func (q_ QuartzFilterView) SizeToFit() {
	objc.Send[objc.ID](q_.ID, objc.Sel("sizeToFit"))
}




