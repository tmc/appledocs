// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit_test

import (
	"github.com/tmc/appledocs/generated/webkit"
)

// Suppress unused import errors
var _ = webkit.NewWebView

// ExampleWebView_DeleteSelection demonstrates using DeleteSelection on a WebView instance.
// Deletes the receiver’s current selection unless it’s collapsed.
func ExampleWebView_DeleteSelection() {
	obj := webkit.NewWebView()
	obj.DeleteSelection()
	// Output:
}
