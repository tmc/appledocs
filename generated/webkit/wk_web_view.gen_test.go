// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit_test

import (
	"github.com/tmc/appledocs/generated/webkit"
)

// Suppress unused import errors
var _ = webkit.NewWebView

// ExampleWebView_GoBack demonstrates using GoBack on a WebView instance.
// Navigates to the back item in the back-forward list.
func ExampleWebView_GoBack() {
	obj := webkit.NewWebView()
	_ = obj.GoBack()
	// Output:
}

// ExampleWebView_GoForward demonstrates using GoForward on a WebView instance.
// Navigates to the forward item in the back-forward list.
func ExampleWebView_GoForward() {
	obj := webkit.NewWebView()
	_ = obj.GoForward()
	// Output:
}

// ExampleWebView_Reload demonstrates using Reload on a WebView instance.
// Reloads the current webpage.
func ExampleWebView_Reload() {
	obj := webkit.NewWebView()
	_ = obj.Reload()
	// Output:
}

// ExampleWebView_ReloadFromOrigin demonstrates using ReloadFromOrigin on a WebView instance.
// Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
func ExampleWebView_ReloadFromOrigin() {
	obj := webkit.NewWebView()
	_ = obj.ReloadFromOrigin()
	// Output:
}

// ExampleWebView_StopLoading demonstrates using StopLoading on a WebView instance.
// Stops loading all resources on the current page.
func ExampleWebView_StopLoading() {
	obj := webkit.NewWebView()
	obj.StopLoading()
	// Output:
}
