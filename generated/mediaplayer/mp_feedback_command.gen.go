// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MPFeedbackCommand */


/* debug [class_header]: Header for MPFeedbackCommand */
// The class instance for the [FeedbackCommand] class.
var (
	FeedbackCommandClass     _FeedbackCommandClass
	FeedbackCommandClassOnce sync.Once
)

func getFeedbackCommandClass() _FeedbackCommandClass {
	FeedbackCommandClassOnce.Do(func() {
		FeedbackCommandClass = _FeedbackCommandClass{objc.GetClass("MPFeedbackCommand")}
	})
	return FeedbackCommandClass
}

type _FeedbackCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FeedbackCommand */
// An interface definition for the [FeedbackCommand] class.
type IFeedbackCommand interface {
	IRemoteCommand
	
/* debug [class_interface_properties]: Properties for FeedbackCommand */
	// properties:
	Active() bool
	SetActive(value bool)
	LocalizedShortTitle() objc.IObject /* cross-framework: NSString */
	SetLocalizedShortTitle(value objc.IObject /* cross-framework: NSString */)
	LocalizedTitle() objc.IObject /* cross-framework: NSString */
	SetLocalizedTitle(value objc.IObject /* cross-framework: NSString */)
	IsActive() bool
	SetIsActive(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FeedbackCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FeedbackCommand */
// Alloc allocates a new instance without initialization.
func (fc _FeedbackCommandClass) Alloc() FeedbackCommand {
	rv := objc.Send[FeedbackCommand](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FeedbackCommandClass) New() FeedbackCommand {
	rv := objc.Send[FeedbackCommand](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FeedbackCommand) Init() FeedbackCommand {
	rv := objc.Send[FeedbackCommand](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FeedbackCommand) Autorelease() FeedbackCommand {
	rv := objc.Send[FeedbackCommand](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFeedbackCommand creates a new FeedbackCommand instance.
func NewFeedbackCommand() FeedbackCommand {
	return getFeedbackCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FeedbackCommand */
// An object that reflects the feedback state for the playing item.
//
// The shared object vends feedback objects for liking, disliking, and bookmarking media items. Use these objects to register handlers for the types of feedback your app supports and to perform the appropriate tasks when that feedback changes. When the currently playing item changes, you can also use this object to set the feedback state for the new item. When the state of a feedback item changes, the system delivers an appropriate event to registered handlers of this object. Your handler code must determine which media item receives the feedback and then apply the update the feedback state for that item. You might also perform other tasks related to receiving feedback. For example, if the user likes the currently playing song, you might update the appropriate UI in your app or use the information to recommend similar songs.


// An object that reflects the feedback state for the playing item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPFeedbackCommand
type FeedbackCommand struct {
	RemoteCommand
}

// FeedbackCommandFrom constructs a [FeedbackCommand] from an unsafe.Pointer.
//
// An object that reflects the feedback state for the playing item.
func FeedbackCommandFrom(ptr unsafe.Pointer) FeedbackCommand {
	return FeedbackCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FeedbackCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FeedbackCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FeedbackCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FeedbackCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FeedbackCommand */

// A Boolean value that indicates whether the feedback’s action is on or off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPFeedbackCommand/isActive
func (f_ FeedbackCommand) Active() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// A Boolean value that indicates whether the feedback’s action is on or off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPFeedbackCommand/isActive
func (f_ FeedbackCommand) SetActive(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setActive:"), value)
}/* debug [instance_properties/setter]: active */


// A shortened version of the string used to describe the context of a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPFeedbackCommand/localizedShortTitle
func (f_ FeedbackCommand) LocalizedShortTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("localizedShortTitle"))
	return rv
}/* debug [instance_properties/getter]: localizedShortTitle */


// A shortened version of the string used to describe the context of a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPFeedbackCommand/localizedShortTitle
func (f_ FeedbackCommand) SetLocalizedShortTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLocalizedShortTitle:"), value)
}/* debug [instance_properties/setter]: localizedShortTitle */


// A localized string used to describe the context of a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPFeedbackCommand/localizedTitle
func (f_ FeedbackCommand) LocalizedTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("localizedTitle"))
	return rv
}/* debug [instance_properties/getter]: localizedTitle */


// A localized string used to describe the context of a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPFeedbackCommand/localizedTitle
func (f_ FeedbackCommand) SetLocalizedTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLocalizedTitle:"), value)
}/* debug [instance_properties/setter]: localizedTitle */


// A Boolean value that indicates whether the feedback’s action is on or off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommand/isactive
func (f_ FeedbackCommand) IsActive() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// A Boolean value that indicates whether the feedback’s action is on or off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommand/isactive
func (f_ FeedbackCommand) SetIsActive(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPFeedbackCommand */



