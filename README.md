# bugboss

[Bugzilla](https://www.bugzilla.org/) cli tool written in [go](https://golang.org/). It uses bugzilla
's [rest api](https://wiki.mozilla.org/Bugzilla:REST_API) for getting data. I am still working on features. Yet to provide installer which does automatic installation.

## Features

- [x] Search a bug with id
- [x] Get bugs assigned to a user by passing user `email id`
- [x] Option to view output in webui

## Config file format

Create a config file at `home` directory named `.bugboss.yaml`. If using a different path or file name you need to pass it in config option.

**Example**

```yaml
bugzilla.url: bugzilla.redhat.com
```


## Build executable

```bash
gi clone https://github.com/slashpai/bugboss.git
cd bugboss
```

```go
go build
```

Move to a preferred path

```bash
mv bugboss /usr/local/bin
```

## Usage

```go
Bugzilla cli to help to interact with bugzilla.
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

## Search

```go
Search a given bugzila id

Usage:
  bugboss search [flags]

Flags:
  -h, --help        help for search
  -n, --id string   Bug ID

Global Flags:
      --bugzilla-url string   bugzilla Url
      --config string         config file (default is $HOME/.bugboss.yaml)
  -w, --ui                    enable webui output
```

**Example**

In this example since config values not passed its taken from config file which is by default `$HOME/.bugboss.yaml`. Look at [config file format section](#config-file-format)

```go
bugboss search --id 1955051                                                                               
```

**To get a web based output**

```go
bugboss search --id 1955051 -w
```

**To override configs in `$HOME/.bugboss.yaml`, pass those flags while executing**

```go
bugboss search --bugzilla-url bugzilla.redhat.com --id 1955051
```

## userBugs

```go
Give user email as input

Usage:
  bugboss userBugs [flags]

Flags:
  -e, --email_id string   Email ID
  -h, --help              help for userBugs

Global Flags:
      --bugzilla-url string   bugzilla Url
      --config string         config file (default is $HOME/.bugboss.yaml)
```

**Example**

**Note: Replace with correct email id**

```go
bugboss userBugs --email_id testuser1@example.com
```

**To get a web based output**

```go
bugboss userBugs --email_id testuser1@example.com -w
```

## TODO

- [ ] Add better documentation in README
- [ ] Provide Installer
- [ ] Improve code
- [ ] Authentication mechanism
- [ ] More features and enhancements
- [ ] Better navigation in UI and more fields

## Contributing

- Fork the project on GitHub
- Clone the project
- Add changes
- Commit and push
- Create a pull request

## License

[Apache License](LICENSE)
