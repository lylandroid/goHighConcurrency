package common

import (
	"github.com/kataras/iris/core/errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

//声明新切片类型
type units []uint32

//返回切片长度
func (x units) Len() int {
	return len(x)
}

//比对两个数大小
func (x units) Less(i, j int) bool {
	return x[i] < x[j]
}

//切片中连个值的交换
func (x units) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

//当Hash环上没有数据时，提示错误
var errEmpty = errors.New("Hash环上没有数据！")

type Consistent struct {
	//hash环，key为哈希值，值存放节点的信息
	Circle map[uint32]string
	//已经排序的节点Hash切片
	SortedHashes units
	//虚拟节点个数，用来增加hash的平衡性
	VirtualNodeCount int
	// map 读写锁
	sync.RWMutex
}

//创建一致性hash算法结构体，设置默认节点数量
func NewConsistent() *Consistent {
	return &Consistent{
		Circle:           make(map[uint32]string),
		VirtualNodeCount: 20,
	}
}

//自动生成key值
func (c *Consistent) generateKey(element string, index int) string {
	return element + strconv.Itoa(index)
}

//获取hash位置
func (c *Consistent) hashKey(key string) uint32 {
	if len(key) < 64 {
		//声明一个数组长度为64
		var srcatch [64]byte
		//拷贝数据到数组中
		copy(srcatch[:], key)
		return crc32.ChecksumIEEE(srcatch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *Consistent) updateSortedHashes() {
	hashes := c.SortedHashes[:0]
	//判断切片容量是否过大，如果过大则重置
	if cap(c.SortedHashes)/(c.VirtualNodeCount*4) > len(c.Circle) {
		hashes = nil
	}
	//添加hashes
	for k := range c.Circle {
		hashes = append(hashes, k)
	}
	//对所有节点Hash值进行排序
	//方便之后进行二分查找
	sort.Sort(hashes)
}

//添加节点
func (c *Consistent) Add(element string) {
	c.Lock()
	defer c.Unlock()
	c.add(element)
}

//添加节点
func (c *Consistent) add(element string) {
	//循环虚拟节点，设置副本
	for i := 0; i < c.VirtualNodeCount; i++ {
		c.Circle[c.hashKey(c.generateKey(element, i))] = element
	}
	//更新排序
	c.updateSortedHashes()
}

//删除节点
func (c *Consistent) remove(element string) {
	for i := 0; i < c.VirtualNodeCount; i++ {
		delete(c.Circle, c.hashKey(c.generateKey(element, i)))
	}
	c.updateSortedHashes()
}

//删除一个节点
func (c *Consistent) Remove(element string) {
	c.Lock()
	defer c.Unlock()
	c.remove(element)
}

//顺时针查找最近的节点
func (c *Consistent) search(key uint32) int {
	//查找算法
	f := func(x int) bool {
		return c.SortedHashes[x] > key
	}
	//使用"二分查找"算法来搜索指定切片满足条件的最小值
	i := sort.Search(len(c.SortedHashes), f)
	//如果超出范围设置i=0
	if i >= len(c.SortedHashes) {
		i = 0
	}
	return i
}

func (c *Consistent) Get(name string) (string, error) {
	c.RLock()
	defer c.RUnlock()
	if len(c.Circle) == 0 {
		return "", errEmpty
	}
	key := c.hashKey(name)
	i := c.search(key)
	return c.Circle[c.SortedHashes[i]], nil
}
