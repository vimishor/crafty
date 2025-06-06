---@meta

---@alias runtime.platform.OS
---| '"darwin"'
---| '"dragonfly"'
---| '"freebsd"'
---| '"linux"'
---| '"netbsd"'
---| '"openbsd"'
---| '"plan9"'
---| '"windows"'
---| '"unknown"'

---@alias runtime.platform.ARCH
---| '"386"'
---| '"amd64"'
---| '"arm"'
---| '"arm64"'
---| '"mips"'
---| '"mips64"'
---| '"ppc64"'
---| '"unknown"'

---
---@class runtime
runtime = {}

---
--- GOOS is the running program's operating system target: one of darwin, freebsd, linux, and so on.
---
--- [View Go doc](https://pkg.go.dev/runtime#pkg-constants)
---
---@type runtime.platform.OS
runtime.OS = "unknown"

---
--- GOARCH is the running program's architecture target: one of 386, amd64, arm, s390x, and so on.
---
--- [View Go doc](https://pkg.go.dev/runtime#pkg-constants)
---
---@type runtime.platform.ARCH
runtime.ARCH = "unknown"

---
---@type string
runtime.LUA_VERSION = "5.1"

return runtime
