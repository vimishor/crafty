---@meta

---@class download_options : table
---@field checksum string
---@field silent boolean

---@class http_util
local _M = {}

---@param url string
---@param fpath? string # Path to file location (default: $XDG_CACHE_HOME/crafty/download/)
---@param opts download_options?
---@return string path # Full path to downloaded file. Empty on error
---@return string|nil error # Error message
function _M.file_download(url, fpath, opts) end

return _M
