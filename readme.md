## 1.初始化
```
go mod init ginweb02
```

## 2.依赖 golang
```
docker pull golang
```

## 3.编译docker image

```
docker build -t ginweb02:v1.0.0 .
```

## 4.运行
```
docker run -d -p 8092:8092 ginweb02:v1.0.0
```


## jenkins
```
docker rm -f ginweb02

docker build -t ginweb02:v1.0.0 . 
docker run -d --name ginweb02 -p 8092:8092 ginweb02:v1.0.0 
```

