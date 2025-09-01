package accesses

import (
	"io"
	"fmt"
	"strconv"
	"bytes"
)

// Custom Accesses
type Accesses string

const (
	AccessesViewProfile                 Accesses = "VIEW_PROFILE"
	AccessesViewProfileBu               Accesses = "VIEW_PROFILE_BU"
	AccessesEditProfile                 Accesses = "EDIT_PROFILE"
	AccessesEditProfileBu               Accesses = "EDIT_PROFILE_BU"
	AccessesCreateProfile               Accesses = "CREATE_PROFILE"
	AccessesViewTeam                    Accesses = "VIEW_TEAM"
	AccessesViewTeamBu                  Accesses = "VIEW_TEAM_BU"
	AccessesViewTeamList                Accesses = "VIEW_TEAM_LIST"
	AccessesCreateTeam                  Accesses = "CREATE_TEAM"
	AccessesEditTeam                    Accesses = "EDIT_TEAM"
	AccessesDeleteTeam                  Accesses = "DELETE_TEAM"
	AccessesViewProfileList             Accesses = "VIEW_PROFILE_LIST"
	AccessesViewProfileListBu           Accesses = "VIEW_PROFILE_LIST_BU"
	AccessesExportFilteredProfileList   Accesses = "EXPORT_FILTERED_PROFILE_LIST"
	AccessesExportFilteredProfileListBu Accesses = "EXPORT_FILTERED_PROFILE_LIST_BU"
	AccessesExportTeamMembers           Accesses = "EXPORT_TEAM_MEMBERS"
	AccessesMigrateProfiles             Accesses = "MIGRATE_PROFILES"
	AccessesCreateAccessRequests        Accesses = "CREATE_ACCESS_REQUESTS"
	AccessesUpdateAccessRequests        Accesses = "UPDATE_ACCESS_REQUESTS"
	AccessesDeleteAccessRequests        Accesses = "DELETE_ACCESS_REQUESTS"
	AccessesViewAccessRequests          Accesses = "VIEW_ACCESS_REQUESTS"
	AccessesViewOrganizationTree        Accesses = "VIEW_ORGANIZATION_TREE"
	AccessesManageEntities              Accesses = "MANAGE_ENTITIES"
)

var AllAccesses = []Accesses{
	AccessesViewProfile,
	AccessesViewProfileBu,
	AccessesEditProfile,
	AccessesEditProfileBu,
	AccessesCreateProfile,
	AccessesViewTeam,
	AccessesViewTeamBu,
	AccessesViewTeamList,
	AccessesCreateTeam,
	AccessesEditTeam,
	AccessesDeleteTeam,
	AccessesViewProfileList,
	AccessesViewProfileListBu,
	AccessesExportFilteredProfileList,
	AccessesExportFilteredProfileListBu,
	AccessesExportTeamMembers,
	AccessesMigrateProfiles,
	AccessesCreateAccessRequests,
	AccessesUpdateAccessRequests,
	AccessesDeleteAccessRequests,
	AccessesViewAccessRequests,
	AccessesViewOrganizationTree,
	AccessesManageEntities,
}

func (e Accesses) IsValid() bool {
	switch e {
	case AccessesViewProfile, AccessesViewProfileBu, AccessesEditProfile, AccessesEditProfileBu, AccessesCreateProfile, AccessesViewTeam, AccessesViewTeamBu, AccessesViewTeamList, AccessesCreateTeam, AccessesEditTeam, AccessesDeleteTeam, AccessesViewProfileList, AccessesViewProfileListBu, AccessesExportFilteredProfileList, AccessesExportFilteredProfileListBu, AccessesExportTeamMembers, AccessesMigrateProfiles, AccessesCreateAccessRequests, AccessesUpdateAccessRequests, AccessesDeleteAccessRequests, AccessesViewAccessRequests, AccessesViewOrganizationTree, AccessesManageEntities:
		return true
	}
	return false
}

func (e Accesses) String() string {
	return string(e)
}

func (e *Accesses) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Accesses(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Accesses", str)
	}
	return nil
}

func (e Accesses) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *Accesses) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return e.UnmarshalGQL(s)
}

func (e Accesses) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	e.MarshalGQL(&buf)
	return buf.Bytes(), nil
}