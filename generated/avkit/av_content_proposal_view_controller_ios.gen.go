//go:build darwin && ios

// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for ContentProposalViewController


// Dismisses the current content proposal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/dismissContentProposal(for:animated:completion:)
func (c_ ContentProposalViewController) DismissContentProposalForActionAnimatedCompletion(action ContentProposalAction, animated bool, block unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("dismissContentProposalForAction:animated:completion:"), action, animated, block)
}

// iOS-only properties

// A prosal of content to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/contentProposal
func (c_ ContentProposalViewController) ContentProposal() IAVContentProposal {
	rv := objc.Send[ContentProposal](c_.ID, objc.Sel("contentProposal"))
	return rv
}

// The date that the system automatically accepts a proposal if the user doesn’t intervene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/dateOfAutomaticAcceptance
func (c_ ContentProposalViewController) DateOfAutomaticAcceptance() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("dateOfAutomaticAcceptance"))
	return rv
}
func (c_ ContentProposalViewController) SetDateOfAutomaticAcceptance(value objc.IObject /* cross-framework: NSDate */) {
	c_.ID.Send(objc.RegisterName("setDateOfAutomaticAcceptance:"), value)
}

// A layout guide that tracks the size and location of the player view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/playerLayoutGuide
func (c_ ContentProposalViewController) PlayerLayoutGuide() objc.IObject /* cross-framework: LayoutGuide */ {
	rv := objc.Send[appkit.LayoutGuide](c_.ID, objc.Sel("playerLayoutGuide"))
	return rv
}

// The player view controller that presents a content proposal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/playerViewController
func (c_ ContentProposalViewController) PlayerViewController() IAVPlayerViewController {
	rv := objc.Send[PlayerViewController](c_.ID, objc.Sel("playerViewController"))
	return rv
}

// The preferred presentation frame of the player view while the content proposal is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController/preferredPlayerViewFrame
func (c_ ContentProposalViewController) PreferredPlayerViewFrame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](c_.ID, objc.Sel("preferredPlayerViewFrame"))
	return rv
}





