proto: protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/Message.proto

## make proto does not work, need to run the command manually