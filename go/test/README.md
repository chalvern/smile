# Go 的测试

关于 Go 语言的测试的几个有意思的点：

1. 所有的测试用例写在名为 `*_test.go` 的文件中。
1. 每个目录可以同时定义 `package` 和 `package_test` 包。比如当前目录中同时存在 `test` 和 `test_test` 包，其中后者是前者的专门定制的测试包（除此之外，不允许在一个目录中定义其他的包）。
1. 我们当然可以把测试用例直接定义在名为 `*_test.go` 文件中定义的 `test` 包里。如果把测试用例定义在 `package_test` 包里，测试用例不能访问 `package` 包里的私有函数或方法（黑盒测试）；如果把测试用例定义在 `package` 包里，测试用例可以测试私有函数或方法（白盒测试）。
1. 同一个包的测试用例里只允许存在一个 `func TestMain(m *testing.M)` 函数，即使存在 `package` 和 `package_test` 两个包，也是如此。
1. 同一个包里所有 `Test*(t *testing.T)` 的函数的执行都是依次按顺序执行的。只有上一个测试用例结束了，才会开始下一个测试用例。（这个特性还是很不错的，让测试用例的执行变得可控了）


## 实际编写 web 测试用例的一些考虑

1. 各个测试用例之间的数据不能相互干扰
1. 构造测试需要的数据（构造场景）

可以参考 [apollo](https://github.com/chalvern/apollo) 中的测试用例的写法看 go 的 web 系统的测试用例的编写方式。