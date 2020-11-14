
# 仅执行一次

## 单例模式

```go
var once sync.Once
var obj *SingletonObj

func GetSingletonObj() *SingletonObj{
    once.Do(func(){
       fmt.Println("Create Singleton Obj")
       obj = &SingleObj{}
    })
    return obj
}
```
