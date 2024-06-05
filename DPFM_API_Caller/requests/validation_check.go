package requests

type ValidationCheck struct {
	InvitationObjectType	string	`json:"InvitationObjectType"`
	InvitationObject		int		`json:"InvitationObject"`
	InvitationOwner			int		`json:"InvitationOwner"`
	InvitationGuest			int		`json:"InvitationGuest"`
	Invitation				int		`json:"Invitation"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	IsCancelled				*bool	`json:"IsCancelled"`
}
