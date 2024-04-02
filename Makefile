pull-image:
	sudo docker pull namely/protoc-all

build-auth:
	sudo docker run --name auth-builder -v `pwd`/auth/:/defs namely/protoc-all -f `pwd`/auth/auth.proto -l go -o ./

clea-up:
	sudo docker rm auth-builder
