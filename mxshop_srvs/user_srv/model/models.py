import uuid

from peewee import *
from user_srv.settings import settings


class BaseModel(Model):
    class Meta:
        database = settings.DB


# 用户模型
class User(BaseModel):
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
    import time
    from datetime import date

    users = User.select().paginate(1, 2)
    for user in User.select().paginate(1,2):
        print(user.nickname)
        if user.birthday:
            print(user.birthday)
            u_time = int(time.mktime(user.birthday.timetuple()))
            print(date.fromtimestamp(u_time))