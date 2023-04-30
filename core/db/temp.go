package db

import (
	"encoding/json"
	"fmt"
	"time"
)

var postsMemDb = []map[string]any{
	{
		"id":       12,
		"title":    "My first post",
		"abstract": "Delectus nesciunt magni consequuntur corporis voluptatem deserunt adipisci dolor. Consectetur aliquam culpa consequatur corporis occaecati impedit possimus quis. Inventore esse debitis quia ratione sunt. Aut pariatur eos ea iure rerum. Dolores explicabo adipisci sit sequi.",
		"body": `Aut voluptates optio maxime cupiditate aut et et. Voluptatum officia itaque. Consequatur modi voluptas et et sequi.
 
Minima quisquam tempore a ut ut at ipsum. Earum omnis voluptatem fuga hic. Cumque modi aut aliquid. Eos est non. Et in aut voluptatem voluptatem et voluptatum. Rem consequatur possimus molestiae at error cupiditate nisi corporis.
 
Dolorem enim voluptatem. Officiis consequatur quaerat sint consequatur. Aut et reprehenderit dolor. Occaecati cupiditate perferendis omnis rem repudiandae placeat.`,
		"created_at":   time.Date(2023, 4, 28, 18, 12, 0, 0, time.Local),
		"published_at": time.Date(2023, 4, 28, 18, 32, 0, 0, time.Local),
		"edited_at":    nil,
	},
	{
		"id":       28,
		"title":    "My Second post",
		"abstract": "Delectus nesciunt magni consequuntur corporis voluptatem deserunt adipisci dolor. Consectetur aliquam culpa consequatur corporis occaecati impedit possimus quis. Inventore esse debitis quia ratione sunt. Aut pariatur eos ea iure rerum. Dolores explicabo adipisci sit sequi.",
		"body": `Aut voluptates optio maxime cupiditate aut et et. Voluptatum officia itaque. Consequatur modi voluptas et et sequi.
 
Minima quisquam tempore a ut ut at ipsum. Earum omnis voluptatem fuga hic. Cumque modi aut aliquid. Eos est non. Et in aut voluptatem voluptatem et voluptatum. Rem consequatur possimus molestiae at error cupiditate nisi corporis.
 
Dolorem enim voluptatem. Officiis consequatur quaerat sint consequatur. Aut et reprehenderit dolor. Occaecati cupiditate perferendis omnis rem repudiandae placeat.`,
		"created_at":   time.Date(2023, 4, 28, 18, 12, 0, 0, time.Local),
		"published_at": time.Date(2023, 4, 28, 18, 32, 0, 0, time.Local),
		"edited_at":    nil,
	},
}

func GetPosts() ([]Post, error) {
	posts := make([]Post, 0)

	for _, p := range postsMemDb {
		data, err := json.Marshal(p)
		if err != nil {
			return posts, err
		}
		var post Post
		err = json.Unmarshal(data, &post)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostById(id int) (Post, error) {
	var post Post
	for _, p := range postsMemDb {
		if p["id"] == id {
			data, err := json.Marshal(p)
			if err != nil {
				return post, err
			}
			var post Post
			err = json.Unmarshal(data, &post)
			if err != nil {
				return post, err
			}
			return post, nil
		}
	}
	return post, fmt.Errorf("now found")
}
