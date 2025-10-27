// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureIndexPicker] class.
var (
	CaptureIndexPickerClass     _CaptureIndexPickerClass
	CaptureIndexPickerClassOnce sync.Once
)

func getCaptureIndexPickerClass() _CaptureIndexPickerClass {
	CaptureIndexPickerClassOnce.Do(func() {
		CaptureIndexPickerClass = _CaptureIndexPickerClass{objc.GetClass("AVCaptureIndexPicker")}
	})
	return CaptureIndexPickerClass
}

type _CaptureIndexPickerClass struct {
	class objc.Class
}





// An interface definition for the [CaptureIndexPicker] class.
type ICaptureIndexPicker interface {
	ICaptureControl
	

	// properties:
	AccessibilityIdentifier() foundation.foundation.INSString
	SetAccessibilityIdentifier(value foundation.foundation.INSString)
	LocalizedIndexTitles() []string
	LocalizedTitle() foundation.foundation.INSString
	NumberOfIndexes() int
	SelectedIndex() int
	SetSelectedIndex(value int)
	SymbolName() foundation.foundation.INSString


	

	// methods:
	SetActionQueueAction(actionQueue objectivec.IObject, action int)


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureIndexPickerClass) Alloc() CaptureIndexPicker {
	rv := objc.Send[CaptureIndexPicker](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureIndexPickerClass) New() CaptureIndexPicker {
	rv := objc.Send[CaptureIndexPicker](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureIndexPicker) Init() CaptureIndexPicker {
	rv := objc.Send[CaptureIndexPicker](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureIndexPicker) Autorelease() CaptureIndexPicker {
	rv := objc.Send[CaptureIndexPicker](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureIndexPicker creates a new CaptureIndexPicker instance.
func NewCaptureIndexPicker() CaptureIndexPicker {
	return getCaptureIndexPickerClass().New()
}





// A control for selecting from a set of mutually exclusive values by index.
//
// Index pickers are appropriate for controls that provide an indexed container of values.


// A control for selecting from a set of mutually exclusive values by index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker
type CaptureIndexPicker struct {
	CaptureControl
}

// CaptureIndexPickerFrom constructs a [CaptureIndexPicker] from an unsafe.Pointer.
//
// A control for selecting from a set of mutually exclusive values by index.
func CaptureIndexPickerFrom(ptr unsafe.Pointer) CaptureIndexPicker {
	return CaptureIndexPicker{
		CaptureControl: CaptureControlFrom(ptr),
	}
}






// Creates an object to select an index from a set of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/init(_:symbolName:localizedIndexTitles:)
func NewCaptureIndexPickerWithLocalizedTitleSymbolNameLocalizedIndexTitles(localizedTitle foundation.foundation.INSString, symbolName foundation.foundation.INSString, localizedIndexTitles []string) CaptureIndexPicker {
	instance := getCaptureIndexPickerClass().Alloc()
	rv := objc.Send[CaptureIndexPicker](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:localizedIndexTitles:"), localizedTitle, symbolName, localizedIndexTitles)
	rv.Autorelease()
	return rv
}


// Creates a control to pick a value from the specified number of indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/init(_:symbolName:numberOfIndexes:)
func NewCaptureIndexPickerWithLocalizedTitleSymbolNameNumberOfIndexes(localizedTitle foundation.foundation.INSString, symbolName foundation.foundation.INSString, numberOfIndexes int) CaptureIndexPicker {
	instance := getCaptureIndexPickerClass().Alloc()
	rv := objc.Send[CaptureIndexPicker](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:numberOfIndexes:"), localizedTitle, symbolName, numberOfIndexes)
	rv.Autorelease()
	return rv
}


// Creates a control to pick a value from the specified number of indices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/init(_:symbolName:numberOfIndexes:localizedTitleTransform:)
func NewCaptureIndexPickerWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransform(localizedTitle foundation.foundation.INSString, symbolName foundation.foundation.INSString, numberOfIndexes int, localizedTitleTransform unsafe.Pointer) CaptureIndexPicker {
	instance := getCaptureIndexPickerClass().Alloc()
	rv := objc.Send[CaptureIndexPicker](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:numberOfIndexes:localizedTitleTransform:"), localizedTitle, symbolName, numberOfIndexes, localizedTitleTransform)
	rv.Autorelease()
	return rv
}

















// Sets the action to perform on the specified dispatch queue when the control’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/setActionQueue:action:
func (c_ CaptureIndexPicker) SetActionQueueAction(actionQueue objectivec.IObject, action int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActionQueue:action:"), actionQueue, action)
}







// A string identifier for this control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/accessibilityIdentifier
func (c_ CaptureIndexPicker) AccessibilityIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("accessibilityIdentifier"))
	return rv
}


// A string identifier for this control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/accessibilityIdentifier
func (c_ CaptureIndexPicker) SetAccessibilityIdentifier(value foundation.foundation.INSString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccessibilityIdentifier:"), value)
}


// The titles to present for each index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/localizedIndexTitles
func (c_ CaptureIndexPicker) LocalizedIndexTitles() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("localizedIndexTitles"))
	return rv
}


// A localized title that describes the control’s action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/localizedTitle
func (c_ CaptureIndexPicker) LocalizedTitle() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedTitle"))
	return rv
}


// The number of index values the control provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/numberOfIndexes
func (c_ CaptureIndexPicker) NumberOfIndexes() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfIndexes"))
	return rv
}


// The currently selected index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/selectedIndex
func (c_ CaptureIndexPicker) SelectedIndex() int {
	rv := objc.Send[int](c_.ID, objc.Sel("selectedIndex"))
	return rv
}


// The currently selected index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/selectedIndex
func (c_ CaptureIndexPicker) SetSelectedIndex(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelectedIndex:"), value)
}


// The name of the SF Symbol that represents this control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/symbolName
func (c_ CaptureIndexPicker) SymbolName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("symbolName"))
	return rv
}







