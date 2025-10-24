// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

// The class instance for the [LevelIndicator] class.
var (
	LevelIndicatorClass     _LevelIndicatorClass
	LevelIndicatorClassOnce sync.Once
)

func getLevelIndicatorClass() _LevelIndicatorClass {
	LevelIndicatorClassOnce.Do(func() {
		LevelIndicatorClass = _LevelIndicatorClass{objc.GetClass("NSLevelIndicator")}
	})
	return LevelIndicatorClass
}

type _LevelIndicatorClass struct {
	class objc.Class
}

// An interface definition for the [LevelIndicator] class.
type ILevelIndicator interface {
	IControl
	// properties:
	CriticalFillColor() IColor
	SetCriticalFillColor(value IColor)
	CriticalValue() float64
	SetCriticalValue(value float64)
	DrawsTieredCapacityLevels() bool
	SetDrawsTieredCapacityLevels(value bool)
	FillColor() IColor
	SetFillColor(value IColor)
	Editable() bool
	SetEditable(value bool)
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
	PlaceholderVisibility() LevelIndicatorPlaceholderVisibility
	SetPlaceholderVisibility(value LevelIndicatorPlaceholderVisibility)
	RatingImage() IImage
	SetRatingImage(value IImage)
	RatingPlaceholderImage() IImage
	SetRatingPlaceholderImage(value IImage)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
	WarningFillColor() IColor
	SetWarningFillColor(value IColor)
	WarningValue() float64
	SetWarningValue(value float64)
	IsEditable() bool
	SetIsEditable(value bool)
	// methods:
	RectOfTickMarkAtIndex(index int) objc.IObject /* cross-framework: Rect */
	TickMarkValueAtIndex(index int) float64
}

// A visual representation of a level or quantity, using discrete values.
//
// A level indicator is similar to an object, but provides a more customized visual feedback to the user. Unlike sliders, level indicators do not have a “knob” indicating the current setting, and they do not allow the user to adjust the current setting. You set the value of the level indicator programmatically. The supported indicator styles include: A capacity style level indicator. The continuous mode for this style is often used to indicate conditions such as how much data is on hard disk. The discrete mode is similar to audio level indicators in audio playback applications. You can specify both a warning value and a critical value that provides additional visual feedback to the user. A ranking style level indicator. This is similar to the star ranking displays provided in iTunes and iPhoto. You can also specify your own ranking image. A relevancy style level indicator. This style is used to display the relevancy of a search result, for example in Mail. uses an to implement much of the control’s functionality. provides cover methods for most of the methods, which call the corresponding cell method.


// A visual representation of a level or quantity, using discrete values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator
type LevelIndicator struct {
	Control
}

// LevelIndicatorFrom constructs a [LevelIndicator] from an unsafe.Pointer.
//
// A visual representation of a level or quantity, using discrete values.
func LevelIndicatorFrom(ptr unsafe.Pointer) LevelIndicator {
	return LevelIndicator{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (lc _LevelIndicatorClass) Alloc() LevelIndicator {
	rv := objc.Send[LevelIndicator](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LevelIndicatorClass) New() LevelIndicator {
	rv := objc.Send[LevelIndicator](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LevelIndicator) Init() LevelIndicator {
	rv := objc.Send[LevelIndicator](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LevelIndicator) Autorelease() LevelIndicator {
	rv := objc.Send[LevelIndicator](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLevelIndicator creates a new LevelIndicator instance.
func NewLevelIndicator() LevelIndicator {
	return getLevelIndicatorClass().New()
}



// Returns the bounding rectangle of the tick mark identified by the specified index (the minimum-value tick mark is at index 0).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/rectOfTickMark(at:)
func (l_ LevelIndicator) RectOfTickMarkAtIndex(index int) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return rv
}


// Returns the receiver’s value represented by the tick mark at the specified index (the minimum-value tick mark has an index of 0).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/tickMarkValue(at:)
func (l_ LevelIndicator) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/criticalFillColor
func (l_ LevelIndicator) CriticalFillColor() IColor {
	rv := objc.Send[Color](l_.ID, objc.Sel("criticalFillColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/criticalFillColor
func (l_ LevelIndicator) SetCriticalFillColor(value IColor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCriticalFillColor:"), value)
}


// The receiver’s critical value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/criticalValue
func (l_ LevelIndicator) CriticalValue() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("criticalValue"))
	return rv
}


// The receiver’s critical value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/criticalValue
func (l_ LevelIndicator) SetCriticalValue(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCriticalValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/drawsTieredCapacityLevels
func (l_ LevelIndicator) DrawsTieredCapacityLevels() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("drawsTieredCapacityLevels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/drawsTieredCapacityLevels
func (l_ LevelIndicator) SetDrawsTieredCapacityLevels(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDrawsTieredCapacityLevels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/fillColor
func (l_ LevelIndicator) FillColor() IColor {
	rv := objc.Send[Color](l_.ID, objc.Sel("fillColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/fillColor
func (l_ LevelIndicator) SetFillColor(value IColor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setFillColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/isEditable
func (l_ LevelIndicator) Editable() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("editable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/isEditable
func (l_ LevelIndicator) SetEditable(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setEditable:"), value)
}


// The appearance of the indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/levelIndicatorStyle
func (l_ LevelIndicator) LevelIndicatorStyle() LevelIndicatorStyle {
	rv := objc.Send[LevelIndicatorStyle](l_.ID, objc.Sel("levelIndicatorStyle"))
	return rv
}


// The appearance of the indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/levelIndicatorStyle
func (l_ LevelIndicator) SetLevelIndicatorStyle(value LevelIndicatorStyle) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLevelIndicatorStyle:"), value)
}


// The receiver’s maximum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/maxValue
func (l_ LevelIndicator) MaxValue() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("maxValue"))
	return rv
}


// The receiver’s maximum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/maxValue
func (l_ LevelIndicator) SetMaxValue(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMaxValue:"), value)
}


// The receiver’s minimum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/minValue
func (l_ LevelIndicator) MinValue() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("minValue"))
	return rv
}


// The receiver’s minimum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/minValue
func (l_ LevelIndicator) SetMinValue(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMinValue:"), value)
}


// The number of major tick marks associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/numberOfMajorTickMarks
func (l_ LevelIndicator) NumberOfMajorTickMarks() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfMajorTickMarks"))
	return rv
}


// The number of major tick marks associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/numberOfMajorTickMarks
func (l_ LevelIndicator) SetNumberOfMajorTickMarks(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfMajorTickMarks:"), value)
}


// The number of tick marks associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/numberOfTickMarks
func (l_ LevelIndicator) NumberOfTickMarks() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfTickMarks"))
	return rv
}


// The number of tick marks associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/numberOfTickMarks
func (l_ LevelIndicator) SetNumberOfTickMarks(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfTickMarks:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/placeholderVisibility-swift.property
func (l_ LevelIndicator) PlaceholderVisibility() LevelIndicatorPlaceholderVisibility {
	rv := objc.Send[LevelIndicatorPlaceholderVisibility](l_.ID, objc.Sel("placeholderVisibility"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/placeholderVisibility-swift.property
func (l_ LevelIndicator) SetPlaceholderVisibility(value LevelIndicatorPlaceholderVisibility) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPlaceholderVisibility:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/ratingImage
func (l_ LevelIndicator) RatingImage() IImage {
	rv := objc.Send[Image](l_.ID, objc.Sel("ratingImage"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/ratingImage
func (l_ LevelIndicator) SetRatingImage(value IImage) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRatingImage:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/ratingPlaceholderImage
func (l_ LevelIndicator) RatingPlaceholderImage() IImage {
	rv := objc.Send[Image](l_.ID, objc.Sel("ratingPlaceholderImage"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/ratingPlaceholderImage
func (l_ LevelIndicator) SetRatingPlaceholderImage(value IImage) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRatingPlaceholderImage:"), value)
}


// Determines how the receiver’s tick marks are aligned with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/tickMarkPosition
func (l_ LevelIndicator) TickMarkPosition() TickMarkPosition {
	rv := objc.Send[TickMarkPosition](l_.ID, objc.Sel("tickMarkPosition"))
	return rv
}


// Determines how the receiver’s tick marks are aligned with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/tickMarkPosition
func (l_ LevelIndicator) SetTickMarkPosition(value TickMarkPosition) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTickMarkPosition:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/warningFillColor
func (l_ LevelIndicator) WarningFillColor() IColor {
	rv := objc.Send[Color](l_.ID, objc.Sel("warningFillColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/warningFillColor
func (l_ LevelIndicator) SetWarningFillColor(value IColor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWarningFillColor:"), value)
}


// The receiver’s warning value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/warningValue
func (l_ LevelIndicator) WarningValue() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("warningValue"))
	return rv
}


// The receiver’s warning value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/warningValue
func (l_ LevelIndicator) SetWarningValue(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWarningValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/iseditable
func (l_ LevelIndicator) IsEditable() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isEditable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/iseditable
func (l_ LevelIndicator) SetIsEditable(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsEditable:"), value)
}



