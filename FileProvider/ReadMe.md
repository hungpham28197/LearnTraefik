# Traefik cấu hình bằng file

Hỏi: Làm  sao để Traefik có thể định tuyến, chuyển hướng, xử lý được request đến các service phía sau?

Đáp: Traefik cần truy vấn hoặc đọc vào API của các provider. Tiếng Anh ngắn gọn là "Configuration discovery in Traefik is achieved through Providers."

Hỏi: Có những loại provider nào?

Đáp: 
1. Label-based: dùng thẻ label để bổ xung thông tin cấu hình vào từng docker container
2. Key-Value-based: mỗi docker container dùng key-value để lưu cấu hình
3. Annotation-based: đánh dấu
4. File-based: dùng file cấu hình

Bài hôm nay chúng ta học cách dùng File provider hay đơn giản là dùng file để cấu hình.
Thông thường chúng ta dùng 2 file cấu hình:
- [traefik.yml](traefik.yml): file cấu hình lúc traefik khởi động lên.
- [traefik_dynamic.yml](traefik_dynamic.yml): file cấu hình định tuyến các request qua middleware rồi đến các service


## Chạy thử ví dụ này

```
docker-compose up -d
```

Mở trình duyệt mở http://localhost/whoami
nếu thấy dialog yêu cầu đăng nhập gõ:
user: root
pass: abc

Bạn sẽ thấy
![](images/localhost_whoami.jpg)

## Giải thích docker-compose.yml

[docker-compose.yml](docker-compose.yml)

```yaml
version: '3'

services:
  reverse-proxy:
    image: traefik:v2.5
    ports:      
      - "80:80"
      - "8080:8080"
    volumes:
# map traefik.yml ở thư mục hiện thời ở host vào thư mục / ở container
      - $PWD/traefik.yml:/traefik.yml
# map traefik_dynamic.yml ở thư mục hiện thời ở host vào thư mục / ở container
      - $PWD/traefik_dynamic.yml:/traefik_dynamic.yml

  whoami:
    image: traefik/whoami
```
## traefik.yml

```yml
entryPoints:
  web:
    address: :80

providers:
  file:
    filename: /traefik_dynamic.yml  # Trỏ tới file cấu hình ở thư mục gốc
api:
  insecure: true  # Bật dashboard lên
```

## traefik_dynamic.yml

```yml
# Cấu hình định tuyến
http:
  routers:
    # trỏ đến whoami servide
    to-whoami:
      rule: "Host(`localhost`) && PathPrefix(`/whoami`)"  # Bắt các request http://localhost/whoami
        # If the rule matches, applies the middleware
      middlewares:
        - test-user
      # If the rule matches, forward to the whoami service (declared below)
      service: whoami

  middlewares:
    # Define an authentication mechanism
    test-user:
      basicAuth:
        users:
          - root:$2y$10$Ja5KhDLxe2wqVPL4rOfx..Ep2Iq3NWH0FIYa6urKdlfIEtohSjS2a

  services:
    whoami:
      loadBalancer:
        servers:
          - url: http://whoami  # Mặc dù có 2 service load balance nhưng luôn chỉ chỏ được 1 cái đầu tiên
```