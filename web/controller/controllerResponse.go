/**
  author: kevin
 */
package controller

import (
	"net/http"
	"path/filepath"
	"html/template"
	"fmt"
)

func showView(w http.ResponseWriter, r *http.Request, templateName string, data interface{})  {
	page := filepath.Join("web", "tpl", templateName)

	// 创建模板实例
	resultTemplate, err := template.ParseFiles(page)
	if err != nil {
		fmt.Println("创建模板实例错误: ", err)
		return
	}

	// 融合数据
	err = resultTemplate.Execute(w, data)
	if err != nil {
		fmt.Println("融合模板数据时发生错误", err)
		return
	}
}