package sql


func GetConfig(key string) (*ConfigModel, error) {

	data := ConfigModel{}
	err := SqlDB.Where(&ConfigModel{Key: key}).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func SetConfig(key string, value string) (*ConfigModel, error) {

	data := ConfigModel{}
	err := SqlDB.Where(ConfigModel{Key: key}).Assign(ConfigModel{Value: value}).FirstOrCreate(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
