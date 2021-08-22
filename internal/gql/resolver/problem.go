package resolver

import "github.com/marki-eriker/kim-worker-go/internal/gql/model"

func NewInternalErrorProblem() model.InternalErrorOutput {
	var errors []model.ProblemInterface
	errors = append(errors, model.InternalErrorProblem{Message: "internal server error"})

	return model.InternalErrorOutput{
		Ok:    false,
		Error: errors,
	}
}

func NewForbiddenErrorProblem() []model.ProblemInterface {
	var errors []model.ProblemInterface
	errors = append(errors, model.ForbiddenErrorProblem{Message: "forbidden"})

	return errors
}

func NewUnauthorizedErrorProblem() []model.ProblemInterface {
	var errors []model.ProblemInterface
	errors = append(errors, model.UnauthorizedErrorProblem{Message: "unauthorized"})

	return errors
}

func NewUnknownErrorProblem(err error) []model.ProblemInterface {
	var errors []model.ProblemInterface
	errors = append(errors, model.UnknowErrorProblem{Message: err.Error()})

	return errors
}

func NewValidationErrorProblem(err map[string]string) []model.ProblemInterface {
	var errors []model.ProblemInterface

	for k, v := range err {
		errors = append(errors, model.ValidationErrorProblem{Field: k, Message: v})
	}

	return errors
}
