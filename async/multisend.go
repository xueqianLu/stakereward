package async

import (
	"crypto/ecdsa"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func (p *PullEvent) Multisend(useraddr string, amount *big.Int) (string, error) {
	privatek := beego.AppConfig.String("privatekey")
	privateKey, err := crypto.HexToECDSA(privatek)
	if err != nil {
		logs.Error("parse private key failed", "err", err)
		return "", err
	}

	sender, err := getAddrFromPrivk(privateKey)
	if err != nil {
		logs.Error("get addr from private key failed", "err", err)
		return "", err
	}
	nonce, err := p.client.NonceAt(p.ctx, sender, nil)
	if err != nil {
		logs.Error("get user nonce failed", "err", err)
		return "", err
	}
	chainId, err := p.client.ChainID(p.ctx)
	if err != nil {
		logs.Error("get chain id failed", "err", err, "use default chainid 269")
		chainId = big.NewInt(269)
	}

	to := common.HexToAddress(useraddr)
	gas := uint64(100000)
	gasPrice := big.NewInt(2000000000)
	tx := types.NewTransaction(nonce, to, amount, gas, gasPrice, []byte{})
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		logs.Error("sign transaction failed", "err", err)
		return "", err
	}

	err = p.client.SendTransaction(p.ctx, signedTx)
	if err != nil {
		logs.Error("send transaction failed", "err", err)
		return "", err
	}
	return signedTx.Hash().String(), nil
}

func getAddrFromPrivk(priv *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := priv.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress, nil
}
