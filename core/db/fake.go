package db

import (
	"encoding/json"
	"time"
	"unknspec/core/models"
)

type FakeDB struct {
	tags  []map[string]any
	posts []map[string]any
}

func NewFakeDB() *FakeDB {
	return &FakeDB{
		tags: []map[string]any{
			{"id": 1, "name": "test"},
			{"id": 2, "name": "test1"},
			{"id": 3, "name": "test2"},
			{"id": 4, "name": "test3"},
		},
		posts: []map[string]any{
			{
				"id":             1,
				"title":          "My first test blog post!!!",
				"abstract":       "Its my first test blog post, and i'm happy about it",
				"body":           "It's my first test blog post",
				"last_edited_at": time.Now(),
				"is_published":   false,
			},
			{
				"id":       2,
				"title":    "Markdown test",
				"abstract": "its just simple markdown document render test",
				"body": `
# H1
 
Lorem ipsum dolor sit amet, *consectetur* adipisicing elit, sed do eiusmod
tempor incididunt ut **labore et dolore magna aliqua**. Ut enim ad minim veniam,
 
 
> quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
 
consequat. ***Duis aute irure dolor*** in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. ~~Excepteur sint occaecat~~ cupidatat nonproident, sunt in culpa qui officia deserunt mollit anim id est laborum.
 
## H2
 
Lorem ipsum dolor sit amet, *consectetur* adipisicing elit, sed do eiusmod
tempor incididunt ut **labore et dolore magna aliqua**. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. 
 
---
 
***Duis aute irure dolor*** in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. ~~Excepteur sint occaecat~~ cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
 
### H3
 
unordered list:
 
* item-1
  * sub-item-1
  * sub-item-2
- item-2
  - sub-item-3
  - sub-item-4
+ item-3
  + sub-item-5
  + sub-item-6
 
 
ordered list:
 
1. item-1
   1. sub-item-1
   2. sub-item-2
2. item-2
   1. sub-item-3
   2. sub-item-4
3. item-3
 
#### Header4
 
Table Header-1 | Table Header-2 | Table Header-3
:--- | :---: | ---:
Table Data-1 | Table Data-2 | Table Data-3
TD-4 | Td-5 | TD-6
Table Data-7 | Table Data-8 | Table Data-9
 
##### Header5
 
You may also want some images right in here like ![GitHub Logo](https://cloud.githubusercontent.com/assets/5456665/13322882/e74f6626-dc00-11e5-921d-f6d024a01eaa.png "GitHub") - you can do that but I would recommend you to use the component "image" and simply split your text.
 
###### Header6
 
Let us do some links - this for example: https://github.com/MinhasKamal/github-markdown-syntax is **NOT** a link but this: is [GitHub](https://github.com/MinhasKamal/github-markdown-syntax)`,
				"last_edited_at": time.Now(),
				"is_published":   false,
			},
		},
	}
}

func (fdb *FakeDB) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	var err error
	for _, post := range fdb.posts {
		var p models.Post
		data, _ := json.Marshal(post)
		err = json.Unmarshal(data, &p)
		if err != nil {
			continue
		}
		posts = append(posts, p)
	}
	return posts, err
}
