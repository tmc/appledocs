// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PositionalSpecifier] class.
var positionalSpecifierClass = _PositionalSpecifierClass{objc.GetClass("NSPositionalSpecifier")}

type _PositionalSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [PositionalSpecifier] class.
type IPositionalSpecifier interface {
	objectivec.IObject
}

// A specifier for an insertion point in a container relative to another object in the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier

type PositionalSpecifier struct {
	objectivec.Object
}

// PositionalSpecifierFrom constructs a [PositionalSpecifier] from an unsafe.Pointer.
//
// A specifier for an insertion point in a container relative to another object in the container.
func PositionalSpecifierFrom(ptr unsafe.Pointer) PositionalSpecifier {
	return PositionalSpecifier{objectivec.Object{objc.ID(ptr)}}
}



