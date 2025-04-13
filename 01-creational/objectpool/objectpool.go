package objectpool

// Object 表示对象池中具体对象.
type Object struct {
	ID int
}

// ObjectPool 表示对象池，存储和管理对象.
type ObjectPool struct {
	objects chan *Object
}

// NewObjectPool 创建对象池并初始化对象.
func NewObjectPool(size int) *ObjectPool {
	pool := &ObjectPool{
		objects: make(chan *Object, size),
	}
	for i := 0; i < size; i++ {
		object := &Object{ID: i}
		pool.objects <- object
	}
	return pool
}

// AcquireObject 从对象从中获取对象.
func (p *ObjectPool) AcquireObject() *Object {
	return <-p.objects
}

// ReleaseObject 将对象归还给对象池.
func (p *ObjectPool) ReleaseObject(object *Object) {
	p.objects <- object
}
