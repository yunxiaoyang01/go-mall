// 自动生成模板PmsProduct
package product

import (
	"github.com/yunxiaoyang01/go-mall/mall-admin/global"
	"time"
)

// 商品信息 结构体
type PmsProduct struct {
      global.GVA_MODEL
      BrandId  *int `json:"brandId" form:"brandId" gorm:"column:brand_id;comment:;size:19;"`
      ProductCategoryId  *int `json:"productCategoryId" form:"productCategoryId" gorm:"column:product_category_id;comment:;size:19;"`
      FeightTemplateId  *int `json:"feightTemplateId" form:"feightTemplateId" gorm:"column:feight_template_id;comment:;size:19;"`
      ProductAttributeCategoryId  *int `json:"productAttributeCategoryId" form:"productAttributeCategoryId" gorm:"column:product_attribute_category_id;comment:;size:19;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:64;"`
      Pic  string `json:"pic" form:"pic" gorm:"column:pic;comment:;size:255;"`
      ProductSn  string `json:"productSn" form:"productSn" gorm:"column:product_sn;comment:货号;size:64;"`
      DeleteStatus  *int `json:"deleteStatus" form:"deleteStatus" gorm:"column:delete_status;comment:删除状态：0->未删除；1->已删除;size:10;"`
      PublishStatus  *int `json:"publishStatus" form:"publishStatus" gorm:"column:publish_status;comment:上架状态：0->下架；1->上架;size:10;"`
      NewStatus  *int `json:"newStatus" form:"newStatus" gorm:"column:new_status;comment:新品状态:0->不是新品；1->新品;size:10;"`
      RecommandStatus  *int `json:"recommandStatus" form:"recommandStatus" gorm:"column:recommand_status;comment:推荐状态；0->不推荐；1->推荐;size:10;"`
      VerifyStatus  *int `json:"verifyStatus" form:"verifyStatus" gorm:"column:verify_status;comment:审核状态：0->未审核；1->审核通过;size:10;"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
      Sale  *int `json:"sale" form:"sale" gorm:"column:sale;comment:销量;size:10;"`
      Price  *float64 `json:"price" form:"price" gorm:"column:price;comment:;size:10;"`
      PromotionPrice  *float64 `json:"promotionPrice" form:"promotionPrice" gorm:"column:promotion_price;comment:促销价格;size:10;"`
      GiftGrowth  *int `json:"giftGrowth" form:"giftGrowth" gorm:"column:gift_growth;comment:赠送的成长值;size:10;"`
      GiftPoint  *int `json:"giftPoint" form:"giftPoint" gorm:"column:gift_point;comment:赠送的积分;size:10;"`
      UsePointLimit  *int `json:"usePointLimit" form:"usePointLimit" gorm:"column:use_point_limit;comment:限制使用的积分数;size:10;"`
      SubTitle  string `json:"subTitle" form:"subTitle" gorm:"column:sub_title;comment:副标题;size:255;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:商品描述;"`
      OriginalPrice  *float64 `json:"originalPrice" form:"originalPrice" gorm:"column:original_price;comment:市场价;size:10;"`
      Stock  *int `json:"stock" form:"stock" gorm:"column:stock;comment:库存;size:10;"`
      LowStock  *int `json:"lowStock" form:"lowStock" gorm:"column:low_stock;comment:库存预警值;size:10;"`
      Unit  string `json:"unit" form:"unit" gorm:"column:unit;comment:单位;size:16;"`
      Weight  *float64 `json:"weight" form:"weight" gorm:"column:weight;comment:商品重量，默认为克;size:10;"`
      PreviewStatus  *int `json:"previewStatus" form:"previewStatus" gorm:"column:preview_status;comment:是否为预告商品：0->不是；1->是;size:10;"`
      ServiceIds  string `json:"serviceIds" form:"serviceIds" gorm:"column:service_ids;comment:以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮;size:64;"`
      Keywords  string `json:"keywords" form:"keywords" gorm:"column:keywords;comment:;size:255;"`
      Note  string `json:"note" form:"note" gorm:"column:note;comment:;size:255;"`
      AlbumPics  string `json:"albumPics" form:"albumPics" gorm:"column:album_pics;comment:画册图片，连产品图片限制为5张，以逗号分割;size:255;"`
      DetailTitle  string `json:"detailTitle" form:"detailTitle" gorm:"column:detail_title;comment:;size:255;"`
      DetailDesc  string `json:"detailDesc" form:"detailDesc" gorm:"column:detail_desc;comment:;"`
      DetailHtml  string `json:"detailHtml" form:"detailHtml" gorm:"column:detail_html;comment:产品详情网页内容;"`
      DetailMobileHtml  string `json:"detailMobileHtml" form:"detailMobileHtml" gorm:"column:detail_mobile_html;comment:移动端网页详情;"`
      PromotionStartTime  *time.Time `json:"promotionStartTime" form:"promotionStartTime" gorm:"column:promotion_start_time;comment:促销开始时间;"`
      PromotionEndTime  *time.Time `json:"promotionEndTime" form:"promotionEndTime" gorm:"column:promotion_end_time;comment:促销结束时间;"`
      PromotionPerLimit  *int `json:"promotionPerLimit" form:"promotionPerLimit" gorm:"column:promotion_per_limit;comment:活动限购数量;size:10;"`
      PromotionType  *int `json:"promotionType" form:"promotionType" gorm:"column:promotion_type;comment:促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购;size:10;"`
      BrandName  string `json:"brandName" form:"brandName" gorm:"column:brand_name;comment:品牌名称;size:255;"`
      ProductCategoryName  string `json:"productCategoryName" form:"productCategoryName" gorm:"column:product_category_name;comment:商品分类名称;size:255;"`
}


// TableName 商品信息 表名
func (PmsProduct) TableName() string {
  return "pms_product"
}

