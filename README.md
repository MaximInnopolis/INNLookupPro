# INN Lookup Pro

INN Lookup Pro is a simple application that performs INN (Individual Taxpayer Identification Number) lookup using gRPC communication. The application consists of a server and a client component, allowing users to perform INN lookup operations.

## Project Structure

```plaintext
INNOLookupPro
├── cmd
│   ├── client
│   │   └── client.go
│   └── server
│       └── server.go
├── logger
│   └── logger.go
├── protos
│   ├── rusprofile_lookup.proto
│   ├── rusprofile_lookup.pb.go
│   └── rusprofile_lookup_grpc.pb.go
├── go.mod
├── Makefile
└── Dockerfile
```

## Building the Project

To build the project, use the provided Makefile. The Makefile includes targets for generating protobuf files, creating a Docker image, and more.

### Build protobuf files

```bash
make proto
```

### Generate Swagger documentation

```bash
make swagger
```

### Build Docker image

```bash
make docker
```

## Running the Project

Once the Docker image is built, you can run the INN Lookup Pro application as a Docker container.

```bash
docker run -p 8080:8080 inn-lookup-pro
```

This will start the gRPC server on port 8080 inside the Docker container.

To interact with the client, you can use the following commands:

- Run client interactively:

```bash
docker exec -it <container_id_or_name> ./client
```

- Run client as a separate command:

```bash
docker exec <container_id_or_name> ./client
```

Replace `<container_id_or_name>` with the actual ID or name of your running container.

## API Documentation

The Swagger documentation for the INN Lookup Pro API is generated during the build process. You can access the Swagger UI by navigating to `http://localhost:{port for swagger}/` in your browser.

## Notes

- Make sure you have Docker installed on your machine before building and running the project.
- The provided Dockerfile creates a multi-stage build for a smaller production image.

Feel free to explore and modify the project according to your requirements.