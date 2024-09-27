// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PullOracle

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PullOracleSignature is an auto generated low-level Go binding around an user-defined struct.
type PullOracleSignature struct {
	R [32]byte
	S [32]byte
	V uint8
}

// PullOracleMetaData contains all meta data concerning the PullOracle contract.
var PullOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"PullOracle__ConsensusRateNotSetForProtocol\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PullOracle__InsufficientSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"name\":\"PullOracle__InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PullOracle__InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"PullOracle__TransmittersNotSetForProtocol\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestTimestamp\",\"type\":\"uint256\"}],\"name\":\"PullOracle__UpdateTimestampTooOld\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endPoint\",\"outputs\":[{\"internalType\":\"contractEndPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structPullOracle.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"dataKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getLastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_protocolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_endPoint\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataKey\",\"type\":\"bytes32\"}],\"name\":\"latestUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"latestPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedData\",\"type\":\"bytes\"}],\"name\":\"updateMultipleAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100315730608052611f9090816100378239608051818181610ef60152818161101d015261127a0152f35b600080fdfe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a71461158057508063248a9ca3146115515780632f2ff15d1461149e57806336568abe1461140c5780633659cfe6146112555780634f1ef28614610fa257806352d1902d14610ee35780636910e33414610d6e578063715018a614610d2357806376bf54ad1461086d5780638da5cb5b1461084457806391d14854146107f757806399e51e1c146107cd578063a217fddf146107b1578063d547741f14610772578063da1f12ab14610753578063e4f6d9bb146101ba578063f0f0fce9146101845763f2fde38b146100f057600080fd5b3461017f57602036600319011261017f576101096115e9565b610111611a1c565b6001600160a01b0381161561012b5761012990611a74565b005b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b600080fd5b3461017f57602036600319011261017f5760043560005261012f6020526040806000206001815491015482519182526020820152f35b3461017f5760c036600319011261017f576001600160401b03602435811061017f5736602360243501121561017f5780602435600401351161017f5736602480356004013560051b813501011161017f57806044351161017f5736602360443501121561017f57604435600401351161017f573660246060604435600401350260443501011161017f5760643560005261012f6020526001604060002001548060a435106107335760405160843560208201526020815261027a8161161a565b6040516102a48161029660208201946064359060a43587611d0c565b03601f198101835282611635565b51902060405160208101918252602081526102be8161161a565b5190206102d060243560040135611d2e565b906102de6040519283611635565b6024803560048101358452019081602084015b602480356004013560051b813501018210610723575050610316816004358095611e1d565b156106c1578260405160208101917b0ca2ba3432b932bab69029b4b3b732b21026b2b9b9b0b3b29d05199960211b8352603c820152603c8152610358816115ff565b51902061012e5461012d5460405163e20aeb2d60e01b815260048101829052909290916001600160a01b0316602083602481845afa92831561061c5760009361068d575b50821561067457604051916343304c0f60e01b8352846004840152604083602481855afa92831561061c57600093610641575b508215610628576103e560443560040135611d6f565b60009560009360005b60443560040135811061047c575b5050505050156104305760643560005261012f60205260406000206084358155600160a43591015560206040516084358152f35b61043991611c72565b9061270f82018092116104665761271060449260405192625293eb60e61b84526004840152046024820152fd5b634e487b7160e01b600052601160045260246000fd5b60406104948260443560040135602460443501611da1565b013560ff8116810361017f5760006020916104bb8460443560040135602460443501611da1565b35906104e6846104d78760443560040135602460443501611da1565b01356040519384938c85611db1565b838052039060015afa1561061c576000516001600160a01b03811661050f575b506001016103ee565b6040516355b5190b60e01b81526020818061052e858960048401611dd0565b0381875afa90811561061c576000916105e2575b501561050657600160005b8b81106105b3575b50156105065790986001916001600160a01b03166105738287611de9565b520197612710898181020481148a151715610466576105958989928c02611dfd565b10156105a15789610506565b505050505050600184808080806103fc565b6001600160a01b03838116906105c9838a611de9565b5116146105d85760010161054d565b505060008b610555565b90506020813d602011610614575b816105fd60209383611635565b8101031261017f5761060e90611d45565b8b610542565b3d91506105f0565b6040513d6000823e3d90fd5b604051630bead28160e21b815260048101869052602490fd5b61066491935060403d60401161066d575b61065c8183611635565b810190611d52565b905091856103cf565b503d610652565b6040516333661b3b60e21b815260048101859052602490fd5b9092506020813d6020116106b9575b816106a960209383611635565b8101031261017f5751918461039c565b3d915061069c565b60405192633734bc9760e01b8452600484015260248301526060604483015260243560040135606483015260018060fb1b03602435600401351161017f576024356004013560051b9060848301376084816024356004013560051b8101030190fd5b81358152602091820191016102f1565b604490604051906342255b7f60e01b825260a43560048301526024820152fd5b3461017f57600036600319011261017f57602061012d54604051908152f35b3461017f57604036600319011261017f576101296004356107916115d3565b908060005260656020526107ac600160406000200154611673565b6119a6565b3461017f57600036600319011261017f57602060405160008152f35b3461017f57600036600319011261017f5761012e546040516001600160a01b039091168152602090f35b3461017f57604036600319011261017f576108106115d3565b600435600052606560205260406000209060018060a01b0316600052602052602060ff604060002054166040519015158152f35b3461017f57600036600319011261017f5760fb546040516001600160a01b039091168152602090f35b3461017f57602036600319011261017f576001600160401b03600435811061017f5736602360043501121561017f576004356004013590811161017f576024369160043501011161017f5760405160208101907b0ca2ba3432b932bab69029b4b3b732b21026b2b9b9b0b3b29d05199960211b825260246004350135603c820152603c81526108fb816115ff565b51902061012e5461012d5460405163e20aeb2d60e01b81526004810182905260219391926001600160a01b0316602082602481845afa91821561061c57600092610cef575b50811561067457604051926343304c0f60e01b8452846004850152604084602481855afa93841561061c57600094610ccc575b5083156106285761098c6044600435013560f81c611d6f565b60009560009360005b6044600435013560f81c8110610b66575b505050505015610b5c5760006001850160043586016024013560f81c5b8083106109cc57005b600160248360043501013560f81c9201926109e683611d2e565b926109f46040519485611635565b808452601f19610a0382611d2e565b013660208601376000905b808210610b38575050836004350160608095019360246040519260448101359283602086015260208552610a418561161a565b6040516020810190610a64816102966064968781013598899101359a8b87611d0c565b5190206040516020810191825260208152610a7e8161161a565b519020610a92816024600435013584611e1d565b15610ad45750505060019495965060005261012f602052604060002090848201918254841115610ac957555501905b9190916109c3565b505050500190610ac1565b899350604051938493633734bc9760e01b855284019160246004350135600486015260248501526044840152815180915260206084840192019060005b818110610b1f575050500390fd5b8251845285945060209384019390920191600101610b11565b90946020816024600193600435010135610b528989611de9565b5201950190610a0e565b9061043991611c72565b60206000610b9760418d600435019d019c6040519181604484930135906064602482013591013560f81c8c85611db1565b838052039060015afa1561061c576000516001600160a01b038116610bc0575b50600101610995565b6040516355b5190b60e01b815260208180610bdf858960048401611dd0565b0381875afa90811561061c57600091610c92575b5015610bb757600160005b8b8110610c63575b5015610bb75790986001916001600160a01b0316610c248287611de9565b520197612710808a02908a8204148a15171561046657610c45888a92611dfd565b1015610c51578a610bb7565b505050505050600185808080806109a6565b6001600160a01b0383811690610c79838a611de9565b511614610c8857600101610bfe565b505060008c610c06565b90506020813d602011610cc4575b81610cad60209383611635565b8101031261017f57610cbe90611d45565b8c610bf3565b3d9150610ca0565b610ce691945060403d60401161066d5761065c8183611635565b90509286610973565b9091506020813d602011610d1b575b81610d0b60209383611635565b8101031261017f57519085610940565b3d9150610cfe565b3461017f57600036600319011261017f57610d3c611a1c565b60fb80546001600160a01b031981169091556000906001600160a01b0316600080516020611efb8339815191528280a3005b3461017f57604036600319011261017f57610d876115d3565b60005460ff8160081c161590818092610ed6575b8015610ebf575b15610e635760ff19811660011760005581610e51575b50610dd360ff60005460081c16610dce81611cac565b611cac565b610ddc33611a74565b60005491610def60ff8460081c16611cac565b60043561012d5561012e80546001600160a01b0319166001600160a01b03909216919091179055610e1c57005b61ff0019166000557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a1005b61ffff19166101011760005582610db8565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b50303b158015610da25750600160ff821614610da2565b50600160ff821610610d9b565b3461017f57600036600319011261017f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610f3c576020604051600080516020611edb8339815191528152f35b60405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c6044820152771b1959081d1a1c9bdd59da0819195b1959d85d1958d85b1b60421b6064820152608490fd5b604036600319011261017f57610fb66115e9565b602435906001600160401b03821161017f573660238301121561017f578160040135610fe181611658565b610fee6040519182611635565b81815260209182820194366024838301011161017f578160009260248693018837830101526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169061104b30831415611aab565b611068600080516020611edb833981519152928284541614611afa565b611070611a1c565b600080516020611e9b8339815191525460ff161561109657505050506101299150611b49565b8492939416906040516352d1902d60e01b81528581600481865afa60009181611226575b506111095760405162461bcd60e51b815260048101879052602e6024820152600080516020611f3b83398151915260448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b949394036111cf5761111a82611b49565b600080516020611f1b833981519152600080a28251158015906111c7575b61113e57005b6000806101299560405195611152876115ff565b602787527f416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c86880152660819985a5b195960ca1b60408801525190845af4903d156111be573d6111a181611658565b906111af6040519283611635565b8152600081943d92013e611bd9565b60609250611bd9565b506001611138565b60405162461bcd60e51b815260048101849052602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608490fd5b9091508681813d831161124e575b61123e8183611635565b8101031261017f575190886110ba565b503d611234565b3461017f5760208060031936011261017f5761126f6115e9565b6001600160a01b03917f000000000000000000000000000000000000000000000000000000000000000083166112a730821415611aab565b6112c4600080516020611edb833981519152918583541614611afa565b6112cc611a1c565b604051828101949091906001600160401b038611838710176113f657856040526000835260ff600080516020611e9b833981519152541660001461131857505050506101299150611b49565b8492939416906040516352d1902d60e01b81528581600481865afa600091816113c7575b5061138b5760405162461bcd60e51b815260048101879052602e6024820152600080516020611f3b83398151915260448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b949394036111cf5761139c82611b49565b600080516020611f1b833981519152600080a28251158015906113bf5761113e57005b506000611138565b9091508681813d83116113ef575b6113df8183611635565b8101031261017f5751908861133c565b503d6113d5565b634e487b7160e01b600052604160045260246000fd5b3461017f57604036600319011261017f576114256115d3565b336001600160a01b0382160361144157610129906004356119a6565b60405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608490fd5b3461017f57604036600319011261017f576004356114ba6115d3565b8160005260656020526114d4600160406000200154611673565b81600052606560205260406000209060018060a01b0316908160005260205260ff604060002054161561150357005b8160005260656020526040600020816000526020526040600020600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4005b3461017f57602036600319011261017f5760043560005260656020526020600160406000200154604051908152f35b3461017f57602036600319011261017f576004359063ffffffff60e01b821680920361017f57602091637965db0b60e01b81149081156115c2575b5015158152f35b6301ffc9a760e01b149050836115bb565b602435906001600160a01b038216820361017f57565b600435906001600160a01b038216820361017f57565b606081019081106001600160401b038211176113f657604052565b604081019081106001600160401b038211176113f657604052565b601f909101601f19168101906001600160401b038211908210176113f657604052565b6001600160401b0381116113f657601f01601f191660200190565b60009080825260209060658252604092838120338252835260ff84822054161561169d5750505050565b338451926116aa846115ff565b602a8452848401908636833784511561194a5760308253845192600193841015611936576078602187015360295b8481116118cc575061189c57865192608084016001600160401b0381118582101761188857885260428452868401946060368737845115611874576030865384518210156118745790607860218601536041915b818311611806575050506117d6576117d29386936117b6936117a760489461177e9a519a8b9576020b1b1b2b9b9a1b7b73a3937b61d1030b1b1b7bab73a1604d1b8c880152518092603788019061195e565b8401917001034b99036b4b9b9b4b733903937b6329607d1b60378401525180938684019061195e565b01036028810187520185611635565b5192839262461bcd60e51b845260048401526024830190611981565b0390fd5b60648587519062461bcd60e51b82528060048301526024820152600080516020611e7b8339815191526044820152fd5b909192600f81166010811015611860576f181899199a1a9b1b9c1cb0b131b232b360811b901a6118368588611c85565b5360041c92801561184c5760001901919061172c565b634e487b7160e01b82526011600452602482fd5b634e487b7160e01b83526032600452602483fd5b634e487b7160e01b81526032600452602490fd5b634e487b7160e01b86526041600452602486fd5b60648688519062461bcd60e51b82528060048301526024820152600080516020611e7b8339815191526044820152fd5b90600f81166010811015611922576f181899199a1a9b1b9c1cb0b131b232b360811b901a6118fa8389611c85565b5360041c90801561190e57600019016116d8565b634e487b7160e01b86526011600452602486fd5b634e487b7160e01b87526032600452602487fd5b634e487b7160e01b85526032600452602485fd5b634e487b7160e01b84526032600452602484fd5b60005b8381106119715750506000910152565b8181015183820152602001611961565b9060209161199a8151809281855285808601910161195e565b601f01601f1916010190565b906000918083526065602052604083209160018060a01b03169182845260205260ff6040842054166119d757505050565b8083526065602052604083208284526020526040832060ff1981541690557ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b339380a4565b60fb546001600160a01b03163303611a3057565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b60fb80546001600160a01b039283166001600160a01b031982168117909255909116600080516020611efb833981519152600080a3565b15611ab257565b60405162461bcd60e51b815260206004820152602c6024820152600080516020611ebb83398151915260448201526b19195b1959d85d1958d85b1b60a21b6064820152608490fd5b15611b0157565b60405162461bcd60e51b815260206004820152602c6024820152600080516020611ebb83398151915260448201526b6163746976652070726f787960a01b6064820152608490fd5b803b15611b7e57600080516020611edb83398151915280546001600160a01b0319166001600160a01b03909216919091179055565b60405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608490fd5b91929015611c3b5750815115611bed575090565b3b15611bf65790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015611c4e5750805190602001fd5b60405162461bcd60e51b8152602060048201529081906117d2906024830190611981565b8181029291811591840414171561046657565b908151811015611c96570160200190565b634e487b7160e01b600052603260045260246000fd5b15611cb357565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b939291611d29906040928652606060208701526060860190611981565b930152565b6001600160401b0381116113f65760051b60200190565b5190811515820361017f57565b919082604091031261017f576020611d6983611d45565b92015190565b90611d7982611d2e565b611d866040519182611635565b8281528092611d97601f1991611d2e565b0190602036910137565b9190811015611c96576060020190565b9094939260ff6060936080840197845216602083015260408201520152565b9081526001600160a01b03909116602082015260400190565b8051821015611c965760209160051b010190565b8115611e07570490565b634e487b7160e01b600052601260045260246000fd5b90926000925b8251841015611e7157611e368484611de9565b5190600082821015611e62575060005260205260406000205b9260001981146104665760010192611e23565b90604092825260205220611e4f565b9150929150149056fe537472696e67733a20686578206c656e67746820696e73756666696369656e744910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914346756e6374696f6e206d7573742062652063616c6c6564207468726f75676820360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0bc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b45524331393637557067726164653a206e657720696d706c656d656e74617469a2646970667358221220c8628de91664f5cfab93549b5f3ed4989c7800879117324e2fd2187dae3e6c1364736f6c63430008130033",
}

// PullOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use PullOracleMetaData.ABI instead.
var PullOracleABI = PullOracleMetaData.ABI

// PullOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PullOracleMetaData.Bin instead.
var PullOracleBin = PullOracleMetaData.Bin

// DeployPullOracle deploys a new Ethereum contract, binding an instance of PullOracle to it.
func DeployPullOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PullOracle, error) {
	parsed, err := PullOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PullOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PullOracle{PullOracleCaller: PullOracleCaller{contract: contract}, PullOracleTransactor: PullOracleTransactor{contract: contract}, PullOracleFilterer: PullOracleFilterer{contract: contract}}, nil
}

// PullOracle is an auto generated Go binding around an Ethereum contract.
type PullOracle struct {
	PullOracleCaller     // Read-only binding to the contract
	PullOracleTransactor // Write-only binding to the contract
	PullOracleFilterer   // Log filterer for contract events
}

// PullOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PullOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PullOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PullOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PullOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PullOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PullOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PullOracleSession struct {
	Contract     *PullOracle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PullOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PullOracleCallerSession struct {
	Contract *PullOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PullOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PullOracleTransactorSession struct {
	Contract     *PullOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PullOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PullOracleRaw struct {
	Contract *PullOracle // Generic contract binding to access the raw methods on
}

// PullOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PullOracleCallerRaw struct {
	Contract *PullOracleCaller // Generic read-only contract binding to access the raw methods on
}

// PullOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PullOracleTransactorRaw struct {
	Contract *PullOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPullOracle creates a new instance of PullOracle, bound to a specific deployed contract.
func NewPullOracle(address common.Address, backend bind.ContractBackend) (*PullOracle, error) {
	contract, err := bindPullOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PullOracle{PullOracleCaller: PullOracleCaller{contract: contract}, PullOracleTransactor: PullOracleTransactor{contract: contract}, PullOracleFilterer: PullOracleFilterer{contract: contract}}, nil
}

// NewPullOracleCaller creates a new read-only instance of PullOracle, bound to a specific deployed contract.
func NewPullOracleCaller(address common.Address, caller bind.ContractCaller) (*PullOracleCaller, error) {
	contract, err := bindPullOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PullOracleCaller{contract: contract}, nil
}

// NewPullOracleTransactor creates a new write-only instance of PullOracle, bound to a specific deployed contract.
func NewPullOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*PullOracleTransactor, error) {
	contract, err := bindPullOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PullOracleTransactor{contract: contract}, nil
}

// NewPullOracleFilterer creates a new log filterer instance of PullOracle, bound to a specific deployed contract.
func NewPullOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*PullOracleFilterer, error) {
	contract, err := bindPullOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PullOracleFilterer{contract: contract}, nil
}

// bindPullOracle binds a generic wrapper to an already deployed contract.
func bindPullOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PullOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PullOracle *PullOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PullOracle.Contract.PullOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PullOracle *PullOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PullOracle.Contract.PullOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PullOracle *PullOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PullOracle.Contract.PullOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PullOracle *PullOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PullOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PullOracle *PullOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PullOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PullOracle *PullOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PullOracle.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PullOracle *PullOracleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PullOracle.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PullOracle *PullOracleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PullOracle.Contract.DEFAULTADMINROLE(&_PullOracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PullOracle *PullOracleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PullOracle.Contract.DEFAULTADMINROLE(&_PullOracle.CallOpts)
}

// EndPoint is a free data retrieval call binding the contract method 0x99e51e1c.
//
// Solidity: function endPoint() view returns(address)
func (_PullOracle *PullOracleCaller) EndPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PullOracle.contract.Call(opts, &out, "endPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EndPoint is a free data retrieval call binding the contract method 0x99e51e1c.
//
// Solidity: function endPoint() view returns(address)
func (_PullOracle *PullOracleSession) EndPoint() (common.Address, error) {
	return _PullOracle.Contract.EndPoint(&_PullOracle.CallOpts)
}

// EndPoint is a free data retrieval call binding the contract method 0x99e51e1c.
//
// Solidity: function endPoint() view returns(address)
func (_PullOracle *PullOracleCallerSession) EndPoint() (common.Address, error) {
	return _PullOracle.Contract.EndPoint(&_PullOracle.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PullOracle *PullOracleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PullOracle.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PullOracle *PullOracleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PullOracle.Contract.GetRoleAdmin(&_PullOracle.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PullOracle *PullOracleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PullOracle.Contract.GetRoleAdmin(&_PullOracle.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PullOracle *PullOracleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PullOracle.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PullOracle *PullOracleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PullOracle.Contract.HasRole(&_PullOracle.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PullOracle *PullOracleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PullOracle.Contract.HasRole(&_PullOracle.CallOpts, role, account)
}

// LatestUpdate is a free data retrieval call binding the contract method 0xf0f0fce9.
//
// Solidity: function latestUpdate(bytes32 dataKey) view returns(uint256 latestPrice, uint256 latestTimestamp)
func (_PullOracle *PullOracleCaller) LatestUpdate(opts *bind.CallOpts, dataKey [32]byte) (struct {
	LatestPrice     *big.Int
	LatestTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _PullOracle.contract.Call(opts, &out, "latestUpdate", dataKey)

	outstruct := new(struct {
		LatestPrice     *big.Int
		LatestTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LatestPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LatestTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestUpdate is a free data retrieval call binding the contract method 0xf0f0fce9.
//
// Solidity: function latestUpdate(bytes32 dataKey) view returns(uint256 latestPrice, uint256 latestTimestamp)
func (_PullOracle *PullOracleSession) LatestUpdate(dataKey [32]byte) (struct {
	LatestPrice     *big.Int
	LatestTimestamp *big.Int
}, error) {
	return _PullOracle.Contract.LatestUpdate(&_PullOracle.CallOpts, dataKey)
}

// LatestUpdate is a free data retrieval call binding the contract method 0xf0f0fce9.
//
// Solidity: function latestUpdate(bytes32 dataKey) view returns(uint256 latestPrice, uint256 latestTimestamp)
func (_PullOracle *PullOracleCallerSession) LatestUpdate(dataKey [32]byte) (struct {
	LatestPrice     *big.Int
	LatestTimestamp *big.Int
}, error) {
	return _PullOracle.Contract.LatestUpdate(&_PullOracle.CallOpts, dataKey)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PullOracle *PullOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PullOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PullOracle *PullOracleSession) Owner() (common.Address, error) {
	return _PullOracle.Contract.Owner(&_PullOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PullOracle *PullOracleCallerSession) Owner() (common.Address, error) {
	return _PullOracle.Contract.Owner(&_PullOracle.CallOpts)
}

// ProtocolId is a free data retrieval call binding the contract method 0xda1f12ab.
//
// Solidity: function protocolId() view returns(bytes32)
func (_PullOracle *PullOracleCaller) ProtocolId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PullOracle.contract.Call(opts, &out, "protocolId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProtocolId is a free data retrieval call binding the contract method 0xda1f12ab.
//
// Solidity: function protocolId() view returns(bytes32)
func (_PullOracle *PullOracleSession) ProtocolId() ([32]byte, error) {
	return _PullOracle.Contract.ProtocolId(&_PullOracle.CallOpts)
}

// ProtocolId is a free data retrieval call binding the contract method 0xda1f12ab.
//
// Solidity: function protocolId() view returns(bytes32)
func (_PullOracle *PullOracleCallerSession) ProtocolId() ([32]byte, error) {
	return _PullOracle.Contract.ProtocolId(&_PullOracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PullOracle *PullOracleCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PullOracle.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PullOracle *PullOracleSession) ProxiableUUID() ([32]byte, error) {
	return _PullOracle.Contract.ProxiableUUID(&_PullOracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PullOracle *PullOracleCallerSession) ProxiableUUID() ([32]byte, error) {
	return _PullOracle.Contract.ProxiableUUID(&_PullOracle.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PullOracle *PullOracleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PullOracle.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PullOracle *PullOracleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PullOracle.Contract.SupportsInterface(&_PullOracle.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PullOracle *PullOracleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PullOracle.Contract.SupportsInterface(&_PullOracle.CallOpts, interfaceId)
}

// GetLastPrice is a paid mutator transaction binding the contract method 0xe4f6d9bb.
//
// Solidity: function getLastPrice(bytes32 merkleRoot, bytes32[] merkleProof, (bytes32,bytes32,uint8)[] signatures, bytes32 dataKey, uint256 price, uint256 timestamp) returns(uint256)
func (_PullOracle *PullOracleTransactor) GetLastPrice(opts *bind.TransactOpts, merkleRoot [32]byte, merkleProof [][32]byte, signatures []PullOracleSignature, dataKey [32]byte, price *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _PullOracle.contract.Transact(opts, "getLastPrice", merkleRoot, merkleProof, signatures, dataKey, price, timestamp)
}

// GetLastPrice is a paid mutator transaction binding the contract method 0xe4f6d9bb.
//
// Solidity: function getLastPrice(bytes32 merkleRoot, bytes32[] merkleProof, (bytes32,bytes32,uint8)[] signatures, bytes32 dataKey, uint256 price, uint256 timestamp) returns(uint256)
func (_PullOracle *PullOracleSession) GetLastPrice(merkleRoot [32]byte, merkleProof [][32]byte, signatures []PullOracleSignature, dataKey [32]byte, price *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _PullOracle.Contract.GetLastPrice(&_PullOracle.TransactOpts, merkleRoot, merkleProof, signatures, dataKey, price, timestamp)
}

// GetLastPrice is a paid mutator transaction binding the contract method 0xe4f6d9bb.
//
// Solidity: function getLastPrice(bytes32 merkleRoot, bytes32[] merkleProof, (bytes32,bytes32,uint8)[] signatures, bytes32 dataKey, uint256 price, uint256 timestamp) returns(uint256)
func (_PullOracle *PullOracleTransactorSession) GetLastPrice(merkleRoot [32]byte, merkleProof [][32]byte, signatures []PullOracleSignature, dataKey [32]byte, price *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _PullOracle.Contract.GetLastPrice(&_PullOracle.TransactOpts, merkleRoot, merkleProof, signatures, dataKey, price, timestamp)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PullOracle *PullOracleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PullOracle.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PullOracle *PullOracleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PullOracle.Contract.GrantRole(&_PullOracle.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PullOracle *PullOracleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PullOracle.Contract.GrantRole(&_PullOracle.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x6910e334.
//
// Solidity: function initialize(bytes32 _protocolId, address _endPoint) returns()
func (_PullOracle *PullOracleTransactor) Initialize(opts *bind.TransactOpts, _protocolId [32]byte, _endPoint common.Address) (*types.Transaction, error) {
	return _PullOracle.contract.Transact(opts, "initialize", _protocolId, _endPoint)
}

// Initialize is a paid mutator transaction binding the contract method 0x6910e334.
//
// Solidity: function initialize(bytes32 _protocolId, address _endPoint) returns()
func (_PullOracle *PullOracleSession) Initialize(_protocolId [32]byte, _endPoint common.Address) (*types.Transaction, error) {
	return _PullOracle.Contract.Initialize(&_PullOracle.TransactOpts, _protocolId, _endPoint)
}

// Initialize is a paid mutator transaction binding the contract method 0x6910e334.
//
// Solidity: function initialize(bytes32 _protocolId, address _endPoint) returns()
func (_PullOracle *PullOracleTransactorSession) Initialize(_protocolId [32]byte, _endPoint common.Address) (*types.Transaction, error) {
	return _PullOracle.Contract.Initialize(&_PullOracle.TransactOpts, _protocolId, _endPoint)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PullOracle *PullOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PullOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PullOracle *PullOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _PullOracle.Contract.RenounceOwnership(&_PullOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PullOracle *PullOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PullOracle.Contract.RenounceOwnership(&_PullOracle.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PullOracle *PullOracleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PullOracle.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PullOracle *PullOracleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PullOracle.Contract.RenounceRole(&_PullOracle.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PullOracle *PullOracleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PullOracle.Contract.RenounceRole(&_PullOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PullOracle *PullOracleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PullOracle.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PullOracle *PullOracleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PullOracle.Contract.RevokeRole(&_PullOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PullOracle *PullOracleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PullOracle.Contract.RevokeRole(&_PullOracle.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PullOracle *PullOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PullOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PullOracle *PullOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PullOracle.Contract.TransferOwnership(&_PullOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PullOracle *PullOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PullOracle.Contract.TransferOwnership(&_PullOracle.TransactOpts, newOwner)
}

// UpdateMultipleAssets is a paid mutator transaction binding the contract method 0x76bf54ad.
//
// Solidity: function updateMultipleAssets(bytes encodedData) returns()
func (_PullOracle *PullOracleTransactor) UpdateMultipleAssets(opts *bind.TransactOpts, encodedData []byte) (*types.Transaction, error) {
	return _PullOracle.contract.Transact(opts, "updateMultipleAssets", encodedData)
}

// UpdateMultipleAssets is a paid mutator transaction binding the contract method 0x76bf54ad.
//
// Solidity: function updateMultipleAssets(bytes encodedData) returns()
func (_PullOracle *PullOracleSession) UpdateMultipleAssets(encodedData []byte) (*types.Transaction, error) {
	return _PullOracle.Contract.UpdateMultipleAssets(&_PullOracle.TransactOpts, encodedData)
}

// UpdateMultipleAssets is a paid mutator transaction binding the contract method 0x76bf54ad.
//
// Solidity: function updateMultipleAssets(bytes encodedData) returns()
func (_PullOracle *PullOracleTransactorSession) UpdateMultipleAssets(encodedData []byte) (*types.Transaction, error) {
	return _PullOracle.Contract.UpdateMultipleAssets(&_PullOracle.TransactOpts, encodedData)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PullOracle *PullOracleTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PullOracle.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PullOracle *PullOracleSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PullOracle.Contract.UpgradeTo(&_PullOracle.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PullOracle *PullOracleTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PullOracle.Contract.UpgradeTo(&_PullOracle.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PullOracle *PullOracleTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PullOracle.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PullOracle *PullOracleSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PullOracle.Contract.UpgradeToAndCall(&_PullOracle.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PullOracle *PullOracleTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PullOracle.Contract.UpgradeToAndCall(&_PullOracle.TransactOpts, newImplementation, data)
}

// PullOracleAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the PullOracle contract.
type PullOracleAdminChangedIterator struct {
	Event *PullOracleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PullOracleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PullOracleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PullOracleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PullOracleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PullOracleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PullOracleAdminChanged represents a AdminChanged event raised by the PullOracle contract.
type PullOracleAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PullOracle *PullOracleFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PullOracleAdminChangedIterator, error) {

	logs, sub, err := _PullOracle.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PullOracleAdminChangedIterator{contract: _PullOracle.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PullOracle *PullOracleFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PullOracleAdminChanged) (event.Subscription, error) {

	logs, sub, err := _PullOracle.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PullOracleAdminChanged)
				if err := _PullOracle.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PullOracle *PullOracleFilterer) ParseAdminChanged(log types.Log) (*PullOracleAdminChanged, error) {
	event := new(PullOracleAdminChanged)
	if err := _PullOracle.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PullOracleBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the PullOracle contract.
type PullOracleBeaconUpgradedIterator struct {
	Event *PullOracleBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PullOracleBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PullOracleBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PullOracleBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PullOracleBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PullOracleBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PullOracleBeaconUpgraded represents a BeaconUpgraded event raised by the PullOracle contract.
type PullOracleBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PullOracle *PullOracleFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*PullOracleBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PullOracle.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &PullOracleBeaconUpgradedIterator{contract: _PullOracle.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PullOracle *PullOracleFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *PullOracleBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PullOracle.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PullOracleBeaconUpgraded)
				if err := _PullOracle.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PullOracle *PullOracleFilterer) ParseBeaconUpgraded(log types.Log) (*PullOracleBeaconUpgraded, error) {
	event := new(PullOracleBeaconUpgraded)
	if err := _PullOracle.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PullOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PullOracle contract.
type PullOracleInitializedIterator struct {
	Event *PullOracleInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PullOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PullOracleInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PullOracleInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PullOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PullOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PullOracleInitialized represents a Initialized event raised by the PullOracle contract.
type PullOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PullOracle *PullOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*PullOracleInitializedIterator, error) {

	logs, sub, err := _PullOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PullOracleInitializedIterator{contract: _PullOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PullOracle *PullOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PullOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _PullOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PullOracleInitialized)
				if err := _PullOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PullOracle *PullOracleFilterer) ParseInitialized(log types.Log) (*PullOracleInitialized, error) {
	event := new(PullOracleInitialized)
	if err := _PullOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PullOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PullOracle contract.
type PullOracleOwnershipTransferredIterator struct {
	Event *PullOracleOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PullOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PullOracleOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PullOracleOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PullOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PullOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PullOracleOwnershipTransferred represents a OwnershipTransferred event raised by the PullOracle contract.
type PullOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PullOracle *PullOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PullOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PullOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PullOracleOwnershipTransferredIterator{contract: _PullOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PullOracle *PullOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PullOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PullOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PullOracleOwnershipTransferred)
				if err := _PullOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PullOracle *PullOracleFilterer) ParseOwnershipTransferred(log types.Log) (*PullOracleOwnershipTransferred, error) {
	event := new(PullOracleOwnershipTransferred)
	if err := _PullOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PullOracleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PullOracle contract.
type PullOracleRoleAdminChangedIterator struct {
	Event *PullOracleRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PullOracleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PullOracleRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PullOracleRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PullOracleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PullOracleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PullOracleRoleAdminChanged represents a RoleAdminChanged event raised by the PullOracle contract.
type PullOracleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PullOracle *PullOracleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PullOracleRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PullOracle.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PullOracleRoleAdminChangedIterator{contract: _PullOracle.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PullOracle *PullOracleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PullOracleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PullOracle.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PullOracleRoleAdminChanged)
				if err := _PullOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PullOracle *PullOracleFilterer) ParseRoleAdminChanged(log types.Log) (*PullOracleRoleAdminChanged, error) {
	event := new(PullOracleRoleAdminChanged)
	if err := _PullOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PullOracleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PullOracle contract.
type PullOracleRoleGrantedIterator struct {
	Event *PullOracleRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PullOracleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PullOracleRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PullOracleRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PullOracleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PullOracleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PullOracleRoleGranted represents a RoleGranted event raised by the PullOracle contract.
type PullOracleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PullOracle *PullOracleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PullOracleRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PullOracle.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PullOracleRoleGrantedIterator{contract: _PullOracle.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PullOracle *PullOracleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PullOracleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PullOracle.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PullOracleRoleGranted)
				if err := _PullOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PullOracle *PullOracleFilterer) ParseRoleGranted(log types.Log) (*PullOracleRoleGranted, error) {
	event := new(PullOracleRoleGranted)
	if err := _PullOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PullOracleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PullOracle contract.
type PullOracleRoleRevokedIterator struct {
	Event *PullOracleRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PullOracleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PullOracleRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PullOracleRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PullOracleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PullOracleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PullOracleRoleRevoked represents a RoleRevoked event raised by the PullOracle contract.
type PullOracleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PullOracle *PullOracleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PullOracleRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PullOracle.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PullOracleRoleRevokedIterator{contract: _PullOracle.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PullOracle *PullOracleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PullOracleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PullOracle.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PullOracleRoleRevoked)
				if err := _PullOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PullOracle *PullOracleFilterer) ParseRoleRevoked(log types.Log) (*PullOracleRoleRevoked, error) {
	event := new(PullOracleRoleRevoked)
	if err := _PullOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PullOracleUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PullOracle contract.
type PullOracleUpgradedIterator struct {
	Event *PullOracleUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PullOracleUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PullOracleUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PullOracleUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PullOracleUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PullOracleUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PullOracleUpgraded represents a Upgraded event raised by the PullOracle contract.
type PullOracleUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PullOracle *PullOracleFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PullOracleUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PullOracle.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PullOracleUpgradedIterator{contract: _PullOracle.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PullOracle *PullOracleFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PullOracleUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PullOracle.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PullOracleUpgraded)
				if err := _PullOracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PullOracle *PullOracleFilterer) ParseUpgraded(log types.Log) (*PullOracleUpgraded, error) {
	event := new(PullOracleUpgraded)
	if err := _PullOracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
