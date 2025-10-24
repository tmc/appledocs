// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CSLocalizedString */


/* debug [class_header]: Header for CSLocalizedString */
// The class instance for the [CSLocalizedString] class.
var (
	CSLocalizedStringClass     _CSLocalizedStringClass
	CSLocalizedStringClassOnce sync.Once
)

func getCSLocalizedStringClass() _CSLocalizedStringClass {
	CSLocalizedStringClassOnce.Do(func() {
		CSLocalizedStringClass = _CSLocalizedStringClass{objc.GetClass("CSLocalizedString")}
	})
	return CSLocalizedStringClass
}

type _CSLocalizedStringClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSLocalizedString */
// An interface definition for the [CSLocalizedString] class.
type ICSLocalizedString interface {
	IString
	
/* debug [class_interface_properties]: Properties for CSLocalizedString */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSLocalizedString */
	// methods:
	LocalizedString() objc.IObject /* cross-framework: String */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSLocalizedString */
// Alloc allocates a new instance without initialization.
func (cc _CSLocalizedStringClass) Alloc() CSLocalizedString {
	rv := objc.Send[CSLocalizedString](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSLocalizedStringClass) New() CSLocalizedString {
	rv := objc.Send[CSLocalizedString](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSLocalizedString) Init() CSLocalizedString {
	rv := objc.Send[CSLocalizedString](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSLocalizedString) Autorelease() CSLocalizedString {
	rv := objc.Send[CSLocalizedString](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSLocalizedString creates a new CSLocalizedString instance.
func NewCSLocalizedString() CSLocalizedString {
	return getCSLocalizedStringClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSLocalizedString */
// An object that displays localized text in search results related to your app.
//
// The class helps you localize text in searchable items. You can use a object in place of an object to display localized text in search results related to your app. For example, you might use the following code to define a object for a searchable item you want to identify as “Song” in English:


// An object that displays localized text in search results related to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSLocalizedString
type CSLocalizedString struct {
	string
}

// CSLocalizedStringFrom constructs a [CSLocalizedString] from an unsafe.Pointer.
//
// An object that displays localized text in search results related to your app.
func CSLocalizedStringFrom(ptr unsafe.Pointer) CSLocalizedString {
	return CSLocalizedString{
		String: stringFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSLocalizedString */

// Initializes a object with the specified dictionary of localized strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSLocalizedString/init(localizedStrings:)
func NewCSLocalizedStringWithLocalizedStrings(localizedStrings objc.IObject /* cross-framework: NSDictionary */) CSLocalizedString {
	instance := getCSLocalizedStringClass().Alloc()
	rv := objc.Send[CSLocalizedString](instance.ID, objc.Sel("initWithLocalizedStrings:"), localizedStrings)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCSLocalizedStringWithLocalizedStrings */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSLocalizedString */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSLocalizedString */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSLocalizedString */

// Returns the localized string for the current language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSLocalizedString/localizedString()
func (c_ CSLocalizedString) LocalizedString() objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](c_.ID, objc.Sel("localizedString"))
	return rv
}/* debug [instance_methods/method]: LocalizedString */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSLocalizedString */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CSLocalizedString */


