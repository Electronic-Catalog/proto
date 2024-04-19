pull-image:
	sudo docker pull namely/protoc-all

build-auth:
	# docker run --name auth-builder -v `pwd`/protobuf/:/defs harbor.tasn.ir/library/namely/protoc-all:latest -f ./auth/auth.protobuf -l go -o ./
	docker run --name auth-builder -v `pwd`/protobuf/:/defs namely/protoc-all -f ./auth/auth.proto -l go -o .
	docker rm auth-builder

build-cmng:
	docker run --name cmng-builder -v `pwd`/protobuf/:/defs namely/protoc-all -f ./cmng/cmng.proto -l go -o .
	docker rm cmng-builder
