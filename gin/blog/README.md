# blog

# swagger
<!-- 
    internal/routers/api/v1 为API接口写入注解
    $ swag init 命令生成swagger文档
    启动服务,访问 http://localhost:8999/swagger/index.html 
-->


# jwt
curl -X POST http://localhost:8999/auth -d 'app_key=biny&app_secret=go-blog'
curl -X GET http://localhost:8999/api/v1/tags?state=6&name=biny -H 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiM2Q1ODZhNDI1ZjU0MmNmY2U1YjlhNzhlNDVmODExZDYiLCJhcHBfc2VjcmV0IjoiZDJjMjI5M2ZhMThjNTc0ZmRkN2E0ZWE4ODNlMTY5YzgiLCJleHAiOjE2NTIzNzI4MzcsImlzcyI6ImJsb2cifQ.mxd5a81oOvSky0DAjBKreMGdChv-x9rj0srUlAqaTIE'

# upload
curl -X POST http://localhost:8999/upload/file -F "file=@/golang/src/github.com/practic-go/README.md" -F type=1
curl -X POST http://localhost:8999/upload/file -F "file=@/golang/src/github.com/practic-go/gin/blog/storage/uploads/biny.jpg" -F type=1