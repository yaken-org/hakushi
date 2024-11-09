package service

import (
	"github.com/yaken-org/hakushi/internal/database"
	"github.com/yaken-org/hakushi/internal/model"
)

// FindAllPosts は全ての Post を取得します。
func FindAllPosts() ([]*model.Post, error) {
	db := database.New()
	var posts []*model.Post

	result := db.Gorm.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

// FindPostsOrderByLikes はいいねの数が多い順に Post を取得します。
// いいねの数が 0 の Post は取得されません。
func FindPostsOrderByLikes() ([]*model.Post, error) {
	db := database.New()

	var posts []*model.Post
	result := db.Gorm.Where("likes > 0").Order("likes desc").Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

// FindPostByID は ID から Post を取得します。
// Post が見つからなかった場合は nil を返します。
func FindPostByID(id int64) (*model.Post, error) {
	db := database.New()

	post := new(model.Post)
	result := db.Gorm.Find(post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return post, nil
}

func FindAPIPostByID(id int64) (*model.APIPost, error) {
	post, err := FindPostByID(id)
	if err != nil {
		return nil, err
	}

	annotations, err := FindAnnotationsByPostID(id)
	if err != nil {
		return nil, err
	}

	tags, err := FindPostRelatedTags(post.ID)
	if err != nil {
		return nil, err
	}

	return post.ToAPIPost(annotations, tags), nil
}

// FindPostByNameRough はタイトルに name を含む Post を取得します。
// name は部分一致で検索されます。
func FindPostByNameRough(name string) ([]*model.Post, error) {
	db := database.New()

	name = "%" + name + "%"
	var posts []*model.Post
	result := db.Gorm.Where("title LIKE ?", name).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

// FindPostsByUserAccountID は UserAccountID から Post を取得します。
// Post は作成日時の降順で取得されます。
func FindPostsByUserAccountID(userAccountID int64) ([]*model.Post, error) {
	db := database.New()

	var posts []*model.Post
	result := db.Gorm.
		Where("user_account_id = ?", userAccountID).
		Order("created_at desc").
		Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

// CreatePost は Post を作成します。
func CreatePost(userAccount model.UserAccount, imageId int64, title string, content string) (*model.Post, error) {
	db := database.New()

	post := &model.Post{
		UserAccountID: userAccount.ID,
		ImageID:       imageId,
		Title:         title,
		Content:       content,
	}

	result := db.Gorm.Create(post)
	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}

func DeletePostById(id int64) error {
	db := database.New()

	result := db.Gorm.Delete(&model.Post{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePost(id int64, title string, content string) (*model.Post, error) {
	db := database.New()

	post := new(model.Post)
	result := db.Gorm.Model(post).Where("id = ?", id).Updates(map[string]interface{}{
		"title":   title,
		"content": content,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return FindPostByID(id)
}

// FindPostRelatedTags は Post に関連する Tag を取得します。
func FindPostRelatedTags(postID int64) ([]*model.Tag, error) {
	db := database.New()

	// PostID に関連する TagID を取得
	var postTags []*model.PostTag
	if result := db.Gorm.Where("post_id = ?", postID).Find(&postTags); result.Error != nil {
		return nil, result.Error
	}
	var tagIDs []int64
	for _, postTag := range postTags {
		tagIDs = append(tagIDs, postTag.TagID)
	}

	// TagID から Tag を取得
	var tags []*model.Tag
	if result := db.Gorm.Where("id IN (?)", tagIDs).Find(&tags); result.Error != nil {
		return nil, result.Error
	}

	return tags, nil
}

// FindPostRelatedTagsByPostIDs は Post に関連する Tag を取得します。
// Post to Tags のマップを返します。
func FindPostRelatedTagsByPostIDs(postIDs []int64) (map[int64][]*model.Tag, error) {
	db := database.New()

	// 複数の PostID から PostTag を取得
	var postTags []*model.PostTag
	if result := db.Gorm.Where("post_id IN (?)", postIDs).Find(&postTags); result.Error != nil {
		return nil, result.Error
	}

	// PostTag から TagID だけのスライスを作成
	var tagIDs []int64
	for _, postTag := range postTags {
		tagIDs = append(tagIDs, postTag.TagID)
	}
	// TagID から Tag を取得
	var tags []*model.Tag
	if result := db.Gorm.Where("id IN (?)", tagIDs).Find(&tags); result.Error != nil {
		return nil, result.Error
	}
	// TagID to Tag のマップを作成
	tagIDToTag := make(map[int64]*model.Tag)
	for _, tag := range tags {
		tagIDToTag[tag.ID] = tag
	}

	// PostID to Tags のマップを作成
	postIDToTags := make(map[int64][]*model.Tag)
	for _, postTag := range postTags {
		tag, ok := tagIDToTag[postTag.TagID]
		if !ok {
			continue
		}
		postIDToTags[postTag.PostID] = append(postIDToTags[postTag.PostID], tag)
	}

	return postIDToTags, nil
}

// FindPostsByTag は Tag から Post を取得します。
func FindPostsByTag(tag *model.Tag) ([]*model.Post, error) {
	db := database.New()

	var postTags []*model.PostTag
	if result := db.Gorm.Where("tag_id = ?", tag.ID).Find(&postTags); result.Error != nil {
		return nil, result.Error
	}

	var postIDs []int64
	for _, postTag := range postTags {
		postIDs = append(postIDs, postTag.PostID)
	}

	var posts []*model.Post
	if result := db.Gorm.Where("id IN (?)", postIDs).Find(&posts); result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

// IncrementPostLikeCount は Post のいいね数をインクリメントします。
func IncrementPostLikeCount(postID int64) (int, error) {
	db := database.New()
	post := new(model.Post)

	result := db.Gorm.Model(post).Where("id = ?", postID).Update("likes", post.Likes+1)
	if result.Error != nil {
		return 0, result.Error
	}

	post, err := FindPostByID(postID)
	if err != nil {
		return 0, err
	}

	return post.Likes, nil
}
