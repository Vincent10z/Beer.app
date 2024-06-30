// File: utils/id_generation.go

package utils

import (
	"fmt"
	"github.com/rs/xid"
)

// IDPrefix represents the prefix for different types of IDs
type IDPrefix string

const (
	UserPrefix           IDPrefix = "usr"
	BreweryPrefix        IDPrefix = "brw"
	BeerStylePrefix      IDPrefix = "sty"
	BeerPrefix           IDPrefix = "ber"
	BeerReviewPrefix     IDPrefix = "brv"
	BreweryReviewPrefix  IDPrefix = "wrv"
	CommentPrefix        IDPrefix = "cmt"
	RecommendationPrefix IDPrefix = "rec"
	FavoritePrefix       IDPrefix = "fav"
	CheckinPrefix        IDPrefix = "chk"
	EventPrefix          IDPrefix = "evn"
	EventAttendeePrefix  IDPrefix = "eat"
	FollowerPrefix       IDPrefix = "flw"
	UserActivityPrefix   IDPrefix = "act"
	BadgePrefix          IDPrefix = "bdg"
	UserBadgePrefix      IDPrefix = "ubg"
	BeerListPrefix       IDPrefix = "blt"
	BeerListItemPrefix   IDPrefix = "bli"
)

// GenerateID generates a new ID with the given prefix
func GenerateID(prefix IDPrefix) string {
	id := xid.New()
	return fmt.Sprintf("%s_%s", prefix, id.String()[:16])
}

// ID generation functions for each type
func GenerateUserID() string           { return GenerateID(UserPrefix) }
func GenerateBreweryID() string        { return GenerateID(BreweryPrefix) }
func GenerateBeerStyleID() string      { return GenerateID(BeerStylePrefix) }
func GenerateBeerID() string           { return GenerateID(BeerPrefix) }
func GenerateBeerReviewID() string     { return GenerateID(BeerReviewPrefix) }
func GenerateBreweryReviewID() string  { return GenerateID(BreweryReviewPrefix) }
func GenerateCommentID() string        { return GenerateID(CommentPrefix) }
func GenerateRecommendationID() string { return GenerateID(RecommendationPrefix) }
func GenerateFavoriteID() string       { return GenerateID(FavoritePrefix) }
func GenerateCheckinID() string        { return GenerateID(CheckinPrefix) }
func GenerateEventID() string          { return GenerateID(EventPrefix) }
func GenerateEventAttendeeID() string  { return GenerateID(EventAttendeePrefix) }
func GenerateFollowerID() string       { return GenerateID(FollowerPrefix) }
func GenerateUserActivityID() string   { return GenerateID(UserActivityPrefix) }
func GenerateBadgeID() string          { return GenerateID(BadgePrefix) }
func GenerateUserBadgeID() string      { return GenerateID(UserBadgePrefix) }
func GenerateBeerListID() string       { return GenerateID(BeerListPrefix) }
func GenerateBeerListItemID() string   { return GenerateID(BeerListItemPrefix) }
