package common

import "github.com/hashicorp/consul/api"

type BackupACLData struct {
	Policies []*api.ACLPolicy
	Roles    []*api.ACLRole
	Tokens   []*api.ACLToken
}
