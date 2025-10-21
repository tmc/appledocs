// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A visual representation of a level or quantity, using discrete values.
//
// A level indicator is similar to an object, but provides a more customized visual feedback to the user. Unlike sliders, level indicators do not have a “knob” indicating the current setting, and they do not allow the user to adjust the current setting. You set the value of the level indicator programmatically. The supported indicator styles include: A capacity style level indicator. The continuous mode for this style is often used to indicate conditions such as how much data is on hard disk. The discrete mode is similar to audio level indicators in audio playback applications. You can specify both a warning value and a critical value that provides additional visual feedback to the user. A ranking style level indicator. This is similar to the star ranking displays provided in iTunes and iPhoto. You can also specify your own ranking image. A relevancy style level indicator. This style is used to display the relevancy of a search result, for example in Mail. uses an to implement much of the control’s functionality. provides cover methods for most of the methods, which call the corresponding cell method.
//
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


//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/criticalfillcolor
func (l_ LevelIndicator) CriticalFillColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("criticalFillColor"))
	return rv
}


// SetCriticalFillColor sets the value of the criticalFillColor property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/criticalfillcolor
func (l_ LevelIndicator) SetCriticalFillColor(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCriticalFillColor:"), value)
}

// The receiver’s critical value.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/criticalvalue
func (l_ LevelIndicator) CriticalValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("criticalValue"))
	return rv
}


// SetCriticalValue sets the value of the criticalValue property.
// The receiver’s critical value.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/criticalvalue
func (l_ LevelIndicator) SetCriticalValue(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCriticalValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/drawstieredcapacitylevels
func (l_ LevelIndicator) DrawsTieredCapacityLevels() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("drawsTieredCapacityLevels"))
	return rv
}


// SetDrawsTieredCapacityLevels sets the value of the drawsTieredCapacityLevels property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/drawstieredcapacitylevels
func (l_ LevelIndicator) SetDrawsTieredCapacityLevels(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDrawsTieredCapacityLevels:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/fillcolor
func (l_ LevelIndicator) FillColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("fillColor"))
	return rv
}


// SetFillColor sets the value of the fillColor property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/fillcolor
func (l_ LevelIndicator) SetFillColor(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setFillColor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/iseditable
func (l_ LevelIndicator) IsEditable() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isEditable"))
	return rv
}


// SetIsEditable sets the value of the isEditable property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/iseditable
func (l_ LevelIndicator) SetIsEditable(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsEditable:"), value)
}

// The appearance of the indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/levelindicatorstyle
func (l_ LevelIndicator) LevelIndicatorStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("levelIndicatorStyle"))
	return rv
}


// SetLevelIndicatorStyle sets the value of the levelIndicatorStyle property.
// The appearance of the indicator.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/levelindicatorstyle
func (l_ LevelIndicator) SetLevelIndicatorStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLevelIndicatorStyle:"), value)
}

// The receiver’s maximum value.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/maxvalue
func (l_ LevelIndicator) MaxValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("maxValue"))
	return rv
}


// SetMaxValue sets the value of the maxValue property.
// The receiver’s maximum value.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/maxvalue
func (l_ LevelIndicator) SetMaxValue(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMaxValue:"), value)
}

// The receiver’s minimum value.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/minvalue
func (l_ LevelIndicator) MinValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("minValue"))
	return rv
}


// SetMinValue sets the value of the minValue property.
// The receiver’s minimum value.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/minvalue
func (l_ LevelIndicator) SetMinValue(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMinValue:"), value)
}

// The number of major tick marks associated with the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/numberofmajortickmarks
func (l_ LevelIndicator) NumberOfMajorTickMarks() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfMajorTickMarks"))
	return rv
}


// SetNumberOfMajorTickMarks sets the value of the numberOfMajorTickMarks property.
// The number of major tick marks associated with the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/numberofmajortickmarks
func (l_ LevelIndicator) SetNumberOfMajorTickMarks(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfMajorTickMarks:"), value)
}

// The number of tick marks associated with the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/numberoftickmarks
func (l_ LevelIndicator) NumberOfTickMarks() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfTickMarks"))
	return rv
}


// SetNumberOfTickMarks sets the value of the numberOfTickMarks property.
// The number of tick marks associated with the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/numberoftickmarks
func (l_ LevelIndicator) SetNumberOfTickMarks(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfTickMarks:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/placeholdervisibility-swift.property
func (l_ LevelIndicator) PlaceholderVisibility() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("placeholderVisibility"))
	return rv
}


// SetPlaceholderVisibility sets the value of the placeholderVisibility property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/placeholdervisibility-swift.property
func (l_ LevelIndicator) SetPlaceholderVisibility(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPlaceholderVisibility:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/ratingimage
func (l_ LevelIndicator) RatingImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("ratingImage"))
	return rv
}


// SetRatingImage sets the value of the ratingImage property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/ratingimage
func (l_ LevelIndicator) SetRatingImage(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRatingImage:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/ratingplaceholderimage
func (l_ LevelIndicator) RatingPlaceholderImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("ratingPlaceholderImage"))
	return rv
}


// SetRatingPlaceholderImage sets the value of the ratingPlaceholderImage property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/ratingplaceholderimage
func (l_ LevelIndicator) SetRatingPlaceholderImage(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRatingPlaceholderImage:"), value)
}

// Determines how the receiver’s tick marks are aligned with it.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/tickmarkposition
func (l_ LevelIndicator) TickMarkPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("tickMarkPosition"))
	return rv
}


// SetTickMarkPosition sets the value of the tickMarkPosition property.
// Determines how the receiver’s tick marks are aligned with it.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/tickmarkposition
func (l_ LevelIndicator) SetTickMarkPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTickMarkPosition:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/warningfillcolor
func (l_ LevelIndicator) WarningFillColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("warningFillColor"))
	return rv
}


// SetWarningFillColor sets the value of the warningFillColor property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/warningfillcolor
func (l_ LevelIndicator) SetWarningFillColor(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWarningFillColor:"), value)
}

// The receiver’s warning value.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/warningvalue
func (l_ LevelIndicator) WarningValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("warningValue"))
	return rv
}


// SetWarningValue sets the value of the warningValue property.
// The receiver’s warning value.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/warningvalue
func (l_ LevelIndicator) SetWarningValue(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWarningValue:"), value)
}



