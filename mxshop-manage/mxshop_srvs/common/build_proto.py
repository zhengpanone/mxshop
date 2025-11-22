#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import re
import subprocess
import sys
from pathlib import Path
from typing import List, Optional

# 尝试导入配置文件，如果不存在则使用默认配置
try:
    from proto_config import PROTO_CONFIG
except ImportError:
    PROTO_CONFIG = {
        "proto_dir": "proto",
        "output_dir": "proto/pb",
        "package_prefix": "common.proto.pb",
        "extra_args": [],
        "specific_files": [],
        "recursive": True,  # 是否递归搜索子目录
        "exclude_patterns": ["*_test.proto", "test_*.proto"]  # 排除的文件模式
    }


class AutoProtoBuilder:
    def __init__(self, config=None):
        self.config = config or PROTO_CONFIG

    def build(self):
        """执行编译和修复"""
        print("🚀 自动搜索并编译 Proto 文件...")
        print(f"📁 源目录: {self.config['proto_dir']}")
        print(f"📁 输出目录: {self.config['output_dir']}")
        print(f"📦 包前缀: {self.config['package_prefix']}")
        print(f"🔍 递归搜索: {'是' if self.config.get('recursive', True) else '否'}")

        # 创建输出目录
        os.makedirs(self.config['output_dir'], exist_ok=True)

        # 自动获取要编译的文件
        proto_files = self._discover_proto_files()

        if not proto_files:
            print("❌ 未找到 proto 文件")
            return False

        # 显示找到的文件
        print(f"\n📄 找到 {len(proto_files)} 个 proto 文件:")
        for i, proto_file in enumerate(proto_files, 1):
            print(f"  {i}. {proto_file}")
        print()

        # 运行 protoc
        if not self._run_protoc(proto_files):
            return False

        # 修复导入
        self._fix_imports()

        print("✅ 编译完成!")
        return True

    def _discover_proto_files(self) -> List[Path]:
        """自动发现 proto 文件"""
        proto_dir = Path(self.config['proto_dir'])

        if not proto_dir.exists():
            print(f"❌ 源目录不存在: {proto_dir}")
            return []

        # 如果指定了具体文件，优先使用
        if self.config.get('specific_files'):
            print("📋 使用指定的文件列表")
            return [Path(f) for f in self.config['specific_files']]

        # 自动搜索 proto 文件
        proto_files = []
        recursive = self.config.get('recursive', True)
        exclude_patterns = self.config.get('exclude_patterns', [])

        if recursive:
            print("🔍 递归搜索 .proto 文件...")
            pattern = "**/*.proto"
        else:
            print("🔍 搜索当前目录下的 .proto 文件...")
            pattern = "*.proto"

        # 搜索文件
        for proto_file in proto_dir.glob(pattern):
            # 检查是否被排除
            if self._should_exclude(proto_file, exclude_patterns):
                print(f"⏭️  跳过: {proto_file} (匹配排除规则)")
                continue
            proto_files.append(proto_file)

        # 按文件名排序
        proto_files.sort(key=lambda x: x.name)

        return proto_files

    def _should_exclude(self, file_path: Path, exclude_patterns: List[str]) -> bool:
        """检查文件是否应该被排除"""
        file_name = file_path.name

        for pattern in exclude_patterns:
            if file_path.match(pattern):
                return True

        return False

    def _run_protoc(self, proto_files: List[Path]) -> bool:
        """运行 protoc"""
        cmd = [
            sys.executable, "-m", "grpc_tools.protoc",
            f"--python_out={self.config['output_dir']}",
            f"--grpc_python_out={self.config['output_dir']}",
            f"--mypy_out={self.config['output_dir']}",
            f"-I={self.config['proto_dir']}"
        ]

        # 添加额外参数
        extra_args = self.config.get('extra_args', [])
        if extra_args:
            print(f"⚙️  额外参数: {' '.join(extra_args)}")
        cmd.extend(extra_args)

        # 添加文件
        cmd.extend([str(f) for f in proto_files])

        print("⚙️  执行编译命令...")
        print(f"   protoc {' '.join(cmd[3:])}")

        try:
            result = subprocess.run(cmd, check=True, capture_output=True, text=True)
            if result.stdout:
                print("📝 编译输出:")
                print(result.stdout)
            return True
        except subprocess.CalledProcessError as e:
            print(f"❌ 编译失败:")
            if e.stderr:
                print(e.stderr)
            if e.stdout:
                print(e.stdout)
            return False
        except FileNotFoundError:
            print("❌ 找不到 grpcio-tools，请安装: pip install grpcio-tools")
            return False

    def _fix_imports(self):
        """修复导入"""
        print("🔧 修复生成文件的导入路径...")

        output_dir = Path(self.config['output_dir'])
        py_files = list(output_dir.glob("*_pb2.py")) + list(output_dir.glob("*_pb2_grpc.py"))

        if not py_files:
            print("⚠️  未找到生成的 Python 文件")
            return

        fixed_count = 0
        for py_file in py_files:
            if self._fix_file_imports(py_file):
                fixed_count += 1
                print(f"  ✓ {py_file.name}")

        print(f"🔄 修复了 {fixed_count}/{len(py_files)} 个文件的导入")

    def _fix_file_imports(self, file_path: Path) -> bool:
        """修复文件导入"""
        try:
            with open(file_path, 'r', encoding='utf-8') as f:
                content = f.read()

            original_content = content
            prefix = self.config['package_prefix']

            # 修复 pb2 导入
            content = re.sub(
                r'^import (\w+_pb2) as (\w+__pb2)$',
                rf'from {prefix} import \1 as \2',
                content, flags=re.MULTILINE
            )

            # 修复 grpc 导入
            content = re.sub(
                r'^import (\w+_pb2_grpc) as (\w+__pb2_grpc)$',
                rf'from {prefix} import \1 as \2',
                content, flags=re.MULTILINE
            )

            # 如果有修改则写回文件
            if content != original_content:
                with open(file_path, 'w', encoding='utf-8') as f:
                    f.write(content)
                return True

            return False

        except Exception as e:
            print(f"❌ 修复 {file_path.name} 失败: {e}")
            return False

    def clean(self):
        """清理生成的文件"""
        print("🧹 清理生成的文件...")

        output_dir = Path(self.config['output_dir'])
        if not output_dir.exists():
            print("📁 输出目录不存在")
            return

        # 要清理的文件模式
        patterns = ["*_pb2.py", "*_pb2.pyi", "*_pb2_grpc.py"]
        removed_files = []

        for pattern in patterns:
            for file_path in output_dir.glob(pattern):
                removed_files.append(file_path.name)
                file_path.unlink()

        if removed_files:
            print(f"🗑️  删除了 {len(removed_files)} 个文件:")
            for filename in sorted(removed_files):
                print(f"  - {filename}")
        else:
            print("📄 没有找到需要清理的文件")

    def list_files(self):
        """列出会被编译的文件"""
        print("📋 扫描 proto 文件...")
        proto_files = self._discover_proto_files()

        if not proto_files:
            print("❌ 未找到 proto 文件")
            return

        print(f"\n📄 找到 {len(proto_files)} 个文件:")
        for i, proto_file in enumerate(proto_files, 1):
            # 显示相对路径和文件大小
            try:
                size = proto_file.stat().st_size
                print(f"  {i:2d}. {proto_file} ({size} bytes)")
            except:
                print(f"  {i:2d}. {proto_file}")


def main():
    """主函数"""
    builder = AutoProtoBuilder()

    # 解析命令行参数
    if len(sys.argv) > 1:
        command = sys.argv[1].lower()

        if command == "clean":
            builder.clean()
            return
        elif command == "list":
            builder.list_files()
            return
        elif command == "help":
            print_help()
            return
        else:
            print(f"❌ 未知命令: {command}")
            print_help()
            return

    # 默认执行编译
    success = builder.build()
    sys.exit(0 if success else 1)


def print_help():
    """显示帮助信息"""
    print("""
🛠️  Proto 自动编译工具

使用方法:
  python build_proto.py           # 自动编译所有 proto 文件
  python build_proto.py clean     # 清理生成的文件
  python build_proto.py list      # 列出会被编译的文件
  python build_proto.py help      # 显示此帮助

配置说明:
  在 proto_config.py 中可以配置:
  - proto_dir: proto 文件目录
  - output_dir: 输出目录
  - package_prefix: 包前缀
  - recursive: 是否递归搜索子目录
  - exclude_patterns: 排除的文件模式
  - specific_files: 指定特定文件(可选)

示例配置 (proto_config.py):
  PROTO_CONFIG = {
      "proto_dir": "protos",
      "output_dir": "generated", 
      "package_prefix": "my.package",
      "recursive": True,
      "exclude_patterns": ["*_test.proto"]
  }
""")


# ============================================
# 示例配置文件 proto_config.py
# ============================================

EXAMPLE_CONFIG = '''
# proto_config.py
"""
Proto 编译配置文件
"""

PROTO_CONFIG = {
    # Proto 文件源目录
    "proto_dir": "common/proto",

    # 生成的 Python 文件输出目录
    "output_dir": "common/proto/pb", 

    # 导入时使用的包前缀
    "package_prefix": "common.proto.pb",

    # 是否递归搜索子目录中的 proto 文件
    "recursive": True,

    # 排除的文件模式(支持通配符)
    "exclude_patterns": [
        "*_test.proto",      # 排除测试文件
        "test_*.proto",      # 排除测试文件
        "deprecated_*.proto", # 排除废弃文件
    ],

    # 指定特定的文件进行编译(可选，如果设置则忽略自动搜索)
    "specific_files": [
        # "common/proto/user.proto",
        # "common/proto/order.proto",
    ],

    # 传递给 protoc 的额外参数
    "extra_args": [
        # "--experimental_allow_proto3_optional",
    ]
}
'''

# 如果运行时加 --create-config 参数，创建示例配置文件
if len(sys.argv) > 1 and sys.argv[1] == "--create-config":
    with open("proto_config.py", "w", encoding="utf-8") as f:
        f.write(EXAMPLE_CONFIG)
    print("✅ 已创建示例配置文件: proto_config.py")
    sys.exit(0)


if __name__ == "__main__":
    main()