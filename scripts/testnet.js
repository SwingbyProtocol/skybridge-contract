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


const Tokens = {
    [ropsten]: {
        WETH: {
            address: '0xc778417e063141139fce010982780140aa0cd5ab',
            decimals: 18
        },
        WBTC: {
            address: '0x442be68395613bdcd19778e761f03261ec46c06d',
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
        BAT: {
            address: '0x0d8775f648430679a709e98d2b0cb6250d2887ef',
            decimals: 18
        }
    },
}

const srcAmountETH = "1000000000000000000"//1 ETHER
const srcAmountBTC = "10000000"//0.1 BTC


//const srcAmount = "147618252344340533"
//UNI Amount: 147618252344340533
//UNI Address: 0x1f9840a85d5af5bf1d1762f925bdaddc4201f984


async function main() {
    let data
    let getPrice = await paraswap.getPrice(
        Tokens[mainnet]['WBTC'],
        Tokens[mainnet]['ETH'],
        srcAmountBTC,
        mainnet
    )

    //console.log(getPrice)

    const slippage = (decimals) => {
        return new BigNumber.from(3).mul(new BigNumber.from(10).pow(decimals - 2)).toString() //Format ERC20 - 0.05
    }

    let decimals = Tokens[mainnet]['UNI'].decimals
    let minDestAmount = new BigNumber.from(getPrice.price).sub(slippage(decimals))


    

    //POST request - build TX data to send to contract
    const txRequest = await paraswap.buildTransaction(
        getPrice.payload,
        Tokens[mainnet]['WBTC'],
        Tokens[mainnet]['ETH'],
        srcAmountBTC,
        minDestAmount.toString(),
        mainnet,
        //"0x202CCe504e04bEd6fC0521238dDf04Bc9E8E15aB", //SWAP contract
        "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC", //user1.address
        true //only params - true for contract -> contract | false for standard transaction object
    )
    data = txRequest.data //params to execute transaction contract -> contract  
    const output = txRequest.config.data
    const { parse } = require('json-parser')
    const parsedOutput = parse(output)
    const contractMethod = parsedOutput.priceRoute.contractMethod


    console.log("Recomended Contract Method:", contractMethod)
    //console.log(data)

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





    /**
     //simpleSwap
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
     */


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



}


main()


















//SWAP: 0xf73D63C3eB97389cB5A28C4aD5e8AC428cb16417


//DAI Amount: 548824612259226186639
//DAI Address: 0xaD6D458402F60fD3Bd25163575031ACDce07538D

//UNI Amount: 147618252344340533
//UNI Address: 0x1f9840a85d5af5bf1d1762f925bdaddc4201f984