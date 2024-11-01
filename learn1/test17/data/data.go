package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DefaultShelfSize = 100
)

// 这里采用数据库的方式
func NewDB() (*gorm.DB, error) {

	dsn := "root:123456@(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("gorm.Open failed:%v\n", err)
		return nil, err
	}

	// 自动迁移数据库，与定义好的模型保持一致
	db.AutoMigrate(&Shelf{}, &Book{})
	return db, nil
}

type Shelf struct {
	ID       int64 `gorm:"primaryKey"`
	Theme    string
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
}

type Book struct {
	ID       int64 `gorm:"primaryKey"`
	Author   string
	Title    string
	ShelfID  int64
	CreateAt time.Time
	UpdateAt time.Time
}

type Bookstore struct {
	DB *gorm.DB
}

// CreateShelves 创建书架
func (b *Bookstore) CreateShelves(ctx context.Context, shelf Shelf) (*Shelf, error) {

	// 检查传进来的数据
	if len(shelf.Theme) <= 0 {
		return nil, errors.New("invalid theme")
	}

	size := shelf.Size
	if size == 0 {
		size = DefaultShelfSize
	}

	newshelf := Shelf{Theme: shelf.Theme, Size: size, CreateAt: time.Now(), UpdateAt: time.Now()}
	err := b.DB.WithContext(ctx).Create(&newshelf).Error
	return &newshelf, err
}

// GetShelf 获取指定书架
func (b *Bookstore) GetShelf(ctx context.Context, id int64) (*Shelf, error) {

	newShelf := Shelf{}
	err := b.DB.WithContext(ctx).First(&newShelf, id).Error
	return &newShelf, err
}

// ListShelves 返回书架列表
func (b *Bookstore) ListShelves(ctx context.Context) ([]*Shelf, error) {

	var newShelves []*Shelf
	err := b.DB.WithContext(ctx).Find(&newShelves).Error
	return newShelves, err
}

// DeleteShelfByID 删除指定书架
func (b *Bookstore) DeleteShelfByID(ctx context.Context, id int64) error {

	return b.DB.WithContext(ctx).Delete(&Shelf{}, id).Error
}

// GetBookListByID 返回图书列表
func (b *Bookstore) GetBookListByID(ctx context.Context, shelfID int64, cursor string, pagesize int64) ([]*Book, error) {

	var List []*Book
	err := b.DB.WithContext(ctx).Where("shelf_id = ? and id > ?", shelfID, cursor).Order("id asc").Limit(int(pagesize)).Find(&List).Error
	return List, err
}

// CreateBook 指定书架上添加书籍
func (b *Bookstore) CreateBookByInfo(ctx context.Context, shelfID int64, rbook Book) (*Book, error) {

	newBook := Book{
		Author:   rbook.Author,
		Title:    rbook.Title,
		ShelfID:  shelfID,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	err := b.DB.WithContext(ctx).Create(&newBook).Error
	return &newBook, err
}

// GetBookByID 根据指定 ID 查询书籍
func (b *Bookstore) GetBookByID(ctx context.Context, shelfID int64, bookID int64) (*Book, error) {

	book := Book{}
	err := b.DB.WithContext(ctx).Where("shelf_id = ? and id = ?", shelfID, bookID).Find(&book).Error
	return &book, err
}

// DeleteBookByID 根据指定 id 删除书籍
func (b *Bookstore) DeleteBookByID(ctx context.Context, shelfID int64, bookID int64) error {

	return b.DB.WithContext(ctx).Delete(&Book{}, "shelf_id = ? and id = ?", shelfID, bookID).Error
}
