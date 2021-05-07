package repo

import "university/model"

// Making more readable response object for the users to consume
func ToPresenter(data *model.Department) *Presenter {
	return &Presenter{
		ID:   data.ID,
		Name: data.Name,
	}
}

func ToPresenterList(data []model.Department) []Presenter {
	res := []Presenter{}
	for _, d := range data {
		res = append(res, *ToPresenter(&d))
	}
	return res
}
