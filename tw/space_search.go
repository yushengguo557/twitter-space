package tw

import (
	"context"
	"errors"
	"fmt"
	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/yushengguo557/twitter-space/models"
	"github.com/yushengguo557/twitter-space/utils"
	"log"
	"strings"
	"time"
)

// SpaceSearch 空间搜索
func (tc *TwitterClient) SpaceSearch(query string) (spaces []models.TwitterSpace, err error) {
	log.Printf("query %s\n", query)
	opts := twitter.SpacesSearchOpts{
		SpaceFields: []twitter.SpaceField{
			twitter.SpaceFieldID,
			twitter.SpaceFieldLang,
			twitter.SpaceFieldParticipantCount,
			twitter.SpaceFieldEndedAt,
			twitter.SpaceFieldState,
			twitter.SpaceFieldTitle,
			twitter.SpaceFieldScheduledStart,
			twitter.SpaceFieldStartedAt,
		},
		Expansions: []twitter.Expansion{
			twitter.ExpansionCreatorID,
		},
		State: twitter.SpaceStateAll,
	}

	log.Println("Callout to spaces search callout")

	spaceResponse, err := tc.SpacesSearch(context.Background(), query, opts)
	if err != nil {
		err = fmt.Errorf("spaces search, err: %w", err)
		return nil, err
	}
	if spaceResponse.Raw == nil {
		return nil, errors.New("no data")
	}
	spaceObjSlice := spaceResponse.Raw.Spaces
	if len(spaceObjSlice) == 0 {
		return nil, errors.New("len(spaceObjSlice) = 0")
	}
	for _, spaceObj := range spaceObjSlice {
		if spaceObj == nil {
			continue
		}
		var startedAt time.Time
		if utils.HasValue(spaceObj.CreatedAt) {
			startedAt, err = time.Parse(time.RFC3339Nano, spaceObj.StartedAt)
			if err != nil {
				err = fmt.Errorf("parse time %s, err: %w", spaceObj.StartedAt, err)
				return nil, err
			}
		}
		var endedAt time.Time
		if utils.HasValue(spaceObj.EndedAt) {
			endedAt, err = time.Parse(time.RFC3339Nano, spaceObj.EndedAt)
			if err != nil {
				err = fmt.Errorf("parse time %s, err: %w", spaceObj.EndedAt, err)
				return nil, err
			}
		}
		var scheduledStart time.Time
		if utils.HasValue(spaceObj.ScheduledStart) {
			endedAt, err = time.Parse(time.RFC3339Nano, spaceObj.ScheduledStart)
			if err != nil {
				err = fmt.Errorf("parse time %s, err: %w", spaceObj.ScheduledStart, err)
				return nil, err
			}
		}
		space := models.TwitterSpace{
			ID:               spaceObj.ID,
			CreatorId:        spaceObj.CreatorID, // CreatorId 很大 不要转换为 int 类型存储
			ParticipantCount: spaceObj.ParticipantCount,
			Title:            spaceObj.Title,
			Description:      spaceObj.Title,
			Tag:              query,
			Lang:             spaceObj.Lang,
			Url:              strings.Join([]string{"https://twitter.com/i/spaces/", spaceObj.ID}, ""),
			Status:           spaceObj.State,
			StartedAt:        startedAt,
			ScheduledStart:   scheduledStart,
			EndedAt:          endedAt,
		}
		spaces = append(spaces, space)
	}
	return spaces, nil
}
