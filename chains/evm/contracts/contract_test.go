package contracts

import (
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"
	"github.comsygmaprotcolsygma-core/chains/evm/transactor"
	"github.comsygmaprotcolsygma-core/mock"
	"go.uber.org/mock/gomock"
)

const testABI = `[{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"approve","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"mint","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"safeTransferFrom","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"},{"internalType":"bytes","name":"_data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"transferFrom","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"approved","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"constant":true,"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"getApproved","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"ownerOf","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"}]`
const testBin = "0x60806040523480156200001157600080fd5b5060405162003622380380620036228339810160408190526200003491620004db565b6000805460ff199081169091556002805490911660ff871617905562000066836200018a602090811b62001ef917901c565b600260016101000a81548160ff021916908360ff1602179055506200009682620001e760201b62001f501760201c565b6002806101000a8154816001600160801b0302191690836001600160801b03160217905550620000d1816200023e60201b62001fa51760201c565b6002805464ffffffffff92909216600160901b0264ffffffffff60901b199092169190911790556200010e60006200010862000297565b620002da565b60005b84518110156200017e57620001697fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4868381518110620001555762000155620005e4565b6020026020010151620002ea60201b60201c565b806200017581620005fa565b91505062000111565b50505050505062000624565b60006101008210620001e35760405162461bcd60e51b815260206004820152601c60248201527f76616c756520646f6573206e6f742066697420696e203820626974730000000060448201526064015b60405180910390fd5b5090565b6000600160801b8210620001e35760405162461bcd60e51b815260206004820152601e60248201527f76616c756520646f6573206e6f742066697420696e20313238206269747300006044820152606401620001da565b6000650100000000008210620001e35760405162461bcd60e51b815260206004820152601d60248201527f76616c756520646f6573206e6f742066697420696e20343020626974730000006044820152606401620001da565b60003360143610801590620002c457506001600160a01b03811660009081526005602052604090205460ff165b15620002d5575060131936013560601c5b919050565b620002e6828262000377565b5050565b60008281526001602052604090206002015462000311906200030b62000297565b620003f2565b620002da5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60448201526e0818591b5a5b881d1bc819dc985b9d608a1b6064820152608401620001da565b60008281526001602090815260409091206200039e91839062001ffc62000421821b17901c565b15620002e657620003ae62000297565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60008281526001602090815260408220620004189184906200201162000438821b17901c565b90505b92915050565b600062000418836001600160a01b0384166200045b565b6001600160a01b0381166000908152600183016020526040812054151562000418565b6000818152600183016020526040812054620004a4575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556200041b565b5060006200041b565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b0381168114620002d557600080fd5b600080600080600060a08688031215620004f457600080fd5b855160ff811681146200050657600080fd5b602087810151919650906001600160401b03808211156200052657600080fd5b818901915089601f8301126200053b57600080fd5b815181811115620005505762000550620004ad565b8060051b604051601f19603f83011681018181108582111715620005785762000578620004ad565b60405291825284820192508381018501918c8311156200059757600080fd5b938501935b82851015620005c057620005b085620004c3565b845293850193928501926200059c565b60408c015160608d01516080909d01519b9e919d509b9a9950975050505050505050565b634e487b7160e01b600052603260045260246000fd5b60006000198214156200061d57634e487b7160e01b600052601160045260246000fd5b5060010190565b612fee80620006346000396000f3fe6080604052600436106102465760003560e01c806391c404ac11610139578063c5b37c22116100b6578063d15ef64e1161007a578063d15ef64e1461074f578063d547741f1461076f578063d7a9cd791461078f578063edc20c3c146107ae578063f8c39e44146107ce578063ffaac0eb146107fe57600080fd5b8063c5b37c2214610676578063c5ec8970146106b4578063ca15c873146106ef578063cb10f2151461070f578063cdb0f73a1461072f57600080fd5b80639debb3bd116100fd5780639debb3bd146105df578063a217fddf146105f4578063a9cf69fa14610609578063bd2a182014610636578063c0331b3e1461065657600080fd5b806391c404ac1461053157806391d1485414610551578063926d7d7f146105715780639d82dd63146105935780639dd694f4146105b357600080fd5b8063541d5548116101c7578063802aabe81161018b578063802aabe81461047957806380ae1c281461048e57806384db809f146104a35780638c0c2631146104f15780639010d07c1461051157600080fd5b8063541d5548146103d15780635a1ad87c146104015780635c975abb146104215780635e1fab0f146104395780637febe63f1461045957600080fd5b806336568abe1161020e57806336568abe146103035780634603ae38146103235780634b0b919d146103435780634e056005146103915780634e0df3f6146103b157600080fd5b806305e2ca171461024b57806317f03ce514610260578063206a98fd14610280578063248a9ca3146102a05780632f2ff15d146102e3575b600080fd5b61025e6102593660046126c3565b610813565b005b34801561026c57600080fd5b5061025e61027b366004612733565b610a0e565b34801561028c57600080fd5b5061025e61029b36600461277f565b610ca1565b3480156102ac57600080fd5b506102d06102bb3660046127fb565b60009081526001602052604090206002015490565b6040519081526020015b60405180910390f35b3480156102ef57600080fd5b5061025e6102fe366004612829565b610f29565b34801561030f57600080fd5b5061025e61031e366004612829565b610fb9565b34801561032f57600080fd5b5061025e61033e36600461289d565b611043565b34801561034f57600080fd5b5061037961035e3660046128fc565b6003602052600090815260409020546001600160401b031681565b6040516001600160401b0390911681526020016102da565b34801561039d57600080fd5b5061025e6103ac3660046127fb565b6110e7565b3480156103bd57600080fd5b506102d06103cc366004612829565b61114a565b3480156103dd57600080fd5b506103f16103ec366004612917565b611176565b60405190151581526020016102da565b34801561040d57600080fd5b5061025e61041c36600461294c565b611190565b34801561042d57600080fd5b5060005460ff166103f1565b34801561044557600080fd5b5061025e610454366004612917565b611246565b34801561046557600080fd5b506103f16104743660046129aa565b6112d2565b34801561048557600080fd5b506102d0611378565b34801561049a57600080fd5b5061025e611396565b3480156104af57600080fd5b506104d96104be3660046127fb565b6004602052600090815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020016102da565b3480156104fd57600080fd5b5061025e61050c3660046129fa565b6113b0565b34801561051d57600080fd5b506104d961052c366004612a28565b61141c565b34801561053d57600080fd5b5061025e61054c3660046127fb565b61143b565b34801561055d57600080fd5b506103f161056c366004612829565b6114d5565b34801561057d57600080fd5b506102d0600080516020612f9983398151915281565b34801561059f57600080fd5b5061025e6105ae366004612917565b6114ed565b3480156105bf57600080fd5b506002546105cd9060ff1681565b60405160ff90911681526020016102da565b3480156105eb57600080fd5b506102d060c881565b34801561060057600080fd5b506102d0600081565b34801561061557600080fd5b50610629610624366004612733565b6115a2565b6040516102da9190612a82565b34801561064257600080fd5b5061025e610651366004612b38565b611670565b34801561066257600080fd5b5061025e610671366004612bca565b6116a6565b34801561068257600080fd5b5060025461069c906201000090046001600160801b031681565b6040516001600160801b0390911681526020016102da565b3480156106c057600080fd5b506002546106d990600160901b900464ffffffffff1681565b60405164ffffffffff90911681526020016102da565b3480156106fb57600080fd5b506102d061070a3660046127fb565b611bc1565b34801561071b57600080fd5b5061025e61072a366004612c38565b611bd8565b34801561073b57600080fd5b5061025e61074a366004612917565b611c6c565b34801561075b57600080fd5b5061025e61076a366004612c58565b611d72565b34801561077b57600080fd5b5061025e61078a366004612829565b611da5565b34801561079b57600080fd5b506002546105cd90610100900460ff1681565b3480156107ba57600080fd5b5061025e6107c9366004612c8d565b611e28565b3480156107da57600080fd5b506103f16107e9366004612917565b60056020526000908152604090205460ff1681565b34801561080a57600080fd5b5061025e611ee1565b61081b612033565b6002546201000090046001600160801b031634146108795760405162461bcd60e51b8152602060048201526016602482015275125b98dbdc9c9958dd08199959481cdd5c1c1b1a595960521b60448201526064015b60405180910390fd5b6000838152600460205260409020546001600160a01b0316806108de5760405162461bcd60e51b815260206004820181905260248201527f7265736f757263654944206e6f74206d617070656420746f2068616e646c65726044820152606401610870565b60ff8516600090815260036020526040812080548290610906906001600160401b0316612ccd565b91906101000a8154816001600160401b0302191690836001600160401b03160217905590506000610935612079565b60405163b07e54bb60e01b815290915083906000906001600160a01b0383169063b07e54bb9061096f908b9087908c908c90600401612d1d565b6000604051808303816000875af115801561098e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109b69190810190612d82565b9050826001600160a01b03167f17bc3181e17a9620a479c24e6c606e474ba84fc036877b768926872e8cd0e11f8a8a878b8b876040516109fb96959493929190612e1b565b60405180910390a2505050505050505050565b610a166120ba565b60ff838116600884901b68ffffffffffffffff0016176000818152600660209081526040808320868452909152808220815160808101909252805493949293919290918391166004811115610a6d57610a6d612a4a565b6004811115610a7e57610a7e612a4a565b8152905461010081046001600160c81b03166020830152600160d01b810460ff166040830152600160d81b900464ffffffffff1660609091015280519091506001816004811115610ad157610ad1612a4a565b1480610aee57506002816004811115610aec57610aec612a4a565b145b610b3a5760405162461bcd60e51b815260206004820152601c60248201527f50726f706f73616c2063616e6e6f742062652063616e63656c6c6564000000006044820152606401610870565b600254606083015164ffffffffff600160901b909204821691610b5f9143911661213e565b64ffffffffff1611610bb35760405162461bcd60e51b815260206004820181905260248201527f50726f706f73616c206e6f7420617420657870697279207468726573686f6c646044820152606401610870565b600480835268ffffffffffffffffff841660009081526006602090815260408083208884529091529020835181548593839160ff1916906001908490811115610bfe57610bfe612a4a565b02179055506020820151815460408085015160609095015164ffffffffff16600160d81b026001600160d81b0360ff909616600160d01b0260ff60d01b196001600160c81b039095166101000294909416610100600160d81b031990931692909217929092179390931692909217905551600080516020612f7983398151915290610c9190889088906004908990612e6c565b60405180910390a1505050505050565b610ca9612180565b610cb1612033565b60008281526004602090815260408083205490516001600160a01b039091169268ffffffffffffffff0060088a901b1660ff8b1617929091610cf99185918a918a9101612ea1565b60408051601f19818403018152918152815160209283012068ffffffffffffffffff851660009081526006845282812082825290935291209091506002815460ff166004811115610d4c57610d4c612a4a565b14610d995760405162461bcd60e51b815260206004820181905260248201527f50726f706f73616c206d757374206861766520506173736564207374617475736044820152606401610870565b805460ff19166003178155838515610e125760405163712467f960e11b81526001600160a01b0382169063e248cff290610ddb908a908d908d90600401612ecd565b600060405180830381600087803b158015610df557600080fd5b505af1158015610e09573d6000803e3d6000fd5b50505050610eef565b60405163712467f960e11b81526001600160a01b0382169063e248cff290610e42908a908d908d90600401612ecd565b600060405180830381600087803b158015610e5c57600080fd5b505af1925050508015610e6d575060015b610eef573d808015610e9b576040519150601f19603f3d011682016040523d82523d6000602084013e610ea0565b606091505b50825460ff191660021783556040517fbd37c1f0d53bb2f33fe4c2104de272fcdeb4d2fef3acdbf1e4ddc3d6833ca37690610edc908390612ee7565b60405180910390a1505050505050610f21565b600080516020612f798339815191528b8b600386604051610f139493929190612e6c565b60405180910390a150505050505b505050505050565b600082815260016020526040902060020154610f479061056c612079565b610fab5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60448201526e0818591b5a5b881d1bc819dc985b9d608a1b6064820152608401610870565b610fb582826121e6565b5050565b610fc1612079565b6001600160a01b0316816001600160a01b0316146110395760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608401610870565b610fb5828261224f565b61104b6122b8565b60005b838110156110e05784848281811061106857611068612efa565b905060200201602081019061107d9190612917565b6001600160a01b03166108fc84848481811061109b5761109b612efa565b905060200201359081150290604051600060405180830381858888f193505050501580156110cd573d6000803e3d6000fd5b50806110d881612f10565b91505061104e565b5050505050565b6110ef6122b8565b6110f881611ef9565b6002805460ff929092166101000261ff00199092169190911790556040518181527fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8906020015b60405180910390a150565b60008281526001602081815260408084206001600160a01b038616855290920190529020545b92915050565b6000611170600080516020612f99833981519152836114d5565b6111986122b8565b60008581526004602081905260409182902080546001600160a01b0319166001600160a01b038a8116918217909255925163de319d9960e01b8152918201889052861660248201526001600160e01b03198086166044830152606482018590528316608482015287919063de319d999060a401600060405180830381600087803b15801561122557600080fd5b505af1158015611239573d6000803e3d6000fd5b5050505050505050505050565b61124e6122b8565b6000611258612079565b9050816001600160a01b0316816001600160a01b031614156112bc5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f742072656e6f756e6365206f6e6573656c660000000000000000006044820152606401610870565b6112c7600083610f29565b610fb5600082610fb9565b68ffffffffffffffffff831660009081526006602090815260408083208584529091528082208151608081019092528054611370929190829060ff16600481111561131f5761131f612a4a565b600481111561133057611330612a4a565b8152905461010081046001600160c81b03166020830152600160d01b810460ff166040830152600160d81b900464ffffffffff1660609091015283612311565b949350505050565b6000611391600080516020612f99833981519152611bc1565b905090565b61139e6122b8565b6113ae6113a9612079565b612334565b565b6113b86122b8565b6040516307b7ed9960e01b81526001600160a01b0382811660048301528391908216906307b7ed99906024015b600060405180830381600087803b1580156113ff57600080fd5b505af1158015611413573d6000803e3d6000fd5b50505050505050565b60008281526001602052604081206114349083612382565b9392505050565b6114436122b8565b6002546201000090046001600160801b03168114156114a45760405162461bcd60e51b815260206004820152601f60248201527f43757272656e742066656520697320657175616c20746f206e657720666565006044820152606401610870565b6114ad81611f50565b6002806101000a8154816001600160801b0302191690836001600160801b0316021790555050565b60008281526001602052604081206114349083612011565b611505600080516020612f99833981519152826114d5565b6115515760405162461bcd60e51b815260206004820152601f60248201527f6164647220646f65736e277420686176652072656c6179657220726f6c6521006044820152606401610870565b611569600080516020612f9983398151915282611da5565b6040516001600160a01b03821681527f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b9060200161113f565b60408051608081018252600080825260208201819052918101829052606081019190915260ff848116600885901b68ffffffffffffffff0016176000818152600660209081526040808320878452909152908190208151608081019092528054929391929091839116600481111561161c5761161c612a4a565b600481111561162d5761162d612a4a565b8152905461010081046001600160c81b03166020830152600160d01b810460ff166040830152600160d81b900464ffffffffff1660609091015295945050505050565b6116786122b8565b60405163025a3c9960e21b815282906001600160a01b03821690630968f264906113e5908590600401612ee7565b6116ae612180565b6116b6612033565b60008381526004602090815260408083205490516001600160a01b039091169268ffffffffffffffff00600889901b1660ff8a16179290916116fe9185918891889101612ea1565b60408051601f19818403018152828252805160209182012068ffffffffffffffffff861660009081526006835283812082825290925282822060808501909352825490945090929190829060ff16600481111561175d5761175d612a4a565b600481111561176e5761176e612a4a565b8152905461010081046001600160c81b0316602080840191909152600160d01b820460ff16604080850191909152600160d81b90920464ffffffffff1660609093019290925260008a815260049092529020549091506001600160a01b03166118195760405162461bcd60e51b815260206004820152601960248201527f6e6f2068616e646c657220666f72207265736f757263654944000000000000006044820152606401610870565b60028151600481111561182e5761182e612a4a565b141561184c57611843898988888b6001610ca1565b505050506110e0565b6000611856612079565b905060018260000151600481111561187057611870612a4a565b11156118ca5760405162461bcd60e51b815260206004820152602360248201527f70726f706f73616c20616c72656164792065786563757465642f63616e63656c6044820152621b195960ea1b6064820152608401610870565b6118d48282612311565b156119195760405162461bcd60e51b81526020600482015260156024820152741c995b185e595c88185b1c9958591e481d9bdd1959605a1b6044820152606401610870565b60008251600481111561192e5761192e612a4a565b141561198e576040805160808101825260018082526000602083018190528284015264ffffffffff431660608301529151909350600080516020612f7983398151915291611981918d918d918890612e6c565b60405180910390a16119f0565b600254606083015164ffffffffff600160901b9092048216916119b39143911661213e565b64ffffffffff1611156119f0576004808352604051600080516020612f79833981519152916119e7918d918d918890612e6c565b60405180910390a15b600482516004811115611a0557611a05612a4a565b14611ad557611a2a611a168261238e565b83602001516001600160c81b0316176123bc565b6001600160c81b0316602083015260408201805190611a4882612f2b565b60ff1690525081516040517f25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b957434364091611a84918d918d918890612e6c565b60405180910390a1600254604083015160ff6101009092048216911610611ad5576002808352604051600080516020612f7983398151915291611acc918d918d918890612e6c565b60405180910390a15b68ffffffffffffffffff8416600090815260066020908152604080832086845290915290208251815484929190829060ff19166001836004811115611b1c57611b1c612a4a565b021790555060208201518154604084015160609094015164ffffffffff16600160d81b026001600160d81b0360ff909516600160d01b0260ff60d01b196001600160c81b039094166101000293909316610100600160d81b0319909216919091179190911792909216919091179055600282516004811115611ba057611ba0612a4a565b1415611bb557611bb58a8a89898c6000610ca1565b50505050505050505050565b600081815260016020526040812061117090612411565b611be06122b8565b60008281526004602081905260409182902080546001600160a01b0319166001600160a01b038781169182179092559251635c7d1b9b60e11b81529182018590528316602482015284919063b8fa373690604401600060405180830381600087803b158015611c4e57600080fd5b505af1158015611c62573d6000803e3d6000fd5b5050505050505050565b611c84600080516020612f99833981519152826114d5565b15611cd15760405162461bcd60e51b815260206004820152601e60248201527f6164647220616c7265616479206861732072656c6179657220726f6c652100006044820152606401610870565b60c8611cdb611378565b10611d215760405162461bcd60e51b81526020600482015260166024820152751c995b185e595c9cc81b1a5b5a5d081c995858da195960521b6044820152606401610870565b611d39600080516020612f9983398151915282610f29565b6040516001600160a01b03821681527f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c59060200161113f565b611d7a6122b8565b6001600160a01b03919091166000908152600560205260409020805460ff1916911515919091179055565b600082815260016020526040902060020154611dc39061056c612079565b6110395760405162461bcd60e51b815260206004820152603060248201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60448201526f2061646d696e20746f207265766f6b6560801b6064820152608401610870565b611e306122b8565b60ff82166000908152600360205260409020546001600160401b0390811690821611611ead5760405162461bcd60e51b815260206004820152602660248201527f446f6573206e6f7420616c6c6f772064656372656d656e7473206f6620746865604482015265206e6f6e636560d01b6064820152608401610870565b60ff919091166000908152600360205260409020805467ffffffffffffffff19166001600160401b03909216919091179055565b611ee96122b8565b6113ae611ef4612079565b61241b565b60006101008210611f4c5760405162461bcd60e51b815260206004820152601c60248201527f76616c756520646f6573206e6f742066697420696e20382062697473000000006044820152606401610870565b5090565b6000600160801b8210611f4c5760405162461bcd60e51b815260206004820152601e60248201527f76616c756520646f6573206e6f742066697420696e20313238206269747300006044820152606401610870565b6000650100000000008210611f4c5760405162461bcd60e51b815260206004820152601d60248201527f76616c756520646f6573206e6f742066697420696e20343020626974730000006044820152606401610870565b6000611434836001600160a01b038416612466565b6001600160a01b03811660009081526001830160205260408120541515611434565b60005460ff16156113ae5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610870565b600033601436108015906120a557506001600160a01b03811660009081526005602052604090205460ff165b156120b5575060131936013560601c5b919050565b60006120c4612079565b90506120d16000826114d5565b806120ef57506120ef600080516020612f99833981519152826114d5565b61213b5760405162461bcd60e51b815260206004820152601e60248201527f73656e646572206973206e6f742072656c61796572206f722061646d696e00006044820152606401610870565b50565b600061143483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506124b5565b61219a600080516020612f9983398151915261056c612079565b6113ae5760405162461bcd60e51b815260206004820181905260248201527f73656e64657220646f65736e277420686176652072656c6179657220726f6c656044820152606401610870565b60008281526001602052604090206121fe9082611ffc565b15610fb55761220b612079565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600082815260016020526040902061226790826124ef565b15610fb557612274612079565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b6122c5600061056c612079565b6113ae5760405162461bcd60e51b815260206004820152601e60248201527f73656e64657220646f65736e277420686176652061646d696e20726f6c6500006044820152606401610870565b60008083602001516001600160c81b031661232b8461238e565b16119392505050565b61233c612033565b6000805460ff191660011790556040516001600160a01b03821681527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589060200161113f565b60006114348383612504565b60006123b26123ab600080516020612f998339815191528461114a565b600161213e565b6001901b92915050565b6000600160c81b8210611f4c5760405162461bcd60e51b815260206004820152601e60248201527f76616c756520646f6573206e6f742066697420696e20323030206269747300006044820152606401610870565b6000611170825490565b61242361252e565b6000805460ff191690556040516001600160a01b03821681527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9060200161113f565b60008181526001830160205260408120546124ad57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155611170565b506000611170565b600081848411156124d95760405162461bcd60e51b81526004016108709190612ee7565b5060006124e68486612f4b565b95945050505050565b6000611434836001600160a01b038416612577565b600082600001828154811061251b5761251b612efa565b9060005260206000200154905092915050565b60005460ff166113ae5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610870565b6000818152600183016020526040812054801561266057600061259b600183612f4b565b85549091506000906125af90600190612f4b565b90508181146126145760008660000182815481106125cf576125cf612efa565b90600052602060002001549050808760000184815481106125f2576125f2612efa565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061262557612625612f62565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050611170565b6000915050611170565b803560ff811681146120b557600080fd5b60008083601f84011261268d57600080fd5b5081356001600160401b038111156126a457600080fd5b6020830191508360208285010111156126bc57600080fd5b9250929050565b600080600080606085870312156126d957600080fd5b6126e28561266a565b93506020850135925060408501356001600160401b0381111561270457600080fd5b6127108782880161267b565b95989497509550505050565b80356001600160401b03811681146120b557600080fd5b60008060006060848603121561274857600080fd5b6127518461266a565b925061275f6020850161271c565b9150604084013590509250925092565b803580151581146120b557600080fd5b60008060008060008060a0878903121561279857600080fd5b6127a18761266a565b95506127af6020880161271c565b945060408701356001600160401b038111156127ca57600080fd5b6127d689828a0161267b565b909550935050606087013591506127ef6080880161276f565b90509295509295509295565b60006020828403121561280d57600080fd5b5035919050565b6001600160a01b038116811461213b57600080fd5b6000806040838503121561283c57600080fd5b82359150602083013561284e81612814565b809150509250929050565b60008083601f84011261286b57600080fd5b5081356001600160401b0381111561288257600080fd5b6020830191508360208260051b85010111156126bc57600080fd5b600080600080604085870312156128b357600080fd5b84356001600160401b03808211156128ca57600080fd5b6128d688838901612859565b909650945060208701359150808211156128ef57600080fd5b5061271087828801612859565b60006020828403121561290e57600080fd5b6114348261266a565b60006020828403121561292957600080fd5b813561143481612814565b80356001600160e01b0319811681146120b557600080fd5b60008060008060008060c0878903121561296557600080fd5b863561297081612814565b955060208701359450604087013561298781612814565b935061299560608801612934565b9250608087013591506127ef60a08801612934565b6000806000606084860312156129bf57600080fd5b833568ffffffffffffffffff811681146129d857600080fd5b92506020840135915060408401356129ef81612814565b809150509250925092565b60008060408385031215612a0d57600080fd5b8235612a1881612814565b9150602083013561284e81612814565b60008060408385031215612a3b57600080fd5b50508035926020909101359150565b634e487b7160e01b600052602160045260246000fd5b60058110612a7e57634e487b7160e01b600052602160045260246000fd5b9052565b6000608082019050612a95828451612a60565b60018060c81b03602084015116602083015260ff604084015116604083015264ffffffffff606084015116606083015292915050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715612b0957612b09612acb565b604052919050565b60006001600160401b03821115612b2a57612b2a612acb565b50601f01601f191660200190565b60008060408385031215612b4b57600080fd5b8235612b5681612814565b915060208301356001600160401b03811115612b7157600080fd5b8301601f81018513612b8257600080fd5b8035612b95612b9082612b11565b612ae1565b818152866020838501011115612baa57600080fd5b816020840160208301376000602083830101528093505050509250929050565b600080600080600060808688031215612be257600080fd5b612beb8661266a565b9450612bf96020870161271c565b93506040860135925060608601356001600160401b03811115612c1b57600080fd5b612c278882890161267b565b969995985093965092949392505050565b600080600060608486031215612c4d57600080fd5b83356129d881612814565b60008060408385031215612c6b57600080fd5b8235612c7681612814565b9150612c846020840161276f565b90509250929050565b60008060408385031215612ca057600080fd5b612ca98361266a565b9150612c846020840161271c565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0380831681811415612cea57612cea612cb7565b6001019392505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8481526001600160a01b0384166020820152606060408201819052600090612d489083018486612cf4565b9695505050505050565b60005b83811015612d6d578181015183820152602001612d55565b83811115612d7c576000848401525b50505050565b600060208284031215612d9457600080fd5b81516001600160401b03811115612daa57600080fd5b8201601f81018413612dbb57600080fd5b8051612dc9612b9082612b11565b818152856020838501011115612dde57600080fd5b6124e6826020830160208601612d52565b60008151808452612e07816020860160208601612d52565b601f01601f19169290920160200192915050565b60ff871681528560208201526001600160401b038516604082015260a060608201526000612e4d60a083018587612cf4565b8281036080840152612e5f8185612def565b9998505050505050505050565b60ff851681526001600160401b038416602082015260808101612e926040830185612a60565b82606083015295945050505050565b6bffffffffffffffffffffffff198460601b168152818360148301376000910160140190815292915050565b8381526040602082015260006124e6604083018486612cf4565b6020815260006114346020830184612def565b634e487b7160e01b600052603260045260246000fd5b6000600019821415612f2457612f24612cb7565b5060010190565b600060ff821660ff811415612f4257612f42612cb7565b60010192915050565b600082821015612f5d57612f5d612cb7565b500390565b634e487b7160e01b600052603160045260246000fdfe968626a768e76ba1363efe44e322a6c4900c5f084e0b45f35e294dfddaa9e0d5e2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4a2646970667358221220b709f2194f4c148f513f57a45e4ba5a064bdbf9c8e2385faecd3a9e498f2bc3164736f6c634300080b0033"

type ContractTestSuite struct {
	suite.Suite
	gomockController *gomock.Controller
	mockClient       *mock.MockClient
	mockTransactor   *mock.MockTransactor
	contract         Contract
}

func TestRunContractTestSuite(t *testing.T) {
	suite.Run(t, new(ContractTestSuite))
}

func (s *ContractTestSuite) SetupSuite()    {}
func (s *ContractTestSuite) TearDownSuite() {}
func (s *ContractTestSuite) SetupTest() {
	s.gomockController = gomock.NewController(s.T())
	s.mockTransactor = mock.NewMockTransactor(s.gomockController)
	s.mockClient = mock.NewMockClient(s.gomockController)
	a, _ := abi.JSON(strings.NewReader(testABI))
	b := common.FromHex(testBin)
	s.contract = NewContract(
		common.Address{}, a, b, s.mockClient, s.mockTransactor,
	)
}
func (s *ContractTestSuite) TearDownTest() {}

func (s *ContractTestSuite) TestContract_PackMethod_ValidRequest_Success() {
	res, err := s.contract.PackMethod("approve", common.Address{}, big.NewInt(10))
	s.Equal(
		common.Hex2Bytes("095ea7b30000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a"),
		res,
	)
	s.Nil(err)
}

func (s *ContractTestSuite) TestContract_PackMethod_InvalidRequest_Fail() {
	res, err := s.contract.PackMethod("invalid_method", common.Address{}, big.NewInt(10))
	s.Equal([]byte{}, res)
	s.NotNil(err)
}

func (s *ContractTestSuite) TestContract_UnpackResult_InvalidRequest_Fail() {
	rawInvalidApproveData := common.Hex2Bytes("095ea7b30000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a")
	res, err := s.contract.UnpackResult("approve", rawInvalidApproveData)
	s.NotNil(err)
	s.Nil(res)
}

func (s *ContractTestSuite) TestContract_ExecuteTransaction_ValidRequest_Success() {
	s.mockTransactor.EXPECT().Transact(
		&common.Address{},
		gomock.Any(),
		transactor.TransactOptions{},
	).Return(&common.Hash{}, nil)
	hash, err := s.contract.ExecuteTransaction(
		"approve",
		transactor.TransactOptions{}, common.Address{}, big.NewInt(10),
	)
	s.Nil(err)
	s.NotNil(hash)
}

func (s *ContractTestSuite) TestContract_ExecuteTransaction_TransactError_Fail() {
	s.mockTransactor.EXPECT().Transact(
		&common.Address{},
		gomock.Any(),
		transactor.TransactOptions{},
	).Return(nil, errors.New("error"))
	hash, err := s.contract.ExecuteTransaction(
		"approve",
		transactor.TransactOptions{}, common.Address{}, big.NewInt(10),
	)
	s.Nil(hash)
	s.Error(err, "error")
}

func (s *ContractTestSuite) TestContract_ExecuteTransaction_InvalidRequest_Fail() {
	hash, err := s.contract.ExecuteTransaction(
		"approve",
		transactor.TransactOptions{}, common.Address{}, // missing one argument
	)
	s.Nil(hash)
	s.Error(err, "error")
}

func (s *ContractTestSuite) TestContract_CallContract_CallContractError_Fail() {
	s.mockClient.EXPECT().CallContract(
		gomock.Any(),
		gomock.Any(),
		nil,
	).Return(nil, errors.New("error"))
	s.mockClient.EXPECT().From().Times(1).Return(common.Address{})

	res, err := s.contract.CallContract("ownerOf", big.NewInt(0))
	if err != nil {
		return
	}
	s.Nil(res)
	s.Error(err, "error")
}

func (s *ContractTestSuite) TestContract_CallContract_InvalidRequest_Fail() {
	res, err := s.contract.CallContract("invalidMethod", big.NewInt(0))
	if err != nil {
		return
	}
	s.Nil(res)
	s.Error(err, "error")
}

func (s *ContractTestSuite) TestContract_CallContract_MissingContract_Fail() {
	s.mockClient.EXPECT().CallContract(
		gomock.Any(),
		gomock.Any(),
		nil,
	).Return(nil, errors.New("error"))
	s.mockClient.EXPECT().From().Times(1).Return(common.Address{})
	res, err := s.contract.CallContract("ownerOf", big.NewInt(0))
	if err != nil {
		return
	}
	s.Nil(res)
	s.Error(err, "error")
}
