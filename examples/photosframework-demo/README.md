# photosframework Demo

Demonstrates the idiomatic Go API for Photos Swift extensions.

## Comparison

### Before (Raw FFI)

From `examples/photos-swift-bindings/`:

```go
// Raw unsafe.Pointer usage
var photosSharedLibrary func() unsafe.Pointer
library := photosSharedLibrary()
defer photosRelease(library)

// No type safety
var photosFetchResultCount func(unsafe.Pointer) int
count := photosFetchResultCount(somePointer)
```

### After (Type-Safe API)

From this example:

```go
// Clean, type-safe Go API
library := photosframework.SharedPhotoLibrary()
defer library.Release()

// Type-safe methods
count := fetchResult.Count()

// Iterator support
iter := fetchResult.Iterator()
for change := iter.Next(); change != nil; change = iter.Next() {
    // Process change
}
```

## Benefits

1. **Type Safety** - No unsafe.Pointer in application code
2. **Memory Management** - Automatic via finalizers
3. **Idiomatic Go** - Methods on structs, not function pointers
4. **Documentation** - Clear API with godoc
5. **Examples** - Built-in example tests

## Usage

```bash
# Copy library to current directory
cp ../../generated/swift/frameworks/photos/.build/release/libPhotosSwift.dylib .

# Run demo
go run main.go
```

## Expected Output

```
=== Photos Framework (Swift Extensions) Demo ===

Test 1: Testing Swift wrapper library
Hello from Photos Swift wrapper!
PHPhotoLibrary is available: PHPhotoLibrary

Test 2: Getting shared photo library
✓ Got photo library: 0x...

Test 3: Type-safe Swift-like API
Available types:
  • photosframework.PhotoLibrary
  • photosframework.FetchResult
  • photosframework.ProjectChangeRequest
  • photosframework.PersistentChangeFetchResult
  • photosframework.PersistentChangeIterator
  • photosframework.PersistentChange

Test 4: Swift extension methods available
  • ProjectChangeRequest.RemoveAssets(assets)
  • PersistentChangeFetchResult.Iterator()
  • PersistentChangeIterator.Next()

=== Success! ===
```

## Package Structure

```
photosframework/
├── photos.go           # Main API
├── example_test.go     # Example usage
├── go.mod
└── README.md

examples/photosframework-demo/
├── main.go             # This demo
├── go.mod
└── README.md
```

## Next Steps

This pattern can be applied to generate packages for:
- `speechframework` (Speech recognition)
- `combineframework` (Reactive programming)
- `swiftuiframework` (SwiftUI views)

All with the same clean, type-safe Go API!
