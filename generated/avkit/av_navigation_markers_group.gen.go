// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NavigationMarkersGroup] class.
var (
	NavigationMarkersGroupClass     _NavigationMarkersGroupClass
	NavigationMarkersGroupClassOnce sync.Once
)

func getNavigationMarkersGroupClass() _NavigationMarkersGroupClass {
	NavigationMarkersGroupClassOnce.Do(func() {
		NavigationMarkersGroupClass = _NavigationMarkersGroupClass{objc.GetClass("AVNavigationMarkersGroup")}
	})
	return NavigationMarkersGroupClass
}

type _NavigationMarkersGroupClass struct {
	class objc.Class
}

// An interface definition for the [NavigationMarkersGroup] class.
type INavigationMarkersGroup interface {
	objectivec.IObject
	// properties:
	DateRangeNavigationMarkers() []avfoundation.objc.IObject /* cross-framework: DateRangeMetadataGroup */
	TimedNavigationMarkers() []avfoundation.objc.IObject /* cross-framework: TimedMetadataGroup */
	Title() string /* primitive/slice/pointer. */
	// methods:
}

// A set of markers for navigating playback of an audiovisual presentation.
//
// The most common form of a navigation markers group is a chapter list; however, you can also provide other sets of markers to allow a user to jump to significant events in the presentation. For example, a “Goals Scored” markers group might summarize key moments in a recorded sporting event. When you associate navigation markers with an object you present with an , the user interface provides options for navigating each group.


// A set of markers for navigating playback of an audiovisual presentation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVNavigationMarkersGroup
type NavigationMarkersGroup struct {
	objectivec.Object
}

// NavigationMarkersGroupFrom constructs a [NavigationMarkersGroup] from an unsafe.Pointer.
//
// A set of markers for navigating playback of an audiovisual presentation.
func NavigationMarkersGroupFrom(ptr unsafe.Pointer) NavigationMarkersGroup {
	return NavigationMarkersGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NavigationMarkersGroupClass) Alloc() NavigationMarkersGroup {
	rv := objc.Send[NavigationMarkersGroup](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NavigationMarkersGroupClass) New() NavigationMarkersGroup {
	rv := objc.Send[NavigationMarkersGroup](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NavigationMarkersGroup) Init() NavigationMarkersGroup {
	rv := objc.Send[NavigationMarkersGroup](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NavigationMarkersGroup) Autorelease() NavigationMarkersGroup {
	rv := objc.Send[NavigationMarkersGroup](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNavigationMarkersGroup creates a new NavigationMarkersGroup instance.
func NewNavigationMarkersGroup() NavigationMarkersGroup {
	return getNavigationMarkersGroupClass().New()
}



// Initializes a navigation markers group with the specified title and array of date range navigation markers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVNavigationMarkersGroup/init(title:dateRangeNavigationMarkers:)
func NewNavigationMarkersGroupWithTitleDateRangeNavigationMarkers(title string /* primitive/slice/pointer. */, navigationMarkers []avfoundation.objc.IObject /* cross-framework DateRangeMetadataGroup */) NavigationMarkersGroup {
	instance := getNavigationMarkersGroupClass().Alloc()
	rv := objc.Send[NavigationMarkersGroup](instance.ID, objc.Sel("initWithTitle:dateRangeNavigationMarkers:"), objc.String(title), navigationMarkers)
	rv.Autorelease()
	return rv
}


// Initializes a navigation markers group with the specified title and array of timed navigation markers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVNavigationMarkersGroup/init(title:timedNavigationMarkers:)
func NewNavigationMarkersGroupWithTitleTimedNavigationMarkers(title string /* primitive/slice/pointer. */, navigationMarkers []avfoundation.objc.IObject /* cross-framework TimedMetadataGroup */) NavigationMarkersGroup {
	instance := getNavigationMarkersGroupClass().Alloc()
	rv := objc.Send[NavigationMarkersGroup](instance.ID, objc.Sel("initWithTitle:timedNavigationMarkers:"), objc.String(title), navigationMarkers)
	rv.Autorelease()
	return rv
}



// The array of date range navigation markers for which the group provides navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVNavigationMarkersGroup/dateRangeNavigationMarkers
func (n_ NavigationMarkersGroup) DateRangeNavigationMarkers() []avfoundation.objc.IObject /* cross-framework: DateRangeMetadataGroup */ {
	rv := objc.Send[[]avfoundation.DateRangeMetadataGroup](n_.ID, objc.Sel("dateRangeNavigationMarkers"))
	return rv
}


// The array of timed navigation markers for which the group provides navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVNavigationMarkersGroup/timedNavigationMarkers
func (n_ NavigationMarkersGroup) TimedNavigationMarkers() []avfoundation.objc.IObject /* cross-framework: TimedMetadataGroup */ {
	rv := objc.Send[[]avfoundation.TimedMetadataGroup](n_.ID, objc.Sel("timedNavigationMarkers"))
	return rv
}


// The title of the marker group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVNavigationMarkersGroup/title
func (n_ NavigationMarkersGroup) Title() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](n_.ID, objc.Sel("title"))
	return rv
}


