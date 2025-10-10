// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

import (
	"github.com/ebitengine/purego"
)

var lib uintptr

func init() {
	var err error
	lib, err = purego.Dlopen("/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

// Functions from CoreGraphics
// CGContextSetInterpolationQuality
// CGBitmapContextCreateAdaptive
// CGBitmapInfoMake
// contentHeadroom
// 
// 
// 
// 
// synchronizeAttributes
// CGContextGetContentToneMappingInfo
// CGContextSetContentToneMappingInfo
// CGEXRToneMappingGammaGetDefaultOptions
// contentHeadroom
// 
// CGRenderingBufferLockBytePtr
// CGRenderingBufferProviderCreate
// CGRenderingBufferProviderCreateWithCFData
// CGRenderingBufferProviderGetSize
// CGRenderingBufferProviderGetTypeID
// CGRenderingBufferUnlockBytePtr
// contentHeadroom
// 
// 
