// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An object that represents appearance characteristics of a widget annotation.
//
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
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/backgroundcolor
func (p_ PDFAppearanceCharacteristics) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The color of the widget’s background.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/backgroundcolor
func (p_ PDFAppearanceCharacteristics) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBackgroundColor:"), value)
}

// The widget identifier for form annotation actions and behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/fieldname
func (p_ PDFAppearanceCharacteristics) FieldName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("fieldName"))
	return rv
}


// SetFieldName sets the value of the fieldName property.
// The widget identifier for form annotation actions and behaviors.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/fieldname
func (p_ PDFAppearanceCharacteristics) SetFieldName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFieldName:"), objc.String(value))
}

// A Boolean value that determines whether the widget is editable.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/isreadonly
func (p_ PDFAppearanceCharacteristics) IsReadOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isReadOnly"))
	return rv
}


// SetIsReadOnly sets the value of the isReadOnly property.
// A Boolean value that determines whether the widget is editable.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/isreadonly
func (p_ PDFAppearanceCharacteristics) SetIsReadOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsReadOnly:"), value)
}

// The string value that the widget reverts to when performing a reset form action.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetdefaultstringvalue
func (p_ PDFAppearanceCharacteristics) WidgetDefaultStringValue() string {
	rv := objc.Send[string](p_.ID, objc.Sel("widgetDefaultStringValue"))
	return rv
}


// SetWidgetDefaultStringValue sets the value of the widgetDefaultStringValue property.
// The string value that the widget reverts to when performing a reset form action.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetdefaultstringvalue
func (p_ PDFAppearanceCharacteristics) SetWidgetDefaultStringValue(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetDefaultStringValue:"), objc.String(value))
}

// The type of widget annotation, such as button, choice, or text.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetfieldtype
func (p_ PDFAppearanceCharacteristics) WidgetFieldType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("widgetFieldType"))
	return rv
}


// SetWidgetFieldType sets the value of the widgetFieldType property.
// The type of widget annotation, such as button, choice, or text.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetfieldtype
func (p_ PDFAppearanceCharacteristics) SetWidgetFieldType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetFieldType:"), value)
}

// The string value of the widget annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetstringvalue
func (p_ PDFAppearanceCharacteristics) WidgetStringValue() string {
	rv := objc.Send[string](p_.ID, objc.Sel("widgetStringValue"))
	return rv
}


// SetWidgetStringValue sets the value of the widgetStringValue property.
// The string value of the widget annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetstringvalue
func (p_ PDFAppearanceCharacteristics) SetWidgetStringValue(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetStringValue:"), objc.String(value))
}

// A dictionary that contains a deep copy of the appearance characteristic key-value pairs.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/appearancecharacteristicskeyvalues
func (p_ PDFAppearanceCharacteristics) AppearanceCharacteristicsKeyValues() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("appearanceCharacteristicsKeyValues"))
	return rv
}


// SetAppearanceCharacteristicsKeyValues sets the value of the appearanceCharacteristicsKeyValues property.
// A dictionary that contains a deep copy of the appearance characteristic key-value pairs.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/appearancecharacteristicskeyvalues
func (p_ PDFAppearanceCharacteristics) SetAppearanceCharacteristicsKeyValues(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAppearanceCharacteristicsKeyValues:"), value)
}

// The border color of the widget annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/bordercolor
func (p_ PDFAppearanceCharacteristics) BorderColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("borderColor"))
	return rv
}


// SetBorderColor sets the value of the borderColor property.
// The border color of the widget annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/bordercolor
func (p_ PDFAppearanceCharacteristics) SetBorderColor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBorderColor:"), value)
}

// The text that the button widget annotation displays when the user isn’t interacting with it.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/caption
func (p_ PDFAppearanceCharacteristics) Caption() string {
	rv := objc.Send[string](p_.ID, objc.Sel("caption"))
	return rv
}


// SetCaption sets the value of the caption property.
// The text that the button widget annotation displays when the user isn’t interacting with it.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/caption
func (p_ PDFAppearanceCharacteristics) SetCaption(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCaption:"), objc.String(value))
}

// The type of button widget annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/controltype
func (p_ PDFAppearanceCharacteristics) ControlType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("controlType"))
	return rv
}


// SetControlType sets the value of the controlType property.
// The type of button widget annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/controltype
func (p_ PDFAppearanceCharacteristics) SetControlType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlType:"), value)
}

// The text that the button widget annotation displays when the user holds down on it.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/downcaption
func (p_ PDFAppearanceCharacteristics) DownCaption() string {
	rv := objc.Send[string](p_.ID, objc.Sel("downCaption"))
	return rv
}


// SetDownCaption sets the value of the downCaption property.
// The text that the button widget annotation displays when the user holds down on it.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/downcaption
func (p_ PDFAppearanceCharacteristics) SetDownCaption(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDownCaption:"), objc.String(value))
}

// The text that the widget annotation displays when the user hovers the pointer over it.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/rollovercaption
func (p_ PDFAppearanceCharacteristics) RolloverCaption() string {
	rv := objc.Send[string](p_.ID, objc.Sel("rolloverCaption"))
	return rv
}


// SetRolloverCaption sets the value of the rolloverCaption property.
// The text that the widget annotation displays when the user hovers the pointer over it.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/rollovercaption
func (p_ PDFAppearanceCharacteristics) SetRolloverCaption(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRolloverCaption:"), objc.String(value))
}

// The number of degrees, in multiples of 90, that the widget annotation rotates counterclockwise relative to the page.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/rotation
func (p_ PDFAppearanceCharacteristics) Rotation() int {
	rv := objc.Send[int](p_.ID, objc.Sel("rotation"))
	return rv
}


// SetRotation sets the value of the rotation property.
// The number of degrees, in multiples of 90, that the widget annotation rotates counterclockwise relative to the page.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfappearancecharacteristics/rotation
func (p_ PDFAppearanceCharacteristics) SetRotation(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRotation:"), value)
}



