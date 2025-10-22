// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	ContentTimeForTransition() unsafe.Pointer
	AutomaticAcceptanceInterval() unsafe.Pointer
	SetAutomaticAcceptanceInterval(value unsafe.Pointer)
	Metadata() avfoundation.MetadataItem
	SetMetadata(value avfoundation.IMetadataItem)
	PreviewImage() appkit.Image
	SetPreviewImage(value appkit.IImage)
	Title() string
	SetTitle(value string)
	Url() foundation.URL
	SetUrl(value foundation.IURL)
	ContentProposal() AVContentProposal
	SetContentProposal(value IAVContentProposal)
	DateOfAutomaticAcceptance() foundation.Date
	SetDateOfAutomaticAcceptance(value foundation.IDate)
	PlayerLayoutGuide() appkit.LayoutGuide
	SetPlayerLayoutGuide(value appkit.ILayoutGuide)
	PreferredPlayerViewFrame() coregraphics.CGRect
	SetPreferredPlayerViewFrame(value coregraphics.CGRect)
}

// An object that describes the content to propose playing after the current item finishes.
//
// A content proposal object models the data about the proposed content such as its title, preview image, presentation time, and content URL. You make a content proposal eligible for presentation by setting it as the of the current .
//
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


// The time within the timeline of the current player item when the content proposal presentation should begin.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposal/contentTimeForTransition
func (c_ ContentProposal) ContentTimeForTransition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("contentTimeForTransition"))
	return rv
}

// The interval between the time playback ends and automatic acceptance of this content proposal.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposal/automaticacceptanceinterval
func (c_ ContentProposal) AutomaticAcceptanceInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("automaticAcceptanceInterval"))
	return rv
}


// SetAutomaticAcceptanceInterval sets the value of the automaticAcceptanceInterval property.
// The interval between the time playback ends and automatic acceptance of this content proposal.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposal/automaticacceptanceinterval
func (c_ ContentProposal) SetAutomaticAcceptanceInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticAcceptanceInterval:"), value)
}

// Optional custom metadata associated with the proposed item.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposal/metadata
func (c_ ContentProposal) Metadata() avfoundation.MetadataItem {
	rv := objc.Send[avfoundation.MetadataItem](c_.ID, objc.Sel("metadata"))
	return rv
}


// SetMetadata sets the value of the metadata property.
// Optional custom metadata associated with the proposed item.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposal/metadata
func (c_ ContentProposal) SetMetadata(value avfoundation.IMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), value)
}

// The preview image of the proposed item.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposal/previewimage
func (c_ ContentProposal) PreviewImage() appkit.Image {
	rv := objc.Send[appkit.Image](c_.ID, objc.Sel("previewImage"))
	return rv
}


// SetPreviewImage sets the value of the previewImage property.
// The preview image of the proposed item.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposal/previewimage
func (c_ ContentProposal) SetPreviewImage(value appkit.IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviewImage:"), value)
}

// The title of the proposed content.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposal/title
func (c_ ContentProposal) Title() string {
	rv := objc.Send[string](c_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title of the proposed content.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposal/title
func (c_ ContentProposal) SetTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// The URL of the proposed content.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposal/url
func (c_ ContentProposal) Url() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL of the proposed content.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposal/url
func (c_ ContentProposal) SetUrl(value foundation.IURL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrl:"), value)
}

// A prosal of content to play.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/contentproposal
func (c_ ContentProposal) ContentProposal() AVContentProposal {
	rv := objc.Send[AVContentProposal](c_.ID, objc.Sel("contentProposal"))
	return rv
}


// SetContentProposal sets the value of the contentProposal property.
// A prosal of content to play.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/contentproposal
func (c_ ContentProposal) SetContentProposal(value IAVContentProposal) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentProposal:"), value)
}

// The date that the system automatically accepts a proposal if the user doesn’t intervene.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/dateofautomaticacceptance
func (c_ ContentProposal) DateOfAutomaticAcceptance() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("dateOfAutomaticAcceptance"))
	return rv
}


// SetDateOfAutomaticAcceptance sets the value of the dateOfAutomaticAcceptance property.
// The date that the system automatically accepts a proposal if the user doesn’t intervene.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/dateofautomaticacceptance
func (c_ ContentProposal) SetDateOfAutomaticAcceptance(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDateOfAutomaticAcceptance:"), value)
}

// A layout guide that tracks the size and location of the player view.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/playerlayoutguide
func (c_ ContentProposal) PlayerLayoutGuide() appkit.LayoutGuide {
	rv := objc.Send[appkit.LayoutGuide](c_.ID, objc.Sel("playerLayoutGuide"))
	return rv
}


// SetPlayerLayoutGuide sets the value of the playerLayoutGuide property.
// A layout guide that tracks the size and location of the player view.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/playerlayoutguide
func (c_ ContentProposal) SetPlayerLayoutGuide(value appkit.ILayoutGuide) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlayerLayoutGuide:"), value)
}

// The preferred presentation frame of the player view while the content proposal is active.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/preferredplayerviewframe
func (c_ ContentProposal) PreferredPlayerViewFrame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("preferredPlayerViewFrame"))
	return rv
}


// SetPreferredPlayerViewFrame sets the value of the preferredPlayerViewFrame property.
// The preferred presentation frame of the player view while the content proposal is active.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcontentproposalviewcontroller/preferredplayerviewframe
func (c_ ContentProposal) SetPreferredPlayerViewFrame(value coregraphics.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredPlayerViewFrame:"), value)
}



