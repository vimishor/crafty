---@meta

---
--- Filesystem utilities module for LUA.
---

---@class fs
local _M = {}

---@param filepath string
---@return boolean
function _M.is_file(filepath) end

---@param src string
---@param dst string
---@return string|nil err
---@nodiscard
function _M.copy_file(src, dst) end

---@param src string
---@param dst string
---@return string|nil err
---@nodiscard
function _M.move_file(src, dst) end

---@param path string
---@return boolean
function _M.is_dir(path) end

---@param src string
---@param dst string
---@return string|nil err
---@nodiscard
function _M.copy_dir(src, dst) end

---@param path string
---@return table<int, string>
function _M.list_dir(path) end

---
--- Create directory recursively at given path.
---
---@param path string
---@return string|nil err
function _M.mkdir(path) end

---@param path string # File or directory
---@param mode string # File mode
---@return string|nil err
function _M.chmod(path, mode) end

return _M
