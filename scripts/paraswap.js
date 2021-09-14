const axios = require('axios');
const { BigNumber, Ethers } = require('ethers')

//call the class with this
const paraURL = "https://apiv4.paraswap.io/v2"
const paraV5 = "https://apiv5.paraswap.io/"

class ParaSwap {

    constructor(apiURL, version) {
        this.apiURL = apiURL
        this.version = version
        this.referrer = 'SkyPools'
    }


    async getPrice(from, to, srcAmount, network) {
        if (this.version == 4) {
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
                    if (e.response != undefined) {
                        console.log("ERROR RESPONSE: ", e.response)
                    } else { console.log("ERROR: ", e) }
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
        } else if (this.version == 5) {
            try {
                const requestURL =
                    `${this.apiURL}/prices/?srcToken=${from.address}&destToken=${to.address}` +
                    `&amount=${srcAmount}&srcDecimals=${from.decimals}&destDecimals` +
                    `=${to.decimals}&side=SELL&network=${network}`
                    //+ `&excludeContractMethods=multiSwap`
                    +  `&includeContractMethods=simpleSwap`
                    const { data } = await axios.get(requestURL, {
                    headers: {
                        'partner': this.referrer,
                    },
                }).catch(e => {
                    if (e.response != undefined) {
                        console.log("ERROR RESPONSE: ", e.response)
                    } else { console.log("ERROR: ", e) }
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

    }


    async buildTransaction(
        pricePayload,
        from,
        to,
        srcAmount,
        minDestAmount,
        network,
        userAddress,
        onlyParams
    ) {
        if (this.version == 4) {
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
                srcToken: from.address,
                receiver: userAddress
            }
            let data = "No Response"
            await axios.post(
                requestURL,
                requestData
            ).then(response => {
                //console.log("RESPONSE: ", response)
                data = response
            }).catch(e => {
                if (e.response != undefined) {
                    console.log("ERROR RESPONSE: ", e.response)
                    data = e.response
                } else { console.log("ERROR: ", e) }
            })
            return data
        } else if (this.version == 5) {
            const requestURL = `${this.apiURL}/transactions/${network}?skipChecks=true&onlyParams=${onlyParams}`;
            const requestData = {
                destDecimals: to.decimals,
                srcDecimals: from.decimals,
                partner: this.referrer,
                userAddress: userAddress,
                destAmount: minDestAmount,
                priceRoute: pricePayload,
                srcAmount: srcAmount,
                destToken: to.address,
                srcToken: from.address,
                receiver: userAddress
            }
            let data = "No Response"
            await axios.post(
                requestURL,
                requestData
            ).then(response => {
                //console.log("RESPONSE: ", response)
                data = response
            }).catch(e => {
                if (e.response != undefined) {
                    console.log("ERROR RESPONSE: ", e.response)
                    data = e.response
                } else { console.log("ERROR: ", e) }
            })
            return data
        }
    }

}

module.exports = ParaSwap;