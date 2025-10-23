// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StringDrawingContext] class.
var (
	StringDrawingContextClass     _StringDrawingContextClass
	StringDrawingContextClassOnce sync.Once
)

func getStringDrawingContextClass() _StringDrawingContextClass {
	StringDrawingContextClassOnce.Do(func() {
		StringDrawingContextClass = _StringDrawingContextClass{objc.GetClass("NSStringDrawingContext")}
	})
	return StringDrawingContextClass
}

type _StringDrawingContextClass struct {
	class objc.Class
}

// An interface definition for the [StringDrawingContext] class.
type IStringDrawingContext interface {
	objectivec.IObject
	// properties:
	ActualScaleFactor() float64 /* primitive/slice/pointer. */
	SetActualScaleFactor(value float64 /* primitive/slice/pointer. */)
	MinimumScaleFactor() float64 /* primitive/slice/pointer. */
	SetMinimumScaleFactor(value float64 /* primitive/slice/pointer. */)
	TotalBounds() objc.IObject /* cross-framework: Rect */
	SetTotalBounds(value objc.IObject /* cross-framework: Rect */)
	// methods:
}

// An object that manages metrics for drawing attributed strings.
//
// Prior to drawing, you can create an instance of this class and use it to specify the minimum scale factor and tracking adjustments for a string. After drawing, you can retrieve the actual values that were used during drawing. To use this class, allocate and initialize a new instance, set the minimum values, and pass your object to one of the corresponding methods that take the context object as a parameter. Upon completion of drawing, you can use the actual drawing values to make adjustments or record where the string was actually drawn.


// An object that manages metrics for drawing attributed strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingContext
type StringDrawingContext struct {
	objectivec.Object
}

// StringDrawingContextFrom constructs a [StringDrawingContext] from an unsafe.Pointer.
//
// An object that manages metrics for drawing attributed strings.
func StringDrawingContextFrom(ptr unsafe.Pointer) StringDrawingContext {
	return StringDrawingContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StringDrawingContextClass) Alloc() StringDrawingContext {
	rv := objc.Send[StringDrawingContext](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StringDrawingContextClass) New() StringDrawingContext {
	rv := objc.Send[StringDrawingContext](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StringDrawingContext) Init() StringDrawingContext {
	rv := objc.Send[StringDrawingContext](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StringDrawingContext) Autorelease() StringDrawingContext {
	rv := objc.Send[StringDrawingContext](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStringDrawingContext creates a new StringDrawingContext instance.
func NewStringDrawingContext() StringDrawingContext {
	return getStringDrawingContextClass().New()
}



// The actual scale factor that the system applied to the font during drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstringdrawingcontext/actualscalefactor
func (s_ StringDrawingContext) ActualScaleFactor() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("actualScaleFactor"))
	return rv
}


// The actual scale factor that the system applied to the font during drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstringdrawingcontext/actualscalefactor
func (s_ StringDrawingContext) SetActualScaleFactor(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setActualScaleFactor:"), value)
}


// The scale factor that determines the smallest font size to use during drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstringdrawingcontext/minimumscalefactor
func (s_ StringDrawingContext) MinimumScaleFactor() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumScaleFactor"))
	return rv
}


// The scale factor that determines the smallest font size to use during drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstringdrawingcontext/minimumscalefactor
func (s_ StringDrawingContext) SetMinimumScaleFactor(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumScaleFactor:"), value)
}


// The most recent bounding rectangle that the system used to draw the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstringdrawingcontext/totalbounds
func (s_ StringDrawingContext) TotalBounds() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("totalBounds"))
	return rv
}


// The most recent bounding rectangle that the system used to draw the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstringdrawingcontext/totalbounds
func (s_ StringDrawingContext) SetTotalBounds(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTotalBounds:"), value)
}



