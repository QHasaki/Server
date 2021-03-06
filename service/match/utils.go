package match

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

func (m *Match) checkArgsandAuth(r *http.Request, req proto.Message) (*model.User, *protocol.Status) {
	data, token, err := utils.GetReqAndToken(r)
	if err != nil {
		return nil, &protocol.Status{
			Code:    protocol.Code_DATA_LOSE,
			Message: err.Error(),
		}
	}
	if err := proto.Unmarshal(data, req); err != nil {
		return nil, &protocol.Status{
			Code:    protocol.Code_DATA_LOSE,
			Message: "failed to unmarshal request body",
		}
	}

	// check token
	userID, err := m.cache.GetUserIDByToken(token)
	if err != nil {
		return nil, &protocol.Status{
			Code:    protocol.Code_UNAUTHORIZATED,
			Message: "invalid token",
		}
	}

	// get user from db, and get the operation auth of the player
	user, err := m.db.GetUserByID(userID)
	if err != nil {
		return nil, &protocol.Status{
			Code:    protocol.Code_UNAUTHORIZATED,
			Message: "failed to get user info",
		}
	}

	if user.Role != module.MANAGER {
		return nil, &protocol.Status{
			Code:    protocol.Code_PERMISSION_DENIED,
			Message: "permission denied",
		}
	}
	return user, nil
}
