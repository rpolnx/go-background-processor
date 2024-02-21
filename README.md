Proof of concept of golang background processor

This repository means to test the project [Go work](https://github.com/gocraft/work)

![Diagram for jobs](./files/diagram_background_processor.jpg)

## Running

Run dependencies
```sh
docker compose up -d
```

- Needs to configure .env first similiar to .env.example
```
    go run main.go
```

### Default host & port to access the app

- http://localhost:8080 # web server
- http://localhost:5040 # work UI
