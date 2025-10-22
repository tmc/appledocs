// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer_test

import (
	"github.com/tmc/appledocs/generated/mediaplayer"
)

// Suppress unused import errors
var _ = mediaplayer.NewContentItem

// ExampleNewContentItemWithIdentifier demonstrates how to create a ContentItem instance using NewContentItemWithIdentifier.
// Sets the identifier for a media item.
func ExampleNewContentItemWithIdentifier() {
	_ = mediaplayer.NewContentItemWithIdentifier(
		"identifier", // identifier string
	)
	// Output:
}
