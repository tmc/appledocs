// Package objc provides a thin compatibility layer over github.com/ebitengine/purego/objc
// for Objective-C runtime integration.
//
// This package re-exports core purego/objc types and adds convenience functions for
// type conversions and common operations. All heavy lifting (FFI, objc_msgSend, etc.)
// is handled by the battle-tested purego/objc package.
//
// # Usage
//
// Import this package along with purego/objc for generated bindings:
//
//	import (
//	    "github.com/ebitengine/purego/objc"
//	    objchelper "github.com/tmc/appledocs/objc"
//	)
//
// The generated bindings use objc.ID, objc.Class, objc.SEL directly from purego/objc.
// This package provides additional helpers for string conversions and other common tasks.
package objc
