# -*- coding: utf-8 -*-
import time

from api.bi import api_key, api_secret
from api.bi.bi_api import BinanceAPI
from api.messager.dingding import Message
from strategic.long_it.run_bid_data import RunLongPlan

binan = BinanceAPI(api_key, api_secret)
runner = RunLongPlan()
msg = Message()


##TODO: change CLI mode
class Run_Main():

    def __init__(self):
        self.coinType = runner.get_coin_pair()  # 交易币种
        pass

    def loop_run(self):
        while True:
            cur_market_price = binan.get_ticker_price(runner.get_coin_pair())  # 当前交易对市价
            grid_buy_price = runner.get_buy_price()  # 当前网格买入价格
            grid_sell_price = runner.get_sell_price()  # 当前网格卖出价格
            quantity = runner.get_quantity()  # 买入量
            step = runner.get_step()  # 当前步数

            if grid_buy_price >= cur_market_price:  # 是否满足买入价
                res = msg.buy_limit_msg(self.coinType, quantity, grid_buy_price)
                if res['orderId']:  # 挂单成功
                    runner.modify_price(grid_buy_price, step + 1)  # 修改data.json中价格、当前步数
                    time.sleep(60 * 2)  # 挂单后，停止运行1分钟
                else:
                    break

            elif grid_sell_price < cur_market_price:  # 是否满足卖出价
                if step == 0:  # setp=0 防止踏空，跟随价格上涨
                    runner.modify_price(grid_sell_price, step)
                else:
                    res = msg.sell_limit_msg(self.coinType, runner.get_quantity(False), grid_sell_price)
                    if res['orderId']:
                        runner.modify_price(grid_sell_price, step - 1)
                        time.sleep(60 * 2)  # 挂单后，停止运行1分钟
                    else:
                        break
            else:
                print("当前市价：{market_price}。未能满足交易,继续运行".format(market_price=cur_market_price))


if __name__ == "__main__":
    instance = Run_Main()
    try:
        instance.loop_run()
    except Exception as e:
        error_info = "报警：币种{coin},服务停止.错误原因{info}".format(coin=instance.coinType, info=str(e))
        msg.dingding_warn(error_info)

# 调试看报错运行下面，正式运行用上面
# if __name__ == "__main__":
# instance = Run_Main()
# instance.loop_run()
