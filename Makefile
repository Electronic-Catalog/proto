pull-image:
	sudo docker pull namely/protoc-all

build-auth:
	# docker run --name auth-builder -v `pwd`/protobuf/:/defs harbor.tasn.ir/library/namely/protoc-all:latest -f ./auth/auth.protobuf -l go -o ./
	docker run --name auth-builder -v `pwd`/protobuf/:/defs namely/protoc-all -f ./auth/auth.proto -l go -o .

clea-up:
	docker rm auth-builder
