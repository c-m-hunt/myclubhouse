package data

type Cost struct {
	ExclTax   float32
	Tax       float32
	TokenCost float32
}

type Organizer struct {
	ID               int
	MembershipNumber string
	Forename         string
	MiddleName       string
	Surname          string
	CompanyName      string
	Branch           string
}

type Event struct {
	ID   int
	Type struct {
		ID        int
		Name      string
		ShortName string
	}
	SeriesID    int
	IsDraft     bool
	Name        string
	Description string
	Organizer
	SecondOrganizer  Organizer
	OrganizingBranch struct {
		ID   int
		Code string
		Name string
	}
	OrganizingSection struct {
		ID   int
		Name string
	}
	VenueName         string
	VenueDirections   string
	VenueLatitude     float32
	VenueLongitude    float32
	StartTime         string
	EndTime           string
	SignUpOpeningTime string
	SignUpClosingTime string
	MinAttendees      int
	MaxAttendees      int
	AttendeeCount     int
	WaitingListCount  int
	HasCosts          bool
	CostStructure     struct {
		CostType          string
		CanPayWithMoney   bool
		CanPayWithTokens  bool
		FixedInstanceCost Cost
		FixedSeriesCost   Cost
		Classes           []struct {
			ID            int
			Ordinal       int
			Code          string
			Name          string
			Description   string
			StartDayIndex int
			DayCount      int
			MetaData      struct{}
			CostOptions   []struct {
				ID                  int
				Name                string
				Description         string
				AppliesToMembers    bool
				AppliesToNonMembers bool
				MinAge              int
				MaxAge              int
				MinAbilityLevel     struct {
					ID    int
					Name  string
					Level int
				}
				MaxAbilityLevel struct {
					ID    int
					Name  string
					Level int
				}
				InvitedMemberLists []struct {
					ID   int
					Name string
				}
				ExcludedMemberLists []struct {
					ID   int
					Name string
				}
				OpensDaysBefore  int
				ClosesDaysBefore int
				InstanceCost     Cost
				SeriesCost       Cost
			}
		}
	}
	AddOns []struct {
		ID       int
		Name     string
		UnitCost Cost
	}
	IsCancelled        bool
	CancellationReason string
	ViewURL            string
}
