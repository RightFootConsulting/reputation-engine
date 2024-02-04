package models

import (
	"time"

	"github.com/google/uuid"
)

type SocialEntityType string

const (
	SE_PERSON   SocialEntityType = "PERSON"
	SE_GROUP    SocialEntityType = "GROUP"
	SE_ORG      SocialEntityType = "ORG"
	SE_BUSINESS SocialEntityType = "BUSINESS"
)

// References a user, org or business on a social media platform
type SocialEntity struct {
	Id         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ExternalId string    `json:"external_id"` // If this is a USER this would be their User ID in the platform
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// KarmaEvents are postive karma awarded for activity on the social platform
// Examples of karma types may include things like:
//
//	ARTICLES, COMMENTS, LIKES, DOWNVOTES, POSTS, FALSE_ACCUSATION and so on.  It is open
//	ended.
//
// The event value is the postive or negative value of the event.
type KarmaEventType struct {
	Id            uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EventTypeName string    `json:"event_type_name"`
	EventValue    int       `json:"event_value"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type KarmaEvent struct {
	Id           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	SocialEntity uuid.UUID `json:"social_entity"  gorm:"type:uuid"`
	EventType    uuid.UUID `json:"karma_event_type" gorm:"karma_event_type"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type BlockReason string

const (
	BR_BORING                     BlockReason = "BORING"                     // Allows users to tune out content they find uninteresting
	BR_OFFENSIVE                  BlockReason = "OFFENSIVE"                  // Allows users to tune out users who they find are offensive
	BR_INNAPROPRIATE_SOLICITATION BlockReason = "INAPPROPRIATE_SOLICITATION" // Allows users to tune out inappropriate solicitations
	BR_MODERATOR_RULES            BlockReason = "MODERATOR_RULES_BREACH"     // Allows users to tune out
	BR_HARRASSMENT                BlockReason = "HARRASSMENT"                // Should be apealable since this is a serious accusation
	BR_FAKE                       BlockReason = "FAKE_PROFILE"               // Should be investigated for veracity
	BR_CRIMINALLY_FRAUDULENT      BlockReason = "CRIMINALLY_FRAUDULENT"      // Should be reported, investigated, and not factored in until confirmed
)

type VisibilityType string

type BlockEvent struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Blocker   uuid.UUID `json:"blocker" gorm:"type:uuid"`
	Blockee   uuid.UUID `json:"blockee" gorm:"type:uuid"`
	Confirmed bool      `json:"confirmed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}