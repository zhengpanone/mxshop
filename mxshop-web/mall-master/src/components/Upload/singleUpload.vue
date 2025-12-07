<template>
  <div class="single-upload-container">
    <el-upload
      :action="uploadUrl"
      :data="useOss ? dataObj : null"
      :headers="uploadHeaders"
      :show-file-list="false"
      :before-upload="beforeUpload"
      :on-success="handleUploadSuccess"
      :on-error="handleUploadError"
      :on-progress="handleProgress"
      :accept="accept"
      :disabled="uploading"
      class="single-upload"
    >
      <div class="upload-content">
        <!-- 已有图片时显示预览 -->
        <div v-if="imageUrl && !uploading" class="image-preview">
          <el-image
            :src="imageUrl"
            fit="cover"
            class="uploaded-image"
            :preview-src-list="[imageUrl]"
            :preview-teleported="true"
          >
            <template #error>
              <div class="image-error">
                <el-icon><Picture /></el-icon>
                <span>加载失败</span>
              </div>
            </template>
          </el-image>

          <!-- 操作按钮遮罩 -->
          <div class="image-actions">
            <el-button
              type="primary"
              :icon="View"
              circle
              size="small"
              @click.stop="handlePreview"
            />
            <el-button
              type="danger"
              :icon="Delete"
              circle
              size="small"
              @click.stop="handleRemove"
            />
          </div>
        </div>

        <!-- 上传中状态 -->
        <div v-else-if="uploading" class="uploading-content">
          <el-progress
            :percentage="uploadProgress"
            type="circle"
            :width="60"
            class="upload-progress"
          />
          <p class="uploading-text">上传中...</p>
        </div>

        <!-- 默认上传区域 -->
        <div v-else class="upload-area">
          <el-icon class="upload-icon"><Plus /></el-icon>
          <div class="upload-text">点击上传图片</div>
          <div v-if="tip" class="upload-tip">{{ tip }}</div>
        </div>
      </div>
    </el-upload>

    <!-- 预览对话框 -->
    <el-dialog
      v-model="dialogVisible"
      title="图片预览"
      width="50%"
      center
      class="preview-dialog"
    >
      <div class="preview-content">
        <el-image :src="imageUrl" fit="contain" class="preview-image">
          <template #error>
            <div class="preview-error">
              <el-icon><Picture /></el-icon>
              <span>图片加载失败</span>
            </div>
          </template>
        </el-image>
      </div>
      <template #footer>
        <el-button @click="dialogVisible = false">关闭</el-button>
        <el-button type="primary" @click="downloadImage">下载图片</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from "vue";
import { Plus, View, Delete, Picture } from "@element-plus/icons-vue";
import { ElMessage, ElMessageBox } from "element-plus";
import type { UploadFile, UploadProps } from "element-plus";

// Props定义
interface Props {
  value: string; //  图片URL（必需）
  accept?: string; // 接受的文件类型
  tip?: string; // 提示文本
  maxSize?: number; // MB
  useOss?: boolean; // 是否使用OSS
  disabled?: boolean; // 是否禁用
}

const props = withDefaults(defineProps<Props>(), {
  accept: "image/*",
  tip: "只能上传jpg/png文件，且不超过10MB",
  maxSize: 10,
  useOss: true,
  disabled: false,
});

// Emits定义
const emit = defineEmits<{
  "update:value": [value: string];
}>();

// 响应式数据
const dialogVisible = ref(false);
const uploading = ref(false);
const uploadProgress = ref(0);

// OSS相关配置
const dataObj = ref({
  policy: "",
  signature: "",
  OSSAccessKeyId: "",
  key: "",
  dir: "",
  host: "",
  success_action_status: "200",
  callback: "",
});

// 上传配置
const ossUploadUrl = "http://macro-oss.oss-cn-shenzhen.aliyuncs.com";
const minioUploadUrl = "http://localhost:8080/minio/upload";

const uploadUrl = computed(() =>
  props.useOss ? dataObj.value.host : minioUploadUrl
);

const uploadHeaders = computed(() => ({}));

// 图片URL
const imageUrl = computed(() => props.value);

// 图片名称
const imageName = computed(() => {
  if (props.value) {
    return props.value.substring(props.value.lastIndexOf("/") + 1);
  }
  return "";
});

// 监听value变化
watch(
  () => props.value,
  (newValue) => {
    if (newValue) {
      uploadProgress.value = 100;
    }
  }
);

// 上传前验证
const beforeUpload: UploadProps["beforeUpload"] = async (file) => {
  // 文件类型验证
  const isImage = file.type.startsWith("image/");
  if (!isImage) {
    ElMessage.error("只能上传图片文件!");
    return false;
  }

  // 文件大小验证
  const isLimitSize = file.size / 1024 / 1024 < props.maxSize;
  if (!isLimitSize) {
    ElMessage.error(`图片大小不能超过 ${props.maxSize}MB!`);
    return false;
  }

  // 如果使用OSS，获取上传策略
  if (props.useOss) {
    try {
      await getOssPolicy();
      return true;
    } catch (error) {
      ElMessage.error("获取上传凭证失败");
      return false;
    }
  }

  uploading.value = true;
  uploadProgress.value = 0;
  return true;
};

// 获取OSS上传策略
const getOssPolicy = async () => {
  return new Promise((resolve, reject) => {
    // 这里应该调用你的API获取OSS策略
    // 模拟API调用
    setTimeout(() => {
      dataObj.value = {
        policy: "mock-policy",
        signature: "mock-signature",
        OSSAccessKeyId: "mock-access-key",
        key: "upload/${filename}",
        dir: "upload",
        host: ossUploadUrl,
        success_action_status: "200",
        callback: "",
      };
      resolve(true);
    }, 100);
  });
};

// 上传进度
const handleProgress = (evt: any) => {
  uploadProgress.value = Math.round(evt.percent);
};

// 上传成功
const handleUploadSuccess = (response: any, file: UploadFile) => {
  uploading.value = false;
  uploadProgress.value = 100;

  let imageUrl = "";

  if (props.useOss) {
    // OSS上传成功，构建图片URL
    imageUrl = `${dataObj.value.host}/${dataObj.value.dir}/${file.name}`;
  } else {
    // MinIO或其他服务上传成功
    imageUrl = response.data?.url || response.url || "";
  }

  if (imageUrl) {
    emit("update:value", imageUrl);
    ElMessage.success("上传成功!");
  } else {
    ElMessage.error("上传失败，未获取到图片地址");
  }
};

// 上传失败
const handleUploadError = (error: any) => {
  uploading.value = false;
  uploadProgress.value = 0;
  console.error("Upload error:", error);
  ElMessage.error("上传失败，请重试!");
};

// 预览图片
const handlePreview = () => {
  if (imageUrl.value) {
    dialogVisible.value = true;
  }
};

// 删除图片
const handleRemove = async () => {
  try {
    await ElMessageBox.confirm("确定要删除这张图片吗？", "确认删除", {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
    });

    emit("update:value", "");
    ElMessage.success("删除成功!");
  } catch {
    // 用户取消删除
  }
};

// 下载图片
const downloadImage = () => {
  if (imageUrl.value) {
    const link = document.createElement("a");
    link.href = imageUrl.value;
    link.download = imageName.value || "image";
    link.target = "_blank";
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  }
};
</script>

<style scoped lang="scss">
.single-upload-container {
  .single-upload {
    :deep(.el-upload) {
      border: 2px dashed #dcdfe6;
      border-radius: 12px;
      cursor: pointer;
      position: relative;
      overflow: hidden;
      transition: all 0.3s ease;
      width: 200px;
      height: 200px;

      &:hover {
        border-color: #409eff;
        box-shadow: 0 2px 12px rgba(64, 158, 255, 0.2);
      }

      &.is-disabled {
        cursor: not-allowed;
        opacity: 0.6;

        &:hover {
          border-color: #dcdfe6;
          box-shadow: none;
        }
      }
    }
  }
}

.upload-content {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: relative;
}

// 图片预览
.image-preview {
  width: 100%;
  height: 100%;
  position: relative;
  border-radius: 10px;
  overflow: hidden;

  &:hover .image-actions {
    opacity: 1;
  }
}

.uploaded-image {
  width: 100%;
  height: 100%;

  :deep(.el-image__inner) {
    transition: transform 0.3s ease;
  }

  &:hover :deep(.el-image__inner) {
    transform: scale(1.05);
  }
}

.image-error {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: #c0c4cc;

  .el-icon {
    font-size: 40px;
    margin-bottom: 8px;
  }

  span {
    font-size: 14px;
  }
}

.image-actions {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  opacity: 0;
  transition: opacity 0.3s ease;

  .el-button {
    background: rgba(255, 255, 255, 0.2);
    border: 1px solid rgba(255, 255, 255, 0.3);
    backdrop-filter: blur(10px);

    &:hover {
      background: rgba(255, 255, 255, 0.3);
      transform: scale(1.1);
    }
  }
}

// 上传中状态
.uploading-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;

  .upload-progress {
    margin-bottom: 16px;
  }

  .uploading-text {
    color: #409eff;
    font-size: 14px;
    margin: 0;
  }
}

// 默认上传区域
.upload-area {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  padding: 20px;
  text-align: center;
}

.upload-icon {
  font-size: 48px;
  color: #c0c4cc;
  margin-bottom: 12px;
  transition: color 0.3s ease;
}

.upload-text {
  color: #606266;
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 8px;
}

.upload-tip {
  color: #909399;
  font-size: 12px;
  line-height: 1.4;
  max-width: 160px;
}

.single-upload:hover .upload-icon {
  color: #409eff;
}

// 预览对话框
.preview-dialog {
  :deep(.el-dialog) {
    border-radius: 12px;
    overflow: hidden;
  }

  :deep(.el-dialog__header) {
    background: linear-gradient(135deg, #f8f9fa, #ffffff);
    border-bottom: 1px solid #f0f0f0;
  }
}

.preview-content {
  text-align: center;
  padding: 20px 0;
}

.preview-image {
  max-width: 100%;
  max-height: 60vh;
  border-radius: 8px;
}

.preview-error {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 200px;
  color: #c0c4cc;

  .el-icon {
    font-size: 48px;
    margin-bottom: 12px;
  }

  span {
    font-size: 16px;
  }
}

// 响应式设计
@media (max-width: 768px) {
  .single-upload-container {
    .single-upload :deep(.el-upload) {
      width: 150px;
      height: 150px;
    }
  }

  .upload-icon {
    font-size: 36px;
  }

  .upload-text {
    font-size: 12px;
  }

  .upload-tip {
    font-size: 11px;
  }

  .preview-dialog {
    :deep(.el-dialog) {
      width: 90% !important;
      margin-top: 5vh;
    }
  }
}

// 小尺寸变体
.single-upload-container.small {
  .single-upload :deep(.el-upload) {
    width: 120px;
    height: 120px;
  }

  .upload-icon {
    font-size: 32px;
  }

  .upload-text {
    font-size: 12px;
  }
}

// 大尺寸变体
.single-upload-container.large {
  .single-upload :deep(.el-upload) {
    width: 300px;
    height: 200px;
  }

  .upload-icon {
    font-size: 60px;
  }

  .upload-text {
    font-size: 16px;
  }
}
</style>
