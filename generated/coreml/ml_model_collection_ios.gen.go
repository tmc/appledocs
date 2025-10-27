//go:build darwin && ios

// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ModelCollection


// iOS-only properties

// The unique identifier of the model collection’s deployment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/deploymentID
func (m_ ModelCollection) DeploymentID() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("deploymentID"))
	return rv
}

// A dictionary of model entries keyed to the models’ identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/entries
func (m_ ModelCollection) Entries() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("entries"))
	return rv
}

// The name of the model collection, unique to the development team.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/identifier
func (m_ ModelCollection) Identifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}





