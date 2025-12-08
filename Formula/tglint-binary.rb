# Binary-based formula (uses pre-built binaries from releases)
class Tglint < Formula
  desc "Formatter and linter for Terragrunt HCL files"
  homepage "https://github.com/roman-senchuk/tglint"
  version "1.0.0"
  license "MIT"

  if OS.mac?
    if Hardware::CPU.intel?
      url "https://github.com/roman-senchuk/tglint/releases/download/v1.0.0/tglint-darwin-amd64"
      sha256 ""
    else
      url "https://github.com/roman-senchuk/tglint/releases/download/v1.0.0/tglint-darwin-arm64"
      sha256 ""
    end
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/roman-senchuk/tglint/releases/download/v1.0.0/tglint-linux-amd64"
      sha256 ""
    else
      url "https://github.com/roman-senchuk/tglint/releases/download/v1.0.0/tglint-linux-arm64"
      sha256 ""
    end
  end

  def install
    binary_name = if OS.mac?
      Hardware::CPU.intel? ? "tglint-darwin-amd64" : "tglint-darwin-arm64"
    else
      Hardware::CPU.intel? ? "tglint-linux-amd64" : "tglint-linux-arm64"
    end
    
    bin.install binary_name => "tglint"
  end

  test do
    system "#{bin}/tglint", "--version"
  end
end
