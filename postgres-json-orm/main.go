// Package main is a toy program that demonstrates how one can leverage
// PostgreSQL's and Go's JSON features to simplify mapping SQL data to Go
// structs without the need for a 3rd-party ORM library.
package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/kylelemons/godebug/pretty"
	_ "github.com/lib/pq"
)

const (
	// schema defines a simple toy schema of blog `posts` that have a 1:N
	// relationship to the `comments` table.
	schema = `
CREATE TEMPORARY TABLE posts (
	post_id serial PRIMARY KEY,
	title text
) ON COMMIT DROP;

CREATE TEMPORARY TABLE comments (
	comment_id serial,
	post_id int REFERENCES posts(post_id),
	body text
) ON COMMIT DROP;
`
	// data populates `posts` and `comments` with some sample data. You may tweak
	// the generate_series() expressions to test with large data sets to explore
	// the performance, but that's not the main concern of this demo.
	data = `
INSERT INTO posts
SELECT
	i AS post_id,
	'Post '||i AS title
FROM generate_series(1, 3) i;

INSERT INTO comments
SELECT
	i AS comment_id,
	i % (SELECT count(*) FROM posts)+1 AS post_id,
	'Comment '||i AS body
FROM generate_series(1, 10) i;
`
)

// Posts is the main type we want to populate with our SQL data.
type Posts []*Post

// Post is a single post that embedds the comments it has.
type Post struct {
	PostID   int
	Title    string
	Comments []*Comment
}

// Comment is a comment that belongs to a Post.
type Comment struct {
	CommentID int
	PostID    int
	Body      string
}

// postsStdlib uses the traditional stdlib based approach of mapping SQL data
// to Go structs. Note how verbose it is, and duplicating the column names (and
// ordering) in both SQL and Go. It's also awkward to have to return each post
// multiple times in the results due to how JOINs work. And how sure are you
// that there isn't a panic-bug hidden somewhere?
func postsStdlib(tx *sql.Tx) (Posts, error) {
	rows, err := tx.Query(`
SELECT
	posts.post_id,
	posts.title,
	comments.comment_id,
	comments.post_id,
	comments.body
FROM posts
JOIN comments USING (post_id)
ORDER BY posts.post_id, comments.comment_id
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts Posts
	for rows.Next() {
		post := &Post{}
		comment := &Comment{}
		err := rows.Scan(
			&post.PostID,
			&post.Title,
			&comment.CommentID,
			&comment.PostID,
			&comment.Body,
		)
		if err != nil {
			return nil, err
		}

		if len(posts) == 0 || posts[len(posts)-1].PostID != post.PostID {
			posts = append(posts, post)
		}

		currentPost := posts[len(posts)-1]
		currentPost.Comments = append(currentPost.Comments, comment)
	}

	return posts, rows.Err()
}

// postsJSON uses PostgreSQL's built-in JSON functions to return a JSON array
// which can be easily unmarshaled using json.Unmarshal. No code duplication,
// and complicated Go logic.
func postsJSON(tx *sql.Tx) (Posts, error) {
	row := tx.QueryRow(`
SELECT coalesce(json_agg(json_build_object(
	'PostID', posts.post_id,
	'Title', posts.title,
	'Comments', (
		SELECT json_agg(json_build_object(
			'CommentID', comments.comment_id,
			'PostID', comments.post_id,
			'Body', comments.body
		) ORDER BY comment_id)
		FROM comments
		WHERE comments.post_id = posts.post_id
	)
) ORDER BY post_id), '[]')
FROM posts
`)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, err
	}
	var posts Posts
	return posts, json.Unmarshal(data, &posts)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	// credentials are controlled by PG* env variables
	db, err := sql.Open("postgres", "")
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := insertSampleData(tx); err != nil {
		return err
	}

	a, err := postsStdlib(tx)
	if err != nil {
		return err
	}

	b, err := postsJSON(tx)
	if err != nil {
		return err
	}

	fmt.Printf("%d %d\n", len(a), len(b))
	if !reflect.DeepEqual(a, b) {
		return errors.New(pretty.Compare(a, b))
	}

	return tx.Commit()
}

func insertSampleData(tx *sql.Tx) error {
	_, err := tx.Exec(schema + data)
	return err
}

func (p Posts) String() string {
	data, _ := json.MarshalIndent(p, "", "  ")
	return string(data)
}
