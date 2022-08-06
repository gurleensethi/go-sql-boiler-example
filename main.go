package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	dbmodels "github.com/gurleensethi/go-sql-boiler-example/db/models"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	ctx := context.Background()
	db := connectDB()

	boil.SetDB(db)

	author := createAuthor(ctx)
	createArticle(ctx, author)
	createArticle(ctx, author)
	selectAuthorWithArticle(ctx, author.ID)
}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:2345/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func createAuthor(ctx context.Context) dbmodels.Author {
	author := dbmodels.Author{
		Name:  "John Doe",
		Email: "johndoe@email.com",
	}

	err := author.InsertG(ctx, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	return author
}

func createArticle(ctx context.Context, author dbmodels.Author) dbmodels.Article {
	article := dbmodels.Article{
		Title:    "Hello World",
		Body:     null.StringFrom("Hello world, this is an article."),
		AuthorID: author.ID,
	}

	err := article.InsertG(ctx, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	return article
}

func selectAuthorWithArticle(ctx context.Context, authorID int) {
	author, err := dbmodels.Authors(dbmodels.AuthorWhere.ID.EQ(authorID)).OneG(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Author: \n\tID:%d \n\tName:%s \n\tEmail:%s\n", author.ID, author.Name, author.Email)

	articles, err := author.Articles().AllG(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range articles {
		fmt.Printf("Article: \n\tID:%d \n\tTitle:%s \n\tBody:%s \n\tCreatedAt:%v\n", a.ID, a.Title, a.Body.String, a.CreatedAt.Time)
	}
}

func selectAuthorWithArticleEager(ctx context.Context, authorID int) {
	author, err := dbmodels.Authors(dbmodels.AuthorWhere.ID.EQ(authorID), qm.Load(dbmodels.AuthorRels.Articles)).OneG(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Author: \n\tID:%d \n\tName:%s \n\tEmail:%s\n", author.ID, author.Name, author.Email)

	for _, a := range author.R.Articles {
		fmt.Printf("Article: \n\tID:%d \n\tTitle:%s \n\tBody:%s \n\tCreatedAt:%v\n", a.ID, a.Title, a.Body.String, a.CreatedAt.Time)
	}
}

func selectAuthorWithArticleJoin(ctx context.Context, authorID int) {
	type AuthorAndArticle struct {
		Article dbmodels.Article `boil:"articles,bind"`
		Author  dbmodels.Author  `boil:"author,bind"`
	}

	authorAndArticles := make([]AuthorAndArticle, 0)

	fmt.Sprintf("%s on %s = %s", dbmodels.TableNames.Article, dbmodels.ArticleTableColumns.ID, dbmodels.AuthorTableColumns.ID)

	err := dbmodels.NewQuery(
		qm.Select("*"),
		qm.From(dbmodels.TableNames.Author),
		qm.InnerJoin("articles on articles.author_id = author.id"),
		dbmodels.AuthorWhere.ID.EQ(authorID),
	).BindG(ctx, &authorAndArticles)
	if err != nil {
		log.Fatal(err)
	}

	for _, authorAndArticle := range authorAndArticles {
		author := authorAndArticle.Author
		a := authorAndArticle.Article

		fmt.Printf("Author: \n\tID:%d \n\tName:%s \n\tEmail:%s\n", author.ID, author.Name, author.Email)
		fmt.Printf("Article: \n\tID:%d \n\tTitle:%s \n\tBody:%s \n\tCreatedAt:%v\n", a.ID, a.Title, a.Body.String, a.CreatedAt.Time)
	}
}
