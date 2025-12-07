package main

// SetProtectCallback allows Android to register a callback for protecting sockets
var protectCallback func(fd int) bool

// SetProtectCallback registers a callback function to mark sockets as protected
func SetProtectCallback(cb func(fd int) bool) {
    protectCallback = cb
}

// ProtectSocket should be called inside your network dialer
func ProtectSocket(fd int) bool {
    if protectCallback != nil {
        return protectCallback(fd)
    }
    return false
}
