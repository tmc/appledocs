// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ContentProposal] class.
var (
	ContentProposalClass     _ContentProposalClass
	ContentProposalClassOnce sync.Once
)

func getContentProposalClass() _ContentProposalClass {
	ContentProposalClassOnce.Do(func() {
		ContentProposalClass = _ContentProposalClass{objc.GetClass("AVContentProposal")}
	})
	return ContentProposalClass
}

type _ContentProposalClass struct {
	class objc.Class
}

// An interface definition for the [ContentProposal] class.
type IContentProposal interface {
	objectivec.IObject
	// properties:
	ContentProposal() IAVContentProposal
	SetContentProposal(value IAVContentProposal)
	DateOfAutomaticAcceptance() objc.IObject /* cross-framework: Date */
	SetDateOfAutomaticAcceptance(value objc.IObject /* cross-framework: Date */)
	PlayerLayoutGuide() objc.IObject /* cross-framework: LayoutGuide */
	SetPlayerLayoutGuide(value objc.IObject /* cross-framework: LayoutGuide */)
	PreferredPlayerViewFrame() objc.IObject /* cross-framework: Rect */
	SetPreferredPlayerViewFrame(value objc.IObject /* cross-framework: Rect */)
	// methods:
}

// An object that describes the content to propose playing after the current item finishes.
//
// A content proposal object models the data about the proposed content such as its title, preview image, presentation time, and content URL. You make a content proposal eligible for presentation by setting it as the of the current .


// An object that describes the content to propose playing after the current item finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposal
type ContentProposal struct {
	objectivec.Object
}

// ContentProposalFrom constructs a [ContentProposal] from an unsafe.Pointer.
//
// An object that describes the content to propose playing after the current item finishes.
func ContentProposalFrom(ptr unsafe.Pointer) ContentProposal {
	return ContentProposal{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ContentProposalClass) Alloc() ContentProposal {
	rv := objc.Send[ContentProposal](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ContentProposalClass) New() ContentProposal {
	rv := objc.Send[ContentProposal](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentProposal) Init() ContentProposal {
	rv := objc.Send[ContentProposal](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentProposal) Autorelease() ContentProposal {
	rv := objc.Send[ContentProposal](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentProposal creates a new ContentProposal instance.
func NewContentProposal() ContentProposal {
	return getContentProposalClass().New()
}



// Creates a new content proposal with the specified transition time, title, and preview image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposal/init(contentTimeForTransition:title:previewImage:)
func NewContentProposalWithContentTimeForTransitionTitlePreviewImage(contentTimeForTransition Time /* not a class type */, title objc.IObject /* cross-framework: NSString */, previewImage objc.IObject /* cross-framework: Image */) ContentProposal {
	instance := getContentProposalClass().Alloc()
	rv := objc.Send[ContentProposal](instance.ID, objc.Sel("initWithContentTimeForTransition:title:previewImage:"), contentTimeForTransition, title, previewImage)
	rv.Autorelease()
	return rv
}



// A prosal of content to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/contentproposal
func (c_ ContentProposal) ContentProposal() IAVContentProposal {
	rv := objc.Send[ContentProposal](c_.ID, objc.Sel("contentProposal"))
	return rv
}


// A prosal of content to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/contentproposal
func (c_ ContentProposal) SetContentProposal(value IAVContentProposal) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentProposal:"), value)
}


// The date that the system automatically accepts a proposal if the user doesn’t intervene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/dateofautomaticacceptance
func (c_ ContentProposal) DateOfAutomaticAcceptance() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("dateOfAutomaticAcceptance"))
	return rv
}


// The date that the system automatically accepts a proposal if the user doesn’t intervene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/dateofautomaticacceptance
func (c_ ContentProposal) SetDateOfAutomaticAcceptance(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDateOfAutomaticAcceptance:"), value)
}


// A layout guide that tracks the size and location of the player view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/playerlayoutguide
func (c_ ContentProposal) PlayerLayoutGuide() objc.IObject /* cross-framework: LayoutGuide */ {
	rv := objc.Send[appkit.LayoutGuide](c_.ID, objc.Sel("playerLayoutGuide"))
	return rv
}


// A layout guide that tracks the size and location of the player view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/playerlayoutguide
func (c_ ContentProposal) SetPlayerLayoutGuide(value objc.IObject /* cross-framework: LayoutGuide */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlayerLayoutGuide:"), value)
}


// The preferred presentation frame of the player view while the content proposal is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/preferredplayerviewframe
func (c_ ContentProposal) PreferredPlayerViewFrame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](c_.ID, objc.Sel("preferredPlayerViewFrame"))
	return rv
}


// The preferred presentation frame of the player view while the content proposal is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/preferredplayerviewframe
func (c_ ContentProposal) SetPreferredPlayerViewFrame(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredPlayerViewFrame:"), value)
}


