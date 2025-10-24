// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureIndexPicker */


/* debug [class_header]: Header for AVCaptureIndexPicker */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureIndexPicker */
// An interface definition for the [CaptureIndexPicker] class.
type ICaptureIndexPicker interface {
	ICaptureControl
	
/* debug [class_interface_properties]: Properties for CaptureIndexPicker */
	// properties:
	AccessibilityIdentifier() objc.IObject /* cross-framework: NSString */
	SetAccessibilityIdentifier(value objc.IObject /* cross-framework: NSString */)
	LocalizedIndexTitles() []string
	LocalizedTitle() objc.IObject /* cross-framework: NSString */
	NumberOfIndexes() int
	SelectedIndex() int
	SetSelectedIndex(value int)
	SymbolName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureIndexPicker */
	// methods:
	SetActionQueueAction(actionQueue objectivec.IObject, action int)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureIndexPicker */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureIndexPicker */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureIndexPicker */

// Creates an object to select an index from a set of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/init(_:symbolName:localizedIndexTitles:)
func NewCaptureIndexPickerWithLocalizedTitleSymbolNameLocalizedIndexTitles(localizedTitle objc.IObject /* cross-framework: NSString */, symbolName objc.IObject /* cross-framework: NSString */, localizedIndexTitles []string) CaptureIndexPicker {
	instance := getCaptureIndexPickerClass().Alloc()
	rv := objc.Send[CaptureIndexPicker](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:localizedIndexTitles:"), localizedTitle, symbolName, localizedIndexTitles)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureIndexPickerWithLocalizedTitleSymbolNameLocalizedIndexTitles */


// Creates a control to pick a value from the specified number of indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/init(_:symbolName:numberOfIndexes:)
func NewCaptureIndexPickerWithLocalizedTitleSymbolNameNumberOfIndexes(localizedTitle objc.IObject /* cross-framework: NSString */, symbolName objc.IObject /* cross-framework: NSString */, numberOfIndexes int) CaptureIndexPicker {
	instance := getCaptureIndexPickerClass().Alloc()
	rv := objc.Send[CaptureIndexPicker](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:numberOfIndexes:"), localizedTitle, symbolName, numberOfIndexes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureIndexPickerWithLocalizedTitleSymbolNameNumberOfIndexes */


// Creates a control to pick a value from the specified number of indices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/init(_:symbolName:numberOfIndexes:localizedTitleTransform:)
func NewCaptureIndexPickerWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransform(localizedTitle objc.IObject /* cross-framework: NSString */, symbolName objc.IObject /* cross-framework: NSString */, numberOfIndexes int, localizedTitleTransform unsafe.Pointer) CaptureIndexPicker {
	instance := getCaptureIndexPickerClass().Alloc()
	rv := objc.Send[CaptureIndexPicker](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:numberOfIndexes:localizedTitleTransform:"), localizedTitle, symbolName, numberOfIndexes, localizedTitleTransform)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureIndexPickerWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransform */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureIndexPicker */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureIndexPicker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureIndexPicker */

// Sets the action to perform on the specified dispatch queue when the control’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/setActionQueue:action:
func (c_ CaptureIndexPicker) SetActionQueueAction(actionQueue objectivec.IObject, action int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActionQueue:action:"), actionQueue, action)
}/* debug [instance_methods/method]: SetActionQueueAction */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureIndexPicker */

// A string identifier for this control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/accessibilityIdentifier
func (c_ CaptureIndexPicker) AccessibilityIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("accessibilityIdentifier"))
	return rv
}/* debug [instance_properties/getter]: accessibilityIdentifier */


// A string identifier for this control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/accessibilityIdentifier
func (c_ CaptureIndexPicker) SetAccessibilityIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccessibilityIdentifier:"), value)
}/* debug [instance_properties/setter]: accessibilityIdentifier */


// The titles to present for each index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/localizedIndexTitles
func (c_ CaptureIndexPicker) LocalizedIndexTitles() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("localizedIndexTitles"))
	return rv
}/* debug [instance_properties/getter]: localizedIndexTitles */


// A localized title that describes the control’s action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/localizedTitle
func (c_ CaptureIndexPicker) LocalizedTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedTitle"))
	return rv
}/* debug [instance_properties/getter]: localizedTitle */


// The number of index values the control provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/numberOfIndexes
func (c_ CaptureIndexPicker) NumberOfIndexes() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfIndexes"))
	return rv
}/* debug [instance_properties/getter]: numberOfIndexes */


// The currently selected index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/selectedIndex
func (c_ CaptureIndexPicker) SelectedIndex() int {
	rv := objc.Send[int](c_.ID, objc.Sel("selectedIndex"))
	return rv
}/* debug [instance_properties/getter]: selectedIndex */


// The currently selected index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/selectedIndex
func (c_ CaptureIndexPicker) SetSelectedIndex(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelectedIndex:"), value)
}/* debug [instance_properties/setter]: selectedIndex */


// The name of the SF Symbol that represents this control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/symbolName
func (c_ CaptureIndexPicker) SymbolName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("symbolName"))
	return rv
}/* debug [instance_properties/getter]: symbolName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureIndexPicker */


