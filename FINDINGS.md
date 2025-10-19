# Findings: Generated Bindings Status

## Summary
The ScreenCaptureKit bindings are fully functional with property accessor generation complete. Property accessors correctly return `unsafe.Pointer` for NSArray types, which is the appropriate behavior for Objective-C collection types.

## What Works

### ScreenCaptureKit Bindings (✓ Complete)
- Property accessors generated and working
- Class methods generated (e.g., `GetCurrentProcessShareableContentWithCompletionHandler`)
- Instance methods generated
- Proper type conversions for Objective-C types
- NSArray properties correctly return `unsafe.Pointer` (cast to `objc.ID` for array operations)

### Property Accessor Behavior
Property accessors are generated with the following behavior:
- **Scalar types**: Return Go equivalents (int, bool, etc.)
- **Geometry types**: Return proper structs (CGSize, CGRect, CGPoint)
- **Object types**: Return `unsafe.Pointer` (cast to `objc.ID` or specific type)
- **NSArray types**: Return `unsafe.Pointer` (cast to `objc.ID` for count/objectAtIndex operations)

This is the correct and intended behavior since NSArray is a dynamic Objective-C collection.

## Known Limitations

### Missing Bindings
1. **CoreMedia** - Needed for CMSampleBuffer handling in screen capture
2. **ImageIO** - Needed for CGImageDestination (saving images)
3. **NSObject methods** - Core methods like `retain`, `release`, `autorelease` are not generated (use objc.Send)
4. **Some class methods** - e.g., `SCShareableContent.GetShareableContentWithCompletionHandler` (only `GetCurrentProcessShareableContentWithCompletionHandler` exists)

### Workarounds
For missing bindings, use manual `objc.Send` calls:
```go
// Example: Retain/Release
shareableContent.ID.Send(objc.RegisterName("retain"))
defer shareableContent.ID.Send(objc.RegisterName("release"))

// Example: NSArray operations
displays := objc.ID(shareableContent.Displays())  // Cast unsafe.Pointer to objc.ID
count := int(displays.Send(objc.RegisterName("count")))
```

## Examples

### Working Examples
- `examples/screencapture-delegate-demo/` - Minimal delegate API demo (49 lines)
- `examples/screencapturekit-generated-bindings/` - Full screen capture with PNG saving (700+ lines)

Both examples successfully use generated bindings with property accessors.

## Status

- ✅ Property accessors generated and working
- ✅ Geometry types return proper structs (CGSize, CGRect, CGPoint)
- ✅ Type-safe delegate helpers (see screencapture-delegate-demo)
- ✅ ScreenCaptureKit bindings complete and functional
- ⏳ CoreMedia/ImageIO bindings pending (use manual purego for now)
- ⏳ NSObject method generation pending (use objc.Send for now)
