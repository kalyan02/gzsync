# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Gzsync < Formula
  desc ""
  homepage ""
  version "0.1.1"
  bottle :unneeded

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/kalyan02/gzsync/releases/download/v0.1.1/gzsync_0.1.1_Darwin_x86_64.tar.gz"
      sha256 "8d8e82c11dcd4e3023fe745809ff4ed9d78dadf909c959d7e7cae45dbf7cc0c2"
    end
    if Hardware::CPU.arm?
      url "https://github.com/kalyan02/gzsync/releases/download/v0.1.1/gzsync_0.1.1_Darwin_arm64.tar.gz"
      sha256 "1bcc8e57c80fc28f9292633d0a2de90a6caaf7dd23e94ab8f6397047a2fb41db"
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/kalyan02/gzsync/releases/download/v0.1.1/gzsync_0.1.1_Linux_x86_64.tar.gz"
      sha256 "f0d9600e58801f26d07a6c668c9de754a6237784bfa6c55d2d4d6f52ff2ae71e"
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/kalyan02/gzsync/releases/download/v0.1.1/gzsync_0.1.1_Linux_arm64.tar.gz"
      sha256 "156592e6e92a9543232d72051f3c8a5aae92677cd563739b06e072e7c95bbd5c"
    end
  end

  def install
    bin.install "gzsync"
  end
end
