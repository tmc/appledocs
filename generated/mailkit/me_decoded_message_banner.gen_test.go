// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit_test

import (
	"github.com/tmc/appledocs/generated/mailkit"
)

// Suppress unused import errors
var _ = mailkit.NewMEDecodedMessageBanner


// ExampleNewMEDecodedMessageBannerWithTitlePrimaryActionTitleDismissable demonstrates how to create a MEDecodedMessageBanner instance using NewMEDecodedMessageBannerWithTitlePrimaryActionTitleDismissable.
func ExampleNewMEDecodedMessageBannerWithTitlePrimaryActionTitleDismissable() {
	_ = mailkit.NewMEDecodedMessageBannerWithTitlePrimaryActionTitleDismissable(
		"title", // title string
		"primaryActionTitle", // primaryActionTitle string
		false, // dismissable bool
	)
	// Output:
}


