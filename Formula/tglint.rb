class Tglint < Formula
  desc "Formatter and linter for Terragrunt HCL files"
  homepage "https://github.com/roman-senchuk/tglint"
  url "https://github.com/roman-senchuk/tglint/archive/v1.0.0.tar.gz"
  sha256 ""
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin/"tglint", "."
  end

  test do
    system "#{bin}/tglint", "--version"
  end
end
