package controller

import (
	"go-service-api/global"
	"go-service-api/model"
	"go-service-api/pkg/response"
	"go-service-api/service"
	"sort"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
}

func NewMenuController() *MenuController {
	return &MenuController{}
}

// 用户获取菜单
func (m *MenuController) GetMenus(c *gin.Context) {
	// 获取用户信息
	var menus []model.Menu
	global.DB.Find(&menus)
	// 对菜单进行树形结构处理
	nodeMap := make(map[int]*model.Menu)
	var rootItems []*model.Menu
	for i, _ := range menus {
		nodeMap[int(menus[i].ID)] = &menus[i]
	}

	for i, _ := range menus {
		if menus[i].ParentId == 0 {
			rootItems = append(rootItems, &menus[i])
		} else {
			if parent, ok := nodeMap[menus[i].ParentId]; ok {
				parent.Children = append(parent.Children, &menus[i])
			}
		}
	}
	// 对菜单进行排序
	sortChildren(rootItems)
	result := make([]map[string]interface{}, 0, len(rootItems))
	for _, item := range rootItems {
		children := make([]map[string]interface{}, 0, len(item.Children))
		for _, child := range item.Children {
			children = append(children, map[string]interface{}{
				"name":      child.Name,
				"path":      child.Path,
				"component": child.Component,
				"meta": map[string]interface{}{
					"title":    child.Meta.Title,
					"icon":     child.Meta.Icon,
					"sort":     child.Meta.OrderNo,
					"hidden":   child.Meta.Hidden,
					"keeplive": child.Meta.Keeplive,
				},
			})
		}

		result = append(result, map[string]interface{}{
			"name":      item.Name,
			"path":      item.Path,
			"component": item.Component,
			"redirect":  item.Redirect,
			"meta": map[string]interface{}{
				"title":  item.Meta.Title,
				"icon":   item.Meta.Icon,
				"sort":   item.Meta.OrderNo,
				"hidden": item.Meta.Hidden,
			},
			"children": children,
		})
	}
	response.SuccessData(c, gin.H{
		"list": result,
	})
}

// 创建菜单
func (m *MenuController) Create(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.Failed(c, err.Error())
		return
	}
	if err := service.NewMenuService().Create(&menu); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}

// 更新菜单
func (m *MenuController) Update(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.Failed(c, err.Error())
		return
	}
	if err := service.NewMenuService().Update(&menu); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}

// 递归排序子节点
func sortChildren(items []*model.Menu) {
	for _, item := range items {
		if len(item.Children) > 0 {
			sort.Slice(item.Children, func(i, j int) bool {
				return item.Children[i].Meta.OrderNo < item.Children[j].Meta.OrderNo
			})
			sortChildren(item.Children)
		}
	}
}

// 获取菜单列表
func (m *MenuController) GetList(c *gin.Context) {
	var menus []model.Menu
	global.DB.Find(&menus)
	// 对菜单进行树形结构处理
	nodeMap := make(map[int]*model.Menu)
	var rootItems []*model.Menu
	for i, _ := range menus {
		nodeMap[int(menus[i].ID)] = &menus[i]
	}

	for i, _ := range menus {
		if menus[i].ParentId == 0 {
			rootItems = append(rootItems, &menus[i])
		} else {
			if parent, ok := nodeMap[menus[i].ParentId]; ok {
				parent.Children = append(parent.Children, &menus[i])
			}
		}
	}
	// 对菜单进行排序
	sortChildren(rootItems)

	response.SuccessData(c, gin.H{
		"items": rootItems,
	})
}

func (m *MenuController) GetMenuList(c *gin.Context) {
	var list []model.Menu
	global.DB.Find(&list)
	response.SuccessData(c, gin.H{
		"list": list,
	})
}

// 删除菜单
func (m *MenuController) Delete(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.Failed(c, err.Error())
		return
	}
	if err := service.NewMenuService().Delete(menu.ID); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}
