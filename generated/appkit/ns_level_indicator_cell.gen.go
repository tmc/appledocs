// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LevelIndicatorCell] class.
var (
	LevelIndicatorCellClass     _LevelIndicatorCellClass
	LevelIndicatorCellClassOnce sync.Once
)

func getLevelIndicatorCellClass() _LevelIndicatorCellClass {
	LevelIndicatorCellClassOnce.Do(func() {
		LevelIndicatorCellClass = _LevelIndicatorCellClass{objc.GetClass("NSLevelIndicatorCell")}
	})
	return LevelIndicatorCellClass
}

type _LevelIndicatorCellClass struct {
	class objc.Class
}

// An interface definition for the [LevelIndicatorCell] class.
type ILevelIndicatorCell interface {
	IActionCell
	// properties:
	CriticalValue() float64 /* primitive/slice/pointer. */
	SetCriticalValue(value float64 /* primitive/slice/pointer. */)
	LevelIndicatorStyle() unsafe.Pointer
	SetLevelIndicatorStyle(value unsafe.Pointer)
	MaxValue() float64 /* primitive/slice/pointer. */
	SetMaxValue(value float64 /* primitive/slice/pointer. */)
	MinValue() float64 /* primitive/slice/pointer. */
	SetMinValue(value float64 /* primitive/slice/pointer. */)
	NumberOfMajorTickMarks() int /* primitive/slice/pointer. */
	SetNumberOfMajorTickMarks(value int /* primitive/slice/pointer. */)
	NumberOfTickMarks() int /* primitive/slice/pointer. */
	SetNumberOfTickMarks(value int /* primitive/slice/pointer. */)
	TickMarkPosition() unsafe.Pointer
	SetTickMarkPosition(value unsafe.Pointer)
	WarningValue() float64 /* primitive/slice/pointer. */
	SetWarningValue(value float64 /* primitive/slice/pointer. */)
	// methods:
}

// is a subclass of that provides several level indicator display styles including: capacity, ranking and relevancy. The capacity style provides both continuous and discrete modes.


// is a subclass of that provides several level indicator display styles including: capacity, ranking and relevancy. The capacity style provides both continuous and discrete modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell
type LevelIndicatorCell struct {
	ActionCell
}

// LevelIndicatorCellFrom constructs a [LevelIndicatorCell] from an unsafe.Pointer.
//
// is a subclass of that provides several level indicator display styles including: capacity, ranking and relevancy. The capacity style provides both continuous and discrete modes.
func LevelIndicatorCellFrom(ptr unsafe.Pointer) LevelIndicatorCell {
	return LevelIndicatorCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (lc _LevelIndicatorCellClass) Alloc() LevelIndicatorCell {
	rv := objc.Send[LevelIndicatorCell](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LevelIndicatorCellClass) New() LevelIndicatorCell {
	rv := objc.Send[LevelIndicatorCell](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LevelIndicatorCell) Init() LevelIndicatorCell {
	rv := objc.Send[LevelIndicatorCell](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LevelIndicatorCell) Autorelease() LevelIndicatorCell {
	rv := objc.Send[LevelIndicatorCell](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLevelIndicatorCell creates a new LevelIndicatorCell instance.
func NewLevelIndicatorCell() LevelIndicatorCell {
	return getLevelIndicatorCellClass().New()
}



// The critical value of the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/criticalvalue
func (l_ LevelIndicatorCell) CriticalValue() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](l_.ID, objc.Sel("criticalValue"))
	return rv
}


// The critical value of the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/criticalvalue
func (l_ LevelIndicatorCell) SetCriticalValue(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCriticalValue:"), value)
}


// The style of the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/levelindicatorstyle
func (l_ LevelIndicatorCell) LevelIndicatorStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("levelIndicatorStyle"))
	return rv
}


// The style of the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/levelindicatorstyle
func (l_ LevelIndicatorCell) SetLevelIndicatorStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLevelIndicatorStyle:"), value)
}


// The maximum value of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/maxvalue
func (l_ LevelIndicatorCell) MaxValue() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](l_.ID, objc.Sel("maxValue"))
	return rv
}


// The maximum value of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/maxvalue
func (l_ LevelIndicatorCell) SetMaxValue(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMaxValue:"), value)
}


// The minimum value of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/minvalue
func (l_ LevelIndicatorCell) MinValue() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](l_.ID, objc.Sel("minValue"))
	return rv
}


// The minimum value of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/minvalue
func (l_ LevelIndicatorCell) SetMinValue(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMinValue:"), value)
}


// The number of major tick marks displayed by the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/numberofmajortickmarks
func (l_ LevelIndicatorCell) NumberOfMajorTickMarks() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfMajorTickMarks"))
	return rv
}


// The number of major tick marks displayed by the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/numberofmajortickmarks
func (l_ LevelIndicatorCell) SetNumberOfMajorTickMarks(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfMajorTickMarks:"), value)
}


// The number of tick marks displayed by the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/numberoftickmarks
func (l_ LevelIndicatorCell) NumberOfTickMarks() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfTickMarks"))
	return rv
}


// The number of tick marks displayed by the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/numberoftickmarks
func (l_ LevelIndicatorCell) SetNumberOfTickMarks(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfTickMarks:"), value)
}


// The placement of tick marks on the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/tickmarkposition
func (l_ LevelIndicatorCell) TickMarkPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("tickMarkPosition"))
	return rv
}


// The placement of tick marks on the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/tickmarkposition
func (l_ LevelIndicatorCell) SetTickMarkPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTickMarkPosition:"), value)
}


// The warning value of the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/warningvalue
func (l_ LevelIndicatorCell) WarningValue() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](l_.ID, objc.Sel("warningValue"))
	return rv
}


// The warning value of the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicatorcell/warningvalue
func (l_ LevelIndicatorCell) SetWarningValue(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWarningValue:"), value)
}



