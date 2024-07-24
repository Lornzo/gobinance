package accountswebsockets

import "github.com/Lornzo/gobinance/usdmfutures/accounts/accountsrests"

type restful interface {
	ListenKeyCreate() *accountsrests.ListenKeyCreate
	ListenKeyDelete() *accountsrests.ListenKeyDelete
	ListenKeyUpdate() *accountsrests.ListenKeyUpdate
}
