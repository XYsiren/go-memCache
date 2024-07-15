# 内存缓存系统

## 内容：

1.支持设定过期时间，精度到秒

2.支持设定最大内存，当内存超过时做出合适的处理

3.支持并发安全

4.按以下接口实现：

```golang
//size : 1KB 100KB 1MB 2MB 1GB
	SetMaxMemory(size string) bool
	//将value写入缓存
	Set(key string, val interface{}, expire time.Duration) bool
	//根据key值获取value
	Get(key string) (interface{}, bool)
	//删除key值
	Del(key string) bool
	//判断key是否存在
	Exists(key string) bool
	//清空所有key
	Flush() bool
	//获取缓存中所有key的数量
	Keys() int64
```

