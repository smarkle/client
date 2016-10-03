// Auto-generated by avdl-compiler v1.3.9 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/tlf_keys.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type CanonicalTlfName string
type CryptKey struct {
	KeyGeneration int     `codec:"KeyGeneration" json:"KeyGeneration"`
	Key           Bytes32 `codec:"Key" json:"Key"`
}

type TLFCryptKeys struct {
	TlfID         TLFID            `codec:"tlfID" json:"tlfID"`
	CanonicalName CanonicalTlfName `codec:"CanonicalName" json:"CanonicalName"`
	CryptKeys     []CryptKey       `codec:"CryptKeys" json:"CryptKeys"`
}

type CanonicalTLFNameAndID struct {
	TlfID         TLFID            `codec:"tlfID" json:"tlfID"`
	CanonicalName CanonicalTlfName `codec:"CanonicalName" json:"CanonicalName"`
}

type GetTLFCryptKeysArg struct {
	TlfName string `codec:"tlfName" json:"tlfName"`
}

type GetPublicCanonicalTLFNameAndIDArg struct {
	TlfName string `codec:"tlfName" json:"tlfName"`
}

type TlfKeysInterface interface {
	// getTLFCryptKeys returns TLF crypt keys from all generations and the TLF ID.
	// TLF ID should not be cached or stored persistently.
	GetTLFCryptKeys(context.Context, string) (TLFCryptKeys, error)
	// getPublicCanonicalTLFNameAndID return the canonical name and TLFID for tlfName.
	// TLF ID should not be cached or stored persistently.
	GetPublicCanonicalTLFNameAndID(context.Context, string) (CanonicalTLFNameAndID, error)
}

func TlfKeysProtocol(i TlfKeysInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.tlfKeys",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getTLFCryptKeys": {
				MakeArg: func() interface{} {
					ret := make([]GetTLFCryptKeysArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetTLFCryptKeysArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetTLFCryptKeysArg)(nil), args)
						return
					}
					ret, err = i.GetTLFCryptKeys(ctx, (*typedArgs)[0].TlfName)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getPublicCanonicalTLFNameAndID": {
				MakeArg: func() interface{} {
					ret := make([]GetPublicCanonicalTLFNameAndIDArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetPublicCanonicalTLFNameAndIDArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetPublicCanonicalTLFNameAndIDArg)(nil), args)
						return
					}
					ret, err = i.GetPublicCanonicalTLFNameAndID(ctx, (*typedArgs)[0].TlfName)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type TlfKeysClient struct {
	Cli rpc.GenericClient
}

// getTLFCryptKeys returns TLF crypt keys from all generations and the TLF ID.
// TLF ID should not be cached or stored persistently.
func (c TlfKeysClient) GetTLFCryptKeys(ctx context.Context, tlfName string) (res TLFCryptKeys, err error) {
	__arg := GetTLFCryptKeysArg{TlfName: tlfName}
	err = c.Cli.Call(ctx, "keybase.1.tlfKeys.getTLFCryptKeys", []interface{}{__arg}, &res)
	return
}

// getPublicCanonicalTLFNameAndID return the canonical name and TLFID for tlfName.
// TLF ID should not be cached or stored persistently.
func (c TlfKeysClient) GetPublicCanonicalTLFNameAndID(ctx context.Context, tlfName string) (res CanonicalTLFNameAndID, err error) {
	__arg := GetPublicCanonicalTLFNameAndIDArg{TlfName: tlfName}
	err = c.Cli.Call(ctx, "keybase.1.tlfKeys.getPublicCanonicalTLFNameAndID", []interface{}{__arg}, &res)
	return
}
