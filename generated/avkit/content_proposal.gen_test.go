// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit_test

import (
	"github.com/tmc/appledocs/generated/avkit"
)

// Suppress unused import errors
var _ = avkit.NewContentProposal

// ExampleNewContentProposalWithContentTimeForTransitionTitlePreviewImage demonstrates how to create a ContentProposal instance using NewContentProposalWithContentTimeForTransitionTitlePreviewImage.
// Creates a new content proposal with the specified transition time, title, and preview image.
func ExampleNewContentProposalWithContentTimeForTransitionTitlePreviewImage() {
	_ = avkit.NewContentProposalWithContentTimeForTransitionTitlePreviewImage(
		avkit.Time{}, // contentTimeForTransition Time
		"title", // title string
		avkit.Image{}, // previewImage Image
	)
	// Output:
}
