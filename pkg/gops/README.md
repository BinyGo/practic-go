# 进程诊断工具 gops

## 安装 
```
$ go get -u github.com/google/gops
```

# gops 命令查看
```
$ go run main.go
$ gops
161662 161510 main* unknown Go version /tmp/go-build4072872873/b001/exe/main
161684 74592  gops  go1.17.6           /golang/bin/gops
161510 23540  go    go1.17.6           /go/go1.17.6/bin/go
```

## 查看指定进程信息
```
$ gops 161662
parent PID:     161510
threads:        6
memory usage:   0.059%
cpu usage:      0.000%
cpu usage (0s): NaN%
username:       biny
cmd+args:       /tmp/go-build4072872873/b001/exe/main
elapsed time:   02:15
local/remote:   127.0.0.1:35069 <-> 0.0.0.0:0 (LISTEN)
local/remote:   :::6060 <-> :::0 (LISTEN)
```

## 查看内存使用情况
```
$ gops memstats 161662
alloc: 1.33MB (1391936 bytes)
total-alloc: 1.33MB (1391936 bytes)
sys: 9.70MB (10175752 bytes)
lookups: 0
mallocs: 642
frees: 13
heap-alloc: 1.33MB (1391936 bytes)
heap-sys: 3.69MB (3866624 bytes)
heap-idle: 1.95MB (2048000 bytes)
heap-in-use: 1.73MB (1818624 bytes)
heap-released: 1.92MB (2015232 bytes)
heap-objects: 629
...
```

## 查看运行时信息
```
$ gops stats 161662
goroutines: 2
OS threads: 7
GOMAXPROCS: 16
num CPU: 16
```

## 查看 trace 信息
```bash 
$ gops trace 161662
# 与 go tool trace 作用基本一致
```


## 查看调用栈信息
```
$ gops stack 161662
goroutine 6 [running]:
runtime/pprof.writeGoroutineStacks({0x6c6380, 0xc000114000})
        /go/go1.17.6/src/runtime/pprof/pprof.go:693 +0x70
runtime/pprof.writeGoroutine({0x6c6380, 0xc000114000}, 0x866c60)
        /go/go1.17.6/src/runtime/pprof/pprof.go:682 +0x2b
...
```
