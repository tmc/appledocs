// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NavigationMarkersGroupClass) Alloc() NavigationMarkersGroup {
	rv := objc.Send[NavigationMarkersGroup](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Initializes a navigation markers group with the specified title and array of date range navigation markers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVNavigationMarkersGroup/init(title:dateRangeNavigationMarkers:)
func NewNavigationMarkersGroupWithTitleDateRangeNavigationMarkers(title objc.IObject /* cross-framework: NSString */, navigationMarkers []avfoundation.DateRangeMetadataGroup) NavigationMarkersGroup {
	instance := getNavigationMarkersGroupClass().Alloc()
	rv := objc.Send[NavigationMarkersGroup](instance.ID, objc.Sel("initWithTitle:dateRangeNavigationMarkers:"), title, navigationMarkers)
	rv.Autorelease()
	return rv
}


// Initializes a navigation markers group with the specified title and array of timed navigation markers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVNavigationMarkersGroup/init(title:timedNavigationMarkers:)
func NewNavigationMarkersGroupWithTitleTimedNavigationMarkers(title objc.IObject /* cross-framework: NSString */, navigationMarkers []avfoundation.TimedMetadataGroup) NavigationMarkersGroup {
	instance := getNavigationMarkersGroupClass().Alloc()
	rv := objc.Send[NavigationMarkersGroup](instance.ID, objc.Sel("initWithTitle:timedNavigationMarkers:"), title, navigationMarkers)
	rv.Autorelease()
	return rv
}



























