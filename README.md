[![Go Reference](https://pkg.go.dev/badge/github.com/moohbr/WebMonitor.svg)](https://pkg.go.dev/github.com/moohbr/WebMonitor)
# WebMonitor

Faced with a considerable number of tools that monitor your site, why use WebMonitor? Why use a lightweight, easy and cost-free tool? I don't know either.

WebMonitor emerged from a need to monitor my infrastructure in a simple and effective way.
It's a tool made by a university student 'n intern, so it's constantly improving.

## Features

- [x] Monitor your site
- [x] Monitor your site with a custom user-agent
- [x] Notify you when your site is down

## Future features

- [ ] Monitor your site with a custom user-agent
- [ ] Notify with custom message/custom interval/custom timeout
- [ ] Create a web interface to manage the application
- [ ] User profiles and authentication


## Contributing

If you want to contribute to this project, you can do it in two ways:

1. Open an issue with a bug report or a feature request
2. Open a pull request with a bug fix or a new feature

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details


## Documentation
### Environment Variables

To run this project, you will need to add the following environment variables to your `.env` file:

- `SMPT_SERVER` - SMTP server address
- `SMTP_PORT` - SMTP server port
- `SMTP_USER` - SMTP server user
- `SMTP_PASSWORD` - SMTP server password
- `SMTP_INSECURE` - SMTP server insecure

### CLI Arguments

```
 Available Commands:
  install     Install database for first time
  show        Show a list of servers or users
  add         Add a new site to monitor or user to notify
  remove      Remove a site or user
  update      Update a site or user
  help        Help about any command

Use "WebMonitor [command] --help" for more information about a command.  

Flags:
  -h, --help      help for WebMonitor
  -v, --verbose   verbose output
```

### Installation


#### Manual

1. Clone the repository
2. Install the dependencies with `go build`
3. Run the script with `WebMonitor`

### Usage


#### Manual

1. Create a `.env` file.
2. Install the dependencies with `go build` and run the script with `webmonitor`

## Questions

If you have any questions, feel free to open an issue or contact me on [email](mailto:moohbr@gmail.com).

## Acknowledgements

 - [Cobra](github.com/spf13/cobra)
 - [SQLite](github.com/mattn/go-sqlite3)
 - [go-mail](github.com/go-mail/mail)
 - [go-dotenv](github.com/joho/godotenv)


## Authors

- [@moohbr](https://www.github.com/moohbr)
