package registry

// Registry 定义注册表结构.
type Registry struct {
	registry map[string]any
}

// Register 方法用于向注册表中注册对象.
func (r *Registry) Register(key string, value any) {
	r.registry[key] = value
}

// Get 方法用于从注册表中检索对象.
func (r *Registry) Get(key string) any {
	return r.registry[key]
}
