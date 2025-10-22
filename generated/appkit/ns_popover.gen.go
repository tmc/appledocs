// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Popover] class.
type IPopover interface {
	IResponder
	PerformClose(sender objectivec.IObject)
	Behavior() PopoverBehavior
	SetBehavior(value PopoverBehavior)
	ContentViewController() NSViewController
	SetContentViewController(value IViewController)
	EffectiveAppearance() NSAppearance
	HasFullSizeContent() bool
	SetHasFullSizeContent(value bool)
	Shown() bool
	Animates() bool
	SetAnimates(value bool)
	Appearance() NSAppearance
	SetAppearance(value IAppearance)
	ContentSize() coregraphics.CGSize
	SetContentSize(value coregraphics.CGSize)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	IsDetached() bool
	SetIsDetached(value bool)
	IsShown() bool
	SetIsShown(value bool)
	PositioningRect() coregraphics.CGRect
	SetPositioningRect(value coregraphics.CGRect)
}

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

// Alloc allocates a new instance without initialization.
func (pc _PopoverClass) Alloc() Popover {
	rv := objc.Send[Popover](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Attempts to close the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/performClose(_:)

func (p_ Popover) PerformClose(sender objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performClose:"), sender)
}


// Specifies the behavior of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/behavior-swift.property

func (p_ Popover) Behavior() PopoverBehavior {
	rv := objc.Send[PopoverBehavior](p_.ID, objc.Sel("behavior"))
	return rv
}


// Specifies the behavior of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/behavior-swift.property

func (p_ Popover) SetBehavior(value PopoverBehavior) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBehavior:"), value)
}


// The view controller that manages the content of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/contentViewController

func (p_ Popover) ContentViewController() NSViewController {
	rv := objc.Send[NSViewController](p_.ID, objc.Sel("contentViewController"))
	return rv
}


// The view controller that manages the content of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/contentViewController

func (p_ Popover) SetContentViewController(value IViewController) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentViewController:"), value)
}


// The appearance that will be used when the popover is displayed onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/effectiveAppearance

func (p_ Popover) EffectiveAppearance() NSAppearance {
	rv := objc.Send[NSAppearance](p_.ID, objc.Sel("effectiveAppearance"))
	return rv
}


// A Boolean value that indicates whether the content view of the popover extends into the arrow region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/hasFullSizeContent

func (p_ Popover) HasFullSizeContent() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasFullSizeContent"))
	return rv
}


// A Boolean value that indicates whether the content view of the popover extends into the arrow region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/hasFullSizeContent

func (p_ Popover) SetHasFullSizeContent(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHasFullSizeContent:"), value)
}


// The display state of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/isShown

func (p_ Popover) Shown() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shown"))
	return rv
}


// Specifies if the popover is to be animated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/animates

func (p_ Popover) Animates() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("animates"))
	return rv
}


// Specifies if the popover is to be animated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/animates

func (p_ Popover) SetAnimates(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAnimates:"), value)
}


// The appearance of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/appearance-swift.property

func (p_ Popover) Appearance() NSAppearance {
	rv := objc.Send[NSAppearance](p_.ID, objc.Sel("appearance"))
	return rv
}


// The appearance of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/appearance-swift.property

func (p_ Popover) SetAppearance(value IAppearance) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAppearance:"), value)
}


// The content size of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/contentsize

func (p_ Popover) ContentSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](p_.ID, objc.Sel("contentSize"))
	return rv
}


// The content size of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/contentsize

func (p_ Popover) SetContentSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentSize:"), value)
}


// The delegate of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/delegate

func (p_ Popover) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/delegate

func (p_ Popover) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether the window created by a popover’s detachment is automatically created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/isdetached

func (p_ Popover) IsDetached() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isDetached"))
	return rv
}


// A Boolean value that indicates whether the window created by a popover’s detachment is automatically created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/isdetached

func (p_ Popover) SetIsDetached(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsDetached:"), value)
}


// The display state of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/isshown

func (p_ Popover) IsShown() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isShown"))
	return rv
}


// The display state of the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/isshown

func (p_ Popover) SetIsShown(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsShown:"), value)
}


// The rectangle within the positioning view relative to which the popover should be positioned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/positioningrect

func (p_ Popover) PositioningRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("positioningRect"))
	return rv
}


// The rectangle within the positioning view relative to which the popover should be positioned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/positioningrect

func (p_ Popover) SetPositioningRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPositioningRect:"), value)
}



