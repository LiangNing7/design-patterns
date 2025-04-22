package memento

import "fmt"

// History 定义历史记录结构体.
type History struct {
	body string // 历史记录内容.
}

// GetBody 获取历史记录内容.
func (h *History) GetBody() string {
	return h.body
}

// SetBody 设置历史记录内容.
func (h *History) SetBody(body string) {
	h.body = body
}

// NewHistory 创建新的历史记录对象实例.
func NewHistory(body string) *History {
	return &History{body: body}
}

// Doc 定义文档结构体.
type Doc struct {
	title string // 文档标题.
	body  string // 文档内容.
}

// Title 获取文档标题.
func (d *Doc) Title() string {
	return d.title
}

// SetTitle 设置文档标题.
func (d *Doc) SetTitle(title string) {
	d.title = title
}

// Body 获取文档内容.
func (d *Doc) Body() string {
	return d.body
}

// SetBody 设置文档内容.
func (d *Doc) SetBody(body string) {
	d.body = body
}

// NewDoc 创建新的文档对象实例.
func NewDoc(title string) *Doc {
	return &Doc{title: title}
}

// CreateHistory 创建文档的历史记录.
func (d *Doc) CreateHistory() *History {
	return &History{body: d.body}
}

// RestoreHistory 恢复文档的历史记录.
func (d *Doc) RestoreHistory(history *History) {
	d.body = history.GetBody()
}

// Editor 定义编辑器结构体.
type Editor struct {
	doc       *Doc       // 当前编辑的文档.
	histories []*History // 编辑历史记录数组.
	position  int        // 当前历史记录位置.
}

// NewEditor 创建新的编辑器实例.
func NewEditor(doc *Doc) *Editor {
	fmt.Println("打开文档" + doc.Title())
	e := &Editor{doc: doc}
	e.histories = make([]*History, 0)
	return e
}

// Append 在文档末尾添加文本内容.
func (e *Editor) Append(text string) {
	e.backup()

	e.doc.SetBody(e.doc.Body() + text + "\n")

	fmt.Println("===> 插入操作，文档内容如下：")
	e.Show()
}

// Save() 保存文档操作.
func (e *Editor) Save() {
	fmt.Println("===> 存盘操作")
}

// Delete 删除文档内容.
func (e *Editor) Delete() {
	e.backup()

	fmt.Println("===> 删除操作，文档内容如下：")
	e.doc.SetBody("")
}

// Show 显示文档内容
func (e *Editor) Show() {
	fmt.Println(e.doc.Body())
}

// backup 备份当前文档状态
func (e *Editor) backup() {
	e.histories = append(e.histories, e.doc.CreateHistory())
	e.position++
}

// Undo 撤销操作，恢复历史状态
func (e *Editor) Undo() {
	// 到头了不可以再撤回
	if e.position == 0 {
		return
	}
	e.position--
	history := e.histories[e.position]
	e.doc.RestoreHistory(history)

	fmt.Println("===> 撤销操作，文档内容如下：")
	e.Show()
}
