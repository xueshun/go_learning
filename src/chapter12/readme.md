
# 扩展与复用

## 复合

Go不支持继承，但可以通过复合的方式来复用

## 匿名类型嵌入

它不是继承，如果把“内部 struct”看作父类，把“外部struct”看作子类，会发现如下问题：
1. 不支持子类替换
2. 子类并不是真正继承了父类方法

- 父类定义的方法无法访问子类的数据和方法