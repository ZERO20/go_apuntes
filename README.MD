# Go

- Habilitar trackeo de dependencia
```
go mod init example/hello
```

- Compilar el paquete
```bash
go build file.go # Compilar

go build -o main main.go # Crear el paquete compilado
```

- Para evitar el tener que realizar múltiples pasos, podemos hacer uso de run que nos permite realizar ambas operaciones y sin dejar el archivo compilado dentro del directorio.
```bash
    go run . # por default para funciones main
```

- Run test
```bash
    go test
```

- Run test with coverage
```bash
    go test -coverage
```

- Run test with coverage html report
```bash
    go test -coverprofile=coverage.out
    go tool cover -html=coverage.out
```

- Install/update package
Use the following command to download the repository to the local file system.
```bash
go get github.com/haccer/subjack # default branch ```master```
```

- Install the executable by using the following command.
```bash
go install -v github.com/haccer/subjack@latest
```

- Update all dependencies
```bash
go get -u
```

Update only package
```bash
go get -u github.com/author/package
```

Update all packages
```bash
go get -u all
```