package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"net/url"
	"bytes"
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/Luxurioust/excelize"
	"io"
	"path"
	"time"

	"strconv"
	"strings"
)

var (
	email  string =""
	html  string =""
	i     int

)

type MainController struct {
	beego.Controller
}

func (c *MainController) Email() {
	c.TplName = "index.html"
}
func (c *MainController ) Uplaoduser() {
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()
	file, head, err := c.GetFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}


	//当期时间格式化
	filename := time.Now().Format("20060102150405")
	//获取文件的后缀
	fileSuffix := path.Ext(head.Filename)


	filePath := "./uplaod/" + filename + fileSuffix
	//创建文件
	fW, err := os.Create(filePath)

	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)

	if err != nil {
		fmt.Println("文件保存失败")
		return
	}
	if fileSuffix == ".xlsx" {

	email,i=fileXlsx(filePath)
		fmt.Println(i,"===========================")
		c.Ctx.WriteString("{\"code\":200,\"count\":"+strconv.Itoa(i)+"}")
		i=0
	return
	}
	if fileSuffix == ".html" {
		dat, _ := ioutil.ReadFile(filePath)
		html=string(dat)
	}
	c.Ctx.WriteString("{\"code\":200}")
}
func (c *MainController ) SendMail(){
	sub:=c.GetString("sub")
	fmt.Println(sub)
	if sub==""{
		c.Ctx.WriteString("{\"code\":410}")
		return
	}
	fmt.Println(email,"-------email--------")
	if email==""{
		fmt.Println(email)
		c.Ctx.WriteString("{\"code\":411}")
		return
	}
	if html==""{
		c.Ctx.WriteString("{\"code\":412}")
		return
	}
	if sub!=""&&email!=""&&html!=""{
		list:=strings.Split(email,";")
		for _,v :=range list{
		sendMail(v,sub,html)
			}
		}
		clear()
		c.Ctx.WriteString("{\"code\":200}")
	}


func sendMail(to string,sub string,html string) {
	RequestURI := "http://api.sendcloud.net/apiv2/mail/send"
	PostParams := url.Values{
		"apiUser": {"dh_market"},
		"apiKey":  {"I5UQX23RJbLZTir2"},
		"from":     {"marketing@datahunter.cn"},
		"fromName": {"DHmarketing"},
		"to":       {to},
		"subject":  {sub},
		"html":     {html},
	}
	PostBody := bytes.NewBufferString(PostParams.Encode())
	ResponseHandler, err := http.Post(RequestURI, "application/x-www-form-urlencoded", PostBody)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ResponseHandler.Body.Close()
	BodyByte, err := ioutil.ReadAll(ResponseHandler.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(BodyByte))
}
func fileXlsx(filePath string)(toemail string,i int) {
	var email string

	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows :=xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for k, colCell := range row {

			if k%2==1{
				i=1+i
				email=email+colCell+";"
				fmt.Print(colCell, "\t")

			}

		}
		fmt.Println()
	}
	return email,i

}

func clear() {
	email=""
	html=""
}
