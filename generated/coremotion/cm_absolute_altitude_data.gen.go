// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AbsoluteAltitudeData] class.
var (
	AbsoluteAltitudeDataClass     _AbsoluteAltitudeDataClass
	AbsoluteAltitudeDataClassOnce sync.Once
)

func getAbsoluteAltitudeDataClass() _AbsoluteAltitudeDataClass {
	AbsoluteAltitudeDataClassOnce.Do(func() {
		AbsoluteAltitudeDataClass = _AbsoluteAltitudeDataClass{objc.GetClass("CMAbsoluteAltitudeData")}
	})
	return AbsoluteAltitudeDataClass
}

type _AbsoluteAltitudeDataClass struct {
	class objc.Class
}

// An interface definition for the [AbsoluteAltitudeData] class.
type IAbsoluteAltitudeData interface {
	ILogItem
	// properties:
	// methods:
}

// Data that records a change in absolute altitude.
//
// Absolute altitude is only available on iPhone 12 and later and Apple Watch 6 or SE and later.


// Data that records a change in absolute altitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAbsoluteAltitudeData
type AbsoluteAltitudeData struct {
	LogItem
}

// AbsoluteAltitudeDataFrom constructs a [AbsoluteAltitudeData] from an unsafe.Pointer.
//
// Data that records a change in absolute altitude.
func AbsoluteAltitudeDataFrom(ptr unsafe.Pointer) AbsoluteAltitudeData {
	return AbsoluteAltitudeData{
		LogItem: LogItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AbsoluteAltitudeDataClass) Alloc() AbsoluteAltitudeData {
	rv := objc.Send[AbsoluteAltitudeData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AbsoluteAltitudeDataClass) New() AbsoluteAltitudeData {
	rv := objc.Send[AbsoluteAltitudeData](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AbsoluteAltitudeData) Init() AbsoluteAltitudeData {
	rv := objc.Send[AbsoluteAltitudeData](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AbsoluteAltitudeData) Autorelease() AbsoluteAltitudeData {
	rv := objc.Send[AbsoluteAltitudeData](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAbsoluteAltitudeData creates a new AbsoluteAltitudeData instance.
func NewAbsoluteAltitudeData() AbsoluteAltitudeData {
	return getAbsoluteAltitudeDataClass().New()
}



