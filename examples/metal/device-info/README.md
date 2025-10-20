# Metal Device Information Example

This example demonstrates how to query information about Metal-capable GPU devices on macOS.

## What it demonstrates

- Getting the system's default Metal device
- Querying device properties (name, type, capabilities)
- Checking hardware features (raytracing, function pointers, MSAA support)
- Working with Metal through the Objective-C bridge
- Proper memory management with retain/release

## Running the example

```bash
go run main.go
```

## Expected Output

```
Metal Device Information
========================

Device Name: Apple M1 Pro
Headless: false
Low Power: false
Removable: false
Registry ID: 4294968320
Max Threads Per Threadgroup: 1024

Feature Support:
  Raytracing: true
  Function Pointers: true
  4x MSAA: true

Recommended Max Working Set Size: 21845 MB (21.33 GB)

Metal device information retrieved successfully!
```

(Output will vary based on your GPU)

## Key Concepts

### Metal Device

A Metal device represents a GPU that can execute Metal commands. Most Macs have at least one Metal-capable device.

### Device Types

- **Low Power**: Integrated GPUs (like Apple Silicon integrated GPU)
- **High Performance**: Discrete GPUs (like AMD Radeon)
- **Headless**: Compute-only GPUs without display output
- **Removable**: External GPUs connected via Thunderbolt

### Device Capabilities

Different Metal devices support different features:
- Raytracing (Apple Silicon M3+, some AMD GPUs)
- Function pointers (shader-level function calls)
- Texture sample counts (MSAA support)
- Working set size (maximum recommended GPU memory usage)

### Memory Management

When working with Objective-C objects through purego:
1. Objects returned from creation functions are retained
2. Must call `release` when done to avoid memory leaks
3. Use `defer` to ensure cleanup happens

## Common Use Cases

- Selecting appropriate GPU for rendering or compute tasks
- Checking feature availability before using advanced Metal features
- Optimizing memory usage based on device capabilities
- Debugging GPU-related issues

## References

- [MTLDevice Documentation](https://developer.apple.com/documentation/metal/mtldevice)
- [Metal Feature Set Tables](https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf)
- [Metal Best Practices Guide](https://developer.apple.com/documentation/metal/gpu_features/understanding_gpu_family)
