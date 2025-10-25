// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSListFormatter */


/* debug [class_header]: Header for NSListFormatter */
// The class instance for the [ListFormatter] class.
var (
	ListFormatterClass     _ListFormatterClass
	ListFormatterClassOnce sync.Once
)

func getListFormatterClass() _ListFormatterClass {
	ListFormatterClassOnce.Do(func() {
		ListFormatterClass = _ListFormatterClass{objc.GetClass("NSListFormatter")}
	})
	return ListFormatterClass
}

type _ListFormatterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ListFormatter */
// An interface definition for the [ListFormatter] class.
type IListFormatter interface {
	IFormatter
	
/* debug [class_interface_properties]: Properties for ListFormatter */
	// properties:
	ItemFormatter() IFormatter
	SetItemFormatter(value IFormatter)
	Locale() ILocale
	SetLocale(value ILocale)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ListFormatter */
	// methods:
	StringForObjectValue(obj objc.IObject) IString
	StringFromItems(items IArray) IString
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ListFormatter */
// Alloc allocates a new instance without initialization.
func (lc _ListFormatterClass) Alloc() ListFormatter {
	rv := objc.Send[ListFormatter](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _ListFormatterClass) New() ListFormatter {
	rv := objc.Send[ListFormatter](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ ListFormatter) Init() ListFormatter {
	rv := objc.Send[ListFormatter](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ ListFormatter) Autorelease() ListFormatter {
	rv := objc.Send[ListFormatter](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewListFormatter creates a new ListFormatter instance.
func NewListFormatter() ListFormatter {
	return getListFormatterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ListFormatter */
// An object that provides locale-correct formatting of a list of items using the appropriate separator and conjunction.
//
// The list formatter isn’t aware of the context where the formatted string will be used and doesn’t provide capitalization customization of the list items. The formatted result may not be grammatically correct if placed in a sentence, and it should only be used in a standalone manner.


// An object that provides locale-correct formatting of a list of items using the appropriate separator and conjunction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ListFormatter
type ListFormatter struct {
	Formatter
}

// ListFormatterFrom constructs a [ListFormatter] from an unsafe.Pointer.
//
// An object that provides locale-correct formatting of a list of items using the appropriate separator and conjunction.
func ListFormatterFrom(ptr unsafe.Pointer) ListFormatter {
	return ListFormatter{
		Formatter: FormatterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ListFormatter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ListFormatter */

// Constructs a formatted string from an array of strings that uses the list format specific to the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ListFormatter/localizedString(byJoining:)
func (lc _ListFormatterClass) LocalizedStringByJoiningStrings(strings []string) IString {
	rv := objc.Send[String](objc.ID(lc.class), objc.Sel("localizedStringByJoiningStrings:"), strings)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalizedStringByJoiningStrings) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ListFormatter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ListFormatter */

// Creates a formatted string for an array of items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ListFormatter/string(for:)
func (l_ ListFormatter) StringForObjectValue(obj objc.IObject) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("stringForObjectValue:"), obj)
	return rv
}/* debug [instance_methods/method]: StringForObjectValue */


// Creates a formatted string for an array of items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ListFormatter/string(from:)
func (l_ ListFormatter) StringFromItems(items IArray) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("stringFromItems:"), items)
	return rv
}/* debug [instance_methods/method]: StringFromItems */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ListFormatter */

// An object that formats each item in the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ListFormatter/itemFormatter
func (l_ ListFormatter) ItemFormatter() IFormatter {
	rv := objc.Send[Formatter](l_.ID, objc.Sel("itemFormatter"))
	return rv
}/* debug [instance_properties/getter]: itemFormatter */


// An object that formats each item in the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ListFormatter/itemFormatter
func (l_ ListFormatter) SetItemFormatter(value IFormatter) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setItemFormatter:"), value)
}/* debug [instance_properties/setter]: itemFormatter */


// The locale to use when formatting items in the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ListFormatter/locale
func (l_ ListFormatter) Locale() ILocale {
	rv := objc.Send[Locale](l_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// The locale to use when formatting items in the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ListFormatter/locale
func (l_ ListFormatter) SetLocale(value ILocale) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLocale:"), value)
}/* debug [instance_properties/setter]: locale */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSListFormatter */



