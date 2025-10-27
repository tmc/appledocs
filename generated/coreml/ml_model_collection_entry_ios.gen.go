//go:build darwin && ios

// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ModelCollectionEntry


// iOS-only properties

// The name of the model, which is unique to the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/Entry/modelIdentifier
func (m_ ModelCollectionEntry) ModelIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("modelIdentifier"))
	return rv
}

// The compiled model’s location on the device’s file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/Entry/modelURL
func (m_ ModelCollectionEntry) ModelURL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("modelURL"))
	return rv
}





