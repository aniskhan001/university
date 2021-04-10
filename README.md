# University [WIP]
A sample project in Go that tries to follow the [CLEAN architecture][clean-architecture] and the [12-factor methodology][12-factor-app] using [SOLID design principle][solid-design-principle]. This application offers to manage entities(departments, clubs, teachers, students, etc.) in a university by providing a REST API interface.

## Configuration
* This project uses `.editorconfig`
* This project uses a separate `dev.Dockerfile` for hot reload
* This project uses a Relational Database (MySQL) to store data & information
* This project uses `infrasctructure/config/config.local.yml` to load application configuration

## Development
```bash
make run
```

This will boot up application dependencies (DB, cache, etc.) in the background with `docker-compose`, and then run the application container in the foreground so we can watch the logs. The default `Dockerfile` is for production use, while `dev.Dockerfile` is being used for local development.

## Hot Reload
To pick up the changes in the running container, this project uses [CompileDaemon][compile-daemon]. It accepts the `go build` command as an argument.

```Dockerfile
ENTRYPOINT ["CompileDaemon", "--build=go build -mod=mod"]
```

Whenever changes are detected, it tries to run the build command again so that it can restart the server. To detect the changes, a volume is attached to the server container, which mounts host files to that volume.

```yml
volumes:
    - ./:/project
```


<!-- external links -->
[clean-architecture]: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
[12-factor-app]: https://en.wikipedia.org/wiki/Twelve-Factor_App_methodology
[solid-design-principle]: https://en.wikipedia.org/wiki/SOLID
[compile-daemon]: https://github.com/githubnemo/CompileDaemon
