// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/appkit"
)

// PWebDownloadDelegate is the WebDownloadDelegate protocol interface.
//
// The   protocol declares one additional method for delegates of  .
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebDownloadDelegate
type PWebDownloadDelegate interface {
	// Optional methods
	DownloadWindowForAuthenticationSheet(download IWebDownload) appkit.Window
	HasDownloadWindowForAuthenticationSheet() bool
}

// WebDownloadDelegate is a delegate implementation builder for the PWebDownloadDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type WebDownloadDelegate struct {
	_DownloadWindowForAuthenticationSheet func(download IWebDownload) appkit.Window
}

// SetDownloadWindowForAuthenticationSheet sets the handler for the DownloadWindowForAuthenticationSheet delegate method.
//
// Returns the window to be used by the authentication sheet.
func (d *WebDownloadDelegate) SetDownloadWindowForAuthenticationSheet(f func(download IWebDownload) appkit.Window) {
	d._DownloadWindowForAuthenticationSheet = f
}

// DownloadWindowForAuthenticationSheet implements the PWebDownloadDelegate interface.
func (d *WebDownloadDelegate) DownloadWindowForAuthenticationSheet(download IWebDownload) appkit.Window {
	if d._DownloadWindowForAuthenticationSheet != nil {
		return d._DownloadWindowForAuthenticationSheet(download)
	}
	var zero appkit.Window
	return zero
}

// HasDownloadWindowForAuthenticationSheet returns true if a handler for DownloadWindowForAuthenticationSheet has been set.
func (d *WebDownloadDelegate) HasDownloadWindowForAuthenticationSheet() bool {
	return d._DownloadWindowForAuthenticationSheet != nil
}
