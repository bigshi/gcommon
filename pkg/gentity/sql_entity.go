package gentity

type QueryCondition struct {
	QueryKey   string
	QueryValue interface{}
}

const (
	OrderMapKeyOfCreateTime = "create_time"
	OrderMapKeyOfUpdateTime = "update_time"
	OrderMapKeyOfId         = "id"

	OrderMapValueOfAsc  = "asc"
	OrderMapValueOfDesc = "desc"

	QueryLimit = 300
)
