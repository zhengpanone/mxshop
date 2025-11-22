import uuid
from datetime import datetime

from peewee import *
# from playhouse.pool import PooledMySQLDatabase
# from playhouse.shortcuts import ReconnectMixin

from userop_srv.settings import settings


# class ReconnectMySQLDatabase(ReconnectMixin, PooledMySQLDatabase):
#     pass


# db = ReconnectMySQLDatabase("mxshop_userop_srv", host="127.0.0.1", port=3306, user="root", password="root")


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


class LeavingMessages(BaseModel):
    """用户留言"""
    MESSAGE_CHOICES = (
        (1, '留言'),
        (2, '投诉'),
        (3, '询问'),
        (4, '售后'),
        (5, '求购')
    )

    user = IntegerField(verbose_name="用户ID")
    message_type = IntegerField(default=1, choices=MESSAGE_CHOICES, verbose_name="留言类型",
                                help_text=u'留言类型:1 留言、2 投诉、3 询问、4 售后、5 求购')
    subject = CharField(max_length=100, default='', verbose_name="主题")
    message = TextField(default='', verbose_name="留言内容", help_text='留言内容')
    file = CharField(max_length=100, verbose_name="上传的文件", help_text='上传的文件')

    class Meta:
        # 绑定上面的数据库实例
        database = settings.DB
        # 设置表名
        table_name = "leaving_messages"


class Address(BaseModel):
    """收件地址"""
    user = IntegerField(verbose_name="用户ID")
    province = CharField(max_length=100, default='', verbose_name="省份")
    city = CharField(max_length=100, default='', verbose_name="城市")
    district = CharField(max_length=100, default='', verbose_name="区域")
    address = CharField(max_length=100, default='', verbose_name="详细地址")
    singer_name = CharField(max_length=100, default='', verbose_name="签收人")
    singer_mobile = CharField(max_length=11, default='', verbose_name="电话")

    class Meta:
        # 绑定上面的数据库实例
        database = settings.DB
        # 设置表名
        table_name = "address"


class UserFav(BaseModel):
    """用户收藏"""
    user = IntegerField(verbose_name="用户ID")
    goods = IntegerField(verbose_name="商品ID")

    class Meta:
        # 绑定上面的数据库实例
        database = settings.DB
        # 设置表名
        table_name = "user_fav"
        primary_key = CompositeKey('user', 'goods')


if __name__ == '__main__':
    settings.DB.create_tables([LeavingMessages, Address, UserFav])
