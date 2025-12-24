<div align="center">

# Sintesi

<blockquote align="center">Simple system info fetch utility.</blockquote>

<p>
<a href="https://github.com/kwame-Owusu/sintesi/blob/main/LICENSE">
  <img alt="GitHub license" src="https://img.shields.io/github/license/kwame-Owusu/sintesi?style=for-the-badge">
</a>
  <br>
  Sintesi is a small and fast tool for getting stats about your system. Written in Go, compiles to Linux and macOS
</p>

<pre>
      ／＞　 フ
     | 　^　^|
    ／` ミ＿xノ
   /　　　　 |
  /　 ヽ　　 ﾉ
 │　　|　|　|
 ／￣  |　|　|
 (￣ヽ＿_ヽ_)__)
 ＼二)
</pre>

</div>

# Why

Wanted to make a utility tool that is fast and simple to showcase important information about
current computer and the cat looks cute as well ;).
# Demo

![sintesi-gif](https://github.com/user-attachments/assets/2079ec9c-d3bc-466a-8c35-5f47c3be9cec)



# Install

Binaries are provided at the releases page [here](https://github.com/kwame-Owusu/sintesi/releases/). Or, you can just run go install `github.com/kwame-Owusu/sintesi@latest`.

Aternatively it can also be installed through my personal homebrew tap:

```
brew tap kwame-Owusu/taps https://github.com/kwame-Owusu/homebrew-taps
brew install sintesi
sintesi --help
```

# Manual Compilation

```bash
git clone https://github.com/kwame-Owusu/sintesi
cd sintesi
go get -d ./...
go build -ldflags "-w -s" # ldflags make the binary smaller

```

# Usage

Run `sintesi`, and system info will be printed to terminal with ansi colors.

# License

Sintesi is licensed under the MIT license. [Read Here](https://github.com/kwame-Owusu/sintesi/blob/main/LICENSE) for more info.
