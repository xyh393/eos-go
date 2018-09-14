package chain

// import (
// 	"fmt"
// 	"github.com/eosspark/eos-go/chain/types"
// 	"github.com/eosspark/eos-go/common"
// 	"github.com/eosspark/eos-go/contracts/system"
// 	"github.com/eosspark/eos-go/db"
// 	"time"
// )

// type AuthorizationManager struct {
// 	//control chain.Controller
// 	db *eosiodb.DataBase
// }

// type PermissionIdType uint64

// func (am *AuthorizationManager) AddIndices() {
// 	am.db.Insert(&types.PermissionObject{})
// 	am.db.Insert(&types.PermissionUsageObject{})
// 	am.db.Insert(&types.PermissionLinkObject{})
// }

// func (am *AuthorizationManager) InitializeDataBase() {

// }

// func (am *AuthorizationManager) CreatePermission(account common.AccountName,
// 	name common.PermissionName,
// 	parent PermissionIdType,
// 	auth common.Authority,
// 	initialCreationTime time.Duration,
// ) types.PermissionObject {
// 	creationTime := initialCreationTime
// 	if creationTime == 1 {
// 		//createTime = pendingBlockTime
// 	}

// 	var permUsage types.PermissionUsageObject
// 	permUsage.LastUsed = creationTime
// 	am.db.Insert(&permUsage)

// 	perm := types.PermissionObject{
// 		UsageId:     permUsage.Id,
// 		Parent:      uint64(parent),
// 		Owner:       account,
// 		Name:        name,
// 		LastUpdated: creationTime,
// 		//Auth:        SharedAuthority(),
// 	}
// 	am.db.Insert(&perm)
// 	return perm
// }

// func (am *AuthorizationManager) ModifyPermission(permission types.PermissionObject, auth common.Authority) {
// 	am.db.Update(&permission, func(data interface{}) error {
// 		//permission.Auth = auth
// 		//permission.LastUpdated = pendingBlockTime
// 		return nil
// 	})
// }

// func (am *AuthorizationManager) RemovePermission() {

// }

// func (am *AuthorizationManager) UpdatePermissionUsage() {
// 	var puo types.PermissionUsageObject
// 	am.db.Update(&puo, func(data interface{}) error {
// 		//puo.LastUsed = pendingBlockTime
// 		return nil
// 	})
// }

// func (am *AuthorizationManager) GetPermissionLastUsed(permission common.Permission) time.Duration {
// 	var puo types.PermissionUsageObject
// 	return puo.LastUsed
// }

// func (am *AuthorizationManager) FindPermission(level common.PermissionLevel) *types.PermissionObject {
// 	var po types.PermissionObject
// 	am.db.Find("", 0, &po)
// 	return &po
// }

// func (am *AuthorizationManager) GetPermission(level common.PermissionLevel) types.PermissionObject {
// 	var po types.PermissionObject
// 	am.db.Find("", 0, &po)
// 	return po
// }

// func (am *AuthorizationManager) LookupLinkedPermission(authorizerAccount common.AccountName,
// 	scope common.AccountName,
// 	actName common.ActionName,
// ) common.PermissionName {
// 	var pn common.PermissionName
// 	return pn
// }

// func (am *AuthorizationManager) LookupMinimumPermission(authorizerAccount common.AccountName,
// 	scope common.AccountName,
// 	actName common.ActionName,
// ) common.PermissionName {
// 	var pn common.PermissionName
// 	return pn
// }

// func (am *AuthorizationManager) CheckUpdateauthAuthorization(update system.UpdateAuth, auths []common.PermissionLevel) {
// 	if len(auths) != 1 {
// 		fmt.Println("error")
// 		return
// 	}
// 	auth := auths[0]
// 	if auth.Actor != update.Account {
// 		fmt.Println("error")
// 		return
// 	}
// 	minPermission := am.FindPermission(common.PermissionLevel{update.Account, update.Permission})
// 	if minPermission == nil {
// 		permission := am.GetPermission(common.PermissionLevel{update.Account, update.Permission})
// 		minPermission = &permission
// 	}
// 	if am.GetPermission(auth).Satisfies(*minPermission) == false {
// 		fmt.Println("error")
// 		return
// 	}
// }

// func (am *AuthorizationManager) CheckDeleteauthAuthorization(del system.DeleteAuth, auths []common.PermissionLevel) {
// 	if len(auths) != 1 {
// 		fmt.Println("error")
// 		return
// 	}
// 	auth := auths[0]
// 	if auth.Actor != del.Account {
// 		fmt.Println("error")
// 		return
// 	}
// 	minPermission := am.GetPermission(common.PermissionLevel{del.Account, del.Permission})
// 	if am.GetPermission(auth).Satisfies(minPermission) == false {
// 		fmt.Println("error")
// 		return
// 	}
// }

// func (am *AuthorizationManager) CheckLinkauthAuthorization(link system.LinkAuth, auths []common.PermissionLevel) {
// 	if len(auths) != 1 {
// 		fmt.Println("error")
// 		return
// 	}
// 	auth := auths[0]
// 	if auth.Actor != link.Account {
// 		fmt.Println("error")
// 		return
// 	}
// 	//待完善
// 	linkPermissionName := am.LookupMinimumPermission(link.Account, link.Code, link.Type)
// 	if &linkPermissionName == nil {
// 		return
// 	}
// 	//待完善
// }

// func (am *AuthorizationManager) CheckUnlinkauthAuthorization(unlink system.UnlinkAuth, auths []common.PermissionLevel) {
// 	if len(auths) != 1 {
// 		fmt.Println("error")
// 		return
// 	}
// 	auth := auths[0]
// 	if auth.Actor != unlink.Account {
// 		fmt.Println("error")
// 		return
// 	}
// 	//待完善
// }

// func (am *AuthorizationManager) CheckCanceldelayAuthorization(canceldelay system.CancelDelay, auths []common.PermissionLevel) {
// 	if len(auths) != 1 {
// 		fmt.Println("error")
// 		return
// 	}
// 	auth := auths[0]
// 	if am.GetPermission(auth).Satisfies(am.GetPermission(canceldelay.CancelingAuth)) == false {
// 		fmt.Println("error")
// 		return
// 	}
// 	//待完善
// }

// func (am *AuthorizationManager) CheckAuthorization(actions []types.Action,
// 	providedKeys []common.PublicKeyType,
// 	providedPermission []common.PermissionLevel,
// 	providedDelay time.Time,
// 	allowUnusedKeys bool,
// ) {
// 	//delayMaxLimit := am.control

// }

// func (am *AuthorizationManager) GetRequiredKeys(trx types.Transaction,
// 	candidateKeys []common.PublicKeyType,
// 	providedDelay time.Time) {
// 	//check := MakeAuthChecker()
// }
