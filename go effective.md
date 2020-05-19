# go effective

## 名字

1.包名尽量简洁

2.不提供 get set 方法

3.单个方法的接口后缀使用"er"，如果写的方法实现了一个众所周知的方法，则使用相同的方法和签名

4.使用混合大小写而不是下划线来命名

##分号

如果换行出现在可以结束的一条语句的符号之后，则插入一个分号。

##控制结构

1.if 语句 省略不必要的 else 语句，往往结束于 break,continue,goto,return

2.for 循环 共三种使用方式

    for init; condition; post { }

    for condition { }

    for { }

在数组，切片，字符串，map，channel 上进行读取时，可以使用 range。仅需使用第一项时可以丢弃第二个，仅需第二项时第一个用下划线表示。

3.switch

- switch 后面不跟表达式 则按照 true 进行匹配，链接成 if else 的结构

        func unhex(c byte) byte {
        switch {
        case '0' <= c && c <= '9':
            return c - '0'
        case 'a' <= c && c <= 'f':
            return c - 'a' + 10
        case 'A' <= c && c <= 'F':
            return c - 'A' + 10
        }
        return 0
        }

- switch 不会自动从一个 case 子句跌落到下一个 case 子句。但是 case 可以使用逗号分隔的列表

  func shouldEscape(c byte) bool {
  switch c {
  case ' ', '?', '&', '=', '#', '+', '%':
  return true
  }
  return false
  }

- 可用于作类型判断

        var t interface{}
        t = functionOfSomeType()
        switch t := t.(type) {
        default:
            fmt.Printf("unexpected type %T", t)
        case bool:
            fmt.Printf("boolean %t\n", t)
        case int:
            fmt.Printf("integer %d\n", t)
        case *bool:
            fmt.Printf("pointer to boolean %t\n", *t)
        case *int:
            fmt.Printf("pointer to integer %d\n", *t)
        }
