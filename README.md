// go项目开发命名原则

文件命名
----

文件命名一律采用小写，不用驼峰式，尽量见名思义，看见文件名就可以知道这个文件下的大概内容。  
其中测试文件以 _test.go 结尾，除测试文件外，命名不出现_。

例子：

> stringutil.go， stringutil_test.go

包名 package
----------

包名用小写, 使用短命名, 尽量和标准库不要冲突。  
包名统一使用单数形式。

变量
--

变量命名一般采用驼峰式，当遇到特有名词（缩写或简称，如 DNS）的时候，特有名词根据是否私有全部大写或小写。

例子：

> apiClient、URLString

常量
--

同变量规则，力求语义表达完整清楚，不要嫌名字长。  
如果模块复杂，为避免混淆，可按功能统一定义在 package 下的一个文件中。

接口
--

单个函数的接口名以 er 为后缀

```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

两个函数的接口名综合两个函数名，如:

```
type WriteFlusher interface {
    Write([]byte) (int, error)
    Flush() error
}
```

三个以上函数的接口名类似于结构体名，如:

```
type Car interface {
    Start() 
    Stop()
    Drive()
}
```

结构体
---

结构体名应该是名词或名词短语，如 Account,Book，避免使用 Manager 这样的。  
如果该数据结构需要序列化，如 json， 则首字母大写， 包括里面的字段。

方法
--

方法名应该是动词或动词短语，采用驼峰式。将功能及必要的参数体现在名字中， 不要嫌长， 如 updateById，getUserInfo.

如果是结构体方法，那么 Receiver 的名称应该缩写，一般使用一个或者两个字符作为 Receiver 的名称。如果 Receiver 是指针， 那么统一使用 p。 如：

```
func (f foo) method() {
    ...
}
```

```
func (p *foo) method() {
    ...
}
```

对于 Receiver 命名应该统一， 要么都使用值， 要么都用指针。

每个包都应该有一个包注释，位于 package 之前。如果同一个包有多个文件，只需要在一个文件中编写即可；如果你想在每个文件中的头部加上注释，需要在版权注释和 Package 前面加一个空行，否则版权注释会作为 Package 的注释。如：

```
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package net
```

每个以大写字母开头（即可以导出）的方法应该有注释，且以该函数名开头。如：

```
// Get 会响应对应路由转发过来的 get 请求
func (c *Controller) Get() {
    ...
}
```

大写字母开头的方法以为着是可供调用的公共方法，如果你的方法想只在本包内掉用，请以小写字母开发。如:

```
func (c *Controller) curl() {
    ...
}
```

注释应该用一个完整的句子，注释的第一个单词应该是要注释的指示符，以便在 godoc 中容易查找。

注释应该以一个句点 . 结束。

# 语言层面
## 关于init函数
> 1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
>
> 2 每个包可以拥有多个init函数
>
> 3 包的每个源文件也可以拥有多个init函数
>
> 4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)
>
> 5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
>
> 6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
>

## 关于指针接收器
>
> 接收者有两种，一种是值接收者，一种是指针接收者。顾名思义，值接收者，是接收者的类型是一个值，是一个副本，方法内部无法对其真正的接收者做更改；指针接收者，接收者的类型是一个指针，是接收者的引用，对这个引用的修改之间影响真正的接收者。下面看一个最基本的例子。