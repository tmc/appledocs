// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [ListFormatter] class.
type IListFormatter interface {
	IFormatter
	ItemFormatter() NSFormatter
	SetItemFormatter(value IFormatter)
	Locale() Locale
	SetLocale(value ILocale)
}

// An object that provides locale-correct formatting of a list of items using the appropriate separator and conjunction.
//
// The list formatter isn’t aware of the context where the formatted string will be used and doesn’t provide capitalization customization of the list items. The formatted result may not be grammatically correct if placed in a sentence, and it should only be used in a standalone manner.
//
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

// Alloc allocates a new instance without initialization.
func (lc _ListFormatterClass) Alloc() ListFormatter {
	rv := objc.Send[ListFormatter](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// An object that formats each item in the list.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/listformatter/itemformatter
func (l_ ListFormatter) ItemFormatter() NSFormatter {
	rv := objc.Send[NSFormatter](l_.ID, objc.Sel("itemFormatter"))
	return rv
}


// SetItemFormatter sets the value of the itemFormatter property.
// An object that formats each item in the list.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/listformatter/itemformatter
func (l_ ListFormatter) SetItemFormatter(value IFormatter) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setItemFormatter:"), value)
}

// The locale to use when formatting items in the list.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/listformatter/locale
func (l_ ListFormatter) Locale() Locale {
	rv := objc.Send[Locale](l_.ID, objc.Sel("locale"))
	return rv
}


// SetLocale sets the value of the locale property.
// The locale to use when formatting items in the list.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/listformatter/locale
func (l_ ListFormatter) SetLocale(value ILocale) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLocale:"), value)
}



