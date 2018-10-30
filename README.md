# repl —- Add more read-eval-print-love in your life

`repl` is an interactive program which wraps non-interactive programs. Inspired by [Chris Wanstrath](https://github.com/defunkt)'s original [Ruby version](https://github.com/defunkt/repl). It comes with line editing capabilities (including history) provided by [liner](https://github.com/peterh/liner).

## Installation

### Manual

```sh
$ go install github.com/citizen428/repl
```

## Examples

Ruby's gem command:

```
$ repl gem
gem>> --version
2.7.6

gem>> list --local
ast (2.4.0)
awesome_print (1.8.0)
bigdecimal (default: 1.3.4)
bundler (1.17.1, 1.16.6)
byebug (10.0.2)
cmath (default: 1.0.0)
[...]
```

Docker:

```
$ repl docker
docker>> version
Client:
Version:           18.06.1-ce
API version:       1.38
Go version:        go1.10.3
Git commit:        e68fc7a
Built:             Tue Aug 21 17:21:31 2018
OS/Arch:           darwin/amd64
Experimental:      false
[...]

docker>> images
REPOSITORY            TAG                 IMAGE ID            CREATED             SIZE
citizen428/unsavory   latest              abe19fd175b0        6 days ago          6.6MB
citizen428/unsavory   v0.2.0              abe19fd175b0        6 days ago          6.6MB
golang                alpine              95ec94706ff6        3 weeks ago         310MB
```

## Usage

```
$ repl
Usage:
  repl cmd [options]

Options:
  -debug
    	Enable debug output
  -histdir string
    	Directory for history file (default "/Users/<username>")
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/citizen428/repl.

## TODO

- [ ] Support completions
- [ ] Add tests
- [ ] Release

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
