package tw

import (
	"context"
	"errors"
	"fmt"
	twitter "github.com/g8rswimmer/go-twitter/v2"
	"github.com/yushengguo557/twitter-space/models"
	"github.com/yushengguo557/twitter-space/utils"
	"log"
	"strings"
	"time"
)

func (tc *TwitterClient) SpaceUser(id string) (users []models.TwitterUser, err error) {
	if !utils.HasValue(id) {
		err = fmt.Errorf("id = %s", id)
		return nil, err
	}
	opts := twitter.SpacesLookupOpts{
		Expansions: []twitter.Expansion{
			twitter.ExpansionInvitedUserIDs,
			twitter.ExpansionSpeakerIDS,
			twitter.ExpansionCreatorID,
			twitter.ExpansionHostIDs,
		},
		UserFields: []twitter.UserField{
			twitter.UserFieldLocation,
			twitter.UserFieldDescription,
			twitter.UserFieldProfileImageURL,
			twitter.UserFieldName,
			twitter.UserFieldUserName,
			twitter.UserFieldID,
		},
	}

	log.Println("Callout to spaces search callout")

	// id列表长度 = 1 访问的API: Spaces lookup by single ID
	// 100 >= id列表长度 > 1 访问的API: Spaces lookup by list of IDs
	var spaceResponse *twitter.SpacesLookupResponse
	spaceResponse, err = tc.SpacesLookup(context.Background(), []string{id}, opts)
	if err != nil {
		log.Printf("spaces search, err: %v\n", err)
	}
	if spaceResponse.Raw == nil {
		return nil, errors.New("no data")
	}
	if spaceResponse.Raw.Includes == nil {
		return nil, errors.New("no data")
	}
	userObjSlice := spaceResponse.Raw.Includes.Users
	if len(userObjSlice) == 0 {
		return nil, errors.New("len(userObjSlice) == 0")
	}
	for _, userObj := range userObjSlice {
		var createdAt time.Time
		if utils.HasValue(userObj.CreatedAt) {
			createdAt, err = time.Parse(time.RFC3339Nano, userObj.CreatedAt)
			if err != nil {
				return nil, err
			}
		}
		user := models.TwitterUser{
			ID:              userObj.ID,
			Name:            userObj.Name,
			Username:        userObj.UserName,
			Location:        userObj.Location,
			Description:     userObj.Description,
			ProfileImageUrl: userObj.ProfileImageURL,
			SpaceId:         id,
			CreatedAt:       createdAt,
			Url:             strings.Join([]string{"https://twitter.com/", userObj.UserName}, ""),
		}
		users = append(users, user)
	}
	return users, nil
}
