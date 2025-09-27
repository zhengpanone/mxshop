from datetime import datetime, timezone

from peewee import Model, DateTimeField, BooleanField, AutoField, Database
from typing import  Type,Any

class SoftDeleteManager:
    """
    软删除查询管理器
    将常用的查询逻辑封装起来，使查询更安全、更易读
    """
    def __init__(self, model_class: Type[Model]):
        self.model_class = model_class

    def all(self):
        """
        获取所有未删除的对象
        """
        return self.model_class.select().where(self.model_class.is_deleted == False)

    def with_deleted(self):
        """
        获取所有对象，包括已删除的
        """
        return self.model_class.select()

    def deleted(self):
        """
        仅获取已删除的对象
        """
        return self.model_class.select().where(self.model_class.is_deleted == True)

    def __getattr__(self, name: str) -> Any:
        """
        将其他方法（如 get、filter）代理给模型类
        """
        return getattr(self.model_class, name)

# 使用 UTC 时间，避免时区问题，推荐在数据库中也使用 UTC
def get_now():
    return datetime.now(timezone.utc)


class TimestampMixin(Model):
    """
        时间戳和软删除混入类 (Mixin)
        提供创建时间、更新时间和软删除字段
    """
    create_time = DateTimeField(default=get_now, verbose_name="添加时间")
    update_time = DateTimeField(default=get_now, verbose_name="更新时间")
    is_deleted = BooleanField(default=False, verbose_name="是否删除")

    class Meta:
        # 这个 Meta 类是 Peewee 用于继承的，并非 Django 的抽象模型
        # 它确保这个 Mixin 类不会被创建成独立的表
        abstract = True


    def save(self, *args: Any, **kwargs: Any) -> int:
        """
        重写 save 方法，在更新时自动更新 update_time
        """
        # _pk 检查是否有主键，如果有则认为是更新操作
        if self._pk is not None:
            self.update_time = get_now()
        return super().save(*args, **kwargs)

    @classmethod
    def delete(cls, permanently: bool = False):
        """
        重写类方法 delete，支持软删除
        :param permanently: 是否永久删除，默认为 False
        """
        if permanently:
            return super().delete()

        return super().update(is_deleted=True, update_time=get_now())

    def delete_instance(self, permanently: bool = False, recursive: bool = False, delete_nullable: bool = False) -> int:
        """
        重写实例方法 delete_instance，支持软删除
        :param permanently: 是否永久删除，默认为 False
        :param recursive: 是否递归删除相关联的对象 (peewee 的原生参数)
        :param delete_nullable: 是否删除可空的外键对象 (peewee 的原生参数)
        """
        if permanently:
            # 调用 peewee 的原生方法进行永久删除
            return super().delete_instance(recursive=recursive, delete_nullable=delete_nullable)
        else:
            self.is_deleted = True
            self.update_time = get_now()
            # 软删除只需保存实例即可
            return self.save()

    # 软删除管理器
    objects = SoftDeleteManager

def create_base_model(_database:Database)-> Type[Model]:
    """
    动态创建带有时间戳和软删除功能的 BaseModel
    :param _database: 数据库实例
    :return: 继承了 Peewee Model 的基类
    """

    class BaseModel(Model):
        class Meta:
            database = _database
            legacy_table_names = False
    return BaseModel


