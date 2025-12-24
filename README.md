<div align="center">

# Sintesi

<blockquote align="center">Simple system info fetch utility.</blockquote>

<p>
  <a href="https://github.com/kwame-Owusu/sintesi/blob/main/LICENSE">
    <img alt="GitHub license" src="https://img.shields.io/github/license/sintesi/sintesi?style=for-the-badge">
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

# Install

Binaries are provided at the releases page [here](https://github.com/kwame-Owusu/sintesi/releases/). Or, you can just run go install `github.com/kwame-Owusu/sintesi@latest`

# Manual Compilation

```bash
Manual Compile

git clone https://github.com/kwame-Owusu/sintesi
cd sintesi
go get -d ./...
go build -ldflags "-w -s" # ldflags make the binary smaller

```
