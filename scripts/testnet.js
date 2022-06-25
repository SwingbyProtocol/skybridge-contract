const { BigNumber } = require("@ethersproject/contracts/node_modules/@ethersproject/bignumber")
const hre = require("hardhat")
const { Ethers } = require('ethers')
const ParaSwap = require('../scripts/paraswap.js')
const paraV5 = "https://apiv5.paraswap.io"
const paraswap = new ParaSwap(paraV5, 5)
const ropsten = 3 //ropsten
const mainnet = 1
const swapAddress = '0xf73D63C3eB97389cB5A28C4aD5e8AC428cb16417'
const swapABI = require('../abi/SwapContract.json')
const { default: Web3 } = require('web3');

//npx hardhat run testnet.js

const Tokens = {
    [ropsten]: {
        WETH: {
            address: '0xc778417e063141139fce010982780140aa0cd5ab',
            decimals: 18
        },
        WBTC: {
            //address: '0x442be68395613bdcd19778e761f03261ec46c06d',
            address: '0x7cb2eac36b4bb7c36640f32e806d33e474d1d427',
            decimals: 8
        },
        DAI: {
            address: '0xaD6D458402F60fD3Bd25163575031ACDce07538D',
            decimals: 18
        },
        UNI: {
            address: '0x1f9840a85d5af5bf1d1762f925bdaddc4201f984',
            decimals: 18
        },
    },
    [mainnet]: {
        ETH: {
            address: '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE',
            decimals: 18,
        },
        WETH: {
            address: '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
            decimals: 18
        },
        USDC: {
            address: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
            decimals: 6,
        },
        MATIC: {
            address: '0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0',
            decimals: 18,
        },
        WBTC: {
            address: '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
            decimals: 8
        },
        UNI: {
            address: '0x1f9840a85d5af5bf1d1762f925bdaddc4201f984',
            decimals: 18
        },
        COMP: {
            address: '0xc00e94cb662c3520282e6f5717214004a7f26888',
            decimals: 18
        },
        MKR: {
            address: '0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2',
            decimals: 18
        },
        SWINGBY: {
            address: '0x8287c7b963b405b7b8d467db9d79eec40625b13a',
            decimals: 18
        }
    },
}

const srcAmountBTC = "10000000"//0.1 BTC //USE FOR FLOW 1
const srcAmountETH = "1000000000000000000"//1 ETHER //USE FOR FLOW 2
const smallSrcAmountETH = "100000000000000000"//0.1 ETHER //USE FOR FLOW 2

const giantAmountFlow2 = "15000000000000000000" // 15 ETH



//const srcAmount = "147618252344340533"
//UNI Amount: 147618252344340533
//UNI Address: 0x1f9840a85d5af5bf1d1762f925bdaddc4201f984


async function main() {
    let data
    let getPrice = await paraswap.getPrice(
        Tokens[ropsten]['WETH'], // From token - CHANGE THIS
        Tokens[ropsten]['WBTC'], // To token - CHANGE THIS
        srcAmountETH, //Change this depending on flow 1 vs flow 2
        ropsten
    )

    //console.log(getPrice)

    const slippage = (decimals) => {
        return new BigNumber.from(3).mul(new BigNumber.from(10).pow(decimals - 2)).toString() //Format ERC20 - 0.05
    }

    let decimals = Tokens[ropsten]['WETH'].decimals
    let minDestAmount = new BigNumber.from(getPrice.price).sub(slippage(decimals))

    //POST request - build TX data to send to contract
    const txRequest = await paraswap.buildTransaction(
        getPrice.payload, //data from GET request
        Tokens[ropsten]['WETH'], // From token - CHANGE THIS
        Tokens[ropsten]['WBTC'], // To token - CHANGE THIS
        srcAmountETH, //Change this depending on flow 1 vs flow 2
        minDestAmount.toString(), //this param is not used for paraswap V5 anymore, redundant
        ropsten,
        "0x202CCe504e04bEd6fC0521238dDf04Bc9E8E15aB", //SWAP contract - flow 2 simpleSwap - uniswap functions don't care about this param
        //"0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC", //user1.address - Flow 1 simpleSwap
        true //only params - true for contract -> contract | false for standard transaction object
    )
    data = txRequest.data //params to execute transaction contract -> contract  
    const output = txRequest.config.data
    const { parse } = require('json-parser')
    const parsedOutput = parse(output)
    const contractMethod = parsedOutput.priceRoute.contractMethod


    console.log("Recomended Contract Method:", contractMethod)
    //console.log(data)

    if (contractMethod == "simpleSwap") {
        console.log(`Getting data for ${contractMethod}`)
        const dataArray = [
            data[0].fromToken,
            data[0].toToken,
            data[0].fromAmount,
            data[0].toAmount,
            data[0].expectedAmount,
            data[0].callees,
            data[0].exchangeData,
            data[0].startIndexes,
            data[0].values,
            data[0].beneficiary,
            data[0].partner,
            data[0].feePercent,
            data[0].permit,
            data[0].deadline,
            data[0].uuid
        ]
        console.log(dataArray)
    } else if (contractMethod == "swapOnUniswap") {
        console.log(`Getting data for ${contractMethod}`)
        console.log(data)
    } else if (contractMethod == "swapOnUniswapFork") {
        console.log(`Getting data for ${contractMethod}`)
        console.log(data)
    }



    /**
     //megaSwap
     let routeArray = [
        data[0].path[0].path[0].adapters[0].route[0].index,
        data[0].path[0].path[0].adapters[0].route[0].targetExchange,
        data[0].path[0].path[0].adapters[0].route[0].percent,
        data[0].path[0].path[0].adapters[0].route[0].payload,
        data[0].path[0].path[0].adapters[0].route[0].networkFee
    ]
    let pathArray = [
        data[0].path[0].fromAmountPercent,
        [
            data[0].path[0].path[0].to,
            data[0].path[0].path[0].totalNetworkFee,
            [ //adapters
                data[0].path[0].path[0].adapters[0].adapter,
                data[0].path[0].path[0].adapters[0].percent,
                data[0].path[0].path[0].adapters[0].networkFee,
                routeArray
            ]
        ]
    ]
    const megaDataArray = [
        data[0].fromToken,
        data[0].fromAmount,
        data[0].toAmount,
        data[0].expectedAmount,
        data[0].beneficiary,
        pathArray,
        data[0].partner,
        data[0].feePercent,
        data[0].permit,
        data[0].deadline,
        data[0].uuid
    ]

    /**
     console.log(dataArray)

    console.log("PATH BELOW")
    console.log(pathArray)

    console.log("ROUTE BELOW")
    console.log(routeArray)
     */



    //console.log(data[0].path[0].path[0].adapters[0].route[0].targetExchange)

    //console.log(dataArray)
    //console.log(data[0].path)
    //console.log(data[0].path[0].path)
    //console.log(data[0].path[0].path[0].adapters)
    //console.log(data[0].path[0].path[0].adapters[0].route)

}


main()


















//SWAP: 0xf73D63C3eB97389cB5A28C4aD5e8AC428cb16417


//DAI Amount: 548824612259226186639
//DAI Address: 0xaD6D458402F60fD3Bd25163575031ACDce07538D

//UNI Amount: 147618252344340533
//UNI Address: 0x1f9840a85d5af5bf1d1762f925bdaddc4201f984