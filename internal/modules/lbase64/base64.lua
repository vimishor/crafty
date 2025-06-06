---@meta

---
--- base64 library
---

---
---@class base64
local _M = {}

---
--- Decode string `s` from base64.
---
--- New line characters (\r and \n) are ignored.
---
---@param s string
---@return string # Encoded string
function _M.encode(s) end

---
--- Perform URL-safe encoding
---
---@param s string
---@return string # Encoded URL
function _M.encode_url(s) end

---@param s string # Encoded string
---@return string # Decoded string
---@return string|nil err
function _M.decode(s) end

---@param s string # Encoded URL
---@return string # Decoded URL
---@return string|nil err
function _M.decode_url(s) end

return _M
