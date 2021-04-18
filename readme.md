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
ls go.mod || go mod init ginweb02
ls go.sum || go mod tidy
docker build -t ginweb02:v1.0.0 . 
docker rm -f ginweb02 || echo '不存在'
docker run -d --name ginweb02 --restart=always -p 8092:8092 ginweb02:v1.0.0 
```

