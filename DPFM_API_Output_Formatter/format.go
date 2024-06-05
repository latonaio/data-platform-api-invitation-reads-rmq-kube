package dpfm_api_output_formatter

import (
	"data-platform-api-invitation-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.Invitation,
			&pm.InvitationType,
			&pm.InvitationOwner,
			&pm.InvitationGuest,
			&pm.InvitationObjectType,
			&pm.InvitationObject,
			&pm.OwnerParticipation,
			&pm.OwnerAttendance,
			&pm.GuestParticipation,
			&pm.GuestAttendance,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			Invitation:				data.Invitation,
			InvitationType:			data.InvitationType,
			InvitationOwner:		data.InvitationOwner,
			InvitationGuest:		data.InvitationGuest,
			InvitationObjectType:	data.InvitationObjectType,
			InvitationObject:		data.InvitationObject,
			OwnerParticipation:		data.OwnerParticipation,
			OwnerAttendance:		data.OwnerAttendance,
			GuestParticipation:		data.GuestParticipation,
			GuestAttendance:		data.GuestAttendance,
			ValidityStartDate:		data.ValidityStartDate,
			ValidityEndDate:		data.ValidityEndDate,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			IsCancelled:			data.IsCancelled,

		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToValidationCheck(rows *sql.Rows) (*[]ValidationCheck, error) {
	defer rows.Close()
	validationCheck := make([]ValidationCheck, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ValidationCheck{}

		err := rows.Scan(
			&pm.InvitationObjectType,
			&pm.InvitationObject,
			&pm.InvitationOwner,
			&pm.InvitationGuest,
			&pm.Invitation,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &validationCheck, err
		}

		data := pm
		validationCheck = append(validationCheck, ValidationCheck{
			InvitationObjectType:	data.InvitationObjectType,
			InvitationObject:		data.InvitationObject,
			InvitationOwner:		data.InvitationOwner,
			InvitationGuest:		data.InvitationGuest,
			Invitation:				data.Invitation,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			IsCancelled:			data.IsCancelled,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &validationCheck, nil
	}

	return &validationCheck, nil
}
