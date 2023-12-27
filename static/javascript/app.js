// Web3js 合约方法部分
let todayNum = $('#todayCount').val()
if (todayNum < 10) {
    $('#oneMin').bind('click', () => {
        let aim = $("#aimAddress").val();
        myContract.methods.withdrawOfOneMin().send({
            from: aim
        })
            .on('receipt', function (receipt) {
                console.log(receipt)
            })
            .on('error', function (error, receipt) {
                if (error == "Error: Returned error: execution reverted"){
                    alert("功能冷却中,请稍后再试")
                }else {
                    console.log(error)
                    alert(error)
                }
            });
    })

    $('#threeMin').bind('click', () => {
        let aim = $("#aimAddress").val();
        myContract.methods.withdrawOfThreeMin().send({
            from: aim
        })
            .on('receipt', function (receipt) {
                console.log(receipt)
            })
            .on('error', function (error, receipt) {
                console.log(error)
                alert("功能冷却中,请稍后再试")
            });
    })

    $('#nineMin').bind('click', () => {
        let aim = $("#aimAddress").val();
        myContract.methods.withdrawOfNineMin().send({
            from: aim
        })
            .on('receipt', function (receipt) {
                console.log(receipt)
            })
            .on('error', function (error, receipt) {
                console.log(error)
                alert("功能冷却中,请稍后再试")
            });
    })
} else {
    $("#showList").click(() => {
        alert("本日以太币发放完毕,请明日再来")
    })
}


myContract.events.Withdraw({
    filter: {},
    fromBlock: 0
}, (error, event) => {
    if (error) {
        console.log(error)
        return
    }
    let aim = $("#aimAddress").val();
    if (aim != "") {
        web3.eth.sendTransaction({
            from: sourceAddress,
            to: aim,
            value: event.returnValues.award
        })
            .on('receipt', function (receipt) {
                console.log(receipt)
            })
            .on('error', function (error, receipt) {
                console.log(error)
            });
        $.ajax({
            url: "/dapp/transaction",
            type: "POST",
            data: {
                "ToAddress": event.returnValues.sender,
                "Award": event.returnValues.award / 1000000000000000000,
                "GetTime": event.returnValues.userLimit,
            },
            success: (data) => {
                alert("提取成功")
            }
        })
    }
})