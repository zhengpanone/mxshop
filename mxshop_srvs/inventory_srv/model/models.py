from datetime import datetime
from playhouse.shortcuts import ReconnectMixin
from playhouse.pool import PooledMySQLDatabase

from peewee import *


class ReconnectMySQLDatabase(ReconnectMixin, PooledMySQLDatabase):
    pass


db = ReconnectMySQLDatabase("mxshop_inventory_srv", host="127.0.0.1", port=3306, user="root", password="root")


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
    def update(cls, __data=None, **update):
        return super().update(update_time=datetime.now())

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
    # from faker import Faker
    # from passlib.hash import pbkdf2_sha256
    #
    # fake = Faker(locale='zh_CN')
    # for i in range(10):
    #     category_name = fake.word().capitalize()
    #     c1 = Category(name=category_name)
    #     c1.save()

    # Category.update(name="1111").where(Category.name=='2222').execute()
    # print(list(Category.select()))
    # Category.delete().where(Category.name=='2222').execute()
    # print(list(Category.select()))
