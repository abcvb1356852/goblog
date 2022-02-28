package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
	"strconv"
	"unicode/utf8"

	"gorm.io/gorm"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
}

// Show 文章详情页面
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4.读取成功 显示文章
		/* 		tmpl, err := template.ParseFiles("./resources/views/articles/show.gohtml") */
		/* 		tmpl, err := template.New("show.gohtml").Funcs(template.FuncMap{
		   			"RouteName2URL":  route.Name2URL,
		   			"Uint64ToString": types.Uint64ToString,
		   		}).ParseFiles("./resources/views/articles/show.gohtml")

		   		if err != nil {
		   			logger.LogError(err)
		   		}

		   		err = tmpl.Execute(w, article)
		   		logger.LogError(err) */

		// 4.0 设置模板相对路径
		/* 		viewDir := "resources/views"

		   		// 4.1 所有布局模板文件 Slice
		   		files, err := filepath.Glob(viewDir + "/layouts/*.gohtml")
		   		logger.LogError(err)

		   		// 4.2 在 Slice 里新增我们的目标文件
		   		newFiles := append(files, viewDir+"/articles/show.gohtml")

		   		// 4.3 解析模板文件
		   		tmpl, err := template.New("show.gohtml").
		   			Funcs(template.FuncMap{
		   				"RouteName2URL":  route.Name2URL,
		   			}).ParseFiles(newFiles...)
		   		logger.LogError(err)

		   		// 4.4 渲染模板，将所有文章的数据传输进去
		   		err = tmpl.ExecuteTemplate(w, "app", article)
		   		logger.LogError(err) */
		// ---  4. 读取成功，显示文章 ---
		view.Render(w, view.D{
			"Article": article,
		}, "articles.show")
	}
}

// Index 文章列表页
func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {

	// 1. 获取结果集
	articles, err := article.GetAll()

	if err != nil {
		// 数据库错误
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	} else {
		// 2. 加载模板
		/* 		tmpl, err := template.ParseFiles("./resources/views/articles/index.gohtml")
		   		if err != nil {
		   			logger.LogError(err)
		   		}

		   		// 3. 渲染模板，将所有文章的数据传输进去
		   		err = tmpl.Execute(w, articles)
		   		logger.LogError(err) */

		// 2.0 设置模板相对路径
		/* 		viewDir := "resources/views"

		   		// 2.1 所有布局模板文件 Slice
		   		files, err := filepath.Glob(viewDir + "/layouts/*.gohtml")
		   		logger.LogError(err)

		   		// 2.2 在 Slice 里新增我们的目标文件
		   		newFiles := append(files, viewDir+"/articles/index.gohtml")

		   		// 2.3 解析模板文件
		   		tmpl, err := template.ParseFiles(newFiles...)
		   		logger.LogError(err)

		   		// 2.4 渲染模板，将所有文章的数据传输进去
		   		err = tmpl.ExecuteTemplate(w, "app", articles)
		   		logger.LogError(err) */
		// ---  2. 加载模板 ---
		view.Render(w, view.D{
			"Articles": articles,
		}, "articles.index")
	}
}

// Create 文章创建页面
func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
}

func validateArticleFormData(title string, body string) map[string]string {
	errors := make(map[string]string)
	// 验证标题
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度需介于 3-40"
	}

	// 验证内容
	if body == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需大于或等于 10 个字节"
	}

	return errors
}

// Store 文章创建页面
func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := validateArticleFormData(title, body)

	// 检查是否有错误
	if len(errors) == 0 {
		_article := article.Article{
			Title: title,
			Body:  body,
		}
		_article.Create()
		if _article.ID > 0 {
			fmt.Fprint(w, "插入成功，ID 为"+strconv.FormatUint(_article.ID, 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章失败，请联系管理员")
		}
	} else {
		/* 		view.Render(w, "articles.create", ArticlesFormData{
			Title:  title,
			Body:   body,
			Errors: errors,
		}) */
		view.Render(w, view.D{
			"Title":  title,
			"Body":   body,
			"Errors": errors,
		}, "articles.create", "articles._form_field")
	}
}

// Edit 文章更新页面
func (*ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {

	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_article, err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		/* 		// 4. 读取成功，显示编辑文章表单
		   		updateURL := route.Name2URL("articles.update", "id", id)
		   		data := ArticlesFormData{
		   			Title:  article.Title,
		   			Body:   article.Body,
		   			URL:    updateURL,
		   			Errors: nil,
		   		}
		   		tmpl, err := template.ParseFiles("./resources/views/articles/edit.gohtml")
		   		logger.LogError(err)

		   		err = tmpl.Execute(w, data)
		   		logger.LogError(err) */
		// 4. 读取成功，显示编辑文章表单
		view.Render(w, view.D{
			"Title":   _article.Title,
			"Body":    _article.Body,
			"Article": _article,
			"Errors":  nil,
		}, "articles.edit", "articles._form_field")
	}
}

// Update 更新文章
func (*ArticlesController) Update(w http.ResponseWriter, r *http.Request) {

	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_article, err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 未出现错误

		// 4.1 表单验证
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")

		errors := validateArticleFormData(title, body)

		if len(errors) == 0 {

			// 4.2 表单验证通过，更新数据
			_article.Title = title
			_article.Body = body

			rowsAffected, err := _article.Update()

			if err != nil {
				// 数据库错误
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
				return
			}

			// √ 更新成功，跳转到文章详情页
			if rowsAffected > 0 {
				showURL := route.Name2URL("articles.show", "id", id)
				http.Redirect(w, r, showURL, http.StatusFound)
			} else {
				fmt.Fprint(w, "您没有做任何更改！")
			}
		} else {

			// 4.3 表单验证不通过，显示理由
			/*
				updateURL := route.Name2URL("articles.update", "id", id)
				data := ArticlesFormData{
					Title:  title,
					Body:   body,
					URL:    updateURL,
					Errors: errors,
				}
				tmpl, err := template.ParseFiles("./resources/views/articles/edit.gohtml")
				logger.LogError(err)

				err = tmpl.Execute(w, data)
				logger.LogError(err) */

			view.Render(w, view.D{
				"Title":   title,
				"Body":    body,
				"Article": _article,
				"Errors":  errors,
			}, "articles.edit", "articles._form_field")
		}
	}
}

func (*ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_article, err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 未出现错误
		rowsAffected, err := _article.Delete()

		// 4.1 发生错误
		if err != nil {
			// 应该是 SQL 报错了
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		} else {
			// 4.2 未发生错误
			if rowsAffected > 0 {
				// 重定向到文章列表页
				indexURL := route.Name2URL("articles.index")
				http.Redirect(w, r, indexURL, http.StatusFound)
			} else {
				// Edge case
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "404 文章未找到")
			}
		}
	}
}
