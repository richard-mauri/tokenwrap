package main

import (
	"os"
	"fmt"
	"github.com/hashicorp/vault/api"
)

type VaultCryptoki struct {
	DerivedRootTokenTTL string
	DerivedRootTokenMaxTTL string
	DerivedRootTokenWrapTTL string
	DerivedRootTokenNumUses int
}

func (v *VaultCryptoki) DeriveRootToken(vaultAddr, rootToken string) (derivedToken string, err error) {
	vaultConfig := api.DefaultConfig()
	vaultConfig.Address = vaultAddr
	client, err := api.NewClient(vaultConfig)
	if err != nil {
		return derivedToken, err
	}

        tokenCreateRequest := &api.TokenCreateRequest {
                TTL: v.DerivedRootTokenTTL,
                ExplicitMaxTTL: v.DerivedRootTokenMaxTTL,
                DisplayName: "VCDerivedRootToken",
                NumUses: 10,
        }

        client.SetToken(rootToken)

	if v.DerivedRootTokenWrapTTL != "" {
		client.SetWrappingLookupFunc(func(op, path string) string {
			fmt.Printf("DeriveRootToken op=%s path=%s wrapTTL=%s\n", op, path, v.DerivedRootTokenWrapTTL)
			// expect op=POST path=auth/token/create

			return v.DerivedRootTokenWrapTTL
		})
	}

        auth := client.Auth()

        tokenAuth := auth.Token()

        secret, err := tokenAuth.Create(tokenCreateRequest)
        if err != nil {
                fmt.Printf("Error: Failure creating derived root token: %v\n", err)
                return derivedToken, err
        }

	fmt.Printf("Created token secret: %+v\n", secret)

	derivedToken, err = secret.TokenID()
	return derivedToken, err
}

func main() {
	v := &VaultCryptoki{
		DerivedRootTokenTTL: "5m",
		DerivedRootTokenMaxTTL: "30m",
		DerivedRootTokenWrapTTL: "1m",
		DerivedRootTokenNumUses: 5,
	}

	vaultAddr := os.Getenv("VAULT_ADDR")
	rootToken := os.Getenv("VAULT_TOKEN")

	token2, err := v.DeriveRootToken(vaultAddr, rootToken)
	if err != nil {
		fmt.Printf("Error from DeriveRootToken; %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Wrapped token: %+v\n", token2)

	v.DerivedRootTokenWrapTTL = ""
	token2, err = v.DeriveRootToken(vaultAddr, rootToken)
	if err != nil {
		fmt.Printf("Error from DeriveRootToken; %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Non wrapped token: %+v\n", token2)
}
