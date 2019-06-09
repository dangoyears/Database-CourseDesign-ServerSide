package dbcd

type GateKeeper struct {
	Token2HumanID map[string]string
}

func (keeper *GateKeeper) addTokenForUser(token, userID string) {
	keeper.Token2HumanID[token] = userID
}

func (keeper *GateKeeper) deleteToken(token string) {
	delete(keeper.Token2HumanID, token)
}
