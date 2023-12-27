const web3 = new Web3('ws://127.0.0.1:8546');
//const web3 = new Web3(Web3.givenProvider)
console.log("web3", web3);

const contractAddress = "0x8aC338520d49F49c04AF4C2e2d0677F791bCCa5D";
const sourceAddress = "0xAa152bba2885BC2c746a2d7804FfaB821d957aa7";

//web3.eth.perconal.newAccount(function(error, result){ .....  })
//web3.eth.getBalance("0xaa152bba2885bc2c746a2d7804ffab821d957aa7").then(console.log);
const abi = [
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "sender",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "award",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "userLimit",
				"type": "uint256"
			}
		],
		"name": "Withdraw",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "timeLimt",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "withdrawOfNineMin",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "withdrawOfOneMin",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "withdrawOfThreeMin",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]

var myContract = new web3.eth.Contract(abi, contractAddress);