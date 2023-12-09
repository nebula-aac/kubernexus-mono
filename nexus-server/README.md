docker build . -t nexus-server

docker run -p 8080:8080 nexus-server