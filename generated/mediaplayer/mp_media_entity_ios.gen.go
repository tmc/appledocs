//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MediaEntity


// Executes a provided block with the fetched values for the given item properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaEntity/enumerateValues(forProperties:using:)
func (m_ MediaEntity) EnumerateValuesForPropertiesUsingBlock(properties unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enumerateValuesForProperties:usingBlock:"), properties, block)
}

// Returns the object specified by the key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaEntity/subscript(_:)
func (m_ MediaEntity) ObjectForKeyedSubscript(key objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}

// Retrieves the value for a specified media property key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaEntity/value(forProperty:)
func (m_ MediaEntity) ValueForProperty(property objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("valueForProperty:"), property)
	return rv
}

// iOS-only properties

// The persistent identifier for a media entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaEntity/persistentID
func (m_ MediaEntity) PersistentID() objc.IObject /* cross-framework: MediaEntityPersistentID */ {
	rv := objc.Send[MediaEntityPersistentID](m_.ID, objc.Sel("persistentID"))
	return rv
}





