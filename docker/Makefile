# docker image for development
build:
	docker build -t chalvern/golang .
	docker tag chalvern/golang chalvern/golang:1.12

# 运行镜像的命令，这里添加了 privileged 用以获取最大权限便于开发
run:
	docker run -it --privileged chalvern/golang:1.12 bash

# 把镜像推送到 dockerHub
push:
	docker push chalvern/golang
	docker push chalvern/golang:1.12