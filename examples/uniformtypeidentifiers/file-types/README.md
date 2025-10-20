# UniformTypeIdentifiers (UTI) File Types Example

This example demonstrates how to work with Uniform Type Identifiers in Go.

## What it demonstrates

- Creating UTType objects from file extensions
- Creating UTType objects from MIME types
- Creating UTType objects from type identifiers
- Working with common file formats
- Understanding document and media types

## Running the example

```bash
go run main.go
```

## Key Concepts

### Uniform Type Identifiers (UTI)

UTIs provide a consistent way to identify data types across macOS and iOS:
- File types (based on extension)
- MIME types (internet media types)
- Pasteboard types (for copy/paste)
- Drag & drop types

### Type Hierarchies

UTIs form a hierarchy:
- `public.data` (base type for all data)
  - `public.content` (human-readable data)
    - `public.text` (text data)
      - `public.plain-text` (plain text files)
      - `public.source-code` (source code)
    - `public.image` (image data)
      - `public.jpeg` (JPEG images)
      - `public.png` (PNG images)

### Creating UTTypes

Three main ways to create UTType objects:

```go
// From file extension
utType := uniformtypeidentifiers.NewUTTypeWithFilenameExtension("pdf")

// From MIME type
utType := uniformtypeidentifiers.NewUTTypeWithMIMEType("application/json")

// From type identifier
utType := uniformtypeidentifiers.NewUTTypeWithIdentifier("public.text")
```

### Common Type Identifiers

- `public.data` - Generic data
- `public.text` - Text documents
- `public.plain-text` - Plain text
- `public.image` - Images
- `public.audio` - Audio files
- `public.video` - Video files
- `public.archive` - Archive files (zip, tar, etc.)
- `public.executable` - Executable programs

## Use Cases

### File Type Detection

Identify file types in your application:

```go
func getFileType(filename string) UTType {
    ext := filepath.Ext(filename)
    return uniformtypeidentifiers.NewUTTypeWithFilenameExtension(ext)
}
```

### File Open Dialogs

Specify allowed file types in open panels:

```go
// Allow only image files
imageType := uniformtypeidentifiers.NewUTTypeWithIdentifier("public.image")
// Use with NSOpenPanel.allowedContentTypes
```

### Drag & Drop

Validate dragged files:

```go
// Check if dragged file is an image
draggedType := /* get from drag operation */
imageType := uniformtypeidentifiers.NewUTTypeWithIdentifier("public.image")
// Check conformance
```

### Document Type Registration

Declare supported document types in your app:

```go
// Register as handler for .myapp files
myAppType := uniformtypeidentifiers.NewUTTypeWithFilenameExtension("myapp")
```

## Type Conformance

UTTypes support hierarchical type checking:
- A JPEG conforms to "public.image"
- A Swift file conforms to "public.source-code"
- A source code file conforms to "public.text"

Use the `ConformsTo` method to check type relationships.

## Best Practices

1. **Use type identifiers over extensions** when possible
2. **Check type conformance** instead of exact matches
3. **Define custom types** for proprietary formats
4. **Register exported types** in Info.plist for custom formats
5. **Use imported types** for types defined by other apps

## References

- [Uniform Type Identifiers Documentation](https://developer.apple.com/documentation/uniformtypeidentifiers)
- [System-Declared Types](https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/UTIRef/Articles/System-DeclaredUniformTypeIdentifiers.html)
- [UTType Documentation](https://developer.apple.com/documentation/uniformtypeidentifiers/uttype)
