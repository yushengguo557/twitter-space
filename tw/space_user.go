package tw

import (
	"context"
	"fmt"
	"github.com/g8rswimmer/go-twitter/v2"
	"log"
	"strings"
	"time"
	"twitter-space/models"
	"twitter-space/utils"
)

func (tc *TwitterClient) SpaceUser(id string) (users []models.TwitterUser, err error) {
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

	fmt.Println("Callout to spaces search callout")

	spaceResponse, err := tc.SpacesLookup(context.Background(), []string{id}, opts)
	if err != nil {
		log.Printf("spaces search, err: %v\n", err)
	}

	userObjSlice := spaceResponse.Raw.Includes.Users
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
