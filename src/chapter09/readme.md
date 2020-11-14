# 字符串与字符编码 常见的字符串函数

## 字符串
1. string 是数据类型，不是引用或指针类型
2. string 是只读的byte slice， len函数只是它所包含的byte数
3. string 的byte数组可以存放任何数据

## Unicode UTF8
1. Unicode是一种字符集（code point）
2. UTF8是Unicode的存储实现（转化为字节序列的规则）


## 常见的字符串函数

1. strings包（https://golang.org/pkg/strings/ ）
2. strconv包（https://golang.org/pkg/strconv/ ）