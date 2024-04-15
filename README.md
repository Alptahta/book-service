# book-service

## About
A simple RESTful API for CRUD functions on book domain.

This project uses:
- Go 1.22.2 with new enhanced routing patterns
- Docker for containerization

## Manual
Build Image From Dockerfile
```
docker build --no-cache --progress=plain -t book-service:multistage -f Dockerfile.multistage .
```

Run Container
```
docker run --detach --publish 8080:8080 --name book-service book-service:multistage
```
