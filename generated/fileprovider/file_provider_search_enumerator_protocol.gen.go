// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// FileProviderSearchEnumeratorProtocol is the NSFileProviderSearchEnumerator protocol.
//
// Availability:
//   - macOS 26.0+
//
// Use this protocol when registering custom classes that conform to NSFileProviderSearchEnumerator.
var FileProviderSearchEnumeratorProtocol *objc.Protocol

func init() {
	FileProviderSearchEnumeratorProtocol = objc.GetProtocol("NSFileProviderSearchEnumerator")
}
