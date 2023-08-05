package tw

import (
	"context"
	"errors"
	"fmt"
	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/yushengguo557/twitter-space/models"
	"log"
)

// GetUserInfoByID 通过ID获取用户信息
func (tc *TwitterClient) GetUserInfoByID(id string) (user *models.TwitterUser, err error) {
	opts := twitter.UserLookupOpts{
		Expansions: []twitter.Expansion{twitter.ExpansionPinnedTweetID},
		UserFields: []twitter.UserField{
			twitter.UserFieldID,
			twitter.UserFieldName,
			twitter.UserFieldUserName,
			twitter.UserFieldLocation,
			twitter.UserFieldDescription,
			twitter.UserFieldProfileImageURL,
			twitter.UserFieldURL,
		},
	}

	fmt.Println("Callout to user lookup callout")

	userResponse, err := tc.UserLookup(context.Background(), []string{id}, opts)
	if err != nil {
		log.Printf("user lookup error: %v\n", err)
	}
	if userResponse == nil {
		return nil, errors.New("no data")
	}
	if userResponse.Raw == nil {
		return nil, errors.New("no data")
	}
	if userResponse.Raw.Users == nil || len(userResponse.Raw.Users) < 1 {
		return nil, errors.New("no data")
	}
	userObj := userResponse.Raw.Users[0]
	return &models.TwitterUser{
		ID:              userObj.ID,
		Name:            userObj.Name,
		Username:        userObj.UserName,
		Location:        userObj.Location,
		Description:     userObj.Description,
		ProfileImageUrl: userObj.ProfileImageURL,
		Url:             "https://twitter.com/" + userObj.UserName,
	}, nil
}
