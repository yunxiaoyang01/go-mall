<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-center: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="center" label="编号" prop="ID" />
        <el-table-column align="center" label="商品图片" prop="pic" width="120">
          <template #default="scope"><img style="height: 80px" :src="scope.row.pic"></template>
        </el-table-column>
        <el-table-column align="center" label="商品名称">
          <template #default="scope">
            <p>{{ scope.row.name }}</p>
            <p>品牌：{{ scope.row.brandName }}</p>
          </template>
        </el-table-column>
        <el-table-column align="center" label="价格/货号">
          <template #default="scope">
            <p>价格：￥{{ scope.row.price }}</p>
            <p>货号：{{ scope.row.productSn }}</p>
          </template>
        </el-table-column>
        <el-table-column label="标签" width="140" align="center">
          <template #default="scope">
            <p>上架：
              <el-switch
                v-model="scope.row.publishStatus"
                :active-value="1"
                :inactive-value="0"
                @change="handlePublishStatusChange(scope.$index, scope.row)"
              />
            </p>
            <p>新品：
              <el-switch
                v-model="scope.row.newStatus"
                :active-value="1"
                :inactive-value="0"
                @change="handleNewStatusChange(scope.$index, scope.row)"
              />
            </p>
            <p>推荐：
              <el-switch
                v-model="scope.row.recommandStatus"
                :active-value="1"
                :inactive-value="0"
                @change="handleRecommendStatusChange(scope.$index, scope.row)"
              />
            </p>
          </template>
        </el-table-column>
        <el-table-column label="排序" width="100" align="center">
          <template #default="scope">{{ scope.row.sort }}</template>
        </el-table-column>
        <el-table-column label="SKU库存" width="100" align="center">
          <template #default="scope">
            <el-button type="primary" icon="edit" circle @click="handleShowSkuEditDialog(scope.$index, scope.row)" />
          </template>
        </el-table-column>
        <el-table-column label="销量" width="100" align="center">
          <template #default="scope">{{ scope.row.sale }}</template>
        </el-table-column>
        <el-table-column label="审核状态" width="100" align="center">
          <template #default="scope">
            <p v-if="scope.row.verifyStatus == 1">审核通过</p>
            <p v-else>审核未通过</p>
            <p>
              <el-button
                type="text"
                @click="handleShowVerifyDetail(scope.$index, scope.row)"
              >审核详情
              </el-button>
            </p>
          </template>
        </el-table-column>
        <el-table-column align="center" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updatePmsProductFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="brandId字段:">
          <el-input v-model.number="formData.brandId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="productCategoryId字段:">
          <el-input v-model.number="formData.productCategoryId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="feightTemplateId字段:">
          <el-input v-model.number="formData.feightTemplateId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="productAttributeCategoryId字段:">
          <el-input v-model.number="formData.productAttributeCategoryId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="name字段:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="pic字段:">
          <el-input v-model="formData.pic" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="货号:">
          <el-input v-model="formData.productSn" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="删除状态：0->未删除；1->已删除:">
          <el-input v-model.number="formData.deleteStatus" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上架状态：0->下架；1->上架:">
          <el-input v-model.number="formData.publishStatus" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="新品状态:0->不是新品；1->新品:">
          <el-input v-model.number="formData.newStatus" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="推荐状态；0->不推荐；1->推荐:">
          <el-input v-model.number="formData.recommandStatus" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核状态：0->未审核；1->审核通过:">
          <el-input v-model.number="formData.verifyStatus" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="排序:">
          <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="销量:">
          <el-input v-model.number="formData.sale" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="price字段:">
          <el-input-number v-model="formData.price" style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="促销价格:">
          <el-input-number v-model="formData.promotionPrice" style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="赠送的成长值:">
          <el-input v-model.number="formData.giftGrowth" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="赠送的积分:">
          <el-input v-model.number="formData.giftPoint" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="限制使用的积分数:">
          <el-input v-model.number="formData.usePointLimit" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="副标题:">
          <el-input v-model="formData.subTitle" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品描述:">
          <el-input v-model="formData.description" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="市场价:">
          <el-input-number v-model="formData.originalPrice" style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="库存:">
          <el-input v-model.number="formData.stock" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="库存预警值:">
          <el-input v-model.number="formData.lowStock" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="单位:">
          <el-input v-model="formData.unit" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品重量，默认为克:">
          <el-input-number v-model="formData.weight" style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="是否为预告商品：0->不是；1->是:">
          <el-input v-model.number="formData.previewStatus" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮:">
          <el-input v-model="formData.serviceIds" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="keywords字段:">
          <el-input v-model="formData.keywords" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="note字段:">
          <el-input v-model="formData.note" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="画册图片，连产品图片限制为5张，以逗号分割:">
          <el-input v-model="formData.albumPics" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="detailTitle字段:">
          <el-input v-model="formData.detailTitle" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="detailDesc字段:">
          <el-input v-model="formData.detailDesc" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="产品详情网页内容:">
          <el-input v-model="formData.detailHtml" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="移动端网页详情:">
          <el-input v-model="formData.detailMobileHtml" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="促销开始时间:">
          <el-date-picker v-model="formData.promotionStartTime" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="促销结束时间:">
          <el-date-picker v-model="formData.promotionEndTime" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="活动限购数量:">
          <el-input v-model.number="formData.promotionPerLimit" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购:">
          <el-input v-model.number="formData.promotionType" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="品牌名称:">
          <el-input v-model="formData.brandName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品分类名称:">
          <el-input v-model="formData.productCategoryName" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'PmsProduct'
}
</script>

<script setup>
import {
  createPmsProduct,
  deletePmsProduct,
  deletePmsProductByIds,
  updatePmsProduct,
  findPmsProduct,
  getPmsProductList
} from '@/api/pmsProduct'
import { fetchList as fetchSkuStockList, update as updateSkuStockList } from '@/api/skuStock'
import { fetchList as fetchProductAttrList } from '@/api/productAttr'
// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  brandId: 0,
  productCategoryId: 0,
  feightTemplateId: 0,
  productAttributeCategoryId: 0,
  name: '',
  pic: '',
  productSn: '',
  deleteStatus: 0,
  publishStatus: 0,
  newStatus: 0,
  recommandStatus: 0,
  verifyStatus: 0,
  sort: 0,
  sale: 0,
  price: 0,
  promotionPrice: 0,
  giftGrowth: 0,
  giftPoint: 0,
  usePointLimit: 0,
  subTitle: '',
  description: '',
  originalPrice: 0,
  stock: 0,
  lowStock: 0,
  unit: '',
  weight: 0,
  previewStatus: 0,
  serviceIds: '',
  keywords: '',
  note: '',
  albumPics: '',
  detailTitle: '',
  detailDesc: '',
  detailHtml: '',
  detailMobileHtml: '',
  promotionStartTime: new Date(),
  promotionEndTime: new Date(),
  promotionPerLimit: 0,
  promotionType: 0,
  brandName: '',
  productCategoryName: '',
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const editSkuInfo = ref({
  dialogVisible: false,
  productId: null,
  productSn: '',
  productAttributeCategoryId: null,
  stockList: [],
  productAttr: [],
  keyword: null
})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getPmsProductList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deletePmsProductFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deletePmsProductByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updatePmsProductFunc = async(row) => {
  const res = await findPmsProduct({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.repmsProduct
    dialogFormVisible.value = true
  }
}

// 删除行
const deletePmsProductFunc = async(row) => {
  const res = await deletePmsProduct({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    brandId: 0,
    productCategoryId: 0,
    feightTemplateId: 0,
    productAttributeCategoryId: 0,
    name: '',
    pic: '',
    productSn: '',
    deleteStatus: 0,
    publishStatus: 0,
    newStatus: 0,
    recommandStatus: 0,
    verifyStatus: 0,
    sort: 0,
    sale: 0,
    price: 0,
    promotionPrice: 0,
    giftGrowth: 0,
    giftPoint: 0,
    usePointLimit: 0,
    subTitle: '',
    description: '',
    originalPrice: 0,
    stock: 0,
    lowStock: 0,
    unit: '',
    weight: 0,
    previewStatus: 0,
    serviceIds: '',
    keywords: '',
    note: '',
    albumPics: '',
    detailTitle: '',
    detailDesc: '',
    detailHtml: '',
    detailMobileHtml: '',
    promotionStartTime: new Date(),
    promotionEndTime: new Date(),
    promotionPerLimit: 0,
    promotionType: 0,
    brandName: '',
    productCategoryName: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createPmsProduct(formData.value)
      break
    case 'update':
      res = await updatePmsProduct(formData.value)
      break
    default:
      res = await createPmsProduct(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    closeDialog()
    getTableData()
  }
}
const handlePublishStatusChange = async(index, row) => {
  const ids = []
  ids.push(row.id)
  this.updatePublishStatus(row.publishStatus, ids)
}

const handleNewStatusChange = async(index, row) => {
  const ids = []
  ids.push(row.id)
  this.updateNewStatus(row.newStatus, ids)
}
const handleRecommendStatusChange = async(index, row) => {
  const ids = []
  ids.push(row.id)
  this.updateRecommendStatus(row.recommandStatus, ids)
}
const handleShowSkuEditDialog = async(index, row) => {
  editSkuInfo.dialogVisible = true
  editSkuInfo.productId = row.id
  editSkuInfo.productSn = row.productSn
  editSkuInfo.productAttributeCategoryId = row.productAttributeCategoryId
  editSkuInfo.keyword = null
  fetchSkuStockList(row.id, { keyword: editSkuInfo.keyword }).then(response => {
    editSkuInfo.stockList = response.data
  })
  if (row.productAttributeCategoryId != null) {
    fetchProductAttrList(row.productAttributeCategoryId, { type: 0 }).then(response => {
      editSkuInfo.productAttr = response.data.list
    })
  }
}
const handleShowVerifyDetail = async(index, row) => {
  console.log('handleShowVerifyDetail', row)
}
</script>

<style>
</style>
