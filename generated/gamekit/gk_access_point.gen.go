// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccessPoint] class.
var (
	AccessPointClass     _AccessPointClass
	AccessPointClassOnce sync.Once
)

func getAccessPointClass() _AccessPointClass {
	AccessPointClassOnce.Do(func() {
		AccessPointClass = _AccessPointClass{objc.GetClass("GKAccessPoint")}
	})
	return AccessPointClass
}

type _AccessPointClass struct {
	class objc.Class
}

// An interface definition for the [AccessPoint] class.
type IAccessPoint interface {
	objectivec.IObject
	// properties:
	Visible() bool
	ShowHighlights() bool
	SetShowHighlights(value bool)
	FrameInScreenCoordinates() objc.IObject /* cross-framework: Rect */
	SetFrameInScreenCoordinates(value objc.IObject /* cross-framework: Rect */)
	IsActive() bool
	SetIsActive(value bool)
	IsFocused() bool
	SetIsFocused(value bool)
	IsPresentingGameCenter() bool
	SetIsPresentingGameCenter(value bool)
	IsVisible() bool
	SetIsVisible(value bool)
	Location() objc.IObject /* cross-framework: Location */
	SetLocation(value objc.IObject /* cross-framework: Location */)
	ParentWindow() objc.IObject /* cross-framework: Window */
	SetParentWindow(value objc.IObject /* cross-framework: Window */)
	// methods:
	TriggerAccessPointForChallengesWithHandler(handler unsafe.Pointer)
	TriggerAccessPointForPlayTogetherWithHandler(handler unsafe.Pointer)
}

// An object that allows players to view and manage their Game Center information from within your game.
//
// The access point displays a control in a corner of your game that opens a Game Center dashboard when the player taps or clicks it. Use the property to get the shared access point object. GameKit attaches the access point to the window you specify in the property, in the corner you specify using the property. If you don’t specify a parent window, GameKit infers an appropriate location. For the location of the access point on visionOS, see . To display highlights, set the property to . Then set to to display the access point control.


// An object that allows players to view and manage their Game Center information from within your game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint
type AccessPoint struct {
	objectivec.Object
}

// AccessPointFrom constructs a [AccessPoint] from an unsafe.Pointer.
//
// An object that allows players to view and manage their Game Center information from within your game.
func AccessPointFrom(ptr unsafe.Pointer) AccessPoint {
	return AccessPoint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AccessPointClass) Alloc() AccessPoint {
	rv := objc.Send[AccessPoint](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccessPointClass) New() AccessPoint {
	rv := objc.Send[AccessPoint](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccessPoint) Init() AccessPoint {
	rv := objc.Send[AccessPoint](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccessPoint) Autorelease() AccessPoint {
	rv := objc.Send[AccessPoint](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccessPoint creates a new AccessPoint instance.
func NewAccessPoint() AccessPoint {
	return getAccessPointClass().New()
}



// Displays the view that allows players to engage each other with challenges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/triggerForChallenges(handler:)
func (a_ AccessPoint) TriggerAccessPointForChallengesWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointForChallengesWithHandler:"), handler)
}


// Displays the view that allows players to engage each other with activities and challenges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/triggerForPlayTogether(handler:)
func (a_ AccessPoint) TriggerAccessPointForPlayTogetherWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("triggerAccessPointForPlayTogetherWithHandler:"), handler)
}


// A Boolean value that indicates whether the access point is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/isVisible
func (a_ AccessPoint) Visible() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("visible"))
	return rv
}


// A Boolean value that indicates whether to display highlights for achievements and current ranks for leaderboards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/showHighlights
func (a_ AccessPoint) ShowHighlights() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("showHighlights"))
	return rv
}


// A Boolean value that indicates whether to display highlights for achievements and current ranks for leaderboards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/showHighlights
func (a_ AccessPoint) SetShowHighlights(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShowHighlights:"), value)
}


// The frame of the access point in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/frameinscreencoordinates
func (a_ AccessPoint) FrameInScreenCoordinates() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](a_.ID, objc.Sel("frameInScreenCoordinates"))
	return rv
}


// The frame of the access point in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/frameinscreencoordinates
func (a_ AccessPoint) SetFrameInScreenCoordinates(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFrameInScreenCoordinates:"), value)
}


// A Boolean value that determines whether to display the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/isactive
func (a_ AccessPoint) IsActive() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isActive"))
	return rv
}


// A Boolean value that determines whether to display the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/isactive
func (a_ AccessPoint) SetIsActive(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsActive:"), value)
}


// A Boolean value that indicates whether the access point is in focus on tvOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/isfocused
func (a_ AccessPoint) IsFocused() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isFocused"))
	return rv
}


// A Boolean value that indicates whether the access point is in focus on tvOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/isfocused
func (a_ AccessPoint) SetIsFocused(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsFocused:"), value)
}


// A Boolean value that indicates whether the game is presenting the Game Center dashboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/ispresentinggamecenter
func (a_ AccessPoint) IsPresentingGameCenter() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPresentingGameCenter"))
	return rv
}


// A Boolean value that indicates whether the game is presenting the Game Center dashboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/ispresentinggamecenter
func (a_ AccessPoint) SetIsPresentingGameCenter(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPresentingGameCenter:"), value)
}


// A Boolean value that indicates whether the access point is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/isvisible
func (a_ AccessPoint) IsVisible() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isVisible"))
	return rv
}


// A Boolean value that indicates whether the access point is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/isvisible
func (a_ AccessPoint) SetIsVisible(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsVisible:"), value)
}


// The corner of the screen to display the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/location-swift.property
func (a_ AccessPoint) Location() objc.IObject /* cross-framework: Location */ {
	rv := objc.Send[corelocation.Location](a_.ID, objc.Sel("location"))
	return rv
}


// The corner of the screen to display the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/location-swift.property
func (a_ AccessPoint) SetLocation(value objc.IObject /* cross-framework: Location */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLocation:"), value)
}


// The window that contains the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/parentwindow
func (a_ AccessPoint) ParentWindow() objc.IObject /* cross-framework: Window */ {
	rv := objc.Send[appkit.Window](a_.ID, objc.Sel("parentWindow"))
	return rv
}


// The window that contains the access point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkaccesspoint/parentwindow
func (a_ AccessPoint) SetParentWindow(value objc.IObject /* cross-framework: Window */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParentWindow:"), value)
}



