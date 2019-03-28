# 短网址

## 简介

本项目基于Golang编写，使用beego框架。使用redis做发号器，将整数转换成62进制的数，实现缩短网址的目的。支持加自定义salt，自定义最短uri长度。目前只完成了一个基础能用的版本，后续加入一些新特性。



## 算法介绍

### 自增序列进制转换算法

也叫永不重复法，本项目使用的算法。大致原理是使用一个十进制的自增ID，转换成62进制的数值，这也就不会出现重复的情况，这个利用的是低进制转高进制，字符数会减少的特性。但如果只是简单的把十进制转成62进制，会遇到以下问题：

1. 短码不固定，随着ID变大，短码也变大。如果要固定，可以设置从指定ID开始。
2. ID是有序的，可以通过62进制再转回十进制，或者通过通过自增ID转62进制之后批量爬库。这个可以通过给ID加salt解决，但如果salt也泄露了，也不安全。

本项目使用第三方包 [hashids](https://hashids.org/go/) ，可以给自增的ID加salt。



### Hash算法

1. 将长网址 `md5` 生成 32 位签名串,分为 4 段, 每段 8 个字节；
2. 对这四段循环处理, 取 8 个字节, 将他看成 16 进制串与 0x3fffffff(30位1) 与操作, 即超过 30 位的忽略处理；
3. 这 30 位分成 6 段, 每 5 位的数字作为字母表的索引取得特定字符, 依次进行获得 6 位字符串；
4. 总的 `md5` 串可以获得 4 个 6 位串,取里面的任意一个就可作为这个长 url 的短 url 地址。

存在碰撞（重复）的可能性，虽然几率很小。



## 部署

1. clone当前项目

2. 安装依赖
   ```shell
   go get -u github.com/astaxie/beego
   go get -u github.com/go-redis/redis
   go get -u github.com/speps/go-hashids
   go get -u github.com/jinzhu/gorm
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
- [ ] 用户自定义短码