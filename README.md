<h1 align="center">🔍 ygrep</h1>

<p align="center">
  Search for keys in all yaml files of a directory(and its subdirectories).
  <br><br>
  <a href="https://cloud.drone.io/dj95/ygrep">
    <img alt="BuildStatus" src="https://cloud.drone.io/api/badges/dj95/ygrep/status.svg" />
  </a>
  <a href="https://github.com/dj95/ygrep/actions?query=workflow%3AGo">
    <img alt="GoActions" src="https://github.com/dj95/ygrep/workflows/Go/badge.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/dj95/ygrep">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/dj95/ygrep" />
  </a>
  <a href="https://github.com/dj95/ygrep/releases">
    <img alt="latest version" src="https://img.shields.io/github/tag/dj95/ygrep.svg" />
  </a>
  <a href="https://codecov.io/gh/dj95/ygrep">
    <img src="https://codecov.io/gh/dj95/ygrep/branch/master/graph/badge.svg" />
  </a>
  <a href="https://sonarcloud.io/dashboard?id=dj95:ygrep">
    <img src="https://sonarcloud.io/api/project_badges/measure?project=dj95:ygrep&metric=sqale_rating" />
  </a>
  <a href="https://sonarcloud.io/dashboard?id=dj95:ygrep">
    <img src="https://sonarcloud.io/api/project_badges/measure?project=dj95:ygrep&metric=alert_status" />
  </a>
</p>


## 📦 Requirements

- Golang(>=1.11) *(for building ygrep)*
- 🚧 Make *(dev dependency)*
- 🚧 staticcheck *(dev dependency)*
- 🚧 golint *(dev dependency)*


## 🔧 Installation

- Download the binary from the release page for your platform or run `go build -o ygrep cmd/ygrep/main.go`
- Copy the binary to a location in your `$PATH` (e.g. for linux `sudo cp ygrep /usr/local/bin/.`)


## 🚀 Usage

ygrep will help you to find key-value pairs in yaml files. For example if you'd like to
search for the all keys containing 'foo' and their related values in the `./test` directory,
you can use the commandline `ygrep -p ./test -e '.*foo.*'`. For searching recursively through
all subdirectories of `./test`, use `ygrep -rp ./test -e '.*foo.*'`.

Feel free to test those commands with the provided yml files in the `./test` directory provided
in this repository. The first example should look like the following listing:

```
$ ./bin/ygrep -p ./test -e 'foo'
:: ./test/foo.yml
foo: bar

```

Please refer to the help page (use `-h` as flag) for more options.


## ✅ Testing

Run `make tests` in order to run the tests for this application.
After running tests you will find a coverage report in the `./report` directory.


## 🤝 Contributing

If you are missing features or find some annoying bugs please feel free to submit an issue or a bugfix within a pull request :)


## 📝 License

© 2020 Daniel Jankowski


This project is licensed under the MIT license.


Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:


The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.


THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
