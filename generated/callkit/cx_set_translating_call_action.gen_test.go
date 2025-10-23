// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXSetTranslatingCallAction

// ExampleNewCXSetTranslatingCallActionWithCallUUIDIsTranslatingLocalLanguageRemoteLanguage demonstrates how to create a CXSetTranslatingCallAction instance using NewCXSetTranslatingCallActionWithCallUUIDIsTranslatingLocalLanguageRemoteLanguage.
func ExampleNewCXSetTranslatingCallActionWithCallUUIDIsTranslatingLocalLanguageRemoteLanguage() {
	_ = callkit.NewCXSetTranslatingCallActionWithCallUUIDIsTranslatingLocalLanguageRemoteLanguage(
		callkit.UUID{}, // uuid UUID
		false, // isTranslating bool
		"localLanguage", // localLanguage string
		"remoteLanguage", // remoteLanguage string
	)
	// Output:
}
// ExampleNewCXSetTranslatingCallActionWithCoder demonstrates how to create a CXSetTranslatingCallAction instance using NewCXSetTranslatingCallActionWithCoder.
// Creates a new action to start or stop translating a call with the provided data.
func ExampleNewCXSetTranslatingCallActionWithCoder() {
	_ = callkit.NewCXSetTranslatingCallActionWithCoder(
		callkit.Coder{}, // aDecoder Coder
	)
	// Output:
}
