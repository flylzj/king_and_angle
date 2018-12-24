# 天使与国王

## 登录

`POST /api/user/token`

need

```json
{
  "username": "5602216028",
  "password": "5602216028"
}
```

成功后返回

```json

{
    "code": 0,
    "message": "success",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo3OH0.C7uzc0laz9xlelzIcbjTr1wUZSAyoRVeUKtQUOj9my4",
    "wish": ""  //没有则为空字符串
}
```

## 国王与天使

### 查看我的愿望

`GET /api/king-and-angle/wish`

成功后返回

```json
{
"code": 0,
"message": "success",
"wish": "",
"wish_status": 0  // status=0未完成   status=1 完成
}
```

### 更新我的愿望状态

`POST /api/king-and-angle/wish_status`

need

```json
{
  "wish_status": 0 // stauts 同上
}
```

成功后返回

```json
{
"code": "0",
"message": "success"
}
```

### 添加愿望

`POST /api/king-and-angle/wish`

need

```json
{
  "wish": ""
}
```

成功后返回

```json
{
  "message": "success",
  "code": 0
}
```

###  查看国王

`GET /api/king-and-angle/king`

成功后返回

```json
{
    "code": 0,
    "data": {
        "king": "xxx",
        "king_username": "xxx",
        "king_wish": "",
        "king_wish_status": 0,
    },
    "message": "success"
}
```

### 查看天使

`GET /api/king-and-angle/angle`

成功后返回

```json
{
    "code": 0,
    "data": {
        "my_wish": "xxx"
    },
    "message": "success"
}
```

### 发送祝福给国王

`POST /api/blessing/king`

need 

```json
{
	"blessing": "123"
}
```

成功后返回

```json
{
    "code": 0,
    "message": "success"
}
```

### 查看天使给我的祝福

`GET /api/blessing/angle`

成功后返回

```json
{
    "code": 0,
    "data": {
        "blessing": ""
    },
    "message": "success"
}
```

### 查看我给国王的祝福


`GET /api/blessing/king`

成功后返回

```json
{
    "code": 0,
    "data": {
        "blessing": ""
    },
    "message": "success"
}
```

## 聊天系统

### 连接websocket

`ws://host/ws`

### 权限验证

send 

```json
{
    "token": ""
}
```

### 发送消息

send

```json
{
    "from": "username",
    "to": "username",
    "message": "message"
}
```



