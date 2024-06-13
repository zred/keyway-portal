package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/labstack/echo-contrib/session"
    "github.com/gorilla/sessions"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "net/http"
    "html/template"
    "io"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

    e.Static("/static", "static")

    tmpl := &Template{
        templates: template.Must(template.ParseGlob("templates/*.html")),
    }
    e.Renderer = tmpl

    // Database setup
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        e.Logger.Fatal(err)
    }
    defer db.Close()
    db.AutoMigrate(&User{})

    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "index.html", map[string]interface{}{
            "title": "Home Page",
        })
    })

    e.Logger.Fatal(e.Start(":1323"))
}

type User struct {
    gorm.Model
    Name  string
    Email string
    Password string
}

