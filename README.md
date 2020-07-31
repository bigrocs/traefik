# traefik demo

## 快速使用
### 下载
```
    git clone https://github.com/bigrocs/traefik.git
    cd traefik
```
### 构建user-api
```
    cd user-api
    make build
```

### 启动 traefik
```
    cd traefik
    make run
```
### 访问user-api
```
    http://127.0.0.1:1080/health
```