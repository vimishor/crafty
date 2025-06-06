local lust = require "lust"
local describe = lust.describe
local it = lust.it
local expect = lust.expect
local base64 = require("base64")

describe("base64 module", function()
    it("encode successfully", function ()
      local b = base64.encode("hello world")
      expect(b).to.be("aGVsbG8gd29ybGQ=")
    end)

    it("decode successfully", function ()
      local b, err  = base64.decode("aGVsbG8gd29ybGQ=")
      expect(err).to.be(nil)
      expect(b).to.be("hello world")
    end)

    it("encode_url successfully", function ()
      local b = base64.encode_url("https://www.example.com/bla?key1=val1&key2=val2#boo")
      expect(b).to.be("aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vYmxhP2tleTE9dmFsMSZrZXkyPXZhbDIjYm9v")
    end)

    it("decode_url successfully", function ()
      local b, err  = base64.decode_url("aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vYmxhP2tleTE9dmFsMSZrZXkyPXZhbDIjYm9v")
      expect(err).to.be(nil)
      expect(b).to.be("https://www.example.com/bla?key1=val1&key2=val2#boo")
    end)
end)

