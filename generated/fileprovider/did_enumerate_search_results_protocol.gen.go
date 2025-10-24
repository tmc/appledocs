// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// didEnumerateSearchResultsProtocol is the didEnumerateSearchResults: protocol.
//
// Availability:
//   - macOS 26.0+
//
// Use this protocol when registering custom classes that conform to didEnumerateSearchResults:.
var didEnumerateSearchResultsProtocol *objc.Protocol

func init() {
	didEnumerateSearchResultsProtocol = objc.GetProtocol("didEnumerateSearchResults:")
}

