// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

package screentime

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [STWebHistory] class.
type ISTWebHistory interface {
	objectivec.IObject
	DeleteAllHistory()
	DeleteHistoryDuringInterval(interval unsafe.Pointer)
	FetchAllHistoryWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchHistoryDuringIntervalCompletionHandler(interval unsafe.Pointer, completionHandler unsafe.Pointer)
}

// The object you use to delete web-usage data.
//
// This class provides an easy way for you to delete web history, including: All history History associated to a specific URL History during a specific time interval
//
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

// Alloc allocates a new instance without initialization.
func (sc _STWebHistoryClass) Alloc() STWebHistory {
	rv := objc.Send[STWebHistory](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a web history instance to delete web-usage data associated to the profile identifier you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/init(profileIdentifier:)
func NewSTWebHistoryWithProfileIdentifier(profileIdentifier unsafe.Pointer) STWebHistory {
	instance := getSTWebHistoryClass().Alloc()
	rv := objc.Send[STWebHistory](instance.ID, objc.Sel("initWithProfileIdentifier:"), profileIdentifier)
	rv.Autorelease()
	return rv
}

// Creates a web history instance to delete web-usage data associated to the bundle identifier you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/init(bundleIdentifier:)
func NewSTWebHistoryWithBundleIdentifierError(bundleIdentifier string, error unsafe.Pointer) STWebHistory {
	instance := getSTWebHistoryClass().Alloc()
	rv := objc.Send[STWebHistory](instance.ID, objc.Sel("initWithBundleIdentifier:error:"), objc.String(bundleIdentifier), error)
	rv.Autorelease()
	return rv
}

// Creates a web history instance to delete web-usage data associated to the bundle identifier and profile identifier you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/init(bundleIdentifier:profileIdentifier:)
func NewSTWebHistoryWithBundleIdentifierProfileIdentifierError(bundleIdentifier string, profileIdentifier unsafe.Pointer, error unsafe.Pointer) STWebHistory {
	instance := getSTWebHistoryClass().Alloc()
	rv := objc.Send[STWebHistory](instance.ID, objc.Sel("initWithBundleIdentifier:profileIdentifier:error:"), objc.String(bundleIdentifier), profileIdentifier, error)
	rv.Autorelease()
	return rv
}


// Deletes all web history associated with the bundle identifier you specified during initialization.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/deleteAllHistory()
func (s_ STWebHistory) DeleteAllHistory() {
	objc.Send[objc.ID](s_.ID, objc.Sel("deleteAllHistory"))
}

// Deletes web history that occurred during the date interval you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/deleteHistory(during:)
func (s_ STWebHistory) DeleteHistoryDuringInterval(interval unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("deleteHistoryDuringInterval:"), interval)
}

// Fetches all web history associated with the bundle identifier and profile identifier you specified during initialization.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/fetchAllHistory(completionHandler:)
func (s_ STWebHistory) FetchAllHistoryWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("fetchAllHistoryWithCompletionHandler:"), completionHandler)
}

// Fetches web history that occurred during the date interval you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebHistory/fetchHistory(during:completionHandler:)
func (s_ STWebHistory) FetchHistoryDuringIntervalCompletionHandler(interval unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("fetchHistoryDuringInterval:completionHandler:"), interval, completionHandler)
}


