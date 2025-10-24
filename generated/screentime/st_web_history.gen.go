// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

package screentime

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class STWebHistory */


/* debug [class_header]: Header for STWebHistory */
// The class instance for the [STWebHistory] class.
var (
	STWebHistoryClass     _STWebHistoryClass
	STWebHistoryClassOnce sync.Once
)

func getSTWebHistoryClass() _STWebHistoryClass {
	STWebHistoryClassOnce.Do(func() {
		STWebHistoryClass = _STWebHistoryClass{objc.GetClass("STWebHistory")}
	})
	return STWebHistoryClass
}

type _STWebHistoryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for STWebHistory */
// An interface definition for the [STWebHistory] class.
type ISTWebHistory interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for STWebHistory */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for STWebHistory */
	// methods:
	DeleteAllHistory()
	DeleteHistoryDuringInterval(interval foundation.DateInterval)
	DeleteHistoryForURL(url objc.IObject /* cross-framework: NSURL */)
	FetchAllHistoryWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchHistoryDuringIntervalCompletionHandler(interval foundation.DateInterval, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for STWebHistory */
// Alloc allocates a new instance without initialization.
func (sc _STWebHistoryClass) Alloc() STWebHistory {
	rv := objc.Send[STWebHistory](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _STWebHistoryClass) New() STWebHistory {
	rv := objc.Send[STWebHistory](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ STWebHistory) Init() STWebHistory {
	rv := objc.Send[STWebHistory](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ STWebHistory) Autorelease() STWebHistory {
	rv := objc.Send[STWebHistory](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSTWebHistory creates a new STWebHistory instance.
func NewSTWebHistory() STWebHistory {
	return getSTWebHistoryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for STWebHistory */
// The object you use to delete web-usage data.
//
// This class provides an easy way for you to delete web history, including: All history History associated to a specific URL History during a specific time interval


// The object you use to delete web-usage data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory
type STWebHistory struct {
	objectivec.Object
}

// STWebHistoryFrom constructs a [STWebHistory] from an unsafe.Pointer.
//
// The object you use to delete web-usage data.
func STWebHistoryFrom(ptr unsafe.Pointer) STWebHistory {
	return STWebHistory{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for STWebHistory */

// Creates a web history instance to delete web-usage data associated to the bundle identifier you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/init(bundleIdentifier:)
func NewSTWebHistoryWithBundleIdentifierError(bundleIdentifier objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) STWebHistory {
	instance := getSTWebHistoryClass().Alloc()
	rv := objc.Send[STWebHistory](instance.ID, objc.Sel("initWithBundleIdentifier:error:"), bundleIdentifier, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSTWebHistoryWithBundleIdentifierError */


// Creates a web history instance to delete web-usage data associated to the bundle identifier and profile identifier you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/init(bundleIdentifier:profileIdentifier:)
func NewSTWebHistoryWithBundleIdentifierProfileIdentifierError(bundleIdentifier objc.IObject /* cross-framework: NSString */, profileIdentifier STWebHistoryProfileIdentifier /* typedef */, error_ unsafe.Pointer) STWebHistory {
	instance := getSTWebHistoryClass().Alloc()
	rv := objc.Send[STWebHistory](instance.ID, objc.Sel("initWithBundleIdentifier:profileIdentifier:error:"), bundleIdentifier, profileIdentifier, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSTWebHistoryWithBundleIdentifierProfileIdentifierError */


// Creates a web history instance to delete web-usage data associated to the profile identifier you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/init(profileIdentifier:)
func NewSTWebHistoryWithProfileIdentifier(profileIdentifier STWebHistoryProfileIdentifier /* typedef */) STWebHistory {
	instance := getSTWebHistoryClass().Alloc()
	rv := objc.Send[STWebHistory](instance.ID, objc.Sel("initWithProfileIdentifier:"), profileIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSTWebHistoryWithProfileIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for STWebHistory */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for STWebHistory */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for STWebHistory */

// Deletes all web history associated with the bundle identifier you specified during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/deleteAllHistory()
func (s_ STWebHistory) DeleteAllHistory() {
	objc.Send[objc.ID](s_.ID, objc.Sel("deleteAllHistory"))
}/* debug [instance_methods/method]: DeleteAllHistory */


// Deletes web history that occurred during the date interval you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/deleteHistory(during:)
func (s_ STWebHistory) DeleteHistoryDuringInterval(interval foundation.DateInterval) {
	objc.Send[objc.ID](s_.ID, objc.Sel("deleteHistoryDuringInterval:"), interval)
}/* debug [instance_methods/method]: DeleteHistoryDuringInterval */


// Deletes all the web history for the URL you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/deleteHistory(for:)
func (s_ STWebHistory) DeleteHistoryForURL(url objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("deleteHistoryForURL:"), url)
}/* debug [instance_methods/method]: DeleteHistoryForURL */


// Fetches all web history associated with the bundle identifier and profile identifier you specified during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/fetchAllHistory(completionHandler:)
func (s_ STWebHistory) FetchAllHistoryWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("fetchAllHistoryWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: FetchAllHistoryWithCompletionHandler */


// Fetches web history that occurred during the date interval you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/fetchHistory(during:completionHandler:)
func (s_ STWebHistory) FetchHistoryDuringIntervalCompletionHandler(interval foundation.DateInterval, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("fetchHistoryDuringInterval:completionHandler:"), interval, completionHandler)
}/* debug [instance_methods/method]: FetchHistoryDuringIntervalCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for STWebHistory */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class STWebHistory */


