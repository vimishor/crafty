local fs = require("fs")

for _, file in pairs(fs.list_dir("/etc")) do
  print(file)
end

local err = fs.mkdir("bla")
if err then
  print("[ERROR] " .. err)
  return 1
end
