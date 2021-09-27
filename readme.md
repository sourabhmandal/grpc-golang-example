## Generate chat.pb.go

protoc --proto_path=./ --go_out=./ --go_opt=paths=source_relative chat.proto ./chat.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative chat.proto
