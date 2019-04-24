# 短网址

## 简介

本项目基于Golang编写，使用beego框架。使用redis做发号器，将整数转换成62进制的数，实现缩短网址的目的。支持加自定义salt，自定义最短uri长度。目前除了基础功能，完成了权限和日志，后续继续加入一些新特性。



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
   
       access_log /wwwroot/publish/shorturl/logs/access.log;
       error_log /wwwroot/publish/shorturl/logs/error.log;
   
       # ssl ------
       listen       443 ssl;
       ssl_certificate      /etc/letsencrypt/live/example.com/fullchain.pem;
       ssl_certificate_key  /etc/letsencrypt/live/example.com/privkey.pem;

       # 关闭 [/favicon.ico] 和 [/robots.txt] 的访问日志。
       # 并且即使它们不存在，也不写入错误日志。
       location = /favicon.ico { access_log off; log_not_found off; }
       location = /robots.txt  { access_log off; log_not_found off; }
   
       location / {
           proxy_set_header X-Forwarded-For $remote_addr;
           proxy_set_header Host            $http_host;
   
           proxy_pass http://127.0.0.1:8080;
       }
   }
   ```

## 接口使用
### 短网址生成接口

**请求地址：** /api/v1/create

**请求方式：** POST

**Content-Type：** application/json; charset=UTF-8

**请求头Headers：**

| 字段  | 类型   | 是否必须 | 说明                       |
| ----- | ------ | -------- | -------------------------- |
| Token | string | 是       | 由数字和字母组成的32位字符 |

**请求体Body：**

| 字段 | 类型   | 是否必须 | 说明    | 示例                  |
| ---- | ------ | -------- | ------- | --------------------- |
| url  | string | 是       | URL地址 | "https://github.com/" |

**代码示例：**

curl

```shell
curl -H "Content-Type:application/json; charset=UTF-8" -H "Token: 你的token" -X POST "http://localhost:8080/api/v1/create" -d '{"url":"你的长网址"}'
```

PHP

```php
<?php
    $url = 'http://localhost:8080/api/v1/create';
    $method = 'POST';
    
    // TODO: 设置Token
    $token = '你的Token';
    
    // TODO：设置待注册长网址
    $bodys = array('url'=>'你的长网址'); 
    
    // 配置headers 
    $headers = array('Content-Type:application/json; charset=UTF-8', 'Token:'.$token);
    
    // 创建连接
    $curl = curl_init($url);
    curl_setopt($curl, CURLOPT_CUSTOMREQUEST, $method);
    curl_setopt($curl, CURLOPT_HTTPHEADER, $headers);
    curl_setopt($curl, CURLOPT_FAILONERROR, false);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($curl, CURLOPT_HEADER, false);
    curl_setopt($curl, CURLOPT_POST, true);
    curl_setopt($curl, CURLOPT_POSTFIELDS, json_encode($bodys));
    
    // 发送请求
    $response = curl_exec($curl);
    curl_close($curl);
    
    // 读取响应
    var_dump($response);
    ?>

```

**响应结果示例：**

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "longUrl": "https://github.com/",
        "shortUrl": "http://127.0.0.1:8080/789Yp146"
    }
}
```

### 短网址还原接口

**请求地址：** /api/v1/query

**请求方式：** POST

**Content-Type：** application/json; charset=UTF-8

**请求头Headers**

| 字段  | 类型   | 是否必须 | 说明                       |
| ----- | ------ | -------- | -------------------------- |
| Token | string | 是       | 由数字和字母组成的32位字符 |

**请求体Body：**

| 字段     | 类型   | 是否必须 | 说明   | 示例                             |
| -------- | ------ | -------- | ------ | -------------------------------- |
| shortUrl | string | 是       | 短网址 | "http://127.0.0.1:8080/RQ9P3L4y" |

**代码示例：**

curl

```shell
curl -H "Content-Type:application/json; charset=UTF-8" -H "Token: 你的token" -X POST "http://localhost:8080/api/v1/query" -d '{"shortUrl":"你的短网址"}'
```

PHP

```php
<?php
    $url = 'http://localhost:8080/api/v1/query';
    $method = 'POST';
    
    // TODO: 设置Token
    $token = '你的Token';
    
    // TODO: 设置还原的短网址
    $bodys = array('shortUrl'=>'你的短网址'); 
    
    // 配置headers 
    $headers = array('Content-Type:application/json; charset=UTF-8', 'Token:'.$token);
    
    // 创建连接
    $curl = curl_init($url);
    curl_setopt($curl, CURLOPT_CUSTOMREQUEST, $method);
    curl_setopt($curl, CURLOPT_HTTPHEADER, $headers);
    curl_setopt($curl, CURLOPT_FAILONERROR, false);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($curl, CURLOPT_HEADER, false);
    curl_setopt($curl, CURLOPT_POST, true);
    curl_setopt($curl, CURLOPT_POSTFIELDS, json_encode($bodys));
    
    // 发送请求
    $response = curl_exec($curl);
    curl_close($curl);
    
    // 读取响应
    var_dump($response);
    ?>

```

**响应结果示例：**

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "longUrl": "https://github.com/",
        "shortUrl": "http://127.0.0.1:8080/789Yp146"
    }
}
```


## 注意
MySQL需要开启大小写敏感


## 未完成功能

- [x] 账户授权模式
- [x] 访问日志
- [ ] 设置短网址有效期
- [ ] 统计功能
- [ ] 短网址加密访问
- [ ] 用户自定义短码
- [ ] 黑名单
- [ ] 内容鉴定