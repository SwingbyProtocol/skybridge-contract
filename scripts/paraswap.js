const axios = require('axios');
const { BigNumber, Ethers } = require('ethers')

//call the class with this
const paraURL = "https://apiv4.paraswap.io/v2"

class ParaSwap {

    constructor(apiURL = paraURL) {
        this.apiURL = apiURL
        this.referrer = 'SkyPools'
    }
    async oldGetPrice(from, to, srcAmount, network) {
        try {
            const requestURL =
                `${this.apiURL}/prices/?from=${from.address}&to=${to.address}` +
                `&amount=${srcAmount}&fromDecimals=${from.decimals}&toDecimals` +
                `=${to.decimals}&side=SELL&network=${network}`
            const { data } = await axios.get(requestURL, {
                headers: {
                    'X-Partner': this.referrer,
                },
            }).catch(e => {
                if(e.response != undefined){
                    console.log("ERROR RESPONSE: ", e.response)
                }else {  console.log("ERROR: ", e) }
            })
            return {
                price: data.priceRoute.destAmount,
                payload: data.priceRoute,
            }
        } catch (e) {
            throw new Error(
                `Paraswap unable to fetch price ${from.address} ${to.address} ${network} ${e.message}`
            )
        }
    }

    async getPrice(from, to, srcAmount, network) {
        
            const requestURL =
                `${this.apiURL}/prices/?from=${from.address}&to=${to.address}` +
                `&amount=${srcAmount}&fromDecimals=${from.decimals}&toDecimals` +
                `=${to.decimals}&side=SELL&network=${network}`
            let data = "No Response"
            await axios.get(requestURL, {
                headers: {
                    'X-Partner': this.referrer,
                },
            }).then(response => {
                data = response
            }).catch(e => {
                if(e.response != undefined){
                    console.log("ERROR RESPONSE: ", e.response)
                }else {  console.log("ERROR: ", e) }
            })

            if (data != 'No Response'){
                return {
                    data
                }
            }else {return data}

    }


    async buildTransaction (
        pricePayload,
        from,
        to,
        srcAmount,
        minDestAmount,
        network,
        userAddress,
        onlyParams
    ) {
        const requestURL = `${this.apiURL}/transactions/${network}?skipChecks=true&onlyParams=${onlyParams}`;
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
            if(e.response != undefined){
                console.log("ERROR RESPONSE: ", e.response)
                data = e.response
            }else { console.log("ERROR: ", e) }
        })
        return data
    }    
}

module.exports = ParaSwap;