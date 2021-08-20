# Fiber JWT
Đây là ví dụ được cải tiến từ [gofiber/jwt](https://github.com/gofiber/jwt)

gofiber/jwt sử dụng package [golang-jwt](https://github.com/golang-jwt/jwt)

## Hướng dẫn chạy thử
Code mẫu chỉ có duy nhật một user là 
```json
{
  "user": "john",
  "pass": "doe"
}
```

Login using username and password to retrieve a token.
```
curl --data "user=john&pass=doe" http://localhost:3000/login
```

Response
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjE5NTcxMzZ9.RB3arc4-OyzASAaUhC2W3ReWaXAt_z2Fd3BN4aWTgEY"
}
```

![](images/login.jpg)
Request a restricted resource using the token in Authorization request header.

```
curl localhost:3000/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjE5NTcxMzZ9.RB3arc4-OyzASAaUhC2W3ReWaXAt_z2Fd3BN4aWTgEY"
```
Response
```
Welcome John Doe
```
![](images/restricted.jpg)
## Cấu hình JWT middleware

```go
app.Use(jwtware.New(jwtware.Config{
		SigningKey:     SecretKey,  //secret key string
		SuccessHandler: handle_when_valid_token, //Hàm xử lý khi JWT hợp lệ
		ErrorHandler:   handle_when_invalid_token, //Hàm xử lý khi JWT không hợp lệ
	}))
```

## Hàm xử lý xự kiện khi phân tích JWT [jwt_handler.go](jwt_handler.go)

```go
func handle_when_valid_token(c *fiber.Ctx) error {
	fmt.Println(c.Get("Authorization")) //Trích xuất ra đoạn header
	c.Next()
	return nil
}

func handle_when_invalid_token(c *fiber.Ctx, e error) error {
	fmt.Println("Error when parsing JWT", e)
	return e
}
```