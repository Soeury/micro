package bookstore

import (
	"context"
	"errors"
	"fmt"
	pan "micro/learn1/test17/panigation"
	"micro/learn1/test17/proto"
	"strconv"
	"time"

	"micro/learn1/test17/data"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

var (
	DefaultCursor   = "0" // 默认游标 ， 第一页从 0 开始
	DefaultPageSize = 2   // 默认每页的数据
)

// bookstore  RPC 服务

type Server struct {
	proto.UnimplementedBookstoreServer
	BS *data.Bookstore
}

// ListShelves 返回所有书架的 rpc 服务
func (s *Server) ListShelves(ctx context.Context, in *emptypb.Empty) (*proto.ListShelvesResponse, error) {

	sl, err := s.BS.ListShelves(ctx)
	if err == gorm.ErrEmptySlice {
		return &proto.ListShelvesResponse{}, nil
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "query failed")
	}

	// 这里吧查询返回的 *Shelf 类型转换为 *proto.Shelf 类型，用来作为返回值
	nsl := make([]*proto.Shelf, 0, len(sl))
	for _, s := range sl {
		nsl = append(nsl, &proto.Shelf{
			Id:    s.ID,
			Theme: s.Theme,
			Size:  s.Size,
		})
	}
	return &proto.ListShelvesResponse{Shelves: nsl}, nil
}

// CreateShelf 创建书架
func (s *Server) CreateShelf(ctx context.Context, in *proto.CreateShelfRequest) (*proto.Shelf, error) {

	if len(in.GetShelf().GetTheme()) <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid theme")
	}

	data := data.Shelf{
		Theme: in.GetShelf().GetTheme(),
		Size:  in.GetShelf().GetSize(),
	}

	shelf, err := s.BS.CreateShelves(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, "created failed")
	}
	return &proto.Shelf{Id: shelf.ID, Theme: shelf.Theme, Size: shelf.Size}, nil
}

// GetShelf 获取指定书架
func (s *Server) GetShelf(ctx context.Context, in *proto.GetShelfRequest) (*proto.Shelf, error) {

	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}

	shelf, err := s.BS.GetShelf(ctx, in.GetShelf())
	if err != nil {
		return nil, status.Error(codes.Internal, "get failed")
	}
	return &proto.Shelf{Id: shelf.ID, Theme: shelf.Theme, Size: shelf.Size}, nil
}

// DeleteShelf 删除指定书架
func (s *Server) DeleteShelf(ctx context.Context, in *proto.DeleteShelfRequest) (*emptypb.Empty, error) {

	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}

	err := s.BS.DeleteShelfByID(ctx, in.GetShelf())
	if err != nil {
		return nil, status.Error(codes.Internal, "delete failed")
	}
	return &emptypb.Empty{}, nil
}

// ListBook 返回图书列表
func (s *Server) ListBook(ctx context.Context, in *proto.ListBookRequest) (*proto.ListBookResponse, error) {

	// 检查参数
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}

	var (
		cursor        string = DefaultCursor
		pageSize      int    = DefaultPageSize
		hasNextPage   bool   = false
		nextPageToken string = ""
		realSize      int    = 0
	)

	// 检查分页 token
	// 用户没有指定分页就使用默认值
	// 指定了分页先检查，检查无效返回错误，检查有效就重新赋值
	if len(in.GetPageToken()) > 0 {
		pageInfo := pan.Token(in.GetPageToken()).Decode()

		if pageInfo.InValid() {
			return nil, status.Error(codes.Internal, "query failed")
		}

		cursor = pageInfo.NextID
		pageSize = int(pageInfo.PageSize)
	}

	// 查询数据库 ， 这里需要知道有没有下一页，有一个很简单的方法就是查询的时候多查一条，然后判断返回的数量和查询的数量
	books, err := s.BS.GetBookListByID(ctx, in.Shelf, cursor, int64(pageSize)+1)
	if err != nil {
		fmt.Printf("s.bs.getbooklist failed:%v\n", err)
		return nil, status.Error(codes.Internal, "query failed")
	}

	// 这里检查返回的数量和查询的页数的比较，用来判断是否有下一页，并且将真实需要的数据切换成 pageSize
	realSize = len(books)
	if len(books) > pageSize {
		hasNextPage = true  // 有下一页
		realSize = pageSize // 返回客户指定的结果
	}

	// []*data.Book  ->  []*proto.Book
	resBooks := make([]*proto.Book, 0, len(books))

	// 这里需要适量返回，因为前面 +1
	// 前面有确定返回数据的真实数量，这里使用 realSize 即可
	for i := 0; i < realSize; i++ {
		resBooks = append(resBooks, &proto.Book{
			ID:     books[i].ID,
			Author: books[i].Author,
			Title:  books[i].Title,
		})
	}

	// 如果有下一页，就生成下一个的 pagetoken
	if hasNextPage {
		nextPageInfo := pan.Page{
			NextID:        strconv.FormatInt(resBooks[realSize-1].ID, 10), // strconv.FormatInt 将一个整数转换为指定进制
			NextTimeAtUTC: time.Now().Unix(),
			PageSize:      int64(pageSize),
		}
		nextPageToken = string(nextPageInfo.Encode())
	}
	return &proto.ListBookResponse{Books: resBooks, NextPageToken: nextPageToken}, nil
}

// CreateBook 添加图书
func (s *Server) CreateBook(ctx context.Context, in *proto.CreateBookRequest) (*proto.Book, error) {

	// 检查参数 id 在数据入库之前已经检查过了
	if in.GetBook().GetTitle() == "" || in.GetBook().GetAuthor() == "" {
		return nil, errors.New("title and author cannot be empty")
	}

	data := data.Book{
		ID:      in.GetBook().GetID(),
		Author:  in.GetBook().GetAuthor(),
		Title:   in.GetBook().GetTitle(),
		ShelfID: in.GetShelf(),
	}

	book, err := s.BS.CreateBookByInfo(ctx, in.GetShelf(), data)
	if err != nil {
		return nil, status.Error(codes.Internal, "query failed")
	}
	return &proto.Book{ID: book.ID, Author: book.Author, Title: book.Title}, nil
}

// GetBook 查询指定图书
func (s *Server) GetBook(ctx context.Context, in *proto.GetBookRequest) (*proto.Book, error) {

	if in.GetBook() <= 0 || in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	book, err := s.BS.GetBookByID(ctx, in.GetShelf(), in.GetBook())
	if err != nil {
		return nil, status.Error(codes.Internal, "query failed")
	}
	return &proto.Book{ID: book.ID, Author: book.Author, Title: book.Title}, nil
}

// DeleteBook 删除指定图书
func (s *Server) DeleteBook(ctx context.Context, in *proto.DeleteBookRequest) (*emptypb.Empty, error) {

	if in.GetBook() <= 0 || in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	if err := s.BS.DeleteBookByID(ctx, in.GetShelf(), in.GetBook()); err != nil {
		return nil, status.Error(codes.Internal, "query failed")
	}

	return &emptypb.Empty{}, nil
}
