// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersonNameComponentsFormatter] class.
var (
	PersonNameComponentsFormatterClass     _PersonNameComponentsFormatterClass
	PersonNameComponentsFormatterClassOnce sync.Once
)

func getPersonNameComponentsFormatterClass() _PersonNameComponentsFormatterClass {
	PersonNameComponentsFormatterClassOnce.Do(func() {
		PersonNameComponentsFormatterClass = _PersonNameComponentsFormatterClass{objc.GetClass("NSPersonNameComponentsFormatter")}
	})
	return PersonNameComponentsFormatterClass
}

type _PersonNameComponentsFormatterClass struct {
	class objc.Class
}

// An interface definition for the [PersonNameComponentsFormatter] class.
type IPersonNameComponentsFormatter interface {
	IFormatter
	PersonNameComponentsFromString(string_ string) unsafe.Pointer
}

// A formatter that provides localized representations of the components of a person’s name.
//
// Each locale has its own set of rules and conventions for how personal names are structured and represented. These rules vary widely across different locales in a several ways, including the sort and display order of given and family names, the use of salutations and honorifics, and other concerns related to the grammar, spelling, punctuation, and formatting. About the only thing that consistent across all locales is that personal names are significant and meaningful. For this reason, names deserve careful and respectful treatment—perhaps more than any other kind of information your app interacts with. Formatters can be configured to represent names in a variety of styles, which are described in detail below. Default ( ) Short ( ) Long ( ) Abbreviated ( ) When determining how to represent a name in a particular style, a formatter takes a number of factors into consideration, in order of priority: Scripts may specify a strict sort or display order of given and family names, and the availability of styles. Users can enable and configure the display of short names, as well as whether or not to display nicknames when available. Users can also override the default sort and display order of given and family names for their current locale. Locales specify a default sort and display order for given and family names. The style property value set for the object. When the behavior specified in one factor conflicts with any other factors, the behavior specified by the factor with the most precedence is used. For example, the U.S. English ( ) locale specifies that names be displayed in “given name followed by the family name” (for example,“John Appleseed”). This behavior would be overridden if the user changed their system preferences to have names displayed as family name followed by given name (for example, “Appleseed, John”), because user-specified preferences take precedence over locale-derived defaults. Furthermore, if the name to be formatted were Japanese (for example, given name: “泰夫”, family name: “木田”), the behavior derived for the name’s script (CJK, for Chinese, Japanese, and Korean languages) would take precedence over any locale-derived defaults or user-specified preferences to have the name displayed as family name followed by given name (for example, “木田 泰夫”). These considerations extend to the availability of certain formatter styles as well. Because developer-specified configurations have the lowest precedence in determining behavior, the value set for the formatter’s style property can be invalidated if it’s not supported for the locale, user preferences, or script. If the specified style is not available, the next longest valid style is used. For example, a name in Arabic script (for example, “أحمد الراجحي”) does not support the Abbreviated style, so the Short style is used instead. A name that contains more than one script (for example, given name: “John”, family name: “王”) is detected to have “Unknown” script, which has its own set of behaviors and characteristics.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter
type PersonNameComponentsFormatter struct {
	Formatter
}

// PersonNameComponentsFormatterFrom constructs a [PersonNameComponentsFormatter] from an unsafe.Pointer.
//
// A formatter that provides localized representations of the components of a person’s name.
func PersonNameComponentsFormatterFrom(ptr unsafe.Pointer) PersonNameComponentsFormatter {
	return PersonNameComponentsFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PersonNameComponentsFormatterClass) Alloc() PersonNameComponentsFormatter {
	rv := objc.Send[PersonNameComponentsFormatter](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersonNameComponentsFormatterClass) New() PersonNameComponentsFormatter {
	rv := objc.Send[PersonNameComponentsFormatter](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersonNameComponentsFormatter) Init() PersonNameComponentsFormatter {
	rv := objc.Send[PersonNameComponentsFormatter](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersonNameComponentsFormatter) Autorelease() PersonNameComponentsFormatter {
	rv := objc.Send[PersonNameComponentsFormatter](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersonNameComponentsFormatter creates a new PersonNameComponentsFormatter instance.
func NewPersonNameComponentsFormatter() PersonNameComponentsFormatter {
	return getPersonNameComponentsFormatterClass().New()
}


// Returns a person name components object from a given string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/personNameComponents(from:)
func (p_ PersonNameComponentsFormatter) PersonNameComponentsFromString(string_ string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("personNameComponentsFromString:"), objc.String(string_))
	return rv
}



