package repo

import "university/model"

// Making more readable response object for the users to consume
func ToResponse(data *model.Club) *Response {
	return &Response{
		ID:         data.ID,
		Name:       data.Name,
		Department: data.Department,
		President:  data.President,
		Secretary:  data.Secretary,
	}
}

func ToListResponse(data []model.Club) []Response {
	res := []Response{}
	for _, d := range data {
		res = append(res, *ToResponse(&d))
	}
	return res
}
