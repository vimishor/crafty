---@meta

---@class ExtractOpts
---@field dry-run boolean # Default false
---@field allow-symlink boolean # Default false

---@class archive
local _M = {}

---@param src string # Path to archive
---@param dst string # A directory where the contents of the archive will be extracted
---@param opts ExtractOpts? # Extract options
---@return string|nil err
---@nodiscard
function _M.extract(src, dst, opts) end

---@param src string # Path to archive
---@param dst string # A directory where the contents of the archive will be decompressed
---@return string|nil err
---@nodiscard
function _M.decompress(src, dst) end

---
--- cherry-pick certain files from an archive
---
---@param src string # Path to archive
---@param filters table # List of files to be extracted and their destination
---@return string|nil err
---@nodiscard
function _M.pick(src, filters) end

return _M
