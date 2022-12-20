# Homepage API Documentation

这里是 Homepage 的 API 的一些约定。

## /api/v1/login

```
PostLoginRequest = {
    username: string;
    password: string;
}

PostLoginResponse = {
    token: string;
}
```