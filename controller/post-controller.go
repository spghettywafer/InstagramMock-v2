package controller

import (
	service "InstagramMock-v2/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	IPostController interface {
		//read, findbyid, create, update, delete,
		//, pageUpdate, pageCreate

		ShowAll(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		PageUpdate(ctx *gin.Context)
		PageCreate(ctx *gin.Context)
	}

	PostController struct {
		postService service.IPostService
	}
)

func NewPostController(s service.IPostService) *PostController {
	return &PostController{s}
}

func (c *PostController) ShowAll(ctx *gin.Context) {
	// var posts []model.Post
	posts, err := c.postService.ShowAll()
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "page-bad-request.html", nil)
		return
	}
	data := gin.H{
		"Data":      posts,
		"TotalPost": len(posts),
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}

func (c *PostController) Create(ctx *gin.Context) {
	file, errFile := ctx.FormFile("file")
	if errFile != nil {
		fmt.Println("Form File Error")
		ctx.HTML(http.StatusBadRequest, "page-bad-request.html", nil)
		return
	}

	var post service.PostCreateDTO = service.PostCreateDTO{
		UrlPhoto: file.Filename,
	}

	path := "upload/" + file.Filename
	if errUpload := ctx.SaveUploadedFile(file, path); errUpload != nil {
		fmt.Println("Save Uploaded File Error")
		ctx.HTML(http.StatusBadRequest, "page-bad-request.html", nil)
		return
	}

	if errDTO := ctx.ShouldBind(&post); errDTO != nil {
		fmt.Println("Bind Body to PostDTO Error")
		ctx.HTML(http.StatusBadRequest, "page-bad-request.html", nil)
		return
	}

	if _, err := c.postService.Create(post); err != nil {
		fmt.Println("DB Create Error")
		ctx.HTML(http.StatusBadRequest, "page-bad-request.html", nil)
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}

func (c *PostController) Update(ctx *gin.Context) {
	var post service.PostUpdateDTO

	id, errParse := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if errParse != nil {
		fmt.Println("Parse String Id to Uint Error")
		ctx.HTML(http.StatusBadRequest, "page-bad-request.html", nil)
		return
	}

	post.ID = id
	if errDTO := ctx.ShouldBind(&post); errDTO != nil {
		fmt.Println("Bind Body to PostDTO Error")
		ctx.HTML(http.StatusBadRequest, "page-bad-request.html", nil)
		return
	}

	_, err := c.postService.FindByID(id)
	if err != nil {
		ctx.HTML(http.StatusNotFound, "page-not-found.html", nil)
		return
	}

	if _, err := c.postService.Update(post); err != nil {
		fmt.Println("Update Post Error")
		ctx.HTML(http.StatusBadRequest, "page-bad-request.html", nil)
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}

func (c *PostController) Delete(ctx *gin.Context) {
	id, errParse := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if errParse != nil {
		fmt.Println("Parse String Id to Uint Error")
		ctx.HTML(http.StatusBadRequest, "page-bad-request.html", nil)
		return
	}

	_, errFindbyID := c.postService.FindByID(id)
	if errFindbyID != nil {
		ctx.HTML(http.StatusNotFound, "page-not-found.html", nil)
		return
	}

	if err := c.postService.Delete(id); err != nil {
		fmt.Println("Update Post Error")
		ctx.HTML(http.StatusBadRequest, "page-bad-request.html", nil)
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}

func (c *PostController) FindByID(ctx *gin.Context) {
	id, errParse := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if errParse != nil {
		fmt.Println("Parse String Id to Uint Error")
		ctx.HTML(http.StatusBadRequest, "page-bad-request.html", nil)
		return
	}

	post, err := c.postService.FindByID(id)
	if err != nil {
		ctx.HTML(http.StatusNotFound, "page-not-found.html", nil)
		return
	}

	ret := gin.H{
		"Data": post,
	}
	ctx.HTML(http.StatusOK, "post-detail.html", ret)
}

func (c *PostController) PageCreate(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.html", nil)
}

func (c *PostController) PageUpdate(ctx *gin.Context) {
	id, errParse := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if errParse != nil {
		fmt.Println("Parse String Id to Uint Error")
		ctx.HTML(http.StatusBadRequest, "page-bad-request.html", nil)
		return
	}

	post, err := c.postService.FindByID(id)
	if err != nil {
		ctx.HTML(http.StatusNotFound, "page-not-found.html", nil)
		return
	}

	ret := gin.H{
		"Data": post,
	}
	ctx.HTML(http.StatusOK, "post-edit.html", ret)
}
