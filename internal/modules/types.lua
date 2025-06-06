---@meta

--
-- In order to maintain compatibility with LUA, numeric
-- data types are used only inside the API, in order to
-- be more specific about the usage and expectations.
--

---@alias uint8 number the set of all unsigned  8-bit integers (0 to 255)
---@alias uint16 number the set of all unsigned 16-bit integers (0 to 65535)
---@alias uint32 number the set of all unsigned 32-bit integers (0 to 4294967295)
---@alias uint64 number the set of all unsigned 64-bit integers (0 to 18446744073709551615)

---@alias int8 number the set of all signed  8-bit integers (-128 to 127)
---@alias int16 number the set of all signed 16-bit integers (-32768 to 32767)
---@alias int32 number the set of all signed 32-bit integers (-2147483648 to 2147483647)
---@alias int64 number the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

---@alias float32 number the set of all IEEE 754 32-bit floating-point numbers
---@alias float64 number the set of all IEEE 754 64-bit floating-point numbers

---@alias byte uint8 alias for uint8
---@alias uint uint32 alias for uint32
---@alias int int32 # alias for int32

---@alias bool boolean

---@class network
network = {}

---@enum (key) network.STATUS
network.STATUS = {
  Unknown = 0,
  Connected = 1,
  Disconnected = 2,
}
