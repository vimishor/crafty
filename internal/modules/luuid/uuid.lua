---@meta

---
--- Generate and validate UUID V1, V4, V6 and V7
---

---@class uuid
local _M = {}

---
--- Generate UUID V1
---
---@return string
---@return string|nil err
function _M.V1() end

---
--- Generate UUID V4
---
---@return string
---@return string|nil err
function _M.V4() end

---
--- Generate UUID V6
---
---@return string
---@return string|nil err
function _M.V6() end

---
--- Generate UUID V7
---
---@return string
---@return string|nil err
function _M.V7() end

---
--- Validate a UUID
---
---@param s string
---@return boolean
---@return (string|nil)? err
function _M.is_valid(s) end


local mt = {
  __call = _M.V4,
}

return setmetatable(_M, mt)
