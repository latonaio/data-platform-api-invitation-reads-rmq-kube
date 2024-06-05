package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-invitation-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-invitation-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var validationCheck *[]dpfm_api_output_formatter.ValidationCheck

	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "Headers":
			func() {
				header = c.Headers(mtx, input, output, errs, log)
			}()
		case "HeadersByInvitationOwner":
			func() {
				header = c.HeadersByInvitationOwner(mtx, input, output, errs, log)
			}()
		case "HeadersByInvitationGuest":
			func() {
				header = c.HeadersByInvitationGuest(mtx, input, output, errs, log)
			}()
		case "ValidationCheck":
			func() {
				validationCheck = c.ValidationCheck(mtx, input, output, errs, log)
			}()
		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:  			header,
		ValidationCheck:	validationCheck,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.Invitation = %d", input.Header.Invitation)

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v", where, *input.Header.IsCancelled)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invitation_header_data AS header
		` + where + ` ORDER BY header.IsCancelled ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Headers(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v", where, *input.Header.IsCancelled)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invitation_header_data AS header
		` + where + ` ORDER BY header.IsCancelled ASC, header.Invitation ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByInvitationOwner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.InvitationOwner = %d", input.Header.InvitationOwner)

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v", where, *input.Header.IsCancelled)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invitation_header_data AS header
		` + where + ` ORDER BY header.IsCancelled ASC, header.Invitation ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByInvitationGuest(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.InvitationGuest = %d", input.Header.InvitationGuest)

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v", where, *input.Header.IsCancelled)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invitation_header_data AS header
		` + where + ` ORDER BY header.IsCancelled ASC, header.Invitation ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ValidationCheck(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ValidationCheck {
	where := fmt.Sprintf("WHERE header.InvitationObjectType = \"%s\"", input.Header.InvitationObjectType)

	where = fmt.Sprintf("%s\nAND header.InvitationObject = %d", where, input.Header.InvitationObject)
	
	where = fmt.Sprintf("%s\nAND header.InvitationOwner = %d", where, input.Header.InvitationOwner)

	where = fmt.Sprintf("%s\nAND header.InvitationGuest = %d", where, input.Header.InvitationGuest)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invitation_header_data AS header
		` + where + ` ORDER BY header.IsCancelled ASC, header.Invitation ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToValidationCheck(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
