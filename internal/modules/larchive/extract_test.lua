local lust = require("lust")
local describe = lust.describe
local it = lust.it
local expect = lust.expect

local archive = require("archive")
local pwd = "internal/modules/larchive"

describe("archive", function()
  it("Should extract everything from good archive", function()
    local err = nil
    local files = {
      "testdata/tar/good1.tar",
      "testdata/zip/good1.zip",
    }

    for _, file in pairs(files) do
      err = archive.extract(pwd .. "/" .. file, "/some/dir", { ["dry-run"] = true })
      expect(err).to_not.be.truthy()
    end
  end)

  it("Should not extract file with absolute path", function()
    local err = nil
    local files = {
      "testdata/tar/absolute1.tar",
      "testdata/tar/absolute2.tar",
      "testdata/zip/absolute1.zip",
      "testdata/zip/absolute2.zip",
    }

    for _, file in pairs(files) do
      err = archive.extract(pwd .. "/" .. file, "/some/dir", { ["dry-run"] = true })
      expect(err).to.be.truthy()
      expect(err).to.match("absolute file path")
    end
  end)

  it("Should not extract file with relative path", function()
    local err = nil
    local files = {
      "testdata/tar/relative0.tar",
      "testdata/tar/relative2.tar",
      "testdata/zip/relative0.zip",
      "testdata/zip/relative2.zip",
    }

    for _, file in pairs(files) do
      err = archive.extract(pwd .. "/" .. file, "/some/dir", { ["dry-run"] = true })
      expect(err).to.be.truthy()
      expect(err).to.match("relative file path")
    end
  end)

  it("Should not extract file with symlink", function()
    local err = nil
    local files = {
      "testdata/tar/dirsymlink.tar",
      "testdata/tar/dirsymlink2a.tar",
      "testdata/tar/dirsymlink2b.tar",
      "testdata/tar/symlink.tar",
      "testdata/zip/dirsymlink.zip",
      "testdata/zip/dirsymlink2a.zip",
      "testdata/zip/dirsymlink2b.zip",
      "testdata/zip/symlink.zip",
    }

    for _, file in pairs(files) do
      err = archive.extract(pwd .. "/" .. file, "/some/dir", { ["dry-run"] = true })
      expect(err).to.be.truthy()
      expect(err).to.match("symlink in file path")
    end
  end)
end)
