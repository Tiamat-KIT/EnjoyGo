package main

import (
	/* "bytes"
	"encoding/json" */
	"fmt"
	"html/template"
	/* "io"
	"net/http"*/
	"log"
	"os" 
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/dstotijn/go-notion"
	"context"
)

func LoadEnv() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// fmt.Println("ログイン確認処理 前")
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Errorf("Fail load: %v",err)
		}
		ctx.Next()
		// fmt.Println("ログイン確認処理 後")
	}
}

func replacer(text string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br />", -1))
}

func main() {
	r := gin.Default()
	r.Use(LoadEnv())
	r.SetFuncMap(
		template.FuncMap{
			"nl2br": replacer,
		},
	)
	r.LoadHTMLGlob("template/*.html")
	r.GET("/",func(ctx *gin.Context) {
		ctx.HTML(200,"index.html",nil)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/tech",func(ctx *gin.Context) {
		/* http_client := &http.Client{}
		req,req_err := http.NewRequest("GET","https://api.notion.com/v1/pages/Tech-Note-cf3fb4b15bf64b4da94cb976dbf36e24",nil)
		if req_err != nil {
			log.Fatal(req_err)
		}
		notion_api_key := os.Getenv("NOTION_API_KEY")
		req.Header.Add("Authorization", "Bearer" + notion_api_key)
		req.Header.Add("Notion-Version", "2021-05-13")

		resp,resp_err := http_client.Do(req)
		if resp_err != nil {
			log.Fatal(resp_err)
		}
		defer resp.Body.Close()
		b,err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var out bytes.Buffer
		json.Indent(&out,b,""," ")
		ctx.JSON(200,out) */
		notion_api_key := os.Getenv("NOTION_API_KEY")
		client := notion.NewClient(notion_api_key)
		page,err := client.FindPageByID(context.Background(),"cf3fb4b15bf64b4da94cb976dbf36e24")
		if err != nil {
			log.Fatal(err)
		}

		ctx.JSON(200,page)
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}

