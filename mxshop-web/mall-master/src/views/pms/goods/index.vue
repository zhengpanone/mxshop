// 搜索卡片 .search-card { border-radius: 16px; border: none; background:
rgba(255, 255, 255, 0.95); backdrop-filter: blur(10px); box-shadow: 0 8px 32px
rgba(0, 0, 0, 0.1); border: 1px solid rgba(255, 255, 255, 0.2); overflow:
hidden; :deep(.el-card__body) { padding: 0; } }
<template>
  <div class="product-list-container">
    <!-- 搜索筛选区域 -->
    <div class="search-section">
      <el-card class="search-card" shadow="hover">
        <div class="search-header">
          <div class="header-info">
            <el-icon class="search-icon"><Search /></el-icon>
            <span class="header-title">筛选搜索</span>
            <el-tag type="primary" class="total-tag">{{ total }} 个商品</el-tag>
          </div>
          <div class="search-actions">
            <el-button type="primary" @click="handleAddProduct" class="add-btn">
              <el-icon><Plus /></el-icon>
              新增商品
            </el-button>
            <el-button
              @click="handleResetSearch"
              :icon="RefreshLeft"
              class="reset-btn"
            >
              重置
            </el-button>
            <el-button
              type="primary"
              @click="handleSearchList"
              :icon="Search"
              :loading="listLoading"
              class="search-btn"
            >
              查询结果
            </el-button>
          </div>
        </div>

        <div class="search-form">
          <el-form
            :model="goodsParams"
            label-width="100px"
            class="search-form-content"
          >
            <el-row :gutter="20">
              <el-col :span="8">
                <el-form-item label="商品名称">
                  <el-input
                    v-model="goodsParams.productName"
                    placeholder="请输入商品名称"
                    clearable
                    @keyup.enter="handleSearchList"
                    class="search-input"
                  >
                    <template #prefix>
                      <el-icon><Search /></el-icon>
                    </template>
                  </el-input>
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="商品分类">
                  <el-cascader
                    v-model="goodsParams.categoryId"
                    placeholder="请选择商品分类"
                    :options="productCateOptions"
                    :props="{
                      value: 'id',
                      label: 'name',
                      children: 'sub_category',
                    }"
                    clearable
                    filterable
                    @change="getBrand"
                    class="category-select"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="商品品牌">
                  <el-select
                    v-model="goodsParams.brandId"
                    placeholder="请选择品牌"
                    clearable
                    filterable
                    class="brand-select"
                  >
                    <el-option
                      v-for="item in brandOptions"
                      :key="item.id"
                      :label="item.name"
                      :value="item.id"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
        </div>
      </el-card>
    </div>

    <!-- 数据表格 -->
    <el-card class="table-card" shadow="hover">
      <div class="table-header">
        <div class="header-info">
          <el-icon class="table-icon"><Grid /></el-icon>
          <span class="header-title">商品列表</span>
        </div>
        <div class="table-actions">
          <!-- 批量操作 -->
          <el-dropdown
            @command="handleBatchCommand"
            v-if="multipleSelection.length > 0"
          >
            <el-button type="warning" size="small" class="batch-btn">
              批量操作 ({{ multipleSelection.length }})
              <el-icon class="el-icon--right"><ArrowDown /></el-icon>
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="publishOn"
                  >批量上架</el-dropdown-item
                >
                <el-dropdown-item command="publishOff"
                  >批量下架</el-dropdown-item
                >
                <el-dropdown-item command="newOn">设为新品</el-dropdown-item>
                <el-dropdown-item command="newOff">取消新品</el-dropdown-item>
                <el-dropdown-item command="hotOn">设为推荐</el-dropdown-item>
                <el-dropdown-item command="hotOff">取消推荐</el-dropdown-item>
                <el-dropdown-item command="delete" divided
                  >批量删除</el-dropdown-item
                >
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>

      <div class="table-container">
        <el-table
          ref="productTable"
          :data="productList"
          v-loading="listLoading"
          @selection-change="handleSelectionChange"
          class="product-table"
          stripe
          header-row-class-name="table-header-row"
        >
          <el-table-column type="selection" width="55" align="center" />

          <el-table-column label="序号" width="80" align="center">
            <template #default="{ $index }">
              <div class="row-number">
                {{
                  (goodsParams.pageNum - 1) * goodsParams.pageSize + $index + 1
                }}
              </div>
            </template>
          </el-table-column>

          <el-table-column label="商品信息" min-width="200">
            <template #default="{ row }">
              <div class="product-info">
                <div class="product-image">
                  <el-image
                    :src="row.pic || 'https://via.placeholder.com/60x60'"
                    fit="cover"
                    class="product-img"
                  >
                    <template #error>
                      <div class="image-slot">
                        <el-icon><Picture /></el-icon>
                      </div>
                    </template>
                  </el-image>
                </div>
                <div class="product-details">
                  <h4 class="product-name">{{ row.name }}</h4>
                  <p class="product-sn">货号: {{ row.productSn || "暂无" }}</p>
                </div>
              </div>
            </template>
          </el-table-column>

          <el-table-column label="品牌" width="120">
            <template #default="{ row }">
              <el-tag
                type="info"
                size="small"
                v-if="row.brand"
                class="brand-tag"
              >
                {{ row.brand.name }}
              </el-tag>
              <span v-else class="text-muted">暂无品牌</span>
            </template>
          </el-table-column>

          <el-table-column label="分类" width="120">
            <template #default="{ row }">
              <el-tag
                type="warning"
                size="small"
                v-if="row.category"
                class="category-tag"
              >
                {{ row.category.name }}
              </el-tag>
              <span v-else class="text-muted">暂无分类</span>
            </template>
          </el-table-column>

          <el-table-column label="价格" width="100" align="right">
            <template #default="{ row }">
              <div class="price-info">
                <span class="price-symbol">¥</span>
                <span class="price-value">{{ row.shop_price || 0 }}</span>
              </div>
            </template>
          </el-table-column>

          <el-table-column label="状态标签" width="150" align="center">
            <template #default="{ row, $index }">
              <div class="status-switches">
                <div class="switch-item">
                  <el-switch
                    v-model="row.publishStatus"
                    :active-value="true"
                    :inactive-value="false"
                    active-text="上架"
                    inactive-text="下架"
                    size="small"
                    @change="handlePublishStatusChange('sale', $index, row)"
                  />
                </div>
                <div class="switch-item">
                  <el-switch
                    v-model="row.newStatus"
                    :active-value="true"
                    :inactive-value="false"
                    active-text="新品"
                    inactive-text=""
                    size="small"
                    @change="handlePublishStatusChange('new', $index, row)"
                  />
                </div>
                <div class="switch-item">
                  <el-switch
                    v-model="row.recommandStatus"
                    :active-value="true"
                    :inactive-value="false"
                    active-text="推荐"
                    inactive-text=""
                    size="small"
                    @change="handlePublishStatusChange('hot', $index, row)"
                  />
                </div>
              </div>
            </template>
          </el-table-column>

          <el-table-column
            label="操作"
            width="160"
            align="center"
            fixed="right"
          >
            <template #default="{ row, $index }">
              <div class="action-buttons">
                <el-button
                  type="primary"
                  size="small"
                  @click="handleShowProduct($index, row)"
                  :icon="View"
                  class="view-btn"
                >
                  查看
                </el-button>
                <el-button
                  type="danger"
                  size="small"
                  @click="handleDelete($index, row)"
                  :icon="Delete"
                  class="delete-btn"
                >
                  删除
                </el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>

    <!-- 分页 -->
    <div class="pagination-container">
      <el-pagination
        v-model:current-page="goodsParams.pageNum"
        v-model:page-size="goodsParams.pageSize"
        :page-sizes="[5, 10, 20, 50]"
        :total="total"
        background
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        class="custom-pagination"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { useRouter, useRoute } from "vue-router";
import {
  fetchList,
  updateDeleteStatus,
  updateNewStatus,
  updateRecommendStatus,
  updatePublishStatus,
} from "@/api/pms/product";
import {
  fetchList as fetchSkuStockList,
  update as updateSkuStockList,
} from "@/api/pms/skuStock";
import { fetchList as fetchProductAttrList } from "@/api/pms/productAttr";
import { fetchList as fetchBrandList } from "@/api/pms/brand";
import { fetchListWithChildren } from "@/api/pms/productCategory";
import {
  getGoodsPageList,
  deleteGoods,
  getBrands,
  putGoodsStatus,
  getBrandsByCate,
} from "@/api/pms/goods";
import { getCategoryList } from "@/api/pms/category";
import { CategoryData, BrandData, GoodsParams } from "@/api/pms/types/category";
import { ProductData } from "@/api/pms/types/product";
import {
  Search,
  View,
  Grid,
  Plus,
  RefreshLeft,
  Delete,
} from "@element-plus/icons-vue";

const router = useRouter();

const goodsParams = reactive<GoodsParams>({
  pageNum: 1,
  pageSize: 20,
  brandId: "",
  categoryId: "",
  productName: "",
});

const editSkuInfo = reactive({
  dialogVisible: false,
  productId: null,
  productSn: "",
  productAttributeCategoryId: null,
  stockList: [],
  productAttr: [],
  keyword: null,
});

const operates = [
  { label: "商品上架", value: "publishOn" },
  { label: "新品", value: "publishOn" },
  { label: "推荐", value: "publishOn" },
  { label: "删除", value: "delete" },
];

const operateType = ref(null);
const productList = ref<ProductData[]>([]);
const total = ref<number>(0);
const listLoading = ref(true);
const selectProductCateValue = ref(null);
const multipleSelection = ref([]);
const productCateOptions = ref<CategoryData[]>([]);
const brandOptions = ref<BrandData[]>([]);

const publishStatusOptions = [
  { value: 1, label: "上架" },
  { value: 0, label: "下架" },
];

const verifyStatusOptions = [
  { value: 1, label: "审核通过" },
  { value: 0, label: "未审核" },
];

const getBrand = async (id: string[]) => {
  if (!id || id.length === 0) {
    goodsParams.categoryId = null;
    goodsParams.brandId = "";
    brandOptions.value = [];
    return;
  }
  const currenntCateId = id[id.length - 1];
  goodsParams.categoryId = currenntCateId;
  goodsParams.brandId = "";
  try {
    const res = await getBrandsByCate(currenntCateId, null);
    brandOptions.value = res.data;
  } catch (error) {
    console.error("获取品牌失败:", error);
    brandOptions.value = [];
  }
};

const getGoodsList = async () => {
  listLoading.value = true;
  try {
    const response = await getGoodsPageList({ ...goodsParams });
    productList.value = response.data.list;
    total.value = response.data.total;
  } finally {
    listLoading.value = false;
  }
};

const getBrandList = async () => {
  const response = await getBrands(null);
  brandOptions.value = response.data.data;
};

const getProductCateList = async () => {
  const response = await getCategoryList({});
  console.log("getProductCateList", JSON.stringify(response.data));
  productCateOptions.value = response.data;
};

const handleSearchEditSku = async () => {
  const response = await fetchSkuStockList(editSkuInfo.productId, {
    keyword: editSkuInfo.keyword,
  });
  editSkuInfo.stockList = response.data;
};

const handleSearchList = () => {
  goodsParams.pageNum = 1;
  getGoodsList();
};

const handleAddProduct = () => {
  router.push({ path: "/product/product_add" });
};

const handleBatchOperate = async () => {
  if (operateType.value == null) {
    ElMessage({
      message: "请选择操作类型",
      type: "warning",
      duration: 1000,
    });
    return;
  }
  if (multipleSelection.value == null || multipleSelection.value.length < 1) {
    ElMessage({
      message: "请选择要操作的商品",
      type: "warning",
      duration: 1000,
    });
    return;
  }
  try {
    await ElMessageBox.confirm("是否要进行该批量操作?", "提示", {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
    });

    const ids = multipleSelection.value.map((item) => item.id);
    switch (operateType.value) {
      case operates[0].value:
        await handleUpdatePublishStatus(true, ids, null);
        break;
      case operates[1].value:
        await handleUpdatePublishStatus(false, ids, null);
        break;
      case operates[2].value:
        await handleUpdateRecommendStatus(1, ids);
        break;
      case operates[3].value:
        await handleUpdateRecommendStatus(0, ids);
        break;
      case operates[4].value:
        await handleUpdateNewStatus(1, ids);
        break;
      case operates[5].value:
        await handleUpdateNewStatus(0, ids);
        break;
      case operates[6].value:
        break;
      case operates[7].value:
        await handleUpdateDeleteStatus(1, ids);
        break;
      default:
        break;
    }
    await getGoodsList();
  } catch (error) {
    console.error(error);
  }
};

const handleSizeChange = (val: number) => {
  goodsParams.pageNum = 1;
  goodsParams.pageSize = val;
  getGoodsList();
};

const handleCurrentChange = (val: number) => {
  goodsParams.pageNum = val;
  getGoodsList();
};

const handleSelectionChange = (val: any) => {
  multipleSelection.value = val;
};

const handlePublishStatusChange = (paramname: any, index: any, row: any) => {
  if (index === 0) {
    handleUpdatePublishStatus(paramname, false, row);
  } else if (index === 1) {
    handleUpdatePublishStatus(paramname, true, row);
  }
};

const handleRecommendStatusChange = (index: any, row: any) => {
  const ids = [row.id];
  handleUpdateRecommendStatus(row.recommandStatus, ids);
};

const handleResetSearch = () => {
  Object.assign(goodsParams, {
    productName: "",
    categoryId: "",
    brandId: "",
    pageNum: 1,
  });
};

const handleDelete = async (index: any, row: any) => {
  try {
    await ElMessageBox.confirm("是否要进行删除商品?", "提示", {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
    });
    await deleteGoods(row.id, null);
    ElMessage({
      message: "删除成功",
      type: "success",
      duration: 1000,
    });
    await getGoodsList();
  } catch (error) {
    console.error(error);
  }
};

const handleUpdateProduct = (index: any, row: any) => {
  router.push({ path: "/updateProduct", query: { id: row.id } });
};

const handleShowProduct = (index: any, row: any) => {
  router.push({ path: "/updateProduct", query: { id: row.id } });
};

const handleShowVerifyDetail = (index: any, row: any) => {
  console.log("handleShowVerifyDetail", row);
};

const handleShowLog = (index: any, row: any) => {
  console.log("handleShowLog", row);
};

const handleUpdatePublishStatus = async (
  paramname: any,
  param: any,
  row: any
) => {
  const params = {
    sale: row.on_sale,
    hot: row.is_hot,
    new: row.is_new,
  };
  params[paramname] = param;

  await putGoodsStatus(row.id, params);
  ElMessage({
    message: "修改成功",
    type: "success",
    duration: 1000,
  });
};

const handleUpdateNewStatus = async (newStatus: any, ids: any) => {
  const params = new URLSearchParams();
  params.append("ids", ids);
  params.append("newStatus", newStatus);

  await updateNewStatus(params);
  ElMessage({
    message: "修改成功",
    type: "success",
    duration: 1000,
  });
};

const handleUpdateRecommendStatus = async (recommendStatus: any, ids: any) => {
  const params = new URLSearchParams();
  params.append("ids", ids);
  params.append("recommendStatus", recommendStatus);

  await updateRecommendStatus(params);
  ElMessage({
    message: "修改成功",
    type: "success",
    duration: 1000,
  });
};

const handleUpdateDeleteStatus = async (deleteStatus: any, ids: any) => {
  const params = new URLSearchParams();
  params.append("ids", ids);
  params.append("deleteStatus", deleteStatus);
  try {
    await updateDeleteStatus(params);
    ElMessage({
      message: "删除成功",
      type: "success",
      duration: 1000,
    });
    await getGoodsList();
  } catch (err) {
    let errorMessage = "发生未知错误";
    console.error(err);
    if (err instanceof Error) {
      ElMessage({
        message: err.message,
        type: "error",
        duration: 1000,
      });
    } else if (typeof err === "string") {
      errorMessage = err;
    } else if (err && typeof err === "object" && "msg" in err) {
      errorMessage = (err as any).msg;
    }
  }
};

onMounted(() => {
  getGoodsList();
  getBrandList();
  getProductCateList();
});
</script>

<style scoped lang="scss">
.product-list-container {
  padding: 24px;
  // background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
}

// 搜索区域
.search-section {
  margin-bottom: 24px;
}

.search-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 24px;
  background: linear-gradient(135deg, #f8f9ff, #ffffff);
  border-bottom: 1px solid rgba(102, 126, 234, 0.1);
}

.header-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.search-icon,
.table-icon {
  color: #667eea;
  font-size: 18px;
}

.header-title {
  font-size: 16px;
  font-weight: 600;
  color: #2c3e50;
}

.total-tag {
  background: linear-gradient(135deg, #667eea, #764ba2);
  border: none;
  color: white;
  font-weight: 500;
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 12px;
}

.search-actions {
  display: flex;
  gap: 12px;

  .add-btn {
    background: linear-gradient(135deg, #667eea, #764ba2);
    border: none;
    border-radius: 8px;
    font-weight: 500;
    box-shadow: 0 4px 15px rgba(102, 126, 234, 0.3);
    transition: all 0.3s ease;

    &:hover {
      transform: translateY(-2px);
      box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
    }

    &:active {
      transform: translateY(0);
    }
  }

  .reset-btn {
    border-radius: 8px;
    border-color: #e2e8f0;
    color: #64748b;
    transition: all 0.3s ease;

    &:hover {
      border-color: #667eea;
      color: #667eea;
      transform: translateY(-1px);
    }
  }

  .search-btn {
    background: linear-gradient(135deg, #667eea, #764ba2);
    border: none;
    border-radius: 8px;
    font-weight: 500;
    transition: all 0.3s ease;

    &:hover {
      transform: translateY(-1px);
      box-shadow: 0 4px 15px rgba(102, 126, 234, 0.3);
    }
  }
}

.search-form {
  padding: 24px;
}

.search-form-content {
  :deep(.el-form-item__label) {
    color: #374151;
    font-weight: 500;
  }

  :deep(.el-input__wrapper) {
    border-radius: 10px;
    border: 2px solid #f1f5f9;
    transition: all 0.3s ease;
    box-shadow: none;

    &:hover {
      border-color: #667eea;
      box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
    }

    &.is-focus {
      border-color: #667eea;
      box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.15);
    }
  }

  :deep(.el-select .el-input__wrapper),
  :deep(.el-cascader .el-input__wrapper) {
    &:hover {
      border-color: #667eea;
      box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
    }
  }
}

// 表格卡片
.table-card {
  border-radius: 16px;
  border: none;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  margin-bottom: 24px;
  overflow: hidden;

  :deep(.el-card__body) {
    padding: 0;
  }
}

.table-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 24px;
  background: linear-gradient(135deg, #f8f9ff, #ffffff);
  border-bottom: 1px solid rgba(102, 126, 234, 0.1);
}

.table-actions {
  .batch-btn {
    background: linear-gradient(135deg, #f59e0b, #f97316);
    border: none;
    border-radius: 8px;
    color: white;
    font-weight: 500;
    transition: all 0.3s ease;

    &:hover {
      transform: translateY(-1px);
      box-shadow: 0 4px 15px rgba(245, 158, 11, 0.3);
    }
  }
}

// 表格样式
.product-table {
  :deep(.table-header-row) {
    background: linear-gradient(135deg, #f8f9ff, #ffffff);

    th {
      color: #374151;
      font-weight: 600;
      border-bottom: 2px solid rgba(102, 126, 234, 0.1);
      font-size: 14px;
    }
  }

  :deep(.el-table__row) {
    transition: all 0.3s ease;

    &:hover {
      background: rgba(102, 126, 234, 0.03);
      transform: translateY(-1px);
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
    }
  }

  :deep(.el-table__body-wrapper) {
    padding: 0 24px 24px;
  }
}

.row-number {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  border-radius: 50%;
  font-size: 12px;
  font-weight: 600;
  box-shadow: 0 2px 8px rgba(102, 126, 234, 0.3);
}

// 商品信息
.product-info {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px 0;
}

.product-image {
  flex-shrink: 0;
}

.product-img {
  width: 64px;
  height: 64px;
  border-radius: 12px;
  border: 2px solid #f1f5f9;
  transition: all 0.3s ease;

  &:hover {
    border-color: #667eea;
    transform: scale(1.05);
  }
}

.image-slot {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #f8f9ff, #f1f5f9);
  color: #94a3b8;
  font-size: 16px;
  border-radius: 10px;
}

.product-details {
  flex: 1;
  min-width: 0;
}

.product-name {
  font-size: 15px;
  font-weight: 600;
  color: #1e293b;
  margin: 0 0 6px 0;
  line-height: 1.4;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  transition: color 0.3s ease;

  &:hover {
    color: #667eea;
  }
}

.product-sn {
  font-size: 12px;
  color: #64748b;
  margin: 0;
  background: #f8fafc;
  padding: 2px 8px;
  border-radius: 6px;
  display: inline-block;
}

// 标签样式
.brand-tag {
  background: linear-gradient(135deg, #06b6d4, #0891b2);
  border: none;
  color: white;
  font-weight: 500;
  border-radius: 8px;
  padding: 4px 10px;
}

.category-tag {
  background: linear-gradient(135deg, #f59e0b, #f97316);
  border: none;
  color: white;
  font-weight: 500;
  border-radius: 8px;
  padding: 4px 10px;
}

// 价格信息
.price-info {
  display: flex;
  align-items: baseline;
  justify-content: flex-end;
  padding: 8px 12px;
  background: linear-gradient(135deg, #fee2e2, #fef7ff);
  border-radius: 10px;
  border: 1px solid #fecaca;
}

.price-symbol {
  font-size: 14px;
  color: #dc2626;
  margin-right: 2px;
  font-weight: 500;
}

.price-value {
  font-size: 18px;
  font-weight: 700;
  color: #dc2626;
}

// 状态开关
.status-switches {
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding: 4px;
}

.switch-item {
  display: flex;
  justify-content: center;

  :deep(.el-switch) {
    --el-switch-on-color: #667eea;
    --el-switch-off-color: #e2e8f0;
    --el-switch-border-color: #e2e8f0;

    .el-switch__core {
      border-radius: 12px;
      transition: all 0.3s ease;
    }

    &.is-checked .el-switch__core {
      background: linear-gradient(135deg, #667eea, #764ba2);
      box-shadow: 0 2px 8px rgba(102, 126, 234, 0.3);
    }
  }
}

// 操作按钮
.action-buttons {
  display: flex;
  flex-direction: column;
  gap: 8px;

  .el-button {
    border-radius: 8px;
    font-size: 12px;
    padding: 6px 12px;
    font-weight: 500;
    transition: all 0.3s ease;
  }

  .view-btn {
    background: linear-gradient(135deg, #667eea, #764ba2);
    border: none;

    &:hover {
      transform: translateY(-1px);
      box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
    }
  }

  .delete-btn {
    background: linear-gradient(135deg, #ef4444, #dc2626);
    border: none;

    &:hover {
      transform: translateY(-1px);
      box-shadow: 0 4px 12px rgba(239, 68, 68, 0.4);
    }
  }
}

.text-muted {
  color: #94a3b8;
  font-size: 12px;
  font-style: italic;
}

// 分页
.pagination-container {
  display: flex;
  justify-content: center;
  padding: 20px 0;
}

.custom-pagination {
  background: rgba(255, 255, 255, 0.95);
  padding: 20px 32px;
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);

  :deep(.el-pagination) {
    .el-pager li {
      border-radius: 8px;
      margin: 0 3px;
      transition: all 0.3s ease;
      border: 1px solid transparent;

      &:hover {
        transform: translateY(-2px);
        background: #667eea;
        color: white;
        box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
      }

      &.is-active {
        background: linear-gradient(135deg, #667eea, #764ba2);
        border-color: transparent;
        color: white;
        box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
      }
    }

    .btn-prev,
    .btn-next {
      border-radius: 8px;
      transition: all 0.3s ease;
      border: 1px solid #e2e8f0;

      &:hover {
        transform: translateY(-2px);
        background: #667eea;
        color: white;
        border-color: #667eea;
        box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
      }
    }

    .el-select .el-input {
      .el-input__wrapper {
        border-radius: 8px;
        border: 1px solid #e2e8f0;
        transition: all 0.3s ease;

        &:hover {
          border-color: #667eea;
        }
      }
    }
  }
}

// 响应式设计
@media (max-width: 1200px) {
  .search-header {
    flex-direction: column;
    gap: 16px;
    align-items: stretch;
  }

  .search-actions {
    align-self: flex-start;
    width: 100%;
    justify-content: space-between;
  }

  .search-form-content {
    .el-col {
      margin-bottom: 16px;
    }
  }
}

@media (max-width: 768px) {
  .section-header {
    padding: 20px;
  }

  .page-title {
    font-size: 24px;
  }

  .search-header,
  .table-header {
    flex-direction: column;
    gap: 12px;
    align-items: flex-start;
    padding: 16px 20px;
  }

  .search-actions,
  .table-actions {
    width: 100%;
    justify-content: flex-end;
    flex-wrap: wrap;
  }

  .search-form {
    padding: 20px;
  }

  .search-form-content {
    .el-col {
      width: 100% !important;
      flex: none;
      max-width: 100%;
    }
  }

  .product-info {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }

  .product-table {
    :deep(.el-table__header-wrapper),
    :deep(.el-table__body-wrapper) {
      overflow-x: auto;
    }

    :deep(.el-table__body-wrapper) {
      padding: 0 16px 16px;
    }
  }

  .status-switches {
    flex-direction: row;
    justify-content: space-around;
    flex-wrap: wrap;
  }

  .action-buttons {
    flex-direction: row;
    justify-content: center;
  }

  .custom-pagination {
    padding: 16px 20px;

    :deep(.el-pagination) {
      .el-pagination__total,
      .el-pagination__jump {
        display: none;
      }
    }
  }
}

@media (max-width: 480px) {
  .header-info {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }

  .search-actions {
    flex-direction: column;
    gap: 8px;
    width: 100%;

    .el-button {
      width: 100%;
      justify-content: center;
    }
  }

  .product-img {
    width: 48px;
    height: 48px;
  }

  .product-name {
    font-size: 13px;
  }

  .price-value {
    font-size: 16px;
  }

  .status-switches {
    gap: 4px;
  }

  .action-buttons {
    .el-button {
      padding: 4px 8px;
      font-size: 11px;
    }
  }
}

// 加载动画
.table-container {
  position: relative;
  min-height: 200px;
}

:deep(.el-loading-mask) {
  border-radius: 16px;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
}

// 空状态
.empty-state {
  text-align: center;
  padding: 80px 0;
  color: #64748b;

  .empty-icon {
    font-size: 72px;
    color: #cbd5e1;
    margin-bottom: 20px;
  }

  .empty-text {
    font-size: 18px;
    margin-bottom: 8px;
    font-weight: 500;
  }

  .empty-description {
    font-size: 14px;
    color: #94a3b8;
  }
}

// 动画效果
@keyframes slideInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.search-section {
  animation: slideInUp 0.6s ease;
}

.table-card {
  animation: slideInUp 0.6s ease;
  animation-delay: 0.1s;
  animation-fill-mode: both;
}

.pagination-container {
  animation: slideInUp 0.6s ease;
  animation-delay: 0.2s;
  animation-fill-mode: both;
}

// 悬浮效果
.search-card,
.table-card,
.custom-pagination {
  transition: all 0.3s ease;

  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15);
  }
}

// 毛玻璃效果增强
.search-card,
.table-card,
.custom-pagination {
  position: relative;

  &::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: inherit;
    backdrop-filter: blur(20px);
    z-index: -1;
    border-radius: inherit;
  }
}
</style>
