package basemodel

type Sort struct {
	By    string `json:"by,omitempty"`
	Order string `json:"order,omitempty"`
}

func (m *Sort) FromProto(pb *Sort) {
	m.By = pb.By
	m.Order = pb.Order
}

func (m *Sort) ToProto() *Sort {
	return &Sort{
		By:    m.By,
		Order: m.Order,
	}
}
