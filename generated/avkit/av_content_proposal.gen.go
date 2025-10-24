// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	DateOfAutomaticAcceptance() foundation.Date
	SetDateOfAutomaticAcceptance(value foundation.Date)
	PlayerLayoutGuide() appkit.LayoutGuide
	SetPlayerLayoutGuide(value appkit.LayoutGuide)
	PreferredPlayerViewFrame() corefoundation.CGRect
	SetPreferredPlayerViewFrame(value corefoundation.CGRect)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _ContentProposalClass) Alloc() ContentProposal {
	rv := objc.Send[ContentProposal](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates a new content proposal with the specified transition time, title, and preview image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposal/init(contentTimeForTransition:title:previewImage:)
func NewContentProposalWithContentTimeForTransitionTitlePreviewImage(contentTimeForTransition objc.IObject /* cross-framework: Time */, title objc.IObject /* cross-framework: NSString */, previewImage appkit.Image) ContentProposal {
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
func (c_ ContentProposal) DateOfAutomaticAcceptance() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("dateOfAutomaticAcceptance"))
	return rv
}


// The date that the system automatically accepts a proposal if the user doesn’t intervene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/dateofautomaticacceptance
func (c_ ContentProposal) SetDateOfAutomaticAcceptance(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDateOfAutomaticAcceptance:"), value)
}


// A layout guide that tracks the size and location of the player view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/playerlayoutguide
func (c_ ContentProposal) PlayerLayoutGuide() appkit.LayoutGuide {
	rv := objc.Send[appkit.LayoutGuide](c_.ID, objc.Sel("playerLayoutGuide"))
	return rv
}


// A layout guide that tracks the size and location of the player view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/playerlayoutguide
func (c_ ContentProposal) SetPlayerLayoutGuide(value appkit.LayoutGuide) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlayerLayoutGuide:"), value)
}


// The preferred presentation frame of the player view while the content proposal is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/preferredplayerviewframe
func (c_ ContentProposal) PreferredPlayerViewFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("preferredPlayerViewFrame"))
	return rv
}


// The preferred presentation frame of the player view while the content proposal is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/preferredplayerviewframe
func (c_ ContentProposal) SetPreferredPlayerViewFrame(value corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredPlayerViewFrame:"), value)
}







