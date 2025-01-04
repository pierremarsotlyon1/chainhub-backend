package utils

import (
	"main/interfaces"

	"github.com/ethereum/go-ethereum/common"
)

var WRAPPERS = []interfaces.Wrapper{
	{
		Name:        "Stake DAO (sdCRV)",
		Address:     STAKEDAO_LOCKERS,
		LogoUrl:     "https://www.defiwars.xyz/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Fstakedao.42f95a34.png&w=48&q=75",
		PoolAddress: common.HexToAddress("0xCA0253A98D16e9C1e3614caFDA19318EE69772D0"),
		OldPoolAddresses: []common.Address{
			common.HexToAddress("0xf7b55C3732aD8b2c2dA7c24f30A69f55c54FB717"),
		},
	},
	{
		Name:             "Yearn (yCRV)",
		Address:          YEARN_LOCKERS,
		LogoUrl:          "https://www.defiwars.xyz/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Fyearn.01d9002d.png&w=48&q=75",
		PoolAddress:      common.HexToAddress("0x99f5acc8ec2da2bc0771c32814eff52b712de1e5"),
		OldPoolAddresses: make([]common.Address, 0),
	},
	{
		Name:             "Convex (cvxCRV)",
		Address:          CONVEX_LOCKERS,
		LogoUrl:          "https://www.defiwars.xyz/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Fconvex.7542789f.jpeg&w=48&q=75",
		PoolAddress:      common.HexToAddress("0x971add32ea87f10bd192671630be3be8a11b8623"),
		OldPoolAddresses: make([]common.Address, 0),
	},
}
