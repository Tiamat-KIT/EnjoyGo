package main

import "github.com/kataras/iris/v12"

type Book struct {
    Title string `json:"title"`
}

func main() {
    app := iris.New()

    // 「/books」から始まるパスに対する処理をグループ化する
    booksAPI := app.Party("/books")
    {
        // GET: http://localhost:8080/books
        booksAPI.Get("/", list)
        // POST: http://localhost:8080/books
        booksAPI.Post("/", create)
    }
	
    app.Listen(":8080")
}

func list(ctx iris.Context) {
    books := []Book{
        {"Mastering Concurrency in Go"},
        {"Go Design Patterns"},
        {"Black Hat Go"},
    }
    ctx.JSON(books)
}

func create(ctx iris.Context) {
    var b Book
    err := ctx.ReadJSON(&b)
    if err != nil {
        ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
            Title("Book creation failure").DetailErr(err))
        return
    }

    println("Received Book: " + b.Title)
    ctx.StatusCode(iris.StatusCreated)
}