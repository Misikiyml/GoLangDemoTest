package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type BookStore struct {
	XMLName xml.Name `xml:"bookstore"`
	Books   []Book   `xml:"book"`
}
type Book struct {
	XMLName  xml.Name `xml:"book"`
	Category string   `xml:"category,attr"`
	Title    string   `xml:"title"`
	Author   string   `xml:"author"`
	Year     string   `xml:"year"`
	Price    string   `xml:"price"`
}

func main() {
	//获取xml文件位置
	content, err := ioutil.ReadFile("./bookstore.xml")
	if err != nil {
		fmt.Println("读文件位置错误信息：", err)
	}

	//将文件转换为对象
	var result BookStore
	err = xml.Unmarshal(content, &result)
	if err != nil {
		fmt.Println("读文件内容错误信息：", err)
	}

	fmt.Println("xml读取:", result)
	fmt.Println()
	for _, books := range result.Books {
		fmt.Println("类别:", books.Category)
		fmt.Println("题目:", books.Title)
		fmt.Println("作者:", books.Author)
		fmt.Println("出版年份:", books.Year)
		fmt.Println("价格:", books.Price)
		fmt.Println()
	}

}
