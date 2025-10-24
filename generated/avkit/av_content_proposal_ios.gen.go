//go:build darwin && ios

// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ContentProposal


// iOS-only properties

// The interval between the time playback ends and automatic acceptance of this content proposal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposal/automaticAcceptanceInterval
func (c_ ContentProposal) AutomaticAcceptanceInterval() float64 {
	rv := objc.Send[TimeInterval](c_.ID, objc.Sel("automaticAcceptanceInterval"))
	return rv
}
func (c_ ContentProposal) SetAutomaticAcceptanceInterval(value float64) {
	c_.ID.Send(objc.RegisterName("setAutomaticAcceptanceInterval:"), value)
}

// The time within the timeline of the current player item when the content proposal presentation should begin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposal/contentTimeForTransition
func (c_ ContentProposal) ContentTimeForTransition() Time /* not a class type */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("contentTimeForTransition"))
	return rv
}

// Optional custom metadata associated with the proposed item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposal/metadata
func (c_ ContentProposal) Metadata() []objc.IObject /* cross-framework: MetadataItem */ {
	rv := objc.Send[[]avfoundation.MetadataItem](c_.ID, objc.Sel("metadata"))
	return rv
}
func (c_ ContentProposal) SetMetadata(value []objc.IObject /* cross-framework: MetadataItem */) {
	c_.ID.Send(objc.RegisterName("setMetadata:"), value)
}

// The preview image of the proposed item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposal/previewImage
func (c_ ContentProposal) PreviewImage() objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](c_.ID, objc.Sel("previewImage"))
	return rv
}

// The title of the proposed content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposal/title
func (c_ ContentProposal) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("title"))
	return rv
}

// The URL of the proposed content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposal/url
func (c_ ContentProposal) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("URL"))
	return rv
}
func (c_ ContentProposal) SetURL(value objc.IObject /* cross-framework: NSURL */) {
	c_.ID.Send(objc.RegisterName("setURL:"), value)
}




