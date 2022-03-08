package lru

import "container/list"

type Cache struct {
	maxBytes  int64                         //最大使用内存
	nbytes    int64                         //当前已使用内存
	ll        *list.List                    //双向链表
	cache     map[string]*list.Element      //键是字符串,值是双向链表中对应节点的指针
	OnEvicted func(key string, value Value) //某条记录被移除时的回调函数
}

//双向链表节点的数据类型
//在链表中仍保存每个值对应的key的好处在于,淘汰队首节点时,需要用key从字典中删除对应的映射
type entry struct {
	key   string
	value Value
}

//为了通用,我们允许值是实现了Value接口的任意类型
//Len() int 返回值所占用的内存大小
type Value interface {
	Len() int
}

//Cache 构造函数
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

//根据key查询Cache
//第一步是从字典中找到对应的双向链表的节点
//第二步，将该节点移动到队尾。
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

//删除最少访问的元素
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

//插入元素
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	//内存不够时删除最久未被访问的元素
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

//cache 中元素个数
func (c *Cache) Len() int {
	return c.ll.Len()
}
