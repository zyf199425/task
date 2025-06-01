package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"blogs/repository"
	"blogs/utils"
)

const CurrentUser string = "currentLoginUser"

func StartWebServer() *gin.Engine {
	router := gin.Default()

	// 用户相关的路由
	userGourp := router.Group("/api/user")
	{
		userGourp.POST("/login", func(c *gin.Context) {
			user := repository.User{}
			err := c.ShouldBindJSON(&user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			res, token, err := Login(&user)
			if !res {
				// 登录失败
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			// 登录成功！
			c.JSON(http.StatusOK, gin.H{"token": token})
		})
		userGourp.POST("/register", func(c *gin.Context) {

			user := repository.User{}
			err := c.ShouldBindJSON(&user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			// 对密码进行加密
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
				return
			}
			user.Password = string(hashedPassword)
			res, err := Register(&user)
			if !res {
				// 注册失败
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			// 注册成功
			c.JSON(http.StatusOK, "注册成功！")
		})
	}

	// 文章相关的路由
	postGourp := router.Group("/api/post", AuthRequired())
	{
		// 创建文章
		postGourp.POST("/create", func(c *gin.Context) {
			var post repository.Post
			err := c.ShouldBindJSON(&post)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			//获取当前登录人
			currentUser, exist := c.Get(CurrentUser)
			if !exist {
				c.JSON(http.StatusBadRequest, gin.H{"error": "请先登录！"})
				c.Abort()
				return
			}
			post.UserId = currentUser.(repository.User).ID
			// 创建文章
			res, err := repository.SavePost(&post)
			if !res {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, "创建文章成功！")
		})
		// 更新文章内容
		postGourp.POST("/update", func(c *gin.Context) {
			var post repository.Post
			err := c.ShouldBindJSON(&post)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			//获取当前登录人
			currentUser, exist := c.Get(CurrentUser)
			if !exist {
				c.JSON(http.StatusBadRequest, gin.H{"error": "请先登录！"})
				c.Abort()
				return
			}
			userId := currentUser.(repository.User).ID
			storePost := repository.QueryPostById(post.ID)
			if storePost.UserId != userId {
				c.JSON(http.StatusBadRequest, gin.H{"error": "您没有权限修改该文章！"})
				c.Abort()
				return
			}
			// 更新文章
			res := repository.UpdatePostById(&post)
			if !res {
				c.JSON(http.StatusBadRequest, gin.H{"error": "更新失败！"})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, "更新文章成功！")
		})

		// 删除文章
		postGourp.POST("/deleteById", func(c *gin.Context) {
			postIdStr := c.Query("id")
			postId, _ := strconv.Atoi(postIdStr)
			storePost := repository.QueryPostById(uint(postId))
			//获取当前登录人
			currentUser, exist := c.Get(CurrentUser)
			if !exist {
				c.JSON(http.StatusBadRequest, gin.H{"error": "请先登录！"})
				c.Abort()
				return
			}
			userId := currentUser.(repository.User).ID
			if storePost.UserId != userId {
				c.JSON(http.StatusBadRequest, gin.H{"error": "您没有权限删除该文章！"})
				c.Abort()
				return
			}
			// 删除文章
			res := repository.DeletePostById(uint(postId))
			if !res {
				c.JSON(http.StatusBadRequest, gin.H{"error": "删除失败！"})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, "删除文章成功！")
		})

		// 根据ID查询文章详情
		postGourp.GET("/getById", func(c *gin.Context) {
			postIdStr := c.Query("id")
			postId, _ := strconv.Atoi(postIdStr)
			storePost := repository.QueryPostById(uint(postId))
			c.JSON(http.StatusOK, storePost)
		})
		// 查询全部文章
		postGourp.GET("/getAll", func(c *gin.Context) {
			pageNumStr := c.Query("pageNum")
			pageSizeStr := c.Query("pageSize")
			if pageNumStr == "" || pageSizeStr == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误！"})
				c.Abort()
				return
			}
			pageNum, _ := strconv.Atoi(pageNumStr)
			pageSize, _ := strconv.Atoi(pageSizeStr)
			posts := repository.QueryPostPage(pageNum, pageSize)
			c.JSON(http.StatusOK, posts)
		})
	}
	// 文章相关的路由
	commentGroup := router.Group("/api/comment", AuthRequired())
	{
		commentGroup.POST("/create", func(c *gin.Context) {
			var comment repository.Comment
			err := c.ShouldBindJSON(&comment)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			//获取当前登录人
			currentUser, exist := c.Get(CurrentUser)
			if !exist {
				c.JSON(http.StatusBadRequest, gin.H{"error": "请先登录！"})
				c.Abort()
				return
			}
			userId := currentUser.(repository.User).ID
			comment.UserId = userId
			// 创建评论
			res, err := repository.SaveComment(&comment)
			if !res {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, "创建评论成功！")
		})

		commentGroup.GET("/getByPostId", func(c *gin.Context) {
			postIdStr := c.Query("postId")
			postId, _ := strconv.Atoi(postIdStr)
			comments := repository.QueryCommentByPostId(uint(postId))
			c.JSON(http.StatusOK, comments)
		})
	}
	return router
}

// 验证登录状态
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请先登录！"})
			c.Abort()
			return
		}
		// 验证token
		res, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的token，请重新登录！"})
			c.Abort()
			return
		}
		if claims, ok := res.Claims.(jwt.MapClaims); ok && res.Valid {
			// 将claims转换为map[string]interface{}类型并返回
			var m map[string]interface{} = map[string]interface{}(claims)

			userId := m["userId"].(float64)
			userName := m["userName"].(string)
			c.Set(CurrentUser, repository.User{ID: uint(userId), UserName: userName})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}
	}
}

// 注册用户
func Register(user *repository.User) (bool, error) {
	var u = repository.QueryUserByName(user.UserName)
	if u.ID != 0 {
		return false, errors.New("用户已存在,请勿重复注册！")
	}
	return repository.SaveUser(user)
}

// 登录
func Login(user *repository.User) (bool, string, error) {
	var u = repository.QueryUserByName(user.UserName)
	if u.ID == 0 {
		return false, "", errors.New("用户不存在,请先注册！")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)); err != nil {
		return false, "", errors.New("密码错误！")
	}
	// 生成token
	token, err := utils.GenerateToken(&u)
	if err != nil {
		return false, "", errors.New("生成token失败！")
	}
	fmt.Println(token)
	return true, token, nil
}
