package ControllersPublisher

type ExchangeName string
type BindingKey string
type TypeUpdateStatus string

type DataPublish struct {
	DataRequest  []byte       `json:"data_request"`
	ExchangeName ExchangeName `json:"exchange_name"`
	Key          BindingKey   `json:"key"`
}
