<template>
  <div class="forgot-container">
    <!-- 背景装饰 -->
    <div class="bg-decoration"></div>

    <el-card class="forgot-card" shadow="always">
      <!-- 卡片顶部装饰条 -->
      <div class="card-header-decoration"></div>

      <!-- 步骤指示器 -->
      <el-steps
        :active="currentStep"
        finish-status="success"
        align-center
        class="steps-container"
      >
        <el-step title="验证手机" description="输入手机号"></el-step>
        <el-step title="验证身份" description="短信验证码"></el-step>
        <el-step title="重置密码" description="设置新密码"></el-step>
        <el-step title="完成" description="密码重置成功"></el-step>
      </el-steps>

      <!-- 页面头部 -->
      <div class="page-header">
        <div class="logo-container">
          <el-icon class="logo-icon" :size="48">
            <Key />
          </el-icon>
        </div>
        <h1 class="page-title">{{ stepTitles[currentStep] }}</h1>
        <p class="page-subtitle">{{ stepSubtitles[currentStep] }}</p>
      </div>

      <!-- 步骤1: 手机号验证 -->
      <div v-show="currentStep === 0" class="step-content">
        <el-form
          ref="phoneFormRef"
          :model="phoneForm"
          :rules="phoneRules"
          class="form-container"
        >
          <el-form-item prop="mobile" class="form-item">
            <el-input
              v-model="phoneForm.mobile"
              size="large"
              placeholder="请输入注册时使用的手机号"
              clearable
            >
              <template #prefix>
                <el-icon class="input-icon">
                  <Iphone />
                </el-icon>
              </template>
            </el-input>
          </el-form-item>

          <el-form-item prop="captcha" class="form-item">
            <div class="captcha-container">
              <el-input
                v-model="phoneForm.captcha"
                size="large"
                placeholder="请输入图形验证码"
                class="captcha-input"
              >
                <template #prefix>
                  <el-icon class="input-icon">
                    <Grid />
                  </el-icon>
                </template>
              </el-input>
              <div class="captcha-image" @click="refreshCaptcha">
                <img
                  v-if="captcha.picPath"
                  :src="captcha.picPath"
                  alt="验证码"
                  class="captcha-img"
                />
                <div v-else class="captcha-placeholder">
                  <el-icon><RefreshRight /></el-icon>
                  <span>{{ captchaText }}</span>
                </div>
              </div>
            </div>
          </el-form-item>

          <el-button
            type="primary"
            size="large"
            class="step-button"
            :loading="loading"
            @click="verifyPhone"
          >
            发送验证码
          </el-button>
        </el-form>
      </div>

      <!-- 步骤2: 短信验证 -->
      <div v-show="currentStep === 1" class="step-content">
        <div class="phone-info">
          <el-icon class="phone-icon"><Message /></el-icon>
          <p>验证码已发送至</p>
          <p class="phone-number">{{ maskedPhone }}</p>
        </div>

        <el-form
          ref="codeFormRef"
          :model="codeForm"
          :rules="codeRules"
          class="form-container"
        >
          <el-form-item prop="smsCode" class="form-item">
            <el-input
              v-model="codeForm.smsCode"
              size="large"
              placeholder="请输入6位短信验证码"
              maxlength="6"
              clearable
              @input="onCodeInput"
            >
              <template #prefix>
                <el-icon class="input-icon">
                  <ChatDotSquare />
                </el-icon>
              </template>
            </el-input>
          </el-form-item>

          <div class="resend-container">
            <span class="resend-text">没收到验证码？</span>
            <el-button
              type="text"
              class="resend-button"
              :disabled="countDown > 0"
              @click="resendCode"
            >
              {{ countDown > 0 ? `${countDown}s后重新发送` : "重新发送" }}
            </el-button>
          </div>

          <el-button
            type="primary"
            size="large"
            class="step-button"
            :loading="loading"
            @click="verifyCode"
          >
            验证
          </el-button>
        </el-form>
      </div>

      <!-- 步骤3: 重置密码 -->
      <div v-show="currentStep === 2" class="step-content">
        <el-form
          ref="passwordFormRef"
          :model="passwordForm"
          :rules="passwordRules"
          class="form-container"
        >
          <el-form-item prop="newPassword" class="form-item">
            <el-input
              v-model="passwordForm.newPassword"
              :type="pwdType1"
              size="large"
              placeholder="请输入新密码"
              clearable
            >
              <template #prefix>
                <el-icon class="input-icon">
                  <Lock />
                </el-icon>
              </template>
              <template #suffix>
                <el-icon class="password-toggle" @click="togglePassword(1)">
                  <View v-if="pwdType1 === 'password'" />
                  <Hide v-else />
                </el-icon>
              </template>
            </el-input>
          </el-form-item>

          <el-form-item prop="confirmPassword" class="form-item">
            <el-input
              v-model="passwordForm.confirmPassword"
              :type="pwdType2"
              size="large"
              placeholder="请确认新密码"
              clearable
              @keyup.enter="resetPassword"
            >
              <template #prefix>
                <el-icon class="input-icon">
                  <Lock />
                </el-icon>
              </template>
              <template #suffix>
                <el-icon class="password-toggle" @click="togglePassword(2)">
                  <View v-if="pwdType2 === 'password'" />
                  <Hide v-else />
                </el-icon>
              </template>
            </el-input>
          </el-form-item>

          <!-- 密码强度指示器 -->
          <div class="password-strength">
            <div class="strength-label">密码强度:</div>
            <div class="strength-bar">
              <div
                class="strength-indicator"
                :class="passwordStrength.level"
                :style="{ width: passwordStrength.width }"
              ></div>
            </div>
            <div class="strength-text">{{ passwordStrength.text }}</div>
          </div>

          <el-button
            type="primary"
            size="large"
            class="step-button"
            :loading="loading"
            @click="resetPassword"
          >
            重置密码
          </el-button>
        </el-form>
      </div>

      <!-- 步骤4: 完成 -->
      <div v-show="currentStep === 3" class="step-content success-content">
        <div class="success-icon">
          <el-icon :size="80" color="#67c23a">
            <CircleCheckFilled />
          </el-icon>
        </div>
        <h2 class="success-title">密码重置成功！</h2>
        <p class="success-text">您的密码已经成功重置，请使用新密码登录</p>

        <div class="success-actions">
          <el-button type="primary" size="large" @click="goToLogin">
            立即登录
          </el-button>
          <el-button size="large" @click="resetProcess"> 重新开始 </el-button>
        </div>
      </div>

      <!-- 返回登录 -->
      <div class="back-to-login">
        <el-button type="text" class="back-button" @click="goToLogin">
          <el-icon><ArrowLeft /></el-icon>
          返回登录
        </el-button>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, computed, onUnmounted, onMounted } from "vue";
import {
  Key,
  Iphone,
  Grid,
  RefreshRight,
  Message,
  ChatDotSquare,
  Lock,
  View,
  Hide,
  CircleCheckFilled,
  ArrowLeft,
} from "@element-plus/icons-vue";
import type { FormInstance, FormRules } from "element-plus";
import { ElMessage } from "element-plus";
import { navigateTo } from "@/utils/support";
import { useRouter, useRoute } from "vue-router";
import { getCaptcha } from "@/api/common";
import { ICaptcha } from "@/api/types/common";

const router = useRouter();
const route = useRoute();

// 表单引用
const phoneFormRef = ref<FormInstance>();
const codeFormRef = ref<FormInstance>();
const passwordFormRef = ref<FormInstance>();

// 验证码数据
const captcha = reactive<ICaptcha>({
  captchaId: "",
  imageBase64: "",
});

// 当前步骤
const currentStep = ref(0);

// 步骤标题和副标题
const stepTitles = ["验证手机号", "验证身份", "重置密码", "重置完成"];
const stepSubtitles = [
  "请输入您注册时使用的手机号码",
  "请输入您收到的短信验证码",
  "请设置您的新密码",
  "密码重置成功，请重新登录",
];

// 手机号验证表单
const phoneForm = reactive({
  mobile: "",
  captcha: "",
});

// 验证码表单
const codeForm = reactive({
  smsCode: "",
});

// 密码重置表单
const passwordForm = reactive({
  newPassword: "",
  confirmPassword: "",
});

// 状态管理
const loading = ref(false);
const captchaText = ref("点击刷新");
const countDown = ref(0);
const pwdType1 = ref("password");
const pwdType2 = ref("password");

// 定时器
let timer: number | null = null;

// 验证规则
const phoneRules: FormRules = {
  mobile: [
    { required: true, message: "请输入手机号码", trigger: "blur" },
    {
      pattern: /^1[3-9]\d{9}$/,
      message: "请输入正确的手机号码",
      trigger: "blur",
    },
  ],
  captcha: [{ required: true, message: "请输入图形验证码", trigger: "blur" }],
};

const codeRules: FormRules = {
  smsCode: [
    { required: true, message: "请输入验证码", trigger: "blur" },
    { pattern: /^\d{6}$/, message: "请输入6位数字验证码", trigger: "blur" },
  ],
};

const passwordRules: FormRules = {
  newPassword: [
    { required: true, message: "请输入新密码", trigger: "blur" },
    { min: 6, message: "密码长度至少6位", trigger: "blur" },
    { max: 20, message: "密码长度不超过20位", trigger: "blur" },
  ],
  confirmPassword: [
    { required: true, message: "请确认新密码", trigger: "blur" },
    {
      validator: (rule, value, callback) => {
        if (value !== passwordForm.newPassword) {
          callback(new Error("两次输入的密码不一致"));
        } else {
          callback();
        }
      },
      trigger: "blur",
    },
  ],
};

// 计算属性
const maskedPhone = computed(() => {
  if (phoneForm.mobile) {
    return phoneForm.mobile.replace(/(\d{3})\d{4}(\d{4})/, "$1****$2");
  }
  return "";
});

// 密码强度计算
const passwordStrength = computed(() => {
  const password = passwordForm.newPassword;
  if (!password) return { level: "weak", width: "0%", text: "" };

  let score = 0;
  // 长度检查
  if (password.length >= 8) score += 2;
  else if (password.length >= 6) score += 1;

  // 复杂性检查
  if (/[a-z]/.test(password)) score += 1;
  if (/[A-Z]/.test(password)) score += 1;
  if (/\d/.test(password)) score += 1;
  if (/[!@#$%^&*(),.?":{}|<>]/.test(password)) score += 1;

  if (score <= 2) return { level: "weak", width: "33%", text: "弱" };
  if (score <= 4) return { level: "medium", width: "66%", text: "中" };
  return { level: "strong", width: "100%", text: "强" };
});

// 方法
const refreshCaptcha = async () => {
  captchaText.value = "刷新中...";
  try {
    const res = await getCaptcha();
    captcha.captchaId = res.data.captchaId;
    captcha.imageBase64 = res.data.imageBase64;
  } catch (error) {
    captchaText.value = "ABCD";
  }
  ElMessage.success("验证码已刷新");
};

const verifyPhone = async () => {
  if (!phoneFormRef.value) return;

  try {
    await phoneFormRef.value.validate();
    loading.value = true;

    // 模拟API调用
    setTimeout(() => {
      loading.value = false;
      currentStep.value = 1;
      startCountDown();
      ElMessage.success("验证码已发送");
    }, 1500);
  } catch (error) {
    console.log("表单验证失败:", error);
  }
};

const startCountDown = () => {
  countDown.value = 60;
  timer = setInterval(() => {
    countDown.value--;
    if (countDown.value <= 0) {
      clearInterval(timer!);
      timer = null;
    }
  }, 1000);
};

const resendCode = () => {
  if (countDown.value > 0) return;

  loading.value = true;
  setTimeout(() => {
    loading.value = false;
    startCountDown();
    ElMessage.success("验证码已重新发送");
  }, 1000);
};

const onCodeInput = (value: string) => {
  // 自动验证6位数字
  if (value.length === 6 && /^\d{6}$/.test(value)) {
    setTimeout(() => verifyCode(), 300);
  }
};

const verifyCode = async () => {
  if (!codeFormRef.value) return;

  try {
    await codeFormRef.value.validate();
    loading.value = true;

    // 模拟API调用
    setTimeout(() => {
      loading.value = false;
      if (codeForm.smsCode === "123456") {
        currentStep.value = 2;
        ElMessage.success("验证成功");
      } else {
        ElMessage.error("验证码错误，请重新输入");
      }
    }, 1500);
  } catch (error) {
    console.log("表单验证失败:", error);
  }
};

const togglePassword = (type: number) => {
  if (type === 1) {
    pwdType1.value = pwdType1.value === "password" ? "text" : "password";
  } else {
    pwdType2.value = pwdType2.value === "password" ? "text" : "password";
  }
};

const resetPassword = async () => {
  if (!passwordFormRef.value) return;

  try {
    await passwordFormRef.value.validate();
    loading.value = true;

    // 模拟API调用
    setTimeout(() => {
      loading.value = false;
      currentStep.value = 3;
      ElMessage.success("密码重置成功");
    }, 2000);
  } catch (error) {
    console.log("表单验证失败:", error);
  }
};

const goToLogin = () => {
  navigateTo(route, router, "/login", undefined, true);
  ElMessage.info("跳转到登录页面");
};

const resetProcess = () => {
  currentStep.value = 0;
  phoneForm.mobile = "";
  phoneForm.captcha = "";
  codeForm.smsCode = "";
  passwordForm.newPassword = "";
  passwordForm.confirmPassword = "";
  countDown.value = 0;
  if (timer) {
    clearInterval(timer);
    timer = null;
  }
};

// 清理定时器
onUnmounted(() => {
  if (timer) {
    clearInterval(timer);
  }
});

onMounted(() => {
  refreshCaptcha();
});
</script>

<style scoped lang="scss">
.forgot-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  position: relative;
  overflow: hidden;
}

// 背景装饰
.bg-decoration {
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background-image: radial-gradient(
      circle at 25% 25%,
      rgba(255, 255, 255, 0.1) 2px,
      transparent 2px
    ),
    radial-gradient(
      circle at 75% 75%,
      rgba(255, 255, 255, 0.1) 1px,
      transparent 1px
    );
  background-size: 100px 100px, 60px 60px;
  animation: float 20s infinite linear;
  pointer-events: none;
}

@keyframes float {
  0% {
    transform: translate(0, 0) rotate(0deg);
  }
  100% {
    transform: translate(-50px, -50px) rotate(360deg);
  }
}

// 忘记密码卡片
.forgot-card {
  width: 520px;
  max-width: 95vw;
  border-radius: 20px;
  border: none;
  backdrop-filter: blur(20px);
  background: rgba(255, 255, 255, 0.95);
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.2);
  overflow: hidden;
  position: relative;
  animation: slideUp 0.8s ease;

  :deep(.el-card__body) {
    padding: 40px;
  }
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

// 卡片顶部装饰条
.card-header-decoration {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #f56c6c, #e6a23c, #67c23a, #409eff);
}

// 步骤指示器
.steps-container {
  margin-bottom: 40px;

  :deep(.el-step__title) {
    font-size: 12px;
    line-height: 1.2;
  }

  :deep(.el-step__description) {
    font-size: 11px;
    margin-top: 2px;
  }
}

// 页面头部
.page-header {
  text-align: center;
  margin-bottom: 40px;
}

.logo-container {
  margin-bottom: 20px;
}

.logo-icon {
  width: 64px;
  height: 64px;
  background: linear-gradient(135deg, #f56c6c, #ff7875);
  border-radius: 16px;
  color: white;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 8px 25px rgba(245, 108, 108, 0.3);
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%,
  100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
}

.page-title {
  font-size: 24px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0 0 8px 0;
}

.page-subtitle {
  font-size: 14px;
  color: #7f8c8d;
  margin: 0;
}

// 步骤内容
.step-content {
  min-height: 280px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

// 表单样式
.form-container {
  width: 100%;

  .form-item {
    margin-bottom: 24px;

    :deep(.el-form-item__content) {
      width: 100%;
    }

    :deep(.el-input) {
      width: 100%;
    }

    :deep(.el-input__wrapper) {
      width: 100%;
      border-radius: 12px;
      box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
      transition: all 0.3s ease;

      &.is-focus {
        transform: translateY(-2px);
        box-shadow: 0 4px 15px rgba(64, 158, 255, 0.2);
      }
    }
  }
}

.input-icon {
  color: #409eff;
  font-size: 18px;
}

.password-toggle {
  cursor: pointer;
  color: #7f8c8d;
  transition: color 0.3s ease;

  &:hover {
    color: #409eff;
  }
}

// 验证码容器
.captcha-container {
  display: flex;
  gap: 12px;
  width: 100%;
  align-items: flex-start;
}

.captcha-input {
  flex: 1;

  :deep(.el-input) {
    width: 100%;
  }
}

.captcha-image {
  width: 120px;
  height: 40px;
  border: 2px solid #dcdfe6;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;

  &:hover {
    border-color: #409eff;
    transform: scale(1.02);
  }
}

.captcha-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.captcha-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  background: linear-gradient(45deg, #f8f9fa, #e9ecef);
  color: #7f8c8d;
  font-size: 12px;

  .el-icon {
    margin-bottom: 4px;
  }
}

// 手机号信息显示
.phone-info {
  text-align: center;
  margin-bottom: 30px;
  padding: 20px;
  background: linear-gradient(135deg, #e8f4fd, #f0f8ff);
  border-radius: 12px;
  border: 1px solid #409eff;
}

.phone-icon {
  font-size: 32px;
  color: #409eff;
  margin-bottom: 10px;
}

.phone-info p {
  margin: 5px 0;
  color: #2c3e50;
}

.phone-number {
  font-size: 18px;
  font-weight: 600;
  font-family: "Monaco", "Consolas", monospace;
}

// 重发验证码
.resend-container {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  margin-bottom: 24px;
}

.resend-text {
  color: #7f8c8d;
  font-size: 14px;
}

.resend-button {
  color: #409eff;
  padding: 0;

  &.is-disabled {
    color: #c0c4cc;
  }
}

// 密码强度指示器
.password-strength {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 24px;
  font-size: 14px;
}

.strength-label {
  color: #7f8c8d;
  white-space: nowrap;
}

.strength-bar {
  flex: 1;
  height: 6px;
  background: #f5f7fa;
  border-radius: 3px;
  overflow: hidden;
}

.strength-indicator {
  height: 100%;
  border-radius: 3px;
  transition: all 0.3s ease;

  &.weak {
    background: #f56c6c;
  }
  &.medium {
    background: #e6a23c;
  }
  &.strong {
    background: #67c23a;
  }
}

.strength-text {
  color: #7f8c8d;
  white-space: nowrap;
  min-width: 20px;
}

// 步骤按钮
.step-button {
  width: 100%;
  height: 50px;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 600;
  transition: all 0.3s ease;

  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 8px 25px rgba(64, 158, 255, 0.4);
  }

  &:active {
    transform: translateY(0);
  }
}

// 成功页面
.success-content {
  text-align: center;
  align-items: center;
}

.success-icon {
  margin-bottom: 20px;
  animation: bounceIn 0.8s ease;
}

@keyframes bounceIn {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.1);
  }
  100% {
    transform: scale(1);
  }
}

.success-title {
  font-size: 24px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0 0 10px 0;
}

.success-text {
  font-size: 16px;
  color: #7f8c8d;
  margin: 0 0 30px 0;
}

.success-actions {
  display: flex;
  gap: 16px;
  justify-content: center;
  width: 100%;

  .el-button {
    flex: 1;
    max-width: 150px;
  }
}

// 返回登录
.back-to-login {
  text-align: center;
  margin-top: 30px;
  padding-top: 20px;
  border-top: 1px solid #f0f0f0;
}

.back-button {
  color: #7f8c8d;
  font-size: 14px;

  &:hover {
    color: #409eff;
  }

  .el-icon {
    margin-right: 4px;
  }
}

// 响应式设计
@media (max-width: 600px) {
  .forgot-card {
    width: 100%;
    margin: 10px;

    :deep(.el-card__body) {
      padding: 30px 20px;
    }
  }

  .steps-container {
    :deep(.el-step__description) {
      display: none;
    }
  }

  .page-title {
    font-size: 20px;
  }

  .captcha-container {
    flex-direction: column;

    .captcha-input {
      width: 100%;
    }
  }

  .captcha-image {
    width: 100%;
    height: 45px;
  }

  .success-actions {
    flex-direction: column;

    .el-button {
      max-width: none;
      width: 100%;
    }
  }

  .password-strength {
    flex-wrap: wrap;
    gap: 5px;
  }
}
</style>
