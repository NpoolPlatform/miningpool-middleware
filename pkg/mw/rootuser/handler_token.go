package rootuser

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/NpoolPlatform/go-service-framework/pkg/secure"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	rootusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	rootuserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/rootuser"
)

const (
	// must be 16 or 32
	TokenSaltLen = 16
)

//nolint:gocognit
func (h *Handler) withAuthTokenEncrypt() error {
	if h.AuthTokenPlain == nil {
		return fmt.Errorf("invalid authtoken")
	}
	authTokenSalt := make([]byte, TokenSaltLen)
	_, err := io.ReadFull(rand.Reader, authTokenSalt)
	if err != nil {
		return err
	}

	secure.SALT = authTokenSalt
	cipher, err := secure.EncryptAES([]byte(*h.AuthTokenPlain))
	if err != nil {
		return err
	}

	_authTokenSalt := base64.RawStdEncoding.EncodeToString(authTokenSalt)
	h.AuthTokenSalt = &_authTokenSalt

	_authToken := base64.RawStdEncoding.EncodeToString(cipher)
	h.AuthToken = &_authToken

	return nil
}

type TokenInfo struct {
	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"miningpool_type"
	MiningpoolTypeStr string            `protobuf:"bytes,40,opt,name=MiningpoolTypeStr,proto3" json:"MiningpoolTypeStr,omitempty" sql:"miningpool_type"`
	MiningpoolType    v1.MiningpoolType `protobuf:"varint,41,opt,name=MiningpoolType,proto3,enum=basetypes.miningpool.v1.MiningpoolType" json:"MiningpoolType,omitempty"`
	// @inject_tag: sql:"auth_token"
	AuthToken      string `protobuf:"bytes,60,opt,name=AuthToken,proto3" json:"AuthToken,omitempty" sql:"auth_token"`
	AuthTokenPlain string
	// @inject_tag: sql:"auth_token_salt"
	AuthTokenSalt string `protobuf:"bytes,60,opt,name=AuthTokenSalt,proto3" json:"AuthTokenSalt,omitempty" sql:"auth_token_salt"`
}

type tokenHandler struct {
	*Handler
	stm   *ent.RootUserSelect
	infos []*TokenInfo
	total uint32
}

func (h *tokenHandler) selectToken(stm *ent.RootUserQuery) {
	h.stm = stm.Select(
		rootuserent.FieldID,
		rootuserent.FieldEntID,
		rootuserent.FieldMiningpoolType,
		rootuserent.FieldAuthToken,
		rootuserent.FieldAuthTokenSalt,
	)
}

func (h *tokenHandler) queryToken(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.RootUser.Query().Where(rootuserent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(rootuserent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(rootuserent.EntID(*h.EntID))
	}
	h.selectToken(stm)
	return nil
}

func (h *tokenHandler) queryTokens(ctx context.Context, cli *ent.Client) error {
	stm, err := rootusercrud.SetQueryConds(cli.RootUser.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectToken(stm)
	return nil
}

func (h *tokenHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *tokenHandler) decryptToken() error {
	for _, v := range h.infos {
		salt, err := base64.RawStdEncoding.DecodeString(v.AuthTokenSalt)
		if err != nil {
			return err
		}

		authToken, err := base64.RawStdEncoding.DecodeString(v.AuthToken)
		if err != nil {
			return err
		}

		secure.SALT = salt
		_plain, err := secure.DecryptAES(authToken)
		if err != nil {
			return err
		}
		v.AuthTokenPlain = string(_plain)
	}
	return nil
}

func (h *tokenHandler) formate() {
	for _, v := range h.infos {
		v.MiningpoolType = v1.MiningpoolType(v1.MiningpoolType_value[v.MiningpoolTypeStr])
	}
}

func (h *Handler) GetAuthToken(ctx context.Context) (*TokenInfo, error) {
	handler := &tokenHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryToken(cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	err = handler.decryptToken()
	if err != nil {
		return nil, err
	}

	handler.formate()
	return handler.infos[0], nil
}

func (h *Handler) GetAuthTokens(ctx context.Context) ([]*TokenInfo, uint32, error) {
	handler := &tokenHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTokens(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(rootuserent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	err = handler.decryptToken()
	if err != nil {
		return nil, 0, err
	}

	handler.formate()
	return handler.infos, handler.total, nil
}
