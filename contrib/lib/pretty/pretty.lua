local _M = {}

local table_prefix = "{"
local table_suffix = "}"
local separator = ": "

local function print_table(tbl, indent)
  local formatting = string.rep("  ", indent)
  for k, v in pairs(tbl) do
    if type(v) == "table" then
      print(formatting .. tostring(k) .. separator .. table_prefix)
      print_table(v, indent+1)
      print(formatting .. table_suffix)
    else
      print(formatting .. tostring(k) .. separator .. tostring(v))
    end
  end
end

function _M.print(data, indent)
  if not indent then indent = 0 end

  if type(data) == "table" then
    print_table(data, indent)
  else
    print(string.rep("  ", indent) .. tostring(data))
  end
end

return _M
