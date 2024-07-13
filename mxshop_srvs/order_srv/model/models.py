from datetime import datetime
from peewee import *
from playhouse.pool import PooledMySQLDatabase
from playhouse.shortcuts import ReconnectMixin

from order_srv.settings import settings


# class ReconnectMySQLDatabase(ReconnectMixin, PooledMySQLDatabase):
#     pass
#
# db = ReconnectMySQLDatabase("mxshop_order_srv", host="127.0.0.1", port=3306, user="root", password="root")


class BaseModel(Model):
    add_time = DateTimeField(default=datetime.now, verbose_name="添加时间")
    update_time = DateTimeField(default=datetime.now, verbose_name="更新时间")
    is_deleted = BooleanField(default=False, verbose_name="是否删除")

    class Meta:
        database = settings.DB

    def save(self, *args, **kwargs):
        if self._pk is not None:
            self.update_time = datetime.now()
        return super().save(*args, **kwargs)

    @classmethod
    def delete(cls, permanently=False):
        if permanently:
            return super().delete()
        return super().update(is_deleted=True, update_time=datetime.now())

    def delete_instance(self, permanently=False, recursive=False, delete_nullable=False):
        if permanently:
            return self.delete(permanently).where(self._pk_expr()).execute()
        else:
            self.is_deleted = True
            self.update_time = datetime.now()
            self.save()

    @classmethod
    def select(cls, *fields):
        return super().select(*fields).where(cls.is_deleted == False)


class ShoppingCart(BaseModel):
    """购物车"""
    user = IntegerField(verbose_name="用户ID")
    goods = IntegerField(verbose_name="商品ID")
    nums = IntegerField(verbose_name="购买数量")
    checked = BooleanField(default=True, verbose_name="是否选中")

    class Meta:
        # 绑定上面的数据库实例
        database = settings.DB
        # 设置表名
        table_name = "shopping_cart"


class OrderInfo(BaseModel):
    """
    订单
    """
    ORDER_STATUS = (
        ("TRADE_SUCCESS", "成功"),
        ("TRADE_CLOSED", "超时关闭"),
        ("WAIT_BUYER_PAY", "交易创建"),
        ("TRADE_FINISHED", "交易结束"),
    )
    PAY_TYPE = (
        ("alipay", "支付宝"),
        ("wechat", "微信"),
    )
    user = IntegerField(verbose_name="用户Id")
    order_sn = CharField(max_length=30, null=True, unique=True, verbose_name="订单号")
    pay_type = CharField(choices=PAY_TYPE, default="wechat", max_length=30, verbose_name="支付方式")
    status = CharField(choices=ORDER_STATUS, default="paying", max_length=30, verbose_name="订单状态")
    trade_no = CharField(max_length=100, unique=True, null=True, verbose_name="交易号")  # 支付宝交易号
    order_amount = FloatField(default=0.0, verbose_name="订单金额")
    pay_time = DateTimeField(null=True, verbose_name="支付时间")

    # 用户信息
    address = CharField(max_length=100, default="", verbose_name="收货地址")
    signer_name = CharField(max_length=20, default="", verbose_name="签收人")
    signer_mobile = CharField(max_length=11, verbose_name="联系电话")
    post = CharField(max_length=20, default="", verbose_name="留言")

    class Meta:
        # 绑定上面的数据库实例
        database = settings.DB
        # 设置表名
        table_name = "order_info"


class OrderGoods(BaseModel):
    """
    订单商品详情
    """
    order = IntegerField(verbose_name="订单id")
    goods = IntegerField(verbose_name="商品Id")
    goods_name = CharField(max_length=20, default="", verbose_name="商品名称")
    goods_image = CharField(max_length=200, default="", verbose_name="商品图片")
    goods_price = DecimalField(default=0, verbose_name="商品价格")
    nums = IntegerField(default=0, verbose_name="商品数量")

    class Meta:
        # 绑定上面的数据库实例
        database = settings.DB
        # 设置表名
        table_name = "order_goods"


if __name__ == '__main__':
    settings.DB.create_tables([ShoppingCart, OrderInfo, OrderGoods])
    # for i in range(1,6):
    #     goods_inv = Inventory(goods=i, stocks=100)
    #     goods_inv.save()

    # goods_info = ((1, 2), (2, 3), (3, 10))
    # for goods_id, num in goods_info:
    #     goods_inv = Inventory.get(Inventory.goods == goods_id)
    #     if goods_inv.stocks < num:
    #         print(f'{goods_id}:库存不足')
    #     else:
    #         goods_inv.stocks = goods_inv.stocks - num
    #         goods_inv.save()
