package dbcd

import (
	"sync"
)

// GateKeeper 保存用户登陆时的凭证。
type GateKeeper struct {
	// Token2HumanID 保存token对应的自然人ID。
	Token2HumanID sync.Map

	// Token2Role 保存token对应的登陆角色。
	// 登陆角色为{"admin", "student", "teacher"}之一。
	Token2Role sync.Map
}

// NewGateKeeper 返回初始化的GateKeeper。
func NewGateKeeper() *GateKeeper {
	var keeper GateKeeper
	return &keeper
}

// RemoveToken 将传入的token从有效token记录中删除，
// 被token被删除后不能用于身份认证。
func (keeper *GateKeeper) RemoveToken(token string) {
	keeper.removeHumanIDToken(token)
	keeper.removeRoleToken(token)
}

// GetRole 返回token认证的角色。
// 当token无效时返回"anonymous"。
func (keeper *GateKeeper) GetRole(token string) string {
	if role, found := keeper.Token2Role.Load(token); found {
		return role.(string)
	}
	return "anonymous"
}

// GetHumanID 返回token对应的HumanID。
// 当token无效时返回空串。
func (keeper *GateKeeper) GetHumanID(token string) string {
	if humanID, found := keeper.Token2Role.Load(token); found {
		return humanID.(string)
	}
	return ""
}

// addHumanIDToken 添加可用于验证HumanID的token。
func (keeper *GateKeeper) addHumanIDToken(token, humanID string) {
	keeper.Token2HumanID.Store(token, humanID)
}

// addRoleToken 添加可用于验证角色的token。
func (keeper *GateKeeper) addRoleToken(token, role string) {
	keeper.Token2Role.Store(token, role)
}

// removeHumanIDToken 移除传入的HumanID token。
func (keeper *GateKeeper) removeHumanIDToken(token string) {
	keeper.Token2HumanID.Delete(token)
}

// removeRoleToken 移除传入的角色token。
func (keeper *GateKeeper) removeRoleToken(token string) {
	keeper.Token2Role.Delete(token)
}
