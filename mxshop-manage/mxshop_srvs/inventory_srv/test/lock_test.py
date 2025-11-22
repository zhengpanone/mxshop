import threading
from datetime import datetime
from playhouse.shortcuts import ReconnectMixin
from playhouse.pool import PooledMySQLDatabase

from peewee import *


class ReconnectMysqlDatabase(ReconnectMixin, PooledMySQLDatabase):
    # python的mro
    pass


db = ReconnectMysqlDatabase('mxshop_inventory_srv', host='127.0.0.1', port=3306, user='root', password='root')


class BaseModel(Model):
    add_time = DateTimeField(default=datetime.now, verbose_name="添加时间")
    update_time = DateTimeField(default=datetime.now, verbose_name="更新时间")
    is_deleted = BooleanField(default=False, verbose_name="是否删除")

    class Meta:
        database = db

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


R = threading.Lock()


def sell():
    goods_list = [(1, 10), (2, 20), (3, 30), (4, 40)]
    # 多线程并发数据不一致问题
    with db.atomic() as txn:
        for goods_id, num in goods_list:
            R.acquire()  # 获取锁
            goods_inv = Inventory.get(Inventory.goods == goods_id)
            import time
            from random import randint
            time.sleep(randint(1, 3))
            if goods_inv.stocks < num:
                print(f'商品：{goods_id}库存不足')
                txn.rollback()
                break
            else:
                # goods_inv.stocks -= num
                query = Inventory.update(stocks=Inventory.stocks - num).where(Inventory.goods == goods_id)
                ok = query.execute()
                if ok:
                    print("更新成功")
                else:
                    print("更新失败")
            R.release()  # 释放锁


if __name__ == '__main__':
    db.create_tables([Inventory])

    t1 = threading.Thread(target=sell)
    t2 = threading.Thread(target=sell)

    t1.start()
    t2.start()
    t1.join()
    t2.join()
