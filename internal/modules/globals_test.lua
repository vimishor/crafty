local lust = require "lust"
local describe = lust.describe
local it = lust.it
local expect = lust.expect

describe("globals", function ()
  it("has USER", function ()
    expect(USER).to.be.a("string")
    expect(USER).to.be.truthy()
  end)

  it("has HOME", function ()
    expect(HOME).to.be.a("string")
    expect(HOME).to.be.truthy()
  end)

  it("has XDG_CONFIG_HOME", function ()
    expect(XDG_CONFIG_HOME).to.be.a("string")
    expect(XDG_CONFIG_HOME).to.be.truthy()
  end)

  it("has XDG_CACHE_HOME", function ()
    expect(XDG_CACHE_HOME).to.be.a("string")
    expect(XDG_CACHE_HOME).to.be.truthy()
  end)

  it("has XDG_DATA_HOME", function ()
    expect(XDG_DATA_HOME).to.be.a("string")
    expect(XDG_DATA_HOME).to.be.truthy()
  end)
end)
