// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets_test

import (
	"github.com/tmc/appledocs/generated/backgroundassets"
)

// Suppress unused import errors
var _ = backgroundassets.NewBAURLDownload

// ExampleNewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifier demonstrates how to create a BAURLDownload instance using NewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifier.
// Creates a download that uses the specified identifier and App Group.
func ExampleNewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifier() {
	_ = backgroundassets.NewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifier(
		"identifier", // identifier string
		backgroundassets.URLRequest{}, // request URLRequest
		"applicationGroupIdentifier", // applicationGroupIdentifier string
	)
	// Output:
}
// ExampleNewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifierPriority demonstrates how to create a BAURLDownload instance using NewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifierPriority.
// Creates a prioritized download that uses the specified identifier and App Group.
func ExampleNewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifierPriority() {
	_ = backgroundassets.NewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifierPriority(
		"identifier", // identifier string
		backgroundassets.URLRequest{}, // request URLRequest
		"applicationGroupIdentifier", // applicationGroupIdentifier string
		backgroundassets.BADownloaderPriority{}, // priority BADownloaderPriority
	)
	// Output:
}
// ExampleNewBAURLDownloadWithIdentifierRequestEssentialFileSizeApplicationGroupIdentifierPriority demonstrates how to create a BAURLDownload instance using NewBAURLDownloadWithIdentifierRequestEssentialFileSizeApplicationGroupIdentifierPriority.
func ExampleNewBAURLDownloadWithIdentifierRequestEssentialFileSizeApplicationGroupIdentifierPriority() {
	_ = backgroundassets.NewBAURLDownloadWithIdentifierRequestEssentialFileSizeApplicationGroupIdentifierPriority(
		"identifier", // identifier string
		backgroundassets.URLRequest{}, // request URLRequest
		false, // essential bool
		0, // fileSize uint
		"applicationGroupIdentifier", // applicationGroupIdentifier string
		backgroundassets.BADownloaderPriority{}, // priority BADownloaderPriority
	)
	// Output:
}
// ExampleNewBAURLDownloadWithIdentifierRequestFileSizeApplicationGroupIdentifier demonstrates how to create a BAURLDownload instance using NewBAURLDownloadWithIdentifierRequestFileSizeApplicationGroupIdentifier.
func ExampleNewBAURLDownloadWithIdentifierRequestFileSizeApplicationGroupIdentifier() {
	_ = backgroundassets.NewBAURLDownloadWithIdentifierRequestFileSizeApplicationGroupIdentifier(
		"identifier", // identifier string
		backgroundassets.URLRequest{}, // request URLRequest
		0, // fileSize uint
		"applicationGroupIdentifier", // applicationGroupIdentifier string
	)
	// Output:
}

