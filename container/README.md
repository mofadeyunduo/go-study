# 容器化

## Docker

```
docker build -t image-name . && docker run -p 8080:8080 -v /Users/piers/config:/data/config --name container-name image-name
```

## 注意点

### 和 MacOS 相关
- VOLUME 必须挂在指定目录
- 同一个 Docker 下，启动 Redis、MySQL 等实例，访问的 host = host.docker.internal 
- Server 在 Docker 中绑定的地址必须是 0.0.0.0