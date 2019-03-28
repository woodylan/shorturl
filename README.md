# 短网址

## 简介

本项目基于Golang编写，使用beego框架。使用redis做发号器，将整数转换成62进制的数，实现缩短网址的目的。支持加自定义salt，自定义最短uri长度。目前只完成了一个基础能用的版本，后续加入一些新特性。



## 部署

1. clone当前项目
2. 安装依赖

```shell
go get -u github.com/go-redis/redis
go get -u github.com/speps/go-hashids
go get -u github.com/jinzhu/gorm
go get -u github.com/astaxie/beego
```

3. 修改配置
```shell
cp conf/app.conf.example conf/app.conf
vim conf/app.conf
```


4. 运行迁移文件
```shell
bee migrate -conn="user:password@tcp(127.0.0.1:3306)/shorturl"
```

5. 运行项目
```shell
bee run
```

6. 打包发布

   Linux包

   ```shell
   bee pack -be GOOS=linux
   ```
   会得到一个可执行压缩包，解压到服务器运行，打开8080端口就可以运行了
   ```shell
   tar -zxvf shorturl.tar.gz
   nohup ./shorturl &
   ```

7. 配置nginx

```nginx
server {
    listen  80;
    server_name www.example.com;
    root /wwwroot/publish/shorturl/public;

    access_log /wwwroot/publish/shorturl/logs/access.log;
    error_log /wwwroot/publish/shorturl/logs/error.log;
    index index.php index.html index.htm;

    # ssl ------
    listen       443 ssl;
    ssl_certificate      /etc/letsencrypt/live/example.com/fullchain.pem;
    ssl_certificate_key  /etc/letsencrypt/live/example.com/privkey.pem;

    location / {
        try_files /_not_exists_ @backend;
    }

    location @backend {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;

        proxy_pass http://127.0.0.1:8080;
    }
}
```



## 使用





## 未完成功能

- [ ] 账户授权模式
- [ ] 设置短网址有效期
- [ ] 统计功能
- [ ] 短网址加密访问