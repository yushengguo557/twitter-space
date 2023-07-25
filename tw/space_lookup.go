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

// SpaceLookup 用Space的ID获取Space信息
func (tc *TwitterClient) SpaceLookup(ids []string) (spaces []models.TwitterSpace, err error) {
	if len(ids) == 0 {
		return nil, errors.New("len(ids) == 0")
	}
	if !utils.HasValue(ids[0]) {
		return nil, fmt.Errorf("id = %s", ids[0])
	}
	opts := twitter.SpacesLookupOpts{
		SpaceFields: []twitter.SpaceField{
			twitter.SpaceFieldStartedAt,
			twitter.SpaceFieldID,
			twitter.SpaceFieldLang,
			twitter.SpaceFieldParticipantCount,
			twitter.SpaceFieldEndedAt,
			twitter.SpaceFieldState,
			twitter.SpaceFieldTitle,
			twitter.SpaceFieldScheduledStart,
		},
		Expansions: []twitter.Expansion{
			twitter.ExpansionCreatorID,
		},
	}

	log.Println("Callout to spaces search callout")

	spaceResponse, err := tc.SpacesLookup(context.Background(), ids, opts)
	if err != nil {
		err = fmt.Errorf("spaces search, err: %v", err)
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
