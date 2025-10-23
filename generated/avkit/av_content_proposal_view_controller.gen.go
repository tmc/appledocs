// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ContentProposalViewController] class.
var (
	ContentProposalViewControllerClass     _ContentProposalViewControllerClass
	ContentProposalViewControllerClassOnce sync.Once
)

func getContentProposalViewControllerClass() _ContentProposalViewControllerClass {
	ContentProposalViewControllerClassOnce.Do(func() {
		ContentProposalViewControllerClass = _ContentProposalViewControllerClass{objc.GetClass("AVContentProposalViewController")}
	})
	return ContentProposalViewControllerClass
}

type _ContentProposalViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [ContentProposalViewController] class.
type IContentProposalViewController interface {
	appkit.IViewController
	DismissContentProposalForActionAnimatedCompletion(action IContentProposalAction, animated bool, block unsafe.Pointer)
	ContentProposal() AVContentProposal
	DateOfAutomaticAcceptance() foundation.NSDate
	SetDateOfAutomaticAcceptance(value foundation.IDate)
	PlayerLayoutGuide() appkit.LayoutGuide
	PlayerViewController() AVPlayerViewController
	PreferredPlayerViewFrame() coregraphics.CGRect
}

// A view controller that proposes content to watch next.
//
// Subclass this class to define the user interface for your content proposal.


// A view controller that proposes content to watch next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController
type ContentProposalViewController struct {
	appkit.ViewController
}

// ContentProposalViewControllerFrom constructs a [ContentProposalViewController] from an unsafe.Pointer.
//
// A view controller that proposes content to watch next.
func ContentProposalViewControllerFrom(ptr unsafe.Pointer) ContentProposalViewController {
	return ContentProposalViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ContentProposalViewControllerClass) Alloc() ContentProposalViewController {
	rv := objc.Send[ContentProposalViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ContentProposalViewControllerClass) New() ContentProposalViewController {
	rv := objc.Send[ContentProposalViewController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentProposalViewController) Init() ContentProposalViewController {
	rv := objc.Send[ContentProposalViewController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentProposalViewController) Autorelease() ContentProposalViewController {
	rv := objc.Send[ContentProposalViewController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentProposalViewController creates a new ContentProposalViewController instance.
func NewContentProposalViewController() ContentProposalViewController {
	return getContentProposalViewControllerClass().New()
}



// Dismisses the current content proposal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/dismissContentProposal(for:animated:completion:)
func (c_ ContentProposalViewController) DismissContentProposalForActionAnimatedCompletion(action IContentProposalAction, animated bool, block unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("dismissContentProposalForAction:animated:completion:"), action, animated, block)
}


// A prosal of content to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/contentProposal
func (c_ ContentProposalViewController) ContentProposal() AVContentProposal {
	rv := objc.Send[AVContentProposal](c_.ID, objc.Sel("contentProposal"))
	return rv
}


// The date that the system automatically accepts a proposal if the user doesn’t intervene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/dateOfAutomaticAcceptance
func (c_ ContentProposalViewController) DateOfAutomaticAcceptance() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("dateOfAutomaticAcceptance"))
	return rv
}


// The date that the system automatically accepts a proposal if the user doesn’t intervene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/dateOfAutomaticAcceptance
func (c_ ContentProposalViewController) SetDateOfAutomaticAcceptance(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDateOfAutomaticAcceptance:"), value)
}


// A layout guide that tracks the size and location of the player view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/playerLayoutGuide
func (c_ ContentProposalViewController) PlayerLayoutGuide() appkit.LayoutGuide {
	rv := objc.Send[appkit.LayoutGuide](c_.ID, objc.Sel("playerLayoutGuide"))
	return rv
}


// The player view controller that presents a content proposal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/playerViewController
func (c_ ContentProposalViewController) PlayerViewController() AVPlayerViewController {
	rv := objc.Send[AVPlayerViewController](c_.ID, objc.Sel("playerViewController"))
	return rv
}


// The preferred presentation frame of the player view while the content proposal is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/preferredPlayerViewFrame
func (c_ ContentProposalViewController) PreferredPlayerViewFrame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("preferredPlayerViewFrame"))
	return rv
}



