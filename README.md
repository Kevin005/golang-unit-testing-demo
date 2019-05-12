# 一个开箱即用的单元测试demo

1. 单元测试
```go
// 执行TestHelloWorld测试程序
// -run 需要测试的函数名，注意和性能测试的文件名区分
// -v 输出测试过程信息
// -count 执行次数
go test -run TestHelloWorld -v -count 1
```
2. 性能(压力)测试
```go
// 执行BenchmarkHelloWorld测试程序
// -run 需要测试的.go文件名，如 hello_world_test.go则文件名为hello_world_test
// -bench 需要测试的函数名
// -count 执行次数
go test -run 文件名 -bench=TestHelloWorld_PrintHello -count=3
```
3. 覆盖率测试
```go
// 生成cover.out文件，并输出覆盖率
go test -v -coverprofile=cover.out
// 根据cover.out文件生成coverage.html可视化文件
go tool cover -html=cover.out -o coverage.html
```