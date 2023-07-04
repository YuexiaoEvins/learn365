package reflect

//A customized json codec

type JsonUnit struct {
	Menu struct {
		Id    string `json:"id"`
		Value string `json:"value"`
		Popup struct {
			Menuitem []struct {
				Value   string `json:"value"`
				Onclick string `json:"onclick"`
			} `json:"menuitem"`
		} `json:"popup"`
	} `json:"menu"`
}

// CustomMarshal todo complete me
func CustomMarshal(t interface{}) (data []byte, err error) {
	//implement json.Marshal
	return nil, nil
}

// CustomUnmarshal todo complete me
func CustomUnmarshal(data []byte, dst interface{}) (err error) {
	//implement json.Marshal
	return nil
}
