// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSLevelIndicator */


/* debug [class_header]: Header for NSLevelIndicator */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LevelIndicator */
// An interface definition for the [LevelIndicator] class.
type ILevelIndicator interface {
	IControl
	
/* debug [class_interface_properties]: Properties for LevelIndicator */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LevelIndicator */
	// methods:
	RectOfTickMarkAtIndex(index int) Rect /* not a class type */
	TickMarkValueAtIndex(index int) float64
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LevelIndicator */
// Alloc allocates a new instance without initialization.
func (lc _LevelIndicatorClass) Alloc() LevelIndicator {
	rv := objc.Send[LevelIndicator](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LevelIndicator */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LevelIndicator *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LevelIndicator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LevelIndicator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LevelIndicator */

// Returns the bounding rectangle of the tick mark identified by the specified index (the minimum-value tick mark is at index 0).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/rectOfTickMark(at:)
func (l_ LevelIndicator) RectOfTickMarkAtIndex(index int) Rect /* not a class type */ {
	rv := objc.Send[Rect](l_.ID, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: RectOfTickMarkAtIndex */


// Returns the receiver’s value represented by the tick mark at the specified index (the minimum-value tick mark has an index of 0).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/tickMarkValue(at:)
func (l_ LevelIndicator) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: TickMarkValueAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LevelIndicator */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/criticalFillColor
func (l_ LevelIndicator) CriticalFillColor() IColor {
	rv := objc.Send[Color](l_.ID, objc.Sel("criticalFillColor"))
	return rv
}/* debug [instance_properties/getter]: criticalFillColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/criticalFillColor
func (l_ LevelIndicator) SetCriticalFillColor(value IColor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCriticalFillColor:"), value)
}/* debug [instance_properties/setter]: criticalFillColor */


// The receiver’s critical value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/criticalValue
func (l_ LevelIndicator) CriticalValue() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("criticalValue"))
	return rv
}/* debug [instance_properties/getter]: criticalValue */


// The receiver’s critical value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/criticalValue
func (l_ LevelIndicator) SetCriticalValue(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCriticalValue:"), value)
}/* debug [instance_properties/setter]: criticalValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/drawsTieredCapacityLevels
func (l_ LevelIndicator) DrawsTieredCapacityLevels() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("drawsTieredCapacityLevels"))
	return rv
}/* debug [instance_properties/getter]: drawsTieredCapacityLevels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/drawsTieredCapacityLevels
func (l_ LevelIndicator) SetDrawsTieredCapacityLevels(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDrawsTieredCapacityLevels:"), value)
}/* debug [instance_properties/setter]: drawsTieredCapacityLevels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/fillColor
func (l_ LevelIndicator) FillColor() IColor {
	rv := objc.Send[Color](l_.ID, objc.Sel("fillColor"))
	return rv
}/* debug [instance_properties/getter]: fillColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/fillColor
func (l_ LevelIndicator) SetFillColor(value IColor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setFillColor:"), value)
}/* debug [instance_properties/setter]: fillColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/isEditable
func (l_ LevelIndicator) Editable() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("editable"))
	return rv
}/* debug [instance_properties/getter]: editable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/isEditable
func (l_ LevelIndicator) SetEditable(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setEditable:"), value)
}/* debug [instance_properties/setter]: editable */


// The appearance of the indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/levelIndicatorStyle
func (l_ LevelIndicator) LevelIndicatorStyle() LevelIndicatorStyle {
	rv := objc.Send[LevelIndicatorStyle](l_.ID, objc.Sel("levelIndicatorStyle"))
	return rv
}/* debug [instance_properties/getter]: levelIndicatorStyle */


// The appearance of the indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/levelIndicatorStyle
func (l_ LevelIndicator) SetLevelIndicatorStyle(value LevelIndicatorStyle) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLevelIndicatorStyle:"), value)
}/* debug [instance_properties/setter]: levelIndicatorStyle */


// The receiver’s maximum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/maxValue
func (l_ LevelIndicator) MaxValue() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("maxValue"))
	return rv
}/* debug [instance_properties/getter]: maxValue */


// The receiver’s maximum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/maxValue
func (l_ LevelIndicator) SetMaxValue(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMaxValue:"), value)
}/* debug [instance_properties/setter]: maxValue */


// The receiver’s minimum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/minValue
func (l_ LevelIndicator) MinValue() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("minValue"))
	return rv
}/* debug [instance_properties/getter]: minValue */


// The receiver’s minimum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/minValue
func (l_ LevelIndicator) SetMinValue(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMinValue:"), value)
}/* debug [instance_properties/setter]: minValue */


// The number of major tick marks associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/numberOfMajorTickMarks
func (l_ LevelIndicator) NumberOfMajorTickMarks() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfMajorTickMarks"))
	return rv
}/* debug [instance_properties/getter]: numberOfMajorTickMarks */


// The number of major tick marks associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/numberOfMajorTickMarks
func (l_ LevelIndicator) SetNumberOfMajorTickMarks(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfMajorTickMarks:"), value)
}/* debug [instance_properties/setter]: numberOfMajorTickMarks */


// The number of tick marks associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/numberOfTickMarks
func (l_ LevelIndicator) NumberOfTickMarks() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfTickMarks"))
	return rv
}/* debug [instance_properties/getter]: numberOfTickMarks */


// The number of tick marks associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/numberOfTickMarks
func (l_ LevelIndicator) SetNumberOfTickMarks(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfTickMarks:"), value)
}/* debug [instance_properties/setter]: numberOfTickMarks */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/placeholderVisibility-swift.property
func (l_ LevelIndicator) PlaceholderVisibility() LevelIndicatorPlaceholderVisibility {
	rv := objc.Send[LevelIndicatorPlaceholderVisibility](l_.ID, objc.Sel("placeholderVisibility"))
	return rv
}/* debug [instance_properties/getter]: placeholderVisibility */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/placeholderVisibility-swift.property
func (l_ LevelIndicator) SetPlaceholderVisibility(value LevelIndicatorPlaceholderVisibility) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPlaceholderVisibility:"), value)
}/* debug [instance_properties/setter]: placeholderVisibility */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/ratingImage
func (l_ LevelIndicator) RatingImage() IImage {
	rv := objc.Send[Image](l_.ID, objc.Sel("ratingImage"))
	return rv
}/* debug [instance_properties/getter]: ratingImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/ratingImage
func (l_ LevelIndicator) SetRatingImage(value IImage) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRatingImage:"), value)
}/* debug [instance_properties/setter]: ratingImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/ratingPlaceholderImage
func (l_ LevelIndicator) RatingPlaceholderImage() IImage {
	rv := objc.Send[Image](l_.ID, objc.Sel("ratingPlaceholderImage"))
	return rv
}/* debug [instance_properties/getter]: ratingPlaceholderImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/ratingPlaceholderImage
func (l_ LevelIndicator) SetRatingPlaceholderImage(value IImage) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRatingPlaceholderImage:"), value)
}/* debug [instance_properties/setter]: ratingPlaceholderImage */


// Determines how the receiver’s tick marks are aligned with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/tickMarkPosition
func (l_ LevelIndicator) TickMarkPosition() TickMarkPosition {
	rv := objc.Send[TickMarkPosition](l_.ID, objc.Sel("tickMarkPosition"))
	return rv
}/* debug [instance_properties/getter]: tickMarkPosition */


// Determines how the receiver’s tick marks are aligned with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/tickMarkPosition
func (l_ LevelIndicator) SetTickMarkPosition(value TickMarkPosition) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTickMarkPosition:"), value)
}/* debug [instance_properties/setter]: tickMarkPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/warningFillColor
func (l_ LevelIndicator) WarningFillColor() IColor {
	rv := objc.Send[Color](l_.ID, objc.Sel("warningFillColor"))
	return rv
}/* debug [instance_properties/getter]: warningFillColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/warningFillColor
func (l_ LevelIndicator) SetWarningFillColor(value IColor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWarningFillColor:"), value)
}/* debug [instance_properties/setter]: warningFillColor */


// The receiver’s warning value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/warningValue
func (l_ LevelIndicator) WarningValue() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("warningValue"))
	return rv
}/* debug [instance_properties/getter]: warningValue */


// The receiver’s warning value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/warningValue
func (l_ LevelIndicator) SetWarningValue(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWarningValue:"), value)
}/* debug [instance_properties/setter]: warningValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/iseditable
func (l_ LevelIndicator) IsEditable() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isEditable"))
	return rv
}/* debug [instance_properties/getter]: isEditable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/iseditable
func (l_ LevelIndicator) SetIsEditable(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsEditable:"), value)
}/* debug [instance_properties/setter]: isEditable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSLevelIndicator */



