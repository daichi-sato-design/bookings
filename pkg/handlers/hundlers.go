package handlers

import (
	"fmt"
	"net/http"

	"github.com/daichi-sato-design/bookings/pkg/config"
	"github.com/daichi-sato-design/bookings/pkg/models"
	"github.com/daichi-sato-design/bookings/pkg/render"
)

// Repo ハンドラーが使用するリポジトリー
var Repo *Repository

// Repository はリポジトリタイプ
type Repository struct{
	App *config.AppConfig
}

// NewRepo は新しいリポジトリを作成
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

// NewHandlers はハンドラーのリポジトリーを設定
func NewHandlers(r *Repository){
	Repo = r
}

// Home ホームページハンドラー
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

// About はaboutページハンドラー
func  (m *Repository)  About(w http.ResponseWriter, r *http.Request){
	// いくつかのロジックを実行
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// テンプレートにデータを送信
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}


// Reservation 予約の作成ページと表示フォームを表示
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{})
}

// Generals 部屋のページをレンダリング
func  (m *Repository) Generals(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

// Majors 部屋のページをレンダリング
func  (m *Repository) Majors(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}

// Availability 検索の作成-可用性ページを表示
func  (m *Repository) Availability(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "search-avaliability.page.html", &models.TemplateData{})
}

// PostAvailability 検索の作成-可用性ページを表示
func  (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request){
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

// Contact コンタクトページをレンダリングします
func  (m *Repository) Contact(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}
