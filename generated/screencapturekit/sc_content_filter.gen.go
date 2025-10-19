// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCContentFilter] class.
var (
	sCContentFilterClass     _SCContentFilterClass
	sCContentFilterClassOnce sync.Once
)

func getSCContentFilterClass() _SCContentFilterClass {
	sCContentFilterClassOnce.Do(func() {
		sCContentFilterClass = _SCContentFilterClass{objc.GetClass("SCContentFilter")}
	})
	return sCContentFilterClass
}

type _SCContentFilterClass struct {
	class objc.Class
}

// An interface definition for the [SCContentFilter] class.
type ISCContentFilter interface {
	objectivec.IObject
}

// An instance that filters the content a stream captures.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter
type SCContentFilter struct {
	objectivec.Object
}

// SCContentFilterFrom constructs a [SCContentFilter] from an unsafe.Pointer.
//
// An instance that filters the content a stream captures.
func SCContentFilterFrom(ptr unsafe.Pointer) SCContentFilter {
	return SCContentFilter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCContentFilterClass) Alloc() SCContentFilter {
	rv := objc.Send[SCContentFilter](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCContentFilterClass) New() SCContentFilter {
	rv := objc.Send[SCContentFilter](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCContentFilter) Init() SCContentFilter {
	rv := objc.Send[SCContentFilter](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCContentFilter) Autorelease() SCContentFilter {
	rv := objc.Send[SCContentFilter](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCContentFilter creates a new SCContentFilter instance.
func NewSCContentFilter() SCContentFilter {
	return getSCContentFilterClass().New()
}


// Creates a filter that captures a display, excluding windows of the specified apps.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:excludingApplications:exceptingWindows:)
func NewSCContentFilterWithDisplayExcludingApplicationsExceptingWindows(display unsafe.Pointer, applications unsafe.Pointer, exceptingWindows unsafe.Pointer) SCContentFilter {
	instance := getSCContentFilterClass().Alloc()
	rv := objc.Send[SCContentFilter](instance.ID, objc.Sel("initWithDisplay:excludingApplications:exceptingWindows:"), display, applications, exceptingWindows)
	rv.Autorelease()
	return rv
}



