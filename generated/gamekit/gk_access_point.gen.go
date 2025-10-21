// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An object that allows players to view and manage their Game Center information from within your game.
//
// The access point displays a control in a corner of your game that opens a Game Center dashboard when the player taps or clicks it. Use the property to get the shared access point object. GameKit attaches the access point to the window you specify in the property, in the corner you specify using the property. If you don’t specify a parent window, GameKit infers an appropriate location. For the location of the access point on visionOS, see . To display highlights, set the property to . Then set to to display the access point control.
//
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


// A Boolean value that indicates whether the access point is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/isVisible
func (a_ AccessPoint) Visible() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("visible"))
	return rv
}

// A Boolean value that indicates whether to display highlights for achievements and current ranks for leaderboards.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/showHighlights
func (a_ AccessPoint) ShowHighlights() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("showHighlights"))
	return rv
}


// SetShowHighlights sets the value of the showHighlights property.
// A Boolean value that indicates whether to display highlights for achievements and current ranks for leaderboards.

//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/showHighlights
func (a_ AccessPoint) SetShowHighlights(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShowHighlights:"), value)
}


