# API性能测试 

## 安装 wrk
```
$ git clone https://github.com/wg/wrk
$ cd wrk
$ make
$ sudo cp ./wrk /usr/bin

$ wrk --help
Usage: wrk <options> <url>
  Options:
    -c, --connections <N>  Connections to keep open
    -d, --duration    <T>  Duration of test
    -t, --threads     <N>  Number of threads to use

    -s, --script      <S>  Load Lua script file
    -H, --header      <H>  Add header to request
        --latency          Print latency statistics
        --timeout     <T>  Socket/request timeout
    -v, --version          Print version details

  Numeric arguments may include a SI unit (1k, 1M, 1G)
  Time arguments may include a time unit (2s, 2m, 2h)
```
## wrk常用的参数有下面这些：
```
-t，线程数（线程数不要太多，是核数的 2 到 4 倍就行，多了反而会因为线程切换过多造成效率降低）。
-c，并发数。
-d，测试的持续时间，默认为 10s。
-T，请求超时时间。
-H，指定请求的 HTTP Header，有些 API 需要传入一些 Header，可通过 wrk 的 -H 参数来传入。
–latency，打印响应时间分布。
-s，指定 Lua 脚本，Lua 脚本可以实现更复杂的请求。

$ wrk -t144 -c30000 -d30s -T30s --latency http://localhost:8080/ping
报错:unable to create thread 7: Too many open files
ulimit -a | grep open
sudo ulimit -n 1024 //修改最大文件打开数
```

## wrk示例
```
biny@BinyPc:wrk$ wrk -t144 -c1000 -d30s -T30s --latency http://localhost:8080/ping
Running 30s test @ http://localhost:8080/ping
  144 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    21.73ms   25.12ms 333.02ms   84.49% //Latency：响应时间，有平均值、标准偏差、最大值、正负一个标准差占比。
    Req/Sec   363.11     98.25     2.96k    76.20% //Req/Sec：每个线程每秒完成的请求数, 同样有平均值、标准偏差、最大值、正负一个标准差占比
  Latency Distribution //响应时间分布
     50%   16.78ms //50% 的响应时间为
     75%   33.24ms //75% 的响应时间为
     90%   51.71ms //90% 的响应时间为
     99%  106.21ms //99% 的响应时间为
  1566592 requests in 30.09s, 212.15MB read //30秒完成总请求(1566592),数据读取总量(212.15MB)
Requests/sec:  52067.10 //每秒查询数（QPS）：每秒查询数 QPS 是对一个特定的查询服务器在规定时间内所处理流量多少的衡量标准。QPS = 并发数 / 平均请求响应时间。
Transfer/sec:      7.05MB //平均每秒读取 7.05MB 数据（吞吐量）
```

# go race 竞态检测
go run -race main.go

# go vet （静态代码检查）