package params

func InitParams() Init {
	return InitManual()
}

func InitManual() Init {
	return Init{Expire: 1000, TokenKey: "test"} 
}
