local archive = require("archive")

local BIN_PATH = HOME .. "/.local/bin"
local MAN_PATH = HOME .. "/.local/share/man"
local CACHE_PATH = XDG_CACHE_HOME .. "/download"

-- Apps version
local RG_VERSION = "14.1.1"

local err = archive.extract(CACHE_PATH .. "/ripgrep-" .. RG_VERSION .. "-x86_64-unknown-linux-musl.tar.gz", "/home/alex/.tmp/test_extract_1")
-- if err then
--   print("[ERROR] " .. err)
--   return 1
-- end
-- print("[OK] " .. CACHE_PATH .. "/ripgrep-" .. RG_VERSION .. "-x86_64-unknown-linux-musl.tar.gz")

err = archive.decompress(CACHE_PATH .. "/restic_0.18.0_linux_amd64.bz2", "/home/alex/.tmp/restic_test")
if err then
  print("[ERROR] " .. err)
  return 1
end

err = archive.pick(CACHE_PATH .. "/ripgrep-" .. RG_VERSION .. "-x86_64-unknown-linux-musl.tar.gz", {
  ["/rg$"] = BIN_PATH .. "/rg",
  ["/doc/rg\\.1$"] = MAN_PATH .. "/man1/rg.1",
})
if err then
  print("[ERROR] " .. err)
  return 1
end
