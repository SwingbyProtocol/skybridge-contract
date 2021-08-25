const axios = require('axios');
const { BigNumber, Ethers } = require('ethers')

//call the class with this
const paraURL = "https://apiv4.paraswap.io/v2"

class ParaSwap {
    
    constructor(apiURL = paraURL) {
        this.apiURL = apiURL
        this.referrer = 'SkyPools'
    }
    async getPrice(from, to, srcAmount, network) {
        try {
            const requestURL =
                `${this.apiURL}/prices/?from=${from.address}&to=${to.address}` +
                `&amount=${srcAmount}&fromDecimals=${from.decimals}&toDecimals` +
                `=${to.decimals}&side=SELL&network=${network}`
            const { data } = await axios.get(requestURL, {
                headers: {
                    'X-Partner': this.referrer,
                },
            })
            return {
                price: data.priceRoute.destAmount,
                payload: data.priceRoute,
            }
        } catch (e) {
            throw new Error(
                `Paraswap unable to fetch TEST price ${from.address} ${to.address} ${network} ${e.message}`
            )
        }
    }

    async buildTransaction(
        pricePayload,
        from,
        to,
        srcAmount,
        minDestAmount,
        network,
        userAddress
    ) {
        try {
            const requestURL = `${this.apiURL}/transactions/${network}`
            const requestData = {
                priceRoute: pricePayload,
                srcToken: from.address,
                destToken: to.address,
                srcAmount: srcAmount,
                destAmount: minDestAmount,
                userAddress: userAddress,
                referrer: this.referrer,
                srcDecimals: from.decimals,
                destDecimals: to.decimals,
            }

            const { data } = await axios.post(requestURL, requestData)
            return {
                from: data.from,
                to: data.to,
                data: data.data,
                gasLimit: '0x' + new BigNumber.from(data.gas),
                value: '0x' + new BigNumber.from(data.value)
            }
        } catch (e) {
            throw new Error(
                `Paraswap unableo to build transaction ${from.address} ${to.address} ${network} ${e.message}`,
            )
        }
    }
}

module.exports = ParaSwap;