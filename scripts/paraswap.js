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



    async buildTransaction (
        pricePayload,
        from,
        to,
        srcAmount,
        minDestAmount,
        network,
        userAddress
    ) {
        const requestURL = `${this.apiURL}/transactions/${network}?skipChecks=true&onlyParams=true`;
        const requestData = {
            toDecimals: to.decimals,
            fromDecimals: from.decimals,
            referrer: this.referrer,
            userAddress: userAddress,
            destAmount: minDestAmount,
            priceRoute: pricePayload,
            srcAmount: srcAmount,
            destToken: to.address,
            srcToken: from.address
        }
        let data = "No Response"
        await axios.post(
            requestURL,
            requestData
        ).then(response => {
            //console.log("RESPONSE: ", response)
            data = response            
        }).catch(e => {
            console.log("ERROR: ", e)
            if(e.response != undefined){
                console.log("ERROR RESPONSE: ", e.response)
            }
        })
        return data
    }    
}

module.exports = ParaSwap;