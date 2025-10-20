# IOSurface Framework Example

This example demonstrates the IOSurface framework for efficient buffer sharing in Go.

## What it demonstrates

- IOSurface creation
- Common use cases
- Surface properties
- Typical workflow
- Framework integration points
- Performance benefits
- Common pixel formats

## Running the example

```bash
go run main.go
# or with e2e flag
go run main.go -e2e
```

## Key Concepts

### IOSurface

IOSurface is a low-level framework for sharing image buffers between:
- CPU and GPU (zero-copy)
- Different processes
- Different frameworks (Metal, Core Animation, Core Image, etc.)

### Why Use IOSurface?

**Without IOSurface:**
- Must copy data between CPU and GPU
- Multiple copies for process sharing
- High memory bandwidth usage
- Inefficient video pipelines

**With IOSurface:**
- Zero-copy sharing (map same memory)
- Direct inter-process sharing
- Hardware-accelerated access
- Efficient video processing

### Creating an IOSurface

```go
// Create with properties dictionary
properties := map[string]interface{}{
    "IOSurfaceWidth": 1920,
    "IOSurfaceHeight": 1080,
    "IOSurfacePixelFormat": 'BGRA', // 32-bit BGRA
    "IOSurfaceBytesPerRow": 1920 * 4,
}

surface := iosurface.NewSurfaceWithProperties(properties)
```

### Surface Properties

Key properties to set when creating:
- **Width/Height**: Dimensions in pixels
- **PixelFormat**: How pixels are stored
- **BytesPerRow**: Stride (may be larger than width * bpp)
- **AllocSize**: Total memory allocation
- **BytesPerElement**: Bytes per pixel

## Use Cases

### GPU-CPU Buffer Sharing

```go
// 1. Create surface
surface := iosurface.NewSurfaceWithProperties(props)

// 2. Render with Metal
metalTexture := device.NewTextureWithIOSurface(surface)
// Render to texture...

// 3. Read from CPU
surface.Lock()
data := surface.BaseAddress()
// Process data...
surface.Unlock()
```

### Inter-Process Sharing

```go
// Process A: Create and share
surface := iosurface.NewSurfaceWithProperties(props)
surfaceID := surface.GetID()
// Send surfaceID to Process B via IPC

// Process B: Access same surface
surface := iosurface.LookupFromID(surfaceID)
// Access same memory as Process A
```

### Video Frame Buffers

```go
// Create surface for video frame
props := map[string]interface{}{
    "IOSurfaceWidth": 1920,
    "IOSurfaceHeight": 1080,
    "IOSurfacePixelFormat": '420v', // YUV 420
}

surface := iosurface.NewSurfaceWithProperties(props)

// Use with AVFoundation
pixelBuffer := CVPixelBufferFromIOSurface(surface)
```

## Framework Integration

### Metal

```go
// Create Metal texture from IOSurface
let descriptor = MTLTextureDescriptor()
descriptor.width = surface.width
descriptor.height = surface.height
descriptor.pixelFormat = .bgra8Unorm

let texture = device.makeTexture(
    descriptor: descriptor,
    iosurface: surface,
    plane: 0
)
```

### Core Animation

```go
// Use IOSurface as CALayer backing
layer.contents = surface
// or
layer.contentsIOSurface = surface
```

### Core Image

```go
// Create CIImage from IOSurface
ciImage := CIImage.init(ioSurface: surface)
```

## Pixel Formats

### RGB/BGRA Formats

| Format | Description | Bytes/Pixel |
|--------|-------------|-------------|
| 'BGRA' | 32-bit BGRA | 4 |
| 'RGBA' | 32-bit RGBA | 4 |
| 'RGBx' | 32-bit RGB + padding | 4 |
| 'BGRx' | 32-bit BGR + padding | 4 |

### YUV/YCbCr Formats

| Format | Description | Use Case |
|--------|-------------|----------|
| '2vuy' | 422 YCbCr 8-bit | Video processing |
| '420v' | 420 YCbCr planar | Video compression |
| '420f' | 420 YCbCr biplanar | Modern video |

## Performance Benefits

1. **Zero-Copy**: GPU and CPU access same memory
2. **Direct Mapping**: No data copying between processes
3. **Page-Aligned**: Optimal memory access patterns
4. **Hardware Acceleration**: GPU can render directly
5. **Efficient Video**: Minimal copies in video pipeline

## Thread Safety

- IOSurface is thread-safe for most operations
- Lock/unlock for CPU access coordination
- Multiple readers OK, exclusive writer needed

## Memory Management

IOSurfaces use reference counting:
- Retain when keeping reference
- Release when done
- Automatic cleanup when refcount reaches 0

## Best Practices

1. **Alignment**: Ensure proper byte alignment for performance
2. **Locking**: Always lock before CPU access, unlock after
3. **Cleanup**: Release surfaces when done
4. **Format Match**: Ensure all users agree on pixel format
5. **Size Limits**: Check hardware limits for dimensions

## Common Workflows

### Render and Capture

```go
// 1. Create surface
surface := iosurface.NewSurfaceWithProperties(props)

// 2. Render with GPU
renderWithMetal(surface)

// 3. Capture CPU-side
surface.Lock()
data := copyPixelData(surface)
surface.Unlock()

// 4. Process data
processImage(data)
```

### Screen Capture

```go
// Create surface for screen capture
surface := iosurface.NewSurfaceWithProperties(screenProps)

// Capture screen content
captureScreen(surface)

// Share with display or encoder
displayLayer.contents = surface
// or
encodeFrame(surface)
```

## Integration with Other APIs

### AVFoundation Video

IOSurface is ideal for video:
- Camera preview buffers
- Video encoding/decoding
- Real-time video effects
- Video frame processing

### Core Animation

Efficient layer backing:
- No copying for layer content
- Direct GPU rendering
- Smooth animations
- Efficient compositing

### Screen Recording

High-performance capture:
- ScreenCaptureKit uses IOSurface
- Minimal overhead
- Direct GPU capture
- Efficient encoding pipeline

## Limitations

- Platform-specific (macOS/iOS)
- Requires understanding of memory management
- Must coordinate locking between users
- Size limits based on hardware
- Pixel format compatibility required

## References

- [IOSurface Documentation](https://developer.apple.com/documentation/iosurface)
- [IOSurface API Reference](https://developer.apple.com/documentation/iosurface/iosurfaceref)
- [Metal IOSurface Integration](https://developer.apple.com/documentation/metal/mtltexture)
- [Core Animation IOSurface](https://developer.apple.com/documentation/quartzcore/calayer)
