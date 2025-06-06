local json = require("json")
local http = require("http")
local client = http.client({
  user_agent = "hihi/v1"
})
local dump = require("pretty").print
local request = http.request("GET", "http://echo.free.beeceptor.com/asd?q")
local response, err = client:do_request(request)
if err then
  error(err)
end
--
if not (response.code == 200) then
  error("did not received a 200 response code")
end

-- print(dump(response.body))

local data = json.decode(response.body)
print(dump(data))
print("Host: " .. data.host)

local http_util = require("http_util")
local _, err2 = http_util.file_download("https://github.com/jqlang/jq/releases/download/jq-1.7.1/jq-linux-mips64", XDG_CACHE_HOME .. "/crafty/download", {
  checksum = "3d4b43d0ab571431fdfc786f951a0547a05f75aed5c076409ed2cb58b7244de1",
})
if err2 then
  error(err2)
  os.exit(1)
end
