package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct {
	Id      string   `json:"id"`      //图书id ISBN ID
	Name    string   `json:"name"`    //图书名称
	Authors []string `json:"authors"` //图书作者
	Press   string   `json:"press"`   //出版社
}

type Store interface {
	Create(*Book) error       //创建图书信息
	Delete(string) error      //删除图书信息
	Update(*Book) error       //修改图书信息
	Get(string) (Book, error) //查询某个图书信息
	GetAll() ([]Book, error)  //获取所有图书信息
}
