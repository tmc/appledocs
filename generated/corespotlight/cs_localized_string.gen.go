// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [CSLocalizedString] class.
type ICSLocalizedString interface {
	foundation.IString
	LocalizedString() string
}

// An object that displays localized text in search results related to your app.
//
// The class helps you localize text in searchable items. You can use a object in place of an object to display localized text in search results related to your app. For example, you might use the following code to define a object for a searchable item you want to identify as “Song” in English:
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSLocalizedString
type CSLocalizedString struct {
	foundation.String
}

// CSLocalizedStringFrom constructs a [CSLocalizedString] from an unsafe.Pointer.
//
// An object that displays localized text in search results related to your app.
func CSLocalizedStringFrom(ptr unsafe.Pointer) CSLocalizedString {
	return CSLocalizedString{
		String: foundation.StringFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CSLocalizedStringClass) Alloc() CSLocalizedString {
	rv := objc.Send[CSLocalizedString](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Initializes a object with the specified dictionary of localized strings.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSLocalizedString/init(localizedStrings:)
func NewCSLocalizedStringWithLocalizedStrings(localizedStrings objc.ID) CSLocalizedString {
	instance := getCSLocalizedStringClass().Alloc()
	rv := objc.Send[CSLocalizedString](instance.ID, objc.Sel("initWithLocalizedStrings:"), localizedStrings)
	rv.Autorelease()
	return rv
}


// Returns the localized string for the current language.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSLocalizedString/localizedString()
func (c_ CSLocalizedString) LocalizedString() string {
	rv := objc.Send[string](c_.ID, objc.Sel("localizedString"))
	return rv
}


