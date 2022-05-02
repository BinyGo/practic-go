# 1.构建镜像
Dockerfile所在的目录下执行docker build构建镜像
docker build -t go-app-img .

# 2.验证镜像(可跳过)
docker run -d -p 3333:3000 --rm --name go-app-container go-app-img

# 3.编写deployment.yaml

# 4.部署应用
kubectl create -f deployment.yaml

# 5.暴露应用 NodePort外部端口 port=xx内部互访端口 target-port容器端口
kubectl expose deployment my-go-app --name=go-app-svc --type=NodePort --target-port=3000
kubectl get svc
    go-app-svc   NodePort    10.108.193.228   <none>        3000:30568/TCP   3s

# 6. 访问 http://localhost:30568 http://localhost:30568/health_check

# 参考资料 https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247485328&idx=1&sn=1615cbb8d60b0df38cf59375cca6cef2&chksm=fa80d607cdf75f11cf025730a492646f3ea375a81fca5e70c46703f37aa2e7d929950693cae6&scene=178&cur_album_id=1394839706508148737#rd