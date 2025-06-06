local lust = require "lust"
local describe = lust.describe
local it = lust.it
local expect = lust.expect
local runtime = require("runtime")

describe("runtime module", function ()
  it("has OS", function ()
    expect(runtime.OS).to.be.a("string")
    expect(runtime.OS).to.be.truthy()
  end)

  it("has ARCH", function ()
    expect(runtime.ARCH).to.be.a("string")
    expect(runtime.ARCH).to.be.truthy()
  end)

  it("has LUA_VERSION", function ()
    expect(runtime.LUA_VERSION).to.be.a("string")
    expect(runtime.LUA_VERSION).to.be.truthy()
  end)
end)
