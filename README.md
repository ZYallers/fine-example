# fine-example
A Example project using Fine framework.

## How to make it run?
Run `make help` command, You will see help information.
```
Choose a command run in project fine-example:
Usage: make <TARGETS> <OPTIONS> ...
Targets:
  go.fmt          Gofmt (reformat) package sources.
  go.deps         Add missing modules and delete unused modules to generate a directory of dependencies.
  go.test         Test packages.
  go.build        Compile the project and generate an executable file.
  go.run          Run packages.
  server.start    Start server.
  server.stop     Stop server.
  server.reload   Reload server.
  server.status   Show server status.
  clean           Clean package.
  swag.fmt        Format Swag interface comments.
  swag.gen        Generate Swagger.json interface document.
  help            Show help message.
Options:
  GOOS                GO compilation system, defaults to the current system.
  GOARCH              GO compilation architecture, defaults to the current architecture.
  GO_BUILD_FLAGS      GO compilation additional parameters, default is -a.
  GO_BUILD_LDFLAGS    GO compilation link parameter, default is -w.
```
Follow the prompts to execute the corresponding commands.

If you want to directly debug and run the project locally, 
execute `make go.run` command.

## Cli tool

```bash
go install github.com/ZYallers/fine/hxs/fine/cmd/fine@latest
```

# Limitation

```
golang version >= 1.17
```

# Documentation
- Gin Web Framework: [https://gin-gonic.com](https://gin-gonic.com)
- Gorm ORM: [https://gorm.io](https://gorm.io/)
- Go Redis: [https://redis.uptrace.dev](https://redis.uptrace.dev)
- GoDoc API: [https://pkg.go.dev/github.com/ZYallers/fine](https://pkg.go.dev/github.com/ZYallers/fine)

# License

`Fine-example` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.