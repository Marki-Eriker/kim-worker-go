package request

import (
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func MapOneRequestToGqlModel(r *Request) *model.Request {
	return &model.Request{
		ID:                       r.ID,
		ServiceTypeID:            r.ServiceTypeID,
		ContractorID:             r.ContractorID,
		OrganizationContactID:    &r.OrganizationContactID,
		ContractMediumType:       r.ContractMediumType,
		ContractFilledTemplateID: &r.ContractFilledTemplateID,
		Status:                   r.Status,
		CreatedAt:                r.CreatedAt,
		BankAccountID:            &r.BankAccountID,
		SignatoryID:              &r.SignatoryID,
	}
}

func MapManyRequestToGqlModels(r []*Request) []*model.Request {
	items := make([]*model.Request, len(r))
	for i, v := range r {
		items[i] = MapOneRequestToGqlModel(v)
	}

	return items
}

func MapOneServiceTypeToGqlModel(st *ServiceType) *model.ServiceType {
	return &model.ServiceType{
		ID:   st.ID,
		Name: st.Name,
	}
}

func MapOneContractorToGqlModel(c *Contractor) *model.Contractor {
	return &model.Contractor{
		ID:             c.ID,
		FillName:       c.FullName,
		ShortName:      &c.ShortName,
		ContractorType: c.ContractorType,
		PersonID:       &c.PersonID,
	}
}

func MapOneOrganizationContactToGqlModel(c *OrganizationContact) *model.OrganizationContact {
	return &model.OrganizationContact{
		ID:    c.ID,
		Phone: &c.Phone,
		Email: &c.Email,
	}
}

func MapOneBankAccountToGqlModel(a *BankAccount) *model.BankAccount {
	return &model.BankAccount{
		ID:                  a.ID,
		AccountNumber:       a.AccountNumber,
		CorrespondentNumber: a.CorrespondentNumber,
		Bik:                 a.Bik,
		BankName:            a.BankName,
	}
}

func MapOneSignatoryToGqlModel(s *Signatory) *model.Signatory {
	return &model.Signatory{
		ID:            s.ID,
		Name:          &s.Name,
		ActingBasis:   &s.ActingBasis,
		WarrantNumber: &s.WarrantNumber,
		WarrantDate:   &s.WarrantDate,
	}
}

func MapOneShipToGqlModel(s *Ship) *model.Ship {
	return &model.Ship{
		ID:                              s.ID,
		Name:                            s.Name,
		HullNumber:                      &s.HullNumber,
		ProjectNumber:                   &s.ProjectNumber,
		Length:                          &s.Length,
		Width:                           &s.Width,
		HullHeight:                      &s.HullHeight,
		Cubic:                           &s.Cubic,
		Flag:                            &s.Flag,
		ShipConfirmParamsCertificateIds: s.ShipConfirmParamsCertificateIds,
		OwnerShipRightsCertificateIds:   s.OwnerShipRightsCertificateIds,
	}
}

func MapManyShipToGqlModel(s []*Ship) []*model.Ship {
	items := make([]*model.Ship, len(s))
	for i, v := range s {
		items[i] = MapOneShipToGqlModel(v)
	}

	return items
}

func MapOnePersonToGqlModel(p *Person) *model.Person {
	return &model.Person{
		ID:    p.ID,
		Email: &p.Email,
		Phone: &p.Phone,
	}
}
