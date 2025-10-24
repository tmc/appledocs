// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSPopover */


/* debug [class_header]: Header for NSPopover */
// The class instance for the [Popover] class.
var (
	PopoverClass     _PopoverClass
	PopoverClassOnce sync.Once
)

func getPopoverClass() _PopoverClass {
	PopoverClassOnce.Do(func() {
		PopoverClass = _PopoverClass{objc.GetClass("NSPopover")}
	})
	return PopoverClass
}

type _PopoverClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Popover */
// An interface definition for the [Popover] class.
type IPopover interface {
	IResponder
	
/* debug [class_interface_properties]: Properties for Popover */
	// properties:
	Animates() bool
	SetAnimates(value bool)
	Appearance() IAppearance
	SetAppearance(value IAppearance)
	ContentSize() Size /* not a class type */
	SetContentSize(value Size /* not a class type */)
	ContentViewController() IViewController
	SetContentViewController(value IViewController)
	Delegate() objc.IObject /* cross-framework: PopoverDelegate */
	SetDelegate(value objc.IObject /* cross-framework: PopoverDelegate */)
	EffectiveAppearance() IAppearance
	SetEffectiveAppearance(value IAppearance)
	HasFullSizeContent() bool
	SetHasFullSizeContent(value bool)
	IsDetached() bool
	SetIsDetached(value bool)
	IsShown() bool
	SetIsShown(value bool)
	PositioningRect() Rect /* not a class type */
	SetPositioningRect(value Rect /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Popover */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Popover */
// Alloc allocates a new instance without initialization.
func (pc _PopoverClass) Alloc() Popover {
	rv := objc.Send[Popover](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PopoverClass) New() Popover {
	rv := objc.Send[Popover](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Popover) Init() Popover {
	rv := objc.Send[Popover](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Popover) Autorelease() Popover {
	rv := objc.Send[Popover](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPopover creates a new Popover instance.
func NewPopover() Popover {
	return getPopoverClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Popover */
// A means to display additional content related to existing content on the screen.
//
// The popover is positioned relative to the existing content and an anchor is used to express the relation between these two units of content. A popover has an appearance that specifies its visual characteristics, as well as a behavior that determines which user interactions will cause the popover to close. A transient popover is closed in response to most user interactions, whereas a semi-transient popover is closed when the user interacts with the window containing the popover’s positioning view. Popovers with application-defined behavior are not usually closed on the developer’s behalf. The system automatically positions each popover relative to its positioning view and moves the popover whenever its positioning view moves. A positioning rectangle within the positioning view can be specified for additional granularity. Popovers can be detached to become a separate window when they are dragged by implementing the appropriate delegate method.


// A means to display additional content related to existing content on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover
type Popover struct {
	Responder
}

// PopoverFrom constructs a [Popover] from an unsafe.Pointer.
//
// A means to display additional content related to existing content on the screen.
func PopoverFrom(ptr unsafe.Pointer) Popover {
	return Popover{
		Responder: ResponderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Popover *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Popover */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Popover */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Popover */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Popover */

// Specifies if the popover is to be animated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/animates
func (p_ Popover) Animates() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("animates"))
	return rv
}/* debug [instance_properties/getter]: animates */


// Specifies if the popover is to be animated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/animates
func (p_ Popover) SetAnimates(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAnimates:"), value)
}/* debug [instance_properties/setter]: animates */


// The appearance of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/appearance-swift.property
func (p_ Popover) Appearance() IAppearance {
	rv := objc.Send[Appearance](p_.ID, objc.Sel("appearance"))
	return rv
}/* debug [instance_properties/getter]: appearance */


// The appearance of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/appearance-swift.property
func (p_ Popover) SetAppearance(value IAppearance) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAppearance:"), value)
}/* debug [instance_properties/setter]: appearance */


// The content size of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/contentsize
func (p_ Popover) ContentSize() Size /* not a class type */ {
	rv := objc.Send[Size](p_.ID, objc.Sel("contentSize"))
	return rv
}/* debug [instance_properties/getter]: contentSize */


// The content size of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/contentsize
func (p_ Popover) SetContentSize(value Size /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentSize:"), value)
}/* debug [instance_properties/setter]: contentSize */


// The view controller that manages the content of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/contentviewcontroller
func (p_ Popover) ContentViewController() IViewController {
	rv := objc.Send[ViewController](p_.ID, objc.Sel("contentViewController"))
	return rv
}/* debug [instance_properties/getter]: contentViewController */


// The view controller that manages the content of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/contentviewcontroller
func (p_ Popover) SetContentViewController(value IViewController) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentViewController:"), value)
}/* debug [instance_properties/setter]: contentViewController */


// The delegate of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/delegate
func (p_ Popover) Delegate() objc.IObject /* cross-framework: PopoverDelegate */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/delegate
func (p_ Popover) SetDelegate(value objc.IObject /* cross-framework: PopoverDelegate */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The appearance that will be used when the popover is displayed onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/effectiveappearance
func (p_ Popover) EffectiveAppearance() IAppearance {
	rv := objc.Send[Appearance](p_.ID, objc.Sel("effectiveAppearance"))
	return rv
}/* debug [instance_properties/getter]: effectiveAppearance */


// The appearance that will be used when the popover is displayed onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/effectiveappearance
func (p_ Popover) SetEffectiveAppearance(value IAppearance) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEffectiveAppearance:"), value)
}/* debug [instance_properties/setter]: effectiveAppearance */


// A Boolean value that indicates whether the content view of the popover extends into the arrow region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/hasfullsizecontent
func (p_ Popover) HasFullSizeContent() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasFullSizeContent"))
	return rv
}/* debug [instance_properties/getter]: hasFullSizeContent */


// A Boolean value that indicates whether the content view of the popover extends into the arrow region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/hasfullsizecontent
func (p_ Popover) SetHasFullSizeContent(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHasFullSizeContent:"), value)
}/* debug [instance_properties/setter]: hasFullSizeContent */


// A Boolean value that indicates whether the window created by a popover’s detachment is automatically created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/isdetached
func (p_ Popover) IsDetached() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isDetached"))
	return rv
}/* debug [instance_properties/getter]: isDetached */


// A Boolean value that indicates whether the window created by a popover’s detachment is automatically created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/isdetached
func (p_ Popover) SetIsDetached(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsDetached:"), value)
}/* debug [instance_properties/setter]: isDetached */


// The display state of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/isshown
func (p_ Popover) IsShown() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isShown"))
	return rv
}/* debug [instance_properties/getter]: isShown */


// The display state of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/isshown
func (p_ Popover) SetIsShown(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsShown:"), value)
}/* debug [instance_properties/setter]: isShown */


// The rectangle within the positioning view relative to which the popover should be positioned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/positioningrect
func (p_ Popover) PositioningRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](p_.ID, objc.Sel("positioningRect"))
	return rv
}/* debug [instance_properties/getter]: positioningRect */


// The rectangle within the positioning view relative to which the popover should be positioned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/positioningrect
func (p_ Popover) SetPositioningRect(value Rect /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPositioningRect:"), value)
}/* debug [instance_properties/setter]: positioningRect */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPopover */



