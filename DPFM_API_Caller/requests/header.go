package requests

type Header struct {
	Invitation				int		`json:"Invitation"`
	InvitationType			string	`json:"InvitationType"`
	InvitationOwner			int		`json:"InvitationOwner"`
	InvitationGuest			int		`json:"InvitationGuest"`
	InvitationObjectType	string	`json:"InvitationObjectType"`
	InvitationObject		int		`json:"InvitationObject"`
	OwnerParticipation		int		`json:"OwnerParticipation"`
	OwnerAttendance			*int	`json:"OwnerAttendance"`
	GuestParticipation		*int	`json:"GuestParticipation"`
	GuestAttendance			*int	`json:"GuestAttendance"`
	ValidityStartDate		string	`json:"ValidityStartDate"`
	ValidityEndDate			string	`json:"ValidityEndDate"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	IsCancelled				*bool	`json:"IsCancelled"`
}
