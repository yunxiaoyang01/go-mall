<template>
  <div>
    <div class="gva-form-box">
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
          <el-input-number v-model="formData.price" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="促销价格:">
          <el-input-number v-model="formData.promotionPrice" :precision="2" clearable></el-input-number>
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
          <el-input-number v-model="formData.originalPrice" :precision="2" clearable></el-input-number>
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
          <el-input-number v-model="formData.weight" :precision="2" clearable></el-input-number>
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
          <el-date-picker v-model="formData.promotionStartTime" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="促销结束时间:">
          <el-date-picker v-model="formData.promotionEndTime" type="date" placeholder="选择日期" clearable></el-date-picker>
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
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
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
  updatePmsProduct,
  findPmsProduct
} from '@/api/pmsProduct'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPmsProduct({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.repmsProduct
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
