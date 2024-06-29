[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/percybolmer/grpcstreams)

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/percybolmer/grpcstreams)

### Reading

1. — [Using GRPC with TLS, Golang and React (No Envoy)](https://programmingpercy.tech/blog/using-grpc-tls-go-react-no-reverse-proxy/)
2. — [Embedding a React application in a Golang binary](https://programmingpercy.tech/blog/embedd-web-application-golang/)
3. — [gRPC Interceptors](https://programmingpercy.tech/blog/grpc-interceptors/)
4. — [Streaming data with gRPC](https://programmingpercy.tech/blog/streaming-data-with-grpc/)

### Updates

Percy Bolmér made [a lovely article about streaming gRPC](https://programmingpercy.tech/blog/streaming-data-with-grpc/)
to a React application without envoy. I found his repository, and dusted it off to get it (mostly) working.

```
cd webapp/hwmonitor
npm run build
cd ../..
go run main.go
```

One problem remains that when I run `go run main.go` in this root directory, the local react application does not load
at `http://localhost:8383`.

However, if in another terminal I do this:
```
cd webapp/hwmonitor
npm start
```
Then in Firefox I can get `localhost:3000` to see the desired results. In Chrome, it complains about CORS because it's
talking to port 8383 from port 3000, which you can relax with CORS extensions.

### Stuff I did to make it up to date

```
export GOBIN=~/go/bin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

brew install protobuf
brew install protoc-gen-js
brew install protoc-gen-grpc-web
cd proto
protoc service.proto --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. --go-grpc_out=. --go_out=.
cp *.js ../webapp/hwmonitor/src/proto
cd ..
```

I also had to update all the Go libraries. gRPC gets real argy-bargy if the versions don't match the protoc plugins.

