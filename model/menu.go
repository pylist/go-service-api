package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Menu struct {
	Model
	ParentId  int       `json:"parentId" gorm:"comment:父菜单ID"`
	Path      string    `json:"path" gorm:"not null;comment:路由path"`        // 是当前路由的路径，会与配置中的父级节点的 path 组成该页面路由的最终路径；如果需要跳转外部链接，可以将path设置为 http 协议开头的路径。
	Name      string    `json:"name" gorm:"not null;unique;comment:菜单名称"`   // 影响多标签 Tab 页的 keep-alive 的能力，如果要确保页面有 keep-alive 的能力，请保证该路由的name与对应页面（SFC)的name保持一致。
	Component string    `json:"component" gorm:"not null;comment:对应前端文件路径"` // 渲染该路由时使用的页面组件
	Redirect  string    `json:"redirect" gorm:"comment:重定向路径"`              // 重定向的路径
	Meta      *MenuMeta `json:"meta" gorm:"type:json;comment:附加属性"`         // 主要用途是路由在菜单上展示的效果的配置
	Children  []*Menu   `json:"children" gorm:"-"`
}

type MenuMeta struct {
	Title            string `json:"title" gorm:"comment:菜单名"`                  // 该路由在菜单上展示的标题
	Icon             string `json:"icon" gorm:"comment:菜单图标"`                  // 该路由在菜单上展示的图标
	Expand           bool   `json:"expand" gorm:"comment:是否展开"`                // 决定该路由在菜单上是否默认展开
	OrderNo          int    `json:"orderNo" gorm:"comment:排序"`                 // 该路由在菜单上展示先后顺序，数字越小越靠前，默认为零
	Hidden           bool   `json:"hidden" gorm:"comment:是否隐藏"`                // 决定该路由是否在菜单上进行展示
	HiddenBreadcrumb bool   `json:"hiddenBreadcrumb" gorm:"comment:是否隐藏面包屑"`   // 如果启用了面包屑，决定该路由是否在面包屑上进行展示
	Single           bool   `json:"single" gorm:"comment:是否单独显示"`              // 如果是多级菜单且只存在一个节点，想在菜单上只展示一级节点，可以使用该配置。请注意该配置需配置在父节点
	FrameSrc         string `json:"frameSrc" gorm:"comment:内嵌iframe的地址"`       // 内嵌 iframe 的地址
	FrameBlank       bool   `json:"frameBlank" gorm:"comment:内嵌iframe是否新窗口打开"` // 内嵌 iframe 的地址是否以新窗口打开
	Keeplive         bool   `json:"keeplive" gorm:"comment:是否缓存"`              // 可决定路由是否开启keep-alive，默认开启
	Sort             int    `json:"sort" gorm:"comment:排序"`
}

func (m *MenuMeta) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &m)
}

func (m MenuMeta) Value() (driver.Value, error) {
	return json.Marshal(m)
}
