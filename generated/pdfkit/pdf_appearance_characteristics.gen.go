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

/* debug [class.gen.go]: Generating class PDFAppearanceCharacteristics */


/* debug [class_header]: Header for PDFAppearanceCharacteristics */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAppearanceCharacteristics */
// An interface definition for the [PDFAppearanceCharacteristics] class.
type IPDFAppearanceCharacteristics interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PDFAppearanceCharacteristics */
	// properties:
	AppearanceCharacteristicsKeyValues() objc.IObject /* cross-framework: NSDictionary */
	BackgroundColor() appkit.Color
	SetBackgroundColor(value appkit.Color)
	BorderColor() appkit.Color
	SetBorderColor(value appkit.Color)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAppearanceCharacteristics */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAppearanceCharacteristics */
// Alloc allocates a new instance without initialization.
func (pc _PDFAppearanceCharacteristicsClass) Alloc() PDFAppearanceCharacteristics {
	rv := objc.Send[PDFAppearanceCharacteristics](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAppearanceCharacteristics */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAppearanceCharacteristics *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAppearanceCharacteristics */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAppearanceCharacteristics */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAppearanceCharacteristics */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAppearanceCharacteristics */

// A dictionary that contains a deep copy of the appearance characteristic key-value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/appearanceCharacteristicsKeyValues
func (p_ PDFAppearanceCharacteristics) AppearanceCharacteristicsKeyValues() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](p_.ID, objc.Sel("appearanceCharacteristicsKeyValues"))
	return rv
}/* debug [instance_properties/getter]: appearanceCharacteristicsKeyValues */


// The background color of the widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/backgroundColor
func (p_ PDFAppearanceCharacteristics) BackgroundColor() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The background color of the widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/backgroundColor
func (p_ PDFAppearanceCharacteristics) SetBackgroundColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The border color of the widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/borderColor
func (p_ PDFAppearanceCharacteristics) BorderColor() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("borderColor"))
	return rv
}/* debug [instance_properties/getter]: borderColor */


// The border color of the widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/borderColor
func (p_ PDFAppearanceCharacteristics) SetBorderColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBorderColor:"), value)
}/* debug [instance_properties/setter]: borderColor */


// The text that the button widget annotation displays when the user isn’t interacting with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/caption
func (p_ PDFAppearanceCharacteristics) Caption() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("caption"))
	return rv
}/* debug [instance_properties/getter]: caption */


// The text that the button widget annotation displays when the user isn’t interacting with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/caption
func (p_ PDFAppearanceCharacteristics) SetCaption(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCaption:"), value)
}/* debug [instance_properties/setter]: caption */


// The type of button widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/controlType
func (p_ PDFAppearanceCharacteristics) ControlType() PDFWidgetControlType {
	rv := objc.Send[PDFWidgetControlType](p_.ID, objc.Sel("controlType"))
	return rv
}/* debug [instance_properties/getter]: controlType */


// The type of button widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/controlType
func (p_ PDFAppearanceCharacteristics) SetControlType(value PDFWidgetControlType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlType:"), value)
}/* debug [instance_properties/setter]: controlType */


// The text that the button widget annotation displays when the user holds down on it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/downCaption
func (p_ PDFAppearanceCharacteristics) DownCaption() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("downCaption"))
	return rv
}/* debug [instance_properties/getter]: downCaption */


// The text that the button widget annotation displays when the user holds down on it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/downCaption
func (p_ PDFAppearanceCharacteristics) SetDownCaption(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDownCaption:"), value)
}/* debug [instance_properties/setter]: downCaption */


// The text that the widget annotation displays when the user hovers the pointer over it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/rolloverCaption
func (p_ PDFAppearanceCharacteristics) RolloverCaption() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("rolloverCaption"))
	return rv
}/* debug [instance_properties/getter]: rolloverCaption */


// The text that the widget annotation displays when the user hovers the pointer over it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/rolloverCaption
func (p_ PDFAppearanceCharacteristics) SetRolloverCaption(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRolloverCaption:"), value)
}/* debug [instance_properties/setter]: rolloverCaption */


// The number of degrees, in multiples of 90, that the widget annotation rotates counterclockwise relative to the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/rotation
func (p_ PDFAppearanceCharacteristics) Rotation() int {
	rv := objc.Send[int](p_.ID, objc.Sel("rotation"))
	return rv
}/* debug [instance_properties/getter]: rotation */


// The number of degrees, in multiples of 90, that the widget annotation rotates counterclockwise relative to the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/rotation
func (p_ PDFAppearanceCharacteristics) SetRotation(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRotation:"), value)
}/* debug [instance_properties/setter]: rotation */


// The widget identifier for form annotation actions and behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/fieldname
func (p_ PDFAppearanceCharacteristics) FieldName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("fieldName"))
	return rv
}/* debug [instance_properties/getter]: fieldName */


// The widget identifier for form annotation actions and behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/fieldname
func (p_ PDFAppearanceCharacteristics) SetFieldName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFieldName:"), value)
}/* debug [instance_properties/setter]: fieldName */


// A Boolean value that determines whether the widget is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/isreadonly
func (p_ PDFAppearanceCharacteristics) IsReadOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isReadOnly"))
	return rv
}/* debug [instance_properties/getter]: isReadOnly */


// A Boolean value that determines whether the widget is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/isreadonly
func (p_ PDFAppearanceCharacteristics) SetIsReadOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsReadOnly:"), value)
}/* debug [instance_properties/setter]: isReadOnly */


// The string value that the widget reverts to when performing a reset form action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetdefaultstringvalue
func (p_ PDFAppearanceCharacteristics) WidgetDefaultStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("widgetDefaultStringValue"))
	return rv
}/* debug [instance_properties/getter]: widgetDefaultStringValue */


// The string value that the widget reverts to when performing a reset form action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetdefaultstringvalue
func (p_ PDFAppearanceCharacteristics) SetWidgetDefaultStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetDefaultStringValue:"), value)
}/* debug [instance_properties/setter]: widgetDefaultStringValue */


// The type of widget annotation, such as button, choice, or text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetfieldtype
func (p_ PDFAppearanceCharacteristics) WidgetFieldType() PDFAnnotationWidgetSubtype /* typedef */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("widgetFieldType"))
	return rv
}/* debug [instance_properties/getter]: widgetFieldType */


// The type of widget annotation, such as button, choice, or text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetfieldtype
func (p_ PDFAppearanceCharacteristics) SetWidgetFieldType(value PDFAnnotationWidgetSubtype /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetFieldType:"), value)
}/* debug [instance_properties/setter]: widgetFieldType */


// The string value of the widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetstringvalue
func (p_ PDFAppearanceCharacteristics) WidgetStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("widgetStringValue"))
	return rv
}/* debug [instance_properties/getter]: widgetStringValue */


// The string value of the widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetstringvalue
func (p_ PDFAppearanceCharacteristics) SetWidgetStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetStringValue:"), value)
}/* debug [instance_properties/setter]: widgetStringValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAppearanceCharacteristics */



