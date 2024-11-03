package controller

import (
	"fmt"
	"go-service-api/global"
	"go-service-api/model"
	"go-service-api/pkg"
	"go-service-api/pkg/orm"
	"go-service-api/pkg/response"
	"go-service-api/pkg/util"
	"go-service-api/service"
	"io"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

type UserCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 创建用户
func (u *UserController) Create(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, err.Error())
		return
	}
	user := req
	if err := service.NewUserService().Create(&user); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}

// 用户注册
type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (u *UserController) Register(c *gin.Context) {
	var req UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
	}
	if err := service.NewUserService().Create(&user); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 用户登录
func (u *UserController) Login(c *gin.Context) {
	var req UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, err.Error())
		return
	}
	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}
	newUser, err := service.NewUserService().Verify(&user)
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	token, err := pkg.NewToken().SigningToken(newUser.Username)
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.SuccessData(c, map[string]interface{}{
		"token": "Bearer " + token,
	})
}

// 用户信息
func (u *UserController) Info(c *gin.Context) {
	username := ""
	if claims, ok := c.Get("claims"); ok {
		fmt.Println(claims)
		username = claims.(*pkg.MyClaims).Username
	}
	var user model.User
	user.Username = username
	if err := service.NewUserService().Info(&user); err != nil {
		response.Failed(c, err.Error())
		return

	}
	response.SuccessData(c, map[string]interface{}{
		"nickName": user.NickName,
		"username": user.Username,
		"avatar":   user.Avatar,
	})
}

func (u *UserController) Logout(c *gin.Context) {
	response.Success(c)
}

// 更新用户
func (u *UserController) Update(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Failed(c, err.Error())
		return
	}
	if err := service.NewUserService().Update(&user); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}

// 删除用户
func (u *UserController) Delete(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Failed(c, err.Error())
		return
	}
	if err := service.NewUserService().Delete(user.ID); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}

// 用户列表
func (u *UserController) List(c *gin.Context) {
	var users []model.User
	username := c.Query("username")
	nickName := c.Query("nickName")
	deptId := c.Query("deptId")
	tx := global.DB.Model(&model.User{})
	if username != "" {
		tx.Where("username like ?", "%"+username+"%")
	}
	if nickName != "" {
		tx.Where("nick_name like ?", "%"+nickName+"%")
	}
	if deptId != "" && deptId != "0" {
		var dept model.Dept
		global.DB.Where("id = ?", deptId).First(&dept)
		deptIds := []int{int(dept.ID)}
		// 如果有父级部门，也要查询出来
		if dept.ParentID != 0 {
			deptIds = append(deptIds, int(dept.ParentID))
		}
		// 查询子部门
		var children []model.Dept
		global.DB.Where("parent_id = ?", dept.ID).Find(&children)
		for _, v := range children {
			deptIds = append(deptIds, int(v.ID))
		}

		tx.Where("dept_id IN ?", deptIds)
	}
	orm.Paginate(c)(tx).Preload("Role").Preload("Dept").Find(&users)
	response.SuccessData(c, gin.H{
		"list": users,
	})
}

func (u *UserController) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Failed(c, err.Error())
		return
	}

	// 读取文件内容
	uploadedFile, _ := file.Open()
	defer uploadedFile.Close()
	data, _ := io.ReadAll(uploadedFile)
	path, err := util.NewAvatarSave("").Save(data)
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	fmt.Println(path)
	response.SuccessData(c, gin.H{
		"url": path,
	})
}

// 重置密码
func (u *UserController) ResetPassword(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}

// 获取用户权限
func (u *UserController) GetPermission(c *gin.Context) {
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
					"sort":     child.Meta.Sort,
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
				"sort":   item.Meta.Sort,
				"hidden": item.Meta.Hidden,
			},
			"children": children,
		})
	}
	response.SuccessData(c, gin.H{
		"menus": result,
	})
}
