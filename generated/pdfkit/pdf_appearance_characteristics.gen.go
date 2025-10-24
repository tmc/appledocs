// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PDFAppearanceCharacteristics] class.
var (
	PDFAppearanceCharacteristicsClass     _PDFAppearanceCharacteristicsClass
	PDFAppearanceCharacteristicsClassOnce sync.Once
)

func getPDFAppearanceCharacteristicsClass() _PDFAppearanceCharacteristicsClass {
	PDFAppearanceCharacteristicsClassOnce.Do(func() {
		PDFAppearanceCharacteristicsClass = _PDFAppearanceCharacteristicsClass{objc.GetClass("PDFAppearanceCharacteristics")}
	})
	return PDFAppearanceCharacteristicsClass
}

type _PDFAppearanceCharacteristicsClass struct {
	class objc.Class
}

// An interface definition for the [PDFAppearanceCharacteristics] class.
type IPDFAppearanceCharacteristics interface {
	objectivec.IObject
	// properties:
	BackgroundColor() objc.IObject /* cross-framework: Color */
	SetBackgroundColor(value objc.IObject /* cross-framework: Color */)
	FieldName() objc.IObject /* cross-framework: NSString */
	SetFieldName(value objc.IObject /* cross-framework: NSString */)
	IsReadOnly() bool
	SetIsReadOnly(value bool)
	WidgetDefaultStringValue() objc.IObject /* cross-framework: NSString */
	SetWidgetDefaultStringValue(value objc.IObject /* cross-framework: NSString */)
	WidgetFieldType() PDFAnnotationWidgetSubtype /* typedef */
	SetWidgetFieldType(value PDFAnnotationWidgetSubtype /* typedef */)
	WidgetStringValue() objc.IObject /* cross-framework: NSString */
	SetWidgetStringValue(value objc.IObject /* cross-framework: NSString */)
	AppearanceCharacteristicsKeyValues() unsafe.Pointer
	SetAppearanceCharacteristicsKeyValues(value unsafe.Pointer)
	BorderColor() objc.IObject /* cross-framework: Color */
	SetBorderColor(value objc.IObject /* cross-framework: Color */)
	Caption() objc.IObject /* cross-framework: NSString */
	SetCaption(value objc.IObject /* cross-framework: NSString */)
	ControlType() PDFWidgetControlType
	SetControlType(value PDFWidgetControlType)
	DownCaption() objc.IObject /* cross-framework: NSString */
	SetDownCaption(value objc.IObject /* cross-framework: NSString */)
	RolloverCaption() objc.IObject /* cross-framework: NSString */
	SetRolloverCaption(value objc.IObject /* cross-framework: NSString */)
	Rotation() int
	SetRotation(value int)
	// methods:
}

// An object that represents appearance characteristics of a widget annotation.


// An object that represents appearance characteristics of a widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics
type PDFAppearanceCharacteristics struct {
	objectivec.Object
}

// PDFAppearanceCharacteristicsFrom constructs a [PDFAppearanceCharacteristics] from an unsafe.Pointer.
//
// An object that represents appearance characteristics of a widget annotation.
func PDFAppearanceCharacteristicsFrom(ptr unsafe.Pointer) PDFAppearanceCharacteristics {
	return PDFAppearanceCharacteristics{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAppearanceCharacteristicsClass) Alloc() PDFAppearanceCharacteristics {
	rv := objc.Send[PDFAppearanceCharacteristics](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAppearanceCharacteristicsClass) New() PDFAppearanceCharacteristics {
	rv := objc.Send[PDFAppearanceCharacteristics](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAppearanceCharacteristics) Init() PDFAppearanceCharacteristics {
	rv := objc.Send[PDFAppearanceCharacteristics](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAppearanceCharacteristics) Autorelease() PDFAppearanceCharacteristics {
	rv := objc.Send[PDFAppearanceCharacteristics](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAppearanceCharacteristics creates a new PDFAppearanceCharacteristics instance.
func NewPDFAppearanceCharacteristics() PDFAppearanceCharacteristics {
	return getPDFAppearanceCharacteristicsClass().New()
}



// The color of the widget’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/backgroundcolor
func (p_ PDFAppearanceCharacteristics) BackgroundColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The color of the widget’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/backgroundcolor
func (p_ PDFAppearanceCharacteristics) SetBackgroundColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The widget identifier for form annotation actions and behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/fieldname
func (p_ PDFAppearanceCharacteristics) FieldName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("fieldName"))
	return rv
}


// The widget identifier for form annotation actions and behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/fieldname
func (p_ PDFAppearanceCharacteristics) SetFieldName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFieldName:"), value)
}


// A Boolean value that determines whether the widget is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/isreadonly
func (p_ PDFAppearanceCharacteristics) IsReadOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isReadOnly"))
	return rv
}


// A Boolean value that determines whether the widget is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/isreadonly
func (p_ PDFAppearanceCharacteristics) SetIsReadOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsReadOnly:"), value)
}


// The string value that the widget reverts to when performing a reset form action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetdefaultstringvalue
func (p_ PDFAppearanceCharacteristics) WidgetDefaultStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("widgetDefaultStringValue"))
	return rv
}


// The string value that the widget reverts to when performing a reset form action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetdefaultstringvalue
func (p_ PDFAppearanceCharacteristics) SetWidgetDefaultStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetDefaultStringValue:"), value)
}


// The type of widget annotation, such as button, choice, or text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetfieldtype
func (p_ PDFAppearanceCharacteristics) WidgetFieldType() PDFAnnotationWidgetSubtype /* typedef */ {
	rv := objc.Send[PDFAnnotationWidgetSubtype](p_.ID, objc.Sel("widgetFieldType"))
	return rv
}


// The type of widget annotation, such as button, choice, or text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetfieldtype
func (p_ PDFAppearanceCharacteristics) SetWidgetFieldType(value PDFAnnotationWidgetSubtype /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetFieldType:"), value)
}


// The string value of the widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetstringvalue
func (p_ PDFAppearanceCharacteristics) WidgetStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("widgetStringValue"))
	return rv
}


// The string value of the widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetstringvalue
func (p_ PDFAppearanceCharacteristics) SetWidgetStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetStringValue:"), value)
}


// A dictionary that contains a deep copy of the appearance characteristic key-value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/appearancecharacteristicskeyvalues
func (p_ PDFAppearanceCharacteristics) AppearanceCharacteristicsKeyValues() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("appearanceCharacteristicsKeyValues"))
	return rv
}


// A dictionary that contains a deep copy of the appearance characteristic key-value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/appearancecharacteristicskeyvalues
func (p_ PDFAppearanceCharacteristics) SetAppearanceCharacteristicsKeyValues(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAppearanceCharacteristicsKeyValues:"), value)
}


// The border color of the widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/bordercolor
func (p_ PDFAppearanceCharacteristics) BorderColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("borderColor"))
	return rv
}


// The border color of the widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/bordercolor
func (p_ PDFAppearanceCharacteristics) SetBorderColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBorderColor:"), value)
}


// The text that the button widget annotation displays when the user isn’t interacting with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/caption
func (p_ PDFAppearanceCharacteristics) Caption() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("caption"))
	return rv
}


// The text that the button widget annotation displays when the user isn’t interacting with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/caption
func (p_ PDFAppearanceCharacteristics) SetCaption(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCaption:"), value)
}


// The type of button widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/controltype
func (p_ PDFAppearanceCharacteristics) ControlType() PDFWidgetControlType {
	rv := objc.Send[PDFWidgetControlType](p_.ID, objc.Sel("controlType"))
	return rv
}


// The type of button widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/controltype
func (p_ PDFAppearanceCharacteristics) SetControlType(value PDFWidgetControlType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlType:"), value)
}


// The text that the button widget annotation displays when the user holds down on it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/downcaption
func (p_ PDFAppearanceCharacteristics) DownCaption() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("downCaption"))
	return rv
}


// The text that the button widget annotation displays when the user holds down on it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/downcaption
func (p_ PDFAppearanceCharacteristics) SetDownCaption(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDownCaption:"), value)
}


// The text that the widget annotation displays when the user hovers the pointer over it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/rollovercaption
func (p_ PDFAppearanceCharacteristics) RolloverCaption() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("rolloverCaption"))
	return rv
}


// The text that the widget annotation displays when the user hovers the pointer over it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/rollovercaption
func (p_ PDFAppearanceCharacteristics) SetRolloverCaption(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRolloverCaption:"), value)
}


// The number of degrees, in multiples of 90, that the widget annotation rotates counterclockwise relative to the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/rotation
func (p_ PDFAppearanceCharacteristics) Rotation() int {
	rv := objc.Send[int](p_.ID, objc.Sel("rotation"))
	return rv
}


// The number of degrees, in multiples of 90, that the widget annotation rotates counterclockwise relative to the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/rotation
func (p_ PDFAppearanceCharacteristics) SetRotation(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRotation:"), value)
}



