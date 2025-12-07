<template>
  <div class="brand-form-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-left">
        <el-button type="text" class="back-button" @click="goBack">
          <el-icon><ArrowLeft /></el-icon>
          返回列表
        </el-button>
        <el-divider direction="vertical" />
        <h1 class="page-title">
          <el-icon class="title-icon"><Shop /></el-icon>
          {{ isEdit ? "编辑品牌" : "新增品牌" }}
        </h1>
      </div>
      <div class="header-right">
        <el-tag :type="isEdit ? 'warning' : 'success'">
          {{ isEdit ? "编辑模式" : "新增模式" }}
        </el-tag>
      </div>
    </div>

    <!-- 表单卡片 -->
    <el-card class="form-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <div class="header-info">
            <el-icon class="header-icon"><EditPen /></el-icon>
            <span class="header-title">品牌信息</span>
          </div>
          <div class="form-progress">
            <span class="progress-text"
              >完成度: {{ completionPercentage }}%</span
            >
            <el-progress
              :percentage="completionPercentage"
              :stroke-width="6"
              :show-text="false"
              class="progress-bar"
            />
          </div>
        </div>
      </template>

      <el-form
        ref="brandFormRef"
        :model="brand"
        :rules="rules"
        label-width="140px"
        class="brand-form"
        label-position="right"
      >
        <el-row :gutter="20">
          <!-- 基本信息 -->
          <el-col :span="24">
            <div class="form-section">
              <div class="section-header">
                <el-icon><Document /></el-icon>
                <span>基本信息</span>
              </div>

              <el-row :gutter="20">
                <el-col :span="12">
                  <el-form-item label="品牌名称" prop="name" required>
                    <el-input
                      v-model="brand.name"
                      placeholder="请输入品牌名称"
                      clearable
                      maxlength="50"
                      show-word-limit
                    />
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-form-item label="品牌首字母" prop="letter">
                    <el-input
                      v-model="brand.letter"
                      placeholder="请输入品牌首字母"
                      maxlength="1"
                      style="text-transform: uppercase"
                      @input="handleLetterInput"
                    />
                  </el-form-item>
                </el-col>
              </el-row>

              <el-row :gutter="20">
                <el-col :span="12">
                  <el-form-item label="排序" prop="sort">
                    <el-input-number
                      v-model="brand.sort"
                      :min="0"
                      :max="9999"
                      placeholder="请输入排序值"
                      class="full-width"
                    />
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <div class="status-group">
                    <el-form-item label="是否显示" prop="showStatus">
                      <el-switch
                        v-model="brand.showStatus"
                        :active-value="1"
                        :inactive-value="0"
                        active-text="显示"
                        inactive-text="隐藏"
                      />
                    </el-form-item>
                  </div>
                </el-col>
              </el-row>
            </div>
          </el-col>

          <!-- 图片信息 -->
          <el-col :span="24">
            <div class="form-section">
              <div class="section-header">
                <el-icon><Picture /></el-icon>
                <span>图片信息</span>
              </div>

              <el-row :gutter="20">
                <el-col :span="12">
                  <el-form-item label="品牌LOGO" prop="logo" required>
                    <div class="upload-container">
                      <SingleUpload
                        :value="brand.logo"
                        @update:value="brand.logo = $event"
                        :preview="true"
                        accept="image/*"
                        tip="建议尺寸: 200x200px，格式: jpg/png"
                      />
                    </div>
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-form-item label="品牌专区大图" prop="bigPic">
                    <div class="upload-container">
                      <SingleUpload
                        :value="brand.bigPic"
                        @update:value="brand.bigPic = $event"
                        :preview="true"
                        accept="image/*"
                        tip="建议尺寸: 800x400px，格式: jpg/png"
                      />
                    </div>
                  </el-form-item>
                </el-col>
              </el-row>
            </div>
          </el-col>

          <!-- 详细信息 -->
          <el-col :span="24">
            <div class="form-section">
              <div class="section-header">
                <el-icon><Reading /></el-icon>
                <span>详细信息</span>
              </div>

              <el-form-item label="品牌故事" prop="brandStory">
                <el-input
                  v-model="brand.brandStory"
                  type="textarea"
                  placeholder="请输入品牌故事..."
                  :autosize="{ minRows: 4, maxRows: 8 }"
                  maxlength="500"
                  show-word-limit
                />
              </el-form-item>

              <el-form-item label="品牌制造商" prop="factoryStatus">
                <el-radio-group
                  v-model="brand.factoryStatus"
                  class="radio-group"
                >
                  <el-radio :label="1" class="radio-item">
                    <el-icon><Check /></el-icon>
                    是制造商
                  </el-radio>
                  <el-radio :label="0" class="radio-item">
                    <el-icon><Close /></el-icon>
                    非制造商
                  </el-radio>
                </el-radio-group>
              </el-form-item>
            </div>
          </el-col>
        </el-row>

        <!-- 操作按钮 -->
        <el-form-item class="form-actions">
          <div class="action-buttons">
            <el-button
              type="primary"
              size="large"
              :loading="submitLoading"
              @click="onSubmit"
            >
              <el-icon><Check /></el-icon>
              {{ isEdit ? "保存修改" : "创建品牌" }}
            </el-button>

            <el-button v-if="!isEdit" size="large" @click="resetForm">
              <el-icon><RefreshLeft /></el-icon>
              重置表单
            </el-button>

            <el-button size="large" @click="goBack">
              <el-icon><Close /></el-icon>
              取消
            </el-button>
          </div>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 帮助信息 -->
    <el-card class="help-card" shadow="never">
      <template #header>
        <div class="help-header">
          <el-icon><QuestionFilled /></el-icon>
          <span>填写说明</span>
        </div>
      </template>
      <div class="help-content">
        <el-alert title="温馨提示" type="info" :closable="false" show-icon>
          <ul class="help-list">
            <li>品牌名称为必填项，建议控制在20个字符以内</li>
            <li>品牌LOGO为必填项，建议上传正方形图片，尺寸200x200px</li>
            <li>品牌首字母用于快速检索，系统会自动转换为大写</li>
            <li>排序值越小，品牌显示越靠前</li>
            <li>品牌故事用于品牌详情页展示，建议控制在500字以内</li>
          </ul>
        </el-alert>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from "vue";
import { useRouter, useRoute } from "vue-router";
import {
  ArrowLeft,
  Shop,
  EditPen,
  Document,
  Picture,
  Reading,
  Check,
  Close,
  RefreshLeft,
  QuestionFilled,
} from "@element-plus/icons-vue";
import type { FormInstance, FormRules } from "element-plus";
import { ElMessage, ElMessageBox } from "element-plus";
// import { createBrand, putBrands, getBrand } from '@/apis/goods'
import SingleUpload from "@/components/Upload/singleUpload.vue";

const router = useRouter();
const route = useRoute();

// 表单引用
const brandFormRef = ref<FormInstance>();

// 状态管理
const submitLoading = ref(false);
const isEdit = ref(false);

// 默认品牌数据
const defaultBrand = {
  name: "",
  letter: "",
  logo: "",
  bigPic: "",
  brandStory: "",
  sort: 0,
  showStatus: 1,
  factoryStatus: 0,
};

// 表单数据
const brand = reactive({ ...defaultBrand });

// 表单验证规则
const rules: FormRules = {
  name: [
    { required: true, message: "请输入品牌名称", trigger: "blur" },
    { min: 2, max: 50, message: "品牌名称长度在2-50个字符", trigger: "blur" },
  ],
  logo: [{ required: true, message: "请上传品牌LOGO", trigger: "change" }],
  letter: [
    { pattern: /^[A-Z]$/, message: "请输入单个英文字母", trigger: "blur" },
  ],
  sort: [{ type: "number", message: "排序必须为数字", trigger: "blur" }],
};

// 计算完成度
const completionPercentage = computed(() => {
  const fields = [
    brand.name,
    brand.logo,
    brand.letter,
    brand.brandStory,
    brand.bigPic,
  ];
  const completed = fields.filter(
    (field) => field && String(field).trim()
  ).length;
  return Math.round((completed / fields.length) * 100);
});

// 初始化
onMounted(() => {
  initForm();
});

// 初始化表单
const initForm = async () => {
  if (route.query.id) {
    isEdit.value = true;
    await loadBrandData(route.query.id as string);
  } else {
    isEdit.value = false;
    resetBrandData();
  }
};

// 加载品牌数据
const loadBrandData = async (id: string) => {
  try {
    // 模拟API调用
    // const response = await getBrand(id)
    // Object.assign(brand, response)

    // 模拟数据
    setTimeout(() => {
      Object.assign(brand, {
        name: "示例品牌",
        letter: "S",
        logo: "https://via.placeholder.com/200x200",
        bigPic: "https://via.placeholder.com/800x400",
        brandStory: "这是一个示例品牌故事...",
        sort: 1,
        showStatus: 1,
        factoryStatus: 1,
      });
    }, 500);
  } catch (error) {
    ElMessage.error("加载品牌数据失败");
  }
};

// 重置品牌数据
const resetBrandData = () => {
  Object.assign(brand, { ...defaultBrand });
};

// 处理首字母输入
const handleLetterInput = (value: string) => {
  brand.letter = value.toUpperCase();
};

// 提交表单
const onSubmit = async () => {
  if (!brandFormRef.value) return;

  try {
    await brandFormRef.value.validate();

    const result = await ElMessageBox.confirm(
      `确定要${isEdit.value ? "保存修改" : "创建品牌"}吗？`,
      "确认操作",
      {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }
    );

    if (result === "confirm") {
      submitLoading.value = true;

      try {
        if (isEdit.value) {
          // await putBrands(route.query.id, brand)
          // 模拟API调用
          await new Promise((resolve) => setTimeout(resolve, 1500));
          ElMessage.success("修改成功");
        } else {
          // await createBrand(brand)
          // 模拟API调用
          await new Promise((resolve) => setTimeout(resolve, 1500));
          ElMessage.success("创建成功");
        }

        goBack();
      } catch (error) {
        ElMessage.error(isEdit.value ? "修改失败" : "创建失败");
      } finally {
        submitLoading.value = false;
      }
    }
  } catch (error) {
    if (error !== "cancel") {
      ElMessage.error("表单验证失败，请检查输入");
    }
  }
};

// 重置表单
const resetForm = () => {
  brandFormRef.value?.resetFields();
  resetBrandData();
  ElMessage.success("表单已重置");
};

// 返回上一页
const goBack = () => {
  router.back();
};
</script>

<style scoped lang="scss">
.brand-form-container {
  padding: 20px;
  background: #f5f7fa;
  min-height: 100vh;
}

// 页面头部
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  background: white;
  padding: 16px 24px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.header-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.back-button {
  color: #666;
  font-size: 14px;
  padding: 0;

  &:hover {
    color: #409eff;
  }

  .el-icon {
    margin-right: 4px;
  }
}

.page-title {
  font-size: 20px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.title-icon {
  color: #409eff;
}

.header-right {
  :deep(.el-tag) {
    font-weight: 500;
  }
}

// 表单卡片
.form-card {
  margin-bottom: 20px;
  border-radius: 12px;
  border: none;

  :deep(.el-card__header) {
    background: linear-gradient(135deg, #f8f9fa, #ffffff);
    border-bottom: 1px solid #f0f0f0;
  }
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.header-icon {
  color: #409eff;
  font-size: 18px;
}

.header-title {
  font-size: 16px;
  font-weight: 600;
  color: #2c3e50;
}

.form-progress {
  display: flex;
  align-items: center;
  gap: 12px;
}

.progress-text {
  font-size: 14px;
  color: #7f8c8d;
  white-space: nowrap;
}

.progress-bar {
  width: 100px;
}

// 表单样式
.brand-form {
  padding: 20px 0;

  :deep(.el-form-item__label) {
    color: #2c3e50;
    font-weight: 500;
  }

  :deep(.el-input__wrapper) {
    border-radius: 8px;
    transition: all 0.3s ease;

    &:hover {
      box-shadow: 0 2px 8px rgba(64, 158, 255, 0.15);
    }

    &.is-focus {
      box-shadow: 0 2px 12px rgba(64, 158, 255, 0.2);
    }
  }

  :deep(.el-textarea__inner) {
    border-radius: 8px;
    transition: all 0.3s ease;

    &:hover {
      box-shadow: 0 2px 8px rgba(64, 158, 255, 0.15);
    }

    &:focus {
      box-shadow: 0 2px 12px rgba(64, 158, 255, 0.2);
    }
  }
}

// 表单分组
.form-section {
  margin-bottom: 32px;

  &:last-child {
    margin-bottom: 0;
  }
}

.section-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 20px;
  padding-bottom: 10px;
  border-bottom: 2px solid #f0f0f0;

  .el-icon {
    color: #409eff;
    font-size: 16px;
  }

  span {
    font-size: 16px;
    font-weight: 600;
    color: #2c3e50;
  }
}

// 上传容器
.upload-container {
  width: 100%;
}

// 状态组
.status-group {
  :deep(.el-form-item__content) {
    justify-content: flex-start;
  }
}

// 单选框组
.radio-group {
  display: flex;
  gap: 24px;
}

.radio-item {
  display: flex;
  align-items: center;
  gap: 4px;

  .el-icon {
    font-size: 14px;
  }

  &.is-checked {
    .el-icon {
      color: #409eff;
    }
  }
}

// 全宽度组件
.full-width {
  width: 100%;
}

// 操作按钮
.form-actions {
  margin-top: 40px;
  padding-top: 20px;
  border-top: 1px solid #f0f0f0;

  :deep(.el-form-item__content) {
    justify-content: center;
  }
}

.action-buttons {
  display: flex;
  gap: 16px;

  .el-button {
    min-width: 120px;
    border-radius: 8px;
    font-weight: 500;

    &.el-button--primary {
      background: linear-gradient(135deg, #409eff, #5cb3ff);
      border: none;

      &:hover {
        transform: translateY(-1px);
        box-shadow: 0 4px 12px rgba(64, 158, 255, 0.4);
      }
    }

    &:not(.el-button--primary) {
      &:hover {
        transform: translateY(-1px);
      }
    }
  }
}

// 帮助卡片
.help-card {
  border-radius: 12px;
  border: none;

  :deep(.el-card__header) {
    background: linear-gradient(135deg, #f0f8ff, #ffffff);
    border-bottom: 1px solid #e8f4fd;
  }
}

.help-header {
  display: flex;
  align-items: center;
  gap: 8px;

  .el-icon {
    color: #409eff;
    font-size: 16px;
  }

  span {
    font-size: 14px;
    font-weight: 600;
    color: #2c3e50;
  }
}

.help-content {
  :deep(.el-alert) {
    border-radius: 8px;
    border: none;
    background: linear-gradient(135deg, #e8f4fd, #f0f8ff);
  }
}

.help-list {
  margin: 0;
  padding-left: 16px;

  li {
    margin: 8px 0;
    color: #7f8c8d;
    font-size: 13px;
    line-height: 1.5;
  }
}

// 响应式设计
@media (max-width: 768px) {
  .brand-form-container {
    padding: 10px;
  }

  .page-header {
    flex-direction: column;
    gap: 12px;
    align-items: flex-start;
  }

  .header-left {
    width: 100%;
  }

  .form-progress {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }

  .progress-bar {
    width: 100%;
  }

  .brand-form {
    :deep(.el-form-item) {
      margin-bottom: 20px;
    }
  }

  .action-buttons {
    flex-direction: column;
    width: 100%;

    .el-button {
      width: 100%;
    }
  }

  .radio-group {
    flex-direction: column;
    gap: 12px;
  }
}

@media (max-width: 480px) {
  .brand-form {
    :deep(.el-form--label-top) .el-form-item__label {
      padding-bottom: 8px;
    }
  }

  .el-col {
    width: 100% !important;
  }
}
</style>
