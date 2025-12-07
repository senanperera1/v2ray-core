package main

// DefaultPageSize defines the memory page size for Android 15+
var DefaultPageSize = 16 * 1024 // 16 KB

// GetPageSize returns the page size
func GetPageSize() int {
    return DefaultPageSize
}
