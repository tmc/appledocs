// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSLevelIndicatorCell */


/* debug [class_header]: Header for NSLevelIndicatorCell */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LevelIndicatorCell */
// An interface definition for the [LevelIndicatorCell] class.
type ILevelIndicatorCell interface {
	IActionCell
	
/* debug [class_interface_properties]: Properties for LevelIndicatorCell */
	// properties:
	CriticalValue() float64
	SetCriticalValue(value float64)
	LevelIndicatorStyle() LevelIndicatorStyle
	SetLevelIndicatorStyle(value LevelIndicatorStyle)
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	NumberOfMajorTickMarks() int
	SetNumberOfMajorTickMarks(value int)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
	WarningValue() float64
	SetWarningValue(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LevelIndicatorCell */
	// methods:
	RectOfTickMarkAtIndex(index int) Rect /* not a class type */
	TickMarkValueAtIndex(index int) float64
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LevelIndicatorCell */
// Alloc allocates a new instance without initialization.
func (lc _LevelIndicatorCellClass) Alloc() LevelIndicatorCell {
	rv := objc.Send[LevelIndicatorCell](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LevelIndicatorCell */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LevelIndicatorCell */

// Initializes the receiver with the style specified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/init(levelIndicatorStyle:)
func NewLevelIndicatorCellWithLevelIndicatorStyle(levelIndicatorStyle LevelIndicatorStyle) LevelIndicatorCell {
	instance := getLevelIndicatorCellClass().Alloc()
	rv := objc.Send[LevelIndicatorCell](instance.ID, objc.Sel("initWithLevelIndicatorStyle:"), levelIndicatorStyle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLevelIndicatorCellWithLevelIndicatorStyle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LevelIndicatorCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LevelIndicatorCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LevelIndicatorCell */

// Returns the bounding rectangle of the tick mark identified by (the minimum-value tick mark is at index 0).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/rectOfTickMark(at:)
func (l_ LevelIndicatorCell) RectOfTickMarkAtIndex(index int) Rect /* not a class type */ {
	rv := objc.Send[Rect](l_.ID, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: RectOfTickMarkAtIndex */


// Returns the receiver’s value represented by the tick mark at index (the minimum-value tick mark has an index of 0).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/tickMarkValue(at:)
func (l_ LevelIndicatorCell) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: TickMarkValueAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LevelIndicatorCell */

// The critical value of the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/criticalValue
func (l_ LevelIndicatorCell) CriticalValue() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("criticalValue"))
	return rv
}/* debug [instance_properties/getter]: criticalValue */


// The critical value of the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/criticalValue
func (l_ LevelIndicatorCell) SetCriticalValue(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCriticalValue:"), value)
}/* debug [instance_properties/setter]: criticalValue */


// The style of the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/levelIndicatorStyle
func (l_ LevelIndicatorCell) LevelIndicatorStyle() LevelIndicatorStyle {
	rv := objc.Send[LevelIndicatorStyle](l_.ID, objc.Sel("levelIndicatorStyle"))
	return rv
}/* debug [instance_properties/getter]: levelIndicatorStyle */


// The style of the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/levelIndicatorStyle
func (l_ LevelIndicatorCell) SetLevelIndicatorStyle(value LevelIndicatorStyle) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLevelIndicatorStyle:"), value)
}/* debug [instance_properties/setter]: levelIndicatorStyle */


// The maximum value of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/maxValue
func (l_ LevelIndicatorCell) MaxValue() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("maxValue"))
	return rv
}/* debug [instance_properties/getter]: maxValue */


// The maximum value of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/maxValue
func (l_ LevelIndicatorCell) SetMaxValue(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMaxValue:"), value)
}/* debug [instance_properties/setter]: maxValue */


// The minimum value of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/minValue
func (l_ LevelIndicatorCell) MinValue() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("minValue"))
	return rv
}/* debug [instance_properties/getter]: minValue */


// The minimum value of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/minValue
func (l_ LevelIndicatorCell) SetMinValue(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMinValue:"), value)
}/* debug [instance_properties/setter]: minValue */


// The number of major tick marks displayed by the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/numberOfMajorTickMarks
func (l_ LevelIndicatorCell) NumberOfMajorTickMarks() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfMajorTickMarks"))
	return rv
}/* debug [instance_properties/getter]: numberOfMajorTickMarks */


// The number of major tick marks displayed by the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/numberOfMajorTickMarks
func (l_ LevelIndicatorCell) SetNumberOfMajorTickMarks(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfMajorTickMarks:"), value)
}/* debug [instance_properties/setter]: numberOfMajorTickMarks */


// The number of tick marks displayed by the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/numberOfTickMarks
func (l_ LevelIndicatorCell) NumberOfTickMarks() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfTickMarks"))
	return rv
}/* debug [instance_properties/getter]: numberOfTickMarks */


// The number of tick marks displayed by the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/numberOfTickMarks
func (l_ LevelIndicatorCell) SetNumberOfTickMarks(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfTickMarks:"), value)
}/* debug [instance_properties/setter]: numberOfTickMarks */


// The placement of tick marks on the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/tickMarkPosition
func (l_ LevelIndicatorCell) TickMarkPosition() TickMarkPosition {
	rv := objc.Send[TickMarkPosition](l_.ID, objc.Sel("tickMarkPosition"))
	return rv
}/* debug [instance_properties/getter]: tickMarkPosition */


// The placement of tick marks on the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/tickMarkPosition
func (l_ LevelIndicatorCell) SetTickMarkPosition(value TickMarkPosition) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTickMarkPosition:"), value)
}/* debug [instance_properties/setter]: tickMarkPosition */


// The warning value of the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/warningValue
func (l_ LevelIndicatorCell) WarningValue() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("warningValue"))
	return rv
}/* debug [instance_properties/getter]: warningValue */


// The warning value of the level indicator control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/warningValue
func (l_ LevelIndicatorCell) SetWarningValue(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWarningValue:"), value)
}/* debug [instance_properties/setter]: warningValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSLevelIndicatorCell */


