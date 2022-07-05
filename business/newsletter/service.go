package newsletter

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strconv"
	"time"

	imgBB "github.com/JohnNON/ImgBB"

	"github.com/mashbens/cps/business/admin"
	"github.com/mashbens/cps/business/newsletter/entity"
)

type NewsRepo interface {
	FindAllNews(search string) (data []entity.News)
	FindNewsByID(newsID string) (entity.News, error)

	InsertNews(news entity.News) (entity.News, error)
	UpdateNews(news entity.News) (entity.News, error)
	DeleteNews(newsID string) error
}

type NewsService interface {
	FindAllNews(search string) (data []entity.News)
	FindNewsByID(newsID string) (*entity.News, error)

	InsertNews(news entity.News) (*entity.News, error)
	UpdateNews(news entity.News) (*entity.News, error)
	DeleteNews(adminId string, newsID string) error
}

type newsService struct {
	newsRepo     NewsRepo
	adminService admin.AdminService
}

func NewNewsService(
	NewsRepo NewsRepo,
	adminService admin.AdminService,
) NewsService {
	return &newsService{
		newsRepo:     NewsRepo,
		adminService: adminService,
	}
}

func (c *newsService) InsertNews(news entity.News) (*entity.News, error) {
	adminID := strconv.Itoa(news.AdminID)
	admin, err := c.adminService.FindAdminByID(adminID)
	if err != nil {
		return nil, err
	}
	_ = admin

	img := c.ImgUpload(news.ImgBB)
	news.Img = img

	n, err := c.newsRepo.InsertNews(news)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func (c *newsService) UpdateNews(news entity.News) (*entity.News, error) {
	adminID := strconv.Itoa(news.AdminID)
	admin, err := c.adminService.FindAdminByID(adminID)
	if err != nil {
		return nil, err
	}
	_ = admin
	img := c.ImgUpload(news.ImgBB)
	news.Img = img

	news, err = c.newsRepo.UpdateNews(news)
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func (c *newsService) DeleteNews(adminID string, newsID string) error {
	admin, err := c.adminService.FindAdminByID(adminID)
	if err != nil {
		return nil
	}
	_ = admin

	news, err := c.FindNewsByID(newsID)
	if err != nil {
		return err
	}

	n := c.newsRepo.DeleteNews(strconv.Itoa(news.ID))
	if n != nil {
		return n
	}
	return n
}
func (c *newsService) FindNewsByID(newsID string) (*entity.News, error) {
	class, err := c.newsRepo.FindNewsByID(newsID)
	if err != nil {
		return nil, err
	}

	return &class, nil
}
func (c *newsService) FindAllNews(search string) (data []entity.News) {
	data = c.newsRepo.FindAllNews(search)
	return
}

func (c *newsService) ImgUpload(file *multipart.FileHeader) string {

	src, err := file.Open()
	if err != nil {
		return fmt.Sprintln(err)
	}

	b, err := io.ReadAll(src)
	if err != nil {
		log.Fatal(err)
	}
	key := "02406488a81ff26d2a22b6306b6b21f9"
	img := imgBB.NewImage(hashSum(b), "60", b)

	bb := imgBB.NewImgBB(key, 5*time.Second)

	r, e := bb.Upload(img)
	if e != nil {
		log.Fatal(e)
	}

	return r.Data.Url
}

func hashSum(b []byte) string {
	sum := md5.Sum(b)
	return hex.EncodeToString(sum[:])
}
