---@meta

---
--- Time library
---
--- Offers time functionality based on Golang's `time` package.
---
---@see https://github.com/vadv/gopher-lua-libs/tree/v0.5.0/time
---

---@class time
local _M = {}

---@return number unix timestamp in seconds
function _M.unix() end

---@return number unix timestamp in nano seconds
function _M.unix_nano() end

---@param duration number How many seconds to sleep
function _M.sleep(duration) end

---@param value string
---@param layout string
---@param location string?
---@return number
---@return (string|nil)? err
function _M.parse(value, layout, location) end

---@param value number
---@param layout string
---@param location string?
---@return string
---@return (string|nil)? err
function _M.format(value, layout, location) end

return _M
