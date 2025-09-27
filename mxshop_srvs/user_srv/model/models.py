from datetime import datetime
from peewee import *

from common.model.base_model import create_base_model, TimestampMixin
from user_srv.settings import settings


BaseModel = create_base_model(settings.DB)

# 用户模型
class User(BaseModel,TimestampMixin):
    GENDER_CHOICES = (
        ('1', '男'),
        ('2', '女')
    )

    ROLE_CHOICES = (
        ('1', '管理员'),
        ('2', '普通用户')
    )
    mobile = CharField(max_length=11, index=True, unique=True, verbose_name="手机号码")
    password = CharField(max_length=255, verbose_name="密码")
    nickname = CharField(max_length=20, null=True, verbose_name="昵称")
    head_url = CharField(max_length=255, null=True, verbose_name="头像")
    birthday = DateField(null=True, verbose_name="出生日期")
    address = CharField(max_length=255, null=True, verbose_name="地址")
    desc = TextField(null=True, verbose_name="个人简介")
    gender = SmallIntegerField(choices=GENDER_CHOICES, null=True, verbose_name="性别")
    role = TextField(choices=ROLE_CHOICES, verbose_name="角色")

    class Meta:
        table_name = 'user'


class Role(BaseModel,TimestampMixin):
    name = CharField(max_length=255, verbose_name="角色名称")
    remark = TextField(null=True, verbose_name="角色备注")
    status =  BooleanField(default=True, verbose_name="是否启用")

    class Meta:
        table_name = 'role'

class User(BaseModel,TimestampMixin):
    GENDER_CHOICES = (
        ('1', '男'),
        ('2', '女')
    )

    ROLE_CHOICES = (
        ('1', '管理员'),
        ('2', '普通用户')
    )
    mobile = CharField(max_length=11, index=True, unique=True, verbose_name="手机号码")
    password = CharField(max_length=255, verbose_name="密码")
    nickname = CharField(max_length=20, null=True, verbose_name="昵称")
    head_url = CharField(max_length=255, null=True, verbose_name="头像")
    birthday = DateField(null=True, verbose_name="出生日期")
    address = CharField(max_length=255, null=True, verbose_name="地址")
    desc = TextField(null=True, verbose_name="个人简介")
    gender = SmallIntegerField(choices=GENDER_CHOICES, null=True, verbose_name="性别")
    role = TextField(choices=ROLE_CHOICES, verbose_name="角色")

    class Meta:
        table_name = 'user'


class Role(BaseModel,TimestampMixin):
    name = CharField(max_length=255, verbose_name="角色名称")
    remark = TextField(null=True, verbose_name="角色备注")
    status =  BooleanField(default=True, verbose_name="是否启用")

    class Meta:
        table_name = 'role'


class UserRole(BaseModel):
    user_id = IntegerField(index=True, verbose_name="用户ID")
    role_id = IntegerField(index=True, verbose_name="角色ID")

    class Meta:
        table_name = "user_role"
        indexes = (
            (("user_id", "role_id"), True),  # 组合唯一索引
        )

class DictType(BaseModel,TimestampMixin):
    system_flag = BooleanField(default=True, verbose_name="是否系统内置")
    dict_code = CharField(max_length=255, null=False,  unique=True, verbose_name="字典标识")
    dict_name = CharField(max_length=255, null=False, verbose_name="字典名称")
    remark = TextField(null=True, verbose_name="备注信息")
    status = BooleanField(default=True, verbose_name="是否启用")

    class Meta:
        table_name = "dict_type"

class DictItem(BaseModel,TimestampMixin):
    dict_type_id =IntegerField(index=True, verbose_name="字典类ID")
    dict_type = CharField(max_length=255, verbose_name="字典类型")
    label = CharField(max_length=255,  verbose_name="标签名")
    item_value = CharField(max_length=255,  verbose_name="数据值")
    description= CharField(max_length=255, verbose_name="描述")
    sort_order = IntegerField(verbose_name="排序")
    remark = TextField(null=True, verbose_name="备注信息")
    status = BooleanField(default=True, verbose_name="是否启用")


    class Meta:
        table_name = "dict_item"
        indexes = (
            (("dict_type", "item_value"), True),  # 组合唯一索引
        )



if __name__ == '__main__':
    # settings.DB.create_tables([User], safe=False)

    # import hashlib
    # m = hashlib.md5()
    # m.update('测试数据MD5'.encode('utf-8'))
    # result = m.hexdigest()
    # print(result)
    # from passlib.hash import pbkdf2_sha256
    # hash_result = pbkdf2_sha256.hash('测试数据MD5'.encode('utf-8'))
    # print(hash_result)
    # verify_result = pbkdf2_sha256.verify('测试数据MD5', hash_result)
    # print(verify_result)

    # from faker import Faker
    # from passlib.hash import pbkdf2_sha256
    #
    # fake = Faker(locale='zh_CN')
    # for i in range(10):
    #     user.py = User()
    #     user.py.mobile = fake.phone_number()
    #     user.py.nickname = fake.name()
    #     user.py.password = pbkdf2_sha256.hash("admin123".encode('utf-8'))
    #     user.py.save()
    # import time
    # from datetime import date
    #
    # users = User.select().paginate(1, 2)
    # for user in User.select().paginate(1,2):
    #     print(user.nickname)
    #     if user.birthday:
    #         print(user.birthday)
    #         u_time = int(time.mktime(user.birthday.timetuple()))
    #         print(date.fromtimestamp(u_time))
    settings.DB.create_tables([DictType,DictItem], safe=True)
    pass