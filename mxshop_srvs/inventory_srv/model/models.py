from datetime import datetime
from playhouse.shortcuts import ReconnectMixin
from playhouse.pool import PooledMySQLDatabase

from peewee import *

from inventory_srv.settings import settings


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


# class Stock(BaseModel):
#     """仓库表"""
#     name = CharField(verbose_name="仓库名")
#     address = CharField(verbose_name="仓库地址")

class Inventory(BaseModel):
    """
    商品库存表
    """
    # stock = PrimaryKeyField(verbose_name="库存Id")
    goods = IntegerField(verbose_name="商品Id", unique=True)
    stocks = IntegerField(verbose_name="库存数量", default=0)
    version = IntegerField(verbose_name="版本号", default=0)  # 分布式锁的乐观锁


class InventoryHistory(BaseModel):
    user = IntegerField(verbose_name="用户Id", unique=True)
    goods = IntegerField(verbose_name="商品Id", unique=True)
    num = IntegerField(verbose_name="数量", unique=True)
    order = CharField(verbose_name="订单id", unique=True)
    status = IntegerField(choices=((1, "未出库"), (2, "已出库")), default=1, verbose_name="出库状态")


if __name__ == '__main__':
    db.create_tables([Inventory])
    # for i in range(1,6):
    #     goods_inv = Inventory(goods=i, stocks=100)
    #     goods_inv.save()

    goods_info = ((1, 2), (2, 3), (3, 10))
    for goods_id, num in goods_info:
        goods_inv = Inventory.get(Inventory.goods == goods_id)
        if goods_inv.stocks < num:
            print(f'{goods_id}:库存不足')
        else:
            goods_inv.stocks = goods_inv.stocks-num
            goods_inv.save()
