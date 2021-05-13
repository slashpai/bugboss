# bugboss

[Bugzilla](https://www.bugzilla.org/) cli tool written in [go](https://golang.org/). It uses bugzilla
's [rest api](https://wiki.mozilla.org/Bugzilla:REST_API) for getting data. I am still working on features. Yet to provide installer which does automatic installation.

## Features

- [x] Search a bug with id

- [x] Get bugs assigned to a user by passing user `email id`

## Config file format

Create a config file at `home` directory named `.bugboss.yaml`. If using a different path or file name you need to pass it in config option.

```yaml
bugzilla:
  url: bugzilla.redhat.com
```

## Build executable

```go
go build
```

Move to a preferred path

```bash
mv bugboss /usr/local/bin
```

## Usage

```go
Buzilla cli to help to interact with bugzilla.
You can quickly search a bugzilla issue instead of waiting for web UI to load

Usage:
  bugboss [command]

Available Commands:
  help        Help about any command
  search      Search a bugzilla id
  userBugs    Search bugs assigned to a user

Flags:
      --bugzilla-url string   bugzilla Url
      --config string         config file (default is $HOME/.bugboss.yaml)
  -h, --help                  help for bugboss
  -w, --ui                    enable webui output

Use "bugboss [command] --help" for more information about a command.
```

## TODO

- [ ] Add better documentation in README
- [ ] Provide Installer
- [ ] Improve code
- [ ] Authentication mechanism
- [ ] More features and enhancements

## Contributing

- Fork the project on GitHub
- Clone the project
- Add changes
- Commit and push
- Create a pull request

## License

[Apache License](LICENSE)
