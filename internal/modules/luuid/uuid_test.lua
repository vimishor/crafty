local lust = require "lust"
local describe = lust.describe
local it = lust.it
local expect = lust.expect
local uuid = require("uuid")

describe("uuid module", function()
    it("create V1", function ()
      local token, err = uuid.V1()
      expect(err).to.be(nil)
      expect(token).to.be.truthy()
    end)

    it("create V4", function ()
      local token, err = uuid.V4()
      expect(err).to.be(nil)
      expect(token).to.be.truthy()
    end)

    it("create V6", function ()
      local token, err = uuid.V6()
      expect(err).to.be(nil)
      expect(token).to.be.truthy()
    end)

    it("create V7", function ()
      local token, err1 = uuid.V7()
      expect(err1).to.be(nil)
      expect(token).to.be.truthy()

      local ok, err2 = uuid.is_valid(token)
      expect(err2).to.be(nil)
      expect(ok).to.be.truthy()
    end)

    it("validates successfully a good V7 UUID", function ()
      local ok, err = uuid.is_valid("0196a4ef-ba2a-748d-9c1e-3ba0bd1add90")
      expect(err).to.be(nil)
      expect(ok).to.be.truthy()
    end)

    it("should not validate a malformed UUID", function ()
      expect(uuid.is_valid("0196a4ef-ba2a-748d-9c1e-3ba0bd1add9")).to.fail()
    end)

    it("should create a V4 as default", function ()
      local token, err = uuid()
      expect(err).to.be(nil)
      expect(token).to.be.truthy()
    end)
end)

